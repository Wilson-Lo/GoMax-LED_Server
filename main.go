package main
import (
    "fmt"
    "log"
	//"time"
   // "net"
    "net/http"
	//"time"
	//"github.com/goburrow/serial"
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
	get_pi4_ipconfig()
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

	log.Fatal(http.ListenAndServeTLS(":10443","server.crt", "server.key",router))

	defer func() {

		}()
}


