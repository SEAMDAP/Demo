package configs

var Server_addr = ""
var Server_port = ""
var Server_URL_phirstPhasePath = "api/sensor/interface"
var Server_URL_secondPhasePath = "api/sensor/instance"
var Server_URL_thirdPhasePath = "api/sensor/data"

var Server_REDIS_fullAddress = "localhost:6379"
var Server_REDIS_pass = ""



//The recommended approach is to use a network scan tool like Wireshark
var LogFileAddress = ""

var Client_number = 10
var Client_maxSensorInstance = 15

var Client_maxLifeTime = 1 * 600 //seconds//1 * 3600 //seconds
var Client_maxWakeTime = (Client_maxLifeTime / 10) // seconds
var Client_maxPhirstPhaseTime = (Client_maxLifeTime / 5)
var Client_maxSecondPhaseTime= Client_maxLifeTime / 4

//var Client_maxSecondPhaseTime_relative = (Client_maxLifeTime - (Client_maxWakeTime+Client_Client_maxPhirstPhaseTime_relative)) / (Client_maxSensorInstance * 2)
var Client_maxCommunicationPeriodRange = []int{ 60, 600} // 1 to 10 minutes

