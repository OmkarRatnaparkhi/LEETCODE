package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BFS / Level Order Traversal
func bfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}
	fmt.Println("Lenght of queue:", len(queue))
	for len(queue) > 0 {
		fmt.Println("Lenght of queue:", len(queue))
		// Dequeue
		node := queue[0]
		queue = queue[1:]

		// Process node
		result = append(result, node.Val)

		// Enqueue children
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

func main() {
	// Example Tree:
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5

	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil},
		},
		&TreeNode{3, nil, nil},
	}

	result := bfs(root)

	fmt.Println(result)

}
