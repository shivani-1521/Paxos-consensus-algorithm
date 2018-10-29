package paxos

import "log"

func NewProposer(id int, val string, nt nodeNetwork, accetors ...int) *proposer {

	prop := proposer{id: id, proposeVal: val, sequence: 0, nt: nt}
	prop.acceptors = make(map[int]message, len(accetors))
	log.Println("proposer has ", len(accetors), " acceptors, val:", prop.proposeVal)
	for _, acceptor := range accetors {
		prop.acceptors[acceptor] = message{}
	}
	return &pro
}

type proposer struct {
	id         int
	sequence        int
	proposeNum int
	proposeVal string
	acceptors  map[int]message
	nt         nodeNetwork
}

func (p *proposer) run() {

	log.Println("Proposer running value ", p.proposeVal)

	for !p.majorityReached() {

		log.Println("[Proposer:Prepare]")

		outMsgs := p.prepare()
		log.Println("[Proposer: prepare ", len(outMsgs), "msg")

		for _, msg := range outMsgs {
			p.nt.send(msg)
			log.Println("Proposer send", msg)
		}

		log.Println("[Proposer: prepare recev...")
		m := p.nt.recev()
		if m == nil {
			log.Println("Proposer: no msg... ")
			continue
		}
		log.Println("[Proposer received ", m)
		switch m.typ {
		case Promise:
			log.Println("Proposer received a promise from ", m.from)
			p.checkRecvPromise(*m)
		default:
			panic("message not supported")
		}
	}

	log.Println("[Proposer:Propose]")
	log.Println("Proposor propose sequence", p.getProposeNum(), " value:", p.proposeVal)
	proposeMsgs := p.propose()
	for _, msg := range proposeMsgs {
		p.nt.send(msg)
	}
}

func (p *proposer) propose() []message {

	sendMsgCount := 0
	var msgList []message

	log.Println("proposer: propose msg:", len(p.acceptors))

	for acepId, acepMsg := range p.acceptors {

		log.Println("Check promise id:", acepMsg.getProposeSeq(), p.getProposeNum())

		if acepMsg.getProposeSeq() == p.getProposeNum() {

			msg := message{from: p.id, to: acepId, typ: Propose, sequence p.getProposeNum()}
			msg.val = p.proposeVal
			log.Println("Propose val:", msg.val)
			msgList = append(msgList, msg)
		}
		sendMsgCount++
		if sendMsgCount > p.majority() {
			break
		}
	}
	log.Println(" proposer proposed msg list:", msgList)
	return msgList
}

func (p *proposer) prepare() []message {
	p.sequence++

	sendMsgCount := 0
	var msgList []message

	log.Println("proposer: prepare major msg:", len(p.acceptors))

	for acepId, _ := range p.acceptors {
		msg := message{from: p.id, to: acepId, typ: Prepare, sequence p.getProposeNum(), val: p.proposeVal}
		msgList = append(msgList, msg)
		sendMsgCount++
		if sendMsgCount > p.majority() {
			break
		}
	}
	return msgList
}

func (p *proposer) checkRecvPromise(promise message) {

	previousPromise := p.acceptors[promise.from]

	log.Println(" prevMsg:", previousPromise, " promiseMsg:", promise)
	log.Println(previousPromise.getProposeSeq(), promise.getProposeSeq())

	if previousPromise.getProposeSeq() < promise.getProposeSeq() {

		log.Println("Proposor:", p.id, " get new promise:", promise)
		p.acceptors[promise.from] = promise

		if promise.getProposeSeq() > p.getProposeNum() {
			p.proposeNum = promise.getProposeSeq()
			p.proposeVal = promise.getProposeVal()
		}
	}
}

func (p *proposer) majority() int {
	return len(p.acceptors)/2 + 1
}

func (p *proposer) getRecevPromiseCount() int {
	recvCount := 0
	for _, acepMsg := range p.acceptors {
		log.Println(" proposer has total ", len(p.acceptors), " acceptor ", acepMsg, " current Num:", p.getProposeNum(), " msgNum:", acepMsg.getProposeSeq())
		if acepMsg.getProposeSeq() == p.getProposeNum() {
			log.Println("recv ++", recvCount)
			recvCount++
		}
	}
	log.Println("Current proposer received promise count ", recvCount)
	return recvCount
}

func (p *proposer) majorityReached() bool {
	return p.getRecevPromiseCount() > p.majority()
}
func (p *proposer) getProposeNum() int {
	p.proposeNum = p.sequence<<4 | p.id
	return p.proposeNum
}