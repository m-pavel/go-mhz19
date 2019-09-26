build:
	go build . -o mhz19-cli
	go build ./mqtt -o mhz19-mqtt

clena:
	rm -rf mhz19-cli
	rm -rf mhz19-mqtt
