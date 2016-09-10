package sanitize

import "testing"

func TestCapitalizeMustAcceptEmptyAndEmptySpaceString(t *testing.T) {
	empty := Capitalize("")
	emptySpace := Capitalize("")

	if empty != "" {
		t.Fail()
	}

	if emptySpace != "" {
		t.Fail()
	}
}

func TestCapitalizeWithOneWord(t *testing.T) {
	actual := Capitalize("TESTE")
	expected := "Teste"

	if actual != expected {
		t.Fail()
	}
}

func TestCapitalizeWithMultipleWords(t *testing.T) {
	actual := Capitalize("TESTE DA SILVA SAURO")
	expected := "Teste Da Silva Sauro"

	if actual != expected {
		t.Fail()
	}
}
