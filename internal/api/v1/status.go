package v1

import (
	"encoding/json"
	"github.com/anpotashev/mpd-ws-api/internal/api"
	"github.com/anpotashev/mpd-ws-api/internal/api/dto"
	"github.com/gorilla/mux"
	"net/http"
)

func (v1 *v1Router) newStatusRouter(router *mux.Router) {
	router.HandleFunc("/get", v1.getStatusHandler).Methods(http.MethodGet)
	router.HandleFunc("/random", v1.setRandom).Methods(http.MethodPost)
	router.HandleFunc("/repeat", v1.setRepeate).Methods(http.MethodPost)
	router.HandleFunc("/single", v1.setSingle).Methods(http.MethodPost)
	router.HandleFunc("/consume", v1.setConsume).Methods(http.MethodPost)
}

func (v1 *v1Router) getStatusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := v1.MpdApi.WithRequestContext(r.Context()).Status()
	var payload interface{}
	if err == nil {
		payload = dto.MapStatus(status)
	}
	checkErrorAndWriteResponseWithPayload(payload, err, w, r)
}

func (v1 *v1Router) setRandom(w http.ResponseWriter, r *http.Request) {
	var request SetRandomRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		api.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
	}
	checkErrorAndWriteResponse(v1.MpdApi.WithRequestContext(r.Context()).Random(request.Enabled), w, r)
}
func (v1 *v1Router) setRepeate(w http.ResponseWriter, r *http.Request) {
	var request SetRepeateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		api.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
	}
	checkErrorAndWriteResponse(v1.MpdApi.WithRequestContext(r.Context()).Repeat(request.Enabled), w, r)
}
func (v1 *v1Router) setSingle(w http.ResponseWriter, r *http.Request) {
	var request SetSingleRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		api.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
	}
	checkErrorAndWriteResponse(v1.MpdApi.WithRequestContext(r.Context()).Single(request.Enabled), w, r)
}
func (v1 *v1Router) setConsume(w http.ResponseWriter, r *http.Request) {
	var request SetConsumeRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		api.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
	}
	checkErrorAndWriteResponse(v1.MpdApi.WithRequestContext(r.Context()).Consume(request.Enabled), w, r)
}
