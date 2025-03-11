#!/bin/bash

go build -ldflags="-s -w" -v
./go-serial-sender temperature /dev/ttyUSB0