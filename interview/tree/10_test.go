package tree

import "testing"

func TestNewTreeNode(t *testing.T) {
	root := NewTreeNode(1, nil, nil)
	PrintTree(root)
}
