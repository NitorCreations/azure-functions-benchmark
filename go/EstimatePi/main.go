package HttpTrigger

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/NitorCreations/azure-functions-go-handler/pkg/function"
)

func estimatePi(n int) float64 {
	total := 0
	totalIn := 0
	for i := 0; i < n; i++ {
			var x = rand.Float64()
			var y = rand.Float64()
			if x*x+y*y < 1 {
					totalIn++
			}
			total++
	}
	return float64(totalIn) * 4.0 / float64(total)
}

func Handle(ctx *function.Context, req *function.HttpRequest) function.HttpResponse {
	t0 := time.Now().UnixMicro()

	n := 100000
	if v, ok := req.Query["n"]; ok {
		i, _ := strconv.ParseInt(v, 10, 32)
		n = int(i)
	}

	pi := estimatePi(n)

	t1 := time.Now().UnixMicro()
	d := float64(t1 - t0) / 1_000
	
	return function.HttpResponse{
		Body: function.H{
			"duration": d,
			"result": pi,
		},
		Headers: function.ContentTypeJson,
	}
}
