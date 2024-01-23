package backupclient

//backup_type ej: "create","update","delete"
func (b *backupClient) BackupOneObjectType(backup_type, table_name string, data_in any) {

	b.e = "Backup Table:: " + table_name + " Type: " + backup_type + " "

	b.object, b.err = b.GetObjectBY("", table_name)
	if b.err != "" {
		b.Log(b.e + b.err)
		return
	}

	b.SendOneRequest("POST", backup_type, b.object.ObjectName, data_in, func(resp []map[string]string, err string) {

		if err != "" {
			// if on_server_too { // necesita respaldo en servidor
			// 	data["create"] = "backup" //estado backup = no respaldado
			// }

			b.Log(b.e+"error respaldo en el servidor", err)
			return
		}

		if len(resp) == 0 {
			b.Log(b.e + "error se esperaba al menos un dato de respuesta desde el servidor después de la eliminación")
		}

	})
}
