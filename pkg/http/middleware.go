package http

import (
	"errors"
	"regexp"
)

var (
	regexpValidate = regexp.MustCompile("[a-z0-9]+([a-z0-9]+)*")
	findAllRegexp  = -1
)

func ValidateProjectNameMiddleware(ctx Context) error {
	var (
		//TODO: сделать константу наименования проекта и его имени(в будущем сделать через регехп поиск урлы)
		project = ctx.Param("project")
		subName = ctx.Param("project-sub-name")
	)

	if project != "" {
		if err := validateMatch(regexpValidate.FindAllString(project, findAllRegexp), project); err != nil {
			return err
		}
	}

	if subName != "" {
		if err := validateMatch(regexpValidate.FindAllString(subName, findAllRegexp), subName); err != nil {
			return err
		}

	}

	if project == "" && subName == "" {
		return errors.New("name repository not found")
	}

	return nil
}

func validateMatch(matches []string, reference string) error {
	length := len(matches)
	if length != 1 {
		return errors.New("name repository not valid")
	}
	if matches[0] != reference {
		return errors.New("name repository not valid")
	}
	return nil
}
