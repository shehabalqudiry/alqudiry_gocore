package templates

const SwaggerTemplate = `package {{ .Package }}

// @Summary List {{ .Module }}
// @Tags {{ .Module }}
// @Accept json
// @Produce json
// @Success 200
// @Router /{{ .Package }} [get]
`