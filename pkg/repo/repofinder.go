package repo

import (
	"fmt"
	"log"
	"os"
)

func FindRepos(root string) []Repository {
	var repositories []Repository
	var paths []string
	findRepoPaths(&paths, root)
	for _, path := range paths {
		repository := Repository{
			Path:  path,
			Root:  root,
			Group: "group",
			Name:  "name",
		}
		repositories = append(repositories, repository)
	}
	return repositories
}

func findRepoPaths(allRepoPaths *[]string, pathToCheck string) bool {
	fileInfo, err := os.Stat(pathToCheck)
	check(err)
	if fileInfo.Mode().IsRegular() {
		return false
	}

	directory, err := os.Open(pathToCheck)
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
			*allRepoPaths = append(*allRepoPaths, fullPath(pathToCheck, currentDir))
			stopDigging = true
			break
		}
	}

	if !stopDigging {
		for i := 0; i < len(directories); i++ {
			currentDir := directories[i]
			stopDigging := findRepoPaths(allRepoPaths, fullPath(pathToCheck, currentDir))
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
