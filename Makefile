build:
	go build -o ./co2-cli .
	go build -o ./co2-mqtt ./mqtt

clean:
	rm -rf co2-cli
	rm -rf co2-mqtt
