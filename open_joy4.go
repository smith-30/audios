package main

import (
	"fmt"

	"github.com/nareix/joy4/av"
	"github.com/nareix/joy4/av/avutil"
	"github.com/nareix/joy4/format"
)

func init() {
	format.RegisterAll()
}

func main() {
	file, _ := avutil.Open("./sample_1mb.mp4")
	defer file.Close()

	streams, _ := file.Streams()
	for _, stream := range streams {
		if stream.Type().IsAudio() {
			astream := stream.(av.AudioCodecData)
			fmt.Println(astream.Type(), astream.SampleRate(), astream.SampleFormat(), astream.ChannelLayout())
		} else if stream.Type().IsVideo() {
			vstream := stream.(av.VideoCodecData)
			fmt.Println(vstream.Type(), vstream.Width(), vstream.Height())
		}
	}

	idx := 0

	for {
		var pkt av.Packet
		var err error
		if pkt, err = file.ReadPacket(); err != nil {
			break
		}

		if streams[pkt.Idx].Type() == av.AAC {
			fmt.Println("pkt", idx, streams[pkt.Idx].Type(), "packet len", len(pkt.Data), "keyframe", pkt.IsKeyFrame)
			idx++
		}
	}

}
