package paxos

import "log"

type learner struct {
	id           int
	acceptedMsgs map[int]message
	nt           nodeNetwork
}

