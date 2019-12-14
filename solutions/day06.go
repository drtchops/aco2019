package solutions

import (
	"fmt"
	"strings"
)

type planet06 struct {
	name     string
	parent   string
	children []string
}

type step06 struct {
	name  string
	steps int
}

// Day06 ...
func Day06(input string) {
	planets := make(map[string]planet06)

	for _, orbit := range strings.Split(input, "\n") {
		parts := strings.Split(orbit, ")")
		parent := parts[0]
		child := parts[1]

		if p, ok := planets[parent]; ok {
			p.children = append(p.children, child)
			planets[parent] = p
		} else {
			p = planet06{
				name:     parent,
				children: []string{child},
			}
			planets[parent] = p
		}

		if p, ok := planets[child]; ok {
			p.parent = parent
			planets[child] = p
		} else {
			p = planet06{
				name:   child,
				parent: parent,
			}
			planets[child] = p
		}
	}

	visited := make(map[string]bool)
	next := make([]step06, 0)

	me := planets["YOU"]
	start := planets[me.parent]
	if start.parent != "" {
		next = append(next, step06{name: start.parent})
	}
	for _, c := range start.children {
		next = append(next, step06{name: c})
	}

	for len(next) > 0 {
		s := next[0]

		if s.name == "SAN" {
			fmt.Println(s.steps)
			break
		}
		visited[s.name] = true

		p := planets[s.name]
		if p.parent != "" && !visited[p.parent] {
			next = append(next, step06{steps: s.steps + 1, name: p.parent})
		}
		for _, c := range p.children {
			if !visited[c] {
				next = append(next, step06{steps: s.steps + 1, name: c})
			}
		}
		next = next[1:]
	}
}
