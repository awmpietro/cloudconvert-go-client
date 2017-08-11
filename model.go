package main

type Data struct {
	Key               string `toml:"key"`
	ProcessUrl        string `toml:"processurl"`
	FileToConvertPath string `toml:"filetoconvertpath"`
	FileToConvertName string `toml:"filetoconvertname"`
	ConvertFrom       string `toml:"convertfrom"`
	ConvertTo         string `toml:"convertto"`
}

type Process struct {
	Apikey       string `json:"apikey"`
	Mode         string `json:"mode"`
	InputFormat  string `json:"inputformat"`
	OutputFormat string `json:"outputformat"`
}

type ProcessResponse struct {
	Url        string `json:"url"`
	Id         string `json:"id"`
	Host       string `json:"host"`
	Expires    string `json:"expires"`
	Maxsize    int64  `json:"maxsize"`
	Concurrent int64  `json:"concurrent"`
	Minutes    int64  `json:"minutes"`
}

type ProcessStart struct {
	Wait         bool   `json:"wait"`
	Input        string `json:"input"`
	File         string `json:"file"`
	Filename     string `json:"filename"`
	Outputformat string `json:"outputformat"`
}

type ProcessStartResponseOutput struct {
	Url      string `json:"url"`
	Filename string `json:"filename"`
}

type ProcessStartResponse struct {
	Id      string                     `json:"id"`
	Message string                     `json:"message"`
	Step    string                     `json:"step"`
	Output  ProcessStartResponseOutput `json:"output"`
}
