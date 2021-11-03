package pokemon

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Summary struct {
	Count int `json:"count"`
}

type Species struct {
	Name string `json:"name"`
}

type Pokemon struct {
	Id int `json:"id"`
	Weight int `json:"weight"`
	Species Species `json:"species"`
}

func GetTotalPokemon() int {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=2000&offset=0")
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	rawIn := json.RawMessage(body)
	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		panic(err)
	}

	var s Summary
	err = json.Unmarshal(bytes, &s)
	if err != nil {
		panic(err)
	}
	return s.Count
}


func GetById(id int) (*Pokemon, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + strconv.Itoa(id))
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rawIn := json.RawMessage(body)
	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		return nil, err
	}

	var p Pokemon
	err = json.Unmarshal(bytes, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
