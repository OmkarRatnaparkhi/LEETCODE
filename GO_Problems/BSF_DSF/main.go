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

//------------------------------------------------------------------------------------

func lowestCommonAncestor(root, p, q *Node) *Node {
    if root == nil || root == p || root == q {
        return root
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)

    if left != nil && right != nil {
        return root
    }
    if left != nil {
        return left
    }
    return right
}

func main() {
	// Example Tree:
	//        1
	//       / \
	//      2   3
	//     / \
	//    4   5

	//&Node{1,nil,nil}
	root := &Node{1,&Node{2,&Node{4,nil,nil},&Node{5,nil,nil}},&Node{3,nil,nil}}
	
	// Let's find LCA of node 4 and node 5 
	p := root.Left.Left 	//4
	q := root.Left.Right 	//5
	
	lca := lowestCommonAncestor(root,p,q)
	fmt.Printf("Lowest Common Ancestor of %d and %d is: %d\n", p.Val, q.Val, lca.Val)
}
