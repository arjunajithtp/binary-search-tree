package main

import (
	"fmt"
	"github.com/arjunajithtp/binary-search-tree/bst"
)

func main() {
	// example usage
	t := bst.BST{}
	t.Insert(100, 110, 150, 50, 25, 50, 75)
	fmt.Println("Preorder: ", t.CollectData(bst.Preorder))
	fmt.Println("Postorder: ", t.CollectData(bst.Postorder))
	fmt.Println("Inorder: ", t.CollectData(bst.Inorder))
}
