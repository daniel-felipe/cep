package cep_test

import (
	"testing"
    "daniel/cep"
)

var address = cep.NewAddress("01454000")


func TestGetLogradouro(t *testing.T) {
    exp := "Avenida Cidade Jardim"

    res := address.GetLogradouro()

    if res != exp {
        t.Errorf("Expected %s, got %s instead.", exp, res)
    }
}

func TestGetComplemento(t *testing.T) {
    exp := "lado par"

    res := address.GetComplemento()

    if res != exp {
        t.Errorf("Expected %s, got %s instead.", exp, res)
    }
}

func TestGetBairro(t *testing.T) {
    exp := "Jardim Paulistano"

    res := address.GetBairro()

    if res != exp {
        t.Errorf("Expected %s, got %s instead.", exp, res)
    }
}

func TestGetLocalidade(t *testing.T) {
    exp := "São Paulo"

    res := address.GetLocalidade()

    if res != exp {
        t.Errorf("Expected %s, got %s instead.", exp, res)
    }
}

func TestGetUf(t *testing.T) {
    exp := "SP"

    res := address.GetUf()

    if res != exp {
        t.Errorf("Expected %s, got %s instead.", exp, res)
    }
}

