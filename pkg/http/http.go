package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ClientHttpParams[payloadStruct any, responseStruct any] struct {
	BaseUrl   string
	Path      string
	Token     string
	Header    map[string]string
	Payload   payloadStruct
	Response  responseStruct
	URLParams map[string]interface{} // params must be ordered from left to right
}

func (receiver *ClientHttpParams[payloadStruct, responseStruct]) Get() (response responseStruct, err error) {
	url := fmt.Sprintf("%s%s", receiver.BaseUrl, receiver.Path)

	req, errConstruct := receiver.constructRequest(url, "GET")

	if errConstruct != nil {
		log.Println(errConstruct)
		err = errConstruct
		return
	}

	log.Printf("request to : %s%s", req.URL.Host, req.URL.Path)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)

	errResp := receiver.constructResponse(res, &response)
	if errResp != nil {
		return
	}
	return
}

func (receiver *ClientHttpParams[payloadStruct, responseStruct]) Post() (response responseStruct, err error) {
	url := fmt.Sprintf("%s%s", receiver.BaseUrl, receiver.Path)

	req, errConstruct := receiver.constructRequest(url, "POST")

	if errConstruct != nil {
		log.Println(errConstruct)
		err = errConstruct
		return
	}

	log.Printf("request to : %s%s", req.URL.Host, req.URL.Path)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)

	errResp := receiver.constructResponse(res, &response)
	if errResp != nil {
		return
	}
	return
}

func (receiver *ClientHttpParams[payloadStruct, responseStruct]) constructResponse(res *http.Response, responseClient *responseStruct) (err error) {
	respBytes, errRead := io.ReadAll(res.Body)
	if errRead != nil {
		log.Println(errRead)
		err = errRead
		return
	}

	err = json.Unmarshal(respBytes, &responseClient)
	if err != nil {
		return
	}
	return
}

func (receiver *ClientHttpParams[payloadStruct, responseStruct]) constructRequest(url string, method string) (req *http.Request, err error) {
	var strPayload []byte

	if len(receiver.URLParams) > 0 {
		for _, params := range receiver.URLParams {
			url += fmt.Sprintf("/%v", params)
		}
	}

	if method != "GET" {
		if strPayload, err = json.Marshal(receiver.Payload); err != nil {
			return
		}
	}

	if req, err = http.NewRequest(method, url, bytes.NewBuffer(strPayload)); err != nil {
		return
	}

	if len(receiver.Header) > 0 {
		for k, v := range receiver.Header {
			req.Header.Add(k, v)
		}
	}

	req.Header.Add("Authorization", receiver.Token)
	req.Header.Add("Accept", "application/json")
	if method != http.MethodGet {
		req.Header.Add("Content-Type", "application/json")
	}

	return
}
