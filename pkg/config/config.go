package config

import (
	"fmt"
	"github.com/hoto/git-my-git/pkg/repo"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"time"
)

const (
	osConfigDir    = ".config"
	appConfigDir   = "git-my-git"
	configFileName = "repositories.yaml"
)

type Repositories struct {
	Repositories []Repository
}
type Repository struct {
	Path             string
	Root             string
	Group            string
	Name             string
	Frequency        int
	CreatedDate      string `yaml:"created_date"`
	LastAccessedDate string `yaml:"last_accessed_date"`
}

func SaveRepos(repositories []repo.Repository) {
	config := Repositories{Repositories: []Repository{}}
	for i := range repositories {
		repository := repositories[i]
		project := Repository{
			Path:        repository.Path,
			Root:        repository.Root,
			Group:       repository.Group,
			Name:        repository.Name,
			CreatedDate: time.Now().Format(time.RFC850),
		}
		config.Repositories = append(config.Repositories, project)
	}

	yamlString, err := yaml.Marshal(&config)
	check(err)
	writeToConfigFile(err, yamlString)
}

func writeToConfigFile(err error, yamlString []byte) {
	homeDir, err := os.UserHomeDir()
	check(err)
	configDir := fmt.Sprintf("%s/%s/%s", homeDir, osConfigDir, appConfigDir)
	err = os.MkdirAll(configDir, os.ModePerm)
	check(err)
	configFile := fmt.Sprintf("%s/%s", configDir, configFileName)
	file, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0644)
	check(err)
	_, err = file.Write(yamlString)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
