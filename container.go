package main

import (
	"api48hours/repository"
	"database/sql"
)

var container = map[string]interface{}{
	"": nil,
}

func GetRepository() *repository.IRepository {
	if container["repository"] == nil {
		database := sql.DB{}
		container["repository"] = repository.NewMySQLRepository(
			&database,
		)
	}
	return container["repository"].(*repository.IRepository)
}
