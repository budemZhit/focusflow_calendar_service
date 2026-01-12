package handler

import (
	"net/http"
	fromdb "github.com/budemZhit/focusflow_calendar_service/iternal/from_db"
)


type SchedulerHandler struct{}

func NewSchedulerHandler() *SchedulerHandler {
	return &SchedulerHandler{}
}

func (h *SchedulerHandler) GetMonthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(fromdb.GetCalendarJson())
}

func (h *SchedulerHandler) GetEventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(fromdb.GetCategories())
}
