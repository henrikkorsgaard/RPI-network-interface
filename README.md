# RPI-wifi-setup

This project seek to develop a way to setup Raspberry PI network interfaces based on availability and interface capabilities. When a Raspberry PI boots this service will do the following:

1. Check for existing network connections (eth/wifi). If succesfully connecting it will print out details to stdout.
2. If unable to connect to an existing network, it will try to setup an AP using hostapd and dnsmasq that allow the user to connect via ssh and/or configure the device.
3. If unable to setup AP, it will print message to stdout and stderr.
