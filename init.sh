#!/usr/bin/zsh

start="package main

import lib \"github.com/thebenkogan/Advent-Of-Code-2015\"

func main() {
	input := lib.GetInput()
}"

if test -d cmd/$1; then
  echo "Directory already exists."
  exit 0
fi

mkdir cmd/$1 && echo $start > cmd/$1/main.go && touch cmd/$1/in.txt