// Package winianatz provides timezone mapping between Microsoft Windows timezone
// identifiers and IANA (Internet Assigned Numbers Authority) timezone names.
//
// This package supports conversion from Microsoft's proprietary timezone formats
// to standardized IANA timezone identifiers, enabling proper timezone handling
// in cross-platform applications that need to work with Microsoft data sources.
//
// The package provides multiple ways to look up timezone mappings:
//   - By Microsoft timezone alias (e.g., "Hawaiian Standard Time")
//   - By Microsoft display name (e.g., "(UTC-10:00) Hawaii")
//   - By IANA timezone name (e.g., "Pacific/Honolulu")
//   - With specific territory codes for regional variations
//
// Example usage:
//
//	entry, err := winianatz.FromMicrosoftAlias("Hawaiian Standard Time")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(entry.IANA) // Output: Pacific/Honolulu
//
//	// Find all entries for a timezone across territories
//	entries := winianatz.AllFromMicrosoftAlias("Central Europe Standard Time")
//	for _, entry := range entries {
//		fmt.Printf("%s (%s) -> %s (%s)\n", entry.MicrosoftAlias, entry.MicrosoftDisplayName, entry.IANA, entry.Territory)
//	}
//
// Following data sources are used programmatically for generating and verifying the mappings:
//   - [Unicode CLDR (Common Locale Data Repository)](https://github.com/unicode-org/cldr/blob/main/common/supplemental/windowsZones.xml)
//   - [Microsoft Graph REST API v1.0](https://learn.microsoft.com/en-us/graph/api/outlookuser-supportedtimezones)
//
// Reference documentation:
//   - [Microsoft Time Zone Index](https://docs.microsoft.com/en-us/windows-hardware/manufacture/desktop/default-time-zones)
package winianatz

import "errors"

// Entry represents a timezone mapping between Microsoft Windows timezone
// identifiers and IANA timezone names for a specific territory.
//
// Each entry contains the Microsoft timezone alias, display name, corresponding
// IANA timezone identifier, and the territory code where this mapping applies.
// Territory "001" represents the global/default mapping.
type Entry struct {
	// MicrosoftAlias is the Microsoft Windows timezone identifier.
	// Examples: "Hawaiian Standard Time", "Central Europe Standard Time"
	MicrosoftAlias string

	// MicrosoftDisplayName is the human-readable timezone display name used by Microsoft.
	// Examples: "(UTC-10:00) Hawaii", "(UTC+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague"
	MicrosoftDisplayName string

	// Territory is the ISO 3166 territory code where this mapping applies.
	// "001" represents the global/default mapping, while specific codes like
	// "US", "CA", "CZ", "SK" represent country-specific mappings.
	Territory string

	// IANA is the standard IANA timezone identifier.
	// Examples: "Pacific/Honolulu", "Europe/Prague"
	IANA string
}

// SupportedTZID represents a supported IANA timezone identifier.
type SupportedTZID struct {
	// IANA is the standard IANA timezone identifier.
	IANA string
}

var (
	indexMicrosoftAlias       map[string][]int
	indexMicrosoftDisplayName map[string][]int
	indexIANA                 map[string][]int
)

func init() {
	indexMicrosoftAlias = make(map[string][]int, len(data))
	indexMicrosoftDisplayName = make(map[string][]int, len(data))
	indexIANA = make(map[string][]int, len(data))

	for i, e := range data {
		indexMicrosoftAlias[e.MicrosoftAlias] = append(indexMicrosoftAlias[e.MicrosoftAlias], i)
		indexMicrosoftDisplayName[e.MicrosoftDisplayName] = append(indexMicrosoftDisplayName[e.MicrosoftDisplayName], i)
		indexIANA[e.IANA] = append(indexIANA[e.IANA], i)
	}
}

// ErrNotFound is returned when no timezone mapping is found for the given identifier.
var ErrNotFound = errors.New("no mapping found")

// FromMicrosoftAlias returns the default timezone mapping for the given Microsoft
// timezone alias. It returns the entry with territory "001" (global/default).
//
// Returns ErrNotFound if the alias exists but has no global mapping,
// or if the alias doesn't exist at all.
//
// Example:
//
//	entry, err := FromMicrosoftAlias("Hawaiian Standard Time")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(entry.IANA) // Output: Pacific/Honolulu
func FromMicrosoftAlias(alias string) (Entry, error) {
	return one001(collect(indexMicrosoftAlias[alias]))
}

// FromMicrosoftDisplayName returns the default timezone mapping for the given
// Microsoft timezone display name. It returns the entry with territory "001" (global/default).
//
// Returns ErrNotFound if the display name exists but has no global mapping,
// or if the display name doesn't exist at all.
//
// Example:
//
//	entry, err := FromMicrosoftDisplayName("(UTC-10:00) Hawaii")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(entry.IANA) // Output: Pacific/Honolulu
func FromMicrosoftDisplayName(name string) (Entry, error) {
	return one001(collect(indexMicrosoftDisplayName[name]))
}

// FromIANA returns the timezone mapping for the given IANA timezone identifier.
// It preferentially returns the entry with territory "001" (global/default) if available,
// otherwise returns the first available entry for any territory.
//
// Returns ErrNotFound if the IANA identifier doesn't exist at all.
//
// Example:
//
//	entry, err := FromIANA("Pacific/Honolulu")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(entry.MicrosoftAlias) // Output: Hawaiian Standard Time
func FromIANA(iana string) (Entry, error) {
	entries := collect(indexIANA[iana])
	if len(entries) == 0 {
		return Entry{}, ErrNotFound
	}

	// Try to find the default territory (001) first
	for _, e := range entries {
		if e.Territory == "001" {
			return e, nil
		}
	}

	// If no default territory exists, return the first available entry
	return entries[0], nil
}

// FromMicrosoftAliasWithTerritory returns the timezone mapping for the given
// Microsoft timezone alias and specific territory code.
//
// Returns ErrNotFound if no mapping exists for the given alias and territory combination.
//
// Example:
//
//	// Get Czech-specific mapping
//	entry, err := FromMicrosoftAliasWithTerritory("Central Europe Standard Time", "CZ")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(entry.IANA) // Output: Europe/Prague
func FromMicrosoftAliasWithTerritory(alias, territory string) (Entry, error) {
	entries := filterTerritory(collect(indexMicrosoftAlias[alias]), territory)
	if len(entries) == 0 {
		return Entry{}, ErrNotFound
	}
	return entries[0], nil
}

// FromMicrosoftDisplayNameWithTerritory returns the timezone mapping for the given
// Microsoft timezone display name and specific territory code.
//
// Returns ErrNotFound if no mapping exists for the given display name and territory combination.
//
// Example:
//
//	// Get US-specific mapping
//	entry, err := FromMicrosoftDisplayNameWithTerritory("(UTC-05:00) Eastern Time (US & Canada)", "US")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(entry.IANA) // Output: America/New_York
func FromMicrosoftDisplayNameWithTerritory(name, territory string) (Entry, error) {
	entries := filterTerritory(collect(indexMicrosoftDisplayName[name]), territory)
	if len(entries) == 0 {
		return Entry{}, ErrNotFound
	}
	return entries[0], nil
}

// FromIANAWithTerritory returns the timezone mapping for the given IANA timezone
// identifier and specific territory code.
//
// Returns ErrNotFound if no mapping exists for the given IANA identifier and territory combination.
//
// Example:
//
//	// Get Canadian-specific mapping
//	entry, err := FromIANAWithTerritory("America/Toronto", "CA")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(entry.MicrosoftAlias) // Output: Eastern Standard Time
func FromIANAWithTerritory(iana, territory string) (Entry, error) {
	entries := filterTerritory(collect(indexIANA[iana]), territory)
	if len(entries) == 0 {
		return Entry{}, ErrNotFound
	}
	return entries[0], nil
}

// AllFromMicrosoftAlias returns all timezone mappings for the given Microsoft
// timezone alias across all territories.
//
// Returns an empty slice if no mappings exist for the given alias.
// The returned entries may include different IANA timezones for different territories.
//
// Example:
//
//	entries := AllFromMicrosoftAlias("Central Europe Standard Time")
//	for _, entry := range entries {
//		fmt.Printf("%s -> %s (%s)\n", entry.MicrosoftAlias, entry.IANA, entry.Territory)
//	}
//	// Output might include:
//	// Central Europe Standard Time -> Europe/Budapest (001)
//	// Central Europe Standard Time -> Europe/Prague (CZ)
//	// Central Europe Standard Time -> Europe/Bratislava (SK)
func AllFromMicrosoftAlias(alias string) []Entry {
	return collect(indexMicrosoftAlias[alias])
}

// AllFromMicrosoftDisplayName returns all timezone mappings for the given Microsoft
// timezone display name across all territories.
//
// Returns an empty slice if no mappings exist for the given display name.
// The returned entries may include different IANA timezones for different territories.
//
// Example:
//
//	entries := AllFromMicrosoftDisplayName("(UTC+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague")
//	for _, entry := range entries {
//		fmt.Printf("%s -> %s (%s)\n", entry.MicrosoftDisplayName, entry.IANA, entry.Territory)
//	}
func AllFromMicrosoftDisplayName(display string) []Entry {
	return collect(indexMicrosoftDisplayName[display])
}

// AllFromIANA returns all timezone mappings for the given IANA timezone identifier
// across all territories.
//
// Returns an empty slice if no mappings exist for the given IANA identifier.
// The returned entries typically have the same Microsoft identifiers across territories.
//
// Example:
//
//	entries := AllFromIANA("Europe/Prague")
//	for _, entry := range entries {
//		fmt.Printf("%s -> %s (%s)\n", entry.IANA, entry.MicrosoftAlias, entry.Territory)
//	}
func AllFromIANA(iana string) []Entry {
	return collect(indexIANA[iana])
}

// GetSupportedTZIDs returns all supported IANA timezone identifiers.
// The returned slice is sorted alphabetically by IANA timezone name.
//
// Each SupportedTZID contains the IANA timezone identifier.
//
// Example:
//
//	tzids := GetSupportedTZIDs()
//	for _, tzid := range tzids {
//		fmt.Printf("%s\n", tzid.IANA)
//	}
func GetSupportedTZIDs() []SupportedTZID {
	// Return a copy to prevent external modification
	result := make([]SupportedTZID, len(supportedTZIDs))
	copy(result, supportedTZIDs)
	return result
}

func collect(idx []int) []Entry {
	out := make([]Entry, 0, len(idx))
	for _, i := range idx {
		out = append(out, data[i])
	}
	return out
}

func filterTerritory(es []Entry, terr string) []Entry {
	out := es[:0]
	for _, e := range es {
		if e.Territory == terr {
			out = append(out, e)
		}
	}
	// ensure we don't alias the original slice capacity unexpectedly
	return append([]Entry(nil), out...)
}

func one001(es []Entry) (Entry, error) {
	for _, e := range es {
		if e.Territory == "001" {
			return e, nil
		}
	}
	return Entry{}, ErrNotFound
}
