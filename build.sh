#!/bin/bash

glide install
go get ./...
go build -race -ldflags "-extldflags '-static'" -o bin/cleaner
