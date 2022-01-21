package seamdap_client

import (
	"../utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func main(){
	fmt.Println("Hello")
}


func NewClient(id uuid.UUID,  wg *sync.WaitGroup, maxTime int, startTime time.Time){
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	timetoWake := rand.Intn(maxTime /3)
	time.Sleep(time.Duration(timetoWake)*time.Second)

	timetoTD := rand.Intn((maxTime/2) - timetoWake)
	time.Sleep(time.Duration(timetoTD)*time.Second)

	//TODO: instanceRegistration() --> return TD_id

	sensorsInstancesNumber := rand.Intn(15)
	var sub_wg sync.WaitGroup
	for s := 0; s < sensorsInstancesNumber; s++ {
		timeToInstance := rand.Intn((maxTime- (timetoWake+timetoTD)) / (sensorsInstancesNumber *2))
		sub_wg.Add(1)
		sensorInstanceRoutine(&sub_wg, uuid.New(), timeToInstance, maxTime, startTime) //TODO: use TD_id
	}
	wg.Wait()

	return
}

func sensorInstanceRoutine( wg *sync.WaitGroup, TDid uuid.UUID, sleepTime int, maxTime int, startTime time.Time ){
	defer wg.Done()
	time.Sleep(time.Duration(sleepTime)*time.Second)

	// TODO: instanceRegistration --> IN_id ?

	// Samples Communication period time
	commPeriod := rand.Intn(20*60) + 10*60 // from 10 up to 30 minutes
	time.Sleep(time.Duration(commPeriod)*time.Second)

	for{
		remainingSeconds := startTime.Add(time.Duration(maxTime)*time.Second).Sub(time.Now()).Seconds()
		if remainingSeconds < float64((2*commPeriod)){
			break
		}

		//TODO: uploadSampling

		time.Sleep(time.Duration(commPeriod)*time.Second)
	}


}

func interfaceRegistration(IdPlot int32, date time.Time) (*http.Response, error) {
	msg := utils.ThingDescription{}
	//url := "http://brie.ce.unipr.it/api/sensor/" + strconv.Itoa(int(IdPlot))
	url := "http://127.0.0.1/api/sensor/interface"
	method := "POST"

	jsonRequest, err := json.Marshal(msg)
	if err != nil {
		//logging.Error(logging.JSONError, "Error on data marshaling", err)
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url,  bytes.NewBuffer(jsonRequest))

	if err != nil {
		fmt.Println(err)
	}
	// TEST *******************
	//fmt.Println(req.Body)
	//return nil, err
	// ************************

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Here insert random UUID")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := client.Do(req)
	if err != nil {
		//logging.Error(logging.HTTPError, "Error on http seamdap_client ", err)
		return nil, err
	}

	return resp, err
}

func instanceRegistration(IdPlot int32, date time.Time) (*http.Response, error) {
	msg := utils.ThingDescription{}
	//url := "http://brie.ce.unipr.it/api/sensor/" + strconv.Itoa(int(IdPlot))
	url := "http://127.0.0.1/api/sensor/instance"
	method := "POST"

	jsonRequest, err := json.Marshal(msg)
	if err != nil {
		//logging.Error(logging.JSONError, "Error on data marshaling", err)
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url,  bytes.NewBuffer(jsonRequest))

	if err != nil {
		fmt.Println(err)
	}
	// TEST *******************
	//fmt.Println(req.Body)
	//return nil, err
	// ************************

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Here insert random UUID")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := client.Do(req)
	if err != nil {
		//logging.Error(logging.HTTPError, "Error on http seamdap_client ", err)
		return nil, err
	}

	return resp, err
}

func uploadSampling(IdPlot int32, date time.Time) (*http.Response, error) {
	msg := utils.ThingDescription{}
	//url := "http://brie.ce.unipr.it/api/sensor/" + strconv.Itoa(int(IdPlot))
	url := "http://brie.ce.unipr.it/api/sensor/" + strconv.Itoa(int(IdPlot))
	method := "POST"

	jsonRequest, err := json.Marshal(msg)
	if err != nil {
		//logging.Error(logging.JSONError, "Error on data marshaling", err)
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url,  bytes.NewBuffer(jsonRequest))

	if err != nil {
		fmt.Println(err)
	}
	// TEST *******************
	//fmt.Println(req.Body)
	//return nil, err
	// ************************

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Here insert random UUID")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := client.Do(req)
	if err != nil {
		//logging.Error(logging.HTTPError, "Error on http seamdap_client ", err)
		return nil, err
	}

	return resp, err
}