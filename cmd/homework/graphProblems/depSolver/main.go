package main

import (
	"errors"
	"fmt"
)

func main() {
	deps := map[string][]string{
		"0": {},
		"1": {"0"},
		"2": {"0"},
		"3": {"1", "2"},
		"4": {"3"},
		"5": {"1", "4"},
		"6": {"3", "0"},
		"7": {},
		"8": {"7", "6", "5", "3"},
	}

	output, err := depSolver(deps)
	if err != nil {
		panic(err)
	}
	for _, s := range output {
		fmt.Println(s + ", ")
	}
}

type node struct {
	value string
	edges []node
}

func (n *node) addEdge(e node) {
	n.edges = append(n.edges, e)
}

var resolved map[string]bool
var visited map[string]bool
var output []string

func depSolver(deps map[string][]string) ([]string, error) {
	output = []string{}
	nodes := make(map[string]*node)
	resolved = make(map[string]bool)
	visited = make(map[string]bool)
	for k, v := range deps {
		nodes[k] = &node{k, []node{}}
		for _, e := range v {
			nodes[k].addEdge(node{e, []node{}})
		}
	}

	for _, v := range nodes {
		// if it's not resolved, resolve it
		if _, ok := resolved[v.value]; !ok {
			err := resolve(v)
			if err != nil {
				return nil, err
			}
		}
	}
	return output, nil
}

func resolve(n *node) error {
	// if already visited, we found a circular ref
	if _, ok := visited[n.value]; ok {
		return errors.New(fmt.Sprintf("circular ref: %s", n.value))
	}

	// if already resolved, stop
	if _, ok := resolved[n.value]; ok {
		return nil
	}

	// visit
	visited[n.value] = true

	// walk edges
	for _, v := range n.edges {
		err := resolve(&v)
		if err != nil {
			return err
		}
	}
	// resolve
	resolved[n.value] = true

	// remove from visited since we are done with this particular depth search
	delete(visited, n.value)

	// add to output
	output = append(output, n.value)

	return nil
}
