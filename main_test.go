package winianatz

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test public API methods with documentation examples

func TestFromMicrosoftAlias(t *testing.T) {
	entry, err := FromMicrosoftAlias("Hawaiian Standard Time")
	require.NoError(t, err)
	assert.Equal(t, "Pacific/Honolulu", entry.IANA)
	assert.Equal(t, "001", entry.Territory)
	assert.Equal(t, "Hawaiian Standard Time", entry.MicrosoftAlias)

	// Test non-existent alias
	_, err = FromMicrosoftAlias("Non Existent Time Zone")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNotFound)
}

func TestFromMicrosoftDisplayName(t *testing.T) {
	entry, err := FromMicrosoftDisplayName("(UTC-10:00) Hawaii")
	require.NoError(t, err)
	assert.Equal(t, "Pacific/Honolulu", entry.IANA)
	assert.Equal(t, "001", entry.Territory)

	// Test non-existent display name
	_, err = FromMicrosoftDisplayName("(UTC-99:00) Non Existent")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNotFound)
}

func TestFromIANA(t *testing.T) {
	entry, err := FromIANA("Pacific/Honolulu")
	require.NoError(t, err)
	assert.Equal(t, "Hawaiian Standard Time", entry.MicrosoftAlias)
	assert.Equal(t, "001", entry.Territory)

	// Test non-existent IANA
	_, err = FromIANA("Non/Existent")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNotFound)
}

func TestFromMicrosoftAliasWithTerritory(t *testing.T) {
	entry, err := FromMicrosoftAliasWithTerritory("Central Europe Standard Time", "CZ")
	require.NoError(t, err)
	assert.Equal(t, "Europe/Prague", entry.IANA)
	assert.Equal(t, "CZ", entry.Territory)

	// Test global mapping
	entry, err = FromMicrosoftAliasWithTerritory("Hawaiian Standard Time", "001")
	require.NoError(t, err)
	assert.Equal(t, "Pacific/Honolulu", entry.IANA)

	// Test non-existent territory
	_, err = FromMicrosoftAliasWithTerritory("Hawaiian Standard Time", "XX")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNotFound)
}

func TestFromMicrosoftDisplayNameWithTerritory(t *testing.T) {
	// Test US-specific mapping
	entry, err := FromMicrosoftDisplayNameWithTerritory("(UTC-05:00) Eastern Time (US & Canada)", "US")
	require.NoError(t, err)
	assert.Equal(t, "America/New_York", entry.IANA)
	assert.Equal(t, "US", entry.Territory)

	// Test Canadian mapping for same display name
	entry, err = FromMicrosoftDisplayNameWithTerritory("(UTC-05:00) Eastern Time (US & Canada)", "CA")
	require.NoError(t, err)
	assert.Equal(t, "America/Toronto", entry.IANA)
	assert.Equal(t, "CA", entry.Territory)
}

func TestFromIANAWithTerritory(t *testing.T) {
	// Test Canadian-specific mapping
	entry, err := FromIANAWithTerritory("America/Toronto", "CA")
	require.NoError(t, err)
	assert.Equal(t, "Eastern Standard Time", entry.MicrosoftAlias)
	assert.Equal(t, "CA", entry.Territory)

	// Test non-existent combination
	_, err = FromIANAWithTerritory("Pacific/Honolulu", "XX")
	assert.Error(t, err)
	assert.ErrorIs(t, err, ErrNotFound)
}

func TestAllFromMicrosoftAlias(t *testing.T) {
	entries := AllFromMicrosoftAlias("Central Europe Standard Time")

	require.NotEmpty(t, entries, "Central Europe Standard Time should have multiple entries")

	territories := make(map[string]bool)
	ianaZones := make(map[string]bool)

	for _, entry := range entries {
		assert.Equal(t, "Central Europe Standard Time", entry.MicrosoftAlias, "All entries should have consistent Microsoft alias")
		territories[entry.Territory] = true
		ianaZones[entry.IANA] = true
	}

	expectedTerritories := []string{"001", "CZ", "SK", "HU"}
	for _, territory := range expectedTerritories {
		assert.True(t, territories[territory], "Territory '%s' should be present in Central Europe entries", territory)
	}

	expectedIANA := []string{"Europe/Budapest", "Europe/Prague", "Europe/Bratislava"}
	for _, iana := range expectedIANA {
		assert.True(t, ianaZones[iana], "IANA zone '%s' should be present in Central Europe entries", iana)
	}

	entries = AllFromMicrosoftAlias("Non Existent Time Zone")
	assert.Empty(t, entries, "Non-existent alias should return empty slice")
}

func TestAllFromMicrosoftDisplayName(t *testing.T) {
	entries := AllFromMicrosoftDisplayName("(UTC+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague")

	require.NotEmpty(t, entries, "Belgrade display name should have multiple entries")

	for _, entry := range entries {
		assert.Equal(t, "(UTC+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague", entry.MicrosoftDisplayName, "All entries should have consistent display name")
		assert.Equal(t, "Central Europe Standard Time", entry.MicrosoftAlias, "All entries should map to Central Europe Standard Time alias")
	}

	entries = AllFromMicrosoftDisplayName("(UTC-99:00) Non Existent")
	assert.Empty(t, entries, "Non-existent display name should return empty slice")
}

func TestAllFromIANA(t *testing.T) {
	entries := AllFromIANA("Europe/Prague")

	require.NotEmpty(t, entries, "Europe/Prague should have entries")

	for _, entry := range entries {
		assert.Equal(t, "Europe/Prague", entry.IANA, "All entries should have consistent IANA identifier")
		assert.Equal(t, "Central Europe Standard Time", entry.MicrosoftAlias, "All entries should map to Central Europe Standard Time alias")
	}

	entries = AllFromIANA("Non/Existent")
	assert.Empty(t, entries, "Non-existent IANA timezone should return empty slice")
}

func TestGetSupportedTZIDs(t *testing.T) {
	tzids := GetSupportedTZIDs()

	require.NotEmpty(t, tzids, "GetSupportedTZIDs should return non-empty slice")

	// Check that the slice is sorted alphabetically
	for i := 1; i < len(tzids); i++ {
		assert.True(t, tzids[i-1].IANA < tzids[i].IANA, "TZIDs should be sorted alphabetically")
	}

	// Check that all entries have IANA field
	for _, tzid := range tzids {
		assert.NotEmpty(t, tzid.IANA, "IANA field should not be empty")
	}

	found := false
	for _, tzid := range tzids {
		if tzid.IANA == "Pacific/Honolulu" {
			found = true
			break
		}
	}
	assert.True(t, found, "Pacific/Honolulu should be in supported TZIDs")
}

// Test private helper functions

func TestCollect(t *testing.T) {
	indices := []int{0, 1, 2}
	entries := collect(indices)
	assert.Len(t, entries, 3, "Should collect exactly 3 entries for 3 indices")

	entries = collect([]int{})
	assert.Empty(t, entries, "Should return empty slice for empty indices")

	entries = collect([]int{0})
	assert.Len(t, entries, 1, "Should collect exactly 1 entry for single index")
}

func TestFilterTerritory(t *testing.T) {
	testEntries1 := []Entry{
		{MicrosoftAlias: "Test Time", Territory: "001", IANA: "Test/Zone1"},
		{MicrosoftAlias: "Test Time", Territory: "US", IANA: "Test/Zone2"},
		{MicrosoftAlias: "Test Time", Territory: "CA", IANA: "Test/Zone3"},
		{MicrosoftAlias: "Test Time", Territory: "US", IANA: "Test/Zone4"},
	}

	filtered := filterTerritory(testEntries1, "US")
	assert.Len(t, filtered, 2, "Should filter exactly 2 US territory entries")

	for _, entry := range filtered {
		assert.Equal(t, "US", entry.Territory, "Filtered entries should only contain US territory")
	}

	testEntries2 := []Entry{
		{MicrosoftAlias: "Test Time", Territory: "001", IANA: "Test/Zone1"},
		{MicrosoftAlias: "Test Time", Territory: "US", IANA: "Test/Zone2"},
		{MicrosoftAlias: "Test Time", Territory: "CA", IANA: "Test/Zone3"},
		{MicrosoftAlias: "Test Time", Territory: "US", IANA: "Test/Zone4"},
	}

	filtered = filterTerritory(testEntries2, "XX")
	assert.Empty(t, filtered, "Non-existent territory should return empty slice")

	testEntries3 := []Entry{
		{MicrosoftAlias: "Test Time", Territory: "001", IANA: "Test/Zone1"},
		{MicrosoftAlias: "Test Time", Territory: "US", IANA: "Test/Zone2"},
		{MicrosoftAlias: "Test Time", Territory: "CA", IANA: "Test/Zone3"},
		{MicrosoftAlias: "Test Time", Territory: "US", IANA: "Test/Zone4"},
	}

	filtered = filterTerritory(testEntries3, "001")
	require.Len(t, filtered, 1, "Should filter exactly 1 global territory entry")
	assert.Equal(t, "001", filtered[0].Territory, "Filtered entry should have correct territory")

	filtered = filterTerritory([]Entry{}, "001")
	assert.Empty(t, filtered, "Empty input should return empty slice")
}

func TestOne001(t *testing.T) {
	testEntries := []Entry{
		{MicrosoftAlias: "Test Time", Territory: "US", IANA: "Test/Zone1"},
		{MicrosoftAlias: "Test Time", Territory: "001", IANA: "Test/Zone2"},
		{MicrosoftAlias: "Test Time", Territory: "CA", IANA: "Test/Zone3"},
	}

	entry, err := one001(testEntries)
	require.NoError(t, err, "Should find global territory entry without error")
	assert.Equal(t, "001", entry.Territory, "Should return entry with global territory")
	assert.Equal(t, "Test/Zone2", entry.IANA, "Should return correct IANA for global territory")

	testEntriesNo001 := []Entry{
		{MicrosoftAlias: "Test Time", Territory: "US", IANA: "Test/Zone1"},
		{MicrosoftAlias: "Test Time", Territory: "CA", IANA: "Test/Zone3"},
	}

	_, err = one001(testEntriesNo001)
	assert.Error(t, err, "Should return error when no global territory present")
	assert.ErrorIs(t, err, ErrNotFound, "Should return ErrNotFound specifically")

	_, err = one001([]Entry{})
	assert.Error(t, err, "Should return error for empty slice")
	assert.ErrorIs(t, err, ErrNotFound, "Should return ErrNotFound for empty slice")
}

// Integration tests

func TestIntegrationRealData(t *testing.T) {
	require.NotEmpty(t, data, "Data slice should be loaded with timezone entries")
	require.NotEmpty(t, indexMicrosoftAlias, "MicrosoftAlias index should be built")
	require.NotEmpty(t, indexMicrosoftDisplayName, "MicrosoftDisplayName index should be built")
	require.NotEmpty(t, indexIANA, "IANA index should be built")

	knownMappings := map[string]string{
		"Hawaiian Standard Time":       "Pacific/Honolulu",
		"Eastern Standard Time":        "America/New_York",
		"Central Europe Standard Time": "Europe/Budapest",
		"GMT Standard Time":            "Europe/London",
	}

	for alias, expectedIANA := range knownMappings {
		entry, err := FromMicrosoftAlias(alias)
		require.NoError(t, err, "Should find known alias '%s'", alias)
		assert.Equal(t, expectedIANA, entry.IANA, "Alias '%s' should map to correct IANA timezone", alias)
		assert.Equal(t, "001", entry.Territory, "Alias '%s' should have global territory", alias)
	}
}

func TestTerritoryVariations(t *testing.T) {
	aliases := []string{
		"Central Europe Standard Time",
		"Eastern Standard Time",
		"Hawaiian Standard Time",
		"West Asia Standard Time",
	}

	for _, alias := range aliases {
		entries := AllFromMicrosoftAlias(alias)
		assert.Greater(t, len(entries), 1, "Alias '%s' should have multiple territory entries", alias)

		found001 := false
		for _, entry := range entries {
			if entry.Territory == "001" {
				found001 = true
				break
			}
		}
		assert.True(t, found001, "Alias '%s' should have global territory '001' entry", alias)
	}
}

func TestRoundTripConversions(t *testing.T) {
	testAliases := []string{
		"Hawaiian Standard Time",
		"Eastern Standard Time",
		"Central Europe Standard Time",
		"Tokyo Standard Time",
	}

	for _, originalAlias := range testAliases {
		entry1, err := FromMicrosoftAlias(originalAlias)
		require.NoError(t, err, "Should find alias '%s' in first lookup", originalAlias)

		entry2, err := FromIANA(entry1.IANA)
		require.NoError(t, err, "Should find IANA '%s' in reverse lookup", entry1.IANA)

		assert.Equal(t, originalAlias, entry2.MicrosoftAlias, "Round-trip conversion should return to original alias: %s -> %s -> %s", originalAlias, entry1.IANA, entry2.MicrosoftAlias)
	}
}

func TestAllDataEntriesIntegrity(t *testing.T) {
	require.NotEmpty(t, data, "Data slice should not be empty")

	for _, e := range data {
		_, err := FromMicrosoftDisplayName(e.MicrosoftDisplayName)
		assert.NoError(t, err, "FromMicrosoftDisplayName should find default entry for display name: %s", e.MicrosoftDisplayName)

		_, err = FromMicrosoftAlias(e.MicrosoftAlias)
		assert.NoError(t, err, "FromMicrosoftAlias should find default entry for alias: %s", e.MicrosoftAlias)

		_, err = FromIANA(e.IANA)
		assert.NoError(t, err, "FromIANA should find default entry for IANA: %s", e.IANA)
	}
}

// Benchmark tests

func BenchmarkFromMicrosoftAlias(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = FromMicrosoftAlias("Hawaiian Standard Time")
	}
}

func BenchmarkFromIANA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = FromIANA("Pacific/Honolulu")
	}
}

func BenchmarkAllFromMicrosoftAlias(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AllFromMicrosoftAlias("Central Europe Standard Time")
	}
}
