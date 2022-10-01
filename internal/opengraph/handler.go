package opengraph

import "net/http"

type OpenGraphHandler struct {
	Service OpenGraphService
}

func (h *OpenGraphHandler) GetOpenGraphTags(w http.ResponseWriter, r *http.Request) {
	l := r.URL.Query().Get("link")
	og, err := h.Service.GetOpenGraphTags(l)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json, err := og.ToJSON()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
