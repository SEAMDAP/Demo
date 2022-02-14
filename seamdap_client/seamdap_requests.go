package seamdap_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/SEAMDAP/Demo/configs"
	"github.com/SEAMDAP/Demo/utils"
	"github.com/google/uuid"
	"net/http"
	"time"
)

/*
	This file contains the request needed to reach the server, performing all the phases in SEAMDAP
*/

func InterfaceRegistration(TD utils.ThingDescription, IdPlot int32, date time.Time, clnt *SEAMDAPClient) (*http.Response, error,time.Time,time.Time) {

	url := fmt.Sprintf("http://%s:%s/%s", configs.Server_addr, configs.Server_port, configs.Server_URL_firstPhasePath)
	method := "POST"

	msg := TD

	jsonRequest, err := json.Marshal(msg)
	if err != nil {
		//logging.Error(logging.JSONError, "Error on data marshaling", err)
		return nil, err, time.Time{}, time.Time{}
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url,  bytes.NewBuffer(jsonRequest))
	if err != nil {
		fmt.Println(err)
	}
	// TEST *******************
	//fmt.Println(req.Body)
	//return nil, err,time.Time{},time.Time{}
	// ************************

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Here insert random UUID")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "gzip")

	timeReq := time.Now()
	resp, err := client.Do(req)
	timeResp := time.Now()
	if err != nil {
		//logging.Error(logging.HTTPError, "Error on http seamdap_client ", err)
		return nil, err, time.Time{}, time.Time{}
	}

	return resp, err, timeReq, timeResp
}

func InstanceRegistration(ins_ utils.InstanceRegistrationRequest) (*http.Response, error, time.Time,time.Time) {

	url := fmt.Sprintf("http://%s:%s/%s", configs.Server_addr, configs.Server_port, configs.Server_URL_secondPhasePath)
	method := "POST"

	msg := ins_

	jsonRequest, err := json.Marshal(msg)
	if err != nil {
		//logging.Error(logging.JSONError, "Error on data marshaling", err)
		return nil, err, time.Time{}, time.Time{}
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url,  bytes.NewBuffer(jsonRequest))

	if err != nil {
		fmt.Println(err)
	}
	// TEST *******************
	//fmt.Println(req.Body)
	//return nil, err,time.Time{},time.Time{}
	// ************************

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Here insert random UUID")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "gzip")

	timeReq := time.Now()
	resp, err := client.Do(req)
	timeResp := time.Now()
	if err != nil {
		//logging.Error(logging.HTTPError, "Error on http seamdap_client ", err)
		return nil, err, time.Time{}, time.Time{}
	}

	return resp, err, timeReq, timeResp
}

func UploadSampling(samp_ utils.Custom, TD_ID uuid.UUID) (*http.Response, error, time.Time,time.Time) {

	url := fmt.Sprintf("http://%s:%s/%s/%s", configs.Server_addr, configs.Server_port, configs.Server_URL_thirdPhasePath,TD_ID.String())
	method := "POST"

	msg := samp_

	jsonRequest, err := json.Marshal(msg)
	if err != nil {
		//logging.Error(logging.JSONError, "Error on data marshaling", err)
		return nil, err, time.Time{}, time.Time{}
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url,  bytes.NewBuffer(jsonRequest))

	if err != nil {
		fmt.Println(err)
	}
	// TEST *******************
	//fmt.Println(req.Body)
	//return nil, err,time.Time{},time.Time{}
	// ************************

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Here insert random UUID")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "gzip")

	timeReq := time.Now()
	resp, err := client.Do(req)
	timeResp := time.Now()
	if err != nil {
		//logging.Error(logging.HTTPError, "Error on http seamdap_client ", err)
		return nil, err, time.Time{}, time.Time{}
	}

	return resp, err, timeReq, timeResp
}