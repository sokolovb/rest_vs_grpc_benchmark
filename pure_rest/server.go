package pure_rest

import (
	"encoding/json"
	"net/http"
	md "rest-vs-grpc/mock_data"
)

const PortRest = "3000"

type handler struct {
	data *md.Data
}

func newHandler(data *md.Data) *handler {
	return &handler{data: data}
}

func StartPureRestServer(data *md.Data, port string) error {
	h := newHandler(data)
	mux := http.NewServeMux()
	mux.HandleFunc("/int", h.integer)
	mux.HandleFunc("/int-slice", h.integerSlice)
	mux.HandleFunc("/string", h.str)
	mux.HandleFunc("/string-slice", h.strSlice)
	mux.HandleFunc("/blob", h.blob)
	mux.HandleFunc("/struct", h.structure)
	return http.ListenAndServe("localhost:"+port, mux)
}

func (h *handler) integer(w http.ResponseWriter, r *http.Request) {
	formResponse(w, h.data.GetInt())
	return
}

func (h *handler) integerSlice(w http.ResponseWriter, r *http.Request) {
	formResponse(w, h.data.GetIntSlice())
	return
}

func (h *handler) str(w http.ResponseWriter, r *http.Request) {
	formResponse(w, h.data.GetString())
	return
}

func (h *handler) strSlice(w http.ResponseWriter, r *http.Request) {
	formResponse(w, h.data.GetStringSlice())
	return
}

func (h *handler) blob(w http.ResponseWriter, r *http.Request) {
	formResponse(w, h.data.GetBlob())
	return
}

func (h *handler) structure(w http.ResponseWriter, r *http.Request) {
	formResponse(w, h.data.GetStruct())
	return
}

func formResponse(w http.ResponseWriter, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(body)
}
