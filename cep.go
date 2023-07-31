package cep

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

    "github.com/jedib0t/go-pretty/v6/table"
)

type Address struct {
    Cep         string `json:"cep"`
    Logradouro  string `json:"logradouro"`
    Complemento string `json:"complemento"`
    Bairro      string `json:"bairro"`
    Localidade  string `json:"localidade"`
    Uf          string `json:"uf"`
}

func NewAddress(cep string) *Address {
    address := &Address{}
    jsonData := address.getAddressJsonString(cep)
    err := json.Unmarshal([]byte(jsonData), address)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v", err)
        os.Exit(1)
    }
    return address
}

func (a Address) getAddressJsonString(cep string) string {
    url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
    res, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v", err) 
        os.Exit(1)
    }
    defer res.Body.Close()
    if res.StatusCode != http.StatusOK {
        fmt.Println("Cannot get address") 
        os.Exit(1)
    }

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v", err) 
        os.Exit(1)
    }
    return string(body)
} 


func (a Address) GetLogradouro() string {
    return a.Logradouro
}

func (a Address) GetComplemento() string {
    return a.Complemento
}

func (a Address) GetBairro() string {
    return a.Bairro
}

func (a Address) GetLocalidade() string {
    return a.Localidade
}

func (a Address) GetUf() string {
    return a.Uf
}

func (a Address) Show() {
    t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
    t.SetStyle(table.StyleLight)
    t.AppendHeader(table.Row{"CEP", a.Cep})
    t.AppendRows([]table.Row{
        {"Logradouro", a.Logradouro},
        {"Complemento", a.Complemento},
        {"Bairro", a.Bairro},
        {"Localidade", a.Localidade},
        {"UF", a.Uf},
    })
    t.Render()
}

