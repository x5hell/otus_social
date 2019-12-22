package template

import (
	"html/template"
	"path/filepath"
	"strings"
)

const ProjectRoot = "./"
const LF = string(filepath.Separator)
const HtmlPath = "template" + LF
const layoutTemplatePath = "layout.html"
const HeaderTemplatePath = "layout-header.html"
const HeaderGuestTemplatePath = "layout-header-guest.html"
const HeaderUserTemplatePath =  "layout-header-user.html"
const LayoutName = "layout"

func OpenGuestTemplate(templateName string) (htmlTemplate *template.Template, err error) {
	htmlTemplate = template.New(templateName)
	return htmlTemplate.ParseFiles(
		getTemplateRelativePath(layoutTemplatePath),
		getTemplateRelativePath(HeaderTemplatePath),
		getTemplateRelativePath(HeaderGuestTemplatePath),
		getTemplateRelativePath(templateName))
}

func OpenUserTemplate(templateName string) (htmlTemplate *template.Template, err error) {
	htmlTemplate = template.New(templateName)
	return htmlTemplate.ParseFiles(
		getTemplateRelativePath(layoutTemplatePath),
		getTemplateRelativePath(HeaderTemplatePath),
		getTemplateRelativePath(HeaderUserTemplatePath),
		getTemplateRelativePath(templateName))
}

func getTemplateRelativePath(templateName string) string {
	filePathParts := []string{ProjectRoot, HtmlPath, templateName}
	return strings.Join(filePathParts, LF)
}