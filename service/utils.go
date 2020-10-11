package service

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"
)

const (
	NORESULTS = "No results for query, %s"
)

// IP2int converts from IP to integer
func IP2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}

// Int2IP converts from integer to ip
func Int2IP(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

//NoResultError custom error for no results
func NoResultError(message string) error {
	return errors.New(fmt.Sprintf(NORESULTS, message))
}

//LogError self explanatory
func LogError(err error) {
	log.Panic(NORESULTS, err)
}
