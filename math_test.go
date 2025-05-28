package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)
	if total != 30 {
		t.Errorf("Resultado %d da soma é invalido. O Resultado esperado é %d", total, 30)
	}
}
