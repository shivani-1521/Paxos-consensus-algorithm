package paxos

import "log"

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

func(n *network)sendTo(from int, m message){
	log.Println("send message from ", m.from, " send to ", m.to, " value ", m.value)
	n.recevQueue[m.to] <- m
}

func(n *network)recevFrom(id int) message{
	retMsg := <-n.recevQueue[id]
	log.Println("Recev msg from ", retMsg.from, " send to ", retMsg.to, " value ", retMsg.value)
	return retMsg
}

type nodeNetwork struct{
	id int
	netw *network
}

func (n *nodeNetwork)send(m message){
	n.netw.sendTo(m.to, m)
}

func (n *nodeNetwork)recev()message{
	return n.netw.recevFrom(n.id)
}