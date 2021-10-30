package main

import (
	"reflect"
	"testing"
)

func TestTopoSortUniqueResult(t *testing.T) {
	type args struct {
		g map[string]StringSet
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty graph",
			args: args{map[string]StringSet{}},
			want: []string{},
		},
		{
			name: "single node",
			args: args{map[string]StringSet{
				"A": newStringSet(),
			}},
			want: []string{"A"},
		},
		{
			name: "two nodes",
			args: args{map[string]StringSet{
				"B": newStringSet("A"),
			}},
			want: []string{"A", "B"},
		},
		{
			name: "multiple nodes",
			args: args{map[string]StringSet{
				"B": newStringSet("A"),
				"E": newStringSet("A", "C", "D"),
				"C": newStringSet("A", "B"),
				"D": newStringSet("C"),
			}},
			want: []string{"A", "B", "C", "D", "E"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := topoSort(tt.args.g)
			if err != nil {
				t.Errorf("toposort: %v", err)
				return
			}
			if len(got) > 0 && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topoSort(%v) = %v, want %v", tt.args.g, got, tt.want)
			}
		})
	}
}

func TestTopoSortNotUniqueResult(t *testing.T) {
	type args struct {
		g map[string]StringSet
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "multiple nodes",
			args: args{map[string]StringSet{
				"algorithms": newStringSet("data structures"),
				"calculus":   newStringSet("linear algebra"),

				"compilers": newStringSet(
					"data structures",
					"formal languages",
					"computer organization",
				),

				"data structures":       newStringSet("discrete math"),
				"databases":             newStringSet("data structures"),
				"discrete math":         newStringSet("intro to programming"),
				"formal languages":      newStringSet("discrete math"),
				"networks":              newStringSet("operating systems"),
				"operating systems":     newStringSet("data structures", "computer organization"),
				"programming languages": newStringSet("data structures", "computer organization"),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := topoSort(tt.args.g)
			if err != nil {
				t.Errorf("toposort: %v", err)
				return
			}
			if !isSorted(tt.args.g, got) {
				t.Errorf("topoSort(%v) = %v", tt.args.g, got)
			}
		})
	}
}

func TestTopoSortExistCycle(t *testing.T) {
	type args struct {
		g map[string]StringSet
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "two nodes",
			args: args{map[string]StringSet{
				"A": newStringSet("B"),
				"B": newStringSet("A"),
			}},
		},
		{
			name: "three nodes",
			args: args{map[string]StringSet{
				"A": newStringSet("B"),
				"B": newStringSet("C"),
				"C": newStringSet("A"),
			}},
		},
		{
			name: "multiple nodes",
			args: args{map[string]StringSet{
				"algorithms": newStringSet("data structures"),
				"calculus":   newStringSet("linear algebra"),

				"compilers": newStringSet(
					"data structures",
					"formal languages",
					"computer organization",
				),

				"data structures":       newStringSet("discrete math"),
				"databases":             newStringSet("data structures"),
				"discrete math":         newStringSet("intro to programming"),
				"formal languages":      newStringSet("discrete math"),
				"networks":              newStringSet("operating systems"),
				"operating systems":     newStringSet("data structures", "computer organization"),
				"programming languages": newStringSet("data structures", "computer organization"),
				"intro to programming":  newStringSet("networks"),
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := topoSort(tt.args.g)
			if err == nil {
				t.Errorf("not detected cycle")
			}
		})
	}
}
