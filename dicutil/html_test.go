package dicutil

import (
	"os"
	"testing"
)

func TestGensite(t *testing.T) {
	data := TemplateData{
		SiteUrl:     "http://dictionary.online-dhamma.net",
		TipitakaURL: "http://tipitaka.online-dhamma.net",
		OgImage:     "https://upload.wikimedia.org/wikipedia/commons/d/df/Dharma_Wheel.svg",
		OgUrl:       "http://dictionary.online-dhamma.net/",
		OgLocale:    "en_US",
	}

	err := CreateHTML(os.Stdout, "index.html", &data, "../../pali/common/locale", "../../paligo/theme/template")
	if err != nil {
		t.Error(err)
	}
}
