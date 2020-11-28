p
ckage configuration

import (
   
    "os"
    "fmt"
    "log"
    
    "gopkg.in/ini.v1" 
)

type Config struct {
    Wifi
    Settings
}

type Wifi struct {
    Ssid string `ini:"ssid"`
    Password string `ini:"password"`
}

type Settings struct {
    Ethernet string `ini:"ethernet"`
    Api string `ini:"api"`
}

var (
    config Config
)

func init(){
    iniFilePath := "rpi-wifi-setup.ini"

    if _, err := os.Stat(iniFilePath); os.IsNotExist(err) {
        config = Config{Wifi{"",""}, Settings{"disabled","disabled"}}
    
        cfg := ini.Empty()
        err = ini.ReflectFrom(cfg, &config)
        if err != nil {
            log.Fatal(err)
        }

        err = cfg.SaveTo(iniFilePath)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("Generating the main ini file")
    } else {
        config = Config{}
        err = ini.MapTo(&config, iniFilePath)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(config.Ethernet)
    }

}

func What(){
    fmt.Println("just to test init funciton")
}

