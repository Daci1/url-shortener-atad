package response

type APIResponse[T any] struct {
	Data Data[T] `json:"data"`
}

type ApiRequest[T any] struct {
	Data Data[T] `json:"data"`
}

type Data[T any] struct {
	Type       string `json:"type"`
	Attributes T      `json:"attributes"`
}

type CreateUrlRequestAttributes struct {
	OriginalUrl string `json:"originalUrl"`
}

type UrlAttributes struct {
	Id          string `json:"-"`
	ShortUrl    string `json:"shortUrl"`
	OriginalUrl string `json:"originalUrl"`
	CreatedAt   string `json:"createdAt"`
	DeletedAt   string `json:"deletedAt,omitempty"`
}
