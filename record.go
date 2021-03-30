package ip2location

import "fmt"

// The Record struct stores all of the available
// geolocation info found in the IP2Location database.
type Record struct {
	CountryShort       string
	CountryLong        string
	Region             string
	City               string
	ISP                string
	Latitude           float32
	Longitude          float32
	Domain             string
	Zipcode            string
	Timezone           string
	NetSpeed           string
	IDDCode            string
	AreaCode           string
	WeatherStationCode string
	WeatherStationName string
	MCC                string
	MNC                string
	MobileBrand        string
	Elevation          float32
	UsageType          string
}

// populate record with message
func loadMessage(mesg string) Record {
	var x Record

	x.CountryShort = mesg
	x.CountryLong = mesg
	x.Region = mesg
	x.City = mesg
	x.ISP = mesg
	x.Domain = mesg
	x.Zipcode = mesg
	x.Timezone = mesg
	x.NetSpeed = mesg
	x.IDDCode = mesg
	x.AreaCode = mesg
	x.WeatherStationCode = mesg
	x.WeatherStationName = mesg
	x.MCC = mesg
	x.MNC = mesg
	x.MobileBrand = mesg
	x.UsageType = mesg

	return x
}

func handleError(rec Record, err error) Record {
	if err != nil {
		fmt.Print(err)
	}
	return rec
}

// PrintRecord is used to output the geolocation data for debugging purposes.
func PrintRecord(x Record) {
	fmt.Printf("country_short: %s\n", x.CountryShort)
	fmt.Printf("country_long: %s\n", x.CountryLong)
	fmt.Printf("region: %s\n", x.Region)
	fmt.Printf("city: %s\n", x.City)
	fmt.Printf("isp: %s\n", x.ISP)
	fmt.Printf("latitude: %f\n", x.Latitude)
	fmt.Printf("longitude: %f\n", x.Longitude)
	fmt.Printf("domain: %s\n", x.Domain)
	fmt.Printf("zipcode: %s\n", x.Zipcode)
	fmt.Printf("timezone: %s\n", x.Timezone)
	fmt.Printf("netSpeed: %s\n", x.NetSpeed)
	fmt.Printf("iddCode: %s\n", x.IDDCode)
	fmt.Printf("areaCode: %s\n", x.AreaCode)
	fmt.Printf("weatherStationCode: %s\n", x.WeatherStationCode)
	fmt.Printf("weatherStationName: %s\n", x.WeatherStationName)
	fmt.Printf("mcc: %s\n", x.MCC)
	fmt.Printf("mnc: %s\n", x.MNC)
	fmt.Printf("mobileBrand: %s\n", x.MobileBrand)
	fmt.Printf("elevation: %f\n", x.Elevation)
	fmt.Printf("usageType: %s\n", x.UsageType)
}
