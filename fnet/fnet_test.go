package fnet_test

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/lenon/gofii/fnet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type httpMock struct {
	mock.Mock
}

func (m *httpMock) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*http.Response), args.Error(1)
}

func sample(name string) *os.File {
	reader, err := os.Open(filepath.Join("testdata", name))
	if err != nil {
		panic(err)
	}
	return reader
}

func TestGetFirstPage(t *testing.T) {
	httpMock := &httpMock{}
	httpMockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       sample("sample.json"),
	}

	httpMock.On("Do", mock.Anything).Return(httpMockResponse, nil)

	client := fnet.NewClient()
	client.HTTPClient = httpMock
	response, err := client.GetFirstPage()

	httpMock.AssertCalled(t, "Do", mock.AnythingOfType("*http.Request"))
	request := httpMock.Calls[0].Arguments.Get(0).(*http.Request)
	queryString := request.URL.Query()

	assert.Equal(t, "GET", request.Method)
	assert.Contains(t, request.URL.String(), "pesquisarGerenciadorDocumentosDados")
	assert.Equal(t, "application/json", request.Header.Get("Accept"))

	assert.Equal(t, "1", queryString.Get("d"))
	assert.Equal(t, "0", queryString.Get("s"))
	assert.Equal(t, "100", queryString.Get("l"))
	assert.Equal(t, "desc", queryString.Get("o[0][dataEntrega]"))
	assert.Equal(t, "1", queryString.Get("tipoFundo"))
	assert.Equal(t, "0", queryString.Get("idCategoriaDocumento"))
	assert.Equal(t, "0", queryString.Get("idTipoDocumento"))
	assert.Equal(t, "0", queryString.Get("idEspecieDocumento"))

	assert.Nil(t, err)
	assert.Equal(t, 3, len(response.Data))
	assert.Equal(t, 789, response.Data[0].Id)
	assert.Equal(t, "14/09/2022 19:35", response.Data[0].SubmissionDate)
	assert.Equal(t, 1333, client.TotalPages())
}

func TestGetFirstPageWithError(t *testing.T) {
	httpMock := &httpMock{}
	httpMock.On("Do", mock.Anything).Return(nil, errors.New("something went terribly wrong"))

	client := fnet.NewClient()
	client.HTTPClient = httpMock
	response, err := client.GetFirstPage()

	assert.Nil(t, response)
	assert.EqualError(t, err, "something went terribly wrong")
}

func TestGetFirstPageMalformedJson(t *testing.T) {
	httpMock := &httpMock{}
	httpMockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       sample("malformed.json"),
	}

	httpMock.On("Do", mock.Anything).Return(httpMockResponse, nil)

	client := fnet.NewClient()
	client.HTTPClient = httpMock
	response, err := client.GetFirstPage()

	assert.Nil(t, response)
	assert.EqualError(t, err, "invalid character '}' after object key")
}

func TestGetPage(t *testing.T) {
	httpMock := &httpMock{}
	httpMockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       sample("sample.json"),
	}

	httpMock.On("Do", mock.Anything).Return(httpMockResponse, nil)

	client := fnet.NewClient()
	client.HTTPClient = httpMock
	response, err := client.GetPage(2)

	httpMock.AssertCalled(t, "Do", mock.AnythingOfType("*http.Request"))
	request := httpMock.Calls[0].Arguments.Get(0).(*http.Request)
	queryString := request.URL.Query()

	assert.Equal(t, "100", queryString.Get("s"))
	assert.Equal(t, "100", queryString.Get("l"))
	assert.Nil(t, err)
	assert.Equal(t, 3, len(response.Data))
}
