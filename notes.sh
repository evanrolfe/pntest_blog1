#===============================================================================
# create the network
#===============================================================================
docker network create -d bridge evan --subnet 10.0.1.1/24 --gateway 10.0.1.1

https://forums.docker.com/t/setting-default-gateway-to-a-container/17420/4
https://unix.stackexchange.com/questions/615255/docker-container-as-network-gateway-not-responding
https://maxammann.org/posts/2020/04/routing-docker-container-over-vpn/

#===============================================================================
# context
#===============================================================================
$ netstat -rn
Kernel IP routing table
Destination     Gateway         Genmask         Flags   MSS Window  irtt Iface
0.0.0.0         192.168.1.1     0.0.0.0         UG        0 0          0 wlp2s0
10.0.1.0        0.0.0.0         255.255.255.0   U         0 0          0 br-e985b298b429

$ ip route
default via 192.168.1.1 dev wlp2s0 proto dhcp metric 600
10.0.1.0/24 dev br-e985b298b429 proto kernel scope link src 10.0.1.1 linkdown
192.168.1.0/24 dev wlp2s0 proto kernel scope link src 192.168.1.101 metric 600

#===============================================================================
# start the app + proxy:
#===============================================================================
docker build -t my_server .
docker build -t proxy .

docker run -d --privileged -p 8092:8092 --network evan --ip 10.0.1.2 -i -t proxy:latest
docker run -d --hostname three -p 8090:8090 --network evan --ip 10.0.1.3 -i -t my_server:latest

docker run -d --hostname four -p 8010:8090 --network evan --ip 10.0.1.4 -i -t my_server:latest

#===============================================================================
# configure the routing policiy on the host machine
#===============================================================================
sudo ip rule add from 10.0.1.0/24 table 200 # the network you like to forward over the gateway container
sudo ip route add default via 10.0.1.2 table 200


sudo ip rule add from 10.0.1.2 table 199 # the network you like to forward over the gateway container
sudo ip route add default via 10.0.1.1 table 199

ip route show table 200
sudo ip route flush table 200

# try: (does not route 1.3 through 1.2)
sudo ip route add 10.0.1.3 via 10.0.1.2 table 200
# ip route add default via 172.20.0.50

#===============================================================================
# configure the routing policiy on the container (THIS WORKS!)
#===============================================================================

# On the app container:

# route external traffic:
docker exec --privileged -it 3e78a70fa578 bash
ip route del default
ip route add default via 10.0.1.2

# route container2container traffic:
ip rule add from 10.0.1.0/24 table 200
ip route add default via 10.0.1.2 table 200

#===============================================================================
# TODO
#===============================================================================
- create a rule which keeps 10.0.1.2 (proxy) using 10.0.1.1 as the gateway
- BUT forwards 10.0.1.3 to 10.0.1.2

see:
https://www.linuxserver.io/blog/routing-docker-host-and-container-traffic-through-wireguard

- show table numbers: cat /etc/iproute2/rt_tables
