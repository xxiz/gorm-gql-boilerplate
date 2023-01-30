package dbrepo

import (
	log "github.com/sirupsen/logrus"
)

func (m *postgresDBRepo) AllUsers() bool {
	log.Info("AllUsers was called")
	return true
}
