package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
)

func main() {
	var cnf Data
	if _, err := toml.DecodeFile("config.toml", &cnf); err != nil {
		fmt.Println(err)
		return
	}
	//Process 1
	process := Process{
		Apikey:       cnf.Key,
		Mode:         "convert",
		InputFormat:  cnf.ConvertFrom,
		OutputFormat: cnf.ConvertTo,
	}
	processJson, err := json.Marshal(process)

	resp, err := http.Post(cnf.ProcessUrl, "application/json", bytes.NewReader(processJson))
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	//Unmarshall json to Go Struct
	var processResponse ProcessResponse
	err = json.NewDecoder(resp.Body).Decode(&processResponse)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	//fmt.Println(processResponse.Url)
	//Process 2
	processStart := ProcessStart{
		Wait:         true,
		Input:        "download",
		File:         cnf.FileToConvertPath,
		Filename:     cnf.FileToConvertName,
		Outputformat: "png",
	}
	processJsonStart, err := json.Marshal(processStart)
	resp, err = http.Post("https:"+processResponse.Url, "application/json", bytes.NewReader(processJsonStart))

	if err != nil {
		fmt.Println("Erro: ", err)
	}

	//Unmarshall json to Go Struct
	var processStartResponse ProcessStartResponse
	err = json.NewDecoder(resp.Body).Decode(&processStartResponse)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	out, err := os.Create("./" + processStartResponse.Output.Filename)
	defer out.Close()
	resp, err = http.Get("https:" + processStartResponse.Output.Url)
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
}
