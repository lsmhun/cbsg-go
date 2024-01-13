package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lsmhun/cbsg-go"
)

const (
	WORKSHOP                   = "--workshop"
	SHORT_WORKSHOP             = "--shortWorkshop"
	FINANCIAL_REPORT           = "--financialReport"
	SENTENCE_GUARANTEED_AMOUNT = "--sentenceGuaranteedAmount="
	DICTIONARY_FILE            = "--dictionaryFile="
	HELP                       = "--help"
	HELP_TEXT                  = "Available options:\n" +
		WORKSHOP + "\n" +
		SHORT_WORKSHOP + "\n" +
		FINANCIAL_REPORT + "\n" +
		SENTENCE_GUARANTEED_AMOUNT + "<ANY_INTEGER>\n\n" +
		DICTIONARY_FILE + "<DICTIONARY_FILE>\n" +
		HELP + "\n"
)

func main() {
	// Corporate Bullshit Generator
	cbsgDictionary := loadCbsgDictionary(os.Args)
	cbsg := cbsg.NewCbsgCore(cbsgDictionary)
	text := getCbsText(cbsg, os.Args)
	fmt.Println(text)
}

func getCbsText(cbsg cbsg.CbsgCore, args []string) string {
	text := ""
	for _, arg := range args {
		if strings.HasPrefix(arg, SENTENCE_GUARANTEED_AMOUNT) {
			amount, err := strconv.Atoi(arg[len(SENTENCE_GUARANTEED_AMOUNT):])
			if err == nil {
				text = cbsg.SentenceGuaranteedAmount(amount)
				break
			}
		}
		switch arg {
		case WORKSHOP:
			text = cbsg.Workshop()
		case SHORT_WORKSHOP:
			text = cbsg.ShortWorkshop()
		case FINANCIAL_REPORT:
			text = cbsg.FinancialReport()
		case HELP:
			text = HELP_TEXT
		}
	}
	if text == "" {
		return cbsg.ShortWorkshop()
	}
	return text
}

func loadCbsgDictionary(args []string) cbsg.CbsgDictionary {
	var dictionaryFile string
	for _, arg := range args {
		if strings.HasPrefix(arg, DICTIONARY_FILE) {
			dictionaryFile = arg[len(DICTIONARY_FILE):]
			break
		}
	}
	if dictionaryFile == "" {
		return cbsg.NewCbsgDictionary()
	}

	return cbsg.NewCustomCbsgDictionary(dictionaryFile)
}
