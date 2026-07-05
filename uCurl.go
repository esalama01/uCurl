package main

import(
	"uCurl/src"
	"os"
)

func main(){
	args := os.Args
	src.Print_protocole(args[1])
}