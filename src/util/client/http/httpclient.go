package http

import (
	"log"
	"net/http"
	"io/ioutil"
	"fmt"
)

func Request(url string) string{
	httpClient := &http.Client{}
    r, err := http.NewRequest("GET", url, nil)
    if err != nil {
    	log.Fatalln("http client", err)
    }

    response, errResp := httpClient.Do(r)
    if errResp !=nil {
    	log.Fatalln("http client request", errResp)
    }

    if response.StatusCode == 200 {
		str, _ := ioutil.ReadAll(response.Body)
	    bodystr := string(str)
	    log.Println(fmt.Sprintf("body string %s", bodystr))
		return bodystr
	 }
	 return "request err..."
}