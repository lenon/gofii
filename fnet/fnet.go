package fnet

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
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
	// the following fields were not included because they are either empty or
	// do not hold any useful information:
	// arquivoEstruturado, assuntos, cnpjAdministrador, cnpjFundo, dda formatoEstruturaDocumento,
	// idEntidadeGerenciadora, idSelectItemConvenio, idSelectNotificacaoConvenio, idTemplate,
	// indicadorFundoAtivoB3, nomeAdministrador, numeroEmissao, ofertaPublica, tipoPedido
	ID                          int    `json:"id"`
	AdditionalInformation       string `json:"informacoesAdicionais"`
	Category                    string `json:"categoriaDocumento"`
	FundDescription             string `json:"descricaoFundo"`
	FundMarketName              string `json:"nomePregao"`
	HighPriority                bool   `json:"altaPrioridade"`
	ReferenceDate               string `json:"dataReferencia"`
	ReferenceDateFormat         string `json:"formatoDataReferencia"`
	Reviewed                    string `json:"analisado"`
	Status                      string `json:"situacaoDocumento"`
	SubCategory1                string `json:"tipoDocumento"`
	SubCategory2                string `json:"especieDocumento"`
	SubmissionDate              string `json:"dataEntrega"`
	SubmissionMethod            string `json:"modalidade"`
	SubmissionMethodDescription string `json:"descricaoModalidade"`
	SubmissionStatus            string `json:"status"`
	SubmissionStatusDescription string `json:"descricaoStatus"`
	Version                     int    `json:"versao"`
}

type Page struct {
	TotalRecords int        `json:"recordsTotal"`
	Data         []Document `json:"data"`
}

func NewClient() *Client {
	return &Client{
		BaseURL:      "https://fnet.bmfbovespa.com.br/fnet/publico/pesquisarGerenciadorDocumentosDados",
		UserAgent:    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:104.0) Gecko/20100101 Firefox/104.0",
		PageSize:     100,
		TotalRecords: 0, // total number of records is only set after the first page request
		HTTPClient: &http.Client{
			Timeout: time.Duration(30) * time.Second, // timeout after 30s
		},
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

	c.TotalRecords = decoded.TotalRecords

	return &decoded, nil
}

func (c *Client) GetFirstPage() (*Page, error) {
	return c.GetPage(1)
}

func (c *Client) TotalPages() int {
	return NumberOfPages(c.TotalRecords, c.PageSize)
}

func (d *Document) ParsedReferenceDate() (time.Time, error) {
	return ParseDate(d.ReferenceDate, d.ReferenceDateFormat)
}

func (d *Document) ParsedSubmissionDate() (time.Time, error) {
	return ParseDate(d.SubmissionDate, DATE_FORMAT_DMY_HM)
}
