package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"

	"basic-go/webook/internal/events"
)

type App struct {
	web       *gin.Engine
	consumers []events.Consumer
	cron      *cron.Cron
}
