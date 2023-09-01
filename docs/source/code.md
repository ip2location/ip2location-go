# IP2Location Go API

## IP2Location Class

```{py:function} OpenDB(binPath)
Load the IP2Location BIN database for lookup.

:param str binPath: (Required) The file path links to IP2Location BIN databases.
```

```{py:function} Get_all(ipAddress)
Retrieve geolocation information for an IP address.

:param str ipAddress: (Required) The IP address (IPv4 or IPv6).
:return: Returns the geolocation information in array. Refer below table for the fields avaliable in the array
:rtype: array

**RETURN FIELDS**

| Field Name       | Description                                                  |
| ---------------- | ------------------------------------------------------------ |
| Country_short    |     Two-character country code based on ISO 3166. |
| Country_long     |     Country name based on ISO 3166. |
| Region           |     Region or state name. |
| City             |     City name. |
| Isp              |     Internet Service Provider or company\'s name. |
| Latitude         |     City latitude. Defaults to capital city latitude if city is unknown. |
| Longitude        |     City longitude. Defaults to capital city longitude if city is unknown. |
| Domain           |     Internet domain name associated with IP address range. |
| ZipCode          |     ZIP code or Postal code. [172 countries supported](https://www.ip2location.com/zip-code-coverage). |
| Timezone         |     UTC time zone (with DST supported). |
| Netspeed         |     Internet connection type. |
| Iddcode         |     The IDD prefix to call the city from another country. |
| Areacode        |     A varying length number assigned to geographic areas for calls between cities. [223 countries supported](https://www.ip2location.com/area-code-coverage). |
| Weatherstationcode     |     The special code to identify the nearest weather observation station. |
| Weatherstationname     |     The name of the nearest weather observation station. |
| Mcc              |     Mobile Country Codes (MCC) as defined in ITU E.212 for use in identifying mobile stations in wireless telephone networks, particularly GSM and UMTS networks. |
| Mnc              |     Mobile Network Code (MNC) is used in combination with a Mobile Country Code(MCC) to uniquely identify a mobile phone operator or carrier. |
| Mobilebrand     |     Commercial brand associated with the mobile carrier. You may click [mobile carrier coverage](https://www.ip2location.com/mobile-carrier-coverage) to view the coverage report. |
| Elevation        |     Average height of city above sea level in meters (m). |
| Usagetype       |     Usage type classification of ISP or company. |
| Addresstype     |     IP address types as defined in Internet Protocol version 4 (IPv4) and Internet Protocol version 6 (IPv6). |
| Category         |     The domain category based on [IAB Tech Lab Content Taxonomy](https://www.ip2location.com/free/iab-categories). |
| District         |     District or county name. |
| Asn              |     Autonomous system number (ASN). BIN databases. |
| As          |     Autonomous system (AS) name. |
```

## IPTools Class

```{py:function} OpenTools ()
Initiate IPTools class.
```

```{py:function} IsIPv4(ipAddress)
Verify if a string is a valid IPv4 address.

:param str ipAddress: (Required) IP address.
:return: Return True if the IP address is a valid IPv4 address or False if it isn't a valid IPv4 address.
:rtype: boolean
```

```{py:function} IsIPv6(ipAddress)
Verify if a string is a valid IPv6 address

:param str ipAddress: (Required) IP address.
:return: Return True if the IP address is a valid IPv6 address or False if it isn't a valid IPv6 address.
:rtype: boolean
```

```{py:function} IPv4ToDecimal(ipAddress)
Translate IPv4 address from dotted-decimal address to decimal format.

:param str ipAddress: (Required) IPv4 address.
:return: Return the decimal format of the IPv4 address.
:rtype: int
```

```{py:function} DecimalToIPv4(ipNumber)
Translate IPv4 address from decimal number to dotted-decimal address.

:param str ip_number: (Required) Decimal format of the IPv4 address.
:return: Returns the dotted-decimal format of the IPv4 address.
:rtype: string
```

```{py:function} IPv6ToDecimal(ipAddress)
Translate IPv6 address from hexadecimal address to decimal format.

:param str ipAddress: (Required) IPv6 address.
:return: Return the decimal format of the IPv6 address.
:rtype: int
```

```{py:function} DecimalToIPv6(ipNumber)
Translate IPv6 address from decimal number into hexadecimal address.

:param str ip_number: (Required) Decimal format of the IPv6 address.
:return: Returns the hexadecimal format of the IPv6 address.
:rtype: string
```

```{py:function} IPv4ToCIDR(ip_from, ip_to)
Convert IPv4 range into a list of IPv4 CIDR notation.

:param str ip_from: (Required) The starting IPv4 address in the range.
:param str ip_to: (Required) The ending IPv4 address in the range.
:return: Returns the list of IPv4 CIDR notation.
:rtype: array
```

```{py:function} CIDRToIPv4(cidr)
Convert IPv4 CIDR notation into a list of IPv4 addresses.

:param str cidr: (Required) IPv4 CIDR notation.
:return: Returns an list of IPv4 addresses.
:rtype: array
```

```{py:function} IPv6ToCIDR(ip_from, ip_to)
Convert IPv6 range into a list of IPv6 CIDR notation.

:param str ip_from: (Required) The starting IPv6 address in the range.
:param str ip_to: (Required) The ending IPv6 address in the range.
:return: Returns the list of IPv6 CIDR notation.
:rtype: array
```

```{py:function} CIDRToIPv6(cidr)
Convert IPv6 CIDR notation into a list of IPv6 addresses.

:param str cidr: (Required) IPv6 CIDR notation.
:return: Returns an list of IPv6 addresses.
:rtype: array
```


```{py:function} CompressIPv6(ipAddress)
Compress a IPv6 to shorten the length.

:param str ipAddress: (Required) IPv6 address.
:return: Returns the compressed version of IPv6 address.
:rtype: str
```

```{py:function} ExpandIPv6(ipAddress)
Expand a shorten IPv6 to full length.

:param str ipAddress: (Required) IPv6 address.
:return: Returns the extended version of IPv6 address.
:rtype: str
```

## Country Class

```{py:function} OpenCountryInfo(CSVFilePath)
Load the IP2Location Country Information CSV file. This database is free for download at <https://www.ip2location.com/free/country-information>.

:param str CSVFilePath: (Required) The file path links to IP2Location Country Information CSV file.
```

```{py:function} GetCountryInfo(CountryCode)
Provide a ISO 3166 country code to get the country information in array. Will return a full list of countries information if country code not provided. 

:param str CountryCode: (Required) The ISO 3166 country code of a country.
:return: Returns the country information in array. Refer below table for the fields avaliable in the array.
:rtype: array

**RETURN FIELDS**

| Field Name       | Description                                                  |
| ---------------- | ------------------------------------------------------------ |
| Country_code     | Two-character country code based on ISO 3166.                |
| Country_alpha3_code | Three-character country code based on ISO 3166.           |
| Country_numeric_code | Three-character country code based on ISO 3166.          |
| Capital          | Capital of the country.                                      |
| Country_demonym  | Demonym of the country.                                      |
| Total_area       | Total area in km{sup}`2`.                                    |
| Population       | Population of year 2014.                                     |
| Idd_code         | The IDD prefix to call the city from another country.        |
| Currency_code    | Currency code based on ISO 4217.                             |
| Currency_name    | Currency name.                                               |
| Currency_symbol  | Currency symbol.                                             |
| Lang_code        | Language code based on ISO 639.                              |
| Lang_name        | Language name.                                               |
| Cctld            | Country-Code Top-Level Domain.                               |
```

## Region Class

```{py:function} OpenRegionInfo(CSVFilePath)
Load the IP2Location ISO 3166-2 Subdivision Code CSV file. This database is free for download at <https://www.ip2location.com/free/iso3166-2>

:param str CSVFilePath: (Required) The file path links to IP2Location ISO 3166-2 Subdivision Code CSV file.
```

```{py:function} GetRegionCode(countryCode, regionName)
Provide a ISO 3166 country code and the region name to get ISO 3166-2 subdivision code for the region.

:param str countryCode: (Required) Two-character country code based on ISO 3166.
:param str regionName: (Required) Region or state name.
:return: Returns the ISO 3166-2 subdivision code of the region.
:rtype: str
```