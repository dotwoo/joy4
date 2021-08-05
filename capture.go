/*
 * @Description:
 * @Author: dotwoo@gmail.com
 * @Github: https://github.com/dotwoo
 * @Date: 2021-07-29 16:33:50
 * @FilePath: /joy4/capture.go
 */

package joy4

import (
	"bytes"
	"errors"
	"fmt"
	"image/jpeg"
	"strings"

	"github.com/dotwoo/joy4/av"
	"github.com/dotwoo/joy4/av/avutil"
	"github.com/dotwoo/joy4/cgo/ffmpeg"
)

var captureErr = errors.New("capture stream error")

func Capture(streamURL string) (out []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("catch panic:", r)
			out = nil
			err = captureErr
			return
		}
	}()
	file, err := avutil.Open(streamURL)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	streams, err := file.Streams()
	if err != nil {
		return nil, err
	}
	var vstream av.VideoCodecData
	for _, stream := range streams {
		if stream.Type().IsVideo() {
			vstream = stream.(av.VideoCodecData)
			// fmt.Println(vstream.Type(), vstream.Width(), vstream.Height())
			break
		}
	}
	dec, err := ffmpeg.NewVideoDecoder(vstream)
	if err != nil {
		return nil, err
	}
	err = dec.Setup()
	if err != nil {
		return nil, err
	}

	for i := 0; i < 10; i++ {
		var pkt av.Packet
		if pkt, err = file.ReadPacket(); err != nil {
			return nil, err
		}
		// fmt.Println("pkt", i, streams[pkt.Idx].Type(), "len", len(pkt.Data), "keyframe", pkt.IsKeyFrame)
		if streams[pkt.Idx].Type() == av.H264 {
			// fmt.Println("pkt", i, streams[pkt.Idx].Type(), "len", len(pkt.Data), "keyframe", pkt.IsKeyFrame)
			img, err := dec.Decode(pkt.Data, i)
			if err != nil {
				if strings.HasSuffix(err.Error(), "11") {
					continue
				}
				return nil, err
			}
			defer img.Free()

			if len(pkt.Data) > 10 {
				buf := new(bytes.Buffer)
				err = jpeg.Encode(buf, &img.Image, nil)
				if err != nil {
					return nil, err
				}
				return buf.Bytes(), nil
			}
		}
	}
	return nil, captureErr
}
