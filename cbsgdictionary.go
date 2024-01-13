package cbsg

import (
	"bufio"
	_ "embed"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

//const DefaultDictionaryFile = "dict/en/cbsg_dictionary.csv"

//go:embed dict/en/cbsg_dictionary.csv
var defaultDictionary string

type CbsgDictionary struct {
	sentenceCache map[string]map[string]int
	mu            sync.RWMutex
}

func NewCbsgDictionary() CbsgDictionary {
	dict := CbsgDictionary{
		sentenceCache: make(map[string]map[string]int),
	}
	myReader := strings.NewReader(defaultDictionary)
	dict.loadDictionaryFromReader(myReader)
	return dict
}

func NewCustomCbsgDictionary(dictionaryFile string) CbsgDictionary {
	dict := CbsgDictionary{
		sentenceCache: make(map[string]map[string]int),
	}
	dict.loadDictionary(dictionaryFile)
	return dict
}

func (cd *CbsgDictionary) SentenceWithWeight(resourceKey string) map[string]int {
	cd.mu.RLock()
	defer cd.mu.RUnlock()
	if cd.sentenceCache == nil {
		return make(map[string]int)
	}
	return cd.sentenceCache[resourceKey]
}

func (cd *CbsgDictionary) loadDictionary(dictionaryFile string) {
	file, err := os.Open(dictionaryFile)
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()
	cd.loadDictionaryFromReader(file)
}

func (cd *CbsgDictionary) loadDictionaryFromReader(readIo io.Reader) {
	scanner := bufio.NewScanner(readIo)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		l := strings.Split(line, ",")
		if len(l) != 3 {
			continue
		}
		key := l[0]
		weight, err := strconv.Atoi(l[1])
		if err != nil {
			continue
		}
		value := l[2]

		cd.mu.Lock()
		if cd.sentenceCache[key] == nil {
			cd.sentenceCache[key] = make(map[string]int)
		}
		cd.sentenceCache[key][value] = weight
		cd.mu.Unlock()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
