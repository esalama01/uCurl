package src

import(
	"fmt"
	"strings"
)

type URL_PARSER struct{ //as in https://www.youtube.com/watch?v=WmMa1GVPWZ0
	Host string //aka domain name || IP address.
	Scheme string //protocol(defaults to http in my case.)
	Port int //sits between host and path. comes after a ":". (if not specified has a default of 80.(cause i m working with http.))
	Path string
}

func Parse(url string){//suppose it's of the format [http://eu.httpbin.org/get/mhido]
	parser := URL_PARSER{
		
	}
}