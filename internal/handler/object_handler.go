package handler

import (
	"encoding/json"
	"net/http"
	"project/internal/service"
)

type ObjectHandler struct {
	objectService service.ObjectService
}

func NewObjectHandler(objectService service.ObjectService) *ObjectHandler {
	return &ObjectHandler{objectService: objectService}
}

func (h *ObjectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/objects" || r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	objects, err := h.objectService.FetchObjects()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(objects)
}
