// This ip2location package provides a fast lookup of country, region, city, latitude, longitude, ZIP code, time zone,
// ISP, domain name, connection type, IDD code, area code, weather station code, station name, MCC, MNC,
// mobile brand, elevation, and usage type from IP address by using IP2Location database.
package ip2location

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"math/big"
	"net"
	"os"
	"strconv"
)

const apiVersion string = "8.4.0"

var countryPosition = [25]uint8{0, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
var regionPosition = [25]uint8{0, 0, 0, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}
var cityPosition = [25]uint8{0, 0, 0, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}
var ispPosition = [25]uint8{0, 0, 3, 0, 5, 0, 7, 5, 7, 0, 8, 0, 9, 0, 9, 0, 9, 0, 9, 7, 9, 0, 9, 7, 9}
var latitudePosition = [25]uint8{0, 0, 0, 0, 0, 5, 5, 0, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}
var longitudePosition = [25]uint8{0, 0, 0, 0, 0, 6, 6, 0, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6}
var domainPosition = [25]uint8{0, 0, 0, 0, 0, 0, 0, 6, 8, 0, 9, 0, 10, 0, 10, 0, 10, 0, 10, 8, 10, 0, 10, 8, 10}
var zipCodePosition = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 7, 7, 7, 7, 0, 7, 7, 7, 0, 7, 0, 7, 7, 7, 0, 7}
var timezonePosition = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 8, 7, 8, 8, 8, 7, 8, 0, 8, 8, 8, 0, 8}
var netSpeedPosition = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 8, 11, 0, 11, 8, 11, 0, 11, 0, 11, 0, 11}
var iddCodePosition = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 12, 0, 12, 0, 12, 9, 12, 0, 12}
var areaCodePosition = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 13, 0, 13, 0, 13, 10, 13, 0, 13}
var weatherStationCodePosition = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 14, 0, 14, 0, 14, 0, 14}
var weatherStationMamePosition = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 15, 0, 15, 0, 15, 0, 15}
var mccPosition = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 9, 16, 0, 16, 9, 16}
var mncPosition = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 17, 0, 17, 10, 17}
var mobileBrandPosition = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 18, 0, 18, 11, 18}
var elevationPosition = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 11, 19, 0, 19}
var usageTypePosition = [25]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 12, 20}

var maxIPV4Range = big.NewInt(4294967295)
var maxIPV6Range = big.NewInt(0)
var fromV4Mapped = big.NewInt(281470681743360)
var toV4Mapped = big.NewInt(281474976710655)
var from6To4 = big.NewInt(0)
var to6To4 = big.NewInt(0)
var fromTeredo = big.NewInt(0)
var toTeredo = big.NewInt(0)
var last32Bits = big.NewInt(4294967295)

const countryShort uint32 = 0x00001
const countryLong uint32 = 0x00002
const region uint32 = 0x00004
const city uint32 = 0x00008
const isp uint32 = 0x00010
const latitude uint32 = 0x00020
const longitude uint32 = 0x00040
const domain uint32 = 0x00080
const zipcode uint32 = 0x00100
const timezone uint32 = 0x00200
const netSpeed uint32 = 0x00400
const iddCode uint32 = 0x00800
const areaCode uint32 = 0x01000
const weatherStationCode uint32 = 0x02000
const weatherStationName uint32 = 0x04000
const mcc uint32 = 0x08000
const mnc uint32 = 0x10000
const mobileBrand uint32 = 0x20000
const elevation uint32 = 0x40000
const usageType uint32 = 0x80000

const all uint32 = countryShort | countryLong | region | city | isp | latitude | longitude | domain | zipcode | timezone | netSpeed | iddCode | areaCode | weatherStationCode | weatherStationName | mcc | mnc | mobileBrand | elevation | usageType

const invalidAddress string = "Invalid IP address."
const missingFile string = "Invalid database file."
const notSupported string = "This parameter is unavailable for selected data file. Please upgrade the data file."

type DBReader interface {
	io.ReadCloser
	io.ReaderAt
}

type DB struct {
	f    DBReader
	meta meta

	countryPositionOffset            uint32
	regionPositionOffset             uint32
	cityPositionOffset               uint32
	ispPositionOffset                uint32
	domainPositionOffset             uint32
	zipcodePositionOffset            uint32
	latitudePositionOffset           uint32
	longitudePositionOffset          uint32
	timezonePositionOffset           uint32
	netSpeedPositionOffset           uint32
	iddCodePositionOffset            uint32
	areaCodePositionOffset           uint32
	weatherStationCodePositionOffset uint32
	weatherStationNamePositionOffset uint32
	mccPositionOffset                uint32
	mncPositionOffset                uint32
	mobileBrandPositionOffset        uint32
	elevationPositionOffset          uint32
	usageTypePositionOffset          uint32

	countryEnabled            bool
	regionEnabled             bool
	cityEnabled               bool
	ispEnabled                bool
	domainEnabled             bool
	zipcodeEnabled            bool
	latitudeEnabled           bool
	longitudeEnabled          bool
	timezoneEnabled           bool
	netSpeedEnabled           bool
	iddCodeEnabled            bool
	areaCodeEnabled           bool
	weatherStationCodeEnabled bool
	weatherStationNameEnabled bool
	mccEnabled                bool
	mncEnabled                bool
	mobileBrandEnabled        bool
	elevationEnabled          bool
	usageTypeEnabled          bool

	metaOK bool
}

// Open takes the path to the IP2Location BIN database file. It will read all the metadata required to
// be able to extract the embedded geolocation data, and return the underlining DB object.
func OpenDB(dbpath string) (*DB, error) {
	f, err := os.Open(dbpath)
	if err != nil {
		return nil, err
	}

	return OpenDBWithReader(f)
}

// OpenDBWithReader takes a DBReader to the IP2Location BIN database file. It will read all the metadata required to
// be able to extract the embedded geolocation data, and return the underlining DB object.
func OpenDBWithReader(reader DBReader) (*DB, error) {
	var db = &DB{}

	maxIPV6Range.SetString("340282366920938463463374607431768211455", 10)
	from6To4.SetString("42545680458834377588178886921629466624", 10)
	to6To4.SetString("42550872755692912415807417417958686719", 10)
	fromTeredo.SetString("42540488161975842760550356425300246528", 10)
	toTeredo.SetString("42540488241204005274814694018844196863", 10)

	db.f = reader

	var err error
	db.meta.databaseType, err = db.readuint8(1)
	if err != nil {
		return fatal(db, err)
	}
	db.meta.databaseColumn, err = db.readuint8(2)
	if err != nil {
		return fatal(db, err)
	}
	db.meta.databaseYear, err = db.readuint8(3)
	if err != nil {
		return fatal(db, err)
	}
	db.meta.databaseMonth, err = db.readuint8(4)
	if err != nil {
		return fatal(db, err)
	}
	db.meta.databaseDay, err = db.readuint8(5)
	if err != nil {
		return fatal(db, err)
	}
	db.meta.ipv4DatabaseCount, err = db.readuint32(6)
	if err != nil {
		return fatal(db, err)
	}
	db.meta.ipv4DatabaseAddr, err = db.readuint32(10)
	if err != nil {
		return fatal(db, err)
	}
	db.meta.ipv6DatabaseCount, err = db.readuint32(14)
	if err != nil {
		return fatal(db, err)
	}
	db.meta.ipv6DatabaseAddr, err = db.readuint32(18)
	if err != nil {
		return fatal(db, err)
	}
	db.meta.ipv4IndexBaseAddr, err = db.readuint32(22)
	if err != nil {
		return fatal(db, err)
	}
	db.meta.ipv6IndexBaseAddr, err = db.readuint32(26)
	if err != nil {
		return fatal(db, err)
	}
	db.meta.ipv4ColumnSize = uint32(db.meta.databaseColumn << 2)              // 4 bytes each column
	db.meta.ipv6ColumnSize = uint32(16 + ((db.meta.databaseColumn - 1) << 2)) // 4 bytes each column, except IPFrom column which is 16 bytes

	dbt := db.meta.databaseType

	if countryPosition[dbt] != 0 {
		db.countryPositionOffset = uint32(countryPosition[dbt]-2) << 2
		db.countryEnabled = true
	}
	if regionPosition[dbt] != 0 {
		db.regionPositionOffset = uint32(regionPosition[dbt]-2) << 2
		db.regionEnabled = true
	}
	if cityPosition[dbt] != 0 {
		db.cityPositionOffset = uint32(cityPosition[dbt]-2) << 2
		db.cityEnabled = true
	}
	if ispPosition[dbt] != 0 {
		db.ispPositionOffset = uint32(ispPosition[dbt]-2) << 2
		db.ispEnabled = true
	}
	if domainPosition[dbt] != 0 {
		db.domainPositionOffset = uint32(domainPosition[dbt]-2) << 2
		db.domainEnabled = true
	}
	if zipCodePosition[dbt] != 0 {
		db.zipcodePositionOffset = uint32(zipCodePosition[dbt]-2) << 2
		db.zipcodeEnabled = true
	}
	if latitudePosition[dbt] != 0 {
		db.latitudePositionOffset = uint32(latitudePosition[dbt]-2) << 2
		db.latitudeEnabled = true
	}
	if longitudePosition[dbt] != 0 {
		db.longitudePositionOffset = uint32(longitudePosition[dbt]-2) << 2
		db.longitudeEnabled = true
	}
	if timezonePosition[dbt] != 0 {
		db.timezonePositionOffset = uint32(timezonePosition[dbt]-2) << 2
		db.timezoneEnabled = true
	}
	if netSpeedPosition[dbt] != 0 {
		db.netSpeedPositionOffset = uint32(netSpeedPosition[dbt]-2) << 2
		db.netSpeedEnabled = true
	}
	if iddCodePosition[dbt] != 0 {
		db.iddCodePositionOffset = uint32(iddCodePosition[dbt]-2) << 2
		db.iddCodeEnabled = true
	}
	if areaCodePosition[dbt] != 0 {
		db.areaCodePositionOffset = uint32(areaCodePosition[dbt]-2) << 2
		db.areaCodeEnabled = true
	}
	if weatherStationCodePosition[dbt] != 0 {
		db.weatherStationCodePositionOffset = uint32(weatherStationCodePosition[dbt]-2) << 2
		db.weatherStationCodeEnabled = true
	}
	if weatherStationMamePosition[dbt] != 0 {
		db.weatherStationNamePositionOffset = uint32(weatherStationMamePosition[dbt]-2) << 2
		db.weatherStationNameEnabled = true
	}
	if mccPosition[dbt] != 0 {
		db.mccPositionOffset = uint32(mccPosition[dbt]-2) << 2
		db.mccEnabled = true
	}
	if mncPosition[dbt] != 0 {
		db.mncPositionOffset = uint32(mncPosition[dbt]-2) << 2
		db.mncEnabled = true
	}
	if mobileBrandPosition[dbt] != 0 {
		db.mobileBrandPositionOffset = uint32(mobileBrandPosition[dbt]-2) << 2
		db.mobileBrandEnabled = true
	}
	if elevationPosition[dbt] != 0 {
		db.elevationPositionOffset = uint32(elevationPosition[dbt]-2) << 2
		db.elevationEnabled = true
	}
	if usageTypePosition[dbt] != 0 {
		db.usageTypePositionOffset = uint32(usageTypePosition[dbt]-2) << 2
		db.usageTypeEnabled = true
	}

	db.metaOK = true

	return db, nil
}

// GetAll will return all geolocation fields based on the queried IP address.
func (d *DB) GetAll(ipaddress string) (Record, error) {
	return d.query(ipaddress, all)
}

// GetCountryShort will return the ISO-3166 country code based on the queried IP address.
func (d *DB) GetCountryShort(ipaddress string) (Record, error) {
	return d.query(ipaddress, countryShort)
}

// GetCountryLong will return the country name based on the queried IP address.
func (d *DB) GetCountryLong(ipaddress string) (Record, error) {
	return d.query(ipaddress, countryLong)
}

// GetRegion will return the region name based on the queried IP address.
func (d *DB) GetRegion(ipaddress string) (Record, error) {
	return d.query(ipaddress, region)
}

// GetCity will return the city name based on the queried IP address.
func (d *DB) GetCity(ipaddress string) (Record, error) {
	return d.query(ipaddress, city)
}

// GetISP will return the Internet Service Provider name based on the queried IP address.
func (d *DB) GetISP(ipaddress string) (Record, error) {
	return d.query(ipaddress, isp)
}

// GetLatitude will return the latitude based on the queried IP address.
func (d *DB) GetLatitude(ipaddress string) (Record, error) {
	return d.query(ipaddress, latitude)
}

// GetLongitude will return the longitude based on the queried IP address.
func (d *DB) GetLongitude(ipaddress string) (Record, error) {
	return d.query(ipaddress, longitude)
}

// GetDomain will return the domain name based on the queried IP address.
func (d *DB) GetDomain(ipaddress string) (Record, error) {
	return d.query(ipaddress, domain)
}

// GetZipCode will return the postal code based on the queried IP address.
func (d *DB) GetZipCode(ipaddress string) (Record, error) {
	return d.query(ipaddress, zipcode)
}

// GetTimezone will return the time zone based on the queried IP address.
func (d *DB) GetTimezone(ipaddress string) (Record, error) {
	return d.query(ipaddress, timezone)
}

// GetNetSpeed will return the Internet connection speed based on the queried IP address.
func (d *DB) GetNetSpeed(ipaddress string) (Record, error) {
	return d.query(ipaddress, netSpeed)
}

// GetIDDCode will return the International Direct Dialing code based on the queried IP address.
func (d *DB) GetIDDCode(ipaddress string) (Record, error) {
	return d.query(ipaddress, iddCode)
}

// GetAreaCode will return the area code based on the queried IP address.
func (d *DB) GetAreaCode(ipaddress string) (Record, error) {
	return d.query(ipaddress, areaCode)
}

// GetWeatherStationCode will return the weather station code based on the queried IP address.
func (d *DB) GetWeatherStationCode(ipaddress string) (Record, error) {
	return d.query(ipaddress, weatherStationCode)
}

// GetWeatherStationName will return the weather station name based on the queried IP address.
func (d *DB) GetWeatherStationName(ipaddress string) (Record, error) {
	return d.query(ipaddress, weatherStationName)
}

// GetMCC will return the mobile country code based on the queried IP address.
func (d *DB) GetMCC(ipaddress string) (Record, error) {
	return d.query(ipaddress, mcc)
}

// GetMNC will return the mobile network code based on the queried IP address.
func (d *DB) GetMNC(ipaddress string) (Record, error) {
	return d.query(ipaddress, mnc)
}

// GetMobileBrand will return the mobile carrier brand based on the queried IP address.
func (d *DB) GetMobileBrand(ipaddress string) (Record, error) {
	return d.query(ipaddress, mobileBrand)
}

// GetElevation will return the elevation in meters based on the queried IP address.
func (d *DB) GetElevation(ipaddress string) (Record, error) {
	return d.query(ipaddress, elevation)
}

// GetUsageType will return the usage type based on the queried IP address.
func (d *DB) GetUsageType(ipaddress string) (Record, error) {
	return d.query(ipaddress, usageType)
}

// checkIP - get IP type and calculate IP number; calculates index too if exists
func (d *DB) checkIP(ip string) (iptype uint32, ipnum *big.Int, ipindex uint32) {
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

				if ipnum.Cmp(fromV4Mapped) >= 0 && ipnum.Cmp(toV4Mapped) <= 0 {
					// ipv4-mapped ipv6 should treat as ipv4 and read ipv4 data section
					iptype = 4
					ipnum.Sub(ipnum, fromV4Mapped)
				} else if ipnum.Cmp(from6To4) >= 0 && ipnum.Cmp(to6To4) <= 0 {
					// 6to4 so need to remap to ipv4
					iptype = 4
					ipnum.Rsh(ipnum, 80)
					ipnum.And(ipnum, last32Bits)
				} else if ipnum.Cmp(fromTeredo) >= 0 && ipnum.Cmp(toTeredo) <= 0 {
					// Teredo so need to remap to ipv4
					iptype = 4
					ipnum.Not(ipnum)
					ipnum.And(ipnum, last32Bits)
				}
			}
		}
	}
	if iptype == 4 {
		if d.meta.ipv4IndexBaseAddr > 0 {
			ipnumtmp.Rsh(ipnum, 16)
			ipnumtmp.Lsh(ipnumtmp, 3)
			ipindex = uint32(ipnumtmp.Add(ipnumtmp, big.NewInt(int64(d.meta.ipv4IndexBaseAddr))).Uint64())
		}
	} else if iptype == 6 {
		if d.meta.ipv6IndexBaseAddr > 0 {
			ipnumtmp.Rsh(ipnum, 112)
			ipnumtmp.Lsh(ipnumtmp, 3)
			ipindex = uint32(ipnumtmp.Add(ipnumtmp, big.NewInt(int64(d.meta.ipv6IndexBaseAddr))).Uint64())
		}
	}
	return
}

// read byte
func (d *DB) readuint8(pos int64) (uint8, error) {
	var retval uint8
	data := make([]byte, 1)
	_, err := d.f.ReadAt(data, pos-1)
	if err != nil {
		return 0, err
	}
	retval = data[0]
	return retval, nil
}

// read unsigned 32-bit integer from slices
func (d *DB) readuint32_row(row []byte, pos uint32) uint32 {
	var retval uint32
	data := row[pos : pos+4]
	retval = binary.LittleEndian.Uint32(data)
	return retval
}

// read unsigned 32-bit integer
func (d *DB) readuint32(pos uint32) (uint32, error) {
	pos2 := int64(pos)
	var retval uint32
	data := make([]byte, 4)
	_, err := d.f.ReadAt(data, pos2-1)
	if err != nil {
		return 0, err
	}
	buf := bytes.NewReader(data)
	err = binary.Read(buf, binary.LittleEndian, &retval)
	if err != nil {
		fmt.Printf("binary read failed: %v", err)
	}
	return retval, nil
}

// read unsigned 128-bit integer
func (d *DB) readuint128(pos uint32) (*big.Int, error) {
	pos2 := int64(pos)
	retval := big.NewInt(0)
	data := make([]byte, 16)
	_, err := d.f.ReadAt(data, pos2-1)
	if err != nil {
		return nil, err
	}

	// little endian to big endian
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	retval.SetBytes(data)
	return retval, nil
}

// read string
func (d *DB) readstr(pos uint32) (string, error) {
	pos2 := int64(pos)
	var retval string
	lenbyte := make([]byte, 1)
	_, err := d.f.ReadAt(lenbyte, pos2)
	if err != nil {
		return "", err
	}
	strlen := lenbyte[0]
	data := make([]byte, strlen)
	_, err = d.f.ReadAt(data, pos2+1)
	if err != nil {
		return "", err
	}
	retval = string(data[:strlen])
	return retval, nil
}

// read float from slices
func (d *DB) readfloat_row(row []byte, pos uint32) float32 {
	var retval float32
	data := row[pos : pos+4]
	bits := binary.LittleEndian.Uint32(data)
	retval = math.Float32frombits(bits)
	return retval
}

// read float
func (d *DB) readfloat(pos uint32) (float32, error) {
	pos2 := int64(pos)
	var retval float32
	data := make([]byte, 4)
	_, err := d.f.ReadAt(data, pos2-1)
	if err != nil {
		return 0, err
	}
	buf := bytes.NewReader(data)
	err = binary.Read(buf, binary.LittleEndian, &retval)
	if err != nil {
		fmt.Printf("binary read failed: %v", err)
	}
	return retval, nil
}

// main query
func (d *DB) query(ipaddress string, mode uint32) (Record, error) {
	x := loadMessage(notSupported) // default message

	// read metadata
	if !d.metaOK {
		x = loadMessage(missingFile)
		return x, nil
	}

	// check IP type and return IP number & index (if exists)
	iptype, ipno, ipindex := d.checkIP(ipaddress)

	if iptype == 0 {
		x = loadMessage(invalidAddress)
		return x, nil
	}

	var err error
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
		baseaddr = d.meta.ipv4DatabaseAddr
		high = d.meta.ipv4DatabaseCount
		maxip = maxIPV4Range
		colsize = d.meta.ipv4ColumnSize
	} else {
		baseaddr = d.meta.ipv6DatabaseAddr
		high = d.meta.ipv6DatabaseCount
		maxip = maxIPV6Range
		colsize = d.meta.ipv6ColumnSize
	}

	// reading index
	if ipindex > 0 {
		low, err = d.readuint32(ipindex)
		if err != nil {
			return x, err
		}
		high, err = d.readuint32(ipindex + 4)
		if err != nil {
			return x, err
		}
	}

	if ipno.Cmp(maxip) >= 0 {
		ipno.Sub(ipno, big.NewInt(1))
	}

	for low <= high {
		mid = ((low + high) >> 1)
		rowoffset = baseaddr + (mid * colsize)
		rowoffset2 = rowoffset + colsize

		if iptype == 4 {
			ipfrom32, err := d.readuint32(rowoffset)
			if err != nil {
				return x, err
			}
			ipfrom = big.NewInt(int64(ipfrom32))

			ipto32, err := d.readuint32(rowoffset2)
			if err != nil {
				return x, err
			}
			ipto = big.NewInt(int64(ipto32))

		} else {
			ipfrom, err = d.readuint128(rowoffset)
			if err != nil {
				return x, err
			}

			ipto, err = d.readuint128(rowoffset2)
			if err != nil {
				return x, err
			}
		}

		if ipno.Cmp(ipfrom) >= 0 && ipno.Cmp(ipto) < 0 {
			var firstcol uint32 = 4 // 4 bytes for ip from
			if iptype == 6 {
				firstcol = 16 // 16 bytes for ipv6
				// rowoffset = rowoffset + 12 // coz below is assuming all columns are 4 bytes, so got 12 left to go to make 16 bytes total
			}

			row := make([]byte, colsize-firstcol) // exclude the ip from field
			_, err := d.f.ReadAt(row, int64(rowoffset+firstcol-1))
			if err != nil {
				return x, err
			}

			if mode&countryShort == 1 && d.countryEnabled {
				// x.CountryShort = readstr(readuint32(rowoffset + countryPositionOffset))
				if x.CountryShort, err = d.readstr(d.readuint32_row(row, d.countryPositionOffset)); err != nil {
					return x, err
				}
			}

			if mode&countryLong != 0 && d.countryEnabled {
				// x.CountryLong = readstr(readuint32(rowoffset + countryPositionOffset) + 3)
				if x.CountryLong, err = d.readstr(d.readuint32_row(row, d.countryPositionOffset) + 3); err != nil {
					return x, err
				}
			}

			if mode&region != 0 && d.regionEnabled {
				// x.Region = readstr(readuint32(rowoffset + regionPositionOffset))
				if x.Region, err = d.readstr(d.readuint32_row(row, d.regionPositionOffset)); err != nil {
					return x, err
				}
			}

			if mode&city != 0 && d.cityEnabled {
				// x.City = readstr(readuint32(rowoffset + cityPositionOffset))
				if x.City, err = d.readstr(d.readuint32_row(row, d.cityPositionOffset)); err != nil {
					return x, err
				}
			}

			if mode&isp != 0 && d.ispEnabled {
				// x.ISP = readstr(readuint32(rowoffset + ispPositionOffset))
				if x.ISP, err = d.readstr(d.readuint32_row(row, d.ispPositionOffset)); err != nil {
					return x, err
				}
			}

			if mode&latitude != 0 && d.latitudeEnabled {
				// x.Latitude = readfloat(rowoffset + latitudePositionOffset)
				x.Latitude = d.readfloat_row(row, d.latitudePositionOffset)
			}

			if mode&longitude != 0 && d.longitudeEnabled {
				// x.Longitude = readfloat(rowoffset + longitudePositionOffset)
				x.Longitude = d.readfloat_row(row, d.longitudePositionOffset)
			}

			if mode&domain != 0 && d.domainEnabled {
				// x.Domain = readstr(readuint32(rowoffset + domainPositionOffset))
				if x.Domain, err = d.readstr(d.readuint32_row(row, d.domainPositionOffset)); err != nil {
					return x, err
				}
			}

			if mode&zipcode != 0 && d.zipcodeEnabled {
				// x.Zipcode = readstr(readuint32(rowoffset + zipcodePositionOffset))
				if x.Zipcode, err = d.readstr(d.readuint32_row(row, d.zipcodePositionOffset)); err != nil {
					return x, err
				}
			}

			if mode&timezone != 0 && d.timezoneEnabled {
				// x.Timezone = readstr(readuint32(rowoffset + timezonePositionOffset))
				if x.Timezone, err = d.readstr(d.readuint32_row(row, d.timezonePositionOffset)); err != nil {
					return x, err
				}
			}

			if mode&netSpeed != 0 && d.netSpeedEnabled {
				// x.NetSpeed = readstr(readuint32(rowoffset + netSpeedPositionOffset))
				if x.NetSpeed, err = d.readstr(d.readuint32_row(row, d.netSpeedPositionOffset)); err != nil {
					return x, err
				}
			}

			if mode&iddCode != 0 && d.iddCodeEnabled {
				// x.IDDCode = readstr(readuint32(rowoffset + iddCodePositionOffset))
				if x.IDDCode, err = d.readstr(d.readuint32_row(row, d.iddCodePositionOffset)); err != nil {
					return x, err
				}
			}

			if mode&areaCode != 0 && d.areaCodeEnabled {
				// x.AreaCode = readstr(readuint32(rowoffset + areaCodePositionOffset))
				if x.AreaCode, err = d.readstr(d.readuint32_row(row, d.areaCodePositionOffset)); err != nil {
					return x, err
				}
			}

			if mode&weatherStationCode != 0 && d.weatherStationCodeEnabled {
				// x.WeatherStationCode = readstr(readuint32(rowoffset + weatherStationCodePositionOffset))
				if x.WeatherStationCode, err = d.readstr(d.readuint32_row(row, d.weatherStationCodePositionOffset)); err != nil {
					return x, err
				}
			}

			if mode&weatherStationName != 0 && d.weatherStationNameEnabled {
				// x.WeatherStationName = readstr(readuint32(rowoffset + weatherStationNamePositionOffset))
				if x.WeatherStationName, err = d.readstr(d.readuint32_row(row, d.weatherStationNamePositionOffset)); err != nil {
					return x, err
				}
			}

			if mode&mcc != 0 && d.mccEnabled {
				// x.MCC = readstr(readuint32(rowoffset + mccPositionOffset))
				if x.MCC, err = d.readstr(d.readuint32_row(row, d.mccPositionOffset)); err != nil {
					return x, err
				}
			}

			if mode&mnc != 0 && d.mncEnabled {
				// x.MNC = readstr(readuint32(rowoffset + mncPositionOffset))
				if x.MNC, err = d.readstr(d.readuint32_row(row, d.mncPositionOffset)); err != nil {
					return x, err
				}
			}

			if mode&mobileBrand != 0 && d.mobileBrandEnabled {
				// x.MobileBrand = readstr(readuint32(rowoffset + mobileBrandPositionOffset))
				if x.MobileBrand, err = d.readstr(d.readuint32_row(row, d.mobileBrandPositionOffset)); err != nil {
					return x, err
				}
			}

			if mode&elevation != 0 && d.elevationEnabled {
				// f, _ := strconv.ParseFloat(readstr(readuint32(rowoffset + elevationPositionOffset)), 32)
				res, err := d.readstr(d.readuint32_row(row, d.elevationPositionOffset))
				if err != nil {
					return x, err
				}

				f, _ := strconv.ParseFloat(res, 32)
				x.Elevation = float32(f)
			}

			if mode&usageType != 0 && d.usageTypeEnabled {
				// x.UsageType = readstr(readuint32(rowoffset + usageTypePositionOffset))
				if x.UsageType, err = d.readstr(d.readuint32_row(row, d.usageTypePositionOffset)); err != nil {
					return x, err
				}
			}

			return x, nil
		} else {
			if ipno.Cmp(ipfrom) < 0 {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return x, nil
}

func (d *DB) Close() {
	_ = d.f.Close()
}

func fatal(db *DB, err error) (*DB, error) {
	_ = db.f.Close()
	return nil, err
}

// ApiVersion returns the version of the component.
func ApiVersion() string {
	return apiVersion
}
