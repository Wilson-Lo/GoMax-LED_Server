package main

import (
     "bufio"
     "os"
     "log"
     "strings"
)

var settingTXTPath = "../setting.txt"
var KEY_NOT_FIND = "notFind"
var KEY_MODE = "mode"
var KEY_BACKGROUND = "background"
var KEY_TEXT = "text"
var KEY_SPEED = "speed"
var KEY_VIVID = "vivid"

// LED Mode setting json structure
type ledMode struct {
    Led_mode int `json:"led_mode"`
}

// Speed setting json structure
type speed struct {
    Speed int `json:"speed"`
}

// Vivid setting json structure
type vivid struct {
    Vivid int `json:"vivid"`
}

// Text content setting json structure
type text struct {
    Content string `json:"content"`
}

// Text RGB setting json structure
type textRGB struct {
    R int `json:"r"`
    G int `json:"g"`
    B int `json:"b"`
}

// GIF
type upLoadGIF struct {
    Base64 string `json:"base64"`
}

// Hostname
type hostName struct {
    Hostname string `json:"hostname"`
}

/**
* Read .txt file line by keyword
*/
func readTXTByKeyWord(keyWord string) string {

	 file, err := os.Open(settingTXTPath)

     if err != nil {
        log.Fatal(err)
     }

     scanner := bufio.NewScanner(file)

     for scanner.Scan() {
        if(strings.Contains(scanner.Text(), keyWord)){
                return scanner.Text()
        }
     }
    // if err := scanner.Err(); err != nil {
     return KEY_NOT_FIND
    // }
}
