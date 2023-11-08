package gc

func (gc *GC) markBlobs() (errs []error) {
	files, err := gc.fs.GetBlobsForDelete()
	if err != nil {
		errs = append(errs, err)
	}
	for _, name := range files {
		if err := gc.redis.SetMarkFile(gc.ctx, "blobs", name); err != nil {
			errs = append(errs, err)
		}
	}

	return
}
