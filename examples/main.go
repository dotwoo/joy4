/*
 * @Description:
 * @Author: dotwoo@gmail.com
 * @Github: https://github.com/dotwoo
 * @Date: 2021-07-22 13:10:06
 * @FilePath: /joy4/examples/main.go
 */
package main

import (
	"fmt"
	"time"

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
	i := 0
	for {
		i++
		start := time.Now()
		bs, err := joy4.Capture(testURL)
		if err != nil {
			fmt.Println("capture error:", err)
			continue
		}

		println("times:", i, "use:", time.Since(start).String(), "write length:", len(bs))
		// ioutil.WriteFile("/tmp/dec.jpg", bs, 0644)
	}
}
