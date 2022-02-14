#!/usr/bin/env python3.6
# -*- coding: utf-8 -*-

# This script could be used to analyze the a simulation CSV results.
# It works with the SEAMDAP Demo log file.
# NB: is recommended to do not use log file information to perform evaluation,
#     instead use some network analysis tool as Wireshark, and then extract and parse his CSV results.

import numpy as np
import matplotlib
# matplotlib.use('agg')
import matplotlib.pyplot as plt
import csv
import math
from scipy.interpolate import make_interp_spline, BSpline

######################## READ THE LOG FILES ########################
####################################################################

# Pre-calculated
TD_mean = (369 + 454 + 309 +822 + 916) / 5
Custom_mean = (520 + 371 + 547 + 274 + 456) / 5
Senml_mean = (67 + 151 + 82 + 139 + 93) / 5

# file path
applicationLogfilePath = "/path/to/file_application_client.log"
parametersLogfilePath = "/path/to/file_application_configs.log"


row_send, row_response, row_wake = [],[],[]
row_application_params,row_client_time_wake ,row_client_time_istances ,row_client_time_1phase ,row_client_time_2phase ,row_client_time_3phasePeriod = [],[],[],[],[],[]

file1_lc, file2_lc = 0,0
client_maxLifeTime, number_TD, number_instance_per_client = 0,0,0


with open(applicationLogfilePath) as csv_file:

    csv_reader = csv.reader(csv_file, delimiter=',')
    line_count = 0
    for row in csv_reader:
        if line_count == 0:
            print("Columns names: ", ",".join(row))
            line_count += 1
        else:
            line_count += 1

            if row[1] == "SEND":
                row_send.append(row)
            elif row[1] == "RESPONSE":
                row_response.append(row)
            elif row[1] == "WAKE":
                row_wake.append(row)
    print('Processed {} lines.'.format(line_count))
    file1_lc  = line_count

with open(parametersLogfilePath) as csv_file:

    csv_reader = csv.reader(csv_file, delimiter=',')
    line_count = 0
    for row in csv_reader:
        if line_count == 0:
            print("Columns names: ", ",".join(row))
            line_count += 1
        else:
            line_count += 1
            
            if row[0] == "APPLICATION":
                if row[1] == "Client_maxLifeTime":
                    client_maxLifeTime = int(row[2]) 
                elif row[1] == "Client_maxLifeTime":
                    client_maxLifeTime = int(row[2]) 
                elif row[1] == "Client_number":
                    number_TD = int(row[2]) 
                row_application_params.append(row)
            elif row[1] == "TIME_TO_WAKE":
                row_client_time_wake.append(row)
            elif row[1] == "GENERATED_INSTANCE": 
                row_client_time_istances.append(row)
            elif row[1] == "TIME_TO_FIRST_PHASE":
                row_client_time_1phase.append(row)
            elif row[1][0:20] == "TIME_TO_SECOND_PHASE":
                row_client_time_2phase.append(row)
            elif row[1][0:23] == "TIME_THIRD_PHASE_PERIOD":
                row_client_time_3phasePeriod.append(row)
            
    print('Processed {} lines.'.format(line_count))
    file2_lc = line_count

print("FILE1: line count = {}, rows saved = {}".format(file1_lc, len(row_send) + len(row_response) + len(row_wake)))
print("FILE2: line count = {}, rows saved = {}".format( file2_lc,  len(row_application_params) + len(row_client_time_wake) + len(row_client_time_istances) + len(row_client_time_1phase) + len(row_client_time_2phase) + len(row_client_time_3phasePeriod) ))


############################# GRAPHICS #############################
####################################################################

factor = 30
# MESSAGE EXCHANGED
lista = [0 for i in range(0, int(client_maxLifeTime/factor))]
for r in row_send:
    lista[math.floor(float(r[3])/factor)] += 1


x = [ i*factor for i in range(0, int(client_maxLifeTime/factor))] 
y = lista

plt.title("Messages") 
plt.xlabel("Time [s]") 
plt.ylabel("Number of Messages") 
plt.plot(x,y) 
plt.show()


# BYTE EXCHANGED
lista2 = [0 for i in range(0, int(client_maxLifeTime/factor))]
lista2_p1 = [0 for i in range(0, int(client_maxLifeTime/factor))]
lista2_p2 = [0 for i in range(0, int(client_maxLifeTime/factor))]
lista2_p3 = [0 for i in range(0, int(client_maxLifeTime/factor))]
for r in row_send:
    if r[2] == "SAMPLE":
        lista2[math.floor(float(r[3])/factor)] += Senml_mean / 100
        lista2_p3[math.floor(float(r[3])/factor)] += Senml_mean / 100
    elif r[2] == "TD":
        lista2[math.floor(float(r[3])/factor)] += TD_mean / 100
        lista2_p1[math.floor(float(r[3])/factor)] += TD_mean / 100
    elif r[2] == "INSTANCE":
        lista2[math.floor(float(r[3])/factor)] += Custom_mean / 100
        lista2_p2[math.floor(float(r[3])/factor)] += Custom_mean / 100
    else:
        print("ERROR: ", r[2])

x2 = [ i*factor for i in range(0, int(client_maxLifeTime/factor))] 
y2 = lista2

plt.title("Network Traffic per Phase") 
plt.xlabel("Time [s]") 
plt.ylabel("Traffic [B]")
# plt.plot(x2,y2)
plt.plot(x2,lista2_p1, label="Phase 1")
plt.plot(x2,lista2_p2, label="Phase 2")
plt.plot(x2,lista2_p3, label="Phase 3")
plt.legend() 
plt.show()

exit(0)

######################## PRINT RESULTS DATA ########################
####################################################################

TD_totalAmount, Custom_totalAmount, Senml_totalAmount = 0,0,0
TD_totalAmount_size, Custom_totalAmount_size, Senml_totalAmount_size = 0,0,0

summedResponseTime, validResponseTime = 0,0
summedResponseTime_p1, validResponseTime_p1 = 0,0
summedResponseTime_p2, validResponseTime_p2 = 0,0
summedResponseTime_p3, validResponseTime_p3 = 0,0

# Some outlier msg exchange could be excluded from the results computing
outlier = 0
outlier_latency_threshold = 10 # seconds

# NB: this computation could be really slow due to the double for loop, with exponential times O(n^2)
for r in row_send:
    if r[2] == "SAMPLE":
        Senml_totalAmount +=1
        Senml_totalAmount_size += TD_mean
    elif r[2] == "TD":
        TD_totalAmount +=1
        TD_totalAmount_size += TD_mean
    elif r[2] == "INSTANCE":
        Custom_totalAmount +=1
        Custom_totalAmount_size += TD_mean
    else:
        print("ERROR: ", r[2])
    
    t1, t2 = 0,0

    # This piece slow down the computation
    # Execution time could be improved up with different techniques like sorting the array using the client id and
    # associating the values by their index.
    for r2 in row_response:
        if r2[4] == r[4]:
            t2 = float(r2[3])
    t1 = float(r[3])
    if t2 > 0:
        # to esclude outlier
        if t2-t1 < outlier_latency_threshold:
            summedResponseTime += t2-t1
            validResponseTime += 1
            if r[2] == "SAMPLE":
                summedResponseTime_p3 += t2-t1
                validResponseTime_p3 += 1
            elif r[2] == "TD":
                summedResponseTime_p1 += t2-t1
                validResponseTime_p1 += 1
            elif r[2] == "INSTANCE":
                summedResponseTime_p2 += t2-t1
                validResponseTime_p2 += 1
        else:
            print("Outlier [{}] Type: {}, Info: {}".format(t2-t1,r[2],r[4]))
            outlier += 1
    else:
        print("ERROR: cannot find ", r[4])

print("Latency")
print("Total: {} s, considering {} messages, excluding {} outliers".format((summedResponseTime / validResponseTime), validResponseTime, outlier))
print("P1: {}, P2: {}, P3:{}".format((summedResponseTime_p1 / validResponseTime_p1),(summedResponseTime_p2 / validResponseTime_p2),(summedResponseTime_p3 / validResponseTime_p3)))
print("Messages exchanged")
print("P1: {}, P2: {}, P3:{}".format(TD_totalAmount, Custom_totalAmount, Senml_totalAmount))
print("Byte exchanged")
print("P1: {}, P2: {}, P3:{}".format(TD_totalAmount_size, Custom_totalAmount_size, Senml_totalAmount_size))
