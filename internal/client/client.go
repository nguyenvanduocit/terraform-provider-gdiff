package client

import (
	"github.com/go-git/go-git/v5"
	"github.com/pkg/errors"
)

type ApiClient struct {
	GitPath string
	repo    *git.Repository
}

func (c *ApiClient) GetRepo() (*git.Repository, error) {
	if c.repo == nil {
		repo, err := git.PlainOpen(c.GitPath)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to open git repo: "+c.GitPath)
		}
		c.repo = repo
	}

	return c.repo, nil
}
