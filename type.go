package ip2location

type Query interface {
	Query(int, uint32) (Record, error)
	Close() error
}

const CountryShort uint32 = 0x00001
const CountryLong uint32 = 0x00002
const Region uint32 = 0x00004
const City uint32 = 0x00008
const Isp uint32 = 0x00010
const Latitude uint32 = 0x00020
const Longitude uint32 = 0x00040
const Domain uint32 = 0x00080
const ZipCode uint32 = 0x00100
const TimeZone uint32 = 0x00200
const NetSpeed uint32 = 0x00400
const IddCode uint32 = 0x00800
const AreaCode uint32 = 0x01000
const WeatherStationCode uint32 = 0x02000
const WeatherStationName uint32 = 0x04000
const Mcc uint32 = 0x08000
const Mnc uint32 = 0x10000
const MobileBrand uint32 = 0x20000
const Elevation uint32 = 0x40000
const UsageType uint32 = 0x80000

const All uint32 = CountryShort | CountryLong | Region | City | Isp | Latitude | Longitude | Domain | ZipCode | TimeZone | NetSpeed | IddCode | AreaCode | WeatherStationCode | WeatherStationName | Mcc | Mnc | MobileBrand | Elevation | UsageType

type Record struct {
	CountryShort       string
	CountryLong        string
	Region             string
	City               string
	Isp                string
	Latitude           float32
	Longitude          float32
	Domain             string
	ZipCode            string
	TimeZone           string
	NetSpeed           string
	IddCode            string
	AreaCode           string
	WeatherStationCode string
	WeatherStationName string
	Mcc                string
	Mnc                string
	MobileBrand        string
	Elevation          float32
	UsageType          string
}
