package git

import (
	"fmt"
	"github.com/go-git/go-git/v5" // with go modules enabled (GO111MODULE=on or outside GOPATH)
)

type client struct {
}

func (c *client) Clone(url, dir string) error {
	// Clone the given repository to the given directory
	Info("git clone %s %s", url, dir)
	r, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
	})

	CheckIfError(err)

	// ... retrieving the commit being pointed by HEAD
	Info("git show-ref --head HEAD")
	ref, err := r.Head()
	CheckIfError(err)
	fmt.Println(ref.Hash())

	_, err = r.Worktree()
	CheckIfError(err)
	return err
}

func (c *client) RemoveClone(dir string) error {

	return nil
}
