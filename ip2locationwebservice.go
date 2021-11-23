package ip2location

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// The IP2LocationResult struct stores all of the available
// geolocation info found in the IP2Location Web Service.
type IP2LocationResult struct {
	Response           string  `json:"response"`
	CountryCode        string  `json:"country_code"`
	CountryName        string  `json:"country_name"`
	RegionName         string  `json:"region_name"`
	CityName           string  `json:"city_name"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	ZipCode            string  `json:"zip_code"`
	TimeZone           string  `json:"time_zone"`
	Isp                string  `json:"isp"`
	Domain             string  `json:"domain"`
	NetSpeed           string  `json:"net_speed"`
	IddCode            string  `json:"idd_code"`
	AreaCode           string  `json:"area_code"`
	WeatherStationCode string  `json:"weather_station_code"`
	WeatherStationName string  `json:"weather_station_name"`
	Mcc                string  `json:"mcc"`
	Mnc                string  `json:"mnc"`
	MobileBrand        string  `json:"mobile_brand"`
	Elevation          int     `json:"elevation"`
	UsageType          string  `json:"usage_type"`
	AddressType        string  `json:"address_type"`
	Category           string  `json:"category"`
	CategoryName       string  `json:"category_name"`
	Geotargeting       struct {
		Metro string `json:"metro"`
	} `json:"geotargeting"`
	Continent struct {
		Name       string   `json:"name"`
		Code       string   `json:"code"`
		Hemisphere []string `json:"hemisphere"`
	} `json:"continent"`
	Country struct {
		Name        string `json:"name"`
		Alpha3Code  string `json:"alpha3_code"`
		NumericCode string `json:"numeric_code"`
		Demonym     string `json:"demonym"`
		Flag        string `json:"flag"`
		Capital     string `json:"capital"`
		TotalArea   string `json:"total_area"`
		Population  string `json:"population"`
		Currency    struct {
			Code   string `json:"code"`
			Name   string `json:"name"`
			Symbol string `json:"symbol"`
		} `json:"currency"`
		Language struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"language"`
		IddCode string `json:"idd_code"`
		Tld     string `json:"tld"`
		IsEu    bool   `json:"is_eu"`
	} `json:"country"`
	CountryGroupings []struct {
		Acronym string `json:"acronym"`
		Name    string `json:"name"`
	} `json:"country_groupings"`
	Region struct {
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"region"`
	City struct {
		Name string `json:"name"`
	} `json:"city"`
	TimeZoneInfo struct {
		Olson       string `json:"olson"`
		CurrentTime string `json:"current_time"`
		GmtOffset   int    `json:"gmt_offset"`
		IsDst       string `json:"is_dst"`
		Sunrise     string `json:"sunrise"`
		Sunset      string `json:"sunset"`
	} `json:"time_zone_info"`
	CreditsConsumed int `json:"credits_consumed"`
}

// The IP2LocationCreditResult struct stores the
// credit balance for the IP2Location Web Service.
type IP2LocationCreditResult struct {
	Response int `json:"response"`
}

// The WS struct is the main object used to query the IP2Location Web Service.
type WS struct {
	apiKey     string
	apiPackage string
	useSSL     bool
}

var regexAPIKey = regexp.MustCompile(`^[\dA-Z]{10}$`)
var regexAPIPackage = regexp.MustCompile(`^WS\d+$`)

const baseURL = "api.ip2location.com/v2/"
const msgInvalidAPIKey = "Invalid API key."
const msgInvalidAPIPackage = "Invalid package name."

// OpenWS initializes with the web service API key, API package and whether to use SSL
func OpenWS(apikey string, apipackage string, usessl bool) (*WS, error) {
	var ws = &WS{}
	ws.apiKey = apikey
	ws.apiPackage = apipackage
	ws.useSSL = usessl

	err := ws.checkParams()

	if err != nil {
		return nil, err
	}

	return ws, nil
}

func (w *WS) checkParams() error {
	if !regexAPIKey.MatchString(w.apiKey) {
		return errors.New(msgInvalidAPIKey)
	}

	if !regexAPIPackage.MatchString(w.apiPackage) {
		return errors.New(msgInvalidAPIPackage)
	}

	return nil
}

// LookUp will return all geolocation fields based on the queried IP address, addon, lang
func (w *WS) LookUp(ipAddress string, addOn string, lang string) (IP2LocationResult, error) {
	var res IP2LocationResult
	err := w.checkParams()

	if err != nil {
		return res, err
	}

	protocol := "https"

	if !w.useSSL {
		protocol = "http"
	}

	// lang param not supported yet due to the inconsistent data type being returned by the API
	myUrl := protocol + "://" + baseURL + "?key=" + w.apiKey + "&package=" + w.apiPackage + "&ip=" + url.QueryEscape(ipAddress) + "&addon=" + url.QueryEscape(addOn)

	resp, err := http.Get(myUrl)

	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return res, err
		}

		err = json.Unmarshal(bodyBytes, &res)

		if err != nil {
			return res, err
		}

		return res, nil
	}

	return res, errors.New("Error HTTP " + strconv.Itoa(int(resp.StatusCode)))
}

// GetCredit will return the web service credit balance.
func (w *WS) GetCredit() (IP2LocationCreditResult, error) {
	var res IP2LocationCreditResult
	err := w.checkParams()

	if err != nil {
		return res, err
	}

	protocol := "https"

	if !w.useSSL {
		protocol = "http"
	}

	myUrl := protocol + "://" + baseURL + "?key=" + w.apiKey + "&check=true"

	resp, err := http.Get(myUrl)

	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			return res, err
		}

		err = json.Unmarshal(bodyBytes, &res)

		if err != nil {
			return res, err
		}

		return res, nil
	}

	return res, errors.New("Error HTTP " + strconv.Itoa(int(resp.StatusCode)))
}
