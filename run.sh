#!/bin/bash
source env/ajira.env

export AJIIRA_NET_SERVICE_PROTOCOL=https
export AJIIRA_NET_SERVICE_HOST_NAME=ajiiranetservice
export AJIIRA_NET_SERVICE_PORT_NUMBER=8080

go run *.go