package ip2location

var defaultDB = &DB{}

// Open takes the path to the IP2Location BIN database file. It will read all the metadata required to
// be able to extract the embedded geolocation data.
//
// Deprecated: No longer being updated.
func Open(dbpath string) {
	db, err := OpenDB(dbpath)
	if err != nil {
		return
	}

	defaultDB = db
}

// GetAll will return all geolocation fields based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetAll(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, all))
}

// GetCountryShort will return the ISO-3166 country code based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetCountryShort(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, countryShort))
}

// GetCountryLong will return the country name based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetCountryLong(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, countryLong))
}

// GetRegion will return the region name based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetRegion(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, region))
}

// GetCity will return the city name based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetCity(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, city))
}

// GetISP will return the Internet Service Provider name based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetISP(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, isp))
}

// GetLatitude will return the latitude based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetLatitude(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, latitude))
}

// GetLongitude will return the longitude based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetLongitude(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, longitude))
}

// GetDomain will return the domain name based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetDomain(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, domain))
}

// GetZipCode will return the postal code based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetZipCode(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, zipcode))
}

// GetTimezone will return the time zone based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetTimezone(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, timezone))
}

// GetNetSpeed will return the Internet connection speed based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetNetSpeed(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, netSpeed))
}

// GetIDDCode will return the International Direct Dialing code based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetIDDCode(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, iddCode))
}

// GetAreaCode will return the area code based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetAreaCode(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, areaCode))
}

// GetWeatherStationCode will return the weather station code based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetWeatherStationCode(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, weatherStationCode))
}

// GetWeatherStationName will return the weather station name based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetWeatherStationName(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, weatherStationName))
}

// GetMCC will return the mobile country code based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetMCC(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, mcc))
}

// GetMNC will return the mobile network code based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetMNC(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, mnc))
}

// GetMobileBrand will return the mobile carrier brand based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetMobileBrand(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, mobileBrand))
}

// GetElevation will return the elevation in meters based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetElevation(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, elevation))
}

// GetUsageType will return the usage type based on the queried IP address.
//
// Deprecated: No longer being updated.
func GetUsageType(ipaddress string) Record {
	return handleError(defaultDB.query(ipaddress, usageType))
}

// Close will close the file handle to the BIN file.
//
// Deprecated: No longer being updated.
func Close() {
	defaultDB.Close()
}