.PHONY: download convert-tzid generate verify test benchmark all clean

download:
	@echo "Downloading reference files..."
	@mkdir -p references
	curl -o references/windowsZones.xml https://raw.githubusercontent.com/unicode-org/cldr/main/common/supplemental/windowsZones.xml
	curl -o references/idList.txt https://raw.githubusercontent.com/unicode-org/cldr/refs/heads/main/tools/cldr-code/src/main/resources/org/unicode/cldr/icu/idList.txt

convert-tzid: references/idList.txt
	@echo "Converting idList.txt to JSON format..."
	@grep "^tzid" references/idList.txt | awk -F';' '{ \
		gsub(/^[ \t]+|[ \t]+$$/, "", $$2); \
		territory = ""; \
		if (NF >= 3) { \
			gsub(/^[ \t]+|[ \t]+$$/, "", $$3); \
			if ($$3 != "") territory = $$3; \
		} \
		if (NF >= 4) { \
			gsub(/^[ \t]+|[ \t]+$$/, "", $$4); \
			if ($$4 != "") territory = $$4; \
		} \
		split($$2, tzs, ", "); \
		for (i in tzs) { \
			gsub(/^[ \t]+|[ \t]+$$/, "", tzs[i]); \
			if (tzs[i] != "") { \
				printf "{\"timezone\":\"%s\",\"territory\":\"%s\"}\n", tzs[i], territory; \
			} \
		} \
	}' | jq -s 'sort_by(.timezone)' > references/TZID.json

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
	rm -f references/TZID.json references/windowsZones.xml references/idList.txt
