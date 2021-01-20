package main

import (
     "bufio"
     "os"
     "log"
     "strings"
)

var settingTXTPath = "../setting.txt"
var KEY_NOT_FIND = "notFind"

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
