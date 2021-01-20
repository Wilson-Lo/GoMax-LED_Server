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
	//LoadPassword()
	//LoadSystemConfig()
	get_pi4_ipconfig()
	//loadEmail()
	//loadEvent()
	//ast_initial_preset()
	router := NewRouter()
	//device_info_load("clear")
	//history_list_load()
	//go ast_node_list()
	go func(){
	    fmt.Println(" go func 1")
		log.Fatal(http.ListenAndServe(":8080",router))
		fmt.Println(" go func 2")
		/*s := &http.Server{
		Addr:           ":80",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		}
		log.Fatal(s.ListenAndServe())*/

	}()
	fmt.Println(" go func 3")
	log.Fatal(http.ListenAndServeTLS(":10443","server.crt", "server.key",router))
	fmt.Println(" go func 4")
	defer func() {
		fmt.Println(" go func 5")
		//device_info_save()
		//history_list_save()
		}()
}


