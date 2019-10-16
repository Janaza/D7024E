#!/bin/sh
#$rval = shuf -i 6001-9999 -n 1
port="-port"
bstrap="-bootstrapIP 10.0.0.2:6000"
go run /go/src/D7024E/main/main.go $port $1 $bstrap
