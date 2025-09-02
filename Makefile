.PHONY: download convert-tzid generate verify test benchmark all clean

download:
	@echo "Downloading reference files..."
	@mkdir -p references
	curl -o references/windowsZones.xml https://raw.githubusercontent.com/unicode-org/cldr/main/common/supplemental/windowsZones.xml
	curl -o references/TZID.txt https://raw.githubusercontent.com/unicode-org/cldr/refs/heads/main/tools/cldr-code/src/main/resources/org/unicode/cldr/util/data/TZID.txt

convert-tzid: references/TZID.txt
	@echo "Converting TZID.txt to JSON format..."
	@grep -v "^#" references/TZID.txt | awk -F'|' 'NF==3 { \
		gsub(/^[ \t]+|[ \t]+$$/, "", $$2); \
		gsub(/^[ \t]+|[ \t]+$$/, "", $$3); \
		split($$2, tzs, ", "); \
		for (i in tzs) { \
			gsub(/^[ \t]+|[ \t]+$$/, "", tzs[i]); \
			if (tzs[i] != "") { \
				printf "{\"timezone\":\"%s\",\"territory\":\"%s\"}\n", tzs[i], $$3; \
			} \
		} \
	}' | jq -s '.' > references/TZID.json

generate:
	go run ./internal/generator

verify:
	go run ./internal/verifier

test:
	go test -v ./...

benchmark:
	go test -bench=. -benchtime=1s ./...

all: download generate verify test

clean:
	rm -f data_gen.go data_gen.json
	rm -f references/TZID.json references/windowsZones.xml references/TZID.txt
