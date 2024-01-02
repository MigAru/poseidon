package gc

func (gc *GC) clearDigests() error {
	digests, err := gc.db.GetDigestsForDelete()
	if err != nil {
		return err
	}

	for _, digest := range digests {
		if err := gc.fs.DeleteDigest(digest); err != nil {
			gc.log.
				WithField("operation", "delete_digest_in_fs").
				WithField("id", digest).
				Error(err)
		}
	}

	return nil
}
