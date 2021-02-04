package main
import (
    "fmt"
    "log"
	"strings"
    "net"
    "net/http"
    "strconv"
    "encoding/hex"
    //"bufio"
   // "os"
)

var aesKey = []byte("qzy159pkn333rty2")

const (
    file  = "/proc/net/route"
)

func main() {

/*
	port, err := serial.Open(
		&serial.Config{
			Address:  "/dev/ttyAMA0",
			BaudRate: 115200,
			DataBits: 8,
			StopBits: 1,
			Parity:   "N",
			Timeout: 1 * time.Second,
	})

	if err != nil {
		log.Fatal("Comport open fail")
	}
	defer port.Close() */


	fmt.Println(" LED Server v1.0.0")
	fmt.Println(" PI MAC address: " + getMacAddrs())
	router := NewRouter()

	go func(){
		log.Fatal(http.ListenAndServe(":8080",router))
		/*s := &http.Server{
		Addr:           ":80",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		}
		log.Fatal(s.ListenAndServe())*/

	}()

    startUDPServer()

	log.Fatal(http.ListenAndServeTLS(":10443","server.crt", "server.key",router))

	defer func() {

		}()
}

func startUDPServer(){

    fmt.Println(" Start UDP Server")
    src := "255.255.255.255:5002"
	listener, err := net.ListenPacket("udp", src)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer listener.Close()

	fmt.Printf("UDP server start and listening on %s.\n", src)

	for {
		buf := make([]byte, 1024)
		n, addr, err := listener.ReadFrom(buf)
		if err != nil {
			continue
		}
		go serve(listener, addr, buf[:n])
	}
}

func serve(listener net.PacketConn, addr net.Addr, buf []byte) {

     netAddrArray := strings.Split(addr.String(), ":")
     //receiveCmd := string(buf)
     // fmt.Printf("%s\t: %s\n", netAddrArray[0], receiveCmd)

     if((len(buf)%16) == 0){
      fmt.Printf("phone ip = %s  receive size = %d\n", netAddrArray[0],len(buf))
          decodeData :=AesDecrypt(buf, aesKey)
          fmt.Printf("after aes decode = %s\n", decodeData)

          if(strings.Contains(decodeData, "ETH_REQ")){
             fmt.Printf("GoMax Device request")
          }else{
             fmt.Printf("Others Device request")
          }

         //get IP address
        addrs, err := net.InterfaceAddrs()

         if err != nil {
           fmt.Println(err)
         }

         var tmp_ip string
         ipCnt := 0

         for _, addr := range addrs {

         	if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
             var ip net.IP
             switch v := addr.(type) {
                 case *net.IPNet:
                     ip = v.IP

                 case *net.IPAddr:
                     ip = v.IP

                 }
                 if ip == nil {
                     continue
                 }
                 ip = ip.To4()
                 if ip == nil {
                     continue
                 }

         		ipCnt = ipCnt+1

         		if ipCnt>2 {
         		continue
         		}
         		tmp_ip = ip.String()
         		fmt.Println("")
                 fmt.Println("ip = " + tmp_ip)
         		}
         }
        var feedBackArray [38]byte

        //device name (LED)
        feedBackArray[5] = 0x4c
        feedBackArray[6] = 0x45
        feedBackArray[7] = 0x44

        //IP
        ipArray := strings.Split(tmp_ip, ".")
        var ipIndex = 27
        for counter := 0; counter < len(ipArray); counter++ {
            if(ipIndex > 30){
               break
            }
            ip, err := strconv.Atoi(ipArray[counter])
            if err != nil {
            }
            feedBackArray[ipIndex] = byte(ip)
            ipIndex++
        }

        //mac address
        data, err := hex.DecodeString(getMacAddrs())
        if(err != nil){
           fmt.Println("mac address error : " , err)
        }

        var macIndex = 21
        for counter := 0; counter < len(data); counter++ {
            if(macIndex > 26){
               break
            }
            feedBackArray[macIndex] = data[counter]
            macIndex++
        }

        //send UDP feedback
        encodeData := AesEncrypt(feedBackArray[:], aesKey)
        sendUDP("255.255.255.255:65088", encodeData)

     }else{

       fmt.Println("non aes data")
     }

}
//(string, error)
func sendUDP(addr string, msg []byte) {
    fmt.Println("sendUDP data size: %d" , len(msg))
	conn, _ := net.Dial("udp", addr)

    _, err := conn.Write(msg)
	if err != nil{
	   fmt.Println("send UDP err ", err)
	}

	// listen for reply
/*	bs := make([]byte, 1024)
	conn.SetDeadline(time.Now().Add(3 * time.Second))
	len, err := conn.Read(bs)
	if err != nil {
		return "", err
	} else {
		return string(bs[:len]), err
	} */
}