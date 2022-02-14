package main

import (
	"fmt"
	"github.com/SEAMDAP/Demo/configs"
	"github.com/SEAMDAP/Demo/seamdap_client"
	"github.com/SEAMDAP/Demo/seamdap_server"
	"github.com/SEAMDAP/Demo/utils"
	"math/rand"
	"os"
	"sync"
	"time"
)

/*
	SEAMDAP Protocol Demo

	The protocol is described in the paper "Seamless Sensor Data Acquisition for the Edge-to-Cloud Continuum".
	The paper was submitted for ACM PODC 2022: The 41st ACM Symposium on Principles of Distributed Computing (PODC 2022),
	and is now under review.

	Description of the Demo

	The demo is composed by two functional parts: the Client and the Server part.
	The Server is a SEAMDAP HTTP-based RESTful seamdap_server, capable of interpreting the SEAMDAP messages, for all the phases.
	He need the access to a Redis-Server to store received data.

	The Client is capable of emulating a dynamic set of SEAMDAP clients (one for each sensors node instance to be registered)
	The Client part is executed as follows. First, N goroutines are created, with N equal to the number of instances,
	which are in charge of obtaining, personalizing and uploading to the seamdap_server a file adhering to the Thing Description
	format. Each goroutine suddenly creates M sub-goroutines, with M equal to the number of instances that will be
	registered using the related TD interface. The instance's sub-goroutine also takes care of loading its samples
	on a regular basis.
	During the execution, spread along the various phases, some inactivity times and periodicities are randomly chosen
	to offset the various activities, respecting some min and max values.

	From the init.go file it is possible to run both the component.
	Usage:
	- go run init.go server
		- then it's possible to type some command. Type "start" to start the seamdap server
	- go run init.go client
*/


func main(){
	if len(os.Args) < 2 {
		fmt.Printf(" Error: missing argument. \nUsage:\t\t %s <modality>\n <modality> must be 'seamdap_server' or 'client' \n", os.Args[0])
		return
	}

	// Every message is still logged in a file,
	// but is preferred to use some network analysis tool to carry out performance evaluation
	defer utils.LOGFILE.Close()


	arg := os.Args[1]
	switch arg {
		case "server":
			fmt.Println("Server Mode")

			seamdap_server.StartAll()

		case "client":
			fmt.Println("Client Mode")

			rand.Seed(time.Now().UnixNano())
			utils.LogConfig_Parameters()
			startTime := time.Now()
			var wg sync.WaitGroup

			// Creation of the Instance registration sub-goroutines
			for i := 0; i < configs.Client_number; i ++{
				wg.Add(1)
				go seamdap_client.NewClient(rand.Intn(100000), i, &wg, configs.Client_maxLifeTime, startTime)
			}
			wg.Wait()
	default:
		fmt.Println("Unknown command: ", arg)
	}

	return
}


