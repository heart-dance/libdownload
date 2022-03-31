package libdownload

type HttpClient interface{}

type httpClient struct {
	UserAgent string
}

func NewHttpClient() HttpClient {
	return &httpClient{
		UserAgent: "Seed",
	}
}
