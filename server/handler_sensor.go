package server

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/paulmach/orb"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-redis/redis"

	"github.com/gPenzotti/SEAMDAP/utils"

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
//type NewSensorRes struct {
//	TDID           utils.RandomId
//	Name          string
//	Owner         string
//	Creation_time time.Time
//}

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

func newSensorInterface(client_redis *redis.Client) func(w http.ResponseWriter, r *http.Request) {
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
		if td.Title == "" {
			fmt.Println("Failed parsing Thing Description")
			return
		}

		id, err := uuid.NewUUID()
		if err !=nil{
			fmt.Println("Error in generating UUID: ", err)
		}
		res := utils.NewSensorRes{
			UID:          id, // Andrebbe creato random qui e restituito al client
			Name:         td.Model,
			Owner:        td.Manufacturer,
			CreationTime: time.Now(),
		}
		err = client_redis.Set(res.UID.String(), res, 0).Err()


		w.WriteHeader(http.StatusOK)
		rs, _ := json.Marshal(res)
		w.Write(rs)

		return

	}
}

func newSensorInstance(client_redis *redis.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		//Read body
		inst_ := utils.InstanceRegistrationRequest{}
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &inst_)
		if err != nil {
			fmt.Println("Failed parsing InstanceRegistrationRequest: ", err.Error())
			return
		}

		// CHECK TD EXISTANCE
		val, err := client_redis.Get(inst_.TDID.String()).Result()
		if err != nil {
			fmt.Println("ERRORE, NESSUN TD TROVATO: ",err)
		}
		fmt.Println(val)


		response := utils.InstanceRegistrationResponse{
			InstanceID:    uuid.New(),
			Endpoint:     "BLBLBL",
			CreationTime: time.Time{},
			BoardName:    "BLBLBL",
			Manufacturer: "BLBLBL",
		}

		instance := utils.SensorInstance{
			UID:          response.InstanceID,
			TD_ID:        inst_.TDID,
			CreationTime: response.CreationTime,
			OwnerID:      inst_.UserID,
		}
		err = client_redis.Set(response.InstanceID.String(), instance, 0).Err()
		if err != nil {
			fmt.Println("ERRORE, Scrittura non riuscita: ",err)
		}

		w.WriteHeader(http.StatusOK)
		rs, _ := json.Marshal(response)
		w.Write(rs)

		return
	}
}

func newSensorSampling(client_redis *redis.Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		string_instance_ID := mux.Vars(r)["instance_id"]
		fmt.Println(string_instance_ID)

		//Read body
		samp_ := utils.Custom{}
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &inst_)
		if err != nil {
			fmt.Println("Failed parsing InstanceRegistrationRequest: ", err.Error())
			return
		}

		// CHECK TD EXISTANCE
		for _,v := range samp_.Record{
			val, err := client_redis.Get(v.Name).Result()
			if err != nil {
				fmt.Println("ERRORE, NESSUN TD TROVATO: ",err)
			}
			fmt.Println(v)
			fmt.Println(val)
		}

		response := utils.SamplingResponse{Status: string_instance_ID}

		w.WriteHeader(http.StatusOK)
		rs, _ := json.Marshal(response)
		w.Write(rs)

		return
	}
}



