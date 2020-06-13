package repo

import "fmt"

type Repository struct {
	Path  string
	Root  string
	Group string
	Name  string
}

func (p Repository) String() string {
	return fmt.Sprintf("Path=[%s], Root=[%s], Group=[%s], Name=[%s]",
		p.Path, p.Root, p.Group, p.Name)
}
