package core

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/viethapascal/go-moralis-sdk/utils"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const BaseUrl = "https://deep-index.moralis.io/api/v2"
const EndpointWeightPath = "info/endpointWeights"

type RequestQuery struct {
	Page           int      `url:"page"`
	Limit          int      `url:"limit"`
	DisableTotal   bool     `url:"disable_total"`
	TokenAddresses []string `url:"token_addresses"`
	Format         string   `url:"format"`
	NormalizedData bool     `url:"normalizeMetadata"`
}
type Params map[string]interface{}

func DefaultQuery() RequestQuery {
	return RequestQuery{
		Page:           1,
		Limit:          100,
		DisableTotal:   true,
		TokenAddresses: nil,
		Format:         "decimal",
		NormalizedData: false,
	}
}

type Moralis struct {
	APIUrl  string
	apiKey  string
	ChainID string

	APIHeader http.Header
	http      *http.Client
	ctx       context.Context

	NFT       *NFTAPI
	Endpoints []EndpointData
	Uri       *UrlBuilder
}

func MoralisAPI() (*Moralis, error) {
	header := http.Header{}
	apiKey := os.Getenv("MORALIS_API_KEY")
	if len(apiKey) == 0 {
		return nil, ErrMissingAPIKey
	}
	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json")
	header.Add("X-API-Key", apiKey)
	m := &Moralis{
		APIUrl:    BaseUrl,
		apiKey:    os.Getenv("MORALIS_API_KEY"),
		ChainID:   "",
		APIHeader: header,
		http:      http.DefaultClient,
		ctx:       context.Background(),
	}
	//enableCostControl, err := strconv.ParseBool(os.Getenv("MORALIS_ENABLE_COST_CONTROL"))
	//if err == nil && enableCostControl {
	//
	//}
	//e, err := m.GetEndpointWeight()
	//if err != nil {
	//	return nil, err
	//}
	//m.Endpoints = e
	m.NFT = newNFTApi(m)
	uri, err := newUrlBuilder(m)
	if err != nil {
		return nil, err
	}
	m.Uri = uri

	return m, nil
}

func (m *Moralis) WithChainID(chainID string) *Moralis {
	m.ChainID = chainID
	return m
}

type RequestOption func(r *http.Request)

func UseDefaultQuery() RequestOption {
	return func(r *http.Request) {
		opt := DefaultQuery()
		q := r.URL.Query()
		q2, _ := query.Values(opt)
		for k := range q2 {
			q.Set(k, q2.Get(k))
		}
		r.URL.RawQuery = q.Encode()
		log.Println(r.URL.Query().Encode())
	}
}

func Page(p int) RequestOption {
	return func(r *http.Request) {
		q := r.URL.Query()
		q.Set("page", strconv.FormatInt(int64(p), 10))
		r.URL.RawQuery = q.Encode()
	}
}

func Limit(p int) RequestOption {
	return func(r *http.Request) {
		q := r.URL.Query()
		q.Set("limit", strconv.FormatInt(int64(p), 10))
		r.URL.RawQuery = q.Encode()
	}
}

func Normalize() RequestOption {
	return func(r *http.Request) {
		q := r.URL.Query()
		q.Set("normalizeMetadata", "true")
		r.URL.RawQuery = q.Encode()
	}
}

func DisableTotal(b bool) RequestOption {
	return func(r *http.Request) {
		q := r.URL.Query()
		q.Set("disable_total", strconv.FormatBool(b))
		r.URL.RawQuery = q.Encode()
	}
}

func TokenAddresses(addr ...string) RequestOption {
	return func(r *http.Request) {
		q := r.URL.Query()
		for i := range addr {
			q.Set(fmt.Sprintf("token_addresses[%d]", i), addr[i])
		}
		r.URL.RawQuery = q.Encode()
	}
}

func Query(data map[string]interface{}) RequestOption {
	return func(r *http.Request) {
		q := r.URL.Query()
		for k := range data {
			q.Set(k, utils.EncodeToQueryString(data[k]))
		}
		r.URL.RawQuery = q.Encode()
	}
}

func (m *Moralis) Do(req *http.Request) (*http.Response, error) {
	log.Println("[REQUEST] Url:", req.URL.String())
	ctx := req.Context()
	req.Header = m.APIHeader
	response, err := m.http.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return nil, err
		}
	}

	return response, nil
}
func (m *Moralis) CheckResponse(response *http.Response, result interface{}) error {

	defer response.Body.Close()

	// If the response contains a client or a server error then return the error.
	if response.StatusCode >= http.StatusBadRequest {
		return errors.New("moralis error. status code:" + response.Status)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %w", err)
	}

	if len(responseBody) > 0 && string(responseBody) != "{}" {
		if err = json.Unmarshal(responseBody, &result); err != nil {
			return fmt.Errorf("failed to unmarshal response payload: %w", err)
		}
	}

	return nil
}

func (m *Moralis) Get(url string, payload interface{}, options ...RequestOption) error {
	var body bytes.Buffer
	if payload != nil {
		if err := json.NewEncoder(&body).Encode(payload); err != nil {
			return fmt.Errorf("encoding request payload failed: %w", err)
		}
	}
	request, _ := http.NewRequest("GET", url, &body)
	for _, opts := range options {
		opts(request)
	}
	response, err := m.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send the request: %w", err)
	}
	return m.CheckResponse(response, payload)
}

func (m Moralis) Post(url string, payload interface{}, result interface{}, options ...RequestOption) error {
	var body bytes.Buffer
	if payload != nil {
		if err := json.NewEncoder(&body).Encode(payload); err != nil {
			return fmt.Errorf("encoding request payload failed: %w", err)
		}
	}
	request, _ := http.NewRequest("POST", url, &body)
	for _, opts := range options {
		opts(request)
	}
	response, err := m.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send the request: %w", err)
	}
	return m.CheckResponse(response, result)
}

func (m *Moralis) GetEndpointWeight() ([]EndpointData, error) {
	urlPath := fmt.Sprintf("%s/%s", m.APIUrl, EndpointWeightPath)
	result := make([]EndpointData, 0)
	err := m.Get(urlPath, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
