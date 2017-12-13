# Cloudconvert Go Client for Cloudconvert

**A simple client written in Go for the Cloudconvert app API **

## Use example:
```
//main.go
package main

import (
	"fmt"
	"github.com/awmpietro/cloudconvgocl"
)

func main() {
  //1 - Criar um mapa com as configuraçes da conversão, conforme abaixo:
	data := make(map[string]string)
	data["key"] = "SUA_API_KEY_CLOUDCONVERT"
	data["from"] = "jpg" //formato atual do arquivo
	data["to"] = "png" //formato para qual o arquivo sera convertido
	data["filePath"] = "https://a-z-animals.com/media/animals/images/original/gopher_4.jpg" //URL do arquivo
	data["fileName"] = "gopher.jpg" //Nome do arquivo que será salvo COM A EXTENSÃO ORIGINAL
	data["input"] = "" //Opcional: O padrão é "download"
	data["mode"] = "" //Opcional: O padrão é "convert"
	data["pathToSave"] = "./assets/" //Local onde o arquivo convertido será salvo
  
  //2 - passa o mapa para a função cloudconvgocl.Convert() que retorna erro ou nil se tudo ocorreu de acordo
	err := cloudconvgocl.Convert(data)
	if err != nil {
		fmt.Println("Erro: ", err)
	}
}
```
### TODO:
More options conversion config
