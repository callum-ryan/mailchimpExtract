package mailchimpExtract

import (
	"strings"
	"net/http"
	"io/ioutil"
	"log"
)

func GetApiRootMailChimp(apiKey string) (apiUrl string) {
	s := strings.Split(apiKey, "-")
	apiUrl = "https://" + s[1] + ".api.mailchimp.com/3.0/"
	return
}

func MakeReq(apiKey string, apiRoot string, endPoint string) (result []byte) {
	client := &http.Client{}
	apiUrl := apiRoot + endPoint

	req, err := http.NewRequest("GET", apiUrl, nil)
	req.SetBasicAuth("apiKey", apiKey)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body
}
