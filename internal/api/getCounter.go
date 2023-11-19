package api

import (
	"counter/internal/model"
	"fmt"
	"net/http"
)

func GetCounterHandler(seqGen *model.Counter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(fmt.Sprintf("%d", seqGen.GetCounter())))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
