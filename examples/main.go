/*
 * @Description:
 * @Author: dotwoo@gmail.com
 * @Github: https://github.com/dotwoo
 * @Date: 2021-07-22 13:10:06
 * @FilePath: /joy4/examples/main.go
 */
package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"

	"github.com/dotwoo/joy4"
	"github.com/dotwoo/joy4/format"
)

func init() {
	format.RegisterAll()
}

var (
	testURL = "rtsp://test.com/pss/ld2342"
)

func main() {
	imgCh := make(chan image.Image, 10)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		err := joy4.StreamRead(ctx, testURL, 2, 10, imgCh)
		if err != nil {

			fmt.Println("capture error:", err)
		}
		close(imgCh)
	}()

	i := 0
	for img := range imgCh {
		var imageBuf bytes.Buffer
		_ = jpeg.Encode(&imageBuf, img, nil)
		ioutil.WriteFile(fmt.Sprintf("/tmp/te%d.jpg", i), imageBuf.Bytes(), 0644)
		println("write", i)
		i++
		if i > 10 {
			close(imgCh)
			cancel()
		}
	}

	println("write ims", i)
	// ioutil.WriteFile("/tmp/dec.jpg", bs, 0644)

}
