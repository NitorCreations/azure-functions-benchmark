package HttpTrigger

import (
	"strconv"
	"time"

	"github.com/NitorCreations/azure-functions-go-handler/pkg/function"
)

func fib(s int) int64 {
	var x, y int64 = 0, 1
	for i := 0; i < s; i++ {
		temp := x
		x = x+y
		y = temp
	}
	return y
}

func Handle(ctx *function.Context, req *function.HttpRequest) function.HttpResponse {
	t0 := time.Now().UnixMicro()

	seq := 30
	if v, ok := req.Query["seq"]; ok {
		i, _ := strconv.ParseInt(v, 10, 32)
		seq = int(i)
	}

	f := fib(seq)

	t1 := time.Now().UnixMicro()
	d := float64(t1 - t0) / float64(1000)
	
	return function.HttpResponse{
		Body: function.H{
			"duration": d,
			"result": f,
		},
		Headers: function.ContentTypeJson,
	}
}
