package di

import (
	"initial-project-go/di/server"
)

func InitApplication() error {
	err := server.InitApiServer()
	if err != nil {
		return err
	}
	return nil
}
