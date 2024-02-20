package httpclient

import (
	"os"
)

type HttpClientInfra struct {
	proxy string
}

func NewHttpClientInfra(proxy string) *HttpClientInfra {
	return &HttpClientInfra{proxy: proxy}
}

func (h *HttpClientInfra) DownloadUrl(webUrl []string) (map[string]*os.File, error) {
	// TODO implement me
	panic("implement me")
}
