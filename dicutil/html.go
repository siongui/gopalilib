package dicutil

import (
	"github.com/siongui/gotm"
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
// The first three parameters are the same as the parameters of
// html/template.ExecuteTemplate method.
func CreateHTML(w io.Writer, name string, data *TemplateData, localeDir, tmplDir string) (err error) {
	gotm.SetupMessagesDomain(localeDir)

	tm := gotm.NewTemplateManager(name)
	err = tm.ParseDirectoryWithGettextFunction(tmplDir)
	if err != nil {
		return
	}

	gotm.SetLocale(data.OgLocale)
	err = tm.ExecuteTemplate(w, name, data)
	return
}
