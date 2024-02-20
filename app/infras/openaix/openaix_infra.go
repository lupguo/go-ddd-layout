package openaix

type OpenAIInfra struct {
	proxy string
	auth  string
}

func NewOpenAIInfra(proxy string, auth string) *OpenAIInfra {
	return &OpenAIInfra{proxy: proxy, auth: auth}
}

func (o OpenAIInfra) GenerateAIImage(promptKey string, keywords string) (imgUrls []string, err error) {
	// TODO implement me
	panic("implement me")
}
