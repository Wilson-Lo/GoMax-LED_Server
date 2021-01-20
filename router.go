package main
import (
	//"time"
	//"fmt"
	"net/http"
	"github.com/gorilla/mux"
	//"net"
    "strings"
    "strconv"
	"encoding/hex"
)

func IPToByte4(ipnr string) [4]byte {
    bits := strings.Split(ipnr, ".")

	var ip4 [4]byte
    b0, _ := strconv.Atoi(bits[0])
    b1, _ := strconv.Atoi(bits[1])
    b2, _ := strconv.Atoi(bits[2])
    b3, _ := strconv.Atoi(bits[3])

	ip4[0] = byte(b0)
	ip4[1] = byte(b1)
	ip4[2] = byte(b2)
	ip4[3] = byte(b3)
    return ip4
}

func MacaddressToByte6(mac string) []byte {//not include ':'

	b_mac , _ := hex.DecodeString(mac )
    return b_mac
}


func NewRouter() *mux.Router {    
	r := mux.NewRouter()
	r.HandleFunc("/api/led/mode", api_GetLEDMode).Methods("GET")
	r.HandleFunc("/api/led/mode", api_SetLEDMode).Methods("POST")
	r.HandleFunc("/api/led/speed", api_GetSpeed).Methods("GET")
	r.HandleFunc("/api/led/speed", api_SetSpeed).Methods("POST")
	r.HandleFunc("/api/led/background_rgb", api_GetBackGroundRGB).Methods("GET")
	r.HandleFunc("/api/led/text_rgb", api_GetTextRGB).Methods("GET")
	r.HandleFunc("/api/led/text", api_GetText).Methods("GET")
	r.HandleFunc("/api/led/text", api_SetText).Methods("POST")
	r.HandleFunc("/api/led/vivid", api_GetVivid).Methods("GET")
	r.HandleFunc("/api/led/vivid", api_SetVivid).Methods("POST")
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./www/"))))
	return r
}

