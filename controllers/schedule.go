package controllers

import (
	"../config"
)

func GiveSchedule(grade, class int) string {
	_ = config.BuildAPIConfig()

	return "No data"
}
