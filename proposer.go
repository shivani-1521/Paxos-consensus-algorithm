package paxos

func NewProposer(id int, val string, net nodeNetwork, acceptors ...int) *proposer{
	pro := proposer{id: id, proposeVal: val, seq: 0, net:net}
	for _, acceptor := range acceptors{
		pro.acceptors[acceptor] = message{}
	}

	return &pro
}

type proposer struct{
	id	int
	sequence	int
	proposeVal	string
	acceptors	map[int]message
	net         nodeNetwork
}

func (p *proposer) run(){
	for  p.majorityReached(){
		outMsgs := p.prepare()
		for _, msg := range outMsgs{
			p.net.send(msg)
		}

		m := p.net.recev()
		for m != nil{
			continue
		}
	}
}

func (p *proposer)prepareMajorityMessages(stag msgType, val string) []message{
	sendMsgCount := 0
	var msgList []message
	for acepId, acepMsg := range p.acceptors{
		if acepMsg.getSeqNumber() == p.proposeNum() {
			msg := message{from: p.id, to: acepId, typ: stag, sequence: p.proposeNum()}
			//Only need value on propose, not in prepare
			if stag == Propose {
				msg.val = value
			}
			msgList = append(msgList, msg)
		}
		sendMsgCount++
		if sendMsgCount > p.majority() {
			break
		}
	}
	return msgList
}