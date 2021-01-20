package main

import (
     "bufio"
     "os"
     "log"
     "strings"
)

var settingTXTPath = "../setting.txt"
var KEY_NOT_FIND = "notFind"

type ledMode struct {
    Led_mode int `json:"led_mode"`
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
