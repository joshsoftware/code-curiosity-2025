package config

import (
	"context"

	"cloud.google.com/go/bigquery"
)

type Bigquery struct {
	Client *bigquery.Client
}

func BigqueryInit(ctx context.Context, appCfg AppConfig) (Bigquery, error) {
	client, err := bigquery.NewClient(ctx, appCfg.BigqueryProject.ProjectID)
	if err != nil {
		return Bigquery{}, err
	}

	bigqueryInstance := Bigquery{
		Client: client,
	}
	return bigqueryInstance, nil
}
