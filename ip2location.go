package ip2location

import (
	"fmt"
	"os"
	"bytes"
	"encoding/binary"
	"math/big"
	"strconv"
	"net"
)

type ip2locationmeta struct {
	databasetype uint8
	databasecolumn uint8
	databaseday uint8
	databasemonth uint8
	databaseyear uint8
	ipv4databasecount uint32
	ipv4databaseaddr uint32
	ipv6databasecount uint32
	ipv6databaseaddr uint32
	ipv4indexbaseaddr uint32
	ipv6indexbaseaddr uint32
	ipv4columnsize uint32
	ipv6columnsize uint32
}

var f *os.File
var meta ip2locationmeta

var country_position = [25]uint8{0, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
var region_position = [25]uint8{0, 0, 0, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}
var city_position = [25]uint8{0, 0, 0, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}
var isp_position = [25]uint8{0, 0, 3, 0, 5, 0, 7, 5, 7, 0, 8, 0, 9, 0, 9, 0, 9, 0, 9, 7, 9, 0, 9, 7, 9}
var latitude_position = [25]uint8{0, 0, 0, 0, 0, 5, 5, 0, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
var longitude_position = [25]uint8{0, 0, 0, 0, 0, 6, 6, 0, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6}
var domain_position = [25]uint8{0, 0, 0, 0, 0, 0, 0, 6, 8, 0, 9, 0, 10,0, 10, 0, 10, 0, 10, 8, 10, 0, 10, 8, 10}
var zipcode_position = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 7, 7, 7, 0, 7, 7, 7, 0, 7, 0, 7, 7, 7, 0, 7}
var timezone_position = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 8, 7, 8, 8, 8, 7, 8, 0, 8, 8, 8, 0, 8}
var netspeed_position = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 11,0, 11,8, 11, 0, 11, 0, 11, 0, 11}
var iddcode_position = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 12, 0, 12, 0, 12, 9, 12, 0, 12}
var areacode_position = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10 ,13 ,0, 13, 0, 13, 10, 13, 0, 13}
var weatherstationcode_position = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 14, 0, 14, 0, 14, 0, 14}
var weatherstationname_position = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 15, 0, 15, 0, 15, 0, 15}
var mcc_position = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 16, 0, 16, 9, 16}
var mnc_position = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10,17, 0, 17, 10, 17}
var mobilebrand_position = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11,18, 0, 18, 11, 18}
var elevation_position = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 19, 0, 19}
var usagetype_position = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 12, 20}

const api_version string = "8.0.3"

var max_ipv4_range = big.NewInt(4294967295)
var max_ipv6_range = big.NewInt(0)

const invalid_address string = "Invalid IP address."
const missing_file string = "Invalid database file."
const not_supported string = "This parameter is unavailable for selected data file. Please upgrade the data file."

var metaok bool

var country_position_offset uint32
var region_position_offset uint32
var city_position_offset uint32
var isp_position_offset uint32
var domain_position_offset uint32
var zipcode_position_offset uint32
var latitude_position_offset uint32
var longitude_position_offset uint32
var timezone_position_offset uint32
var netspeed_position_offset uint32
var iddcode_position_offset uint32
var areacode_position_offset uint32
var weatherstationcode_position_offset uint32
var weatherstationname_position_offset uint32
var mcc_position_offset uint32
var mnc_position_offset uint32
var mobilebrand_position_offset uint32
var elevation_position_offset uint32
var usagetype_position_offset uint32

var country_enabled bool
var region_enabled bool
var city_enabled bool
var isp_enabled bool
var domain_enabled bool
var zipcode_enabled bool
var latitude_enabled bool
var longitude_enabled bool
var timezone_enabled bool
var netspeed_enabled bool
var iddcode_enabled bool
var areacode_enabled bool
var weatherstationcode_enabled bool
var weatherstationname_enabled bool
var mcc_enabled bool
var mnc_enabled bool
var mobilebrand_enabled bool
var elevation_enabled bool
var usagetype_enabled bool

// get IP type and calculate IP number; calculates index too if exists
func checkip(ip string) (iptype uint32, ipnum *big.Int, ipindex uint32) {
	iptype = 0
	ipnum = big.NewInt(0)
	ipnumtmp := big.NewInt(0)
	ipindex = 0
	ipaddress := net.ParseIP(ip)
	
	if ipaddress != nil {
		v4 := ipaddress.To4()
		
		if v4 != nil {
			iptype = 4
			ipnum.SetBytes(v4)
		} else {
			v6 := ipaddress.To16()
			
			if v6 != nil {
				iptype = 6
				ipnum.SetBytes(v6)
			}
		}
	}
	if iptype == 4 {
		if meta.ipv4indexbaseaddr > 0 {
			ipnumtmp.Rsh(ipnum, 16)
			ipnumtmp.Lsh(ipnumtmp, 3)
			ipindex = uint32(ipnumtmp.Add(ipnumtmp, big.NewInt(int64(meta.ipv4indexbaseaddr))).Uint64())
		}
	} else if iptype == 6 {
		if meta.ipv6indexbaseaddr > 0 {
			ipnumtmp.Rsh(ipnum, 112)
			ipnumtmp.Lsh(ipnumtmp, 3)
			ipindex = uint32(ipnumtmp.Add(ipnumtmp, big.NewInt(int64(meta.ipv6indexbaseaddr))).Uint64())
		}
	}
	return 
}

// read byte
func readuint8(pos int64) uint8 {
	var retval uint8
	data := make([]byte, 1)
	_, err := f.ReadAt(data, pos - 1)
	if err != nil {
		fmt.Println("File read failed:", err)
	}
	retval = data[0]
	return retval
}

// read unsigned 32-bit integer
func readuint32(pos uint32) uint32 {
	pos2 := int64(pos)
	var retval uint32
	data := make([]byte, 4)
	_, err := f.ReadAt(data, pos2 - 1)
	if err != nil {
		fmt.Println("File read failed:", err)
	}
	buf := bytes.NewReader(data)
	err = binary.Read(buf, binary.LittleEndian, &retval)
	if err != nil {
		fmt.Println("Binary read failed:", err)
	}
	return retval
}

// read unsigned 128-bit integer
func readuint128(pos uint32) *big.Int {
	pos2 := int64(pos)
	retval := big.NewInt(0)
	data := make([]byte, 16)
	_, err := f.ReadAt(data, pos2 - 1)
	if err != nil {
		fmt.Println("File read failed:", err)
	}
	
	// little endian to big endian
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	retval.SetBytes(data)
	return retval
}

// read string
func readstr(pos uint32) string {
	pos2 := int64(pos)
	var retval string
	lenbyte := make([]byte, 1)
	_, err := f.ReadAt(lenbyte, pos2)
	if err != nil {
		fmt.Println("File read failed:", err)
	}
	strlen := lenbyte[0]
	data := make([]byte, strlen)
	_, err = f.ReadAt(data, pos2 + 1)
	if err != nil {
		fmt.Println("File read failed:", err)
	}
	retval = string(data[:strlen])
	return retval
}

// read float
func readfloat(pos uint32) float32 {
	pos2 := int64(pos)
	var retval float32
	data := make([]byte, 4)
	_, err := f.ReadAt(data, pos2 - 1)
	if err != nil {
		fmt.Println("File read failed:", err)
	}
	buf := bytes.NewReader(data)
	err = binary.Read(buf, binary.LittleEndian, &retval)
	if err != nil {
		fmt.Println("Binary read failed:", err)
	}
	return retval
}

// initialize the component with the database path
func Open(dbpath string) {
	max_ipv6_range.SetString("340282366920938463463374607431768211455", 10)
	
	var err error
	f, err = os.Open(dbpath)
	if err != nil {
		return
	}
	
	meta.databasetype = readuint8(1)
	meta.databasecolumn = readuint8(2)
	meta.databaseyear = readuint8(3)
	meta.databasemonth = readuint8(4)
	meta.databaseday = readuint8(5)
	meta.ipv4databasecount = readuint32(6)
	meta.ipv4databaseaddr = readuint32(10)
	meta.ipv6databasecount = readuint32(14)
	meta.ipv6databaseaddr = readuint32(18)
	meta.ipv4indexbaseaddr = readuint32(22)
	meta.ipv6indexbaseaddr = readuint32(26)
	meta.ipv4columnsize = uint32(meta.databasecolumn << 2) // 4 bytes each column
	meta.ipv6columnsize = uint32(16 + ((meta.databasecolumn - 1) << 2)) // 4 bytes each column, except IPFrom column which is 16 bytes
	
	dbt := meta.databasetype
	
	// since both IPv4 and IPv6 use 4 bytes for the below columns, can just do it once here
	if country_position[dbt] != 0 {
		country_position_offset = uint32(country_position[dbt] - 1) << 2
		country_enabled = true
	}
	if region_position[dbt] != 0 {
		region_position_offset = uint32(region_position[dbt] - 1) << 2
		region_enabled = true
	}
	if city_position[dbt] != 0 {
		city_position_offset = uint32(city_position[dbt] - 1) << 2
		city_enabled = true
	}
	if isp_position[dbt] != 0 {
		isp_position_offset = uint32(isp_position[dbt] - 1) << 2
		isp_enabled = true
	}
	if domain_position[dbt] != 0 {
		domain_position_offset = uint32(domain_position[dbt] - 1) << 2
		domain_enabled = true
	}
	if zipcode_position[dbt] != 0 {
		zipcode_position_offset = uint32(zipcode_position[dbt] - 1) << 2
		zipcode_enabled = true
	}
	if latitude_position[dbt] != 0 {
		latitude_position_offset = uint32(latitude_position[dbt] - 1) << 2
		latitude_enabled = true
	}
	if longitude_position[dbt] != 0 {
		longitude_position_offset = uint32(longitude_position[dbt] - 1) << 2
		longitude_enabled = true
	}
	if timezone_position[dbt] != 0 {
		timezone_position_offset = uint32(timezone_position[dbt] - 1) << 2
		timezone_enabled = true
	}
	if netspeed_position[dbt] != 0 {
		netspeed_position_offset = uint32(netspeed_position[dbt] - 1) << 2
		netspeed_enabled = true
	}
	if iddcode_position[dbt] != 0 {
		iddcode_position_offset = uint32(iddcode_position[dbt] - 1) << 2
		iddcode_enabled = true
	}
	if areacode_position[dbt] != 0 {
		areacode_position_offset = uint32(areacode_position[dbt] - 1) << 2
		areacode_enabled = true
	}
	if weatherstationcode_position[dbt] != 0 {
		weatherstationcode_position_offset = uint32(weatherstationcode_position[dbt] - 1) << 2
		weatherstationcode_enabled = true
	}
	if weatherstationname_position[dbt] != 0 {
		weatherstationname_position_offset = uint32(weatherstationname_position[dbt] - 1) << 2
		weatherstationname_enabled = true
	}
	if mcc_position[dbt] != 0 {
		mcc_position_offset = uint32(mcc_position[dbt] - 1) << 2
		mcc_enabled = true
	}
	if mnc_position[dbt] != 0 {
		mnc_position_offset = uint32(mnc_position[dbt] - 1) << 2
		mnc_enabled = true
	}
	if mobilebrand_position[dbt] != 0 {
		mobilebrand_position_offset = uint32(mobilebrand_position[dbt] - 1) << 2
		mobilebrand_enabled = true
	}
	if elevation_position[dbt] != 0 {
		elevation_position_offset = uint32(elevation_position[dbt] - 1) << 2
		elevation_enabled = true
	}
	if usagetype_position[dbt] != 0 {
		usagetype_position_offset = uint32(usagetype_position[dbt] - 1) << 2
		usagetype_enabled = true
	}
	
	metaok = true
}

// close database file handle
func Close() {
	f.Close()
}

// get api version
func Api_version() string {
	return api_version
}

// populate record with message
func loadmessage (mesg string) Record {
	var x Record
	
	x.CountryShort = mesg
	x.CountryLong = mesg
	x.Region = mesg
	x.City = mesg
	x.Isp = mesg
	x.Domain = mesg
	x.ZipCode = mesg
	x.TimeZone = mesg
	x.NetSpeed = mesg
	x.IddCode = mesg
	x.AreaCode = mesg
	x.WeatherStationCode = mesg
	x.WeatherStationName = mesg
	x.Mcc = mesg
	x.Mnc = mesg
	x.MobileBrand = mesg
	x.UsageType = mesg
	
	return x
}

// get All fields
func Get_all(ipaddress string) Record {
	return query(ipaddress, All)
}

// get country code
func Get_country_short(ipaddress string) Record {
	return query(ipaddress, CountryShort)
}

// get country name
func Get_country_long(ipaddress string) Record {
	return query(ipaddress, CountryLong)
}

// get Region
func Get_region(ipaddress string) Record {
	return query(ipaddress, Region)
}

// get City
func Get_city(ipaddress string) Record {
	return query(ipaddress, City)
}

// get Isp
func Get_isp(ipaddress string) Record {
	return query(ipaddress, Isp)
}

// get Latitude
func Get_latitude(ipaddress string) Record {
	return query(ipaddress, Latitude)
}

// get Longitude
func Get_longitude(ipaddress string) Record {
	return query(ipaddress, Longitude)
}

// get Domain
func Get_domain(ipaddress string) Record {
	return query(ipaddress, Domain)
}

// get zip code
func Get_zipcode(ipaddress string) Record {
	return query(ipaddress, ZipCode)
}

// get time zone
func Get_timezone(ipaddress string) Record {
	return query(ipaddress, TimeZone)
}

// get net speed
func Get_netspeed(ipaddress string) Record {
	return query(ipaddress, NetSpeed)
}

// get idd code
func Get_iddcode(ipaddress string) Record {
	return query(ipaddress, IddCode)
}

// get area code
func Get_areacode(ipaddress string) Record {
	return query(ipaddress, AreaCode)
}

// get weather station code
func Get_weatherstationcode(ipaddress string) Record {
	return query(ipaddress, WeatherStationCode)
}

// get weather station name
func Get_weatherstationname(ipaddress string) Record {
	return query(ipaddress, WeatherStationName)
}

// get mobile country code
func Get_mcc(ipaddress string) Record {
	return query(ipaddress, Mcc)
}

// get mobile network code
func Get_mnc(ipaddress string) Record {
	return query(ipaddress, Mnc)
}

// get mobile carrier brand
func Get_mobilebrand(ipaddress string) Record {
	return query(ipaddress, MobileBrand)
}

// get Elevation
func Get_elevation(ipaddress string) Record {
	return query(ipaddress, Elevation)
}

// get usage type
func Get_usagetype(ipaddress string) Record {
	return query(ipaddress, UsageType)
}

// main query
func query(ipaddress string, mode uint32) Record {
	x := loadmessage(not_supported) // default message
	
	// read metadata
	if !metaok {
		x = loadmessage(missing_file)
		return x
	}
	
	// check IP type and return IP number & index (if exists)
	iptype, ipno, ipindex := checkip(ipaddress)
	
	if iptype == 0 {
		x = loadmessage(invalid_address)
		return x
	}
	
	var colsize uint32
	var baseaddr uint32
	var low uint32
	var high uint32
	var mid uint32
	var rowoffset uint32
	var rowoffset2 uint32
	ipfrom := big.NewInt(0)
	ipto := big.NewInt(0)
	maxip := big.NewInt(0)
	
	if iptype == 4 {
		baseaddr = meta.ipv4databaseaddr
		high = meta.ipv4databasecount
		maxip = max_ipv4_range
		colsize = meta.ipv4columnsize
	} else {
		baseaddr = meta.ipv6databaseaddr
		high = meta.ipv6databasecount
		maxip = max_ipv6_range
		colsize = meta.ipv6columnsize
	}
	
	// reading index
	if ipindex > 0 {
		low = readuint32(ipindex)
		high = readuint32(ipindex + 4)
	}
	
	if ipno.Cmp(maxip)>=0 {
		ipno = ipno.Sub(ipno, big.NewInt(1))
	}
	
	for low <= high {
		mid = ((low + high) >> 1)
		rowoffset = baseaddr + (mid * colsize)
		rowoffset2 = rowoffset + colsize
		
		if iptype == 4 {
			ipfrom = big.NewInt(int64(readuint32(rowoffset)))
			ipto = big.NewInt(int64(readuint32(rowoffset2)))
		} else {
			ipfrom = readuint128(rowoffset)
			ipto = readuint128(rowoffset2)
		}
		
		if ipno.Cmp(ipfrom)>=0 && ipno.Cmp(ipto)<0 {
			if iptype == 6 {
				rowoffset = rowoffset + 12 // coz below is assuming All columns are 4 bytes, so got 12 left to go to make 16 bytes total
			}
			
			if mode&CountryShort == 1 && country_enabled {
				x.CountryShort = readstr(readuint32(rowoffset + country_position_offset))
			}
			
			if mode&CountryLong != 0 && country_enabled {
				x.CountryLong = readstr(readuint32(rowoffset + country_position_offset) + 3)
			}
			
			if mode&Region != 0 && region_enabled {
				x.Region = readstr(readuint32(rowoffset + region_position_offset))
			}
			
			if mode&City != 0 && city_enabled {
				x.City = readstr(readuint32(rowoffset + city_position_offset))
			}
			
			if mode&Isp != 0 && isp_enabled {
				x.Isp = readstr(readuint32(rowoffset + isp_position_offset))
			}
			
			if mode&Latitude != 0 && latitude_enabled {
				x.Latitude = readfloat(rowoffset + latitude_position_offset)
			}
			
			if mode&Longitude != 0 && longitude_enabled {
				x.Longitude = readfloat(rowoffset + longitude_position_offset)
			}
			
			if mode&Domain != 0 && domain_enabled {
				x.Domain = readstr(readuint32(rowoffset + domain_position_offset))
			}
			
			if mode&ZipCode != 0 && zipcode_enabled {
				x.ZipCode = readstr(readuint32(rowoffset + zipcode_position_offset))
			}
			
			if mode&TimeZone != 0 && timezone_enabled {
				x.TimeZone = readstr(readuint32(rowoffset + timezone_position_offset))
			}
			
			if mode&NetSpeed != 0 && netspeed_enabled {
				x.NetSpeed = readstr(readuint32(rowoffset + netspeed_position_offset))
			}
			
			if mode&IddCode != 0 && iddcode_enabled {
				x.IddCode = readstr(readuint32(rowoffset + iddcode_position_offset))
			}
			
			if mode&AreaCode != 0 && areacode_enabled {
				x.AreaCode = readstr(readuint32(rowoffset + areacode_position_offset))
			}
			
			if mode&WeatherStationCode != 0 && weatherstationcode_enabled {
				x.WeatherStationCode = readstr(readuint32(rowoffset + weatherstationcode_position_offset))
			}
			
			if mode&WeatherStationName != 0 && weatherstationname_enabled {
				x.WeatherStationName = readstr(readuint32(rowoffset + weatherstationname_position_offset))
			}
			
			if mode&Mcc != 0 && mcc_enabled {
				x.Mcc = readstr(readuint32(rowoffset + mcc_position_offset))
			}
			
			if mode&Mnc != 0 && mnc_enabled {
				x.Mnc = readstr(readuint32(rowoffset + mnc_position_offset))
			}
			
			if mode&MobileBrand != 0 && mobilebrand_enabled {
				x.MobileBrand = readstr(readuint32(rowoffset + mobilebrand_position_offset))
			}
			
			if mode&Elevation != 0 && elevation_enabled {
				f, _ := strconv.ParseFloat(readstr(readuint32(rowoffset + elevation_position_offset)), 32)
				x.Elevation = float32(f)
			}
			
			if mode&UsageType != 0 && usagetype_enabled {
				x.UsageType = readstr(readuint32(rowoffset + usagetype_position_offset))
			}
			
			return x
		} else {
			if ipno.Cmp(ipfrom)<0 {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return x
}

// for debugging purposes
func Printrecord(x Record) {
	fmt.Printf("country_short: %s\n", x.CountryShort)
	fmt.Printf("country_long: %s\n", x.CountryLong)
	fmt.Printf("Region: %s\n", x.Region)
	fmt.Printf("City: %s\n", x.City)
	fmt.Printf("Isp: %s\n", x.Isp)
	fmt.Printf("Latitude: %f\n", x.Latitude)
	fmt.Printf("Longitude: %f\n", x.Longitude)
	fmt.Printf("Domain: %s\n", x.Domain)
	fmt.Printf("ZipCode: %s\n", x.ZipCode)
	fmt.Printf("TimeZone: %s\n", x.TimeZone)
	fmt.Printf("NetSpeed: %s\n", x.NetSpeed)
	fmt.Printf("IddCode: %s\n", x.IddCode)
	fmt.Printf("AreaCode: %s\n", x.AreaCode)
	fmt.Printf("WeatherStationCode: %s\n", x.WeatherStationCode)
	fmt.Printf("WeatherStationName: %s\n", x.WeatherStationName)
	fmt.Printf("Mcc: %s\n", x.Mcc)
	fmt.Printf("Mnc: %s\n", x.Mnc)
	fmt.Printf("MobileBrand: %s\n", x.MobileBrand)
	fmt.Printf("Elevation: %f\n", x.Elevation)
	fmt.Printf("UsageType: %s\n", x.UsageType)
}
