package main

import (
	"fmt"
	"github.com/hoto/git-my-git/pkg/config"
	"github.com/hoto/git-my-git/pkg/repo"
)

func main() {
	config.ParseArgsAndFlags()

	var allRepos []repo.Repository
	for _, rootPath := range config.ProjectRoots {
		repos := repo.FindRepos(rootPath)
		allRepos = addAll(allRepos, repos)
	}
	fmt.Println(allRepos)
	config.SaveRepos(allRepos)
}

func addAll(container []repo.Repository, toAdd []repo.Repository) []repo.Repository {
	for _, repository := range toAdd {
		container = append(container, repository)
	}
	return container
}
