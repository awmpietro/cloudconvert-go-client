package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

//Boot
func Boot(data map[string]string) {
	processUrl := "https://api.cloudconvert.com/process"
	//Example:
	data = make(map[string]string)
	data["key"] = ""
	data["from"] = ""
	data["to"] = ""
	data["filePath"] = ""
	data["filePath"] = ""
	data["fileName"] = ""
	data["outputFmt"] = ""
	data["input"] = ""
	data["mode"] = ""
	data["pathToSave"] = ""
	if data["input"] == "" {
		data["input"] = "download"
	}
	if data["mode"] == "" {
		data["mode"] = "convert"
	}
	if data["pathToSave"] == "" {
		data["pathToSave"] = "/"
	}

	process := Process{
		Apikey:       data["key"],
		Mode:         data["mode"],
		InputFormat:  data["from"],
		OutputFormat: data["to"],
	}
	processJSON, err := json.Marshal(process)
	if err != nil {
		fmt.Println("Erro: ", err)
	}

	resp, err := http.Post(processUrl, "application/json", bytes.NewReader(processJSON))

	if err != nil {
		fmt.Println("Erro: ", err)
	}
	defer resp.Body.Close()
	//Unmarshall json to Go Struct
	var processResponse ProcessResponse
	err = json.NewDecoder(resp.Body).Decode(&processResponse)
	if err != nil {
		fmt.Println("Erro: ", err)
	}

	processStart(processResponse, data)

}

//ProcessStart
func processStart(processResponse ProcessResponse, data map[string]string) {
	processStart := ProcessStart{
		Wait:         true,
		Input:        data["input"],
		File:         data["filePath"],
		Filename:     data["fileName"],
		Outputformat: data["outputFmt"],
	}
	processJSONStart, err := json.Marshal(processStart)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
	resp, err := http.Post("https:"+processResponse.Url, "application/json", bytes.NewReader(processJSONStart))

	if err != nil {
		fmt.Println("Erro: ", err)
	}

	//Unmarshall json to Go Struct
	var processStartResponse ProcessStartResponse
	err = json.NewDecoder(resp.Body).Decode(&processStartResponse)
	if err != nil {
		fmt.Println("Erro: ", err)
	}

	fileDownload(processStartResponse, data)
}

//fileDownload
func fileDownload(processStartResponse ProcessStartResponse, data map[string]string) error {
	out, err := os.Create(data["pathToSave"] + processStartResponse.Output.Filename)
	if err != nil {
		fmt.Println("Erro: ", err)
		return err
	}
	defer out.Close()
	resp, err := http.Get("https:" + processStartResponse.Output.Url)
	if err != nil {
		fmt.Println("Erro: ", err)
		return err
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Erro: ", err)
		return err
	}
	return nil
}
