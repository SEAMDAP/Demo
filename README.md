# SEAMDAP Protocol Demo

This is the repository of the SEAMDAP Protocol.

## Brief Description of the Protocol
The SEAMDAP protocol is described in the paper "Seamless Sensor Data Acquisition for the Edge-to-Cloud Continuum".
The paper was submitted for ACM PODC 2022: The 41st ACM Symposium on Principles of Distributed Computing (PODC 2022),
and is now under review. \

Read the paper before continuing: _link-and-info-to-be-added_ (Info are hidden due to **double blind** review policy).

## Description of the Demo

The demo is composed by two functional parts: the Client and the Server part.\
\
The **Server** is a SEAMDAP HTTP-based RESTful seamdap_server, capable of interpreting the SEAMDAP messages, for all the phases.
He need the access to a Redis-Server to store received data.

\
The **Client** is capable of emulating a dynamic set of SEAMDAP clients (one for each sensors node instance to be registered)
The Client part is executed as follows. First, N goroutines are created, with N equal to the number of instances,
which are in charge of obtaining, personalizing and uploading to the seamdap_server a file adhering to the Thing Description format. Each goroutine suddenly creates M sub-goroutines, with M equal to the number of instances that will be
registered using the related TD interface. The instance's sub-goroutine also takes care of loading its samples
on a regular basis. During the execution, spread along the various phases, some inactivity times and periodicities are randomly chosen
to offset the various activities, respecting some min and max values.

### Repository Organization
\
The code is unique to both parties. It is necessary to to copy it into the two selected nodes to host the relative entities,
and then start them separately, after having properly set the configuration file.\
The repo is mainly organized as follow:
- _init.go_: the main file where the Demo starts
- _/seamdap_server_: directory containing the Server code.
  - _seamdap_server.go_ contains the main cycle, routing and functionality of the server.
  - _seamdap_handler.go_ contains the handler function associated to each API.
- _/seamdap_client_: directory containing the Client code.
  - _seamdap_client.go_ contains the main logic of the emulated clients.
  - _seamdap_requests.go_ contains the code for requests generations.
- _/configs/settings.go_: file containing the settings of the Demo. Modify this file as needed.
- _/utils_: directory containing the common code used by both Server and Client.
  - _messages.go_ contains the pre-formatted message prototypes sended by the client in phase 1 and 3.
  - _models.go_ contains the models, in the form of structs, of the main entities that are used in the Demo, and utility
    functions used to deal with them.
  - _senml_seamdap.go_ contains Go struct of a SenML message (SEAMDAP version).   
  - _simplelog.go_ contains the code for Logging operations.
  - _thingdescription.go_ contains Go simple struct of a Thing Description (TD) message .

### How to Start the Demo
\
From the init.go file it is possible to run both the component.

To start the Server node: 
```bash
go run init.go server
```
  - After the startup, the seamdap_server accepts command-line read instructions. the accepted commands are:
    - _start_: starts the seamdap_server
    - _stop_: shut down the seamdap_server
    - _help_: print help text
    - _status_: print status of the seamdap_server
    - _exit_: stops the seamdap_server and close the program

To start the Clients node:    
```bash
go run init.go client
```

Go 1.16 version or above is recommended.
