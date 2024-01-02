package gc

func (gc *GC) clearBlobs() error {
	blobs, err := gc.fs.GetBlobsForDelete()
	if err != nil {
		return err
	}

	for _, blob := range blobs {
		if err := gc.fs.DeleteBlob(blob); err != nil {
			gc.log.
				WithField("operation", "delete_blob_in_fs").
				WithField("id", blob).
				Error(err)
		}
	}
	return nil
}
