package server

import (
	"encoding/json"
	"fmt"
	"github.com/paulmach/orb"
	"io/ioutil"
	"time"
	"net/http"

	"../utils"

	geojson "github.com/paulmach/orb/geojson"

)


//Get only useful data in InsertSensorData()
type DataSave struct {
	ID               int64
	ThingDescription []byte
}

/*
Response for NewSensor()
*/
type NewSensorRes struct {
	UID           utils.RandomId
	Name          string
	Owner         string
	Creation_time time.Time
}

/*
Request for AddSensor()
*/
type extServer struct {
	URL string  `json:"Url"`
	Period int32  `json:"Period"`
}

type AddSensorTestReq struct {
	TD			utils.ThingDescription  `json:"TD"`
	UserID		int64  `json:"UserID"`
	PlotID		int64  `json:"PlotID"`
	Name 		string 	`json:"Name"`
	Position	orb.Point  `json:"Position"`
	Area		*geojson.FeatureCollection  `json:"Area"`
	Server		extServer  `json:"Server"`
}

/*
Return info for ActiveSensorList()
*/
type RetSensorInfo struct {
	UID           utils.RandomId
	Creation_time time.Time
	ProductName   string
	Owner         string
	Note          string
}

func newSensorInterface() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		//Read body
		td := utils.ThingDescription{}
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &td)
		if err != nil {
			fmt.Println("Failed parsing Thing Description: ", err.Error())
			return
		}

		// Checks
		// if TD title not set, raise request error
		if td.Title == nil {
			fmt.Println("Failed parsing Thing Description")
			return
		}


		res := NewSensorRes{UID: s.SensorUID, Name: s.ProductName, Owner: s.Owner, Creation_time: s.Creation_time}
		w.WriteHeader(http.StatusOK)
		rs, _ := json.Marshal(res)
		w.Write(rs)

		return

	}
}
func newSensorInstance() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		//Token to check if request is authorized
		//tokenIF := r.Header.Get("Authorization")

		//Read body
		body, _ := ioutil.ReadAll(r.Body)
		sen := utils.Custom{}
		err := json.Unmarshal([]byte(body), &sen)
		if err != nil {
			fmt.Println("Failed parsing InsertSensorData: ", err.Error())
			return
		}


		fmt.Printf(" Added %d data in %s.\n")

		return
	}
}

func newSensorSampling() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		//Token to check if request is authorized
		//tokenIF := r.Header.Get("Authorization")

		//Read body
		body, _ := ioutil.ReadAll(r.Body)
		sen := utils.Custom{}
		err := json.Unmarshal([]byte(body), &sen)
		if err != nil {
			fmt.Println("Failed parsing InsertSensorData: ", err.Error())
			return
		}


		fmt.Printf(" Added %d data in %s.\n")

		return
	}
}



