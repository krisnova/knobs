package rtmp

import (
	"fmt"

	"github.com/kris-nova/knobs/stream"
	"github.com/kris-nova/rtmp/server"
)

type RTMPStreamer struct {
	Options *stream.StreamerOptions
}

func New(options *stream.StreamerOptions) stream.Streamer {
	return &RTMPStreamer{
		Options: options,
	}
}

func (s *RTMPStreamer) Stream() error {

	sconn, err := server.New(s.Options.StreamListenAddr)
	if err != nil {
		fmt.Errorf("unable to init server: %v", err)
	}

	go sconn.Accept()
	defer sconn.Close()

	for {
		//select {
		//case _ := <-sconn.Clients():
		//	// ...
		//case _ := <-sconn.Errs():
		//	// ...
		//}
	}

	return nil
}
