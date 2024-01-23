package backupclient

import "github.com/cdvelop/model"

var pending_table = &model.Object{
	ObjectName: "bkp.pending",
	Table:      "pending",
	Fields: []model.Field{
		{Name: "id_pending", Legend: "Pendientes de respaldo", Unique: false},
		{Name: "action", Legend: "Acción Pendiente"}, //create,update,delete
		{Name: "object_name", Legend: "Nombre Objeto"},
		{Name: "contend", Legend: "Contenido"},
	},
}

func AddBackupHandlerAdapter(h *model.MainHandler, c *Config) (b *backupClient) {

	m := &model.Module{
		ModuleName:             "backup",
		Title:                  "Respaldos",
		FrontendModuleHandlers: model.FrontendModuleHandlers{},
		Areas:                  map[string]string{},
		Inputs:                 []*model.Input{},
		MainHandler:            h,
	}

	m.AddObjectsToModule(pending_table)

	h.AddModules(m)

	b = &backupClient{
		Config:                c,
		Logger:                h,
		FetchAdapter:          h,
		DataBaseAdapter:       h,
		ObjectsHandlerAdapter: h,
		backupRespond:         func(err string) {},
		backups:               []*backup{},
		remaining:             0,
		err:                   "",
	}

	h.BackupHandlerAdapter = b

	return b
}
