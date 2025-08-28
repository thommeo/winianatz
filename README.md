# winianatz

A Go package that provides support for converting date timezones from the non-standard Windows timezone format into IANA timezones and UTC.

This package is auto-generated from the [Unicode CLDR project's windowsZones.xml](https://github.com/unicode-org/cldr/blob/main/common/supplemental/windowsZones.xml) file. Repository updates the data on a weekly basis. If data or code change is detected, a pull request will be generated automatically.

See [verification report](./VERIFY.md) for cross check results between [Microsoft Graph API supported timezones](./references/msgraph-supported-timezones-windows.json) and IANA timezones.

Inspired by: https://github.com/thinkovation/windowsiana

## Installation

```bash
go get github.com/thommeo/winianatz
```

## Usage

```go
package main

import (
    "fmt"
    "log"
    "github.com/thommeo/winianatz"
)

func main() {
    // Convert Windows timezone to IANA
    windowsTZ := "(UTC+08:00) Beijing, Chongqing, Hong Kong, Urumqi"
    ianaTZ := winianatz.WinIANA[windowsTZ]
    fmt.Printf("Windows: %s\nIANA: %s\n", windowsTZ, ianaTZ)

    // Parse time with Windows timezone
    t, err := winianatz.TimezoneParseWindows("2024-01-15T10:30:00", windowsTZ)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("UTC Time: %s\n", t.UTC())
}
```
