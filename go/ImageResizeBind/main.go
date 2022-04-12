package HttpTrigger

import (
	"bytes"
	"encoding/json"
	"image"
	"image/jpeg"
	"io"
	"log"
	"time"

	"golang.org/x/image/draw"

	"github.com/NitorCreations/azure-functions-go-handler/pkg/function"
)

func resize(r io.Reader, w io.Writer) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	
	rect := image.Rect(0, 0, 300, 200)
	dst := image.NewRGBA(rect)
	
	draw.NearestNeighbor.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)

	opts := jpeg.Options{
		Quality: 80,
	}

	return jpeg.Encode(w, dst, &opts)
}

func Handle(ctx *function.Context) function.HttpResponse {
	t0 := time.Now().UnixMicro()

	data, ok := ctx.Data["srcImage"]
	if !ok {
		log.Fatalln("No input data")
	}

	b := make([]byte, len(data))
	if err := json.Unmarshal(data, &b); err != nil {
		log.Fatalln(err)
	}
	
	r := bytes.NewReader(b)
	w := bytes.NewBuffer([]byte{})
	if err := resize(r, w); err != nil {
		log.Fatalln(err)
	}
	
	ctx.Outputs["dstImage"] = w.Bytes()
	
	t1 := time.Now().UnixMicro()
	d := float64(t1 - t0) / float64(1000)
	
	return function.HttpResponse{
		Body: function.H{
			"duration": d,
		},
		Headers: function.ContentTypeJson,
	}
}
