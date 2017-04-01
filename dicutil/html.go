package dicutil

import (
	"github.com/siongui/gotemplateutil"
	"io"
)

// Template data for webpage of Pāli Dictionary
type TemplateData struct {
	SiteUrl     string
	TipitakaURL string
	OgImage     string
	OgUrl       string
	OgLocale    string
}

// Pāli Dictionary is single page application (SPA).
// This method create HTML for SPA.
// The first three parameters is the same as html/template.ExecuteTemplate
// method.
func CreateHTML(w io.Writer, name string, data *TemplateData, localeDir, tmplDir string, isdev bool) (err error) {
	if isdev {
		data.SiteUrl = ""
	}

	gossg.SetupMessagesDomain(localeDir)

	tmpl, err := gossg.ParseDirectoryWithGettextFunction(tmplDir)
	if err != nil {
		return
	}

	gossg.SetLocale(data.OgLocale)
	err = tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		return
	}

	return
}
