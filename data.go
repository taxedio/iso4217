package iso4217

type Country struct {
	Name   string
	Alpha2 string
	Alpha3 string
}

type ISOEntry struct {
	countryName  []Country
	currencyName string
	numCode      string
	minorUnit    string
}

var (
	isocodes = map[string]ISOEntry{
		"EUR": {
			countryName: []Country{{
				Name:   "åland islands",
				Alpha2: "AX",
				Alpha3: "ALA",
			}, {
				Name:   "andorra",
				Alpha2: "AD",
				Alpha3: "AND",
			}, {
				Name:   "austria",
				Alpha2: "AT",
				Alpha3: "AUT",
			}, {
				Name:   "belgium",
				Alpha2: "BE",
				Alpha3: "BEL",
			}, {
				Name:   "cyprus",
				Alpha2: "CY",
				Alpha3: "CYP",
			}, {
				Name:   "estonia",
				Alpha2: "EE",
				Alpha3: "EST",
			}, {
				Name:   "european union",
				Alpha2: "XX",
				Alpha3: "XXX",
			}, {
				Name:   "finland",
				Alpha2: "FI",
				Alpha3: "FIN",
			}, {
				Name:   "france",
				Alpha2: "FR",
				Alpha3: "FRA",
			}, {
				Name:   "french guiana",
				Alpha2: "GF",
				Alpha3: "GUF",
			}, {
				Name:   "french southern territories (the)",
				Alpha2: "TF",
				Alpha3: "ATF",
			}, {
				Name:   "germany",
				Alpha2: "DE",
				Alpha3: "DEU",
			}, {
				Name:   "greece",
				Alpha2: "GR",
				Alpha3: "GRC",
			}, {
				Name:   "guadeloupe",
				Alpha2: "GP",
				Alpha3: "GLP",
			}, {
				Name:   "holy see (the)",
				Alpha2: "VA",
				Alpha3: "VAT",
			}, {
				Name:   "ireland",
				Alpha2: "IE",
				Alpha3: "IRL",
			}, {
				Name:   "italy",
				Alpha2: "IT",
				Alpha3: "ITA",
			}, {
				Name:   "latvia",
				Alpha2: "LV",
				Alpha3: "LVA",
			}, {
				Name:   "lithuania",
				Alpha2: "LT",
				Alpha3: "LTU",
			}, {
				Name:   "luxembourg",
				Alpha2: "LU",
				Alpha3: "LUX",
			}, {
				Name:   "malta",
				Alpha2: "MT",
				Alpha3: "MLT",
			}, {
				Name:   "martinique",
				Alpha2: "MQ",
				Alpha3: "MTQ",
			}, {
				Name:   "mayotte",
				Alpha2: "YT",
				Alpha3: "MYT",
			}, {
				Name:   "monaco",
				Alpha2: "MC",
				Alpha3: "MCO",
			}, {
				Name:   "montenegro",
				Alpha2: "ME",
				Alpha3: "MNE",
			}, {
				Name:   "netherlands (the)",
				Alpha2: "NL",
				Alpha3: "NLD",
			}, {
				Name:   "portugal",
				Alpha2: "PT",
				Alpha3: "PRT",
			}, {
				Name:   "réunion",
				Alpha2: "RE",
				Alpha3: "REU",
			}, {
				Name:   "saint barthélemy",
				Alpha2: "BL",
				Alpha3: "BLM",
			}, {
				Name:   "saint martin (french part)",
				Alpha2: "MF",
				Alpha3: "MAF",
			}, {
				Name:   "saint pierre and miquelon",
				Alpha2: "PM",
				Alpha3: "SPM",
			}, {
				Name:   "san marino",
				Alpha2: "SM",
				Alpha3: "SMR",
			}, {
				Name:   "slovakia",
				Alpha2: "SK",
				Alpha3: "SVK",
			}, {
				Name:   "slovenia",
				Alpha2: "SI",
				Alpha3: "SVN",
			}, {
				Name:   "spain",
				Alpha2: "ES",
				Alpha3: "ESP",
			},
			},
			currencyName: "Euro",
			numCode:      "978",
			minorUnit:    "2",
		}, "ALL": {
			countryName: []Country{{
				Name:   "albania",
				Alpha2: "AL",
				Alpha3: "ALB",
			}},
			currencyName: "Lek",
			numCode:      "008",
			minorUnit:    "2",
		}, "DZD": {
			countryName: []Country{{
				Name:   "algeria",
				Alpha2: "DZ",
				Alpha3: "DZA",
			}},
			currencyName: "Algerian Dinar",
			numCode:      "012",
			minorUnit:    "2",
		}, "USD": {
			countryName: []Country{{
				Name:   "american samoa",
				Alpha2: "AS",
				Alpha3: "ASM",
			}, {
				Name:   "bonaire, sint eustatius and saba",
				Alpha2: "XX",
				Alpha3: "XXX",
			}, {
				Name:   "british indian ocean territory (the)",
				Alpha2: "IO",
				Alpha3: "IOT",
			}, {
				Name:   "ecuador",
				Alpha2: "EC",
				Alpha3: "ECU",
			}, {
				Name:   "el salvador",
				Alpha2: "SV",
				Alpha3: "SLV",
			}, {
				Name:   "guam",
				Alpha2: "GU",
				Alpha3: "GUM",
			}, {
				Name:   "haiti",
				Alpha2: "HT",
				Alpha3: "HTI",
			}, {
				Name:   "marshall islands (the)",
				Alpha2: "MH",
				Alpha3: "MHL",
			}, {
				Name:   "micronesia (federated states of)",
				Alpha2: "FM",
				Alpha3: "FSM",
			}, {
				Name:   "northern mariana islands (the)",
				Alpha2: "MP",
				Alpha3: "MNP",
			}, {
				Name:   "palau",
				Alpha2: "PW",
				Alpha3: "PLW",
			}, {
				Name:   "panama",
				Alpha2: "PA",
				Alpha3: "PAN",
			}, {
				Name:   "puerto rico",
				Alpha2: "PR",
				Alpha3: "PRI",
			}, {
				Name:   "timor-leste",
				Alpha2: "TL",
				Alpha3: "TLS",
			}, {
				Name:   "turks and caicos islands (the)",
				Alpha2: "TC",
				Alpha3: "TCA",
			}, {
				Name:   "united states minor outlying islands (the)",
				Alpha2: "UM",
				Alpha3: "UMI",
			}, {
				Name:   "united states of america (the)",
				Alpha2: "US",
				Alpha3: "USA",
			}, {
				Name:   "virgin islands (british)",
				Alpha2: "VG",
				Alpha3: "VGB",
			}, {
				Name:   "virgin islands (u.s.)",
				Alpha2: "VI",
				Alpha3: "VIR",
			}},
			currencyName: "US Dollar",
			numCode:      "840",
			minorUnit:    "2",
		}, "AOA": {
			countryName: []Country{{
				Name:   "angola",
				Alpha2: "AO",
				Alpha3: "AGO",
			}},
			currencyName: "Kwanza",
			numCode:      "973",
			minorUnit:    "2",
		}, "XCD": {
			countryName: []Country{{
				Name:   "anguilla",
				Alpha2: "AI",
				Alpha3: "AIA",
			}, {
				Name:   "antigua and barbuda",
				Alpha2: "AG",
				Alpha3: "ATG",
			}, {
				Name:   "dominica",
				Alpha2: "DM",
				Alpha3: "DMA",
			}, {
				Name:   "grenada",
				Alpha2: "GD",
				Alpha3: "GRD",
			}, {
				Name:   "montserrat",
				Alpha2: "MS",
				Alpha3: "MSR",
			}, {
				Name:   "saint kitts and nevis",
				Alpha2: "KN",
				Alpha3: "KNA",
			}, {
				Name:   "saint lucia",
				Alpha2: "LC",
				Alpha3: "LCA",
			}, {
				Name:   "saint vincent and the grenadines",
				Alpha2: "VC",
				Alpha3: "VCT",
			}},
			currencyName: "East Caribbean Dollar",
			numCode:      "951",
			minorUnit:    "2",
		}, "UNKNOWN": {
			countryName: []Country{{
				Name:   "antarctica",
				Alpha2: "AQ",
				Alpha3: "ATA",
			}, {
				Name:   "palestine, state of",
				Alpha2: "PS",
				Alpha3: "PSE",
			}, {
				Name:   "south georgia and the south sandwich islands",
				Alpha2: "GS",
				Alpha3: "SGS",
			}},
			currencyName: "No universal currency",
			numCode:      "UNKNOWN",
			minorUnit:    "UNKNOWN",
		}, "ARS": {
			countryName: []Country{{
				Name:   "argentina",
				Alpha2: "AR",
				Alpha3: "ARG",
			}},
			currencyName: "Argentine Peso",
			numCode:      "032",
			minorUnit:    "2",
		}, "AMD": {
			countryName: []Country{{
				Name:   "armenia",
				Alpha2: "AM",
				Alpha3: "ARM",
			}},
			currencyName: "Armenian Dram",
			numCode:      "051",
			minorUnit:    "2",
		}, "AWG": {
			countryName: []Country{{
				Name:   "aruba",
				Alpha2: "AW",
				Alpha3: "ABW",
			}},
			currencyName: "Aruban Florin",
			numCode:      "533",
			minorUnit:    "2",
		}, "AUD": {
			countryName: []Country{{
				Name:   "australia",
				Alpha2: "AU",
				Alpha3: "AUS",
			}, {
				Name:   "christmas island",
				Alpha2: "CX",
				Alpha3: "CXR",
			}, {
				Name:   "cocos (keeling) islands (the)",
				Alpha2: "CC",
				Alpha3: "CCK",
			}, {
				Name:   "heard island and mcdonald islands",
				Alpha2: "HM",
				Alpha3: "HMD",
			}, {
				Name:   "kiribati",
				Alpha2: "KI",
				Alpha3: "KIR",
			}, {
				Name:   "nauru",
				Alpha2: "NR",
				Alpha3: "NRU",
			}, {
				Name:   "norfolk island",
				Alpha2: "NF",
				Alpha3: "NFK",
			}, {
				Name:   "tuvalu",
				Alpha2: "TV",
				Alpha3: "TUV",
			}},
			currencyName: "Australian Dollar",
			numCode:      "036",
			minorUnit:    "2",
		}, "AZN": {
			countryName: []Country{{
				Name:   "azerbaijan",
				Alpha2: "AZ",
				Alpha3: "AZE",
			}},
			currencyName: "Azerbaijan Manat",
			numCode:      "944",
			minorUnit:    "2",
		}, "BSD": {
			countryName: []Country{{
				Name:   "bahamas (the)",
				Alpha2: "BS",
				Alpha3: "BHS",
			}},
			currencyName: "Bahamian Dollar",
			numCode:      "044",
			minorUnit:    "2",
		}, "BHD": {
			countryName: []Country{{
				Name:   "bahrain",
				Alpha2: "BH",
				Alpha3: "BHR",
			}},
			currencyName: "Bahraini Dinar",
			numCode:      "048",
			minorUnit:    "3",
		}, "BDT": {
			countryName: []Country{{
				Name:   "bangladesh",
				Alpha2: "BD",
				Alpha3: "BGD",
			}},
			currencyName: "Taka",
			numCode:      "050",
			minorUnit:    "2",
		}, "BBD": {
			countryName: []Country{{
				Name:   "barbados",
				Alpha2: "BB",
				Alpha3: "BRB",
			}},
			currencyName: "Barbados Dollar",
			numCode:      "052",
			minorUnit:    "2",
		}, "BYN": {
			countryName: []Country{{
				Name:   "belarus",
				Alpha2: "BY",
				Alpha3: "BLR",
			}},
			currencyName: "Belarusian Ruble",
			numCode:      "933",
			minorUnit:    "2",
		}, "BZD": {
			countryName: []Country{{
				Name:   "belize",
				Alpha2: "BZ",
				Alpha3: "BLZ",
			}},
			currencyName: "Belize Dollar",
			numCode:      "084",
			minorUnit:    "2",
		}, "XOF": {
			countryName: []Country{{
				Name:   "benin",
				Alpha2: "BJ",
				Alpha3: "BEN",
			}, {
				Name:   "burkina faso",
				Alpha2: "BF",
				Alpha3: "BFA",
			}, {
				Name:   "côte d'ivoire",
				Alpha2: "CI",
				Alpha3: "CIV",
			}, {
				Name:   "guinea-bissau",
				Alpha2: "GW",
				Alpha3: "GNB",
			}, {
				Name:   "mali",
				Alpha2: "ML",
				Alpha3: "MLI",
			}, {
				Name:   "niger (the)",
				Alpha2: "NE",
				Alpha3: "NER",
			}, {
				Name:   "senegal",
				Alpha2: "SN",
				Alpha3: "SEN",
			}, {
				Name:   "togo",
				Alpha2: "TG",
				Alpha3: "TGO",
			}},
			currencyName: "CFA Franc BCEAO",
			numCode:      "952",
			minorUnit:    "0",
		}, "BMD": {
			countryName: []Country{{
				Name:   "bermuda",
				Alpha2: "BM",
				Alpha3: "BMU",
			}},
			currencyName: "Bermudian Dollar",
			numCode:      "060",
			minorUnit:    "2",
		}, "INR": {
			countryName: []Country{{
				Name:   "bhutan",
				Alpha2: "BT",
				Alpha3: "BTN",
			}, {
				Name:   "india",
				Alpha2: "IN",
				Alpha3: "IND",
			}},
			currencyName: "Indian Rupee",
			numCode:      "356",
			minorUnit:    "2",
		}, "BTN": {
			countryName: []Country{{
				Name:   "bhutan",
				Alpha2: "BT",
				Alpha3: "BTN",
			}},
			currencyName: "Ngultrum",
			numCode:      "064",
			minorUnit:    "2",
		}, "BOB": {
			countryName: []Country{{
				Name:   "bolivia (plurinational state of)",
				Alpha2: "BO",
				Alpha3: "BOL",
			}},
			currencyName: "Boliviano",
			numCode:      "068",
			minorUnit:    "2",
		}, "BOV": {
			countryName: []Country{{
				Name:   "bolivia (plurinational state of)",
				Alpha2: "BO",
				Alpha3: "BOL",
			}},
			currencyName: "Mvdol",
			numCode:      "984",
			minorUnit:    "2",
		}, "BAM": {
			countryName: []Country{{
				Name:   "bosnia and herzegovina",
				Alpha2: "BA",
				Alpha3: "BIH",
			}},
			currencyName: "Convertible Mark",
			numCode:      "977",
			minorUnit:    "2",
		}, "BWP": {
			countryName: []Country{{
				Name:   "botswana",
				Alpha2: "BW",
				Alpha3: "BWA",
			}},
			currencyName: "Pula",
			numCode:      "072",
			minorUnit:    "2",
		}, "NOK": {
			countryName: []Country{{
				Name:   "bouvet island",
				Alpha2: "BV",
				Alpha3: "BVT",
			}, {
				Name:   "norway",
				Alpha2: "NO",
				Alpha3: "NOR",
			}, {
				Name:   "svalbard and jan mayen",
				Alpha2: "SJ",
				Alpha3: "SJM",
			}},
			currencyName: "Norwegian Krone",
			numCode:      "578",
			minorUnit:    "2",
		}, "BRL": {
			countryName: []Country{{
				Name:   "brazil",
				Alpha2: "BR",
				Alpha3: "BRA",
			}},
			currencyName: "Brazilian Real",
			numCode:      "986",
			minorUnit:    "2",
		}, "BND": {
			countryName: []Country{{
				Name:   "brunei darussalam",
				Alpha2: "BN",
				Alpha3: "BRN",
			}},
			currencyName: "Brunei Dollar",
			numCode:      "096",
			minorUnit:    "2",
		}, "BGN": {
			countryName: []Country{{
				Name:   "bulgaria",
				Alpha2: "BG",
				Alpha3: "BGR",
			}},
			currencyName: "Bulgarian Lev",
			numCode:      "975",
			minorUnit:    "2",
		}, "BIF": {
			countryName: []Country{{
				Name:   "burundi",
				Alpha2: "BI",
				Alpha3: "BDI",
			}},
			currencyName: "Burundi Franc",
			numCode:      "108",
			minorUnit:    "0",
		}, "CVE": {
			countryName: []Country{{
				Name:   "cabo verde",
				Alpha2: "CV",
				Alpha3: "CPV",
			}},
			currencyName: "Cabo Verde Escudo",
			numCode:      "132",
			minorUnit:    "2",
		}, "KHR": {
			countryName: []Country{{
				Name:   "cambodia",
				Alpha2: "KH",
				Alpha3: "KHM",
			}},
			currencyName: "Riel",
			numCode:      "116",
			minorUnit:    "2",
		}, "XAF": {
			countryName: []Country{{
				Name:   "cameroon",
				Alpha2: "CM",
				Alpha3: "CMR",
			}, {
				Name:   "central african republic (the)",
				Alpha2: "CF",
				Alpha3: "CAF",
			}, {
				Name:   "chad",
				Alpha2: "TD",
				Alpha3: "TCD",
			}, {
				Name:   "congo (the)",
				Alpha2: "CG",
				Alpha3: "COG",
			}, {
				Name:   "equatorial guinea",
				Alpha2: "GQ",
				Alpha3: "GNQ",
			}, {
				Name:   "gabon",
				Alpha2: "GA",
				Alpha3: "GAB",
			}},
			currencyName: "CFA Franc BEAC",
			numCode:      "950",
			minorUnit:    "0",
		}, "CAD": {
			countryName: []Country{{
				Name:   "canada",
				Alpha2: "CA",
				Alpha3: "CAN",
			}},
			currencyName: "Canadian Dollar",
			numCode:      "124",
			minorUnit:    "2",
		}, "KYD": {
			countryName: []Country{{
				Name:   "cayman islands (the)",
				Alpha2: "KY",
				Alpha3: "CYM",
			}},
			currencyName: "Cayman Islands Dollar",
			numCode:      "136",
			minorUnit:    "2",
		}, "CLP": {
			countryName: []Country{{
				Name:   "chile",
				Alpha2: "CL",
				Alpha3: "CHL",
			}},
			currencyName: "Chilean Peso",
			numCode:      "152",
			minorUnit:    "0",
		}, "CLF": {
			countryName: []Country{{
				Name:   "chile",
				Alpha2: "CL",
				Alpha3: "CHL",
			}},
			currencyName: "Unidad de Fomento",
			numCode:      "990",
			minorUnit:    "4",
		}, "CNY": {
			countryName: []Country{{
				Name:   "china",
				Alpha2: "CN",
				Alpha3: "CHN",
			}},
			currencyName: "Yuan Renminbi",
			numCode:      "156",
			minorUnit:    "2",
		}, "COP": {
			countryName: []Country{{
				Name:   "colombia",
				Alpha2: "CO",
				Alpha3: "COL",
			}},
			currencyName: "Colombian Peso",
			numCode:      "170",
			minorUnit:    "2",
		}, "COU": {
			countryName: []Country{{
				Name:   "colombia",
				Alpha2: "CO",
				Alpha3: "COL",
			}},
			currencyName: "Unidad de Valor Real",
			numCode:      "970",
			minorUnit:    "2",
		}, "KMF": {
			countryName: []Country{{
				Name:   "comoros (the)",
				Alpha2: "KM",
				Alpha3: "COM",
			}},
			currencyName: "Comorian Franc ",
			numCode:      "174",
			minorUnit:    "0",
		}, "CDF": {
			countryName: []Country{{
				Name:   "congo (the democratic republic of the)",
				Alpha2: "CD",
				Alpha3: "COD",
			}},
			currencyName: "Congolese Franc",
			numCode:      "976",
			minorUnit:    "2",
		}, "NZD": {
			countryName: []Country{{
				Name:   "cook islands (the)",
				Alpha2: "CK",
				Alpha3: "COK",
			}, {
				Name:   "new zealand",
				Alpha2: "NZ",
				Alpha3: "NZL",
			}, {
				Name:   "niue",
				Alpha2: "NU",
				Alpha3: "NIU",
			}, {
				Name:   "pitcairn",
				Alpha2: "PN",
				Alpha3: "PCN",
			}, {
				Name:   "tokelau",
				Alpha2: "TK",
				Alpha3: "TKL",
			}},
			currencyName: "New Zealand Dollar",
			numCode:      "554",
			minorUnit:    "2",
		}, "CRC": {
			countryName: []Country{{
				Name:   "costa rica",
				Alpha2: "CR",
				Alpha3: "CRI",
			}},
			currencyName: "Costa Rican Colon",
			numCode:      "188",
			minorUnit:    "2",
		}, "HRK": {
			countryName: []Country{{
				Name:   "croatia",
				Alpha2: "HR",
				Alpha3: "HRV",
			}},
			currencyName: "Kuna",
			numCode:      "191",
			minorUnit:    "2",
		}, "CUP": {
			countryName: []Country{{
				Name:   "cuba",
				Alpha2: "CU",
				Alpha3: "CUB",
			}},
			currencyName: "Cuban Peso",
			numCode:      "192",
			minorUnit:    "2",
		}, "CUC": {
			countryName: []Country{{
				Name:   "cuba",
				Alpha2: "CU",
				Alpha3: "CUB",
			}},
			currencyName: "Peso Convertible",
			numCode:      "931",
			minorUnit:    "2",
		}, "ANG": {
			countryName: []Country{{
				Name:   "curaçao",
				Alpha2: "CW",
				Alpha3: "CUW",
			}, {
				Name:   "sint maarten (dutch part)",
				Alpha2: "SX",
				Alpha3: "SXM",
			}},
			currencyName: "Netherlands Antillean Guilder",
			numCode:      "532",
			minorUnit:    "2",
		}, "CZK": {
			countryName: []Country{{
				Name:   "czechia",
				Alpha2: "CZ",
				Alpha3: "CZE",
			}},
			currencyName: "Czech Koruna",
			numCode:      "203",
			minorUnit:    "2",
		}, "DKK": {
			countryName: []Country{{
				Name:   "denmark",
				Alpha2: "DK",
				Alpha3: "DNK",
			}, {
				Name:   "faroe islands (the)",
				Alpha2: "FO",
				Alpha3: "FRO",
			}, {
				Name:   "greenland",
				Alpha2: "GL",
				Alpha3: "GRL",
			}},
			currencyName: "Danish Krone",
			numCode:      "208",
			minorUnit:    "2",
		}, "DJF": {
			countryName: []Country{{
				Name:   "djibouti",
				Alpha2: "DJ",
				Alpha3: "DJI",
			}},
			currencyName: "Djibouti Franc",
			numCode:      "262",
			minorUnit:    "0",
		}, "DOP": {
			countryName: []Country{{
				Name:   "dominican republic (the)",
				Alpha2: "DO",
				Alpha3: "DOM",
			}},
			currencyName: "Dominican Peso",
			numCode:      "214",
			minorUnit:    "2",
		}, "EGP": {
			countryName: []Country{{
				Name:   "egypt",
				Alpha2: "EG",
				Alpha3: "EGY",
			}},
			currencyName: "Egyptian Pound",
			numCode:      "818",
			minorUnit:    "2",
		}, "SVC": {
			countryName: []Country{{
				Name:   "el salvador",
				Alpha2: "SV",
				Alpha3: "SLV",
			}},
			currencyName: "El Salvador Colon",
			numCode:      "222",
			minorUnit:    "2",
		}, "ERN": {
			countryName: []Country{{
				Name:   "eritrea",
				Alpha2: "ER",
				Alpha3: "ERI",
			}},
			currencyName: "Nakfa",
			numCode:      "232",
			minorUnit:    "2",
		}, "SZL": {
			countryName: []Country{{
				Name:   "eswatini",
				Alpha2: "SZ",
				Alpha3: "SWZ",
			}},
			currencyName: "Lilangeni",
			numCode:      "748",
			minorUnit:    "2",
		}, "ETB": {
			countryName: []Country{{
				Name:   "ethiopia",
				Alpha2: "ET",
				Alpha3: "ETH",
			}},
			currencyName: "Ethiopian Birr",
			numCode:      "230",
			minorUnit:    "2",
		}, "FKP": {
			countryName: []Country{{
				Name:   "falkland islands (the) [malvinas]",
				Alpha2: "FK",
				Alpha3: "FLK",
			}},
			currencyName: "Falkland Islands Pound",
			numCode:      "238",
			minorUnit:    "2",
		}, "FJD": {
			countryName: []Country{{
				Name:   "fiji",
				Alpha2: "FJ",
				Alpha3: "FJI",
			}},
			currencyName: "Fiji Dollar",
			numCode:      "242",
			minorUnit:    "2",
		}, "XPF": {
			countryName: []Country{{
				Name:   "french polynesia",
				Alpha2: "PF",
				Alpha3: "PYF",
			}, {
				Name:   "new caledonia",
				Alpha2: "NC",
				Alpha3: "NCL",
			}, {
				Name:   "wallis and futuna",
				Alpha2: "WF",
				Alpha3: "WLF",
			}},
			currencyName: "CFP Franc",
			numCode:      "953",
			minorUnit:    "0",
		}, "GMD": {
			countryName: []Country{{
				Name:   "gambia (the)",
				Alpha2: "GM",
				Alpha3: "GMB",
			}},
			currencyName: "Dalasi",
			numCode:      "270",
			minorUnit:    "2",
		}, "GEL": {
			countryName: []Country{{
				Name:   "georgia",
				Alpha2: "GE",
				Alpha3: "GEO",
			}},
			currencyName: "Lari",
			numCode:      "981",
			minorUnit:    "2",
		}, "GHS": {
			countryName: []Country{{
				Name:   "ghana",
				Alpha2: "GH",
				Alpha3: "GHA",
			}},
			currencyName: "Ghana Cedi",
			numCode:      "936",
			minorUnit:    "2",
		}, "GIP": {
			countryName: []Country{{
				Name:   "gibraltar",
				Alpha2: "GI",
				Alpha3: "GIB",
			}},
			currencyName: "Gibraltar Pound",
			numCode:      "292",
			minorUnit:    "2",
		}, "GTQ": {
			countryName: []Country{{
				Name:   "guatemala",
				Alpha2: "GT",
				Alpha3: "GTM",
			}},
			currencyName: "Quetzal",
			numCode:      "320",
			minorUnit:    "2",
		}, "GBP": {
			countryName: []Country{{
				Name:   "guernsey",
				Alpha2: "GG",
				Alpha3: "GGY",
			}, {
				Name:   "isle of man",
				Alpha2: "IM",
				Alpha3: "IMN",
			}, {
				Name:   "jersey",
				Alpha2: "JE",
				Alpha3: "JEY",
			}, {
				Name:   "united kingdom of great britain and northern ireland (the)",
				Alpha2: "GB",
				Alpha3: "GBR",
			}},
			currencyName: "Pound Sterling",
			numCode:      "826",
			minorUnit:    "2",
		}, "GNF": {
			countryName: []Country{{
				Name:   "guinea",
				Alpha2: "GN",
				Alpha3: "GIN",
			}},
			currencyName: "Guinean Franc",
			numCode:      "324",
			minorUnit:    "0",
		}, "GYD": {
			countryName: []Country{{
				Name:   "guyana",
				Alpha2: "GY",
				Alpha3: "GUY",
			}},
			currencyName: "Guyana Dollar",
			numCode:      "328",
			minorUnit:    "2",
		}, "HTG": {
			countryName: []Country{{
				Name:   "haiti",
				Alpha2: "HT",
				Alpha3: "HTI",
			}},
			currencyName: "Gourde",
			numCode:      "332",
			minorUnit:    "2",
		}, "HNL": {
			countryName: []Country{{
				Name:   "honduras",
				Alpha2: "HN",
				Alpha3: "HND",
			}},
			currencyName: "Lempira",
			numCode:      "340",
			minorUnit:    "2",
		}, "HKD": {
			countryName: []Country{{
				Name:   "hong kong",
				Alpha2: "HK",
				Alpha3: "HKG",
			}},
			currencyName: "Hong Kong Dollar",
			numCode:      "344",
			minorUnit:    "2",
		}, "HUF": {
			countryName: []Country{{
				Name:   "hungary",
				Alpha2: "HU",
				Alpha3: "HUN",
			}},
			currencyName: "Forint",
			numCode:      "348",
			minorUnit:    "2",
		}, "ISK": {
			countryName: []Country{{
				Name:   "iceland",
				Alpha2: "IS",
				Alpha3: "ISL",
			}},
			currencyName: "Iceland Krona",
			numCode:      "352",
			minorUnit:    "0",
		}, "IDR": {
			countryName: []Country{{
				Name:   "indonesia",
				Alpha2: "ID",
				Alpha3: "IDN",
			}},
			currencyName: "Rupiah",
			numCode:      "360",
			minorUnit:    "2",
		}, "XDR": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "SDR (Special Drawing Right)",
			numCode:      "960",
			minorUnit:    "N.A.",
		}, "IRR": {
			countryName: []Country{{
				Name:   "iran (islamic republic of)",
				Alpha2: "IR",
				Alpha3: "IRN",
			}},
			currencyName: "Iranian Rial",
			numCode:      "364",
			minorUnit:    "2",
		}, "IQD": {
			countryName: []Country{{
				Name:   "iraq",
				Alpha2: "IQ",
				Alpha3: "IRQ",
			}},
			currencyName: "Iraqi Dinar",
			numCode:      "368",
			minorUnit:    "3",
		}, "ILS": {
			countryName: []Country{{
				Name:   "israel",
				Alpha2: "IL",
				Alpha3: "ISR",
			}},
			currencyName: "New Israeli Sheqel",
			numCode:      "376",
			minorUnit:    "2",
		}, "JMD": {
			countryName: []Country{{
				Name:   "jamaica",
				Alpha2: "JM",
				Alpha3: "JAM",
			}},
			currencyName: "Jamaican Dollar",
			numCode:      "388",
			minorUnit:    "2",
		}, "JPY": {
			countryName: []Country{{
				Name:   "japan",
				Alpha2: "JP",
				Alpha3: "JPN",
			}},
			currencyName: "Yen",
			numCode:      "392",
			minorUnit:    "0",
		}, "JOD": {
			countryName: []Country{{
				Name:   "jordan",
				Alpha2: "JO",
				Alpha3: "JOR",
			}},
			currencyName: "Jordanian Dinar",
			numCode:      "400",
			minorUnit:    "3",
		}, "KZT": {
			countryName: []Country{{
				Name:   "kazakhstan",
				Alpha2: "KZ",
				Alpha3: "KAZ",
			}},
			currencyName: "Tenge",
			numCode:      "398",
			minorUnit:    "2",
		}, "KES": {
			countryName: []Country{{
				Name:   "kenya",
				Alpha2: "KE",
				Alpha3: "KEN",
			}},
			currencyName: "Kenyan Shilling",
			numCode:      "404",
			minorUnit:    "2",
		}, "KPW": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "North Korean Won",
			numCode:      "408",
			minorUnit:    "2",
		}, "KRW": {
			countryName: []Country{{
				Name:   "korea (the republic of)",
				Alpha2: "KR",
				Alpha3: "KOR",
			}},
			currencyName: "Won",
			numCode:      "410",
			minorUnit:    "0",
		}, "KWD": {
			countryName: []Country{{
				Name:   "kuwait",
				Alpha2: "KW",
				Alpha3: "KWT",
			}},
			currencyName: "Kuwaiti Dinar",
			numCode:      "414",
			minorUnit:    "3",
		}, "KGS": {
			countryName: []Country{{
				Name:   "kyrgyzstan",
				Alpha2: "KG",
				Alpha3: "KGZ",
			}},
			currencyName: "Som",
			numCode:      "417",
			minorUnit:    "2",
		}, "LAK": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "Lao Kip",
			numCode:      "418",
			minorUnit:    "2",
		}, "LBP": {
			countryName: []Country{{
				Name:   "lebanon",
				Alpha2: "LB",
				Alpha3: "LBN",
			}},
			currencyName: "Lebanese Pound",
			numCode:      "422",
			minorUnit:    "2",
		}, "LSL": {
			countryName: []Country{{
				Name:   "lesotho",
				Alpha2: "LS",
				Alpha3: "LSO",
			}},
			currencyName: "Loti",
			numCode:      "426",
			minorUnit:    "2",
		}, "ZAR": {
			countryName: []Country{{
				Name:   "lesotho",
				Alpha2: "LS",
				Alpha3: "LSO",
			}, {
				Name:   "namibia",
				Alpha2: "NA",
				Alpha3: "NAM",
			}, {
				Name:   "south africa",
				Alpha2: "ZA",
				Alpha3: "ZAF",
			}},
			currencyName: "Rand",
			numCode:      "710",
			minorUnit:    "2",
		}, "LRD": {
			countryName: []Country{{
				Name:   "liberia",
				Alpha2: "LR",
				Alpha3: "LBR",
			}},
			currencyName: "Liberian Dollar",
			numCode:      "430",
			minorUnit:    "2",
		}, "LYD": {
			countryName: []Country{{
				Name:   "libya",
				Alpha2: "LY",
				Alpha3: "LBY",
			}},
			currencyName: "Libyan Dinar",
			numCode:      "434",
			minorUnit:    "3",
		}, "CHF": {
			countryName: []Country{{
				Name:   "liechtenstein",
				Alpha2: "LI",
				Alpha3: "LIE",
			}, {
				Name:   "switzerland",
				Alpha2: "CH",
				Alpha3: "CHE",
			}},
			currencyName: "Swiss Franc",
			numCode:      "756",
			minorUnit:    "2",
		}, "MOP": {
			countryName: []Country{{
				Name:   "macao",
				Alpha2: "MO",
				Alpha3: "MAC",
			}},
			currencyName: "Pataca",
			numCode:      "446",
			minorUnit:    "2",
		}, "MKD": {
			countryName: []Country{{
				Name:   "north macedonia",
				Alpha2: "MK",
				Alpha3: "MKD",
			}},
			currencyName: "Denar",
			numCode:      "807",
			minorUnit:    "2",
		}, "MGA": {
			countryName: []Country{{
				Name:   "madagascar",
				Alpha2: "MG",
				Alpha3: "MDG",
			}},
			currencyName: "Malagasy Ariary",
			numCode:      "969",
			minorUnit:    "2",
		}, "MWK": {
			countryName: []Country{{
				Name:   "malawi",
				Alpha2: "MW",
				Alpha3: "MWI",
			}},
			currencyName: "Malawi Kwacha",
			numCode:      "454",
			minorUnit:    "2",
		}, "MYR": {
			countryName: []Country{{
				Name:   "malaysia",
				Alpha2: "MY",
				Alpha3: "MYS",
			}},
			currencyName: "Malaysian Ringgit",
			numCode:      "458",
			minorUnit:    "2",
		}, "MVR": {
			countryName: []Country{{
				Name:   "maldives",
				Alpha2: "MV",
				Alpha3: "MDV",
			}},
			currencyName: "Rufiyaa",
			numCode:      "462",
			minorUnit:    "2",
		}, "MRU": {
			countryName: []Country{{
				Name:   "mauritania",
				Alpha2: "MR",
				Alpha3: "MRT",
			}},
			currencyName: "Ouguiya",
			numCode:      "929",
			minorUnit:    "2",
		}, "MUR": {
			countryName: []Country{{
				Name:   "mauritius",
				Alpha2: "MU",
				Alpha3: "MUS",
			}},
			currencyName: "Mauritius Rupee",
			numCode:      "480",
			minorUnit:    "2",
		}, "XUA": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "ADB Unit of Account",
			numCode:      "965",
			minorUnit:    "N.A.",
		}, "MXN": {
			countryName: []Country{{
				Name:   "mexico",
				Alpha2: "MX",
				Alpha3: "MEX",
			}},
			currencyName: "Mexican Peso",
			numCode:      "484",
			minorUnit:    "2",
		}, "MXV": {
			countryName: []Country{{
				Name:   "mexico",
				Alpha2: "MX",
				Alpha3: "MEX",
			}},
			currencyName: "Mexican Unidad de Inversion (UDI)",
			numCode:      "979",
			minorUnit:    "2",
		}, "MDL": {
			countryName: []Country{{
				Name:   "moldova (the republic of)",
				Alpha2: "MD",
				Alpha3: "MDA",
			}},
			currencyName: "Moldovan Leu",
			numCode:      "498",
			minorUnit:    "2",
		}, "MNT": {
			countryName: []Country{{
				Name:   "mongolia",
				Alpha2: "MN",
				Alpha3: "MNG",
			}},
			currencyName: "Tugrik",
			numCode:      "496",
			minorUnit:    "2",
		}, "MAD": {
			countryName: []Country{{
				Name:   "morocco",
				Alpha2: "MA",
				Alpha3: "MAR",
			}, {
				Name:   "western sahara",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "Moroccan Dirham",
			numCode:      "504",
			minorUnit:    "2",
		}, "MZN": {
			countryName: []Country{{
				Name:   "mozambique",
				Alpha2: "MZ",
				Alpha3: "MOZ",
			}},
			currencyName: "Mozambique Metical",
			numCode:      "943",
			minorUnit:    "2",
		}, "MMK": {
			countryName: []Country{{
				Name:   "myanmar",
				Alpha2: "MM",
				Alpha3: "MMR",
			}},
			currencyName: "Kyat",
			numCode:      "104",
			minorUnit:    "2",
		}, "NAD": {
			countryName: []Country{{
				Name:   "namibia",
				Alpha2: "NA",
				Alpha3: "NAM",
			}},
			currencyName: "Namibia Dollar",
			numCode:      "516",
			minorUnit:    "2",
		}, "NPR": {
			countryName: []Country{{
				Name:   "nepal",
				Alpha2: "NP",
				Alpha3: "NPL",
			}},
			currencyName: "Nepalese Rupee",
			numCode:      "524",
			minorUnit:    "2",
		}, "NIO": {
			countryName: []Country{{
				Name:   "nicaragua",
				Alpha2: "NI",
				Alpha3: "NIC",
			}},
			currencyName: "Cordoba Oro",
			numCode:      "558",
			minorUnit:    "2",
		}, "NGN": {
			countryName: []Country{{
				Name:   "nigeria",
				Alpha2: "NG",
				Alpha3: "NGA",
			}},
			currencyName: "Naira",
			numCode:      "566",
			minorUnit:    "2",
		}, "OMR": {
			countryName: []Country{{
				Name:   "oman",
				Alpha2: "OM",
				Alpha3: "OMN",
			}},
			currencyName: "Rial Omani",
			numCode:      "512",
			minorUnit:    "3",
		}, "PKR": {
			countryName: []Country{{
				Name:   "pakistan",
				Alpha2: "PK",
				Alpha3: "PAK",
			}},
			currencyName: "Pakistan Rupee",
			numCode:      "586",
			minorUnit:    "2",
		}, "PAB": {
			countryName: []Country{{
				Name:   "panama",
				Alpha2: "PA",
				Alpha3: "PAN",
			}},
			currencyName: "Balboa",
			numCode:      "590",
			minorUnit:    "2",
		}, "PGK": {
			countryName: []Country{{
				Name:   "papua new guinea",
				Alpha2: "PG",
				Alpha3: "PNG",
			}},
			currencyName: "Kina",
			numCode:      "598",
			minorUnit:    "2",
		}, "PYG": {
			countryName: []Country{{
				Name:   "paraguay",
				Alpha2: "PY",
				Alpha3: "PRY",
			}},
			currencyName: "Guarani",
			numCode:      "600",
			minorUnit:    "0",
		}, "PEN": {
			countryName: []Country{{
				Name:   "peru",
				Alpha2: "PE",
				Alpha3: "PER",
			}},
			currencyName: "Sol",
			numCode:      "604",
			minorUnit:    "2",
		}, "PHP": {
			countryName: []Country{{
				Name:   "philippines (the)",
				Alpha2: "PH",
				Alpha3: "PHL",
			}},
			currencyName: "Philippine Peso",
			numCode:      "608",
			minorUnit:    "2",
		}, "PLN": {
			countryName: []Country{{
				Name:   "poland",
				Alpha2: "PL",
				Alpha3: "POL",
			}},
			currencyName: "Zloty",
			numCode:      "985",
			minorUnit:    "2",
		}, "QAR": {
			countryName: []Country{{
				Name:   "qatar",
				Alpha2: "QA",
				Alpha3: "QAT",
			}},
			currencyName: "Qatari Rial",
			numCode:      "634",
			minorUnit:    "2",
		}, "RON": {
			countryName: []Country{{
				Name:   "romania",
				Alpha2: "RO",
				Alpha3: "ROU",
			}},
			currencyName: "Romanian Leu",
			numCode:      "946",
			minorUnit:    "2",
		}, "RUB": {
			countryName: []Country{{
				Name:   "russian federation (the)",
				Alpha2: "RU",
				Alpha3: "RUS",
			}},
			currencyName: "Russian Ruble",
			numCode:      "643",
			minorUnit:    "2",
		}, "RWF": {
			countryName: []Country{{
				Name:   "rwanda",
				Alpha2: "RW",
				Alpha3: "RWA",
			}},
			currencyName: "Rwanda Franc",
			numCode:      "646",
			minorUnit:    "0",
		}, "SHP": {
			countryName: []Country{{
				Name:   "saint helena, ascension and tristan da cunha",
				Alpha2: "SH",
				Alpha3: "SHN",
			}},
			currencyName: "Saint Helena Pound",
			numCode:      "654",
			minorUnit:    "2",
		}, "WST": {
			countryName: []Country{{
				Name:   "samoa",
				Alpha2: "WS",
				Alpha3: "WSM",
			}},
			currencyName: "Tala",
			numCode:      "882",
			minorUnit:    "2",
		}, "STN": {
			countryName: []Country{{
				Name:   "sao tome and principe",
				Alpha2: "ST",
				Alpha3: "STP",
			}},
			currencyName: "Dobra",
			numCode:      "930",
			minorUnit:    "2",
		}, "SAR": {
			countryName: []Country{{
				Name:   "saudi arabia",
				Alpha2: "SA",
				Alpha3: "SAU",
			}},
			currencyName: "Saudi Riyal",
			numCode:      "682",
			minorUnit:    "2",
		}, "RSD": {
			countryName: []Country{{
				Name:   "serbia",
				Alpha2: "RS",
				Alpha3: "SRB",
			}},
			currencyName: "Serbian Dinar",
			numCode:      "941",
			minorUnit:    "2",
		}, "SCR": {
			countryName: []Country{{
				Name:   "seychelles",
				Alpha2: "SC",
				Alpha3: "SYC",
			}},
			currencyName: "Seychelles Rupee",
			numCode:      "690",
			minorUnit:    "2",
		}, "SLL": {
			countryName: []Country{{
				Name:   "sierra leone",
				Alpha2: "SL",
				Alpha3: "SLE",
			}},
			currencyName: "Leone",
			numCode:      "694",
			minorUnit:    "2",
		}, "SGD": {
			countryName: []Country{{
				Name:   "singapore",
				Alpha2: "SG",
				Alpha3: "SGP",
			}},
			currencyName: "Singapore Dollar",
			numCode:      "702",
			minorUnit:    "2",
		}, "XSU": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "Sucre",
			numCode:      "994",
			minorUnit:    "N.A.",
		}, "SBD": {
			countryName: []Country{{
				Name:   "solomon islands",
				Alpha2: "SB",
				Alpha3: "SLB",
			}},
			currencyName: "Solomon Islands Dollar",
			numCode:      "090",
			minorUnit:    "2",
		}, "SOS": {
			countryName: []Country{{
				Name:   "somalia",
				Alpha2: "SO",
				Alpha3: "SOM",
			}},
			currencyName: "Somali Shilling",
			numCode:      "706",
			minorUnit:    "2",
		}, "SSP": {
			countryName: []Country{{
				Name:   "south sudan",
				Alpha2: "SS",
				Alpha3: "SSD",
			}},
			currencyName: "South Sudanese Pound",
			numCode:      "728",
			minorUnit:    "2",
		}, "LKR": {
			countryName: []Country{{
				Name:   "sri lanka",
				Alpha2: "LK",
				Alpha3: "LKA",
			}},
			currencyName: "Sri Lanka Rupee",
			numCode:      "144",
			minorUnit:    "2",
		}, "SDG": {
			countryName: []Country{{
				Name:   "sudan (the)",
				Alpha2: "SD",
				Alpha3: "SDN",
			}},
			currencyName: "Sudanese Pound",
			numCode:      "938",
			minorUnit:    "2",
		}, "SRD": {
			countryName: []Country{{
				Name:   "suriname",
				Alpha2: "SR",
				Alpha3: "SUR",
			}},
			currencyName: "Surinam Dollar",
			numCode:      "968",
			minorUnit:    "2",
		}, "SEK": {
			countryName: []Country{{
				Name:   "sweden",
				Alpha2: "SE",
				Alpha3: "SWE",
			}},
			currencyName: "Swedish Krona",
			numCode:      "752",
			minorUnit:    "2",
		}, "CHE": {
			countryName: []Country{{
				Name:   "switzerland",
				Alpha2: "CH",
				Alpha3: "CHE",
			}},
			currencyName: "WIR Euro",
			numCode:      "947",
			minorUnit:    "2",
		}, "CHW": {
			countryName: []Country{{
				Name:   "switzerland",
				Alpha2: "CH",
				Alpha3: "CHE",
			}},
			currencyName: "WIR Franc",
			numCode:      "948",
			minorUnit:    "2",
		}, "SYP": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "Syrian Pound",
			numCode:      "760",
			minorUnit:    "2",
		}, "TWD": {
			countryName: []Country{{
				Name:   "taiwan (province of china)",
				Alpha2: "TW",
				Alpha3: "TWN",
			}},
			currencyName: "New Taiwan Dollar",
			numCode:      "901",
			minorUnit:    "2",
		}, "TJS": {
			countryName: []Country{{
				Name:   "tajikistan",
				Alpha2: "TJ",
				Alpha3: "TJK",
			}},
			currencyName: "Somoni",
			numCode:      "972",
			minorUnit:    "2",
		}, "TZS": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "Tanzanian Shilling",
			numCode:      "834",
			minorUnit:    "2",
		}, "THB": {
			countryName: []Country{{
				Name:   "thailand",
				Alpha2: "TH",
				Alpha3: "THA",
			}},
			currencyName: "Baht",
			numCode:      "764",
			minorUnit:    "2",
		}, "TOP": {
			countryName: []Country{{
				Name:   "tonga",
				Alpha2: "TO",
				Alpha3: "TON",
			}},
			currencyName: "Pa’anga",
			numCode:      "776",
			minorUnit:    "2",
		}, "TTD": {
			countryName: []Country{{
				Name:   "trinidad and tobago",
				Alpha2: "TT",
				Alpha3: "TTO",
			}},
			currencyName: "Trinidad and Tobago Dollar",
			numCode:      "780",
			minorUnit:    "2",
		}, "TND": {
			countryName: []Country{{
				Name:   "tunisia",
				Alpha2: "TN",
				Alpha3: "TUN",
			}},
			currencyName: "Tunisian Dinar",
			numCode:      "788",
			minorUnit:    "3",
		}, "TRY": {
			countryName: []Country{{
				Name:   "turkey",
				Alpha2: "TR",
				Alpha3: "TUR",
			}},
			currencyName: "Turkish Lira",
			numCode:      "949",
			minorUnit:    "2",
		}, "TMT": {
			countryName: []Country{{
				Name:   "turkmenistan",
				Alpha2: "TM",
				Alpha3: "TKM",
			}},
			currencyName: "Turkmenistan New Manat",
			numCode:      "934",
			minorUnit:    "2",
		}, "UGX": {
			countryName: []Country{{
				Name:   "uganda",
				Alpha2: "UG",
				Alpha3: "UGA",
			}},
			currencyName: "Uganda Shilling",
			numCode:      "800",
			minorUnit:    "0",
		}, "UAH": {
			countryName: []Country{{
				Name:   "ukraine",
				Alpha2: "UA",
				Alpha3: "UKR",
			}},
			currencyName: "Hryvnia",
			numCode:      "980",
			minorUnit:    "2",
		}, "AED": {
			countryName: []Country{{
				Name:   "united arab emirates (the)",
				Alpha2: "AE",
				Alpha3: "ARE",
			}},
			currencyName: "UAE Dirham",
			numCode:      "784",
			minorUnit:    "2",
		}, "USN": {
			countryName: []Country{{
				Name:   "united states of america (the)",
				Alpha2: "US",
				Alpha3: "USA",
			}},
			currencyName: "US Dollar (Next day)",
			numCode:      "997",
			minorUnit:    "2",
		}, "UYU": {
			countryName: []Country{{
				Name:   "uruguay",
				Alpha2: "UY",
				Alpha3: "URY",
			}},
			currencyName: "Peso Uruguayo",
			numCode:      "858",
			minorUnit:    "2",
		}, "UYI": {
			countryName: []Country{{
				Name:   "uruguay",
				Alpha2: "UY",
				Alpha3: "URY",
			}},
			currencyName: "Uruguay Peso en Unidades Indexadas (UI)",
			numCode:      "940",
			minorUnit:    "0",
		}, "UYW": {
			countryName: []Country{{
				Name:   "uruguay",
				Alpha2: "UY",
				Alpha3: "URY",
			}},
			currencyName: "Unidad Previsional",
			numCode:      "927",
			minorUnit:    "4",
		}, "UZS": {
			countryName: []Country{{
				Name:   "uzbekistan",
				Alpha2: "UZ",
				Alpha3: "UZB",
			}},
			currencyName: "Uzbekistan Sum",
			numCode:      "860",
			minorUnit:    "2",
		}, "VUV": {
			countryName: []Country{{
				Name:   "vanuatu",
				Alpha2: "VU",
				Alpha3: "VUT",
			}},
			currencyName: "Vatu",
			numCode:      "548",
			minorUnit:    "0",
		}, "VES": {
			countryName: []Country{{
				Name:   "venezuela (bolivarian republic of)",
				Alpha2: "VE",
				Alpha3: "VEN",
			}},
			currencyName: "Bolívar Soberano",
			numCode:      "928",
			minorUnit:    "2",
		}, "VED": {
			countryName: []Country{{
				Name:   "venezuela (bolivarian republic of)",
				Alpha2: "VE",
				Alpha3: "VEN",
			}},
			currencyName: "Bolívar Soberano",
			numCode:      "926",
			minorUnit:    "2",
		}, "VND": {
			countryName: []Country{{
				Name:   "viet nam",
				Alpha2: "VN",
				Alpha3: "VNM",
			}},
			currencyName: "Dong",
			numCode:      "704",
			minorUnit:    "0",
		}, "YER": {
			countryName: []Country{{
				Name:   "yemen",
				Alpha2: "YE",
				Alpha3: "YEM",
			}},
			currencyName: "Yemeni Rial",
			numCode:      "886",
			minorUnit:    "2",
		}, "ZMW": {
			countryName: []Country{{
				Name:   "zambia",
				Alpha2: "ZM",
				Alpha3: "ZMB",
			}},
			currencyName: "Zambian Kwacha",
			numCode:      "967",
			minorUnit:    "2",
		}, "ZWL": {
			countryName: []Country{{
				Name:   "zimbabwe",
				Alpha2: "ZW",
				Alpha3: "ZWE",
			}},
			currencyName: "Zimbabwe Dollar",
			numCode:      "932",
			minorUnit:    "2",
		}, "XBA": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "Bond Markets Unit European Composite Unit (EURCO)",
			numCode:      "955",
			minorUnit:    "N.A.",
		}, "XBB": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "Bond Markets Unit European Monetary Unit (E.M.U.-6)",
			numCode:      "956",
			minorUnit:    "N.A.",
		}, "XBC": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "Bond Markets Unit European Unit of Account 9 (E.U.A.-9)",
			numCode:      "957",
			minorUnit:    "N.A.",
		}, "XBD": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "Bond Markets Unit European Unit of Account 17 (E.U.A.-17)",
			numCode:      "958",
			minorUnit:    "N.A.",
		}, "XTS": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "Codes specifically reserved for testing purposes",
			numCode:      "963",
			minorUnit:    "N.A.",
		}, "XXX": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "The codes assigned for transactions where no currency is involved",
			numCode:      "999",
			minorUnit:    "N.A.",
		}, "XAU": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "Gold",
			numCode:      "959",
			minorUnit:    "N.A.",
		}, "XPD": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "Palladium",
			numCode:      "964",
			minorUnit:    "N.A.",
		}, "XPT": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "Platinum",
			numCode:      "962",
			minorUnit:    "N.A.",
		}, "XAG": {
			countryName: []Country{{
				Name:   "unknown",
				Alpha2: "XX",
				Alpha3: "XXX",
			}},
			currencyName: "Silver",
			numCode:      "961",
			minorUnit:    "N.A.",
		}}
)
