package http

func GetProjectName(ctx Context) string {
	project := ctx.Param("project")
	subName := ctx.Param("project-sub-name")
	if subName != "" {
		project += "." + subName
	}
	return project
}
