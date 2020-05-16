package rtmp

import (
	"github.com/kris-nova/knobs/stream"
)

type RTMPStreamer struct {
}

func New() stream.Streamer {
	return &RTMPStreamer{}
}

func (s *RTMPStreamer) Stream() error {
	return nil
}
