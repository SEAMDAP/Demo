package configs

// This file contains the different configuration parameters to customize each Demo run.

// SERVER
var Server_addr = "127.0.0.1" // replace with the real server location
var Server_port = "8000"
var Server_URL_firstPhasePath = "api/sensor/interface"
var Server_URL_secondPhasePath = "api/sensor/instance"
var Server_URL_thirdPhasePath = "api/sensor/data"

// CLIENT
var Client_number = 10 //30
var Client_maxSensorInstance = 30 //10
var Client_maxLifeTime = 18 // max time of the entire client pool simulation
// These times values indicates the deadline to be respected in the completition of the related actions. A random value
// is chosen (limited by the deadline)
var Client_maxWakeTime = (Client_maxLifeTime / 10)  // max time in seconds to wait before wake
var Client_maxFirstPhaseTime = (Client_maxLifeTime / 8)  // max time in seconds to wait before start phase 1
var Client_maxSecondPhaseTime= (Client_maxLifeTime / 5)  // max time in seconds to wait before start phase 2
var Client_maxCommunicationPeriodRange = []int{ 15, 30} // min and max time in seconds to wait after each repetition of phase 3


// REDIS SERVER
var Server_REDIS_fullAddress = "localhost:6379"
var Server_REDIS_pass = ""

// LOGGING
var LogFileAddress = "./" //replace with the chosen path

