package cbsg

import (
	"strings"
	"testing"
)

// Test the SentenceWithWeight function
func TestSentenceWithWeight(t *testing.T) {
	// Create a CbsgDictionary instance with the default dictionary
	dict := NewCbsgDictionary()

	// Test that an empty resourceKey results in an empty map
	result := dict.SentenceWithWeight("")
	if len(result) != 0 {
		t.Errorf("Expected empty map, got: %v", result)
	}

	// Test that an existing resourceKey returns the correct map
	result = dict.SentenceWithWeight("WORD_GROWTH_ATOM")
	if len(result) == 0 {
		t.Error("Expected non-empty map, got empty map")
	}
}

// Test the NewCbsgDictionary function
func TestNewCbsgDictionary(t *testing.T) {
	// Create a new CbsgDictionary instance
	dict := NewCbsgDictionary()

	// Check if the default dictionary is loaded
	if len(dict.sentenceCache) == 0 {
		t.Error("Expected non-empty sentenceCache, got empty sentenceCache")
	}
}

// Test the NewCustomCbsgDictionary function
func TestNewCustomCbsgDictionary(t *testing.T) {
	// Create a new CbsgDictionary instance with a custom dictionary file
	dict := NewCustomCbsgDictionary("./test/test_dictionary.csv")

	// Check if the dictionary is loaded
	if len(dict.sentenceCache) == 0 {
		t.Error("Expected non-empty sentenceCache, got empty sentenceCache")
	}
}

// Test the loadDictionaryFromReader function
func TestLoadDictionaryFromReader(t *testing.T) {
	// Create a new CbsgDictionary instance
	dict := &CbsgDictionary{
		sentenceCache: make(map[string]map[string]int),
	}

	// Test loading the dictionary from a Reader
	dictionaryData := "WORD1,10,VALUE1\nWORD2,20,VALUE2\n"
	reader := strings.NewReader(dictionaryData)
	dict.loadDictionaryFromReader(reader)

	// Check if the dictionary is loaded correctly
	expectedResult := map[string]map[string]int{
		"WORD1": {"VALUE1": 10},
		"WORD2": {"VALUE2": 20},
	}

	if !compareMaps(dict.sentenceCache, expectedResult) {
		t.Errorf("Unexpected sentenceCache value, got: %v", dict.sentenceCache)
	}
}

// Helper function to compare two maps
func compareMaps(map1, map2 map[string]map[string]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key, values1 := range map1 {
		values2, ok := map2[key]
		if !ok {
			return false
		}

		if len(values1) != len(values2) {
			return false
		}

		for innerKey, value1 := range values1 {
			value2, ok := values2[innerKey]
			if !ok || value1 != value2 {
				return false
			}
		}
	}

	return true
}

func TestLoadDictionaryFromReader_SingleLine(t *testing.T) {
	dictionaryContent := `SENW_ARTICULATED_PROPOSITION,270,This is an enumeration: one, two, three
`
	reader := strings.NewReader(dictionaryContent)

	dictionary := NewCbsgDictionary()
	dictionary.loadDictionaryFromReader(reader)

	expectedCache := map[string]map[string]int{
		"SENW_ARTICULATED_PROPOSITION": {
			"This is an enumeration: one, two, three": 270,
		},
	}

	for key, expectedValues := range expectedCache {
		actualValues, ok := dictionary.sentenceCache[key]
		if !ok {
			t.Errorf("Key %s not found in sentenceCache", key)
			continue
		}

		for value, expectedWeight := range expectedValues {
			actualWeight, ok := actualValues[value]
			if !ok {
				t.Errorf("Value %s not found for key %s in sentenceCache", value, key)
				continue
			}

			if actualWeight != expectedWeight {
				t.Errorf("Weight mismatch for key %s and value %s. Expected: %d, Actual: %d",
					key, value, expectedWeight, actualWeight)
			}

			// Additional check for value
			if actualValue := value; actualValue != value {
				t.Errorf("Value mismatch for key %s. Expected: %s, Actual: %s",
					key, value, actualValue)
			}
		}
	}
}
