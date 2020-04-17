package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjdghks994/tucker_golang_web/deco"
	"github.com/tjdghks994/tucker_golang_web/myapp"
)

func TestIndexPage(t *testing.T) {
	assert := assert.New(t)

	mux := myapp.NewHandler()
	decoHandler := deco.NewDecoHandler(mux, logger)
	decoHandler = deco.NewDecoHandler(decoHandler, logger2)

	ts := httptest.NewServer(decoHandler)
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, err := ioutil.ReadAll(resp.Body)
	assert.NoError(err)
	assert.Equal(string(data), "Hello World")
}

func TestDecorator(t *testing.T) {
	assert := assert.New(t)

	mux := myapp.NewHandler()
	decoHandler := deco.NewDecoHandler(mux, logger)
	decoHandler = deco.NewDecoHandler(decoHandler, logger2)

	ts := httptest.NewServer(decoHandler)
	defer ts.Close()

	buf := &bytes.Buffer{}
	log.SetOutput(buf)

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	r := bufio.NewReader(buf)
	// 첫번째 줄 읽기
	line, _, err := r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "Started")
	// 두번째 줄 읽기
	line, _, err = r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "Started")
	// 세번째 줄 읽기
	line, _, err = r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "Completed")
}
