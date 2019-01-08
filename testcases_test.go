package vat_test

var invalidVATNumberFormat = []struct {
	VATNumber   string
	CountryCode string
}{
	{"", ""},
	{"A", ""},
	{"AB123A01", ""},
	{"ATU1234567", "AT"},
	{"BE012345678", "BE"},
	{"BE123456789", "BE"},
	{"BG1234567", "BG"},
	{"CHE-156.730.098 MWST", ""},
	{"CHE-156.730.098", ""},
	{"CHE156730098MWST", ""},
	{"CHE156730098", ""},
	{"CY1234567X", "CY"},
	{"CZ1234567", "CZ"},
	{"DE12345678", "DE"},
	{"DK1234567", "DK"},
	{"EE12345678", "EE"},
	{"EL12345678", "GR"},
	{"ESX1234567", "ES"},
	{"FI1234567", "FI"},
	{"FR1234567890", "FR"},
	{"GB99999997", "GB"},
	{"HU1234567", "HU"},
	{"HR1234567890", "HR"},
	{"IE123456X", "IE"},
	{"IT1234567890", "IT"},
	{"LT12345678", "LT"},
	{"LU1234567", "LU"},
	{"LV1234567890", "LV"},
	{"MT1234567", "MT"},
	{"NL12345678B12", "NL"},
	{"PL123456789", "PL"},
	{"PT12345678", "PT"},
	{"RO1", "RO"}, // Romania has a really weird VAT format...
	{"SE12345678901", "SE"},
	{"SI1234567", "SI"},
	{"SK123456789", "SK"},
}

var validVATNumberFormat = []struct {
	VATNumber   string
	CountryCode string
}{
	{"ATU12345678", "AT"},
	{"ATU15673009", "AT"},
	{"BE0123456789", "BE"},
	{"BE1234567891", "BE"},
	{"BE0999999999", "BE"},
	{"BE9999999999", "BE"},
	{"BG123456789", "BG"},
	{"BG1234567890", "BG"},
	{"CY12345678X", "CY"},
	{"CY15673009L", "CY"},
	{"CZ12345678", "CZ"},
	{"DE123456789", "DE"},
	{"DK12345678", "DK"},
	{"EE123456789", "EE"},
	{"EL123456789", "GR"},
	{"ESX12345678", "ES"},
	{"FI12345678", "FI"},
	{"FR12345678901", "FR"},
	{"GB999999973", "GB"},
	{"GB156730098481", "GB"},
	{"GBGD549", "GB"},
	{"GBHA549", "GB"},
	{"HU12345678", "HU"},
	{"HR12345678901", "HR"},
	{"IE1234567X", "IE"},
	{"IT12345678901", "IT"},
	{"LT123456789", "LT"},
	{"LU26375245", "LU"},
	{"LU12345678", "LU"},
	{"LV12345678901", "LV"},
	{"MT12345678", "MT"},
	{"NL123456789B01", "NL"},
	{"NL123456789B12", "NL"},
	{"PL1234567890", "PL"},
	{"PT123456789", "PT"},
	{"RO123456789", "RO"},
	{"SE123456789012", "SE"},
	{"SI12345678", "SI"},
	{"SK1234567890", "SK"},
}

var existentVATNumbers = []struct {
	VATNumber   string
	CountryCode string
	Name        string
	Address     string
}{
	{"IE6388047V", "IE", "GOOGLE IRELAND LIMITED", "3RD FLOOR, GORDON HOUSE, BARROW STREET, DUBLIN 4"},
	{"EL094160738", "GR", "ΑΕ ΤΟΥΡΙΣΤΙΚΕΣ ΓΡΑΜΜΕΣ ΑΙΓΑΙΟΥ ΑΝΩΝΥΜΟΣ ΕΤΑΙΡΙΑ ||AGL", "ΚΟΛΟΚΟΤΡΩΝΗ 99 18535 - ΠΕΙΡΑΙΑΣ"},
	{"GB270600730", "GB", "SFD SA", "HM REVENUE AND CUSTOMS, RUBY HOUSE, 8 RUBY PLACE, ABERDEEN, AB10 1ZP"},
}

var nonexistentVATNumbers = []string{}
