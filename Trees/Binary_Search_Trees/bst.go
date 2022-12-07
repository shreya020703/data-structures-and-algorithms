/*
	Binary Search tree implementation
	In computer science, a binary search tree (BST), also called an ordered or sorted binary tree, is a rooted binary tree data structure
	with the key of each internal node being greater than all the keys in the respective node's left subtree and less than the ones
	in its right subtree. The time complexity of operations on the binary search tree is directly proportional to the height of the tree.

	The complexity analysis of BST shows that, on average, the insert, delete and search
	takes O(log n) fot n nodes. In the worst case, they degrade to that of a singly linked list O(n).

	Source(https://en.wikipedia.org/wiki/Binary_search_tree)

	Opeartions on BST:
	1 Find/ Find Minimum / Find Maximum element in binary search trees
	2 Inserting an element in binary search trees
	3 Deleting an element from binary search trees

*/
package main

import (
	"fmt"
	"math"
	"math/rand"
)

type BSTNode struct {
	left *BSTNode
	data int
	right *BSTNode
}

// ConstructBST returns a new, random binary tree
func ConstructBST(n, k int) *BSTNode {
	var root *BSTNode
	for _, v := range rand.Perm(n) {
		root = Insert(root, (1 + v) * k)
	}
	return root
}

// Insert, inserts an element in binary tree
func Insert(root *BSTNode, data int) *BSTNode {
	if root == nil {
		return &BSTNode{nil, data, nil}
	}
	if data < root.data {
		root.left = Insert(root.left, data)
		return root
	}
	root.right = Insert(root.right, data)
	return root
}

// Time Complexity: O(n). Space Complexity: O(n).
// Approach: start with root and keep moving left or right recursively
// if data we are searching is same as current node's data then return current node
// if data is less than node's data then resurse left, else recurse right
// if no data then it will return nil
func SearchElementRecursive(root *BSTNode, data int) *BSTNode {
	if root == nil {
		return root
	}
	if data < root.data {
		return SearchElementRecursive(root.left, data)
	} else if data > root.data {
		return SearchElementRecursive(root.right, data)
	}
	return root
}

// Time Complexity: O(n). Space Complexity: O(1).
// Approach: start with root and keep moving left or right
// if data we are searching is same as current node's data then return current node
// if data is less than node's data then resurse left, else recurse right
// if no data then it will return nil
func SearchElementNonRecursive(root *BSTNode, data int) *BSTNode {
	if root == nil {
		return root
	}
	for root != nil {
		if data == root.data {
			return root
		} else if data < root.data {
			root = root.left
		} else {
			root = root.right
		}
	}
	return nil
}

// FindMin: finds min element in BST,
// Approach:  start with root and keep moving left
// Min element is the left most node in BST
// Time Complexity: O(n). Space Complexity: O(n).
func FindMinRecursive(root *BSTNode) *BSTNode {
	if root == nil {
		return nil
	} else if root.left == nil {
		return root
	} else {
		return FindMinRecursive(root.left)
	}
}

// FindMIn: finds max element in BST,
// Approach:  start with root and keep moving right
// Max element is the right most node in BST
// Time Complexity: O(n). Space Complexity: O(1).
func FindMinNonRecursive(root *BSTNode) *BSTNode {
	if root == nil {
		return root
	}
	for root.left != nil {
		root = root.left
	}
	return root
}

// FindMax: finds min element in BST,
// Approach:  start with root and keep moving left
// min element is the left most node in BST
// Time Complexity: O(n). Space Complexity: O(n).
func FindMaxRecursive(root *BSTNode) *BSTNode {
	if root == nil {
		return nil
	}
	if root.right == nil {
		return root
	}
	return FindMaxRecursive(root.right)
}

// FindMIn: finds max element in BST,
// Approach:  start with root and keep moving right
// Max element is the right most node in BST
// Time Complexity: O(n). Space Complexity: O(1).
func FindMaxNonRecursive(root *BSTNode) *BSTNode {
	if root == nil {
		return nil
	}
	for root.right != nil {
		root  = root.right
	}
	return root
}
func DeleteMin(root *BSTNode) *BSTNode{
	if root.left == nil {
		return root.right
	}
	root.left = DeleteMin(root.left)
	return root
}

// DeleteFromBST deleted given value from bst and doesn't break the node
// Time Complexity: O(n). Space Complexity: O(n).
func DeleteFromBST(root *BSTNode, data int) *BSTNode {
	if root == nil {
		return nil
	}
	// traverse left or right recursively until you find data
	if data < root.data {
		DeleteFromBST(root.left, data)
	} else if data > root.data {
		DeleteFromBST(root.right, data)
	} else {
		// straight forward conddition no child set root as nil and return root
		if root.right == nil && root.left == nil{
			fmt.Println("hit")
			root = nil
			return root
		}
		// if left is nil and right has children set root as root.right 
		if root.left == nil && root.right != nil {
			temp := root.right
			root = temp
			return root
		}
		// if right is nil and left has children set root as root.left
		if root.right == nil && root.left != nil {
			fmt.Println(" boom")
			temp := root.left
			root = temp
			return root 
		}
		// if both left and right have children
		// find min in the right sub tree, 
		// set root.data as minimum value you found you found in right sub tree
		// call DeleteFromBST and delete min value
		temp := FindMinNonRecursive(root.right)
		DeleteMin(root.right)
		root.data = temp.data
		root.right = DeleteFromBST(root.right, temp.data)
		
	}
	return root
}


// Pre-order traversal
// Preorder traversal is defined as follows:
// 1 Visit the root.
// 2 Traverse the left subtree in Preorder.
// 3 Traverse the right subtree in Preorder.
// Time Complexity: O(n). Space Complexity: O(n).
// The nodes of tree would be visited in the order: 1 2 4 5 3 6 7
func PreOrder(root *BSTNode) {
	if root == nil {
		return
	}
	fmt.Print(root.data)
	PreOrder(root.left)
	PreOrder(root.right)
}
// IsBST: checks whether a given tree is a valid BST or not
// Time Complexity: O(n2). Space Complexity: O(n).
func IsBST(root *BSTNode) bool {
	if root == nil {
		return true
	}
	// if max on left tree is greater then root then return false
	max := FindMaxRecursive(root.left)
	if root.left != nil && max.data > root.data {
		return false
	}
	// if min on right tree is less then root then return false
	min := FindMinNonRecursive(root.right)
	if root.right != nil && min.data < root.data {
		return false
	}
	// recursively check if left or right is not a BST
	if !IsBST(root.left) || !IsBST(root.right) {
		return false
	}
	return true
}
// Time Complexity: O(n). Space Complexity: O(n).
func IsBSTOptimal(root *BSTNode, min, max int) bool {
	if root == nil {
		return true
	}
	return (
		root.data > min && 
		root.data < max && 
		IsBSTOptimal(root.left, min, root.data) &&  
		IsBSTOptimal(root.right, root.data, max))
}

// IsBSTInorder: Using Inorder traversal we can solve this problem 
// Inorder traversal gives us sorted values, while traversing the BST in inorder,
// at each node check the condition that its key value should be greater than 
// the key value of its previous visited node, initialize prev with minimum integer value
// Time Complexity: O(n). Space Complexity: O(n).
func IsBSTInorder(root *BSTNode, prev *int) bool {
	if root == nil {
		return true
	}
	if !IsBSTInorder(root.left, prev) {
		return false
	}
	// compare root value with prev visited value
	if root.data < *prev {
		return false
	}
	// set prev value as root's value
	*prev = root.data
	return IsBSTInorder(root.right, prev)
}

func main() {
	tree := ConstructBST(10, 1)
	fmt.Println(tree)
	fmt.Println(SearchElementRecursive(tree, 5))
	fmt.Println(SearchElementNonRecursive(tree, 5))
	fmt.Println(FindMinRecursive(tree))
	fmt.Println(FindMinNonRecursive(tree))
	fmt.Println(FindMaxRecursive(tree))
	fmt.Println(FindMaxNonRecursive(tree))
	PreOrder(tree)
	
	tree = DeleteFromBST(tree, 6)
	fmt.Println()
	PreOrder(tree)
	//fmt.Println(SearchElementRecursive(tree, 7))
	fmt.Print(IsBST(tree))
	fmt.Print(IsBSTOptimal(tree, math.MinInt32, math.MaxInt32))
	var prev = math.MinInt32
	fmt.Println(IsBSTInorder(tree, &prev))
}