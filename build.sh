#!/bin/bash

glide install
go get ./...
GOOS=linux go build -ldflags "-extldflags '-static'" -o bin/cleaner
go build -race -ldflags "-extldflags '-static'" -o bin/cleaner-darwin
