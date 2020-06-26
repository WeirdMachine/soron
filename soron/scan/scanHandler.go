package scan

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ScanHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	body, _ := ioutil.ReadAll(r.Body)
	_, _ = fmt.Fprintf(w, "%s", string(body))
}
