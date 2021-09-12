package data

import (
	"errors"
	"io"
	"quran-indonesia/models"

	"github.com/jszwec/csvutil"
)

//handle quran csv
func handleQuran(dec *csvutil.Decoder, filter func(interface{}) (interface{}, bool, bool)) (interface{}, error) {
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
			slices = append(slices, item.(models.Quran))
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

//handle list csv
func handleList(dec *csvutil.Decoder, filter func(interface{}) (interface{}, bool, bool)) (interface{}, error) {
	slices := []models.List{}

	for {
		var unit models.List

		err := dec.Decode(&unit)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err.Error())
		}

		item, valid, isItFinish := filter(unit)

		if valid {
			slices = append(slices, item.(models.List))
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
