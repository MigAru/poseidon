package http

func GetProjectName(ctx Context) string {
	var project string
	name := ctx.Param("project")
	subName := ctx.Param("sub-name")
	projectName := name
	if subName != "" {
		projectName += "." + subName
	}
	return project
}
