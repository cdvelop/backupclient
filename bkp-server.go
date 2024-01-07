package backupclient

//backup_type ej: "create","update","delete"
func (b *backupClient) BackupOneObjectType(backup_type, table_name string, data_in any) {

	b.e = "BackupObject: " + table_name + " Type: " + backup_type + " "

	b.object, b.err = b.GetObjectBY("", table_name)
	if b.err != "" {
		b.Log(b.e + b.err)
		return
	}

	b.SendOneRequest("POST", backup_type, b.object.ObjectName, data_in, func(resp []map[string]string, err string) {

		if err != "" {
			// if backup_required { // necesita respaldo en servidor
			// 	data["create"] = "backup" //estado backup = no respaldado
			// }

			b.Log(b.e+"error respaldo en el servidor acción:", backup_type, err)
			return
		}

		b.Log(b.e+"RESPUESTA ACCIÓN DE RESPALDO:", backup_type, "EN EL SERVIDOR:", resp)

	})
}
