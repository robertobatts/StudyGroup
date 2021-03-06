package main

type Tree struct {
	Value       int
	Left, Right *Tree
}

func Sort(values []int) []int {
	var root *Tree
	for _, value := range values {
		root = Add(root, value)
	}
	return AppendValues(values[:0], root)
}

func AppendValues(values []int, t *Tree) []int {
	if t != nil {
		values = AppendValues(values, t.Left)
		values = append(values, t.Value)
		values = AppendValues(values, t.Right)
	}
	return values
}

func Add(tree *Tree, value int) *Tree {
	if tree == nil {
		tree = new(Tree)
		tree.Value = value
		return tree
	}
	if value < tree.Value {
		tree.Left = Add(tree.Left, value)
	} else {
		tree.Right = Add(tree.Right, value)
	}
	return tree
}
