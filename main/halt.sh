#!/bin/bash
docker container stop $(docker container ls -aq)
docker container rm $(docker container ls -aq)
docker network rm d7024e_kademlia_network