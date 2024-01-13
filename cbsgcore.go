package cbsg

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"unicode"
)

var (
	varRegexp = regexp.MustCompile(`\$[a-zA-Z\d]+`)
	randGen   = rand.New(rand.NewSource(42)) // Random generator with seed
)

type CbsgCore struct {
	// You might define necessary data structures or fields here
	dict CbsgDictionary
}

func NewCbsgCore(dictionary CbsgDictionary) CbsgCore {
	myCbsgCore := CbsgCore{
		dict: dictionary,
	}
	return myCbsgCore
}

func NewDefaultCbsgCore() CbsgCore {
	dictionary := NewCbsgDictionary()
	myCbsgCore := CbsgCore{
		dict: dictionary,
	}
	return myCbsgCore
}

func (c *CbsgCore) SentenceGuaranteedAmount(count int) string {
	var sb strings.Builder
	for i := 0; i < count; i++ {
		sb.WriteString(c.sentence())
		sb.WriteString(" ")
	}
	return c.cleanup(sb.String())
}

func (c *CbsgCore) Workshop() string {
	return c.SentenceGuaranteedAmount(500)
}

func (c *CbsgCore) ShortWorkshop() string {
	return c.SentenceGuaranteedAmount(5)
}

func (c *CbsgCore) FinancialReport() string {
	return c.sentences()
}

///////////////// TEXT ELEMENTS ///////////////////

func (c *CbsgCore) sentence() string {
	propositions := c.articulatedPropositions()
	if propositions == "" {
		return propositions
	}
	return strings.ToUpper(propositions[:1]) + propositions[1:] + "."
}

func (c *CbsgCore) sentences() string {

	var ret string
	pm := rand.Intn(10)
	limit := 30
	if rand.Intn(2) == 1 {
		limit = 30 + pm
	} else {
		limit = 30 - pm
	}

	cnt := 0
	for limit != cnt {
		ret += c.sentence() + " "
		cnt++
	}

	return c.cleanup(ret)
}

func (c *CbsgCore) cleanup(text string) string {
	// Implement the logic for cleaning up text
	// Similar to Java's cleanup method
	return strings.Join(strings.Fields(text), " ")
}

func (c *CbsgCore) makeEventualPlural(word string, plural bool) string {
	if len(word) < 3 || !plural {
		return word
	}

	// Check for parentheses abbreviation
	abbrIndex := strings.Index(word, " (")
	if abbrIndex != -1 {
		return c.makeEventualPlural(word[:abbrIndex], plural) + word[abbrIndex:]
	}

	// Handle specific word transformations
	switch word {
	case "matrix":
		return "matrices"
	case "analysis":
		return "analyses"
	}

	// Handle general plural forms
	last := len(word) - 1
	switch {
	case strings.HasSuffix(word, "gh"):
		return word + "s"
	case strings.HasSuffix(word, "s"), strings.HasSuffix(word, "x"), strings.HasSuffix(word, "z"), strings.HasSuffix(word, "h"):
		return word + "es"
	case strings.HasSuffix(word, "y") && !strings.HasSuffix(strings.ToLower(word), "ay") &&
		!strings.HasSuffix(strings.ToLower(word), "ey") &&
		!strings.HasSuffix(strings.ToLower(word), "iy") &&
		!strings.HasSuffix(strings.ToLower(word), "oy") &&
		!strings.HasSuffix(strings.ToLower(word), "uy"):
		return word[:last] + "ies"
	}

	return word + "s"
}

func (c *CbsgCore) buildPluralVerb(verb string, plural bool) string {
	last := len(strings.TrimSpace(verb)) - 1
	if plural || last < 0 {
		return verb
	}

	sExtension := verb[:last+1] + "s" + verb[last+1:]
	esExtension := verb[:last+1] + "es" + verb[last+1:]

	switch verb[last] {
	case 's', 'o', 'z':
		return esExtension
	case 'h':
		if last > 0 && (verb[last-1] == 'c' || verb[last-1] == 's') {
			return esExtension
		}
		return sExtension
	case 'y':
		if last > 0 && unicode.Is(unicode.Latin, rune(verb[last-1])) {
			return sExtension
		}
		return verb[:last] + "ies" + verb[last+1:]
	}

	return sExtension
}

func (c *CbsgCore) sillyAbbreviationGeneratorSAS(s string) string {
	words := strings.Split(s, " ")
	var abbreviation strings.Builder

	for _, word := range words {
		abbreviation.WriteString(string(word[0]))
	}

	return abbreviation.String()
}

func (c *CbsgCore) abbreviate(s string, probability float64) string {
	if rand.Float64() < probability {
		return s + " (" + c.sillyAbbreviationGeneratorSAS(s) + ")"
	}
	return s
}

func (c *CbsgCore) weightedChoice(key string) string {
	return c.weightedChoiceFromMap(c.dict.SentenceWithWeight(key))
}

func (c *CbsgCore) weightedChoiceFromMap(choices map[string]int) string {
	totalWeight := 0
	for _, weight := range choices {
		totalWeight += weight
	}
	if totalWeight == 0 {
		return ""
	}

	randomWeight := rand.Intn(totalWeight)
	for choice, weight := range choices {
		randomWeight -= weight
		if randomWeight < 0 {
			return choice
		}
	}

	return ""
}

func (c *CbsgCore) addIndefiniteArticle(word string, plural bool) string {
	if plural || len(word) == 0 {
		return word
	}

	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, v := range vowels {
		if rune(word[0]) == v {
			return "an " + word
		}
	}

	return "a " + word
}

func (c *CbsgCore) evaluateValues(template string) string {
	var values []interface{}
	matches := varRegexp.FindAllString(template, -1)

	for _, match := range matches {
		// It will evaluate $thingAtom to its values
		values = append(values, c.templateFunction(match))
	}

	templateReplace := varRegexp.ReplaceAllString(template, "%s")
	return fmt.Sprintf(templateReplace, values...)
}

func (c *CbsgCore) person() string {
	personTemplate := c.weightedChoice(SENW_PERSON)
	if personTemplate == "" {
		return c.weightedChoice(WORD_PERSON)
	}
	return c.evaluateValues(personTemplate)
}

func (c *CbsgCore) personPlural() string {
	return c.weightedChoice(WORD_PERSON_PLURAL)
}

func (c *CbsgCore) boss() string {

	department := c.weightedChoice(WORD_BOSS_DEPARTMENT)
	departmentOrTopRole := c.weightedChoice(SENW_BOSS_DEPT)

	if rand.Intn(4) == 1 {
		managing := c.weightedChoice(SENW_BOSS_MANAGING)
		vice := c.weightedChoice(SENW_BOSS_VICE)
		co := c.weightedChoice(SENW_BOSS_CO)
		title := fmt.Sprintf("%s %s %s", c.weightedChoice(SENW_BOSS_TITLE), co, vice)

		age := c.weightedChoice(SENW_BOSS_AGE)
		executive := c.weightedChoice(SENW_BOSS_EXECUTIVE)
		return managing + age + executive + title + " of " + department
	}

	groupal := c.weightedChoice(SENW_BOSS_GROUPAL)
	officerOrCatalyst := c.weightedChoice(SENW_BOSS_OFFICER)
	return groupal + c.abbreviate("Chief "+departmentOrTopRole+" "+
		officerOrCatalyst, 0.6)
}

func (c *CbsgCore) timelessEvent() string {
	return c.weightedChoice(WORD_TIMELESS_EVENT)
}

func (c *CbsgCore) growthAtom() string {
	return c.weightedChoice(WORD_GROWTH_ATOM)
}

func (c *CbsgCore) growth() string {
	superlative := c.weightedChoice(WORD_GROWTH)
	return superlative + " " + c.growthAtom()
}

// Define functions corresponding to each Java method
func (c *CbsgCore) innerPersonVerbHavingThingComplement() string {
	return c.weightedChoice(WORD_PERSON_INNER_HAVING_THING_COMPLEMENT)
}

func (c *CbsgCore) personVerbHavingThingComplement(plural bool, infinite bool) string {
	if infinite {
		innerPerson := c.innerPersonVerbHavingThingComplement()
		return c.buildPluralVerb(innerPerson, plural)
	}
	return "" // Return appropriate default value or handle the case
}

func (c *CbsgCore) personVerbHavingBadThingComplement(plural bool) string {
	inner := c.weightedChoice(WORD_PERSON_HAVING_BAD_THING_COMPLEMENT)
	return c.buildPluralVerb(inner, plural)
}

func (c *CbsgCore) personVerbHavingPersonComplement(plural bool) string {
	inner := c.weightedChoice(WORD_PERSON_HAVING_PERSON_COMPLEMENT)
	return c.buildPluralVerb(inner, plural)
}

func (c *CbsgCore) thingVerbAndDefiniteEnding(plural bool) string {
	inner := c.weightedChoice(WORD_THING_VERB_DEFINITE_ENDING)
	return c.buildPluralVerb(inner, plural)
}

func (c *CbsgCore) thingVerbHavingThingComplement(plural bool) string {
	inner := c.weightedChoice(WORD_THING_VERB_HAVING_COMPLEMENT)
	return c.buildPluralVerb(inner, plural)
}

func (c *CbsgCore) thingVerbAndEnding() string {
	weightedVerbAndEnding := c.weightedChoice(SENW_THING_VERB_ENDING)
	if weightedVerbAndEnding == "" {
		return c.thingVerbAndDefiniteEnding(false)
	}
	return c.evaluateValues(weightedVerbAndEnding)
}

func (c *CbsgCore) thingVerbAndEndingPlural() string {
	weightedVerbAndEnding := c.weightedChoice(SENW_THING_VERB_ENDING_PLURAL)
	if weightedVerbAndEnding == "" {
		return c.thingVerbAndDefiniteEnding(true)
	}
	return c.evaluateValues(weightedVerbAndEnding)
}

func (c *CbsgCore) personVerbAndEnding(plural bool, infinitive bool) string {
	complSP := rand.Intn(2) == 1
	r := rand.Intn(95)
	if r < 10 {
		return c.personVerbAndDefiniteEnding(plural, infinitive)
	} else if r < 15 {
		return c.personVerbHavingBadThingComplement(plural) + " " +
			c.addRandomArticle(c.badThings(), plural)
	}
	complement := c.addRandomArticle(c.thing(), false)
	if complSP {
		complement = c.addRandomArticle(c.thingPlural(), true)
	}
	return c.personVerbHavingThingComplement(plural, infinitive) + " " + complement
}

func (c *CbsgCore) thingInner() string {
	weightedThingInner := c.weightedChoice(SENW_THING_INNER)
	if weightedThingInner == "" {
		return c.weightedChoice(WORD_THING_INNER)
	}
	res := c.evaluateValues(weightedThingInner)
	senwOrg := c.dict.SentenceWithWeight(SENW_ORG)
	for org := range senwOrg {
		if strings.Contains(res, org) {
			return res
		}
	}
	return c.abbreviate(res, 0.5)
}

func (c *CbsgCore) matrixOrSO() string {
	return c.weightedChoice(SENW_ORG)
}

func (c *CbsgCore) thingAtom() string {
	weightedThingAtom := c.weightedChoice(SENW_THING_ATOM)
	if weightedThingAtom == "" {
		return c.weightedChoice(WORD_THING_ATOM)
	}
	if strings.Contains(weightedThingAtom, "$") {
		return c.evaluateValues(weightedThingAtom)
	}
	return c.abbreviate(weightedThingAtom, 0.5)
}

func (c *CbsgCore) thingAtomPlural() string {
	weightedThingAtom := c.weightedChoice(SENW_THING_ATOM_PLURAL)
	if weightedThingAtom == "" {
		return c.weightedChoice(WORD_THING_ATOM_PLURAL)
	}
	return c.makeEventualPlural(c.evaluateValues(weightedThingAtom), true)
}

func (c *CbsgCore) badThings() string {
	return c.weightedChoice(WORD_BAD_THINGS)
}

func (c *CbsgCore) thingAdjective() string {
	return c.weightedChoice(WORD_THING_ADJECTIVE)
}

func (c *CbsgCore) eventualAdverb() string {
	if rand.Intn(4) == 1 {
		return c.weightedChoice(WORD_ADVERB_EVENTUAL)
	}
	return ""
}

func (c *CbsgCore) thing() string {
	weightedThing := c.weightedChoice(SENW_THING)
	return c.evaluateValues(weightedThing)
}

func (c *CbsgCore) thingPlural() string {
	weightedThing := c.weightedChoice(SENW_THING_PLURAL)
	return c.evaluateValues(weightedThing)
}

func (c *CbsgCore) addRandomArticle(word string, plural bool) string {
	weightedRandomArticle := c.weightedChoice(SENW_ADD_RANDOM_ARTICLE)
	if weightedRandomArticle == "" {
		return c.addIndefiniteArticle(word, plural)
	}
	return c.evaluateValues(fmt.Sprintf(weightedRandomArticle, word))
}

func (c *CbsgCore) faukon() string {
	weightedProposition := c.weightedChoice(SENW_FAUKON)
	if weightedProposition == "" {
		return c.weightedChoice(WORD_FAUKON)
	}
	return c.evaluateValues(weightedProposition)
}

func (c *CbsgCore) eventualPostfixedAdverb() string {
	weightedProposition := c.weightedChoice(SENW_EVENTUAL_POSTFIXED_ADVERB)
	if weightedProposition == "" {
		return c.weightedChoice(WORD_ADVERB_EVENTUAL_POSTFIXED)
	}
	return c.evaluateValues(weightedProposition)
}

func (c *CbsgCore) personVerbAndDefiniteEnding(plural, infinitive bool) string {
	weightedProposition := c.weightedChoice(SENW_PERSON_VERB_AND_DEFINITE_ENDING)
	inner := weightedProposition
	if weightedProposition == "" {
		inner = c.weightedChoice(WORD_PERSON_VERB_DEFINITE_ENDING)
	} else {
		inner = c.evaluateValues(weightedProposition)
	}

	if infinitive {
		return inner
	}
	return c.buildPluralVerb(inner, plural)
}

func (c *CbsgCore) proposition() string {
	weightedProposition := c.weightedChoice(SENW_PROPOSITION)
	return c.evaluateValues(weightedProposition)
}

func (c *CbsgCore) articulatedPropositions() string {
	weightedProposition := c.weightedChoice(SENW_ARTICULATED_PROPOSITION)
	return c.evaluateValues(weightedProposition)
}

func (c *CbsgCore) templateFunction(templateName string) string {
	var result string

	switch templateName {
	case "$faukon":
		result = c.faukon()
	case "$sentence":
		result = c.sentence()
	case "$thing":
		result = c.thing()
	case "$thingPlural":
		result = c.thingPlural()
	case "$thingRandom":
		if rand.Intn(2) == 0 {
			result = c.thing()
		} else {
			result = c.thingPlural()
		}
	case "$thingInner":
		result = c.thingInner()
	case "$thingAdjective":
		result = c.thingAdjective()
	case "$thingAtom":
		result = c.thingAtom()
	case "$thingAtomRandom":
		if rand.Intn(2) == 0 {
			result = c.thingAtom()
		} else {
			result = c.thingAtomPlural()
		}
	case "$thingAtomPlural":
		result = c.thingAtomPlural()
	case "$thingVerbAndEnding":
		result = c.thingVerbAndEnding()
	case "$thingVerbAndEndingPlural":
		result = c.thingVerbAndEndingPlural()
	case "$thingVerbHavingThingComplement":
		result = c.thingVerbHavingThingComplement(false)
	case "$thingVerbHavingThingComplementPlural":
		result = c.thingVerbHavingThingComplement(true)
	case "$thingWithRandomArticle":
		result = c.addRandomArticle(c.thing(), false)
	case "$thingWithRandomArticlePlural":
		result = c.addRandomArticle(c.thingPlural(), true)
	case "$thingWithRandomArticleRandom":
		if rand.Intn(2) == 0 {
			result = c.addRandomArticle(c.thing(), false)
		} else {
			result = c.addRandomArticle(c.thingPlural(), true)
		}
	case "$badThingWithRandomArticle":
		result = c.addRandomArticle(c.badThings(), false)
	case "$badThingWithRandomArticlePlural":
		result = c.addRandomArticle(c.badThings(), true)
	case "$person":
		result = c.person()
	case "$personPlural":
		result = c.personPlural()
	case "$personRandom":
		if rand.Intn(2) == 0 {
			result = c.person()
		} else {
			result = c.personPlural()
		}
	case "$personInfinitiveVerbAndEnding":
		result = c.personVerbAndEnding(true, true)
	case "$personVerbAndEnding":
		result = c.personVerbAndEnding(false, false)
	case "$personVerbHavingPersonComplement":
		result = c.personVerbHavingPersonComplement(false)
	case "$personVerbHavingPersonComplementPlural":
		result = c.personVerbHavingPersonComplement(true)
	case "$personVerbHavingThingComplement":
		result = c.personVerbHavingThingComplement(false, false)
	case "$personVerbHavingBadThingComplement":
		result = c.personVerbHavingBadThingComplement(false)
	case "$personVerbHavingBadThingComplementPlural":
		result = c.personVerbHavingBadThingComplement(true)
	case "$boss":
		result = c.boss()
	case "$eventualAdverb":
		result = c.eventualAdverb()
	case "$eventualPostfixedAdverb":
		result = c.eventualPostfixedAdverb()
	case "$growthAtom":
		result = c.growthAtom()
	case "$addIndefiniteArticleGrowth":
		result = c.addIndefiniteArticle(c.growth(), false)
	case "$addIndefiniteArticleGrowthPlural":
		result = c.addIndefiniteArticle(c.growth(), true)
	case "$addIndefiniteArticleThing":
		result = c.addIndefiniteArticle(c.thing(), false)
	case "$addIndefiniteArticleThingPlural":
		result = c.addIndefiniteArticle(c.thingPlural(), true)
	case "$addIndefiniteArticleThingRandom":
		if rand.Intn(2) == 0 {
			result = c.addIndefiniteArticle(c.thing(), false)
		} else {
			result = c.addIndefiniteArticle(c.thingPlural(), true)
		}
	case "$proposition":
		result = c.proposition()
	case "$matrixOrSOPlural":
		result = c.makeEventualPlural(c.matrixOrSO(), true)
	case "$matrixOrSO":
		result = c.matrixOrSO()
	case "$timelessEvent":
		result = c.timelessEvent()
	case "$personVerbAndDefiniteEnding":
		result = c.personVerbAndDefiniteEnding(false, false)
	case "$personVerbAndDefiniteEndingInf":
		result = c.personVerbAndDefiniteEnding(false, true)
	case "$personVerbAndDefiniteEndingPlural":
		result = c.personVerbAndDefiniteEnding(true, false)
	case "$personVerbAndDefiniteEndingPluralInf":
		result = c.personVerbAndDefiniteEnding(true, true)
	}

	return result
}
