build: build-mqtt build-http build-cli

build-mqtt:
		go build -o ./co2-mqtt ./mqtt
		
build-http:
		go build -o ./co2-http ./http

build-cli:
		go build -o ./co2-cli ./cli

clean:
		rm -rf co2-cli
		rm -rf co2-mqtt
