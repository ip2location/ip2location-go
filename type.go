package ip2location

// IP2Location интерфейс для работы с Query
type IP2Location interface {
	Query(int, uint32) (Record, error)
	Close() error
}

// CountryShort код страны
const CountryShort uint32 = 0x00001

// CountryLong страна
const CountryLong uint32 = 0x00002

// Region регион
const Region uint32 = 0x00004

// City город
const City uint32 = 0x00008

// ISP интернет провайдер
const ISP uint32 = 0x00010

// Latitude широта
const Latitude uint32 = 0x00020

// Longitude долгота
const Longitude uint32 = 0x00040

// Domain домэн
const Domain uint32 = 0x00080

// ZipCode индекс
const ZipCode uint32 = 0x00100

// TimeZone часовой пояс
const TimeZone uint32 = 0x00200

// NetSpeed скорость интернета
const NetSpeed uint32 = 0x00400

// IddCode код страны
const IddCode uint32 = 0x00800

// AreaCode код города(области)
const AreaCode uint32 = 0x01000

// WeatherStationCode код метеостанции
const WeatherStationCode uint32 = 0x02000

// WeatherStationName имя метеостанции
const WeatherStationName uint32 = 0x04000

// Mcc мобильный код страны
const Mcc uint32 = 0x08000

// Mnc код мобильной сети
const Mnc uint32 = 0x10000

// MobileBrand мобильный брэнд (оператор)
const MobileBrand uint32 = 0x20000

// Elevation высота
const Elevation uint32 = 0x40000

// UsageType тип использования
const UsageType uint32 = 0x80000

// All все поля
const All uint32 = CountryShort | CountryLong | Region | City | ISP | Latitude | Longitude | Domain | ZipCode | TimeZone | NetSpeed | IddCode | AreaCode | WeatherStationCode | WeatherStationName | Mcc | Mnc | MobileBrand | Elevation | UsageType

// Record содержит данные о местоположении
type Record struct {
	CountryShort       string
	CountryLong        string
	Region             string
	City               string
	ISP                string
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
