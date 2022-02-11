package configs

var Server_addr = ""
var Server_port = "8000"
var Server_URL_firstPhasePath = "api/sensor/interface"
var Server_URL_secondPhasePath = "api/sensor/instance"
var Server_URL_thirdPhasePath = "api/sensor/data"

var Server_REDIS_fullAddress = "localhost:6379"
var Server_REDIS_pass = ""



var LogFileAddress = ""

var Client_number = 15
var Client_maxSensorInstance = 20

var Client_maxLifeTime = 1 * 3600 //seconds
var Client_maxWakeTime = (Client_maxLifeTime / 10) // seconds
var Client_maxFirstPhaseTime = (Client_maxLifeTime / 5)
var Client_maxSecondPhaseTime= Client_maxLifeTime / 4

var Client_maxCommunicationPeriodRange = []int{ 60, 600} // 1 to 10 minutes

