package server

import (
	"bufio"
	"fmt"
	"github.com/gPenzotti/SEAMDAP/configs"
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


type Server struct {
	HttpServer *http.Server
	Listener   net.Listener
}
var Port = configs.Server_port
var done = make(chan bool, 1)
var initServer = make(chan bool, 1)
var server_flag = false
var server *Server

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", s.HttpServer.Addr) //sempre in ascolto
	if err != nil {
		return err
	}
	s.Listener = listener

	// Uso un'altra goroutine così sulla principale posso aspettare sul canale "done" e intercettare i segnali
	go s.HttpServer.Serve(s.Listener)
	// Serve accepts incoming HTTP connections on the listener l, creating a new service goroutine for each.
	// The service goroutines read requests and then call handler to reply to them.

	fmt.Println("Server now listening")

	return nil
}

func (s *Server) Shutdown() error {

	if s.Listener != nil {
		err := s.Listener.Close()
		s.Listener = nil
		if err != nil {
			return err
		}
	}
	//s.Db.Close()
	fmt.Println("Shutting down server")
	return nil
}




//////////////////////////MAIN
//////////////////////////MAIN
//////////////////////////MAIN


func StartAll() { //MAIN
	done = make(chan bool, 1)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("$ SEAMDAP Server ready ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(input, done)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}



func runServer() {
	fmt.Println("Starting Server...")

	if Port == "" {
		fmt.Println("ERROR: missing port")
	}

	client_redis := redis.NewClient(&redis.Options{
		Addr: configs.Server_REDIS_fullAddress,
		Password: configs.Server_REDIS_pass,
		DB: 0,
	})
	HTTPServer := Routing(client_redis)
	server =&Server{ HttpServer: HTTPServer}
	server.ListenAndServe()

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

func stopServer(arrCommandStr []string) {
	if server_flag == true {
		done <- true
		server.Shutdown()
		server_flag = false
	} else {
		fmt.Println("Server is not running. Use 'start' to run the server.")
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
	fmt.Println("'start': turns the server on")
	fmt.Println("'stop': turns the server off")
	fmt.Println("'status': shows server, services and cron status")
	fmt.Println("'help': mmm....helps :)")
	fmt.Println("'exit': closes the program and gracefully shuts the server down if needed")
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
	//subRouterApi := router.PathPrefix("/api").Subrouter()
	//subRouterApiSensor := subRouterApi.PathPrefix("/sensor").Subrouter()
	//subRouterApiSensor.HandleFunc("/interface", newSensorInterface(client_redis)).Methods("POST").Schemes("http")
	//subRouterApiSensor.HandleFunc("/instance", newSensorInstance(client_redis)).Methods("POST").Schemes("http")
	//subRouterApiSensor.HandleFunc("/data", newSensorSampling(client_redis)).Methods("POST").Schemes("http")
	//subRouterApiSensor.HandleFunc("/data/{TD_id}", newSensorSampling(client_redis)).Methods("POST").Schemes("http")


	router.HandleFunc(fmt.Sprintf("/%s",configs.Server_URL_firstPhasePath), newSensorInterface(client_redis)).Methods("POST").Schemes("http")
	router.HandleFunc(fmt.Sprintf("/%s",configs.Server_URL_secondPhasePath), newSensorInstance(client_redis)).Methods("POST").Schemes("http")
	router.HandleFunc(fmt.Sprintf("/%s",configs.Server_URL_thirdPhasePath), newSensorSampling(client_redis)).Methods("POST").Schemes("http")
	router.HandleFunc(fmt.Sprintf("/%s/%s",configs.Server_URL_thirdPhasePath, "{TD_id}"), newSensorSampling(client_redis)).Methods("POST").Schemes("http")
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err1 := route.GetPathTemplate()
		met, err2 := route.GetMethods()
		fmt.Println(tpl, err1, met, err2)
		return nil
	})
	HTTPServer := &http.Server{Addr: ":" + Port, Handler: router}

	return HTTPServer
}