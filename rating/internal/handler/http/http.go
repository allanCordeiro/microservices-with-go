package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/allancordeiro/movieapp/rating/internal/controller/rating"
	model "github.com/allancordeiro/movieapp/rating/pkg"
)

// Handler defines a rating service controller
type Handler struct {
	ctrl *rating.Controller
}

// New creates a new rating service HTTP handler
func New(ctrl *rating.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// Handle handles PUT and GET /rating requests
func (h *Handler) Hande(w http.ResponseWriter, r *http.Request) {
	recordID := model.RecordID(r.FormValue("id"))
	if recordID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	recordType := model.RecordType(r.FormValue("type"))
	if recordType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAggregateRating(r.Context(), recordID, recordType)
		if err != nil && errors.Is(err, rating.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("response encode error: %v\n", err)
		}
	case http.MethodPut:
		userId := model.UserID(r.FormValue("userId"))
		v, err := strconv.ParseFloat(r.FormValue("value"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		ratingData := &model.Rating{
			UserID: userId,
			Value:  model.RatingValue(v),
		}

		if err := h.ctrl.PutRating(r.Context(), recordID, recordType, ratingData); err != nil {
			log.Printf("repository put error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
