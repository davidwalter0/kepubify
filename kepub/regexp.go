package kepub

import (
	"regexp"
)

var sentenceStart *regexp.Regexp
var quoteAttrStart *regexp.Regexp
var quoteEnd *regexp.Regexp
var quoteStart *regexp.Regexp
var err error
var sentenceStartString = `([.?,!;:>])([A-Za-z])`
var quoteAttrStartString = `([A-Za-z.?,!;:])(<)`
var quoteEndString = `([.?,!;:])[”"]`
var quoteStartString = `>(“)`
var replaceSentenceStart = []byte(`$1 $2`)
var replaceAttrStart = []byte(`$1 $2`)
var replaceEnd = []byte(`$1” `)
var replaceStart = []byte(`> "`)

func AddRegexpSpace(text []byte) []byte {
	sentenceStart = regexp.MustCompile(sentenceStartString)
	quoteAttrStart = regexp.MustCompile(quoteAttrStartString)
	quoteEnd = regexp.MustCompile(quoteEndString)
	quoteStart = regexp.MustCompile(quoteStartString)
	text = quoteEnd.ReplaceAll(text, replaceEnd)
	text = quoteStart.ReplaceAll(text, replaceStart)
	text = quoteAttrStart.ReplaceAll(text, replaceAttrStart)
	text = sentenceStart.ReplaceAll(text, replaceSentenceStart)
	return text
}
