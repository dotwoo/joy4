/*
 * @Description:
 * @Author: dotwoo@gmail.com
 * @Github: https://github.com/dotwoo
 * @Date: 2021-08-02 13:30:47
 * @FilePath: /joy4/stream_capture.go
 */
package joy4

import (
	"context"
	"image"
	"time"

	"github.com/dotwoo/joy4/av"
	"github.com/dotwoo/joy4/av/avutil"
	"github.com/dotwoo/joy4/cgo/ffmpeg"
)

func StreamRead(ctx context.Context, streamURL string, capturePerSecond int, brokeErrs int, imgCh chan image.Image) (err error) {
	if capturePerSecond < 0 {
		capturePerSecond = 2
	}
	if brokeErrs < 1 {
		brokeErrs = 10
	}
	var pps = 1
	var duOfp = time.Duration(0)
	file, err := avutil.Open(streamURL)
	if err != nil {
		return err
	}
	defer file.Close()

	streams, err := file.Streams()
	if err != nil {
		return err
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
		return err
	}
	err = dec.Setup()
	if err != nil {
		return err
	}

	// var err error
	var skipP = 0
	i := 0
	for {
		if ctx.Err() != nil {
			return nil
		}

		var pkt av.Packet
		if pkt, err = file.ReadPacket(); err != nil {
			return err
		}

		if streams[pkt.Idx].Type() != vstream.Type() {
			continue
		}

		i++
		if duOfp == 0 && pkt.Time > 0 {
			duOfp = pkt.Time
			pps = int(time.Second/duOfp) / capturePerSecond
			skipP = pps - 1
			// fmt.Println("pps:", pps, "timeOfp:", duOfp)
		}
		if skipP > 1 {
			skipP--
			err = dec.FeedFrame(pkt.Data, i-1)
			if err != nil {
				brokeErrs--
				if brokeErrs == 0 {
					return err
				}
			}
			continue
		} else {
			skipP = pps - 1
		}

		// fmt.Println("pkt", i, streams[pkt.Idx].Type(), "len", len(pkt.Data), "keyframe", pkt.IsKeyFrame, pkt.Time, pkt.CompositionTime)
		if !pkt.IsKeyFrame {
			i = 1
		}
		// fmt.Println(pkt.Time, "len", len(pkt.Data), "keyframe", pkt.IsKeyFrame)
		img, err := dec.Decode(pkt.Data, i-1)
		if err != nil {
			brokeErrs--
			if brokeErrs == 0 {
				return err
			}
			continue
		}

		imgCh <- &img.Image
		img.Free()
	}
	// fmt.Println("pkt", streams[pkt.Idx].Type(), "len", len(pkt.Data), "keyframe", pkt.IsKeyFrame)
}
