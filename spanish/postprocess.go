package spanish

import (
	"github.com/olegator77/snowball/snowballword"
)

// Applies transformations necessary after
// a word has been completely processed.
//
func postprocess(word *snowballword.SnowballWord) {

	removeAccuteAccents(word)
}
