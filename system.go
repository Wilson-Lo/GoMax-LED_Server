package main
import (
	"fmt"
	"net"
    "strings"
    "os/exec"
	"log"
	"os"
	"strconv"
	"net/mail"
	"encoding/binary"
	"bufio"
)

const (
    file  = "/proc/net/route"
    line  = 1    // line containing the gateway addr. (first line: 0)
    sep   = "\t" // field separator
    field = 2    // field containing hex gateway address (first field: 0)
)

type systemConfig struct {	
	UI string `json:"ui_type"`
	Timezone string `json:"tz"`
	IP_mode string `json:"ip_mode"`
	TIME string  `json:"time"`
	IP string `json:"ip"`
	MASK string `json:"mask"`
	GATEWAY string `json:"gw"`
	
}

type Event_log struct {	
	Type string `json:"type"`
	Time string `json:"time"`
	Event_body string `json:"event"`
}

type emialGroup struct {	
	INDEX string `json:"index"`
	NAME string `json:"name"`
	EMAIL string `json:"email"`
}

type emialList struct {	
	Mail []emialGroup `json:"mail"`
}

type event_node struct {	
	EMAIL string `json:"email_index"`
	MAC string  `json:"mac"`
	TYPE string  `json:"type"`
}

type LoginAuth struct {
    username, password string
}

var event_connect_map =make(map[string]*event_node)
var event_videolost_map =make(map[string]*event_node)

var event_num , event_index int = 0 , 0
var event_array[520] Event_log
var pi4_mac string
var pi4_password string ="admin"
var eventEmail_map[8] emialGroup
var fw_file string
var sys_config systemConfig

func encodeRFC2047(String string) string{
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}


func get_pi4_ipconfig(){

	ipCnt := 0
	/*mgmtInterface, err := net.InterfaceByName("eth0")
    if err != nil {
        fmt.Println("Unable to find interface")
        os.Exit(-1)
    }*/
	addrs, err := net.InterfaceAddrs()//mgmtInterface.Addrs()
	if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	
	var tmp_ip,tmp_mask string
		
	for _, addr := range addrs {
	
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
        var ip net.IP
        var mask net.IPMask
        switch v := addr.(type) {
        case *net.IPNet:
            ip = v.IP
            mask = v.Mask
        case *net.IPAddr:
            ip = v.IP
            mask = ip.DefaultMask()
        }
        if ip == nil {
            continue
        }
        ip = ip.To4()
        if ip == nil {
            continue
        }	
		cleanMask := fmt.Sprintf("%d.%d.%d.%d", mask[0], mask[1], mask[2], mask[3])
        //fmt.Println(ip, cleanMask,ipCnt)
		ipCnt = ipCnt+1
		
		if ipCnt>2 {
		continue
		}
		tmp_ip = ip.String()
		tmp_mask = cleanMask
		
		}
    }

	if sys_config.IP_mode == "STATIC"{
	
		 if tmp_ip!=sys_config.IP{
		 
		 fmt.Println("static IP not found",tmp_ip,sys_config.IP)
		 cmd := exec.Command("sudo","ip","adderss","add",sys_config.IP+"/24","dev eth0:0")
		 out,err := cmd.Output()
		 if err != nil {
			fmt.Println(err)
		 }
		 fmt.Println(string(out))
		 }
	
	
	}else{
	
		sys_config.IP = tmp_ip
		sys_config.MASK = tmp_mask
	
	}


	//
	file, err := os.Open(file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {

        // jump to line containing the agteway address
        for i := 0; i < line; i++ {
            scanner.Scan()
        }

        // get field containing gateway address
        tokens := strings.Split(scanner.Text(), sep)
        gatewayHex := "0x" + tokens[field]

        // cast hex address to uint32
        d, _ := strconv.ParseInt(gatewayHex, 0, 64)
        d32 := uint32(d)

        // make net.IP address from uint32
        ipd32 := make(net.IP, 4)
        binary.LittleEndian.PutUint32(ipd32, d32)
        //fmt.Printf("%T --> %[1]v\n", ipd32)

        // format net.IP to dotted ipV4 string
        gw := net.IP(ipd32).String()
		sys_config.GATEWAY = gw
        //fmt.Printf("%T --> %[1]v\n", ip)

        // exit scanner
        break
    }
}

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
