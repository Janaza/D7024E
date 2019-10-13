#!/bin/bash
docker build . -t kadlab:latest
docker container stop $(docker container ls -aq)
docker container rm $(docker container ls -aq)
docker network rm d7024e_kademlia_network
docker network create --ipam-driver=default --subnet=10.0.0.0/24 --gateway=10.0.0.1 d7024e_kademlia_network  

docker run --net d7024e_kademlia_network --ip 10.0.0.2 --volume /home/d7024e-user/go/src/D7024E/:/go/src/D7024E --volume  /home/d7024e-user/go/src/D7024E/main/init/bootstrap/:/init --name bootstrap_node -t -d kadlab
port=6000
for i in $(seq 1 50)
do
    port=$((port+1))
    ip="10.0.0.$((i+2))"
    docker run --net d7024e_kademlia_network --ip "$ip" --volume /home/d7024e-user/go/src/D7024E/:/go/src/D7024E --volume /home/d7024e-user/go/src/D7024E/main/init/node/:/init --name "node_$i" -t -d kadlab "$port"
done

