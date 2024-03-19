#!/bin/bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 GO1111MODULE=on go build -o get_lrc .