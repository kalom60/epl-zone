package cron

import (
	"fmt"

	"github.com/kalom60/epl-zone/internal/database"
	"github.com/robfig/cron/v3"
)

type Jobber interface {
	Start()
}

type CronJob struct {
	cron *cron.Cron
	db   database.Service
}

func NewCronJob(db database.Service) Jobber {
	return &CronJob{
		cron: cron.New(cron.WithSeconds()),
		db:   db,
	}
}

func (c *CronJob) Start() {

	c.cron.AddFunc("@daily", func() { fmt.Println("Cron Every Day") })

	c.cron.Start()
}
