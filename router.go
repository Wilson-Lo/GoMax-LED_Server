package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {    
	r := mux.NewRouter()
	r.HandleFunc("/api/led/all", api_GetALLSetting).Methods("GET")
	r.HandleFunc("/api/led/mode", api_GetLEDMode).Methods("GET")
	r.HandleFunc("/api/led/mode", api_SetLEDMode).Methods("POST")
	r.HandleFunc("/api/led/speed", api_GetSpeed).Methods("GET")
	r.HandleFunc("/api/led/speed", api_SetSpeed).Methods("POST")
	r.HandleFunc("/api/led/background_rgb", api_GetBackGroundRGB).Methods("GET")
	r.HandleFunc("/api/led/background_rgb", api_SetBackGroundRGB).Methods("POST")
	r.HandleFunc("/api/led/text_rgb", api_GetTextRGB).Methods("GET")
	r.HandleFunc("/api/led/text_rgb", api_SetTextRGB).Methods("POST")
	r.HandleFunc("/api/led/text", api_GetText).Methods("GET")
	r.HandleFunc("/api/led/text", api_SetText).Methods("POST")
	r.HandleFunc("/api/led/vivid", api_GetVivid).Methods("GET")
	r.HandleFunc("/api/led/vivid", api_SetVivid).Methods("POST")
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./www/"))))
	return r
}

