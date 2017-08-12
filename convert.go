package cloudconvgocl

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func Convert(data map[string]string) error {
	processUrl := "https://api.cloudconvert.com/process"
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
		return err
	}

	resp, err := http.Post(processUrl, "application/json", bytes.NewReader(processJSON))

	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var processResponse ProcessResponse
	err = json.NewDecoder(resp.Body).Decode(&processResponse)
	if err != nil {
		return err
	}

	err = processStart(processResponse, data)
	return err

}

func processStart(processResponse ProcessResponse, data map[string]string) error {
	processStart := ProcessStart{
		Wait:         true,
		Input:        data["input"],
		File:         data["filePath"],
		Filename:     data["fileName"],
		Outputformat: data["outputFmt"],
	}
	processJSONStart, err := json.Marshal(processStart)
	if err != nil {
		return err
	}
	resp, err := http.Post("https:"+processResponse.Url, "application/json", bytes.NewReader(processJSONStart))

	if err != nil {
		return err
	}

	var processStartResponse ProcessStartResponse
	err = json.NewDecoder(resp.Body).Decode(&processStartResponse)
	if err != nil {
		return err
	}

	err = fileDownload(processStartResponse, data)
	return err
}

func fileDownload(processStartResponse ProcessStartResponse, data map[string]string) error {
	out, err := os.Create(data["pathToSave"] + processStartResponse.Output.Filename)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get("https:" + processStartResponse.Output.Url)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
