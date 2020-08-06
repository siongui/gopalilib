// Package dicmgr provides high-level methods to access dictionary data.
package dicmgr

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

const HtmlTemplateSuggestedWords = `
{{range $word := .}}
<div>{{$word}}</div>
{{end}}`
