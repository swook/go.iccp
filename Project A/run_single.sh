#!/bin/bash

# Check if go command exists
if ! $(command -v go > /dev/null 2>&1); then
	echo "Error: go command not accessible. You can install go from http://golang.org/doc/install"
	exit
fi

# Check if $GOPATH set
# If go env setup properly, should exist
if [ ! ${GOPATH+x} ]; then
	echo "Error: \$GOPATH not set. Please see http://golang.org/doc/install#install"
	exit
fi

# Check if plotinum installed, if not then install it
if [ ! -d $GOPATH"/src/code.google.com/p/plotinum/plot" ]; then
	echo "Installing plotinum library..."
	go get code.google.com/p/plotinum/plotutil
fi

cd single/
go run *.go
