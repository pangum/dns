package dns

import (
	"github.com/goexl/dns"
)

// Client 客户端
type Client struct {
	*dns.Client
}

func newClient() (client *Client) {
	client = new(Client)
	client.Client = dns.New()

	return
}
