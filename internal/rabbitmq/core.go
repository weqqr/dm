package rabbitmq

import (
	"dm/internal/domain"
)

type Config struct {
	Address string `toml:"address"`
}

type Core struct {
	storage *ServerStorage
}

func (c *Core) FetchResponse(request *domain.StreamRequest) (*domain.StreamResponse, error) {
	response, err := c.storage.FetchResponse(StreamRequest{
		Id: request.GetId(),
	})

	if err != nil {
		return nil, err
	}

	return &domain.StreamResponse{
		Result: response.Result,
	}, nil
}
