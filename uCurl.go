package main

import(
	"uCurl/src"
	"os"
	"fmt"
)

func main(){
	args := os.Args
	server := src.Parse(args[1])
	fmt.Println(server.Connect())
}