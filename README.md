# go-mhz19
  * Go bindings for the MH-Z19 CO2 Module
  * MQTT client for the hass.io integration
  * For some bizarre reason the default for Pi3 using the latest 4.4.9 kernel is to DISABLE UART. To enable it you need to change enable_uart=1 in /boot/config.txt. 
  * Getty service must be stopped sudo service serial-getty@ttyS0 stop
  * And disabled sudo systemctl mask serial-getty@ttyS0.service
