![si-ethereum-isabelschoeps-png](https://github.com/ip2location/ip2location-go/assets/127110010/78fe8ecb-86d0-4499-b752-1c15d9f5714d)

# IP2Location Go Package

This Go package provides a fast lookup of country, region, city, latitude, longitude, ZIP code, time zone, ISP, domain name, connection type, IDD code, area code, weather station code, station name, mcc, mnc, mobile brand, elevation, usage type, address type, IAB category, district, autonomous system number (ASN) and autonomous system (AS) from IP address by using IP2Location database. This package uses a file based database available at IP2Location.com. This database simply contains IP blocks as keys, and other information such as country, region, city, latitude, longitude, ZIP code, time zone, ISP, domain name, connection type, IDD code, area code, weather station code, station name, mcc, mnc, mobile brand, elevation, usage type, address type, IAB category, district, autonomous system number (ASN) and autonomous system (AS) as values. It supports both IP address in IPv4 and IPv6.

This package can be used in many types of projects such as:

 - select the geographically closest mirror
 - analyze your web server logs to determine the countries of your visitors
 - credit card fraud detection
 - software export controls
 - display native language and currency 
 - prevent password sharing and abuse of service 
 - geotargeting in advertisement

The database will be updated in monthly basis for the greater accuracy. Free LITE databases are available at https://lite.ip2location.com/ upon registration.

The paid databases are available at https://www.ip2location.com under Premium subscription package.

As an alternative, this package can also call the IP2Location Web Service. This requires an API key. If you don't have an existing API key, you can subscribe for one at the below:
https://www.ip2location.com/web-service/ip2location

Developer Documentation
=====================

To learn more about installation, usage, and code examples, please visit the developer documentation at [https://ip2location-go.readthedocs.io/en/latest/](https://ip2location-go.readthedocs.io/en/latest/).
