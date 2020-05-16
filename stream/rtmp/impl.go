package rtmp



type RTMPStreamer struct {


}


func New() stream.Streamer {
	return &RTMPStreamer{}
}