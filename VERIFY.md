# Timezone conversion verification report

## Microsoft Graph API Timezone Verification

This is a cross check results between [Microsoft Graph API supported timezones](./references/msgraph-supported-timezones-windows.json) and IANA timezones.

The conversion is implemented based on [Unicode CLDR project's windowsZones.xml](https://github.com/unicode-org/cldr/blob/main/common/supplemental/windowsZones.xml).

| Microsoft Alias | Microsoft Display Name |
|-----------------|------------------------|
| Dateline Standard Time<br>✅ Etc/GMT+12 (001) | (UTC-12:00) International Date Line West<br>✅ Etc/GMT+12 (001) |
| UTC-11<br>✅ Etc/GMT+11 (001) | (UTC-11:00) Coordinated Universal Time-11<br>✅ Etc/GMT+11 (001) |
| Samoa Standard Time<br>✅ Pacific/Apia (001) | (UTC+13:00) Samoa<br>✅ Pacific/Apia (001) |
| Aleutian Standard Time<br>✅ America/Adak (001) | (UTC-10:00) Aleutian Islands<br>✅ America/Adak (001) |
| Hawaiian Standard Time<br>✅ Pacific/Honolulu (001) | (UTC-10:00) Hawaii<br>✅ Pacific/Honolulu (001) |
| Marquesas Standard Time<br>✅ Pacific/Marquesas (001) | (UTC-09:30) Marquesas Islands<br>✅ Pacific/Marquesas (001) |
| Alaskan Standard Time<br>✅ America/Anchorage (001) | (UTC-09:00) Alaska<br>✅ America/Anchorage (001) |
| UTC-09<br>✅ Etc/GMT+9 (001) | (UTC-09:00) Coordinated Universal Time-09<br>✅ Etc/GMT+9 (001) |
| Yukon Standard Time<br>✅ America/Whitehorse (001) | (UTC-07:00) Yukon<br>✅ America/Whitehorse (001) |
| Pacific Standard Time (Mexico)<br>✅ America/Tijuana (001) | (UTC-08:00) Baja California<br>✅ America/Tijuana (001) |
| UTC-08<br>✅ Etc/GMT+8 (001) | (UTC-08:00) Coordinated Universal Time-08<br>✅ Etc/GMT+8 (001) |
| Pacific Standard Time<br>✅ America/Los_Angeles (001) | (UTC-08:00) Pacific Time (US & Canada)<br>✅ America/Los_Angeles (001) |
| US Mountain Standard Time<br>✅ America/Phoenix (001) | (UTC-07:00) Arizona<br>✅ America/Phoenix (001) |
| Mountain Standard Time (Mexico)<br>✅ America/Mazatlan (001) | (UTC-07:00) La Paz, Mazatlan<br>❌ IANA Not found |
| Mountain Standard Time<br>✅ America/Denver (001) | (UTC-07:00) Mountain Time (US & Canada)<br>✅ America/Denver (001) |
| Eastern Standard Time (Mexico)<br>✅ America/Cancun (001) | (UTC-05:00) Chetumal<br>✅ America/Cancun (001) |
| Central America Standard Time<br>✅ America/Guatemala (001) | (UTC-06:00) Central America<br>✅ America/Guatemala (001) |
| Central Standard Time<br>✅ America/Chicago (001) | (UTC-06:00) Central Time (US & Canada)<br>✅ America/Chicago (001) |
| Easter Island Standard Time<br>✅ Pacific/Easter (001) | (UTC-06:00) Easter Island<br>✅ Pacific/Easter (001) |
| Central Standard Time (Mexico)<br>✅ America/Mexico_City (001) | (UTC-06:00) Guadalajara, Mexico City, Monterrey<br>✅ America/Mexico_City (001) |
| Canada Central Standard Time<br>✅ America/Regina (001) | (UTC-06:00) Saskatchewan<br>✅ America/Regina (001) |
| SA Pacific Standard Time<br>✅ America/Bogota (001) | (UTC-05:00) Bogota, Lima, Quito, Rio Branco<br>✅ America/Bogota (001) |
| Eastern Standard Time<br>✅ America/New_York (001) | (UTC-05:00) Eastern Time (US & Canada)<br>✅ America/New_York (001) |
| Haiti Standard Time<br>✅ America/Port-au-Prince (001) | (UTC-05:00) Haiti<br>✅ America/Port-au-Prince (001) |
| Cuba Standard Time<br>✅ America/Havana (001) | (UTC-05:00) Havana<br>✅ America/Havana (001) |
| US Eastern Standard Time<br>✅ America/Indianapolis (001) | (UTC-05:00) Indiana (East)<br>✅ America/Indianapolis (001) |
| Turks And Caicos Standard Time<br>✅ America/Grand_Turk (001) | (UTC-05:00) Turks and Caicos<br>✅ America/Grand_Turk (001) |
| Venezuela Standard Time<br>✅ America/Caracas (001) | (UTC-04:00) Caracas<br>✅ America/Caracas (001) |
| Magallanes Standard Time<br>✅ America/Punta_Arenas (001) | (UTC-03:00) Punta Arenas<br>✅ America/Punta_Arenas (001) |
| Paraguay Standard Time<br>✅ America/Asuncion (001) | (UTC-03:00) Asuncion<br>❌ IANA Not found |
| Atlantic Standard Time<br>✅ America/Halifax (001) | (UTC-04:00) Atlantic Time (Canada)<br>✅ America/Halifax (001) |
| Central Brazilian Standard Time<br>✅ America/Cuiaba (001) | (UTC-04:00) Cuiaba<br>✅ America/Cuiaba (001) |
| SA Western Standard Time<br>✅ America/La_Paz (001) | (UTC-04:00) Georgetown, La Paz, Manaus, San Juan<br>✅ America/La_Paz (001) |
| Pacific SA Standard Time<br>✅ America/Santiago (001) | (UTC-04:00) Santiago<br>✅ America/Santiago (001) |
| Newfoundland Standard Time<br>✅ America/St_Johns (001) | (UTC-03:30) Newfoundland<br>✅ America/St_Johns (001) |
| Tocantins Standard Time<br>✅ America/Araguaina (001) | (UTC-03:00) Araguaina<br>✅ America/Araguaina (001) |
| E. South America Standard Time<br>✅ America/Sao_Paulo (001) | (UTC-03:00) Brasilia<br>✅ America/Sao_Paulo (001) |
| SA Eastern Standard Time<br>✅ America/Cayenne (001) | (UTC-03:00) Cayenne, Fortaleza<br>✅ America/Cayenne (001) |
| Argentina Standard Time<br>✅ America/Buenos_Aires (001) | (UTC-03:00) City of Buenos Aires<br>✅ America/Buenos_Aires (001) |
| Greenland Standard Time<br>✅ America/Godthab (001) | (UTC-02:00) Greenland<br>❌ IANA Not found |
| Montevideo Standard Time<br>✅ America/Montevideo (001) | (UTC-03:00) Montevideo<br>✅ America/Montevideo (001) |
| Saint Pierre Standard Time<br>✅ America/Miquelon (001) | (UTC-03:00) Saint Pierre and Miquelon<br>✅ America/Miquelon (001) |
| Bahia Standard Time<br>✅ America/Bahia (001) | (UTC-03:00) Salvador<br>✅ America/Bahia (001) |
| UTC-02<br>✅ Etc/GMT+2 (001) | (UTC-02:00) Coordinated Universal Time-02<br>✅ Etc/GMT+2 (001) |
| Mid-Atlantic Standard Time<br>❌ IANA Not found | (UTC-02:00) Mid-Atlantic - Old<br>❌ IANA Not found |
| Azores Standard Time<br>✅ Atlantic/Azores (001) | (UTC-01:00) Azores<br>✅ Atlantic/Azores (001) |
| Cape Verde Standard Time<br>✅ Atlantic/Cape_Verde (001) | (UTC-01:00) Cabo Verde Is.<br>✅ Atlantic/Cape_Verde (001) |
| UTC<br>✅ Etc/UTC (001) | (UTC) Coordinated Universal Time<br>✅ Etc/UTC (001) |
| GMT Standard Time<br>✅ Europe/London (001) | (UTC+00:00) Dublin, Edinburgh, Lisbon, London<br>✅ Europe/London (001) |
| Greenwich Standard Time<br>✅ Atlantic/Reykjavik (001) | (UTC+00:00) Monrovia, Reykjavik<br>✅ Atlantic/Reykjavik (001) |
| Morocco Standard Time<br>✅ Africa/Casablanca (001) | (UTC+01:00) Casablanca<br>✅ Africa/Casablanca (001) |
| W. Europe Standard Time<br>✅ Europe/Berlin (001) | (UTC+01:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna<br>✅ Europe/Berlin (001) |
| Central Europe Standard Time<br>✅ Europe/Budapest (001) | (UTC+01:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague<br>✅ Europe/Budapest (001) |
| Romance Standard Time<br>✅ Europe/Paris (001) | (UTC+01:00) Brussels, Copenhagen, Madrid, Paris<br>✅ Europe/Paris (001) |
| Central European Standard Time<br>✅ Europe/Warsaw (001) | (UTC+01:00) Sarajevo, Skopje, Warsaw, Zagreb<br>✅ Europe/Warsaw (001) |
| W. Central Africa Standard Time<br>✅ Africa/Lagos (001) | (UTC+01:00) West Central Africa<br>✅ Africa/Lagos (001) |
| Libya Standard Time<br>✅ Africa/Tripoli (001) | (UTC+02:00) Tripoli<br>✅ Africa/Tripoli (001) |
| Namibia Standard Time<br>✅ Africa/Windhoek (001) | (UTC+02:00) Windhoek<br>✅ Africa/Windhoek (001) |
| GTB Standard Time<br>✅ Europe/Bucharest (001) | (UTC+02:00) Athens, Bucharest<br>✅ Europe/Bucharest (001) |
| Middle East Standard Time<br>✅ Asia/Beirut (001) | (UTC+02:00) Beirut<br>✅ Asia/Beirut (001) |
| Egypt Standard Time<br>✅ Africa/Cairo (001) | (UTC+02:00) Cairo<br>✅ Africa/Cairo (001) |
| E. Europe Standard Time<br>✅ Europe/Chisinau (001) | (UTC+02:00) Chisinau<br>✅ Europe/Chisinau (001) |
| Syria Standard Time<br>✅ Asia/Damascus (001) | (UTC+03:00) Damascus<br>❌ IANA Not found |
| West Bank Standard Time<br>✅ Asia/Hebron (001) | (UTC+02:00) Gaza, Hebron<br>✅ Asia/Hebron (001) |
| South Africa Standard Time<br>✅ Africa/Johannesburg (001) | (UTC+02:00) Harare, Pretoria<br>✅ Africa/Johannesburg (001) |
| FLE Standard Time<br>✅ Europe/Kiev (001) | (UTC+02:00) Helsinki, Kyiv, Riga, Sofia, Tallinn, Vilnius<br>✅ Europe/Kiev (001) |
| Israel Standard Time<br>✅ Asia/Jerusalem (001) | (UTC+02:00) Jerusalem<br>✅ Asia/Jerusalem (001) |
| South Sudan Standard Time<br>✅ Africa/Juba (001) | (UTC+02:00) Juba<br>✅ Africa/Juba (001) |
| Kaliningrad Standard Time<br>✅ Europe/Kaliningrad (001) | (UTC+02:00) Kaliningrad<br>✅ Europe/Kaliningrad (001) |
| Sudan Standard Time<br>✅ Africa/Khartoum (001) | (UTC+02:00) Khartoum<br>✅ Africa/Khartoum (001) |
| Jordan Standard Time<br>✅ Asia/Amman (001) | (UTC+03:00) Amman<br>❌ IANA Not found |
| Turkey Standard Time<br>✅ Europe/Istanbul (001) | (UTC+03:00) Istanbul<br>✅ Europe/Istanbul (001) |
| Belarus Standard Time<br>✅ Europe/Minsk (001) | (UTC+03:00) Minsk<br>✅ Europe/Minsk (001) |
| Arabic Standard Time<br>✅ Asia/Baghdad (001) | (UTC+03:00) Baghdad<br>✅ Asia/Baghdad (001) |
| Arab Standard Time<br>✅ Asia/Riyadh (001) | (UTC+03:00) Kuwait, Riyadh<br>✅ Asia/Riyadh (001) |
| Russian Standard Time<br>✅ Europe/Moscow (001) | (UTC+03:00) Moscow, St. Petersburg<br>✅ Europe/Moscow (001) |
| E. Africa Standard Time<br>✅ Africa/Nairobi (001) | (UTC+03:00) Nairobi<br>✅ Africa/Nairobi (001) |
| Volgograd Standard Time<br>✅ Europe/Volgograd (001) | (UTC+03:00) Volgograd<br>❌ IANA Not found |
| Astrakhan Standard Time<br>✅ Europe/Astrakhan (001) | (UTC+04:00) Astrakhan, Ulyanovsk<br>✅ Europe/Astrakhan (001) |
| Russia Time Zone 3<br>✅ Europe/Samara (001) | (UTC+04:00) Izhevsk, Samara<br>✅ Europe/Samara (001) |
| Saratov Standard Time<br>✅ Europe/Saratov (001) | (UTC+04:00) Saratov<br>✅ Europe/Saratov (001) |
| Iran Standard Time<br>✅ Asia/Tehran (001) | (UTC+03:30) Tehran<br>✅ Asia/Tehran (001) |
| Arabian Standard Time<br>✅ Asia/Dubai (001) | (UTC+04:00) Abu Dhabi, Muscat<br>✅ Asia/Dubai (001) |
| Azerbaijan Standard Time<br>✅ Asia/Baku (001) | (UTC+04:00) Baku<br>✅ Asia/Baku (001) |
| Mauritius Standard Time<br>✅ Indian/Mauritius (001) | (UTC+04:00) Port Louis<br>✅ Indian/Mauritius (001) |
| Georgian Standard Time<br>✅ Asia/Tbilisi (001) | (UTC+04:00) Tbilisi<br>✅ Asia/Tbilisi (001) |
| Caucasus Standard Time<br>✅ Asia/Yerevan (001) | (UTC+04:00) Yerevan<br>✅ Asia/Yerevan (001) |
| Afghanistan Standard Time<br>✅ Asia/Kabul (001) | (UTC+04:30) Kabul<br>✅ Asia/Kabul (001) |
| West Asia Standard Time<br>✅ Asia/Tashkent (001) | (UTC+05:00) Ashgabat, Tashkent<br>✅ Asia/Tashkent (001) |
| Qyzylorda Standard Time<br>✅ Asia/Qyzylorda (001) | (UTC+05:00) Astana<br>❌ IANA Not found |
| Ekaterinburg Standard Time<br>✅ Asia/Yekaterinburg (001) | (UTC+05:00) Ekaterinburg<br>✅ Asia/Yekaterinburg (001) |
| Pakistan Standard Time<br>✅ Asia/Karachi (001) | (UTC+05:00) Islamabad, Karachi<br>✅ Asia/Karachi (001) |
| India Standard Time<br>✅ Asia/Calcutta (001) | (UTC+05:30) Chennai, Kolkata, Mumbai, New Delhi<br>✅ Asia/Calcutta (001) |
| Sri Lanka Standard Time<br>✅ Asia/Colombo (001) | (UTC+05:30) Sri Jayawardenepura<br>✅ Asia/Colombo (001) |
| Nepal Standard Time<br>✅ Asia/Katmandu (001) | (UTC+05:45) Kathmandu<br>✅ Asia/Katmandu (001) |
| Central Asia Standard Time<br>✅ Asia/Bishkek (001) | (UTC+06:00) Bishkek<br>❌ IANA Not found |
| Bangladesh Standard Time<br>✅ Asia/Dhaka (001) | (UTC+06:00) Dhaka<br>✅ Asia/Dhaka (001) |
| Omsk Standard Time<br>✅ Asia/Omsk (001) | (UTC+06:00) Omsk<br>✅ Asia/Omsk (001) |
| Altai Standard Time<br>✅ Asia/Barnaul (001) | (UTC+07:00) Barnaul, Gorno-Altaysk<br>✅ Asia/Barnaul (001) |
| N. Central Asia Standard Time<br>✅ Asia/Novosibirsk (001) | (UTC+07:00) Novosibirsk<br>✅ Asia/Novosibirsk (001) |
| Tomsk Standard Time<br>✅ Asia/Tomsk (001) | (UTC+07:00) Tomsk<br>✅ Asia/Tomsk (001) |
| Myanmar Standard Time<br>✅ Asia/Rangoon (001) | (UTC+06:30) Yangon (Rangoon)<br>✅ Asia/Rangoon (001) |
| SE Asia Standard Time<br>✅ Asia/Bangkok (001) | (UTC+07:00) Bangkok, Hanoi, Jakarta<br>✅ Asia/Bangkok (001) |
| W. Mongolia Standard Time<br>✅ Asia/Hovd (001) | (UTC+07:00) Hovd<br>✅ Asia/Hovd (001) |
| North Asia Standard Time<br>✅ Asia/Krasnoyarsk (001) | (UTC+07:00) Krasnoyarsk<br>✅ Asia/Krasnoyarsk (001) |
| China Standard Time<br>✅ Asia/Shanghai (001) | (UTC+08:00) Beijing, Chongqing, Hong Kong, Urumqi<br>✅ Asia/Shanghai (001) |
| North Asia East Standard Time<br>✅ Asia/Irkutsk (001) | (UTC+08:00) Irkutsk<br>✅ Asia/Irkutsk (001) |
| Singapore Standard Time<br>✅ Asia/Singapore (001) | (UTC+08:00) Kuala Lumpur, Singapore<br>✅ Asia/Singapore (001) |
| W. Australia Standard Time<br>✅ Australia/Perth (001) | (UTC+08:00) Perth<br>✅ Australia/Perth (001) |
| Taipei Standard Time<br>✅ Asia/Taipei (001) | (UTC+08:00) Taipei<br>✅ Asia/Taipei (001) |
| Ulaanbaatar Standard Time<br>✅ Asia/Ulaanbaatar (001) | (UTC+08:00) Ulaanbaatar<br>✅ Asia/Ulaanbaatar (001) |
| Transbaikal Standard Time<br>✅ Asia/Chita (001) | (UTC+09:00) Chita<br>✅ Asia/Chita (001) |
| North Korea Standard Time<br>✅ Asia/Pyongyang (001) | (UTC+09:00) Pyongyang<br>✅ Asia/Pyongyang (001) |
| Aus Central W. Standard Time<br>✅ Australia/Eucla (001) | (UTC+08:45) Eucla<br>✅ Australia/Eucla (001) |
| Tokyo Standard Time<br>✅ Asia/Tokyo (001) | (UTC+09:00) Osaka, Sapporo, Tokyo<br>✅ Asia/Tokyo (001) |
| Korea Standard Time<br>✅ Asia/Seoul (001) | (UTC+09:00) Seoul<br>✅ Asia/Seoul (001) |
| Yakutsk Standard Time<br>✅ Asia/Yakutsk (001) | (UTC+09:00) Yakutsk<br>✅ Asia/Yakutsk (001) |
| Cen. Australia Standard Time<br>✅ Australia/Adelaide (001) | (UTC+09:30) Adelaide<br>✅ Australia/Adelaide (001) |
| AUS Central Standard Time<br>✅ Australia/Darwin (001) | (UTC+09:30) Darwin<br>✅ Australia/Darwin (001) |
| E. Australia Standard Time<br>✅ Australia/Brisbane (001) | (UTC+10:00) Brisbane<br>✅ Australia/Brisbane (001) |
| AUS Eastern Standard Time<br>✅ Australia/Sydney (001) | (UTC+10:00) Canberra, Melbourne, Sydney<br>✅ Australia/Sydney (001) |
| West Pacific Standard Time<br>✅ Pacific/Port_Moresby (001) | (UTC+10:00) Guam, Port Moresby<br>✅ Pacific/Port_Moresby (001) |
| Tasmania Standard Time<br>✅ Australia/Hobart (001) | (UTC+10:00) Hobart<br>✅ Australia/Hobart (001) |
| Vladivostok Standard Time<br>✅ Asia/Vladivostok (001) | (UTC+10:00) Vladivostok<br>✅ Asia/Vladivostok (001) |
| Bougainville Standard Time<br>✅ Pacific/Bougainville (001) | (UTC+11:00) Bougainville Island<br>✅ Pacific/Bougainville (001) |
| Magadan Standard Time<br>✅ Asia/Magadan (001) | (UTC+11:00) Magadan<br>✅ Asia/Magadan (001) |
| Sakhalin Standard Time<br>✅ Asia/Sakhalin (001) | (UTC+11:00) Sakhalin<br>✅ Asia/Sakhalin (001) |
| Lord Howe Standard Time<br>✅ Australia/Lord_Howe (001) | (UTC+10:30) Lord Howe Island<br>✅ Australia/Lord_Howe (001) |
| Russia Time Zone 10<br>✅ Asia/Srednekolymsk (001) | (UTC+11:00) Chokurdakh<br>✅ Asia/Srednekolymsk (001) |
| Norfolk Standard Time<br>✅ Pacific/Norfolk (001) | (UTC+11:00) Norfolk Island<br>✅ Pacific/Norfolk (001) |
| Central Pacific Standard Time<br>✅ Pacific/Guadalcanal (001) | (UTC+11:00) Solomon Is., New Caledonia<br>✅ Pacific/Guadalcanal (001) |
| Russia Time Zone 11<br>✅ Asia/Kamchatka (001) | (UTC+12:00) Anadyr, Petropavlovsk-Kamchatsky<br>✅ Asia/Kamchatka (001) |
| New Zealand Standard Time<br>✅ Pacific/Auckland (001) | (UTC+12:00) Auckland, Wellington<br>✅ Pacific/Auckland (001) |
| UTC+12<br>✅ Etc/GMT-12 (001) | (UTC+12:00) Coordinated Universal Time+12<br>✅ Etc/GMT-12 (001) |
| Fiji Standard Time<br>✅ Pacific/Fiji (001) | (UTC+12:00) Fiji<br>✅ Pacific/Fiji (001) |
| Kamchatka Standard Time<br>❌ IANA Not found | (UTC+12:00) Petropavlovsk-Kamchatsky - Old<br>❌ IANA Not found |
| Chatham Islands Standard Time<br>✅ Pacific/Chatham (001) | (UTC+12:45) Chatham Islands<br>✅ Pacific/Chatham (001) |
| UTC+13<br>✅ Etc/GMT-13 (001) | (UTC+13:00) Coordinated Universal Time+13<br>✅ Etc/GMT-13 (001) |
| Tonga Standard Time<br>✅ Pacific/Tongatapu (001) | (UTC+13:00) Nuku'alofa<br>✅ Pacific/Tongatapu (001) |
| Line Islands Standard Time<br>✅ Pacific/Kiritimati (001) | (UTC+14:00) Kiritimati Island<br>✅ Pacific/Kiritimati (001) |


## IANA Timezone ID Verification

This section verifies all IANA timezone IDs from [CLDR TZID.txt](./references/TZID.txt) against the winianatz FromIANA function.

**Summary:**
- Total IANA timezones tested: 402
- Found in winianatz: 384
- Missing from winianatz: 18
- Coverage: 95.5%

### All IANA Timezone Results

| IANA Timezone | Territory | Result |
|---------------|-----------|--------|
| Africa/Abidjan | CI | ✅ Greenwich Standard Time |
| Africa/Accra | GH | ✅ Greenwich Standard Time |
| Africa/Addis_Ababa | ET | ✅ E. Africa Standard Time |
| Africa/Algiers | DZ | ✅ W. Central Africa Standard Time |
| Africa/Asmera | ER | ✅ E. Africa Standard Time |
| Africa/Bamako | ML | ✅ Greenwich Standard Time |
| Africa/Bangui | CF | ✅ W. Central Africa Standard Time |
| Africa/Banjul | GM | ✅ Greenwich Standard Time |
| Africa/Bissau | GW | ✅ Greenwich Standard Time |
| Africa/Blantyre | MW | ✅ South Africa Standard Time |
| Africa/Brazzaville | CG | ✅ W. Central Africa Standard Time |
| Africa/Bujumbura | BI | ✅ South Africa Standard Time |
| Africa/Cairo | EG | ✅ Egypt Standard Time |
| Africa/Casablanca | MA | ✅ Morocco Standard Time |
| Africa/Ceuta | ES | ✅ Romance Standard Time |
| Africa/Conakry | GN | ✅ Greenwich Standard Time |
| Africa/Dakar | SN | ✅ Greenwich Standard Time |
| Africa/Dar_es_Salaam | TZ | ✅ E. Africa Standard Time |
| Africa/Djibouti | DJ | ✅ E. Africa Standard Time |
| Africa/Douala | CM | ✅ W. Central Africa Standard Time |
| Africa/El_Aaiun | EH | ✅ Morocco Standard Time |
| Africa/Freetown | SL | ✅ Greenwich Standard Time |
| Africa/Gaborone | BW | ✅ South Africa Standard Time |
| Africa/Harare | ZW | ✅ South Africa Standard Time |
| Africa/Johannesburg | ZA | ✅ South Africa Standard Time |
| Africa/Kampala | UG | ✅ E. Africa Standard Time |
| Africa/Khartoum | SD | ✅ Sudan Standard Time |
| Africa/Kigali | RW | ✅ South Africa Standard Time |
| Africa/Kinshasa | CD | ✅ W. Central Africa Standard Time |
| Africa/Lagos | NG | ✅ W. Central Africa Standard Time |
| Africa/Libreville | GA | ✅ W. Central Africa Standard Time |
| Africa/Lome | TG | ✅ Greenwich Standard Time |
| Africa/Luanda | AO | ✅ W. Central Africa Standard Time |
| Africa/Lubumbashi | CD | ✅ South Africa Standard Time |
| Africa/Lusaka | ZM | ✅ South Africa Standard Time |
| Africa/Malabo | GQ | ✅ W. Central Africa Standard Time |
| Africa/Maputo | MZ | ✅ South Africa Standard Time |
| Africa/Maseru | LS | ✅ South Africa Standard Time |
| Africa/Mbabane | SZ | ✅ South Africa Standard Time |
| Africa/Mogadishu | SO | ✅ E. Africa Standard Time |
| Africa/Monrovia | LR | ✅ Greenwich Standard Time |
| Africa/Nairobi | KE | ✅ E. Africa Standard Time |
| Africa/Ndjamena | TD | ✅ W. Central Africa Standard Time |
| Africa/Niamey | NE | ✅ W. Central Africa Standard Time |
| Africa/Nouakchott | MR | ✅ Greenwich Standard Time |
| Africa/Ouagadougou | BF | ✅ Greenwich Standard Time |
| Africa/Porto-Novo | BJ | ✅ W. Central Africa Standard Time |
| Africa/Sao_Tome | ST | ✅ Sao Tome Standard Time |
| Africa/Timbuktu | ML | ❌ Not found |
| Africa/Tripoli | LY | ✅ Libya Standard Time |
| Africa/Tunis | TN | ✅ W. Central Africa Standard Time |
| Africa/Windhoek | NA | ✅ Namibia Standard Time |
| America/Adak | US | ✅ Aleutian Standard Time |
| America/Anchorage | US | ✅ Alaskan Standard Time |
| America/Anguilla | AI | ✅ SA Western Standard Time |
| America/Antigua | AG | ✅ SA Western Standard Time |
| America/Araguaina | BR | ✅ Tocantins Standard Time |
| America/Aruba | AW | ✅ SA Western Standard Time |
| America/Asuncion | PY | ✅ Paraguay Standard Time |
| America/Barbados | BB | ✅ SA Western Standard Time |
| America/Belem | BR | ✅ SA Eastern Standard Time |
| America/Belize | BZ | ✅ Central America Standard Time |
| America/Boa_Vista | BR | ✅ SA Western Standard Time |
| America/Bogota | CO | ✅ SA Pacific Standard Time |
| America/Boise | US | ✅ Mountain Standard Time |
| America/Buenos_Aires | AR | ✅ Argentina Standard Time |
| America/Cambridge_Bay | CA | ✅ Mountain Standard Time |
| America/Cancun | MX | ✅ Eastern Standard Time (Mexico) |
| America/Caracas | VE | ✅ Venezuela Standard Time |
| America/Catamarca | AR | ✅ Argentina Standard Time |
| America/Cayenne | GF | ✅ SA Eastern Standard Time |
| America/Cayman | KY | ✅ SA Pacific Standard Time |
| America/Chicago | US | ✅ Central Standard Time |
| America/Chihuahua | MX | ✅ Central Standard Time (Mexico) |
| America/Cordoba | AR | ✅ Argentina Standard Time |
| America/Costa_Rica | CR | ✅ Central America Standard Time |
| America/Cuiaba | BR | ✅ Central Brazilian Standard Time |
| America/Curacao | AN | ✅ SA Western Standard Time |
| America/Danmarkshavn | GL | ✅ Greenwich Standard Time |
| America/Dawson | CA | ✅ Yukon Standard Time |
| America/Dawson_Creek | CA | ✅ US Mountain Standard Time |
| America/Denver | US | ✅ Mountain Standard Time |
| America/Detroit | US | ✅ Eastern Standard Time |
| America/Dominica | DM | ✅ SA Western Standard Time |
| America/Edmonton | CA | ✅ Mountain Standard Time |
| America/Eirunepe | BR | ✅ SA Pacific Standard Time |
| America/El_Salvador | SV | ✅ Central America Standard Time |
| America/Fortaleza | BR | ✅ SA Eastern Standard Time |
| America/Glace_Bay | CA | ✅ Atlantic Standard Time |
| America/Godthab | GL | ✅ Greenland Standard Time |
| America/Goose_Bay | CA | ✅ Atlantic Standard Time |
| America/Grand_Turk | TC | ✅ Turks And Caicos Standard Time |
| America/Grenada | GD | ✅ SA Western Standard Time |
| America/Guadeloupe | GP | ✅ SA Western Standard Time |
| America/Guatemala | GT | ✅ Central America Standard Time |
| America/Guayaquil | EC | ✅ SA Pacific Standard Time |
| America/Guyana | GY | ✅ SA Western Standard Time |
| America/Halifax | CA | ✅ Atlantic Standard Time |
| America/Havana | CU | ✅ Cuba Standard Time |
| America/Hermosillo | MX | ✅ US Mountain Standard Time |
| America/Indiana/Knox | US | ✅ Central Standard Time |
| America/Indiana/Marengo | US | ✅ US Eastern Standard Time |
| America/Indiana/Vevay | US | ✅ US Eastern Standard Time |
| America/Indianapolis | US | ✅ US Eastern Standard Time |
| America/Inuvik | CA | ✅ Mountain Standard Time |
| America/Iqaluit | CA | ✅ Eastern Standard Time |
| America/Jamaica | JM | ✅ SA Pacific Standard Time |
| America/Jujuy | AR | ✅ Argentina Standard Time |
| America/Juneau | US | ✅ Alaskan Standard Time |
| America/Kentucky/Monticello | US | ✅ Eastern Standard Time |
| America/La_Paz | BO | ✅ SA Western Standard Time |
| America/Lima | PE | ✅ SA Pacific Standard Time |
| America/Los_Angeles | US | ✅ Pacific Standard Time |
| America/Louisville | US | ✅ Eastern Standard Time |
| America/Maceio | BR | ✅ SA Eastern Standard Time |
| America/Managua | NI | ✅ Central America Standard Time |
| America/Manaus | BR | ✅ SA Western Standard Time |
| America/Martinique | MQ | ✅ SA Western Standard Time |
| America/Mazatlan | MX | ✅ Mountain Standard Time (Mexico) |
| America/Mendoza | AR | ✅ Argentina Standard Time |
| America/Menominee | US | ✅ Central Standard Time |
| America/Merida | MX | ✅ Central Standard Time (Mexico) |
| America/Mexico_City | MX | ✅ Central Standard Time (Mexico) |
| America/Miquelon | PM | ✅ Saint Pierre Standard Time |
| America/Monterrey | MX | ✅ Central Standard Time (Mexico) |
| America/Montevideo | UY | ✅ Montevideo Standard Time |
| America/Montserrat | MS | ✅ SA Western Standard Time |
| America/Nassau | BS | ✅ Eastern Standard Time |
| America/New_York | US | ✅ Eastern Standard Time |
| America/Nipigon | CA | ❌ Not found |
| America/Nome | US | ✅ Alaskan Standard Time |
| America/Noronha | BR | ✅ UTC-02 |
| America/North_Dakota/Center | US | ✅ Central Standard Time |
| America/Panama | PA | ✅ SA Pacific Standard Time |
| America/Pangnirtung | CA | ❌ Not found |
| America/Paramaribo | SR | ✅ SA Eastern Standard Time |
| America/Phoenix | US | ✅ US Mountain Standard Time |
| America/Port-au-Prince | HT | ✅ Haiti Standard Time |
| America/Port_of_Spain | TT | ✅ SA Western Standard Time |
| America/Porto_Velho | BR | ✅ SA Western Standard Time |
| America/Puerto_Rico | PR | ✅ SA Western Standard Time |
| America/Rainy_River | CA | ❌ Not found |
| America/Rankin_Inlet | CA | ✅ Central Standard Time |
| America/Recife | BR | ✅ SA Eastern Standard Time |
| America/Regina | CA | ✅ Canada Central Standard Time |
| America/Rio_Branco | BR | ✅ SA Pacific Standard Time |
| America/Santiago | CL | ✅ Pacific SA Standard Time |
| America/Santo_Domingo | DO | ✅ SA Western Standard Time |
| America/Sao_Paulo | BR | ✅ E. South America Standard Time |
| America/Scoresbysund | GL | ✅ Azores Standard Time |
| America/St_Johns | CA | ✅ Newfoundland Standard Time |
| America/St_Kitts | KN | ✅ SA Western Standard Time |
| America/St_Lucia | LC | ✅ SA Western Standard Time |
| America/St_Thomas | VI | ✅ SA Western Standard Time |
| America/St_Vincent | VC | ✅ SA Western Standard Time |
| America/Swift_Current | CA | ✅ Canada Central Standard Time |
| America/Tegucigalpa | HN | ✅ Central America Standard Time |
| America/Thule | GL | ✅ Atlantic Standard Time |
| America/Thunder_Bay | CA | ❌ Not found |
| America/Tijuana | MX | ✅ Pacific Standard Time (Mexico) |
| America/Tortola | VG | ✅ SA Western Standard Time |
| America/Vancouver | CA | ✅ Pacific Standard Time |
| America/Whitehorse | CA | ✅ Yukon Standard Time |
| America/Winnipeg | CA | ✅ Central Standard Time |
| America/Yakutat | US | ✅ Alaskan Standard Time |
| America/Yellowknife | CA | ❌ Not found |
| Antarctica/Casey | AQ | ✅ Central Pacific Standard Time |
| Antarctica/Davis | AQ | ✅ SE Asia Standard Time |
| Antarctica/DumontDUrville | AQ | ✅ West Pacific Standard Time |
| Antarctica/Mawson | AQ | ✅ West Asia Standard Time |
| Antarctica/McMurdo | AQ | ✅ New Zealand Standard Time |
| Antarctica/Palmer | AQ | ✅ SA Eastern Standard Time |
| Antarctica/Rothera | AQ | ✅ SA Eastern Standard Time |
| Antarctica/Syowa | AQ | ✅ E. Africa Standard Time |
| Antarctica/Vostok | AQ | ✅ Central Asia Standard Time |
| Asia/Aden | YE | ✅ Arab Standard Time |
| Asia/Almaty | KZ | ✅ West Asia Standard Time |
| Asia/Amman | JO | ✅ Jordan Standard Time |
| Asia/Anadyr | RU | ✅ Russia Time Zone 11 |
| Asia/Aqtau | KZ | ✅ West Asia Standard Time |
| Asia/Aqtobe | KZ | ✅ West Asia Standard Time |
| Asia/Ashgabat | TM | ✅ West Asia Standard Time |
| Asia/Baghdad | IQ | ✅ Arabic Standard Time |
| Asia/Bahrain | BH | ✅ Arab Standard Time |
| Asia/Baku | AZ | ✅ Azerbaijan Standard Time |
| Asia/Bangkok | TH | ✅ SE Asia Standard Time |
| Asia/Beirut | LB | ✅ Middle East Standard Time |
| Asia/Bishkek | KG | ✅ Central Asia Standard Time |
| Asia/Brunei | BN | ✅ Singapore Standard Time |
| Asia/Calcutta | IN | ✅ India Standard Time |
| Asia/Chongqing | CN | ❌ Not found |
| Asia/Colombo | LK | ✅ Sri Lanka Standard Time |
| Asia/Damascus | SY | ✅ Syria Standard Time |
| Asia/Dhaka | BD | ✅ Bangladesh Standard Time |
| Asia/Dili | TL | ✅ Tokyo Standard Time |
| Asia/Dubai | AE | ✅ Arabian Standard Time |
| Asia/Dushanbe | TJ | ✅ West Asia Standard Time |
| Asia/Gaza | PS | ✅ West Bank Standard Time |
| Asia/Harbin | CN | ❌ Not found |
| Asia/Hong_Kong | HK | ✅ China Standard Time |
| Asia/Hovd | MN | ✅ W. Mongolia Standard Time |
| Asia/Irkutsk | RU | ✅ North Asia East Standard Time |
| Asia/Jakarta | ID | ✅ SE Asia Standard Time |
| Asia/Jayapura | ID | ✅ Tokyo Standard Time |
| Asia/Jerusalem | IL | ✅ Israel Standard Time |
| Asia/Kabul | AF | ✅ Afghanistan Standard Time |
| Asia/Kamchatka | RU | ✅ Russia Time Zone 11 |
| Asia/Karachi | PK | ✅ Pakistan Standard Time |
| Asia/Kashgar | CN | ❌ Not found |
| Asia/Katmandu | NP | ✅ Nepal Standard Time |
| Asia/Krasnoyarsk | RU | ✅ North Asia Standard Time |
| Asia/Kuala_Lumpur | MY | ✅ Singapore Standard Time |
| Asia/Kuching | MY | ✅ Singapore Standard Time |
| Asia/Kuwait | KW | ✅ Arab Standard Time |
| Asia/Macau | MO | ✅ China Standard Time |
| Asia/Magadan | RU | ✅ Magadan Standard Time |
| Asia/Makassar | ID | ✅ Singapore Standard Time |
| Asia/Manila | PH | ✅ Singapore Standard Time |
| Asia/Muscat | OM | ✅ Arabian Standard Time |
| Asia/Nicosia | CY | ✅ GTB Standard Time |
| Asia/Novosibirsk | RU | ✅ N. Central Asia Standard Time |
| Asia/Omsk | RU | ✅ Omsk Standard Time |
| Asia/Oral | KZ | ✅ West Asia Standard Time |
| Asia/Phnom_Penh | KH | ✅ SE Asia Standard Time |
| Asia/Pontianak | ID | ✅ SE Asia Standard Time |
| Asia/Pyongyang | KP | ✅ North Korea Standard Time |
| Asia/Qatar | QA | ✅ Arab Standard Time |
| Asia/Qostanay | KZ | ✅ West Asia Standard Time |
| Asia/Qyzylorda | KZ | ✅ Qyzylorda Standard Time |
| Asia/Rangoon | MM | ✅ Myanmar Standard Time |
| Asia/Riyadh | SA | ✅ Arab Standard Time |
| Asia/Saigon | VN | ✅ SE Asia Standard Time |
| Asia/Sakhalin | RU | ✅ Sakhalin Standard Time |
| Asia/Samarkand | UZ | ✅ West Asia Standard Time |
| Asia/Seoul | KR | ✅ Korea Standard Time |
| Asia/Shanghai | CN | ✅ China Standard Time |
| Asia/Singapore | SG | ✅ Singapore Standard Time |
| Asia/Taipei | TW | ✅ Taipei Standard Time |
| Asia/Tashkent | UZ | ✅ West Asia Standard Time |
| Asia/Tbilisi | GE | ✅ Georgian Standard Time |
| Asia/Tehran | IR | ✅ Iran Standard Time |
| Asia/Thimphu | BT | ✅ Bangladesh Standard Time |
| Asia/Tokyo | JP | ✅ Tokyo Standard Time |
| Asia/Ulaanbaatar | MN | ✅ Ulaanbaatar Standard Time |
| Asia/Urumqi | CN | ✅ Central Asia Standard Time |
| Asia/Vientiane | LA | ✅ SE Asia Standard Time |
| Asia/Vladivostok | RU | ✅ Vladivostok Standard Time |
| Asia/Yakutsk | RU | ✅ Yakutsk Standard Time |
| Asia/Yekaterinburg | RU | ✅ Ekaterinburg Standard Time |
| Asia/Yerevan | AM | ✅ Caucasus Standard Time |
| Atlantic/Azores | PT | ✅ Azores Standard Time |
| Atlantic/Bermuda | BM | ✅ Atlantic Standard Time |
| Atlantic/Canary | ES | ✅ GMT Standard Time |
| Atlantic/Cape_Verde | CV | ✅ Cape Verde Standard Time |
| Atlantic/Faeroe | FO | ✅ GMT Standard Time |
| Atlantic/Jan_Mayen | SJ | ❌ Not found |
| Atlantic/Madeira | PT | ✅ GMT Standard Time |
| Atlantic/Reykjavik | IS | ✅ Greenwich Standard Time |
| Atlantic/South_Georgia | GS | ✅ UTC-02 |
| Atlantic/St_Helena | SH | ✅ Greenwich Standard Time |
| Atlantic/Stanley | FK | ✅ SA Eastern Standard Time |
| Australia/Adelaide | AU | ✅ Cen. Australia Standard Time |
| Australia/Brisbane | AU | ✅ E. Australia Standard Time |
| Australia/Broken_Hill | AU | ✅ Cen. Australia Standard Time |
| Australia/Darwin | AU | ✅ AUS Central Standard Time |
| Australia/Hobart | AU | ✅ Tasmania Standard Time |
| Australia/Lindeman | AU | ✅ E. Australia Standard Time |
| Australia/Lord_Howe | AU | ✅ Lord Howe Standard Time |
| Australia/Melbourne | AU | ✅ AUS Eastern Standard Time |
| Australia/Perth | AU | ✅ W. Australia Standard Time |
| Australia/Sydney | AU | ✅ AUS Eastern Standard Time |
| Etc/GMT | ZZ | ✅ UTC |
| Etc/GMT+1 | ZZ | ✅ Cape Verde Standard Time |
| Etc/GMT+10 | ZZ | ✅ Hawaiian Standard Time |
| Etc/GMT+11 | ZZ | ✅ UTC-11 |
| Etc/GMT+12 | ZZ | ✅ Dateline Standard Time |
| Etc/GMT+2 | ZZ | ✅ UTC-02 |
| Etc/GMT+3 | ZZ | ✅ SA Eastern Standard Time |
| Etc/GMT+4 | ZZ | ✅ SA Western Standard Time |
| Etc/GMT+5 | ZZ | ✅ SA Pacific Standard Time |
| Etc/GMT+6 | ZZ | ✅ Central America Standard Time |
| Etc/GMT+7 | ZZ | ✅ US Mountain Standard Time |
| Etc/GMT+8 | ZZ | ✅ UTC-08 |
| Etc/GMT+9 | ZZ | ✅ UTC-09 |
| Etc/GMT-1 | ZZ | ✅ W. Central Africa Standard Time |
| Etc/GMT-10 | ZZ | ✅ West Pacific Standard Time |
| Etc/GMT-11 | ZZ | ✅ Central Pacific Standard Time |
| Etc/GMT-12 | ZZ | ✅ UTC+12 |
| Etc/GMT-13 | ZZ | ✅ UTC+13 |
| Etc/GMT-14 | ZZ | ✅ Line Islands Standard Time |
| Etc/GMT-2 | ZZ | ✅ South Africa Standard Time |
| Etc/GMT-3 | ZZ | ✅ E. Africa Standard Time |
| Etc/GMT-4 | ZZ | ✅ Arabian Standard Time |
| Etc/GMT-5 | ZZ | ✅ West Asia Standard Time |
| Etc/GMT-6 | ZZ | ✅ Central Asia Standard Time |
| Etc/GMT-7 | ZZ | ✅ SE Asia Standard Time |
| Etc/GMT-8 | ZZ | ✅ Singapore Standard Time |
| Etc/GMT-9 | ZZ | ✅ Tokyo Standard Time |
| Etc/UCT | ZZ | ❌ Not found |
| Etc/UTC | ZZ | ✅ UTC |
| Europe/Amsterdam | NL | ✅ W. Europe Standard Time |
| Europe/Andorra | AD | ✅ W. Europe Standard Time |
| Europe/Athens | GR | ✅ GTB Standard Time |
| Europe/Belfast | GB | ❌ Not found |
| Europe/Belgrade | YU | ✅ Central Europe Standard Time |
| Europe/Berlin | DE | ✅ W. Europe Standard Time |
| Europe/Bratislava | SK | ✅ Central Europe Standard Time |
| Europe/Brussels | BE | ✅ Romance Standard Time |
| Europe/Bucharest | RO | ✅ GTB Standard Time |
| Europe/Budapest | HU | ✅ Central Europe Standard Time |
| Europe/Chisinau | MD | ✅ E. Europe Standard Time |
| Europe/Copenhagen | DK | ✅ Romance Standard Time |
| Europe/Dublin | IE | ✅ GMT Standard Time |
| Europe/Gibraltar | GI | ✅ W. Europe Standard Time |
| Europe/Helsinki | FI | ✅ FLE Standard Time |
| Europe/Istanbul | TR | ✅ Turkey Standard Time |
| Europe/Kaliningrad | RU | ✅ Kaliningrad Standard Time |
| Europe/Kiev | UA | ✅ FLE Standard Time |
| Europe/Lisbon | PT | ✅ GMT Standard Time |
| Europe/Ljubljana | SI | ✅ Central Europe Standard Time |
| Europe/London | GB | ✅ GMT Standard Time |
| Europe/Luxembourg | LU | ✅ W. Europe Standard Time |
| Europe/Madrid | ES | ✅ Romance Standard Time |
| Europe/Malta | MT | ✅ W. Europe Standard Time |
| Europe/Minsk | BY | ✅ Belarus Standard Time |
| Europe/Monaco | MC | ✅ W. Europe Standard Time |
| Europe/Moscow | RU | ✅ Russian Standard Time |
| Europe/Oslo | NO | ✅ W. Europe Standard Time |
| Europe/Paris | FR | ✅ Romance Standard Time |
| Europe/Prague | CZ | ✅ Central Europe Standard Time |
| Europe/Riga | LV | ✅ FLE Standard Time |
| Europe/Rome | IT | ✅ W. Europe Standard Time |
| Europe/Samara | RU | ✅ Russia Time Zone 3 |
| Europe/San_Marino | SM | ✅ W. Europe Standard Time |
| Europe/Sarajevo | BA | ✅ Central European Standard Time |
| Europe/Simferopol | UA | ✅ Russian Standard Time |
| Europe/Skopje | MK | ✅ Central European Standard Time |
| Europe/Sofia | BG | ✅ FLE Standard Time |
| Europe/Stockholm | SE | ✅ W. Europe Standard Time |
| Europe/Tallinn | EE | ✅ FLE Standard Time |
| Europe/Tirane | AL | ✅ Central Europe Standard Time |
| Europe/Uzhgorod | UA | ❌ Not found |
| Europe/Vaduz | LI | ✅ W. Europe Standard Time |
| Europe/Vatican | VA | ✅ W. Europe Standard Time |
| Europe/Vienna | AT | ✅ W. Europe Standard Time |
| Europe/Vilnius | LT | ✅ FLE Standard Time |
| Europe/Warsaw | PL | ✅ Central European Standard Time |
| Europe/Zagreb | HR | ✅ Central European Standard Time |
| Europe/Zaporozhye | UA | ❌ Not found |
| Europe/Zurich | CH | ✅ W. Europe Standard Time |
| Indian/Antananarivo | MG | ✅ E. Africa Standard Time |
| Indian/Chagos | IO | ✅ Central Asia Standard Time |
| Indian/Christmas | CX | ✅ SE Asia Standard Time |
| Indian/Cocos | CC | ✅ Myanmar Standard Time |
| Indian/Comoro | KM | ✅ E. Africa Standard Time |
| Indian/Kerguelen | TF | ✅ West Asia Standard Time |
| Indian/Mahe | SC | ✅ Mauritius Standard Time |
| Indian/Maldives | MV | ✅ West Asia Standard Time |
| Indian/Mauritius | MU | ✅ Mauritius Standard Time |
| Indian/Mayotte | YT | ✅ E. Africa Standard Time |
| Indian/Reunion | RE | ✅ Mauritius Standard Time |
| NONE | BV | ❌ Not found |
| NONE2 | HM | ❌ Not found |
| Pacific/Apia | WS | ✅ Samoa Standard Time |
| Pacific/Auckland | NZ | ✅ New Zealand Standard Time |
| Pacific/Chatham | NZ | ✅ Chatham Islands Standard Time |
| Pacific/Easter | CL | ✅ Easter Island Standard Time |
| Pacific/Efate | VU | ✅ Central Pacific Standard Time |
| Pacific/Enderbury | KI | ✅ UTC+13 |
| Pacific/Fakaofo | TK | ✅ UTC+13 |
| Pacific/Fiji | FJ | ✅ Fiji Standard Time |
| Pacific/Funafuti | TV | ✅ UTC+12 |
| Pacific/Galapagos | EC | ✅ Central America Standard Time |
| Pacific/Gambier | PF | ✅ UTC-09 |
| Pacific/Guadalcanal | SB | ✅ Central Pacific Standard Time |
| Pacific/Guam | GU | ✅ West Pacific Standard Time |
| Pacific/Honolulu | US | ✅ Hawaiian Standard Time |
| Pacific/Johnston | UM | ❌ Not found |
| Pacific/Kiritimati | KI | ✅ Line Islands Standard Time |
| Pacific/Kosrae | FM | ✅ Central Pacific Standard Time |
| Pacific/Kwajalein | MH | ✅ UTC+12 |
| Pacific/Majuro | MH | ✅ UTC+12 |
| Pacific/Marquesas | PF | ✅ Marquesas Standard Time |
| Pacific/Midway | UM | ✅ UTC-11 |
| Pacific/Nauru | NR | ✅ UTC+12 |
| Pacific/Niue | NU | ✅ UTC-11 |
| Pacific/Norfolk | NF | ✅ Norfolk Standard Time |
| Pacific/Noumea | NC | ✅ Central Pacific Standard Time |
| Pacific/Pago_Pago | AS | ✅ UTC-11 |
| Pacific/Palau | PW | ✅ Tokyo Standard Time |
| Pacific/Pitcairn | PN | ✅ UTC-08 |
| Pacific/Ponape | FM | ✅ Central Pacific Standard Time |
| Pacific/Port_Moresby | PG | ✅ West Pacific Standard Time |
| Pacific/Rarotonga | CK | ✅ Hawaiian Standard Time |
| Pacific/Saipan | MP | ✅ West Pacific Standard Time |
| Pacific/Tahiti | PF | ✅ Hawaiian Standard Time |
| Pacific/Tarawa | KI | ✅ UTC+12 |
| Pacific/Tongatapu | TO | ✅ Tonga Standard Time |
| Pacific/Truk | FM | ✅ West Pacific Standard Time |
| Pacific/Wake | UM | ✅ UTC+12 |
| Pacific/Wallis | WF | ✅ UTC+12 |
| Pacific/Yap | FM | ❌ Not found |


## Missing Microsoft Aliases

The following Microsoft Aliases are reported by Microsoft Graph API as supported but are missing in the current conversion implementation.

- Mid-Atlantic Standard Time
- Kamchatka Standard Time


## Missing Microsoft Display Names

The following Microsoft Display Names are reported by Microsoft Graph API as supported but are missing in the current conversion implementation.

- (UTC-07:00) La Paz, Mazatlan
- (UTC-03:00) Asuncion
- (UTC-02:00) Greenland
- (UTC-02:00) Mid-Atlantic - Old
- (UTC+03:00) Damascus
- (UTC+03:00) Amman
- (UTC+03:00) Volgograd
- (UTC+05:00) Astana
- (UTC+06:00) Bishkek
- (UTC+12:00) Petropavlovsk-Kamchatsky - Old


## Missing IANA Timezones

The following IANA timezones from CLDR are not found in winianatz:

- Africa/Timbuktu
- America/Nipigon
- America/Pangnirtung
- America/Rainy_River
- America/Thunder_Bay
- America/Yellowknife
- Asia/Chongqing
- Asia/Harbin
- Asia/Kashgar
- Atlantic/Jan_Mayen
- Etc/UCT
- Europe/Belfast
- Europe/Uzhgorod
- Europe/Zaporozhye
- NONE
- NONE2
- Pacific/Johnston
- Pacific/Yap
