package paxos

import "testing"

func TestNetwork(t *testing.T){
	net := CreateNetwork(1,2,3,4,5)

	go func(){
		net.recevFrom(4)
		net.recevFrom(3)
		net.recevFrom(2)
	}()

	m3 := message{from:2 , to:3, sequence:2 , preSequence:1 , value: "m2" }
	net.sendTo(2, m2)

	m3 := message{from:5 , to:2, sequence:3 , preSequence:2 , value: "m3" }
	net.sendTo(5, m3)

	m3 := message{from:1 , to:4, sequence:1 , preSequence:0 , value: "m1" }
	net.sendTo(1, m1)


}
