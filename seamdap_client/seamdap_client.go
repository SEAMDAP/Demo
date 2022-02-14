package seamdap_client

import (
	"encoding/json"
	"fmt"
	"github.com/SEAMDAP/Demo/configs"
	"github.com/SEAMDAP/Demo/utils"
	"github.com/google/uuid"
	"io/ioutil"
	"math/rand"
	"strconv"
	"sync"
	"time"
)
//////////////////////////CLIENT
//////////////////////////CLIENT
//////////////////////////CLIENT

/*
	A dynamic pool of Clients are emulated using the concurrency of the goroutine.
*/

var LOG = true // Enable log on file
var VERBOSE = true // Enable print

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
		UserAgent: "CLIENT_N"+strconv.Itoa(index)+"_ID"+strconv.Itoa(id), //unused
		MSG_index: rand.Intn(5),
		StartTime:startTime,
	}

	// First Inactivity time
	timetoWake := rand.Intn(configs.Client_maxWakeTime)
	time.Sleep(time.Duration(timetoWake)*time.Second)

	if LOG{
		utils.LogClientWake(MyClient.ID, time.Now().Sub(MyClient.StartTime))
		utils.LogConfig_RndTime(MyClient.ID, "TIME_TO_WAKE", timetoWake)
	}

	// Second Inactivity time
	timetoTD := rand.Intn(configs.Client_maxFirstPhaseTime - timetoWake)
	time.Sleep(time.Duration(timetoTD)*time.Second)

	if LOG{
		utils.LogConfig_RndTime(MyClient.ID, "TIME_TO_FIRST_PHASE", timetoTD)
	}

	// PHASE 1
	TD_id,err := TD_CreateAndRegister(&MyClient)
	if err != nil {
		fmt.Printf("[INIT Clients] ERROR: %s \n ", err)
		panic(err)
		return
	}

	MyClient.TDId = TD_id

	sensorsInstancesNumber := configs.Client_maxSensorInstance// or rand.Intn(configs.Client_maxSensorInstance)
	if LOG{
		utils.LogConfig_InstanceNumber(MyClient.ID, sensorsInstancesNumber)
	}

	var sub_wg sync.WaitGroup
	for s := 0; s < sensorsInstancesNumber; s++ {

		// Creation of Sensor Node Instances goroutine
		timeToInstance := rand.Intn(configs.Client_maxSecondPhaseTime - (timetoTD + timetoWake))
		if LOG{
			utils.LogConfig_RndTime(MyClient.ID, "TIME_TO_SECOND_PHASE_"+strconv.Itoa(s), timeToInstance)
		}

		commPeriod := rand.Intn( configs.Client_maxCommunicationPeriodRange[1] - configs.Client_maxCommunicationPeriodRange[0]) + configs.Client_maxCommunicationPeriodRange[0]
		if LOG{
			utils.LogConfig_RndTime(MyClient.ID, "TIME_THIRD_PHASE_PERIOD_"+strconv.Itoa(s), commPeriod)
		}

		sub_wg.Add(1)
		go sensorInstanceSubRoutine(&sub_wg, MyClient, commPeriod, timeToInstance, maxTime, startTime)
	}
	sub_wg.Wait()

	return
}

func sensorInstanceSubRoutine( sub_wg *sync.WaitGroup,client SEAMDAPClient, commPeriod int, sleepTime int, maxTime int, startTime time.Time ){
	defer sub_wg.Done()

	// Third Inactivity time
	time.Sleep(time.Duration(sleepTime)*time.Second)

	// PHASE 2
	instID, err := INSTANCE_CreateAndRegister(&client)
	if err != nil{
		fmt.Printf("[client %d] ERROR: %s \n", client.ID, err.Error())
		panic(err)
	}


	for{
		//PHASE 3

		remainingSeconds := startTime.Add(time.Duration(maxTime)*time.Second).Sub(time.Now()).Seconds()
		if remainingSeconds < float64(2*commPeriod){
			break
		}
		//TEST
		//fmt.Printf("[client %d] REMAINING SEC: %f, 2*CommPeriod: %f, check : %t \n", client.ID, remainingSeconds,float64(2*commPeriod), remainingSeconds < float64(2*commPeriod) )

		stat, err := SAMPLE_CreateAndUpload(&client, instID)
		if err != nil{
			fmt.Printf("[client %d] ERRORE: %s \n", client.ID, err.Error())
		}
		if VERBOSE{
			fmt.Printf("[client %d] Sample Upload Complete. Response: %s \n", client.ID, stat)
		}

		// Inactive for samples upload period
		time.Sleep(time.Duration(commPeriod)*time.Second)
	}
	fmt.Printf("[client %d] Closing.\n", client.ID)
	return

}

func TD_CreateAndRegister(client *SEAMDAPClient) (uuid.UUID,error){
	if VERBOSE{
		fmt.Printf("[client %d] PHASE 1: Registering TD. \n", client.ID)
	}

	TD := utils.TestMessagesTD[client.MSG_index]
	TD.ID = client.TDId.String()
	TD.Title ="TD_TITLE_EXAMPLE_" + strconv.Itoa(client.Index)
	TD.Model = "TD_MODEL_EXAMPLE_" + strconv.Itoa(client.Index)
	TD.Description = "TD_DESC_EXAMPLE_" + strconv.Itoa(client.Index)

	resp,err, timeReq, timeResp := InterfaceRegistration(TD, 0, time.Now(), client)
	if err != nil{
		fmt.Printf("[client %d] ERROR: %s \n", client.ID, err.Error())
		return uuid.New(), err
	}
	if resp.StatusCode != 200{
		fmt.Printf("[client %d] ERROR: Failed request - status code != 200. Desc: %s \n",  client.ID, err.Error())
		return uuid.New(), err
	}

	ns := utils.NewInterfaceResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ns)
	if err != nil {
		fmt.Printf("[client %d] Failed parsing NewInterfaceResponse: %s \n", client.ID, err.Error())
		return uuid.New(), err
	}

	if VERBOSE {
		fmt.Printf("[client %d] Received response: %v \n", client.ID, ns)
	}

	if LOG{
		utils.LogClientSendTD(client.ID, timeReq.Sub(client.StartTime), ns.UID.String())
		utils.LogClientSendTD_response(client.ID, timeResp.Sub(client.StartTime), ns.UID.String())
	}

	return ns.UID, nil
}

func INSTANCE_CreateAndRegister(client *SEAMDAPClient) (int,error){
	if VERBOSE{
		fmt.Printf("[client %d] PHASE 2: Registering Instance.\n", client.ID)
	}

	instance_request := utils.InstanceRegistrationRequest{
		TDID:      client.TDId,
		UserID:    client.ID,
		BoardName: "BOARD_NAME_"+strconv.Itoa(client.Index),
		Area:      utils.NewGeojsonFeature(),
	}

	resp, err, timeReq, timeResp := InstanceRegistration(instance_request)
	if err != nil{
		fmt.Printf("[client %d] ERROR: %s \n", client.ID, err.Error())
		return 0, err
	}
	if resp.StatusCode != 200{
		fmt.Printf("[client %d] ERROR: Failed request - status code != 200. Desc: %s \n", client.ID, err.Error())
		return 0, err
	}

	ns := utils.InstanceRegistrationResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ns)
	if err != nil {
		fmt.Printf("[client %d] ERROR: Failed parsing InstanceRegistrationResponse. Desc: %s \n", client.ID, err.Error())
		return 0, err
	}
	if VERBOSE{
		fmt.Printf("[client %d] Received response: %v \n", client.ID, ns)
	}

	if LOG{
		utils.LogClientSendInstance(client.ID, timeReq.Sub(client.StartTime), strconv.Itoa(ns.InstanceID))
		utils.LogClientSendInstance_response(client.ID, timeResp.Sub(client.StartTime), strconv.Itoa(ns.InstanceID))
	}

	return ns.InstanceID, nil
}

func SAMPLE_CreateAndUpload(client *SEAMDAPClient, instanceID int) (string,error){

	if VERBOSE{
		fmt.Printf("[client %d] PHASE 3: Uploading sample.\n", client.ID)
	}

	rec,err := utils.GetSENML(client.MSG_index)
	if err != nil {
		fmt.Printf("[client %d] Error in generating Senml: %s \n", client.ID, err.Error())
	}
	rec.TimeRecord = time.Now().Format("2006.01.02T15:04:05")
	rec.Name = strconv.Itoa(instanceID)

	msg := utils.Custom{
		Record: []utils.SenMLPos{rec},
	}

	resp, err, timeReq, timeResp := UploadSampling(msg, client.TDId)
	if err != nil{
		fmt.Printf("[client %d] ERROR: %s \n ",client.ID, err.Error())
		return "", err
	}
	if resp.StatusCode != 200{
		fmt.Printf("[client %d] ERROR: Failed request - status code != 200. Desc: %s \n", client.ID, err.Error())
		return "", err
	}

	ns := utils.SamplingResponse{}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ns)
	if err != nil {
		fmt.Printf("[client %d] Failed parsing InstanceRegistrationResponse: %s \n", client.ID, err.Error())
		return "", err
	}

	if VERBOSE{
		fmt.Printf("[client %d] Received response: %v \n", client.ID, ns)
	}

	if LOG{
		temporaryid:=rand.Intn(1000000)
		utils.LogClientSendSample(client.ID, timeReq.Sub(client.StartTime), strconv.Itoa(temporaryid))
		utils.LogClientSendSample_response(client.ID, timeResp.Sub(client.StartTime), strconv.Itoa(temporaryid))
	}

	return ns.Status, nil
}