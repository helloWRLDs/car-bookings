package encoders

import (
	"fmt"
	errs "helloWRLDs/bookings/pkg/web/errors"
	"net/http"
)

func SendErr(w http.ResponseWriter, error *errs.Error) {
	w.Header().Set("Content-Type", "plain/text")
	w.WriteHeader(error.Code())
	w.Write([]byte(fmt.Sprintf("%d %s", error.Code(), error.Error())))
}
