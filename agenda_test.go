package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddContato(t *testing.T) {

	Agenda := map[string]string{}

	nome := "João"
	tel := "1199876-9087"

	AddContato(Agenda, nome, tel)

	assert.NotNil(t, Agenda)
	assert.Equal(t, tel, Agenda["João"])
}
