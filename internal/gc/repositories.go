package gc

import (
	"github.com/MigAru/poseidon/internal/database/structs"
)

func (gc *GC) clearRepositories() error {
	repositories, err := gc.db.GetRepositoriesForDelete()
	if err != nil {
		return err
	}

	for _, repository := range repositories {
		if err := gc.clearRepository(repository); err != nil {
			gc.log.
				WithField("operation", "clear_repository").
				WithField("id", repository.ID).
				Error(err)
		}
	}

	return nil
}

func (gc *GC) clearRepository(repository *structs.Repository) error {
	use, err := gc.db.DigestUseAnotherRepository(repository.Digest)
	if err != nil {
		return err
	}

	if !use {
		if err := gc.db.MarkDeleteDigest(nil, repository.Digest); err != nil {
			return err
		}
	}

	if err := gc.db.DeleteRepository(repository.ID); err != nil {
		return err
	}

	return nil
}
