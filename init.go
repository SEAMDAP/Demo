package main

import (
	"./seamdap_client"
	"fmt"
	"github.com/google/uuid"
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

	clientMaxLifeTime := 24*3600 //seconds
	startTime := time.Now()
	// Attivare tutti i seamdap_client: per ogni seamdap_client generare un UUID univoco
	var wg sync.WaitGroup

	id, err := uuid.NewUUID()
	fmt.Println(id, err, clientMaxLifeTime)

	for i :=0; i < 10; i ++{
		wg.Add(1)
		go seamdap_client.NewClient(uuid.New(), &wg, clientMaxLifeTime, startTime)
	}
	wg.Wait()
	return
}


