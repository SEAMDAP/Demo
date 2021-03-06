package utils

import (
	"fmt"
	"github.com/SEAMDAP/Demo/configs"
	"os"
	"strconv"
	"time"
)

// LOGGING functions
// NB: Use only as a referende. The recommended approach is to use a network scan tool like Wireshark fo evaluations.

var LOGFILE *os.File
var LOGFILE_CONFIGS *os.File


func LogConfig_Parameters(){

	l := map[string]string{
		"Server_addr" :                        configs.Server_addr,
		"Server_port" :                        configs.Server_port,
		"Server_URL_firstPhasePath" :          configs.Server_URL_firstPhasePath,
		"Server_URL_secondPhasePath" :         configs.Server_URL_secondPhasePath,
		"Server_URL_thirdPhasePath" :          configs.Server_URL_thirdPhasePath,
		"Server_REDIS_fullAddress" :           configs.Server_REDIS_fullAddress,
		"Server_REDIS_pass" :                  configs.Server_REDIS_pass,
		"LogFileAddress" :                     configs.LogFileAddress,
		"Client_number" :                      strconv.Itoa(configs.Client_number),
		"Client_maxSensorInstance" :           strconv.Itoa(configs.Client_maxSensorInstance),
		"Client_maxLifeTime" :                 strconv.Itoa(configs.Client_maxLifeTime),
		"Client_maxWakeTime" :                 strconv.Itoa(configs.Client_maxWakeTime),
		"Client_maxFirstPhaseTime" :           strconv.Itoa(configs.Client_maxFirstPhaseTime),
		"Client_maxSecondPhaseTime" :          strconv.Itoa(configs.Client_maxSecondPhaseTime),
		"Client_maxCommunicationPeriodRange" : fmt.Sprintf("[%d - %d]", configs.Client_maxCommunicationPeriodRange[0], configs.Client_maxCommunicationPeriodRange[1]),
	}

	for k,v := range l {
		log:= fmt.Sprintf("APPLICATION,%s,%s\n", k, v)
		if _, err := LOGFILE_CONFIGS.Write([]byte(log)); err != nil {
			fmt.Println(err)
		}
	}
}

func LogConfig_InstanceNumber(id int, ins_number int){
	log:= fmt.Sprintf("CLIENT_%d,%s,%d\n", id, "GENERATED_INSTANCE", ins_number)
	if _, err := LOGFILE_CONFIGS.Write([]byte(log)); err != nil {
		fmt.Println(err)
	}
}

func LogConfig_RndTime(id int, timeType string, tm int){
	log:= fmt.Sprintf("CLIENT_%d,%s,%d\n", id, timeType, tm)
	if _, err := LOGFILE_CONFIGS.Write([]byte(log)); err != nil {
		fmt.Println(err)
	}
}


func LogClientWake(id int, tm time.Duration){
	log:= fmt.Sprintf("CLIENT_%d,%s,%s, %f,%s\n", id, "WAKE", "nil", tm.Seconds(), "nil")
	if _, err := LOGFILE.Write([]byte(log)); err != nil {
		fmt.Println(err)
	}
}

func LogClientSendTD(id int, tm time.Duration, TDID string){
	log:= fmt.Sprintf("CLIENT_%d,%s,%s,%f,TDID=%s\n", id, "SEND", "TD", tm.Seconds(), TDID)
	if _, err := LOGFILE.Write([]byte(log)); err != nil {
		fmt.Println(err)
	}
}

func LogClientSendTD_response(id int, tm time.Duration, TDID string){
	log:= fmt.Sprintf("CLIENT_%d,%s,%s,%f,TDID=%s\n", id, "RESPONSE","TD", tm.Seconds(), TDID)
	if _, err := LOGFILE.Write([]byte(log)); err != nil {
		fmt.Println(err)
	}
}

func LogClientSendInstance(id int, tm time.Duration, INID string){
	log:= fmt.Sprintf("CLIENT_%d,%s,%s,%f,INSTANCE_ID=%s\n", id, "SEND","INSTANCE", tm.Seconds(), INID)
	if _, err := LOGFILE.Write([]byte(log)); err != nil {
		fmt.Println(err)
	}
}

func LogClientSendInstance_response(id int, tm time.Duration, INID string){
	log:= fmt.Sprintf("CLIENT_%d,%s,%s,%f,INSTANCE_ID=%s\n", id, "RESPONSE","INSTANCE", tm.Seconds(), INID)
	if _, err := LOGFILE.Write([]byte(log)); err != nil {
		fmt.Println(err)
	}
}
func LogClientSendSample(id int, tm time.Duration, INID string){
	log:= fmt.Sprintf("CLIENT_%d,%s,%s,%f,SAMPLE_ID=%s\n", id, "SEND", "SAMPLE", tm.Seconds(), INID)
	if _, err := LOGFILE.Write([]byte(log)); err != nil {
		fmt.Println(err)
	}
}

func LogClientSendSample_response(id int, tm time.Duration, INID string){
	log:= fmt.Sprintf("CLIENT_%d,%s,%s,%f,SAMPLE_ID=%s\n", id, "RESPONSE", "SAMPLE", tm.Seconds(), INID)
	if _, err := LOGFILE.Write([]byte(log)); err != nil {
		fmt.Println(err)
	}
}


func init() {
	timeF := time.Now().Format("2006-01-02_15:04:05")
	logFile := fmt.Sprintf("%s/application_client_%s.log", configs.LogFileAddress, timeF)
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil{
		panic(err)
		return
	}

	logFile2 := fmt.Sprintf("%s/application_configs_%s.log", configs.LogFileAddress, timeF)
	file2, err := os.OpenFile(logFile2, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil{
		panic(err)
		return
	}

	LOGFILE = file
	LOGFILE_CONFIGS = file2

	// Columns name
	if _, err := LOGFILE.Write([]byte("CLIENT, ACTION, TYPE, TIME, INFO \n")); err != nil {
		panic(err)
	}
	if _, err := LOGFILE_CONFIGS.Write([]byte("TYPE, PARAMETER, VALUE \n")); err != nil {
		panic(err)
	}

	return

}


