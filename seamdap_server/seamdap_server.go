package seamdap_server

import (
	"bufio"
	"fmt"
	"github.com/SEAMDAP/Demo/configs"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"

)
//////////////////////////SERVER
//////////////////////////SERVER
//////////////////////////SERVER

/*
	The Server is a simple SEAMDAP HTTP-based RESTful seamdap_server.
	After the startup, the seamdap_server accepts command-line read instructions. the accepted commands are:
	- start: starts the seamdap_server
	- stop: shut down the seamdap_server
	- help: print help text
	- status: print status of the seamdap_server
	- exit: stops the seamdap_server and close the program

*/


type Server struct {
	HttpServer *http.Server
	Listener   net.Listener
	Redis	*redis.Client
}

var Port = configs.Server_port
var done = make(chan bool, 1)
var initServer = make(chan bool, 1)
var server_flag = false
var server *Server

func (s *Server) ListenAndServe() error {

	// Serve accepts incoming HTTP connections on the listener l, creating a new service goroutine for each.
	// The service goroutines read requests and then call handler to reply to them.

	listener, err := net.Listen("tcp", s.HttpServer.Addr)
	if err != nil {
		return err
	}
	s.Listener = listener

	go s.HttpServer.Serve(s.Listener)

	fmt.Println("Server now listening")
	return nil
}

//////////////////////////MAIN
//////////////////////////MAIN
//////////////////////////MAIN


func StartAll() {
	//MAIN
	done = make(chan bool, 1)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type \"start\" to start the server")
	for {
		fmt.Println("$ SEAMDAP Server ready ")
		input, err := reader.ReadString('\n')
		if err != nil {
			if _, err_ := fmt.Fprintln(os.Stderr, err); err_ != nil {
				panic(err_)
			}
		}
		err = runCommand(input, done)
		if err != nil {
			if _, err_ := fmt.Fprintln(os.Stderr, err); err_ != nil {
				panic(err_)
			}
		}
	}
}

func runServer() {
	fmt.Println("Starting Server...")

	if strings.Compare(Port, "") == 0{
		fmt.Println("ERROR: missing port")
	}

	// A REDIS db is used to store some incoming data. Update the address and password as needed in configs/settings.go
	client_redis := redis.NewClient(&redis.Options{
		Addr: configs.Server_REDIS_fullAddress,
		Password: configs.Server_REDIS_pass,
		DB: 0,
	})

	HTTPServer := Routing(client_redis)
	server =&Server{ HttpServer: HTTPServer, Redis: client_redis}
	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}

	initServer <- true
	<-done
}

func runCommand(input string, done chan bool) error {
	input = strings.TrimSuffix(input, "\n")
	arrCommandStr := strings.Fields(input)

	if len(arrCommandStr) == 0 {
		fmt.Print("")
	} else {
		function, match := admin_cmd[arrCommandStr[0]]
		if match == true {
			function(arrCommandStr)
			return nil
		} else {
			fmt.Println(arrCommandStr[0] + " doesn't match to any command")
		}

		cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		return cmd.Run()
	}
	return nil
}

var admin_cmd = map[string]func(arrCommandString []string){
	"start":  startServer,
	"stop":   stopServer,
	"help": help,
	"exit": exit,
	"status":statusServer,
}

func startServer(arrCommandStr []string) {
	if server_flag == true {
		fmt.Println("Server already running")
	} else {
		go runServer()
		<-initServer
		server_flag = true
	}
}

func (s *Server) Shutdown() error {

	if s.Listener != nil {
		err := s.Listener.Close()
		s.Listener = nil
		if err != nil {
			return err
		}
	}

	if err :=s.Redis.Close();err != nil {
		panic(err)
	}
	fmt.Println("Shutting down seamdap_server")
	return nil
}

func stopServer(arrCommandStr []string) {
	if server_flag == true {

		done <- true
		if err:= server.Shutdown(); err != nil {
			panic(err)
		}
		server_flag = false
	} else {
		fmt.Println("Server is not running. Use 'start' to run the seamdap_server.")
	}
}

func statusServer(arrCommandStr []string) {
	if server_flag == true {
		fmt.Println("\nON - Server Status")
	} else {
		fmt.Println("\nOFF - Server Status")
	}

}

func help(arrCommandStr []string) {
	fmt.Println("Command List:")
	fmt.Println("'start': turns the seamdap_server on")
	fmt.Println("'stop': turns the seamdap_server off")
	fmt.Println("'status': shows seamdap_server, services and cron status")
	fmt.Println("'help': mmm....helps :)")
	fmt.Println("'exit': closes the program and gracefully shuts the seamdap_server down if needed")
}

func exit(arrCommandStr []string) {
	stopServer(arrCommandStr)
	os.Exit(0)
}


//////////////////////////ROUTING
//////////////////////////ROUTING
//////////////////////////ROUTING

func Routing(client_redis *redis.Client) *http.Server {

	router := mux.NewRouter()

	router.HandleFunc(fmt.Sprintf("/%s",configs.Server_URL_firstPhasePath), newSensorInterface(client_redis)).Methods("POST").Schemes("http")
	router.HandleFunc(fmt.Sprintf("/%s",configs.Server_URL_secondPhasePath), newSensorInstance(client_redis)).Methods("POST").Schemes("http")
	router.HandleFunc(fmt.Sprintf("/%s",configs.Server_URL_thirdPhasePath), newSensorSampling(client_redis)).Methods("POST").Schemes("http")
	router.HandleFunc(fmt.Sprintf("/%s/%s",configs.Server_URL_thirdPhasePath, "{TD_id}"), newSensorSampling(client_redis)).Methods("POST").Schemes("http")

	// CHECK
	//router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	//	tpl, err1 := route.GetPathTemplate()
	//	met, err2 := route.GetMethods()
	//	fmt.Println(tpl, err1, met, err2)
	//	return nil
	//})
	HTTPServer := &http.Server{Addr: ":" + Port, Handler: router}

	return HTTPServer
}