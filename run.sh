#!/usr/bin/zsh

go build -o ./bin/$1 cmd/$1/*.go && ./bin/$1 $1 $2