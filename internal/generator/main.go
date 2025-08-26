package main

import (
	"bytes"
	"encoding/xml"
	"go/format"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

const (
	cldrURL = "https://raw.githubusercontent.com/unicode-org/cldr/main/common/supplemental/windowsZones.xml"
	xmlFile = "windowsZones.xml"
	goFile  = "windowsiana.go"
)

type SupplementalData struct {
	XMLName      xml.Name     `xml:"supplementalData"`
	WindowsZones WindowsZones `xml:"windowsZones"`
}

type WindowsZones struct {
	MapTimezones MapTimezones `xml:"mapTimezones"`
}

type MapTimezones struct {
	MapZone []MapZone `xml:"mapZone"`
}

type MapZone struct {
	Other     string `xml:"other,attr"`
	Territory string `xml:"territory,attr"`
	Type      string `xml:"type,attr"`
	Comment   string `xml:",comment"`
}

type XMLNode struct {
	XMLName  xml.Name
	Content  string    `xml:",chardata"`
	Comment  string    `xml:",comment"`
	Children []XMLNode `xml:",any"`
}

const fileTemplate = `// Package winianatz provides support in converting date timezones from
// the non-standard windows timezone format into IANA format.
//
// Originally written by Gary Barnett (2018), released into the Public Domain.
// Modified by Artem Danilov (2025). This modified version is released under
// the MIT License. See LICENSE for details.
package winianatz

import (
	"errors"
	"time"
)

// WinIANA is a pseudo constant that provides a mapping between the Windows timezones and the IANA zones
// Source: https://github.com/unicode-org/cldr/blob/main/common/supplemental/windowsZones.xml
var WinIANA = map[string]string{
{{range .WinIANA}}	"{{.Key}}": "{{.Value}}",
{{end}}}

// MicrosoftIANA is a mapping between Microsoft timezone identifiers and IANA zones
// Source: https://github.com/unicode-org/cldr/blob/main/common/supplemental/windowsZones.xml
var MicrosoftIANA = map[string]string{
{{range .MicrosoftIANA}}	"{{.Key}}": "{{.Value}}",
{{end}}}

// TimezoneParseWindows accepts a timestring in the format "2006-01-02T15:04:05" as the tstring
// parameter and a windows time zone (eg "(UTC+12:00) Fiji") as the timezone. It will return
// a timezoned date, which will correctly handle daylight savings time if it's in force at the given date
func TimezoneParseWindows(tstring string, tzone string) (time.Time, error) {
	ianazone := WinIANA[tzone]
	if ianazone == "" {
		var t time.Time
		return t, errors.New("Could not match windows timezone to IANA timezone")
	}
	return TimezoneParseIANA(tstring, WinIANA[tzone])
}

// TimezoneParseIANA accepts a timestring in the format "2006-01-02T15:04:05" as the tstring
// parameter and am IANA time zone (eg  "Pacific/Fiji") as the timezone. It will return
// a timezoned date, which will correctly handle daylight savings time if it's in force at the given date
func TimezoneParseIANA(tstring string, tzone string) (time.Time, error) {
	var zulutime time.Time
	it, err := time.Parse("2006-01-02T15:04:05", tstring)

	if err != nil {
		return zulutime, err
	}
	loc, err := time.LoadLocation(tzone)
	if err != nil {
		return zulutime, err
	}
	zulutime = time.Date(it.Year(), it.Month(), it.Day(), it.Hour(), it.Minute(), it.Second(), 0, loc)
	return zulutime, nil
}

// StripTimezoneFromDate removes the same point in time and date, but strips apart the timezone
func StripTimezoneFromDate(indate time.Time) time.Time {
	return time.Unix(indate.Unix(), 0)
}
`

func main() {
	log.Println("Downloading windowsZones.xml from CLDR...")
	if err := downloadFile(cldrURL, xmlFile); err != nil {
		log.Fatalf("Failed to download XML: %v", err)
	}

	log.Println("Parsing XML...")
	mappings, microsoftMappings, err := parseXML(xmlFile)
	if err != nil {
		log.Fatalf("Failed to parse XML: %v", err)
	}

	log.Println("Generating windowsiana.go...")
	if err := generateGoFile(mappings, microsoftMappings); err != nil {
		log.Fatalf("Failed to generate Go file: %v", err)
	}

	log.Println("Successfully generated windowsiana.go")
}

func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func parseXML(filepath string) (map[string]string, map[string]string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, nil, err
	}

	mappings := make(map[string]string)
	microsoftMappings := make(map[string]string)

	decoder := xml.NewDecoder(bytes.NewReader(data))
	var currentComment string

	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}

		switch t := token.(type) {
		case xml.Comment:
			comment := strings.TrimSpace(string(t))
			if strings.HasPrefix(comment, "(UTC") {
				currentComment = comment
			}
		case xml.StartElement:
			if t.Name.Local == "mapZone" {
				var mapZone MapZone
				if err := decoder.DecodeElement(&mapZone, &t); err != nil {
					return nil, nil, err
				}

				// Only process entries with territory="001" (default/primary mapping)
				if mapZone.Territory == "001" && currentComment != "" {
					ianaZone := mapZone.Type
					// Use the first timezone if multiple are specified (space-separated)
					if strings.Contains(ianaZone, " ") {
						ianaZone = strings.Fields(ianaZone)[0]
					}
					mappings[currentComment] = ianaZone
					microsoftMappings[mapZone.Other] = ianaZone
					currentComment = ""
				}
			}
		}
	}

	return mappings, microsoftMappings, nil
}

type KeyValue struct {
	Key   string
	Value string
}

func generateGoFile(mappings map[string]string, microsoftMappings map[string]string) error {
	var sortedKVs []KeyValue
	for k, v := range mappings {
		sortedKVs = append(sortedKVs, KeyValue{Key: k, Value: v})
	}

	// Helper to extract offset as minutes from key string using regexp
	re := regexp.MustCompile(`\(UTC([+-]\d{2}:\d{2})\)`)
	parseOffset := func(key string) int {
		// Handle exactly "(UTC) ..." as offset 0
		if strings.HasPrefix(key, "(UTC)") {
			return 0
		}
		matches := re.FindStringSubmatch(key)
		if len(matches) != 2 {
			return 0
		}
		offset := matches[1] // e.g. +05:30 or -08:00
		sign := 1
		if offset[0] == '-' {
			sign = -1
		}
		parts := strings.Split(offset[1:], ":")
		if len(parts) != 2 {
			return 0
		}
		hours, err1 := strconv.Atoi(parts[0])
		minutes, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return 0
		}
		return sign * (hours*60 + minutes)
	}

	// Sort from max negative offset to max positive offset, then alphabetically for same offset
	sort.SliceStable(sortedKVs, func(i, j int) bool {
		oi := parseOffset(sortedKVs[i].Key)
		oj := parseOffset(sortedKVs[j].Key)
		if oi == oj {
			return sortedKVs[i].Key < sortedKVs[j].Key
		}
		return oi < oj // lower (more negative) comes first
	})

	// Sort Microsoft mappings alphabetically by key
	var sortedMsKVs []KeyValue
	for k, v := range microsoftMappings {
		sortedMsKVs = append(sortedMsKVs, KeyValue{Key: k, Value: v})
	}
	sort.SliceStable(sortedMsKVs, func(i, j int) bool {
		return sortedMsKVs[i].Key < sortedMsKVs[j].Key
	})

	templateData := struct {
		WinIANA       []KeyValue
		MicrosoftIANA []KeyValue
	}{
		WinIANA:       sortedKVs,
		MicrosoftIANA: sortedMsKVs,
	}

	tmpl, err := template.New("gofile").Parse(fileTemplate)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, templateData); err != nil {
		return err
	}

	// Format the generated code using go/format
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	return os.WriteFile(goFile, formatted, 0644)
}
