package types

import (
	"github.com/google/go-github/github"
)

type Repo struct {
	githubRepo *github.Repository
}

func (r *Repo) Stars() float64 {
	if r.githubRepo.StargazersCount != nil {
		return float64(*r.githubRepo.StargazersCount)
	}
	return 0
}

func CreateRepo(githubRepo *github.Repository) *Repo {
	repo := &Repo{
		githubRepo: githubRepo,
	}
	return repo
}
