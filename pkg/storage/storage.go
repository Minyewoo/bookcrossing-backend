package storage

import (
	"bookcrossing-backend/ent"
	"context"
	"log"
)

type DBCredentials struct {
	DriverName, DataSourceName string
}

func SetupStorage(creds *DBCredentials) *ent.Client {
	storage, err := ent.Open(creds.DriverName, creds.DataSourceName)
	if err != nil {
		log.Fatalf("failed opening db connection: %s %s", creds.DriverName, creds.DataSourceName)
	}

	if err := storage.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return storage
}
