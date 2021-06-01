package spanish

import (
	"github.com/inteli-ai-dev/snowball/snowballword"
)

// Applies transformations necessary after
// a word has been completely processed.
//
func postprocess(word *snowballword.SnowballWord) {

	removeAccuteAccents(word)
}
