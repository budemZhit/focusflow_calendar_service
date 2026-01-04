package router

import (
	"log/slog"
	"net/http"

	"github.com/budemZhit/focusflow_calendar_service/iternal/handler"
	commonmw "github.com/budemZhit/focusflow_common/middleware"

	"github.com/go-chi/chi"
)

func InitRouter(log *slog.Logger) *chi.Mux {

	handler := handler.NewSchedulerHandler()
	return setupRouter(handler, log)
}

func setupRouter(schedulerHandler *handler.SchedulerHandler, log *slog.Logger) *chi.Mux {
	r := chi.NewRouter()

	commonmw.InitBase(r, log)

	r.Get("/calendar/month", schedulerHandler.GetMonthHandler)
	r.Get("/calendar/events", schedulerHandler.GetEventsHandler)
	http.ListenAndServe(":8080", r)

	return r
}
