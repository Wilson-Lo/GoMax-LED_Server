package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"os"
	"net"
	"log"
	"io/ioutil"
	"strconv"
	"strings"
	"bufio"
	"crypto/md5"
	"os/exec"
    "encoding/base64"
    "time"
)

//use to fw update base64
var fw_file string

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
        if(strings.Contains(scanner.Text(), KEY_TEXT)){
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
    fmt.Print(" ")
    fmt.Print(textRGBObject.G)
    fmt.Print(" ")
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
    var setting_line = readTXTByKeyWord(KEY_BACKGROUND)
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
    fmt.Print(" ")
    fmt.Print(textRGBObject.G)
    fmt.Print(" ")
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
    var setting_line = readTXTByKeyWord(KEY_MODE)
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
    var setting_line = readTXTByKeyWord(KEY_SPEED)
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
    var setting_line = readTXTByKeyWord(KEY_VIVID)
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

        if(strings.Contains(scanner.Text(), KEY_VIVID)){
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

/**
*  Get All Settings
*/
func api_GetALLSetting(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)

    file, err := os.Open(settingTXTPath)
    if err != nil {
       log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    var isLastLine = false
    var mode string
    var speed string
    var text_content string
    var vivid = false
    var background_rgb [3] string
    var text_rgb [3] string

    for scanner.Scan() {

       if(strings.Contains(scanner.Text(), KEY_MODE)){
           data := strings.Split(scanner.Text(), " ")
           fmt.Println("mode = " + data[1])
           mode = data[1]
       }

       if(strings.Contains(scanner.Text(), KEY_BACKGROUND)){
           data := strings.Split(scanner.Text(), " ")
           fmt.Println("background rgb = " + data[1] + " " + data[2] + " " + data[3] + " " )
           background_rgb[0] = data[1]
           background_rgb[1] = data[2]
           background_rgb[2] = data[3]
       }

       if(strings.Contains(scanner.Text(), KEY_TEXT)){
            data := strings.Split(scanner.Text(), " ")
            fmt.Println("text rgb = " + data[1] + " " + data[2] + " " + data[3] + " " )
            text_rgb[0] = data[1]
            text_rgb[1] = data[2]
            text_rgb[2] = data[3]
       }

       if(strings.Contains(scanner.Text(), KEY_SPEED)){
            data := strings.Split(scanner.Text(), " ")
            fmt.Println("speed = " + data[1])
            speed = data[1]
       }

       if(isLastLine){
          text_content = scanner.Text()
          fmt.Println("text content = " + scanner.Text())
       }

       if(strings.Contains(scanner.Text(), KEY_VIVID)){
            data := strings.Split(scanner.Text(), " ")
            fmt.Println("vivid = " + data[1])
            isLastLine = true
             switch data[1] {
             case "0":
                  vivid = false
                  break;

             case "1":
                  vivid = true
                  break;
              }
       }
    }

    //get hostname
    hostname,err := os.Hostname()
    if err != nil {
       hostname = ""
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
        fmt.Println("ip = " + tmp_ip)
		}
    }

    fmt.Fprintf(w,"{\"led_mode\":"+ mode + ",\"ip\":\"" + tmp_ip + "\",\"hostname\":\"" + hostname + "\",\"speed\":" + speed + ",\"vivid\":" + strconv.FormatBool(vivid) + ", \"content\": \"" + text_content + "\","+"\"text_rgb\":{\"r\":" + text_rgb[0] + ",\"g\":" + text_rgb[1] + ",\"b\":" + text_rgb[2] + "},"+ "\"background_rgb\":{\"r\":" + background_rgb[0] + ",\"g\":" + background_rgb[1] + ",\"b\":" + background_rgb[2] + "}}");
	w.(http.Flusher).Flush()
}

/**
*  Upload GIF
*/
func api_UploadGIF(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
    fmt.Println("api_UploadGIF")
    jsonDecoder := json.NewDecoder(r.Body)
	var gifObject upLoadGIF
    err := jsonDecoder.Decode(&gifObject)
    if err != nil {
       fmt.Println("api_UploadGIF JSON error !")
       fmt.Fprintf(w,"{\"result\":\"failed\"}")
       w.(http.Flusher).Flush()
       panic(err)
       return
    }
    fmt.Println(gifObject.Base64)
    unbased, err := base64.StdEncoding.DecodeString(gifObject.Base64)
    if err != nil {
        fmt.Println("api_UploadGIF base64 error !")
        fmt.Fprintf(w,"{\"result\":\"failed\"}")
    	return
    }
    _ = ioutil.WriteFile("../led.gif", unbased, 0644)
   	fmt.Fprintf(w,"{\"result\":\"ok\"}")
   	w.(http.Flusher).Flush()
}

/**
*  Set Hostname
*  file location: /etc/hostname
*/
func api_SetHostName(w http.ResponseWriter, r *http.Request) {

    fmt.Println("api_SetHostName")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
    jsonDecoder := json.NewDecoder(r.Body)
	var hostnameObject hostName
    err := jsonDecoder.Decode(&hostnameObject)
    if err != nil {
       fmt.Println("api_SetHostName JSON error !")
       fmt.Fprintf(w,"{\"result\":\"failed\"}")
       w.(http.Flusher).Flush()
       panic(err)
       return
    }
    fmt.Println(hostnameObject.Hostname)
    cmd := exec.Command( "sudo", "sed", "-i", "1c " + hostnameObject.Hostname, "/etc/hostname")
	_,_ = cmd.Output()
	cmd2 := exec.Command( "sudo", "sed", "-i", "7c 127.0.0.1       " + hostnameObject.Hostname, "/etc/hosts")
    _,_ = cmd2.Output()
   	fmt.Fprintf(w,"{\"result\":\"ok\"}")

   	w.(http.Flusher).Flush()
    time.Sleep(1500 * time.Millisecond)
    go rebootShell()
}

/**
*  Get Hostname
*/
func api_GetHostName(w http.ResponseWriter, r *http.Request) {

    fmt.Println("api_GetHostName")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

    hostname,err := os.Hostname()
    if err == nil {
       fmt.Fprintf(w,"{\"result\":true,\"hostname\":\"" + hostname + "\"}")
       w.(http.Flusher).Flush()
    }else{
       fmt.Fprintf(w,"{\"result\":false,\"hostname\":\"\"}")
       w.(http.Flusher).Flush()
    }
}

/**
*  Get Config file base64 string
*/
func api_GetConfigFileBase64(w http.ResponseWriter, r *http.Request) {

    fmt.Println("api_GetConfigFileBase64")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

     txtByteArray, err := ioutil.ReadFile(settingTXTPath) // just pass the file name
     if err != nil {
        fmt.Println(err)
     }else{
        fmt.Println("get txt to byte array successful !")
     }

     if txtByteArray != nil {
        encodedBase64String := base64.StdEncoding.EncodeToString(txtByteArray)
        fmt.Fprintf(w,"{\"base64\":\"" + encodedBase64String + "\"}")
        w.(http.Flusher).Flush()
     }else{
        fmt.Fprintf(w,"{\"result\":\"failed\"}")
        w.(http.Flusher).Flush()
     }
}

/**
*  POST web & golang server file base64 string
*/
func api_fw_update(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")//delete this line when FW is publishing **************************
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w,"{\"result\":\"%v\"}", fw_update(r.FormValue("total"),r.FormValue("index"),r.FormValue("md5"),r.FormValue("base64")))
	w.(http.Flusher).Flush()
}

//fw update functions
func fw_update(total,index, md5Data , data64 string)(string){

	b64_ := strings.Replace(data64, " ", "+",-1)

	if index=="1"{//clear
	fw_file ="";
	}
	fw_file = fw_file + b64_

	m_data := []byte(b64_)
	md5_Data := md5.Sum(m_data)

	md5str1 := fmt.Sprintf("%X", md5_Data) //byte to hex
	fmt.Println("Ori:"+md5Data)
	fmt.Println("data:"+md5str1)
	if md5Data!=md5str1{
	return "failed"
	}

	if total == index{

	unbased, err := base64.StdEncoding.DecodeString(fw_file)
	if err != nil {
    	panic("Cannot decode b64")
		return "failed"
	}
	_ = ioutil.WriteFile("./tmp/fw.zip",unbased, 0644)
	fmt.Println("wait for save firmware...")
	time.Sleep(3000 * time.Millisecond)
	sync()
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("start update...")
	go runUpdateShell()
	time.Sleep(3000 * time.Millisecond)
	sync()
	}
	return "OK";
}

//linux command: sync
func sync(){
    cmd := exec.Command("sync")
	out,err := cmd.Output()
    if err != nil {
        fmt.Println(err)
    }
	fmt.Println(string(out))
}

//linux run update script
func runUpdateShell(){

	command := `sh update.sh`
    cmd := exec.Command("sh", "update.sh")
    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
        return
    }
    fmt.Println("Execute Shell:%s finished with output:\n%s", command, string(output))
}

//linux run reboot script
func rebootShell(){

	command := `sh reboot.sh`
    cmd := exec.Command("sh", "reboot.sh")
    output, err := cmd.Output()
    if err != nil {
        fmt.Printf("Execute reboot shell:%s failed with error:%s", command, err.Error())
        return
    }
    fmt.Println("Execute reboot shell:%s finished with output:\n%s", command, string(output))
}