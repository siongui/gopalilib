package dicmgr

import (
	"bytes"
	"html/template"

	"github.com/siongui/gopalilib/lib"
)

// HtmlTemplateBookNameWordExps is html template for word definition.
//
// Message | Bulma
// https://bulma.io/documentation/components/message/
const HtmlTemplateBookNameWordExps = `
{{range $bnwe := .}}
<article class="message">
  <div class="message-header">
    <p>{{$bnwe.BookName}}</p>
  </div>
  <div class="message-body">
    {{$bnwe.Explanation}}
  </div>
</article>
{{end}}`

// HtmlTemplateSuggestedWords is html template for suggested words
const HtmlTemplateSuggestedWords = `
{{range $word := .}}
<div>{{$word}}</div>
{{end}}`

// GetWordDefinitionHtml returns the html string of word definition according to
// setting and window.navigator.languages
func GetWordDefinitionHtml(wi lib.BookIdWordExps, setting lib.PaliSetting, navigatorLanguages string) string {
	bnwes := lib.IdExps2BookNameWordExps(
		lib.BookIdWordExps2IdExpsAccordingToSetting(wi, bookIdAndInfos, setting, navigatorLanguages),
		bookIdAndInfos)

	t1, err := template.New("wordExplanation").Parse(HtmlTemplateBookNameWordExps)
	if err != nil {
		return err.Error()
	}
	// Google Search: go html template output string
	// https://groups.google.com/forum/#!topic/golang-nuts/dSFHCV-e6Nw
	var buf bytes.Buffer
	err = t1.Execute(&buf, bnwes)
	if err != nil {
		return err.Error()
	}
	return buf.String()
}

// GetSuggestedWordsHtml returns the html string of suggested words.
func GetSuggestedWordsHtml(words []string) string {
	t1, _ := template.New("suggestedWords").Parse(HtmlTemplateSuggestedWords)
	var buf bytes.Buffer
	err := t1.Execute(&buf, words)
	if err != nil {
		return err.Error()
	}
	return buf.String()
}
