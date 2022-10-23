package helper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
)

type Disaster struct {
	Status Status `json:"status"`
}

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func RandomNumber(min, max int) int {
	return min + rand.Intn(max-min)
}

func WindStatus(number int) string {
	disasterType := "Wind "
	desc := fmt.Sprint(number, " meter per second")
	if number < 5 {
		return disasterType + "(Safe): " + desc
	} else if number >= 6 && number < 10 {
		return disasterType + "(Alert): " + desc
	} else {
		return disasterType + "(Dangerous): " + desc
	}
}

func WaterStatus(number int) string {
	disasterType := "Water "
	desc := fmt.Sprint(number, " meter")
	if number < 5 {
		return disasterType + "(Safe): " + desc
	} else if number >= 6 && number < 10 {
		return disasterType + "(Alert): " + desc
	} else {
		return disasterType + "(Dangerous): " + desc
	}
}

func GetDisaster() Disaster {
	fileBytes, err := ioutil.ReadFile("./data/data.json")

	if err != nil {
		panic(err)
	}

	var data Disaster

	err = json.Unmarshal(fileBytes, &data)

	if err != nil {
		panic(err)
	}

	return data
}
