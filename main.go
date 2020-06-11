package main

import (
	"fmt"
	"github.com/hoto/git-my-git/config"
	"github.com/hoto/git-my-git/io"
)

func main() {
	config.ParseArgsAndFlags()

	var gitRepos []string
	for _, rootPath := range config.ProjectRoots {
		io.FindGitRepositories(&gitRepos, rootPath)
	}
	fmt.Printf("Found %d repos.\n", len(gitRepos))
	fmt.Printf("Paths=%s", gitRepos)
}
