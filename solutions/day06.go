package solutions

import (
	"fmt"
	"strings"
)

type planet struct {
	name     string
	parent   string
	children []string
}

type step struct {
	name  string
	steps int
}

// Day06 ...
func Day06(input string) {
	planets := make(map[string]planet)

	for _, orbit := range strings.Split(input, "\n") {
		parts := strings.Split(orbit, ")")
		parent := parts[0]
		child := parts[1]

		if p, ok := planets[parent]; ok {
			p.children = append(p.children, child)
			planets[parent] = p
		} else {
			p = planet{
				name:     parent,
				children: []string{child},
			}
			planets[parent] = p
		}

		if p, ok := planets[child]; ok {
			p.parent = parent
			planets[child] = p
		} else {
			p = planet{
				name:   child,
				parent: parent,
			}
			planets[child] = p
		}
	}

	visited := make(map[string]bool)
	next := make([]step, 0)

	me := planets["YOU"]
	start := planets[me.parent]
	if start.parent != "" {
		next = append(next, step{name: start.parent})
	}
	for _, c := range start.children {
		next = append(next, step{name: c})
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
			next = append(next, step{steps: s.steps + 1, name: p.parent})
		}
		for _, c := range p.children {
			if !visited[c] {
				next = append(next, step{steps: s.steps + 1, name: c})
			}
		}
		next = next[1:]
	}
}
