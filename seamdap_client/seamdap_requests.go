package seamdap_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gPenzotti/SEAMDAP/utils"
	"github.com/google/uuid"
	"net/http"
	"time"
)


func  InterfaceRegistration(TD utils.ThingDescription, IdPlot int32, date time.Time) (*http.Response, error) {
	msg := TD
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

func InstanceRegistration(ins_ utils.InstanceRegistrationRequest) (*http.Response, error) {
	msg := ins_
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

func UploadSampling(samp_ utils.Custom, instance_ID uuid.UUID) (*http.Response, error) {
	msg := samp_
	//url := "http://brie.ce.unipr.it/api/sensor/" + strconv.Itoa(int(IdPlot))
	url := "http://brie.ce.unipr.it/api/sensor/data/" + instance_ID.String() // DEVO METTERE QUI L'ID SENNO DOVE?
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