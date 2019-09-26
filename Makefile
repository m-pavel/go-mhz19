build:
	go build -o ./mhz19-cli .
	go build -o ./mhz19-mqtt ./mqtt

clean:
	rm -rf mhz19-cli
	rm -rf mhz19-mqtt
