package network

import (
    "net"
    
    "github.com/henrikkorsgaard/wifi"
)

func CheckNetworkStatus(){
    return
}

//DetectAvailableNetworkInterfaces will return an array of available network interfaces
func DetectAvailableNetworkInterfaces(){

}

//What is the scenario here!
//if ip && correct network -> do nothing
//if ip && !correct network -> reconnect
//if no ip && network available -> connect to correct network
//if no ip && !network available || !network defined -> set up as ap with server+api


//We need to detect ethernet interfaces?
//We need to detect wifi interfaces?
//We need to figure out if a wifi interface is connected, has IP and if it's the right ap


//IsSSIDAvailble will return a boolean indicating if the provided SSID is available
func IsSSIDAvailable(ssid string) (isAvailable bool) {
    return
}

//detectAvailableStations should return a list of SSIDs nearby to match
func detectAvailableStations() (ssids []string){
    return
}


//GetConnectedWifiInterfaces returns the available wifi interfaces that has an IP associated.
func GetConnectedWifiInterfaces()(connectedInterfaces []net.Interface, err error){
    //We use the wifi library as it allow us to get only wifi devices. 
    //This gives us a base for searching for the net.Interfaces using the Golang net library.
    c, err := wifi.New()
    defer c.Close()
    if err != nil {
        return connectedInterfaces, err
    }


    ifaces, err := c.Interfaces()
    if err != nil {
        return connectedInterfaces, err
    }

    for _, iface := range ifaces {

        if iface.Name != "" {
            netiface, err := net.InterfaceByName(iface.Name)
            if err != nil {
                return connectedInterfaces, err
            }

            addrs, err := netiface.Addrs()
            if err != nil {
                return connectedInterfaces, err
            }

            for _, addr := range addrs {
                if ipnet, ok := addr.(*net.IPNet); ok {
                    if ipnet.IP.To4() != nil {
                        connectedInterfaces = append(connectedInterfaces, *netiface)
                    }
                }
            }
        }
    }

    return connectedInterfaces, err
}
