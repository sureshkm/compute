package compute

import (
	"os"
)

type Args struct {
	Incoming chan Packet
	Outgoing chan Packet
	//Done chan bool
	Container map[string]interface{}
	Logger    *Log
}

type Computes interface {
	String() string
	Execute(*Args)
}

type Packet map[string]interface{}

func Run(computes ...Computes) {
	//done := make()
	in := make(chan Packet, 10000)
	logger := Logger(os.Stdout)
	logger.logf("Initializing Compute ...")
	//done := make(chan bool)
	var indx = 1
	for _, compute := range computes {
		out := make(chan Packet, 10000)
		arg := Args{Incoming: in, Outgoing: out, Logger: logger}
		//for i := 0; i < indx; i++ {
		logger.logf("Initializing Compute: %s", compute.String())
		go compute.Execute(&arg)
		//}
		in = out
		indx += 1
	}

	for {
		_ = <-in
	}
}

func NewPacket() Packet {
	packet := make(Packet)
	return packet
}
