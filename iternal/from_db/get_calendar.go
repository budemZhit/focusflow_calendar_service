package fromdb

import (
	_ "embed"
)

//go:embed build_data/get_calendar.json
var getCalendarJSON []byte


func GetCalendarJson()[]byte{
	return getCalendarJSON
}