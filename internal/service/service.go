package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func DefinOfContent(text string) string {
	var result string
	

	if strings.ContainsAny(text, "ЙЦУКЕНГШЩЗХЪЭЖДЛОРПАВЫФЯЧСМИТЬБЮ") {
		result = morse.ToMorse(text)
		return result
	}
	
	result = morse.ToText(text)
	return result 
}