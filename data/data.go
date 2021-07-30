package data

import (
	_ "embed"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"quran-indonesia/models"

	"github.com/jszwec/csvutil"
)

func csvFilepath(name string) string {
	return filepath.Join("data", name)
}

var QURAN = csvFilepath("quran.csv")

func openFile(filepath string) *os.File {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	return file
}

func unmarshall(fname string, filter func(models.Quran) (models.Quran, bool, bool)) (interface{}, error) {

	file := openFile(fname)

	csvReader := csv.NewReader(file)
	csvReader.Comma = ','

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	result, err := doLoopingWithFilter(dec, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//loop the csv with filter parameter
func doLoopingWithFilter(dec *csvutil.Decoder, filter func(models.Quran) (models.Quran, bool, bool)) ([]models.Quran, error) {
	slices := []models.Quran{}

	for {
		var unit models.Quran

		err := dec.Decode(&unit)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err.Error())
		}

		item, valid, isItFinish := filter(unit)

		if valid {
			slices = append(slices, item)
		}
		if isItFinish {
			break
		}
	}

	if len(slices) == 0 {
		return nil, errors.New("NO DATA")
	}

	return slices, nil
}

func UnmarshallQuranCSV(filter func(models.Quran) (models.Quran, bool, bool)) ([]models.Quran, error) {
	result, err := unmarshall(QURAN, filter)
	if err != nil {
		return nil, err
	}
	return result.([]models.Quran), nil
}
