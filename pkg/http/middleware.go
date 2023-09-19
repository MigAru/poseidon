package http

import (
	"errors"
	registryErrors "github.com/MigAru/poseidon/pkg/registry/errors"
	http2 "net/http"
	"regexp"
)

var (
	regexpValidate = regexp.MustCompile("[a-z0-9]+(?:[._-][a-z0-9]+)*")
	findAllRegexp  = -1
)

func ValidateProjectName(ctx Context) error {
	var (
		project = ctx.Param("name")
		subName = ctx.Param("sub-name")
	)

	if project != "" {
		if err := validateMatch(regexpValidate.FindAllString(project, findAllRegexp), project); err != nil {
			ctx.JSON(http2.StatusBadRequest, registryErrors.NameInvalid)
			return err
		}
	}

	if subName != "" {
		if err := validateMatch(regexpValidate.FindAllString(subName, findAllRegexp), subName); err != nil {
			ctx.JSON(http2.StatusBadRequest, registryErrors.NameInvalid)
			return err
		}

	}

	if project != "" && subName != "" {
		ctx.JSON(http2.StatusBadRequest, registryErrors.NameInvalid)
		return errors.New("name repository not found")
	}

	return nil
}

func validateMatch(matches []string, reference string) error {
	length := len(matches)
	if length > 1 || length <= 0 {
		return errors.New("implemented me")
	}
	if matches[0] != reference {
		return errors.New("implemented me")
	}
	return nil
}
