package internal

type GitClient interface {
	Clone(repo string) error
}