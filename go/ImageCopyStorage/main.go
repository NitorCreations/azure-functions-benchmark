package HttpTrigger

import (
	"bytes"
	"context"
	"log"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/NitorCreations/azure-functions-go-handler/pkg/function"
)

func Handle(ctx *function.Context) function.HttpResponse {
	t0 := time.Now().UnixMicro()

	connStr, ok := os.LookupEnv("StorageConnectionString")
	if !ok || connStr == "" {
		log.Fatal("StorageConnectionString string not set")
	}

	serviceClient, err := azblob.NewServiceClientFromConnectionString(connStr, nil)
	if err != nil {
		log.Fatal(err)
	}

	containerClient := serviceClient.NewContainerClient("images")
	srcBlob := containerClient.NewBlockBlobClient("src.jpeg")
	dstBlob := containerClient.NewBlockBlobClient("dst.jpeg")

	context := context.Background()
		
	get, err := srcBlob.Download(context, nil)
	if err != nil {
    log.Fatal(err)
	}

	data := &bytes.Buffer{}
	reader := get.Body(&azblob.RetryReaderOptions{})
	_, err = data.ReadFrom(reader)
	if err != nil {
		log.Fatal(err)
	}
	err = reader.Close()
	if err != nil {
		log.Fatal(err)
	}

	_, err = dstBlob.UploadBufferToBlockBlob(
		context, data.Bytes(), azblob.HighLevelUploadToBlockBlobOption{})
	if err != nil {
		log.Fatal(err)
	}
	
	t1 := time.Now().UnixMicro()
	d := float64(t1 - t0) / float64(1000)
	
	return function.HttpResponse{
		Body: function.H{
			"duration": d,
		},
		Headers: function.ContentTypeJson,
	}
}
