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


// ----------------------------------------------------------------------------------------------------
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Preorder DFS: Root -> Left -> Right				//PRE - root - left - right
func dfsPreorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val) // Process root
	dfsPreorder(root.Left, result)      // Go left
	dfsPreorder(root.Right, result)     // Go right
}

// Inorder DFS: Left -> Root -> Right				//Inorder - left - root - right
func dfsInorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	dfsInorder(root.Left, result)       // Go left
	*result = append(*result, root.Val) // Process root
	dfsInorder(root.Right, result)      // Go right
}

// Postorder DFS: Left -> Right -> Root				//Post - left - right - root
func dfsPostorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	dfsPostorder(root.Left, result)     // Go left
	dfsPostorder(root.Right, result)    // Go right
	*result = append(*result, root.Val) // Process root
}

func main() {
	// Example Tree:
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5

	//	root := &TreeNode{1,&TreeNode{2,&TreeNode{4,nil,nil},&TreeNode{5,nil,nil}},&TreeNode{3,nil,nil}}
	root := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4, nil, nil},
			&TreeNode{5, nil, nil},
		},
		&TreeNode{3, nil, nil},
	}

	// Run DFS traversals
	var preorder []int
	dfsPreorder(root, &preorder)
	fmt.Println("Preorder:", preorder) //Preorder:  [1 2 4 5 3]

	var  inorder []int
	dfsInorder(root, &inorder)
	fmt.Println("Inorder:", inorder) //Inorder:   [4 2 5 1 3]

	var postorder []int
	dfsPostorder(root, &postorder)
	fmt.Println("Postorder:", postorder)//Postorder: [4 5 2 3 1]

}

//traverse - graph ya tree problems

//sub array 
//repeated strings
//contiguous 
//for above three use slinding window problems
