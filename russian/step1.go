package russian

import (
	"github.com/olegator77/snowball/snowballword"
	// "log"
)

var ruSuff0 = strs2runes("ся", "сь")

var ruSuff1 = strs2runes("иями", "ями", "иях", "иям", "ием", "ией", "ами", "ях",
	"ям", "ья", "ью", "ье", "ом", "ой", "ов", "ия", "ию",
	"ий", "ии", "ие", "ем", "ей", "еи", "ев", "ах", "ам",
	"я", "ю", "ь", "ы", "у", "о", "й", "и", "е", "а")

// Step 1 is the removal of standard suffixes, all of which must
// occur in RV.
//
//
// Search for a PERFECTIVE GERUND ending. If one is found remove it, and
// that is then the end of step 1. Otherwise try and remove a REFLEXIVE
// ending, and then search in turn for (1) an ADJECTIVAL, (2) a VERB or
// (3) a NOUN ending. As soon as one of the endings (1) to (3) is found
// remove it, and terminate step 1.
//
func step1(word *snowballword.SnowballWord) bool {

	// `stop` will be used to signal early termination
	var stop bool

	// Search for a PERFECTIVE GERUND ending
	stop = removePerfectiveGerundEnding(word)
	if stop {
		return true
	}

	// Next remove reflexive endings
	word.RemoveFirstSuffixInR(word.RVstart, ruSuff0)

	// Next remove adjectival endings
	stop = removeAdjectivalEnding(word)
	if stop {
		return true
	}

	// Next remove verb endings
	stop = removeVerbEnding(word)
	if stop {
		return true
	}

	// Next remove noun endings
	suffix, _ := word.RemoveFirstSuffixInR(word.RVstart, ruSuff1)
	if suffix != "" {
		return true
	}

	return false
}

var ruSuff2 = strs2runes("ившись", "ывшись", "вшись", "ивши", "ывши", "вши", "ив", "ыв", "в")

// Remove perfective gerund endings and return true if one was removed.
//
func removePerfectiveGerundEnding(word *snowballword.SnowballWord) bool {
	suffix, suffixRunes := word.FirstSuffixInR(word.RVstart, len(word.RS), ruSuff2)
	switch suffix {
	case "в", "вши", "вшись":

		// These are "Group 1" perfective gerund endings.
		// Group 1 endings must follow а (a) or я (ia) in RV.
		if precededByARinRV(word, len(suffixRunes)) == false {
			suffix = ""
		}

	}

	if suffix != "" {
		word.RemoveLastNRunes(len(suffixRunes))
		return true
	}
	return false
}

// Remove adjectival endings and return true if one was removed.
//
func strs2runes(strings ...string) [][]rune {
	ret := make([][]rune, 0, len(strings))
	for _, str := range strings {
		ret = append(ret, []rune(str))
	}

	return ret
}

var ruSuff3 = strs2runes("ими", "ыми", "его", "ого", "ему", "ому", "ее", "ие",
	"ые", "ое", "ей", "ий", "ый", "ой", "ем", "им", "ым",
	"ом", "их", "ых", "ую", "юю", "ая", "яя", "ою", "ею")

var ruSuff4 = strs2runes("ивш", "ывш", "ующ", "ем", "нн", "вш", "ющ", "щ")

func removeAdjectivalEnding(word *snowballword.SnowballWord) bool {

	// Remove adjectival endings.  Start by looking for
	// an adjective ending.
	//
	suffix, _ := word.RemoveFirstSuffixInR(word.RVstart, ruSuff3)
	if suffix != "" {

		// We found an adjective ending.  Remove optional participle endings.
		//
		newSuffix, newSuffixRunes := word.FirstSuffixInR(word.RVstart, len(word.RS), ruSuff4)
		switch newSuffix {
		case "ем", "нн", "вш", "ющ", "щ":

			// These are "Group 1" participle endings.
			// Group 1 endings must follow а (a) or я (ia) in RV.
			if precededByARinRV(word, len(newSuffixRunes)) == false {
				newSuffix = ""
			}
		}

		if newSuffix != "" {
			word.RemoveLastNRunes(len(newSuffixRunes))
		}
		return true
	}
	return false
}

var ruSuff5 = strs2runes("уйте", "ейте", "ыть", "ыло", "ыли", "ыла", "уют", "ует",
	"нно", "йте", "ишь", "ить", "ите", "ило", "или", "ила",
	"ешь", "ете", "ены", "ено", "ена", "ят", "ют", "ыт", "ым",
	"ыл", "ую", "уй", "ть", "ны", "но", "на", "ло", "ли", "ла",
	"ит", "им", "ил", "ет", "ен", "ем", "ей", "ю", "н", "л", "й")

// Remove verb endings and return true if one was removed.
//
func removeVerbEnding(word *snowballword.SnowballWord) bool {
	suffix, suffixRunes := word.FirstSuffixInR(word.RVstart, len(word.RS), ruSuff5)
	switch suffix {
	case "ла", "на", "ете", "йте", "ли", "й", "л", "ем", "н",
		"ло", "но", "ет", "ют", "ны", "ть", "ешь", "нно":

		// These are "Group 1" verb endings.
		// Group 1 endings must follow а (a) or я (ia) in RV.
		if precededByARinRV(word, len(suffixRunes)) == false {
			suffix = ""
		}

	}

	if suffix != "" {
		word.RemoveLastNRunes(len(suffixRunes))
		return true
	}
	return false
}

// There are multiple classes of endings that must be
// preceded by а (a) or я (ia) in RV in order to be removed.
//
func precededByARinRV(word *snowballword.SnowballWord, suffixLen int) bool {
	idx := len(word.RS) - suffixLen - 1
	if idx >= word.RVstart && (word.RS[idx] == 'а' || word.RS[idx] == 'я') {
		return true
	}
	return false
}
