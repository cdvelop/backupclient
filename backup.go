package backupclient

func (b *backupClient) BackupDataBase(callback func(err string)) {
	// reset valores a 0
	b.backups = []*backup{}
	b.backupRespond = callback
	b.remaining = len(b.GetAllObjects(true))

	b.Log("RESPALDANDO BASE DE DATOS INDEX DB NO IMPLEMENTADO")

	// callback("NO RESPALDADO AUN"))

	// b.addNewObjectsCreated()

}

// func (d *backupClient) addNewObjectsCreated() {

// 	for i, o := range b.getObjectsDB() {
// 		index := i // Captura el valor de i en esta iteración
// 		table := o.Table
// 		b.ReadAsyncDataDB(model.ReadParams{
// 			FROM_TABLE: table,
// 			WHERE:      []map[string]string{{"create": "backup", "update": "backup", "delete": "backup"}},
// 			RETURN_ANY: true,
// 		}, func(r *model.ReadResults, err string) {

// 			if err != "" {
// 				b.Log(err)
// 				return
// 			}

// 			if len(r.ResultsAny) != 0 {
// 				b.Log(r.ResultsAny)

// 				new := backup{
// 					object:   o,
// 					data:     r.ResultsAny,
// 					finished: false,
// 					err:      "",
// 				}

// 				b.backups = append(b.backups, new)

// 				b.Log("BACKUP REQUERIDO", table)
// 			}
// 			b.remaining--

// 			// finish
// 			b.finishReadData(index, table)
// 		})

// 	}

// }

// func (d *backupClient) finishReadData(index int, table string) {
// 	b.Log("INDICE ACTUAL:", index, table)

// 	b.Log("LECTURA RESTANTE:", b.remaining)

// 	if b.remaining == 0 {
// 		b.Log("LECTURA FINALIZADA")
// 		if len(b.backups) != 0 {
// 			b.Log("BACKUP A REALZAR:", len(b.backups))
// 			b.prepareToSendData()

// 		} else {
// 			b.Log("BACKUP OK NADA PARA ENVIAR")
// 			b.backupRespond("")
// 		}
// 	}
// }
