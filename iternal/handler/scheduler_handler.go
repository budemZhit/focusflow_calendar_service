package handler

import (
	_ "embed"
	"net/http"
)

//go:embed build_data/get_calendar.json
var getCalendarJSON []byte

//go:embed build_data/get_categories.json
var getCategoriesJSON []byte

type SchedulerHandler struct{}

func NewSchedulerHandler() *SchedulerHandler {
	return &SchedulerHandler{}
}

func (h *SchedulerHandler) GetMonthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(getCalendarJSON)
}

func (h *SchedulerHandler) GetEventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(getCategoriesJSON)
}
