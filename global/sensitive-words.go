package global

import (
	_ "embed"
	"github.com/importcjj/sensitive"
)

//go:embed sensitive-words.txt
var SensitiveWords string

var SensitiveWordsFilter *sensitive.Filter

func InitSensitiveWords() {
	SensitiveWordsFilter = sensitive.New()
}
