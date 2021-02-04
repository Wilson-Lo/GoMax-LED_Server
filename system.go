package main
import (
	"fmt"
	"net"
    "strings"
)

/**
* Get PI MAC Address
*/
func getMacAddrs() (macAddrs string) {
    netInterfaces, err := net.Interfaces()
    if err != nil {
        fmt.Printf("fail to get net interfaces: %v", err)
        return ""
    }

    for _, netInterface := range netInterfaces {
        macAddr := netInterface.HardwareAddr.String()
        if len(macAddr) == 0 {
            continue
        }

        macAddrs = macAddr
		macAddrs = strings.Replace(string(macAddrs[:]), ":", "",-1)
		macAddrs = strings.ToUpper(macAddrs)
		return macAddrs

    }
    return macAddrs
}