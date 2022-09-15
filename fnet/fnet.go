package fnet

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	BaseURL      string
	UserAgent    string
	PageSize     int
	TotalRecords int
	HTTPClient   HTTPClient
}

type Document struct {
	ID                          int    `json:"id"`
	AdditionalInformation       string `json:"informacoesAdicionais"`
	DocumentCategory            string `json:"categoriaDocumento"`
	DocumentStatus              string `json:"situacaoDocumento"`
	DocumentSubCategory1        string `json:"tipoDocumento"`
	DocumentSubCategory2        string `json:"especieDocumento"`
	FundDescription             string `json:"descricaoFundo"`
	HighPriority                bool   `json:"altaPrioridade"`
	MarketName                  string `json:"nomePregao"`
	ReferenceDate               string `json:"dataReferencia"`
	ReferenceDateFormat         string `json:"formatoDataReferencia"`
	Reviewed                    string `json:"analisado"`
	Status                      string `json:"status"`
	StatusDescription           string `json:"descricaoStatus"`
	SubmissionDate              string `json:"dataEntrega"`
	SubmissionMethod            string `json:"modalidade"`
	SubmissionMethodDescription string `json:"descricaoModalidade"`
	Version                     int    `json:"versao"`
}

type Page struct {
	RecordsTotal int        `json:"recordsTotal"`
	Data         []Document `json:"data"`
}

func NewClient() *Client {
	return &Client{
		BaseURL:      "https://fnet.bmfbovespa.com.br/fnet/publico/pesquisarGerenciadorDocumentosDados",
		UserAgent:    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:104.0) Gecko/20100101 Firefox/104.0",
		PageSize:     100,
		TotalRecords: 0, // total number of records is only set after the first page request
		HTTPClient:   &http.Client{},
	}
}

func (c *Client) buildRequest(page int) *http.Request {
	request, err := http.NewRequest("GET", c.BaseURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	offset := PageOffset(page, c.PageSize)

	qs := url.Values{}
	qs.Add("d", "1") // seems to be important for the UI only
	qs.Add("s", strconv.Itoa(offset))
	qs.Add("l", strconv.Itoa(c.PageSize))
	qs.Add("o[0][dataEntrega]", "desc") // order by submission date desc
	qs.Add("tipoFundo", "1")            // 1 = FII
	qs.Add("idCategoriaDocumento", "0") // 0 = all
	qs.Add("idTipoDocumento", "0")      // 0 = all
	qs.Add("idEspecieDocumento", "0")   // 0 = all

	request.URL.RawQuery = qs.Encode()
	request.Header.Add("User-Agent", c.UserAgent)
	request.Header.Add("Accept", "application/json")

	return request
}

func (c *Client) GetPage(page int) (*Page, error) {
	request := c.buildRequest(page)
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	var decoded Page

	err = json.NewDecoder(response.Body).Decode(&decoded)
	if err != nil {
		return nil, err
	}

	c.TotalRecords = decoded.RecordsTotal

	return &decoded, nil
}

func (c *Client) GetFirstPage() (*Page, error) {
	return c.GetPage(1)
}

func (c *Client) TotalPages() int {
	return NumberOfPages(c.TotalRecords, c.PageSize)
}
