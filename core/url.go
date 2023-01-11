package core

import (
	"encoding/json"
	"fmt"
	slices2 "github.com/go-moralis-sdk/utils/slices"
	"regexp"
	"strings"
)

type UrlBuilder struct {
	*Moralis
	BaseUrl   string
	Path      string
	Endpoints EndpointList
}

var re, _ = regexp.Compile(`\{([a-z_]+)}`)

func newUrlBuilder(m *Moralis) (*UrlBuilder, error) {
	u := &UrlBuilder{
		Moralis: m,
		BaseUrl: BaseUrl,
		Path:    "",
	}
	endpoints, err := u.GetEndpointWeight()
	if err != nil {
		return nil, err
	}
	u.Endpoints = endpoints
	return u, nil
}

type EndpointData struct {
	Endpoint      string `json:"endpoint"`
	Path          string `json:"path"`
	Price         int    `json:"price"`
	RateLimitCost int    `json:"rateLimitCost"`
}
type EndpointList []EndpointData

func IsEndpoint(e *EndpointData, name string) bool {
	return e.Endpoint == name
}

func (e EndpointData) String() string {
	bytes, _ := json.MarshalIndent(e, "", "\t")
	return string(bytes)
}

func (ub *UrlBuilder) GetEndPoint(name string) *EndpointData {
	return slices2.FirstObject(ub.Endpoints, func(e EndpointData) bool { return e.Endpoint == name })
}
func (ub *UrlBuilder) Encode(name string, params map[string]string) string {
	path := ub.GetEndPoint(name).Encode(params)
	return fmt.Sprintf("%s/%s?chain=%s", ub.BaseUrl, path, ub.ChainID)
}
func (e *EndpointData) Encode(params map[string]string) string {
	//e := GetEndPoint(eps, ep)
	groups := re.FindAllStringSubmatch(e.Path, -1)
	p := e.Path
	for i := range groups {
		p = strings.Replace(p, groups[i][0], params[groups[i][1]], -1)
	}
	return p
}
