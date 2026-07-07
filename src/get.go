package src

import(
	"strconv"
	"strings"
)

type URL_PARSER struct{ //as in https://www.youtube.com/watch?v=WmMa1GVPWZ0
	Host string //aka domain name || IP address.
	Scheme string //protocol(defaults to http in my case.)
	Port int //sits between host and path. comes after a ":". (if not specified has a default of 80.(cause i m working with http.))
	Path string
}

func Parse(url string) (URL_PARSER){//suppose it's of the format [http://eu.httpbin.org/get/mhido]
	host, path, _ := strings.Cut(url[7:], "/") //what's between // and /
	_,port,found :=  strings.Cut(host, ":")

	parser := URL_PARSER{
		Host : host,
		Scheme : "HTTP/1.1",
		Path : path,
	}
	if (found){
		port, _ := strconv.Atoi(port)
		parser.Port = port
	}else{
		parser.Port = 80
	}
	return parser
}