#!/bin/sh
# Test the Golang application
cd $GITHUB_WORKSPACE

go test ./... -v

