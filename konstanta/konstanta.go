package konstanta

import "fmt"

var AYA = "aya"
var SURA = "sura"
var TAKE = "take"
var TAKEVALUE = "takevalue"
var ALL = "all"
var FROM = "from"
var FROMVALUE = "fromvalue"

func TakeFromPath() string {
	return fmt.Sprintf("/%s/:%s/%s/:%s/%s/:%s", AYA, SURA, TAKE, TAKEVALUE, FROM, FROMVALUE)
}

func AllPath() string {
	return fmt.Sprintf("//%s/:%s/%s", AYA, SURA, ALL)
}

func ListSuraPath() string {
	return fmt.Sprintf("/%s", SURA)
}
