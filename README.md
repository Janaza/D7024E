# D7024E
A distrubuted systems using docker containers and the Kademila algorithm.

## Building the system

#### Prerequisite
- Docker
- Go

#### Clone git

Start by cloning the git to your own workspace:

```bash
git clone https://github.com/Janaza/D7024E.git
```

#### Building main

Change directory to main:

```bash
cd D7024E\main
```

Run the following command to build the main:

```bash
go build .\main.go
```

#### Run main

You are now ready to run the main using the following command:

```bash
.\main.exe -port *port* -bootstrapIP *IP* (optional)
```

where <-port> specifies which port to use and <-bootstrapIP> specifies which ip to connect to, just starts listener if ommited.


#### Docker script

However, the main itself is not an effective way to spin up many nodes so you 
can use a docker script that provides you with 50 nodes all part of the same Kademlia network.

Start cluster with ```main/run.sh``` (OBS terminates any running containers!)
``` chmod +x run.sh ``` if necessary

## Command Line Interface

The system provides a Command Line Interface (CLI) in order to save values, retrieve values, 
terminate a node and retrieve the k closest contacts of from the node.

#### Put
The CLI for saving objects to the system.
```bash
PUT <value>
```
PUT takes a single argument <value>, that is the content of the value you are uploading. After 
Successful save it outputs the hash of that object.

#### Get
The CLI for receiving an object that is stored within the Kademlia network.
```bash
GET <hash>
```  
GET takes a single argument <hash>, if the hash is stored somewhere in the system you get
the value returned, else you receive a list of the k closest contacts.

#### Exit
CLI for terminating the node.

```bash
EXIT
```
Terminates the node after three seconds.

#### Contacts
CLI for retrieving the k closest contacts from a node.
```bash
CONTACTS
```
