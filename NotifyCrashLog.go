// Package p contains an HTTP Cloud Function.
package p

import (
	"fmt"
	"net/http"
)

func NotifyCrashLog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
