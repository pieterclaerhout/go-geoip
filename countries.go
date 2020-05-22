package geoip

import (
	"errors"
	"strings"
)

var countryCodes map[string]string

var countriesForRegionWestUS []string
var countriesForRegionSouthBrazil []string
var countriesForRegionJapanEast []string
var countriesForRegionSouthEastAsia []string

func init() {

	countryCodes = make(map[string]string)
	countryCodes["A1"] = "Anonymous Proxy"
	countryCodes["A2"] = "Satellite Provider"
	countryCodes["O1"] = "Other Country"
	countryCodes["AD"] = "Andorra"
	countryCodes["AE"] = "United Arab Emirates"
	countryCodes["AF"] = "Afghanistan"
	countryCodes["AG"] = "Antigua and Barbuda"
	countryCodes["AI"] = "Anguilla"
	countryCodes["AL"] = "Albania"
	countryCodes["AM"] = "Armenia"
	countryCodes["AO"] = "Angola"
	countryCodes["AP"] = "Asia/Pacific Region"
	countryCodes["AQ"] = "Antarctica"
	countryCodes["AR"] = "Argentina"
	countryCodes["AS"] = "American Samoa"
	countryCodes["AT"] = "Austria"
	countryCodes["AU"] = "Australia"
	countryCodes["AW"] = "Aruba"
	countryCodes["AX"] = "Aland Islands"
	countryCodes["AZ"] = "Azerbaijan"
	countryCodes["BA"] = "Bosnia and Herzegovina"
	countryCodes["BB"] = "Barbados"
	countryCodes["BD"] = "Bangladesh"
	countryCodes["BE"] = "Belgium"
	countryCodes["BF"] = "Burkina Faso"
	countryCodes["BG"] = "Bulgaria"
	countryCodes["BH"] = "Bahrain"
	countryCodes["BI"] = "Burundi"
	countryCodes["BJ"] = "Benin"
	countryCodes["BL"] = "Saint Bartelemey"
	countryCodes["BM"] = "Bermuda"
	countryCodes["BN"] = "Brunei Darussalam"
	countryCodes["BO"] = "Bolivia"
	countryCodes["BQ"] = "Bonaire, Saint Eustatius and Saba"
	countryCodes["BR"] = "Brazil"
	countryCodes["BS"] = "Bahamas"
	countryCodes["BT"] = "Bhutan"
	countryCodes["BV"] = "Bouvet Island"
	countryCodes["BW"] = "Botswana"
	countryCodes["BY"] = "Belarus"
	countryCodes["BZ"] = "Belize"
	countryCodes["CA"] = "Canada"
	countryCodes["CC"] = "Cocos (Keeling) Islands"
	countryCodes["CD"] = "Congo, The Democratic Republic of the"
	countryCodes["CF"] = "Central African Republic"
	countryCodes["CG"] = "Congo"
	countryCodes["CH"] = "Switzerland"
	countryCodes["CI"] = "Cote d'Ivoire"
	countryCodes["CK"] = "Cook Islands"
	countryCodes["CL"] = "Chile"
	countryCodes["CM"] = "Cameroon"
	countryCodes["CN"] = "China"
	countryCodes["CO"] = "Colombia"
	countryCodes["CR"] = "Costa Rica"
	countryCodes["CU"] = "Cuba"
	countryCodes["CV"] = "Cape Verde"
	countryCodes["CW"] = "Curacao"
	countryCodes["CX"] = "Christmas Island"
	countryCodes["CY"] = "Cyprus"
	countryCodes["CZ"] = "Czech Republic"
	countryCodes["DE"] = "Germany"
	countryCodes["DJ"] = "Djibouti"
	countryCodes["DK"] = "Denmark"
	countryCodes["DM"] = "Dominica"
	countryCodes["DO"] = "Dominican Republic"
	countryCodes["DZ"] = "Algeria"
	countryCodes["EC"] = "Ecuador"
	countryCodes["EE"] = "Estonia"
	countryCodes["EG"] = "Egypt"
	countryCodes["EH"] = "Western Sahara"
	countryCodes["ER"] = "Eritrea"
	countryCodes["ES"] = "Spain"
	countryCodes["ET"] = "Ethiopia"
	countryCodes["EU"] = "Europe"
	countryCodes["FI"] = "Finland"
	countryCodes["FJ"] = "Fiji"
	countryCodes["FK"] = "Falkland Islands (Malvinas)"
	countryCodes["FM"] = "Micronesia, Federated States of"
	countryCodes["FO"] = "Faroe Islands"
	countryCodes["FR"] = "France"
	countryCodes["GA"] = "Gabon"
	countryCodes["GB"] = "United Kingdom"
	countryCodes["GD"] = "Grenada"
	countryCodes["GE"] = "Georgia"
	countryCodes["GF"] = "French Guiana"
	countryCodes["GG"] = "Guernsey"
	countryCodes["GH"] = "Ghana"
	countryCodes["GI"] = "Gibraltar"
	countryCodes["GL"] = "Greenland"
	countryCodes["GM"] = "Gambia"
	countryCodes["GN"] = "Guinea"
	countryCodes["GP"] = "Guadeloupe"
	countryCodes["GQ"] = "Equatorial Guinea"
	countryCodes["GR"] = "Greece"
	countryCodes["GS"] = "South Georgia and the South Sandwich Islands"
	countryCodes["GT"] = "Guatemala"
	countryCodes["GU"] = "Guam"
	countryCodes["GW"] = "Guinea-Bissau"
	countryCodes["GY"] = "Guyana"
	countryCodes["HK"] = "Hong Kong"
	countryCodes["HM"] = "Heard Island and McDonald Islands"
	countryCodes["HN"] = "Honduras"
	countryCodes["HR"] = "Croatia"
	countryCodes["HT"] = "Haiti"
	countryCodes["HU"] = "Hungary"
	countryCodes["ID"] = "Indonesia"
	countryCodes["IE"] = "Ireland"
	countryCodes["IL"] = "Israel"
	countryCodes["IM"] = "Isle of Man"
	countryCodes["IN"] = "India"
	countryCodes["IO"] = "British Indian Ocean Territory"
	countryCodes["IQ"] = "Iraq"
	countryCodes["IR"] = "Iran, Islamic Republic of"
	countryCodes["IS"] = "Iceland"
	countryCodes["IT"] = "Italy"
	countryCodes["JE"] = "Jersey"
	countryCodes["JM"] = "Jamaica"
	countryCodes["JO"] = "Jordan"
	countryCodes["JP"] = "Japan"
	countryCodes["KE"] = "Kenya"
	countryCodes["KG"] = "Kyrgyzstan"
	countryCodes["KH"] = "Cambodia"
	countryCodes["KI"] = "Kiribati"
	countryCodes["KM"] = "Comoros"
	countryCodes["KN"] = "Saint Kitts and Nevis"
	countryCodes["KP"] = "Korea, Democratic People's Republic of"
	countryCodes["KR"] = "Korea, Republic of"
	countryCodes["KW"] = "Kuwait"
	countryCodes["KY"] = "Cayman Islands"
	countryCodes["KZ"] = "Kazakhstan"
	countryCodes["LA"] = "Lao People's Democratic Republic"
	countryCodes["LB"] = "Lebanon"
	countryCodes["LC"] = "Saint Lucia"
	countryCodes["LI"] = "Liechtenstein"
	countryCodes["LK"] = "Sri Lanka"
	countryCodes["LR"] = "Liberia"
	countryCodes["LS"] = "Lesotho"
	countryCodes["LT"] = "Lithuania"
	countryCodes["LU"] = "Luxembourg"
	countryCodes["LV"] = "Latvia"
	countryCodes["LY"] = "Libyan Arab Jamahiriya"
	countryCodes["MA"] = "Morocco"
	countryCodes["MC"] = "Monaco"
	countryCodes["MD"] = "Moldova, Republic of"
	countryCodes["ME"] = "Montenegro"
	countryCodes["MF"] = "Saint Martin"
	countryCodes["MG"] = "Madagascar"
	countryCodes["MH"] = "Marshall Islands"
	countryCodes["MK"] = "Macedonia"
	countryCodes["ML"] = "Mali"
	countryCodes["MM"] = "Myanmar"
	countryCodes["MN"] = "Mongolia"
	countryCodes["MO"] = "Macao"
	countryCodes["MP"] = "Northern Mariana Islands"
	countryCodes["MQ"] = "Martinique"
	countryCodes["MR"] = "Mauritania"
	countryCodes["MS"] = "Montserrat"
	countryCodes["MT"] = "Malta"
	countryCodes["MU"] = "Mauritius"
	countryCodes["MV"] = "Maldives"
	countryCodes["MW"] = "Malawi"
	countryCodes["MX"] = "Mexico"
	countryCodes["MY"] = "Malaysia"
	countryCodes["MZ"] = "Mozambique"
	countryCodes["NA"] = "Namibia"
	countryCodes["NC"] = "New Caledonia"
	countryCodes["NE"] = "Niger"
	countryCodes["NF"] = "Norfolk Island"
	countryCodes["NG"] = "Nigeria"
	countryCodes["NI"] = "Nicaragua"
	countryCodes["NL"] = "Netherlands"
	countryCodes["NO"] = "Norway"
	countryCodes["NP"] = "Nepal"
	countryCodes["NR"] = "Nauru"
	countryCodes["NU"] = "Niue"
	countryCodes["NZ"] = "New Zealand"
	countryCodes["OM"] = "Oman"
	countryCodes["PA"] = "Panama"
	countryCodes["PE"] = "Peru"
	countryCodes["PF"] = "French Polynesia"
	countryCodes["PG"] = "Papua New Guinea"
	countryCodes["PH"] = "Philippines"
	countryCodes["PK"] = "Pakistan"
	countryCodes["PL"] = "Poland"
	countryCodes["PM"] = "Saint Pierre and Miquelon"
	countryCodes["PN"] = "Pitcairn"
	countryCodes["PR"] = "Puerto Rico"
	countryCodes["PS"] = "Palestinian Territory"
	countryCodes["PT"] = "Portugal"
	countryCodes["PW"] = "Palau"
	countryCodes["PY"] = "Paraguay"
	countryCodes["QA"] = "Qatar"
	countryCodes["RE"] = "Reunion"
	countryCodes["RO"] = "Romania"
	countryCodes["RS"] = "Serbia"
	countryCodes["RU"] = "Russian Federation"
	countryCodes["RW"] = "Rwanda"
	countryCodes["SA"] = "Saudi Arabia"
	countryCodes["SB"] = "Solomon Islands"
	countryCodes["SC"] = "Seychelles"
	countryCodes["SD"] = "Sudan"
	countryCodes["SE"] = "Sweden"
	countryCodes["SG"] = "Singapore"
	countryCodes["SH"] = "Saint Helena"
	countryCodes["SI"] = "Slovenia"
	countryCodes["SJ"] = "Svalbard and Jan Mayen"
	countryCodes["SK"] = "Slovakia"
	countryCodes["SL"] = "Sierra Leone"
	countryCodes["SM"] = "San Marino"
	countryCodes["SN"] = "Senegal"
	countryCodes["SO"] = "Somalia"
	countryCodes["SR"] = "Suriname"
	countryCodes["SS"] = "South Sudan"
	countryCodes["ST"] = "Sao Tome and Principe"
	countryCodes["SV"] = "El Salvador"
	countryCodes["SX"] = "Sint Maarten"
	countryCodes["SY"] = "Syrian Arab Republic"
	countryCodes["SZ"] = "Swaziland"
	countryCodes["TC"] = "Turks and Caicos Islands"
	countryCodes["TD"] = "Chad"
	countryCodes["TF"] = "French Southern Territories"
	countryCodes["TG"] = "Togo"
	countryCodes["TH"] = "Thailand"
	countryCodes["TJ"] = "Tajikistan"
	countryCodes["TK"] = "Tokelau"
	countryCodes["TL"] = "Timor-Leste"
	countryCodes["TM"] = "Turkmenistan"
	countryCodes["TN"] = "Tunisia"
	countryCodes["TO"] = "Tonga"
	countryCodes["TR"] = "Turkey"
	countryCodes["TT"] = "Trinidad and Tobago"
	countryCodes["TV"] = "Tuvalu"
	countryCodes["TW"] = "Taiwan"
	countryCodes["TZ"] = "Tanzania, United Republic of"
	countryCodes["UA"] = "Ukraine"
	countryCodes["UG"] = "Uganda"
	countryCodes["UM"] = "United States Minor Outlying Islands"
	countryCodes["US"] = "United States"
	countryCodes["UY"] = "Uruguay"
	countryCodes["UZ"] = "Uzbekistan"
	countryCodes["VA"] = "Holy See (Vatican City State)"
	countryCodes["VC"] = "Saint Vincent and the Grenadines"
	countryCodes["VE"] = "Venezuela"
	countryCodes["VG"] = "Virgin Islands, British"
	countryCodes["VI"] = "Virgin Islands, U.S."
	countryCodes["VN"] = "Vietnam"
	countryCodes["VU"] = "Vanuatu"
	countryCodes["WF"] = "Wallis and Futuna"
	countryCodes["WS"] = "Samoa"
	countryCodes["YE"] = "Yemen"
	countryCodes["YT"] = "Mayotte"
	countryCodes["ZA"] = "South Africa"
	countryCodes["ZM"] = "Zambia"
	countryCodes["ZW"] = "Zimbabwe"

	countriesForRegionWestUS = []string{"CA", "US", "MX", "CU", "DO", "GT", "HN", "NI", "CR", "PA", "CO", "VE", "GY", "EC", "PE", "BO", "PY", "UY", "CL", "AR"}
	countriesForRegionSouthBrazil = []string{"BR"}
	countriesForRegionJapanEast = []string{"JP"}
	countriesForRegionSouthEastAsia = []string{"AU", "BN", "KH", "ID", "LA", "MY", "MM", "PH", "SG", "TH", "VN", "CN", "HK", "MO", "MN", "KP", "KR", "TW"}

}

// CountryCodeToName translates an ISO country to the English name
func CountryCodeToName(code string) (string, error) {
	for countryCode, countryName := range countryCodes {
		if strings.ToLower(code) == strings.ToLower(countryCode) {
			return countryName, nil
		}
	}
	return "", errors.New("Unknown country code: " + code)
}

// CountryCodeToRegion translates an ISO country code to a region
// Defaults to "west-europe"
func CountryCodeToRegion(code string) string {
	if sliceContainsString(countriesForRegionWestUS, code) {
		return regionWestUS
	} else if sliceContainsString(countriesForRegionSouthBrazil, code) {
		return regionSouthBrazil
	} else if sliceContainsString(countriesForRegionJapanEast, code) {
		return regionJapanEast
	} else if sliceContainsString(countriesForRegionSouthEastAsia, code) {
		return regionSouthEastAsia
	}
	return regionWestEurope
}

// CountryNameToCode translates the country name to it's ISO country code
func CountryNameToCode(name string) (string, error) {
	for countryCode, countryName := range countryCodes {
		if strings.ToLower(name) == strings.ToLower(countryName) {
			return countryCode, nil
		}
	}
	return "", errors.New("Unknown country name: " + name)
}

// sliceContainsString checks if the slice contains the given string in a case-sensitive way
func sliceContainsString(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}
