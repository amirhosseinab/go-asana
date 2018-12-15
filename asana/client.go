package asana

type (
	Client struct {
	}

	Authorization struct {
	}
)

func NewClient(auth Authorization) (*Client, error) {
	return &Client{}, nil
}
