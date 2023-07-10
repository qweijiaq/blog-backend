package utils

import (
	"log"
	"net"

	"github.com/sirupsen/logrus"
)

func GetIPList() (IPList []string) {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, i2 := range interfaces {
		address, err := i2.Addrs()
		if err != nil {
			logrus.Error(err)
			continue
		}
		for _, addr := range address {
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			ip4 := ipNet.IP.To4()
			if ip4 == nil {
				continue
			}
			IPList = append(IPList, ip4.String())
		}
	}
	return
}
