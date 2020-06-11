package io

import (
	"fmt"
	"log"
	"os"
)

func FindGitRepositories(gitRepos *[]string, path string) bool {
	fileInfo, err := os.Stat(path)
	check(err)
	if fileInfo.Mode().IsRegular() {
		return false
	}

	directory, err := os.Open(path)
	check(err)
	directories, err := directory.Readdirnames(0)
	check(err)
	if len(directories) == 0 {
		return true //stop digging
	}

	stopDigging := false
	for i := 0; i < len(directories); i++ {
		currentDir := directories[i]
		if currentDir == ".git" {
			*gitRepos = append(*gitRepos, fullPath(path, currentDir))
			stopDigging = true
			break
		}
	}

	if !stopDigging {
		for i := 0; i < len(directories); i++ {
			currentDir := directories[i]
			stopDigging := FindGitRepositories(gitRepos, fullPath(path, currentDir))
			if stopDigging {
				continue
			}
		}
	}

	return stopDigging
}

func fullPath(startingDir string, currentDir string) string {
	return fmt.Sprintf("%s/%s", startingDir, currentDir)
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
