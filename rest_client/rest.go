package rest_client

import (
	"fmt"
	"net/http"
	"strings"
	"io"
	"io/ioutil"
)

type RestClient struct {
	host string
	c    *http.Client
}

func NewRestClient(port string) *RestClient {
	return &RestClient{
		host: "http://localhost:" + port,
		c:    &http.Client{},
	}
}

func (rest *RestClient) GetInt() error {
	return rest.doGet(strings.Join([]string{rest.host, "int"}, "/"))
}

func (rest *RestClient) GetIntSlice() error {
	return rest.doGet(strings.Join([]string{rest.host, "int-slice"}, "/"))
}

func (rest *RestClient) GetString() error {
	return rest.doGet(strings.Join([]string{rest.host, "string"}, "/"))
}

func (rest *RestClient) GetStringSlice() error {
	return rest.doGet(strings.Join([]string{rest.host, "string-slice"}, "/"))
}

func (rest *RestClient) GetBlob() error {
	return rest.doGet(strings.Join([]string{rest.host, "blob"}, "/"))
}

func (rest *RestClient) GetStruct() error {
	return rest.doGet(strings.Join([]string{rest.host, "struct"}, "/"))
}

func (rest *RestClient) GetStructSlices() error {
	return rest.doGet(strings.Join([]string{rest.host, "struct-slices"}, "/"))
}

func (rest *RestClient) GetStructStructs() error {
	return rest.doGet(strings.Join([]string{rest.host, "struct-structs"}, "/"))
}

func (rest *RestClient) GetFile() error {
	return rest.doGet(strings.Join([]string{rest.host, "file"}, "/"))
}

func (rest *RestClient) doGet(endpoint string) error {
	resp, err := rest.c.Get(endpoint)
	defer resp.Body.Close()

	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got status code %d", resp.StatusCode)
	}
	_, err = io.Copy(ioutil.Discard, resp.Body)
	return err
}
