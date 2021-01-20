package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"os"
	"log"
	"io/ioutil"
	"strconv"
	"strings"
	"bufio"
)

/**
*  Get Text RGB
*/
func api_GetTextRGB(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	//vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

    file, err := os.Open("../setting.txt")
    if err != nil {
       log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)
    var first_line string
    for scanner.Scan() {
        if(strings.Contains(scanner.Text(), "text")){
           first_line = scanner.Text()
           break
        }
    }

    if err := scanner.Err(); err != nil {
            log.Fatal(err)
    }

    fmt.Println(first_line)
    data := strings.Split(first_line, " ")
  //  if data[1] != nil {
        w.WriteHeader(http.StatusOK)
    	fmt.Fprintf(w,"{\"type\":\"text\", \"r\":" + data[1] + ",\"g\":" + data[2] + ",\"b\":" + data[3] + "}")
    //}
	w.(http.Flusher).Flush()
}

/**
*  Set Text RGB
*/
func api_SetTextRGB(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonDecoder := json.NewDecoder(r.Body)
	var textRGBObject textRGB
    err := jsonDecoder.Decode(&textRGBObject)
    if err != nil {
       fmt.Println("Set Text RGB failed !")
       fmt.Fprintf(w,"{\"result\":\"failed\"}")
       w.(http.Flusher).Flush()
       panic(err)
       return
    }

    fmt.Print("New Text RGB ")
    fmt.Print(textRGBObject.R)
    fmt.Print("")
    fmt.Print(textRGBObject.G)
    fmt.Print("")
    fmt.Println(textRGBObject.B)

    input, err := ioutil.ReadFile(settingTXTPath)
    if err != nil {
       log.Fatalln(err)
    }

    lines := strings.Split(string(input), "\n")
    lines[2] = "text " + strconv.Itoa(textRGBObject.R) + " " + strconv.Itoa(textRGBObject.G) + " " + strconv.Itoa(textRGBObject.B)
    output := strings.Join(lines, "\n")
    err = ioutil.WriteFile(settingTXTPath, []byte(output), 0644)
    if err != nil {
        log.Fatalln(err)
        fmt.Fprintf(w,"{\"result\":\"failed\"}")
        w.(http.Flusher).Flush()
        return
    }
	fmt.Fprintf(w,"{\"result\":\"ok\"}")
	w.(http.Flusher).Flush()
}

/**
*  Get Background RGB
*/
func api_GetBackGroundRGB(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
    var setting_line = readTXTByKeyWord("background")
    fmt.Println(setting_line)

    result := setting_line == KEY_NOT_FIND

    if(!result){
        data := strings.Split(setting_line, " ")
        fmt.Fprintf(w,"{\"type\":\"background\", \"r\":" + data[1] + ",\"g\":" + data[2] + ",\"b\":" + data[3] + "}")
    }else{
        fmt.Fprintf(w,"{\"type\":\"background\", \"r\":-1,\"g\":-1,\"b\":-1}")
    }
    w.(http.Flusher).Flush()
}

/**
*  Set Background RGB
*/
func api_SetBackGroundRGB(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonDecoder := json.NewDecoder(r.Body)
	var textRGBObject textRGB
    err := jsonDecoder.Decode(&textRGBObject)
    if err != nil {
       fmt.Println("Set Background RGB failed !")
       fmt.Fprintf(w,"{\"result\":\"failed\"}")
       w.(http.Flusher).Flush()
       panic(err)
       return
    }

    fmt.Print("New Text RGB ")
    fmt.Print(textRGBObject.R)
    fmt.Print("")
    fmt.Print(textRGBObject.G)
    fmt.Print("")
    fmt.Println(textRGBObject.B)

    input, err := ioutil.ReadFile(settingTXTPath)
    if err != nil {
       log.Fatalln(err)
    }

    lines := strings.Split(string(input), "\n")
    lines[1] = "background " + strconv.Itoa(textRGBObject.R) + " " + strconv.Itoa(textRGBObject.G) + " " + strconv.Itoa(textRGBObject.B)
    output := strings.Join(lines, "\n")
    err = ioutil.WriteFile(settingTXTPath, []byte(output), 0644)
    if err != nil {
        log.Fatalln(err)
        fmt.Fprintf(w,"{\"result\":\"failed\"}")
        w.(http.Flusher).Flush()
        return
    }
	fmt.Fprintf(w,"{\"result\":\"ok\"}")
	w.(http.Flusher).Flush()
}

/**
*  Get LED Mode
*/
func api_GetLEDMode(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	//vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    var setting_line = readTXTByKeyWord("mode")
    fmt.Println(setting_line)
    result := setting_line == KEY_NOT_FIND

    if(!result){
       data := strings.Split(setting_line, " ")
       fmt.Fprintf(w,"{\"led_mode\":" + data[1] + "}")
    }else{
       fmt.Fprintf(w,"{\"led_mode\":-1}")
    }
	w.(http.Flusher).Flush()
}

/**
*  Set LED Mode
*/
func api_SetLEDMode(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonDecoder := json.NewDecoder(r.Body)
	var ledModeObject ledMode
    err := jsonDecoder.Decode(&ledModeObject)
    if err != nil {
       fmt.Println("Set LED Mode failed !")
       fmt.Fprintf(w,"{\"result\":\"failed\"}")
       w.(http.Flusher).Flush()
       panic(err)
       return
    }

    fmt.Print("New LED Mode ")
    fmt.Println(ledModeObject.Led_mode)

    input, err := ioutil.ReadFile(settingTXTPath)
    if err != nil {
       log.Fatalln(err)
    }

    lines := strings.Split(string(input), "\n")
    lines[0] = "mode " + strconv.Itoa(ledModeObject.Led_mode)
    output := strings.Join(lines, "\n")
    err = ioutil.WriteFile(settingTXTPath, []byte(output), 0644)
    if err != nil {
        log.Fatalln(err)
        fmt.Fprintf(w,"{\"result\":\"failed\"}")
        w.(http.Flusher).Flush()
        return
    }
	fmt.Fprintf(w,"{\"result\":\"ok\"}")
	w.(http.Flusher).Flush()
}

/**
*  Get Speed ( 0: fast ~ 5:slow )
*/
func api_GetSpeed(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    var setting_line = readTXTByKeyWord("speed")
    fmt.Println(setting_line)
    result := setting_line == KEY_NOT_FIND

    if(!result){
       data := strings.Split(setting_line, " ")
       fmt.Fprintf(w,"{\"speed\":" + data[1] + "}")
    }else{
       fmt.Fprintf(w,"{\"speed\":-1}")
    }
	w.(http.Flusher).Flush()
}

/**
*  Set Speed
*/
func api_SetSpeed(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonDecoder := json.NewDecoder(r.Body)
	var speedObject speed
    err := jsonDecoder.Decode(&speedObject)
    if err != nil {
       fmt.Println("Set Speed failed !")
       fmt.Fprintf(w,"{\"result\":\"failed\"}")
       w.(http.Flusher).Flush()
       panic(err)
       return
    }

    fmt.Print("New Speed ")
    fmt.Println(speedObject.Speed)

    input, err := ioutil.ReadFile(settingTXTPath)
    if err != nil {
       log.Fatalln(err)
    }

    lines := strings.Split(string(input), "\n")
    lines[3] = "speed " + strconv.Itoa(speedObject.Speed)
    output := strings.Join(lines, "\n")
    err = ioutil.WriteFile(settingTXTPath, []byte(output), 0644)
    if err != nil {
        log.Fatalln(err)
        fmt.Fprintf(w,"{\"result\":\"failed\"}")
        w.(http.Flusher).Flush()
        return
    }
	fmt.Fprintf(w,"{\"result\":\"ok\"}")
	w.(http.Flusher).Flush()
}

/**
*  Get Vivid ( 0: Off, 1: On ,Color on or off)
*/
func api_GetVivid(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
    var setting_line = readTXTByKeyWord("vivid")
    fmt.Println(setting_line)
    result := setting_line == KEY_NOT_FIND

    if(!result){
       data := strings.Split(setting_line, " ")
       switch data[1] {
         case "0":
              fmt.Fprintf(w,"{\"vivid\":false}")
              break;

         case "1":
              fmt.Fprintf(w,"{\"vivid\":true}")
              break;
        }
    }else{
       fmt.Fprintf(w,"{\"vivid\":-1}")
    }
	w.(http.Flusher).Flush()
}

/**
*  Set Vivid ( 0: Off, 1: On ,Color on or off)
*/
func api_SetVivid(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonDecoder := json.NewDecoder(r.Body)
	var vividObject vivid
    err := jsonDecoder.Decode(&vividObject)
    if err != nil {
       fmt.Println("Set Vivid failed !")
       fmt.Fprintf(w,"{\"result\":\"failed\"}")
       w.(http.Flusher).Flush()
       panic(err)
       return
    }

    fmt.Print("New Vivid ")
    fmt.Println(vividObject.Vivid)

    input, err := ioutil.ReadFile(settingTXTPath)
    if err != nil {
       log.Fatalln(err)
    }

    lines := strings.Split(string(input), "\n")
    lines[4] = "vivid " + strconv.Itoa(vividObject.Vivid)
    output := strings.Join(lines, "\n")
    err = ioutil.WriteFile(settingTXTPath, []byte(output), 0644)
    if err != nil {
        log.Fatalln(err)
        fmt.Fprintf(w,"{\"result\":\"failed\"}")
        w.(http.Flusher).Flush()
        return
    }
	fmt.Fprintf(w,"{\"result\":\"ok\"}")
	w.(http.Flusher).Flush()
}

/**
*  Get Text Content
*/
func api_GetText(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	//vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

    file, err := os.Open(settingTXTPath)
    if err != nil {
       log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)
    var last_line string
    var isLastLine = false
    for scanner.Scan() {

        if(isLastLine){
            last_line = scanner.Text()
        }

        if(strings.Contains(scanner.Text(), "vivid")){
           isLastLine = true
        }
    }

    if err := scanner.Err(); err != nil {
       log.Fatal(err)
    }

    fmt.Println("text = " + last_line)

    fmt.Fprintf(w,"{\"content\":\"" + last_line + "\"}")
	w.(http.Flusher).Flush()
}

/**
*  Set Text Content
*/
func api_SetText(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonDecoder := json.NewDecoder(r.Body)
	var textObject text
    err := jsonDecoder.Decode(&textObject)
    if err != nil {
       fmt.Println("Set Text content failed !")
       fmt.Fprintf(w,"{\"result\":\"failed\"}")
       w.(http.Flusher).Flush()
       panic(err)
       return
    }

    fmt.Print("New Text ")
    fmt.Println(textObject.Content)

    input, err := ioutil.ReadFile(settingTXTPath)
    if err != nil {
       log.Fatalln(err)
    }

    lines := strings.Split(string(input), "\n")
    lines[5] = textObject.Content
    output := strings.Join(lines, "\n")
    err = ioutil.WriteFile(settingTXTPath, []byte(output), 0644)
    if err != nil {
        log.Fatalln(err)
        fmt.Fprintf(w,"{\"result\":\"failed\"}")
        w.(http.Flusher).Flush()
        return
    }
	fmt.Fprintf(w,"{\"result\":\"ok\"}")
	w.(http.Flusher).Flush()
}