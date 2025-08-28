package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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

type TimeZoneResult struct {
	Alias         string
	DisplayName   string
	AliasResult   string
	DisplayResult string
}

type ReportData struct {
	TotalCount          int
	TimeZones           []TimeZoneResult
	MissingAliases      []string
	MissingDisplayNames []string
}

const reportTemplate = `# Timezone conversion verification report

This is a cross check results between [Microsoft Graph API supported timezones](./references/msgraph-supported-timezones-windows.json) and IANA timezones.

The conversion is implemented based on [Unicode CLDR project's windowsZones.xml](https://github.com/unicode-org/cldr/blob/main/common/supplemental/windowsZones.xml).

| Microsoft Alias | Microsoft Display Name |
|-----------------|------------------------|
{{range .TimeZones}}| {{.Alias}}<br>{{.AliasResult}} | {{.DisplayName}}<br>{{.DisplayResult}} |
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
{{end}}`

const (
	reportFileName = "VERIFY.md"
	jsonFile       = "references/msgraph-supported-timezones-windows.json"
)

func main() {
	data, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	var msGraphData MSGraphResponse
	err = json.Unmarshal(data, &msGraphData)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	// Prepare template data
	reportData := ReportData{
		TotalCount: len(msGraphData.Value),
		TimeZones:  make([]TimeZoneResult, 0, len(msGraphData.Value)),
	}

	// Process each timezone
	for _, tz := range msGraphData.Value {
		result := TimeZoneResult{
			Alias:       tz.Alias,
			DisplayName: tz.DisplayName,
		}

		// Check alias mapping
		if entry, err := winianatz.FromMicrosoftAlias(tz.Alias); err == nil {
			result.AliasResult = fmt.Sprintf("✅ %s (%s)", entry.IANA, entry.Territory)
		} else {
			result.AliasResult = "❌ IANA Not found"
			reportData.MissingAliases = append(reportData.MissingAliases, tz.Alias)
		}

		// Check display name mapping
		if entry, err := winianatz.FromMicrosoftDisplayName(tz.DisplayName); err == nil {
			result.DisplayResult = fmt.Sprintf("✅ %s (%s)", entry.IANA, entry.Territory)
		} else {
			result.DisplayResult = "❌ IANA Not found"
			reportData.MissingDisplayNames = append(reportData.MissingDisplayNames, tz.DisplayName)
		}

		reportData.TimeZones = append(reportData.TimeZones, result)
	}

	// Parse and execute template
	tmpl, err := template.New("report").Parse(reportTemplate)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Create output file
	file, err := os.Create(reportFileName)
	if err != nil {
		log.Fatalf("Error creating report file: %v", err)
	}
	defer file.Close()

	// Execute template
	err = tmpl.Execute(file, reportData)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	fmt.Printf("%s generated successfully\n", reportFileName)
}
