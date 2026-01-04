package main

import (
	"github.com/budemZhit/focusflow_calendar_service/iternal/router"
	"github.com/budemZhit/focusflow_common/logger"
)

func main() {
	log := logger.Init("calendar_service", "debug")
	r := router.InitRouter(log)
	_ = r
}
