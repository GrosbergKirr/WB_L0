package v1

import (
	"fmt"
	"github.com/go-chi/render"
	"log/slog"
	"net/http"
	"project0/models"
)

type Request struct {
	Uid string `json:"order_uid"`
}

func OrderGetter(log *slog.Logger, cache map[string]models.Order) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("Get_handler, json decode error")
			w.WriteHeader(http.StatusBadRequest)
		}
		id := req.Uid
		fmt.Println(id)

		StatusGetRespOK(w, r, cache[id])
		log.Debug("render json")
	}
}
func StatusGetRespOK(w http.ResponseWriter, r *http.Request, order models.Order) {
	render.JSON(w, r, order)
}
