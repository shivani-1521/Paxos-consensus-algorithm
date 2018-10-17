package paxos

type msgType int

const(
	Prepare msgType = iota + 1
	Propose
	Promise
	Accept
)

type message struct{
	from			int
	to				int
	type    		msgType
	sequence		int
	preSequence		int
	value			string

}

func (m *message)getSeqNumber() int{
	return m.sequence
}

func (m *message)getProposeValue() string{
	return m.value
}

//if

func (m *message)getProposeSeq() int{
	switch m.typ {
	case Promise:
		return m.preSequence
	case Accept:
		return m.sequence
	default:
		panic("message type not supported")
	}
}