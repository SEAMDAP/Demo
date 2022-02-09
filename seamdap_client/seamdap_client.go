package seamdap_client

import (
	"encoding/json"
	"fmt"
	"github.com/gPenzotti/SEAMDAP/test"
	"github.com/gPenzotti/SEAMDAP/utils"
	"github.com/google/uuid"
	"io/ioutil"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main(){
	fmt.Println("Hello")
}

type SEAMDAPClient struct{
	ID int
	Index int
	TDId uuid.UUID
	UserAgent string
	MSG_index int
}

func NewClient(id uuid.UUID, index int,  wg *sync.WaitGroup, maxTime int, startTime time.Time){
	defer wg.Done()
	rand.Seed(time.Now().UnixNano() + int64(index))

	clientID:=rand.Intn(100000)
	MyClient := SEAMDAPClient{
		ID:    clientID,
		Index: index,
		TDId:  uuid.UUID{},
		UserAgent: "CLIENT_N"+strconv.Itoa(index)+"_ID"+strconv.Itoa(clientID),
		MSG_index: rand.Intn(5),
	}

	timetoWake := rand.Intn(maxTime /3)
	//time.Sleep(time.Duration(timetoWake)*time.Second)//TODO: SLEEP

	timetoTD := rand.Intn((maxTime/2) - timetoWake)
	//time.Sleep(time.Duration(timetoTD)*time.Second) //TODO: SLEEP

	//TODO: interfaceRegistration() --> return TD_id
	TD_id,err := TD_CreateAndRegister(&MyClient)
	fmt.Println(TD_id)
	if err != nil {
		fmt.Println("ERRORE: ", err)
		return
	}

	MyClient.TDId = TD_id

	sensorsInstancesNumber := rand.Intn(15)
	var sub_wg sync.WaitGroup
	for s := 0; s < sensorsInstancesNumber; s++ {
		timeToInstance := rand.Intn((maxTime- (timetoWake+timetoTD)) / (sensorsInstancesNumber *2))
		sub_wg.Add(1)
		sensorInstanceSubRoutine(&sub_wg, MyClient, timeToInstance, maxTime, startTime)
	}
	wg.Wait()

	return
}

func sensorInstanceSubRoutine( wg *sync.WaitGroup,client SEAMDAPClient, sleepTime int, maxTime int, startTime time.Time ){
	defer wg.Done()
	//time.Sleep(time.Duration(sleepTime)*time.Second) //TODO: SLEEP

	// TODO: instanceRegistration --> IN_id ?
	instID, err := INSTANCE_CreateAndRegister(&client)
	if err != nil{
		fmt.Println(err)
	}
	// Samples Communication period time
	commPeriod := rand.Intn(20*60) + 10*60 // from 10 up to 30 minutes
	//time.Sleep(time.Duration(commPeriod)*time.Second) //TODO: SLEEP

	for{

		remainingSeconds := startTime.Add(time.Duration(maxTime)*time.Second).Sub(time.Now()).Seconds()
		if remainingSeconds < float64((2*commPeriod)){
			break
		}

		//TODO: uploadSampling
		stat, err := SAMPLE_CreateAndUpload(&client, instID)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println("UPLOAD COMPLETE: ",stat)
		//time.Sleep(time.Duration(commPeriod)*time.Second)
		time.Sleep(time.Duration(10)*time.Second)
	}


}

func TD_CreateAndRegister(client *SEAMDAPClient) (uuid.UUID,error){
	fmt.Println("Registering TD...")

	TD := test.TestMessagesTD[client.MSG_index]

	resp,err := InterfaceRegistration(TD, 0, time.Now())
	if err != nil{
		fmt.Println("ERROR ERROR ", err)
		return uuid.New(), err
	}
	if resp.StatusCode != 200{
		fmt.Println("Failed request - status code != 200: ", err)
		return uuid.New(), err
	}

	ns := utils.NewSensorRes{}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ns)
	if err != nil {
		fmt.Println("Failed parsing NewSensorRes: ", err.Error())
		return uuid.New(), err
	}

	fmt.Println("Received response: ", ns)
	return ns.UID, nil
}

func INSTANCE_CreateAndRegister(client *SEAMDAPClient) (uuid.UUID,error){
	fmt.Println("Registering Instance...")

	// Creazione del messaggio di istanza
	instance_request := utils.InstanceRegistrationRequest{
		TDID:      client.TDId,
		UserID:    client.ID,
		BoardName: "BOARD_NAME_"+strconv.Itoa(client.Index),
		Area:      utils.NewGeojsonFeature(),
	}

	// Invio al Server e ricezione risposta

	resp, err := InstanceRegistration(instance_request)
	if err != nil{
		fmt.Println("ERROR ERROR ", err)
		return uuid.New(), err
	}
	if resp.StatusCode != 200{
		fmt.Println("Failed request - status code != 200: ", err)
		return uuid.New(), err
	}

	// Parsing risposta e ritorno valori necessari
	ns := utils.InstanceRegistrationResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ns)
	if err != nil {
		fmt.Println("Failed parsing InstanceRegistrationResponse: ", err.Error())
		return uuid.New(), err
	}
	fmt.Println("Received response: ", ns)
	return ns.InstanceID, nil
}

func SAMPLE_CreateAndUpload(client *SEAMDAPClient, instanceID uuid.UUID) (string,error){

	fmt.Println("Uploading sample...")

	// Creazione del messaggio di sampling

	rec,err := test.GetSENML(client.MSG_index)
	if err != nil {
		fmt.Println("Error in generating Senml: ", err)
	}
	rec.TimeRecord = time.Now().Format("2006.01.02T15:04:05")
	rec.Name = instanceID.String()

		msg := utils.Custom{
		Record: []utils.SenMLPos{rec},
	}


	// Invio al Server e ricezione risposta

	resp, err := UploadSampling(msg, instanceID)
	if err != nil{
		fmt.Println("ERROR ERROR ", err)
		return "", err
	}
	if resp.StatusCode != 200{
		fmt.Println("Failed request - status code != 200: ", err)
		return "", err
	}

	// Parsing risposta e ritorno valori necessari
	ns := utils.SamplingResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ns)
	if err != nil {
		fmt.Println("Failed parsing InstanceRegistrationResponse: ", err.Error())
		return "", err
	}

	fmt.Println("Received response: ", ns)
	return ns.Status, nil
}