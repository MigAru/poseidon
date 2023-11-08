package gc

import "context"

func (gc *GC) clear() (errs []error) {

	files, err := gc.redis.GetMarkedFiles(context.Background(), "blobs")
	if err != nil {
		errs = append(errs, err)
	}

	for _, name := range files {
		if err := gc.fs.DeleteBlob(name); err != nil {
			errs = append(errs, err)
		}
	}

	return
}
