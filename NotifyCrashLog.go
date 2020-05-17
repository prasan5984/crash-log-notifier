package b2c

import (
	"fmt"
	"net/http"
)

func NotifyCrashLog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
