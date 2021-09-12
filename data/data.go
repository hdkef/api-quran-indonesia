package data

import (
	_ "embed"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"quran-indonesia/konstanta"
	"quran-indonesia/models"

	"github.com/jszwec/csvutil"
)

func csvFilepath(name string) string {
	return filepath.Join("data", name)
}

var QURAN = csvFilepath("quran.csv")
var LIST = csvFilepath("list.csv")

func openFile(filepath string) *os.File {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	return file
}

func unmarshall(unittype string, fname string, filter func(interface{}) (interface{}, bool, bool)) (interface{}, error) {

	file := openFile(fname)

	csvReader := csv.NewReader(file)
	csvReader.Comma = ','

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	result, err := doLoopingWithFilter(unittype, dec, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//loop the csv with filter parameter according to type
func doLoopingWithFilter(unittype string, dec *csvutil.Decoder, filter func(interface{}) (interface{}, bool, bool)) (interface{}, error) {
	switch unittype {
	case konstanta.AYA:
		return handleQuran(dec, filter)
	case konstanta.SURA:
		return handleList(dec, filter)
	default:
		return nil, errors.New("NO match type")
	}
}

func UnmarshallQuranCSV(filter func(interface{}) (interface{}, bool, bool)) ([]models.Quran, error) {
	result, err := unmarshall(konstanta.AYA, QURAN, filter)
	if err != nil {
		return nil, err
	}
	return result.([]models.Quran), nil
}

func UnmarshallListCSV() []models.List {
	result, err := unmarshall(konstanta.SURA, LIST, func(q interface{}) (interface{}, bool, bool) {
		return q, true, false
	})
	if err != nil {
		panic(err)
	}
	return result.([]models.List)
}
