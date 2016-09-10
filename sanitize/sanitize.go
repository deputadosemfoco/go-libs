package sanitize

import (
	"bytes"
	"strings"
)

var transliterations = map[rune]string{
	'À': "A",
	'Á': "A",
	'Â': "A",
	'Ã': "A",
	'Ä': "A",
	'Å': "AA",
	'Æ': "AE",
	'Ç': "C",
	'È': "E",
	'É': "E",
	'Ê': "E",
	'Ë': "E",
	'Ì': "I",
	'Í': "I",
	'Î': "I",
	'Ï': "I",
	'Ð': "D",
	'Ł': "L",
	'Ñ': "N",
	'Ò': "O",
	'Ó': "O",
	'Ô': "O",
	'Õ': "O",
	'Ö': "O",
	'Ø': "OE",
	'Ù': "U",
	'Ú': "U",
	'Ü': "U",
	'Û': "U",
	'Ý': "Y",
	'Þ': "Th",
	'ß': "ss",
	'à': "a",
	'á': "a",
	'â': "a",
	'ã': "a",
	'ä': "a",
	'å': "aa",
	'æ': "ae",
	'ç': "c",
	'è': "e",
	'é': "e",
	'ê': "e",
	'ë': "e",
	'ì': "i",
	'í': "i",
	'î': "i",
	'ï': "i",
	'ð': "d",
	'ł': "l",
	'ñ': "n",
	'ń': "n",
	'ò': "o",
	'ó': "o",
	'ô': "o",
	'õ': "o",
	'ō': "o",
	'ö': "o",
	'ø': "oe",
	'ś': "s",
	'ù': "u",
	'ú': "u",
	'û': "u",
	'ū': "u",
	'ü': "u",
	'ý': "y",
	'þ': "th",
	'ÿ': "y",
	'ż': "z",
	'Œ': "OE",
	'œ': "oe",
}

// NoAccents removes all accents from string
func NoAccents(s string) string {
	b := bytes.NewBufferString("")

	for _, c := range s {
		if val, ok := transliterations[c]; ok {
			b.WriteString(val)
		} else {
			b.WriteRune(c)
		}
	}

	return b.String()
}

// CleanStr removes all special characters from a string
func CleanStr(s string) string {
	s = strings.Trim(s, " ")
	s = NoAccents(s)
	s = strings.Replace(s, "'", "", -1)
	s = strings.Replace(s, "|", "", -1)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, "&", "", -1)
	s = strings.Replace(s, "¨", "", -1)

	return s
}

// Capitalize upper cases the first letters of each word from a string
func Capitalize(str string) string {
	str = strings.Trim(str, " ")

	if str == "" {
		return ""
	}

	arr := strings.Split(str, " ")
	b := bytes.NewBufferString("")

	for _, word := range arr {
		if len(word) == 1 {
			b.WriteString(strings.ToUpper(word))
		} else {
			b.WriteString(strings.ToUpper(word[0:1]))
			b.WriteString(strings.ToLower(word[1:len(word)]))
		}

		b.WriteString(" ")
	}

	return strings.Trim(b.String(), " ")
}
