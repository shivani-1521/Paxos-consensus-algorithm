package paxos

import "log"

func NewAcceptor(id int, nt nodeNetwork, learners ...int) acceptor {
	newAcceptr := acceptor{id: id, nt: nt}
	newAcceptr.learners = learners
	return newAcceptr
}

type acceptor struct {
	id         int
	learners   []int
	acceptMsg  message
	promiseMsg message
	nt         nodeNetwork
}

