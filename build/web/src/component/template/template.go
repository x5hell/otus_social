package template

import (
	"html/template"
	"model"
	"path/filepath"
	"strings"
)

const ProjectRoot = "./"
const LF = string(filepath.Separator)
const HtmlPath = "template" + LF
const layoutTemplatePath = "layout.html"
const HeaderTemplatePath = "layout-header.html"
const LayoutName = "layout"

func OpenUserTemplate(templateName string) (htmlTemplate *template.Template, err error) {
	htmlTemplate = template.New(templateName)
	return htmlTemplate.
		Funcs(model.GetTemplateFunctions()).
		ParseFiles(
			getTemplateRelativePath(layoutTemplatePath),
			getTemplateRelativePath(HeaderTemplatePath),
			getTemplateRelativePath(templateName))
}

func getTemplateRelativePath(templateName string) string {
	filePathParts := []string{ProjectRoot, HtmlPath, templateName}
	return strings.Join(filePathParts, LF)
}