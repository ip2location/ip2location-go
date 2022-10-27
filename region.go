package ip2location

import (
	"encoding/csv"
	"errors"
	"os"
	"strings"
)

// The RegionInfoRecord struct stores all of the available
// region info found in the region information CSV file.
type RegionInfoRecord struct {
	Country_code string
	Name         string
	Code         string
}

// The RI struct is the main object used to read the region information CSV.
type RI struct {
	resultsMap map[string][]RegionInfoRecord
}

// OpenRegionInfo initializes with the path to the region information CSV file.
func OpenRegionInfo(csvFile string) (*RI, error) {
	var ri = &RI{}

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

	ri.resultsMap = make(map[string][]RegionInfoRecord)
	var headerArr []string
	var resultsArr []RegionInfoRecord

	for i, line := range data {
		if i == 0 { // headers
			for _, field := range line {
				headerArr = append(headerArr, field)
			}
		} else {
			var rec RegionInfoRecord
			for j, field := range line {
				switch headerArr[j] {
				case "country_code":
					rec.Country_code = field
				case "subdivision_name":
					rec.Name = field
				case "code":
					rec.Code = field
				}
			}
			if rec.Name == "" {
				return nil, errors.New("Invalid region information CSV file.")
			}
			resultsArr = append(resultsArr, rec)
		}
	}
	for _, elem := range resultsArr {
		if _, ok := ri.resultsMap[elem.Country_code]; !ok {
			var arr []RegionInfoRecord
			ri.resultsMap[elem.Country_code] = arr
		}
		ri.resultsMap[elem.Country_code] = append(ri.resultsMap[elem.Country_code], elem)
	}
	return ri, nil
}

// GetRegionCode returns the region code for the specified country and region name
func (r *RI) GetRegionCode(countryCode string, regionName string) (string, error) {
	if len(r.resultsMap) == 0 {
		return "", errors.New("No record available.")
	}

	if arr, ok := r.resultsMap[countryCode]; ok {
		for _, elem := range arr {
			if strings.ToUpper(elem.Name) == strings.ToUpper(regionName) {
				return elem.Code, nil
			}
		}
	}
	return "", errors.New("No record found.")
}
