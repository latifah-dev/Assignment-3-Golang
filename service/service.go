package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"Tugas3/helper"
)

func UpdateJSON(disaster helper.Disaster) {
	/*
		update (write) the JSON file
	*/
	const FilePath = "./data/data.json"

	disaster.Status.Water = helper.RandomNumber(1, 100)
	disaster.Status.Wind = helper.RandomNumber(1, 100)

	disasterByte, err := json.Marshal(disaster)

	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(FilePath, disasterByte, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", disaster)
}
