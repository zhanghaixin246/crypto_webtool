package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://localhost:8083/", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp,err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, "200 OK", resp.Status)
	assert.Equal(t, "{\"code\":0,\"data\":{},\"msg\":\"It works!\"}", string(body))
}
