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

# Check if plotinum installed, and if not, install
if [ ! -d $GOPATH"/src/code.google.com/p/plotinum/plotutil" ]; then
	echo "Installing plotinum library..."
	go get code.google.com/p/plotinum/plotutil
fi

# Check if my libraries are installed, and if not, install
# Check if plotinum installed, if not then install it
if [ ! -d $GOPATH"/src/github.com/swook/go.iccp/matrix" ]; then
	echo "Installing go.iccp/matrix library..."
	go get github.com/swook/go.iccp/matrix
fi

if [ ! -d $GOPATH"/src/github.com/swook/gogsl" ]; then
	echo "Installing gogsl library..."
	go get github.com/swook/gogsl
fi

go run *.go
