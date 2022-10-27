package ip2location

import (
	"encoding/csv"
	"errors"
	"os"
)

// The CountryInfoRecord struct stores all of the available
// country info found in the country information CSV file.
type CountryInfoRecord struct {
	Country_code         string
	Country_name         string
	Country_alpha3_code  string
	Country_numeric_code string
	Capital              string
	Country_demonym      string
	Total_area           string
	Population           string
	Idd_code             string
	Currency_code        string
	Currency_name        string
	Currency_symbol      string
	Lang_code            string
	Lang_name            string
	Cctld                string
}

// The CI struct is the main object used to read the country information CSV.
type CI struct {
	resultsArr []CountryInfoRecord
	resultsMap map[string]CountryInfoRecord
}

// OpenCountryInfo initializes with the path to the country information CSV file.
func OpenCountryInfo(csvFile string) (*CI, error) {
	var ci = &CI{}

	_, err := os.Stat(csvFile)
	if os.IsNotExist(err) {
		return nil, errors.New("The CSV file '" + csvFile + "' is not found.")
	}

	f, err := os.Open(csvFile)
	if err != nil {
		return nil, errors.New("Unable to read '" + csvFile + "'.")
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.New("Unable to read '" + csvFile + "'.")
	}

	ci.resultsMap = make(map[string]CountryInfoRecord)
	var headerArr []string

	for i, line := range data {
		if i == 0 { // headers
			for _, field := range line {
				headerArr = append(headerArr, field)
			}
		} else {
			var rec CountryInfoRecord
			for j, field := range line {
				switch headerArr[j] {
				case "country_code":
					rec.Country_code = field
				case "country_name":
					rec.Country_name = field
				case "country_alpha3_code":
					rec.Country_alpha3_code = field
				case "country_numeric_code":
					rec.Country_numeric_code = field
				case "capital":
					rec.Capital = field
				case "country_demonym":
					rec.Country_demonym = field
				case "total_area":
					rec.Total_area = field
				case "population":
					rec.Population = field
				case "idd_code":
					rec.Idd_code = field
				case "currency_code":
					rec.Currency_code = field
				case "currency_name":
					rec.Currency_name = field
				case "currency_symbol":
					rec.Currency_symbol = field
				case "lang_code":
					rec.Lang_code = field
				case "lang_name":
					rec.Lang_name = field
				case "cctld":
					rec.Cctld = field
				}
			}
			if rec.Country_code == "" {
				return nil, errors.New("Invalid country information CSV file.")
			}
			ci.resultsArr = append(ci.resultsArr, rec)
			ci.resultsMap[rec.Country_code] = rec
		}
	}
	return ci, nil
}

// GetCountryInfo returns the country information for the specified country or all countries if not specified
func (c *CI) GetCountryInfo(countryCode ...string) ([]CountryInfoRecord, error) {
	if len(c.resultsArr) == 0 {
		return nil, errors.New("No record available.")
	}

	if len(countryCode) == 1 {
		cc := countryCode[0]
		if rec, ok := c.resultsMap[cc]; ok {
			var x []CountryInfoRecord
			x = append(x, rec)
			return x, nil // return record
		} else {
			return nil, errors.New("No record found.")
		}
	} else {
		return c.resultsArr, nil // return all countries
	}
}
