package b2c

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func NotifyCrashLog(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client, _ := storage.NewClient(ctx, option.WithoutAuthentication())
	bkt := client.Bucket("crash-logs")
	obj := bkt.Object("crash-file-1")
	storeWriter := obj.NewWriter(ctx)
	fmt.Fprintf(storeWriter, "Test data\n")
}
