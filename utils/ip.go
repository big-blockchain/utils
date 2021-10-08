package utils

import (
	"github.com/mohong122/ip2region/binding/golang/ip2region"
	"strings"
)

var ipRegion *ip2region.Ip2Region

func IpToCity(ip string) string {

	if ipRegion == nil {
		ipRegion, _ = ip2region.New("./ip2region.db")
	}

	ipfo, _ := ipRegion.BinarySearch(ip)

	if ipfo.City != "" {
		if len(ipfo.City) > 2 {
			if strings.Contains(ipfo.City, "市") {
				return strings.Replace(ipfo.City, "市", "", 1)
			}
			if strings.Contains(ipfo.City, "州") {
				return strings.Replace(ipfo.City, "州", "", 1)
			}
		}

	}
	return ipfo.City
}

func IpToCountry(ip string) string {

	if ipRegion == nil {
		ipRegion, _ = ip2region.New("./ip2region.db")
	}

	ipfo, _ := ipRegion.BinarySearch(ip)

	return ipfo.Country
}
