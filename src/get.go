package src

import(
	"fmt"
	"strings"
)

func Print_protocole(url string){//suppose it's of the format [http://eu.httpbin.org/get]
	full_address := url[7:]	//what's between // and /
	address, _, _ := strings.Cut(full_address, "/")
	fmt.Printf("connecting to %s\nSending request GET /get HTTP/1.1\nHost:%s\nAccept: */*\n",address, address)
}