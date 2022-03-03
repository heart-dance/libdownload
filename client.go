package down

type Client interface{}

type client struct {
	HttpClient HttpClient
}

func NewClient() Client {
	return &client{
		HttpClient: NewHttpClient(),
	}
}
