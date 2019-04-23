#!/bin/sh

go build -gcflags "all=-N -l" -o /server main.go
dlv --listen=:2345 --headless=true --api-version=2 exec /server