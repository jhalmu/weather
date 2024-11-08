package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	XMLName xml.Name `xml:"target:region"`
}

func main() {
	res, err := http.Get("https://opendata.fmi.fi/wfs/fin?service=WFS&version=2.0.0&request=getFeature&storedquery_id=fmi::observations::weather::timevaluepair&place=Helsinki&")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode !=200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
		if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}