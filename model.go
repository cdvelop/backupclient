package backupclient

import "github.com/cdvelop/model"

type backupClient struct {
	*Config
	model.Logger
	model.FetchAdapter
	model.DataBaseAdapter
	model.ObjectsHandlerAdapter

	backupRespond func(err string)

	backups   []*backup
	remaining int

	object *model.Object

	ok  bool
	e   string
	err string
}

type backup struct {
	object   *model.Object
	data     []map[string]interface{}
	finished bool
	err      string
}

type Config struct {
	ImmediateServerBackup bool
}
