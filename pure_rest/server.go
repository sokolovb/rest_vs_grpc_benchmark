package pure_rest

import (
	"encoding/json"
	md "github.com/sokolovb/rest_vs_grpc_benchmark/mock_data"
	"net/http"
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
	mux.HandleFunc("/struct-slices", h.structureSlices)
	mux.HandleFunc("/struct-structs", h.structureStructs)
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

func (h *handler) structureSlices(w http.ResponseWriter, r *http.Request) {
	formResponse(w, h.data.GetStructSlices())
	return
}

func (h *handler) structureStructs(w http.ResponseWriter, r *http.Request) {
	formResponse(w, h.data.GetStructStructs())
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
