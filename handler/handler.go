package handler

import (
	"Tugas3/helper"
	"html/template"
	"net/http"
	"path"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	var FilePath = path.Join("view", "index.html")

	var tmpl, err = template.ParseFiles(FilePath)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	disaster := helper.GetDisaster()
	numWater := disaster.Status.Water
	numWind := disaster.Status.Wind

	waterDesc := helper.WaterStatus(numWater)
	windDesc := helper.WindStatus(numWind)

	var statusDesc = map[string]interface{}{
		"waterDesc": waterDesc,
		"windDesc":  windDesc,
	}

	err = tmpl.Execute(w, statusDesc)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
