package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	order := topoSort(prereqs)
	for k, v := range order {
		fmt.Printf("%d\t%s\n", k, v)
	}
}

func topoSort(m map[string][]string) []string {
	isVisit := make(map[string]bool)
	var order []string
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !isVisit[item] {
				isVisit[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
