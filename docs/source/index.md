[![Go Report Card](https://goreportcard.com/badge/github.com/ip2location/ip2location-go/v9)](https://goreportcard.com/report/github.com/ip2location/ip2location-go/v9)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/ip2location/ip2location-go/v9)](https://pkg.go.dev/github.com/ip2location/ip2location-go/v9)

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

## Table of contents
 ```{eval-rst}
 .. toctree::

   self
   quickstart
   code
 ```