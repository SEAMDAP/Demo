package seamdap_client

import (
	"encoding/json"
	"fmt"
	"github.com/gPenzotti/SEAMDAP/configs"
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
	StartTime time.Time
}

func NewClient(id int, index int,  wg *sync.WaitGroup, maxTime int, startTime time.Time){
	defer wg.Done()
	rand.Seed(time.Now().UnixNano() + int64(index))

	MyClient := SEAMDAPClient{
		ID:    id,
		Index: index,
		TDId:  uuid.UUID{},
		UserAgent: "CLIENT_N"+strconv.Itoa(index)+"_ID"+strconv.Itoa(id),
		MSG_index: rand.Intn(5),
		StartTime:startTime,
	}

	timetoWake := rand.Intn(configs.Client_maxWakeTime)
	time.Sleep(time.Duration(timetoWake)*time.Second)
	configs.LogClientWake(MyClient.ID, time.Now().Sub(MyClient.StartTime))
	configs.LogConfig_RndTime(MyClient.ID, "TIME_TO_WAKE", timetoWake)

	timetoTD := rand.Intn(configs.Client_maxFirstPhaseTime - timetoWake)
	time.Sleep(time.Duration(timetoTD)*time.Second) //TODO: SLEEP
	configs.LogConfig_RndTime(MyClient.ID, "TIME_TO_FIRST_PHASE", timetoTD)


	//TODO: interfaceRegistration() --> return TD_id
	TD_id,err := TD_CreateAndRegister(&MyClient)
	fmt.Println(TD_id)
	if err != nil {
		fmt.Println("ERRORE: ", err)
		return
	}

	MyClient.TDId = TD_id

	sensorsInstancesNumber := rand.Intn(configs.Client_maxSensorInstance)
	configs.LogConfig_InstanceNumber(MyClient.ID, sensorsInstancesNumber)

	var sub_wg sync.WaitGroup
	for s := 0; s < sensorsInstancesNumber; s++ {

		timeToInstance := rand.Intn(configs.Client_maxSecondPhaseTime - (timetoTD + timetoWake))
		configs.LogConfig_RndTime(MyClient.ID, "TIME_TO_SECOND_PHASE_"+strconv.Itoa(s), timeToInstance)

		commPeriod := rand.Intn( configs.Client_maxCommunicationPeriodRange[1] - configs.Client_maxCommunicationPeriodRange[0]) + configs.Client_maxCommunicationPeriodRange[0]
		configs.LogConfig_RndTime(MyClient.ID, "TIME_THIRD_PHASE_PERIOD_"+strconv.Itoa(s), commPeriod)


		sub_wg.Add(1)
		go sensorInstanceSubRoutine(&sub_wg, MyClient, commPeriod, timeToInstance, maxTime, startTime)
	}
	sub_wg.Wait()

	return
}

func sensorInstanceSubRoutine( sub_wg *sync.WaitGroup,client SEAMDAPClient, commPeriod int, sleepTime int, maxTime int, startTime time.Time ){
	defer sub_wg.Done()
	time.Sleep(time.Duration(sleepTime)*time.Second) //TODO: SLEEP

	// TODO: instanceRegistration --> IN_id ?
	instID, err := INSTANCE_CreateAndRegister(&client)
	if err != nil{
		fmt.Println(err)
	}
	// Samples Communication period time

	for{

		// Check and sleep
		remainingSeconds := startTime.Add(time.Duration(maxTime)*time.Second).Sub(time.Now()).Seconds()
		fmt.Printf("REMAINING SEC: %f, 2*CommPeriod: %f, check : %t \n",remainingSeconds,float64(2*commPeriod), remainingSeconds < float64(2*commPeriod) )
		if remainingSeconds < float64(2*commPeriod){
			fmt.Println("FINITO")
			break
		}
		//TODO: uploadSampling
		stat, err := SAMPLE_CreateAndUpload(&client, instID)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println("UPLOAD COMPLETE: ",stat)

		//fmt.Println("SLEEPIAMO PER SECONDI:", commPeriod)
		time.Sleep(time.Duration(commPeriod)*time.Second) //TODO: SLEEP
	}
	fmt.Println("SUPER FINITO")
	return

}

func TD_CreateAndRegister(client *SEAMDAPClient) (uuid.UUID,error){
	fmt.Println("Registering TD...")

	TD := configs.TestMessagesTD[client.MSG_index]
	TD.ID = client.TDId.String()
	TD.Title ="TD_TITLE_EXAMPLE_" + strconv.Itoa(client.Index)
	TD.Model = "TD_MODEL_EXAMPLE_" + strconv.Itoa(client.Index)
	TD.Description = "TD_DESC_EXAMPLE_" + strconv.Itoa(client.Index)

	resp,err, timeReq, timeResp := InterfaceRegistration(TD, 0, time.Now(), client)
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

	configs.LogClientSendTD(client.ID, timeReq.Sub(client.StartTime), ns.UID.String())
	configs.LogClientSendTD_response(client.ID, timeResp.Sub(client.StartTime), ns.UID.String())

	return ns.UID, nil
}

func INSTANCE_CreateAndRegister(client *SEAMDAPClient) (int,error){
	fmt.Println("Registering Instance...")

	// Creazione del messaggio di istanza
	instance_request := utils.InstanceRegistrationRequest{
		TDID:      client.TDId,
		UserID:    client.ID,
		BoardName: "BOARD_NAME_"+strconv.Itoa(client.Index),
		Area:      utils.NewGeojsonFeature(),
	}

	// Invio al Server e ricezione risposta

	resp, err, timeReq, timeResp := InstanceRegistration(instance_request)
	if err != nil{
		fmt.Println("ERROR ERROR ", err)
		return 0, err
	}
	if resp.StatusCode != 200{
		fmt.Println("Failed request - status code != 200: ", err)
		return 0, err
	}

	// Parsing risposta e ritorno valori necessari
	ns := utils.InstanceRegistrationResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ns)
	if err != nil {
		fmt.Println("Failed parsing InstanceRegistrationResponse: ", err.Error())
		return 0, err
	}
	fmt.Println("Received response: ", ns)

	configs.LogClientSendInstance(client.ID, timeReq.Sub(client.StartTime), strconv.Itoa(ns.InstanceID))
	configs.LogClientSendInstance_response(client.ID, timeResp.Sub(client.StartTime), strconv.Itoa(ns.InstanceID))

	return ns.InstanceID, nil
}

func SAMPLE_CreateAndUpload(client *SEAMDAPClient, instanceID int) (string,error){

	fmt.Println("Uploading sample...")

	// Creazione del messaggio di sampling

	rec,err := configs.GetSENML(client.MSG_index)
	if err != nil {
		fmt.Println("Error in generating Senml: ", err)
	}
	rec.TimeRecord = time.Now().Format("2006.01.02T15:04:05")
	rec.Name = strconv.Itoa(instanceID)

	msg := utils.Custom{
		Record: []utils.SenMLPos{rec},
	}


	// Invio al Server e ricezione risposta
	resp, err, timeReq, timeResp := UploadSampling(msg, client.TDId)
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

	tempid:=rand.Intn(1000000)
	configs.LogClientSendSample(client.ID, timeReq.Sub(client.StartTime), strconv.Itoa(tempid))
	configs.LogClientSendSample_response(client.ID, timeResp.Sub(client.StartTime), strconv.Itoa(tempid))

	return ns.Status, nil
}