package service

import (
	"x-ui-scratch/database"
	"x-ui-scratch/xray"
)

type InboundService struct {
	xrayApi xray.XrayAPI
}

func (s *InboundService) AddTraffic(inboundTraffics []*xray.Traffic, clientTraffics []*xray.ClientTraffic) (error, bool) {
	var err error
	db := database.GetDB()
	tx := db.Begin()

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	panic("todo AddTraffic")
}
