package models

type Quran struct {
	ID        string `json:"id" csv:"id"`
	Sura      string `json:"sura" csv:"sura"`
	Aya       string `json:"aya" csv:"aya"`
	Indonesia string `json:"indonesia" csv:"indonesia"`
	Footnotes string `json:"footnotes" csv:"footnotes"`
	Arabic    string `json:"arabic" csv:"arabic"`
}
