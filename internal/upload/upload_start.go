package upload

import "time"

func (u *Upload) start() {
	timer := time.NewTimer(u.timeout)
	for {
		select {
		case <-timer.C:
			u.log.Info("download is dead. exit routine")
			break
		case chunk := <-u.Queue:
			if err := u.fs.UploadBlob(u.ID, chunk); err != nil {
				u.Errors = append(u.Errors, err)
				u.log.Error(err.Error())
			}

			ok := timer.Reset(u.timeout)
			if !ok {
				u.log.Error("unable to update timer. download is dead. exit routine")
				break
			}
		}
	}
}
