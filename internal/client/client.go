package client

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/pkg/errors"
	"io"
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

func (c *ApiClient) GetLastCommit(path string) (*plumbing.Hash, error) {
	repo, err := c.GetRepo()
	if err != nil {
		return nil, err
	}

	lastTagHash, err := repo.ResolveRevision("tag~1")
	if err != nil {
		return nil, errors.WithMessage(err, "failed to resolve last tag")
	}

	if lastTagHash.IsZero() {
		return nil, errors.New("last tag is zero")
	}

	// get the last commit that changed the path
	commit, err := repo.CommitObject(*lastTagHash)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get commit object")
	}

	files, err := commit.Files()
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get commit files")
	}

	for {
		file, err := files.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, errors.WithMessage(err, "failed to get next file")
		}

		if file.Name == path {
			return &commit.Hash, nil
		}
	}

	return nil, nil
}
