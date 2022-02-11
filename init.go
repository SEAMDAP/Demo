package main

import (
	"fmt"
	"github.com/gPenzotti/SEAMDAP/configs"
	"github.com/gPenzotti/SEAMDAP/seamdap_client"
	"github.com/gPenzotti/SEAMDAP/server"
	"math/rand"
	"os"
	"sync"
	"time"
)

/*
	Ogni seamdap_client è una go routine che prima fa la registrazione della classe, poi di n istanze e poi comunica diverse volte i dati rilevati sul path fornito.
	Ogni richiesta dovrebbe contenere un token univoco e identificativo in modo da associre le richieste e le risposte
	in modo semplice (il server ovviamente deve replicare con queste).
	Possiamo fare la simulazione con gli hub solo in modalità seamdap_client?
*/


func main(){

	if len(os.Args) < 2 {
		fmt.Printf(" Error: missing argument. \nUsage:\t\t %s <modality>\n <modality> must be in ['server','client'] \n", os.Args[0])
		return
	}


	defer configs.LOGFILE.Close()

	serv := func() {
		server.StartAll()
	}
	client := func() {


		rand.Seed(time.Now().UnixNano())

		clientMaxLifeTime := configs.Client_maxLifeTime //seconds
		numberOfClient := configs.Client_number
		startTime := time.Now()

		configs.LogConfig_Parameters()

		// Attivare tutti i seamdap_client: per ogni seamdap_client genera un UUID univoco
		var wg sync.WaitGroup

		for i := 0; i < numberOfClient; i ++{
			wg.Add(1)
			go seamdap_client.NewClient(rand.Intn(100000), i, &wg, clientMaxLifeTime, startTime)
		}
		wg.Wait()
	}

	arg := os.Args[1]
	switch arg {
	case "server":
		fmt.Println("Only Server")
		serv()

	case "client":
		fmt.Println("Only Client")
		client()

	}





	return
}


