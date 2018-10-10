package mailchimpExtract

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"log"
	"encoding/json"
)

type BatchReponse struct {
	Id                 string      `json:"id"`
	Status             string      `json:"status"`
	TotalOperations    int         `json:"total_operations"`
	FinishedOperations int         `json:"finished_operations"`
	ErroredOperations  int         `json:"errored_operations"`
	SubmittedAt        string      `json:"submitted_at"`
	CompletedAt        string      `json:"completed_at"`
	ResponseBodyUrl    string      `json:"response_body_url"`
	Links              interface{} `json:"_links"`
}

type BatchRequest struct {
	Operations []BatchOperation `json:"operations"`
}

type BatchOperation struct {
	Method      string      `json:"method"`
	Path        string      `json:"path"`
	Params      interface{} `json:"params,omitempty" bson:"omitempty"`
	Body        string      `json:"body,omitempty" bson:"omitempty"`
	OperationId string      `json:"operation_id,omitempty" bson:"omitempty"`
}

func MakeBatchReq(apiKey, apiRoot string, requestBody []byte) (BatchReponse) {
	client := &http.Client{}
	apiUrl := apiRoot + "batches/"
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBody))
	req.SetBasicAuth("apiKey", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	var batchResponse BatchReponse
	json.Unmarshal(body, &batchResponse)
	return batchResponse
}

/*func BatchRequestHandler(batchReqResp BatchReponse) {
	if batchReqResp.Status == "finished" {
		batchReqResp.ResponseBodyUrl
	}
}*/
