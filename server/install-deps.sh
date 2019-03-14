#!/bin/bash

echo "Installing dependencies to $GOPATH/pkg"
for line in `cat ./deps.list`; do
    echo "Installing $line"
    go get $line
done