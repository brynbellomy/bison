package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"reflect"

	"golang.org/x/net/http2"
)

func EnvVarOrPanic(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic("Missing environment variable " + key)
	}
	return val
}

func JSONReq(method string, url string, body any, headers http.Header, response any) (http.Header, error) {
	if headers == nil {
		headers = http.Header{}
	}

	headers["Accept"] = []string{"application/json"}
	headers["Content-Type"] = []string{"application/json"}

	var bs []byte
	var err error
	if body != nil && !reflect.ValueOf(body).IsZero() {
		bs, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	bs, respHeader, err := HTTPReq(method, url, ioutil.NopCloser(bytes.NewReader(bs)), headers)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bs, &response)
	if err != nil {
		return nil, err
	}

	return respHeader, nil
}

func HTTPReq(method string, url string, body io.ReadCloser, headers http.Header) ([]byte, http.Header, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, nil, err
	}

	req.Header = http.Header(headers)

	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Request:\n\n%s\n\n", string(reqDump))

	c := &http.Client{
		Transport: &http2.Transport{},
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Response (status", resp.StatusCode, resp.Status, "):\n\n", string(respDump), "\n")
	fmt.Println(string(bs))
	fmt.Println("--------------------------------------------------------------------------------------------------------")

	return bs, resp.Header, nil
}
