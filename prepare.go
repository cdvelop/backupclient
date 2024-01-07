package backupclient

func (b *backupClient) prepareToSendData() {

	for _, bkp := range b.backups {

		b.Log("OBJETO", b.object.ObjectName)
		// var o *model.Object
		// for _, v := range d.objects {
		// 	if v.Table == o.Table {
		// 		o = v
		// 		break
		// 	}
		// }

		for _, item := range bkp.data {

			if b.object.Table == "file" {
				b.Log("TIPO FILE ENVIÓ FORM DATA", item)

				// d.http.SendAllRequests()

			} else {
				b.Log("ENVIÓ NORMAL JSON", item)

				if _, update := item["update"]; update {

					b.Log("UPDATE", item)

				} else {

					b.Log("CREATE", item)

				}
			}

		}

	}

}
