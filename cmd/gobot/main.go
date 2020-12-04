package main

import (
	"flag"
	"gobot"
)

// Variables used for command line parameters
var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	gobot.Gobot(Token)
}
