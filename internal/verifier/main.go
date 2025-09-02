package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"text/template"

	winianatz "github.com/thommeo/winianatz"
)

type TimeZoneInfo struct {
	Alias       string `json:"alias"`
	DisplayName string `json:"displayName"`
}

type MSGraphResponse struct {
	Context string         `json:"@odata.context"`
	Value   []TimeZoneInfo `json:"value"`
}

type TZIDEntry struct {
	Timezone  string `json:"timezone"`
	Territory string `json:"territory"`
}

type TimeZoneResult struct {
	Alias         string
	DisplayName   string
	AliasResult   string
	DisplayResult string
}

type TZIDResult struct {
	IANA      string
	Territory string
	Result    string
}

type ReportData struct {
	TotalCount          int
	TimeZones           []TimeZoneResult
	MissingAliases      []string
	MissingDisplayNames []string
	TZIDTotal           int
	TZIDResults         []TZIDResult
	TZIDFound           int
	TZIDMissing         int
	MissingTZIDs        []string
}

const reportTemplate = `# Timezone conversion verification report

## Microsoft Graph API Timezone Verification

This is a cross check results between [Microsoft Graph API supported timezones](./references/msgraph-supported-timezones-windows.json) and IANA timezones.

The conversion is implemented based on [Unicode CLDR project's windowsZones.xml](https://github.com/unicode-org/cldr/blob/main/common/supplemental/windowsZones.xml).

| Microsoft Alias | Microsoft Display Name |
|-----------------|------------------------|
{{range .TimeZones}}| {{.Alias}}<br>{{.AliasResult}} | {{.DisplayName}}<br>{{.DisplayResult}} |
{{end}}

## IANA Timezone ID Verification

This section verifies all IANA timezone IDs from [CLDR TZID.txt](./references/TZID.txt) against the winianatz FromIANA function.

**Summary:**
- Total IANA timezones tested: {{.TZIDTotal}}
- Found in winianatz: {{.TZIDFound}}
- Missing from winianatz: {{.TZIDMissing}}
- Coverage: {{printf "%.1f%%" (div (mul (float64 .TZIDFound) 100.0) (float64 .TZIDTotal))}}

### All IANA Timezone Results

| IANA Timezone | Territory | Result |
|---------------|-----------|--------|
{{range .TZIDResults}}| {{.IANA}} | {{.Territory}} | {{.Result}} |
{{end}}

## Missing Microsoft Aliases

The following Microsoft Aliases are reported by Microsoft Graph API as supported but are missing in the current conversion implementation.

{{if .MissingAliases}}{{range .MissingAliases}}- {{.}}
{{end}}{{else}}None
{{end}}

## Missing Microsoft Display Names

The following Microsoft Display Names are reported by Microsoft Graph API as supported but are missing in the current conversion implementation.

{{if .MissingDisplayNames}}{{range .MissingDisplayNames}}- {{.}}
{{end}}{{else}}None
{{end}}

## Missing IANA Timezones

The following IANA timezones from CLDR are not found in winianatz{{if gt (len .MissingTZIDs) 100}} (first 100 of {{len .MissingTZIDs}}){{end}}:

{{if .MissingTZIDs}}{{range $i, $tz := .MissingTZIDs}}{{if lt $i 100}}- {{$tz}}
{{end}}{{end}}{{else}}None
{{end}}`

const (
	reportFileName = "VERIFY.md"
	jsonFile       = "references/msgraph-supported-timezones-windows.json"
	tzidFile       = "references/TZID.json"
)

func loadMSGraphData() ([]TimeZoneInfo, error) {
	data, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("error reading JSON file: %v", err)
	}

	var msGraphData MSGraphResponse
	err = json.Unmarshal(data, &msGraphData)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return msGraphData.Value, nil
}

func loadTZIDData() ([]TZIDEntry, error) {
	tzidData, err := os.ReadFile(tzidFile)
	if err != nil {
		return nil, fmt.Errorf("error reading TZID JSON file: %v", err)
	}

	var tzidEntries []TZIDEntry
	err = json.Unmarshal(tzidData, &tzidEntries)
	if err != nil {
		return nil, fmt.Errorf("error parsing TZID JSON: %v", err)
	}

	return tzidEntries, nil
}

func processMSGraphTimezones(timezones []TimeZoneInfo) ([]TimeZoneResult, []string, []string) {
	results := make([]TimeZoneResult, 0, len(timezones))
	var missingAliases []string
	var missingDisplayNames []string

	for _, tz := range timezones {
		result := TimeZoneResult{
			Alias:       tz.Alias,
			DisplayName: tz.DisplayName,
		}

		// Check alias mapping
		if entry, err := winianatz.FromMicrosoftAlias(tz.Alias); err == nil {
			result.AliasResult = fmt.Sprintf("✅ %s (%s)", entry.IANA, entry.Territory)
		} else {
			result.AliasResult = "❌ IANA Not found"
			missingAliases = append(missingAliases, tz.Alias)
		}

		// Check display name mapping
		if entry, err := winianatz.FromMicrosoftDisplayName(tz.DisplayName); err == nil {
			result.DisplayResult = fmt.Sprintf("✅ %s (%s)", entry.IANA, entry.Territory)
		} else {
			result.DisplayResult = "❌ IANA Not found"
			missingDisplayNames = append(missingDisplayNames, tz.DisplayName)
		}

		results = append(results, result)
	}

	return results, missingAliases, missingDisplayNames
}

func processTZIDEntries(tzidEntries []TZIDEntry) ([]TZIDResult, int, int, []string) {
	results := make([]TZIDResult, 0, len(tzidEntries))
	var found, missing int
	var missingTZIDs []string

	for _, tzid := range tzidEntries {
		result := TZIDResult{
			IANA:      tzid.Timezone,
			Territory: tzid.Territory,
		}

		// Check if IANA timezone is found in winianatz
		if entry, err := winianatz.FromIANA(tzid.Timezone); err == nil {
			result.Result = fmt.Sprintf("✅ %s", entry.MicrosoftAlias)
			found++
		} else {
			result.Result = "❌ Not found"
			missing++
			missingTZIDs = append(missingTZIDs, tzid.Timezone)
		}

		results = append(results, result)
	}

	// Sort results by IANA timezone name for better readability
	sort.Slice(results, func(i, j int) bool {
		return results[i].IANA < results[j].IANA
	})

	// Sort missing TZIDs for better readability
	sort.Strings(missingTZIDs)

	return results, found, missing, missingTZIDs
}

func createTemplateWithFunctions() *template.Template {
	return template.New("report").Funcs(template.FuncMap{
		"div":     func(a, b float64) float64 { return a / b },
		"mul":     func(a, b float64) float64 { return a * b },
		"sub":     func(a, b int) int { return a - b },
		"add":     func(a, b int) int { return a + b },
		"lt":      func(a, b int) bool { return a < b },
		"float64": func(i int) float64 { return float64(i) },
	})
}

func generateReport(reportData ReportData) error {
	tmpl := createTemplateWithFunctions()

	tmpl, err := tmpl.Parse(reportTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	// Create output file
	file, err := os.Create(reportFileName)
	if err != nil {
		return fmt.Errorf("error creating report file: %v", err)
	}
	defer file.Close()

	// Execute template
	err = tmpl.Execute(file, reportData)
	if err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	return nil
}

func printSummary(reportData ReportData) {
	fmt.Printf("%s generated successfully\n", reportFileName)
	fmt.Printf("Microsoft Graph API: %d timezones processed\n", reportData.TotalCount)
	fmt.Printf("TZID verification: %d/%d found (%.1f%% coverage)\n",
		reportData.TZIDFound, reportData.TZIDTotal,
		float64(reportData.TZIDFound)/float64(reportData.TZIDTotal)*100)
}

func main() {
	// Load data
	msGraphTimezones, err := loadMSGraphData()
	if err != nil {
		log.Fatal(err)
	}

	tzidEntries, err := loadTZIDData()
	if err != nil {
		log.Fatal(err)
	}

	// Process Microsoft Graph API timezones
	timeZoneResults, missingAliases, missingDisplayNames := processMSGraphTimezones(msGraphTimezones)

	// Process TZID entries
	tzidResults, tzidFound, tzidMissing, missingTZIDs := processTZIDEntries(tzidEntries)

	// Prepare report data
	reportData := ReportData{
		TotalCount:          len(msGraphTimezones),
		TimeZones:           timeZoneResults,
		MissingAliases:      missingAliases,
		MissingDisplayNames: missingDisplayNames,
		TZIDTotal:           len(tzidEntries),
		TZIDResults:         tzidResults,
		TZIDFound:           tzidFound,
		TZIDMissing:         tzidMissing,
		MissingTZIDs:        missingTZIDs,
	}

	// Generate and save report
	if err := generateReport(reportData); err != nil {
		log.Fatal(err)
	}

	// Print summary
	printSummary(reportData)
}
