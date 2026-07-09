package main

import(
	"uCurl/src"
	"os"
	"fmt"
)

func main(){
	args := os.Args
	if (len(args) == 2){
		server := src.Parse(args[1])
		fmt.Println(server.Connect())
	}else{
		switch args[1]{
		case "-v":
			server := src.Parse(args[2])
			fmt.Println(server.Connect_Verbose())
		default:
			fmt.Println("Invalid Command")
		}
	}
}