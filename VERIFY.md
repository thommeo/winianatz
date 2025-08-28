# Timezone conversion verification report

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
