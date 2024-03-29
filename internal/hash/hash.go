package hash

import (
	"hash/fnv"
)

type Hasher interface {
	Get(key string) (uint32, error)
}

type Client struct{}

func New() *Client {
	return &Client{}
}

func (c *Client) Get(key string) (uint32, error) {
	hasher := fnv.New32()
	_, err := hasher.Write([]byte(key))
	if err != nil {
		return 0, err
	}
	return hasher.Sum32(), nil
}
