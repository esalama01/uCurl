package src

import(
	"strings"
	"net"
)

type URL_PARSER struct{ //as in https://www.youtube.com/watch?v=WmMa1GVPWZ0
	Host string //aka domain name || IP address.
	Protocol string //protocol(defaults to http in my case.)
	Port string //sits between host and path. comes after a ":". (if not specified has a default of 80.(cause i m working with http.))
	Path string
}

func Parse(url string) (URL_PARSER){//suppose it's of the format [http://eu.httpbin.org/get/mhido]
	host, path, _ := strings.Cut(url[7:], "/") //what's between // and /
	_,port,found :=  strings.Cut(host, ":")

	parser := URL_PARSER{
		Host : host,
		Protocol : "HTTP/1.1",
		Path : path,
	}
	if (found){
		parser.Port = port
	}else{
		parser.Port = "80"
	}
	return parser
}

func (server URL_PARSER) Connect() (string){ //d'apres https://dzone.com/articles/socket-programming-in-go
	//establish connection
	connection, err := net.Dial("tcp", server.Host + ":" + server.Port) //i defaulted it to tcp bec i m not running twitch lol
	if err != nil {
        panic(err)
    }
	//read the data
	buffer := make([]byte, 1024)
	l, err := connection.Read(buffer)
	if err != nil {
        panic(err)
    }
	response := string(buffer[:l])
	defer connection.Close()
	return response
}