package paxos

import (
	"log"
	"time"
)

func createNetwork(nodes ...int)*network{
	net := network{recevQueue: make(map[int]chan message, 0)}

	for _, node := range nodes{
		net.recevQueue[node] = make(chan message, 1024)
	}

	return &net
}

type network struct{
	recevQueue map[int]chan message
}

func(n *network)sendTo(m message){
	log.Println("send message from ", m.from, " send to ", m.to, " value ", m.value, "typ:", m.typ)
	n.recevQueue[m.to] <- m
}

func(n *network)recevFrom(id int) *message{
	select{
	case retMsg := <-n.recevQueue[id]:
		log.Println("Receive message from ", retMsg.from, " send to ", retMsg.to, " val ", retMsg.val, "typ:", retMsg.typ)
		return &retMsg
		case <-time.After(time.Second):
			log.Println("id:", id, " no message, time out")
			return nil
	}
}

type nodeNetwork struct{
	id int
	netw *network
}

func (n *nodeNetwork)send(m message){
	n.netw.sendTo(m)
}

func (n *nodeNetwork)recev() *message{
	return n.netw.recevFrom(n.id)
}

func (n *network)getNodeNetwork (id int) nodeNetwork{
	return nodeNetwork{id : id, net: n}
}