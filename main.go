package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	update "github.com/inconshreveable/go-update"
)

type Contact struct {
	Email  string
	Github string
}

type Info struct {
	Version     string
	Description string
	Contact     Contact
}

func main() {
	programInfo := Info{
		"1.0.0",
		"This script checks for new version",
		Contact{
			"example@gmail.com",
			"http://github.com/example.com",
		},
	}

	needUpdate, err := needUpdate(programInfo)
	if err != nil {
		fmt.Println(err)
	}

	URL := "https://dl.equinox.io/code7unner/updater/stable"
	if needUpdate {
		fmt.Println("Start update")
		// call update function
		doUpdate(URL)
	} else {
		fmt.Println("Update not need")
	}
}

// Update binary file if version is updated
func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		return err
	}
	return err
}

// Check ersion of app with github sevice
func needUpdate(cerVer Info) (bool, error) {
	resp, err := http.Get("http://meromen.github.io/go-tasks/updater/getInfo.json")
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var getInfo Info

	err = json.Unmarshal(body, &getInfo)
	if err != nil {
		return false, err
	}

	// Parsing string version to int
	intCurVer, _ := strconv.ParseInt(strings.Replace(cerVer.Version, ".", "", -1), 10, 64)
	intGetVer, _ := strconv.ParseInt(strings.Replace(getInfo.Version, ".", "", -1), 10, 64)

	anwser := intGetVer > intCurVer

	return anwser, nil
}
