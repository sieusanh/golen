package BinaryTree

import ("fmt")

// https://vietcodes.github.io/algo/bst

// https://www.geeksforgeeks.org/remove-bst-keys-outside-the-given-range/



type Node struct {
	left *Node
	right *Node
	key int64
}

// cur: 20, go left, go right
// cur: 10, go left, go right
// cur:  5, go left, go right
// cur:  null, return left, go right
// cur:  null, return right, end of 5
// cur: return left, go right
// cur: 17, go left, go right
// cur: 15, go left, go right
// cur: return left, go right
// cur: null, return right, end of 15
// cur: 19, go left, go right
// cur: return left, go right
// cur: null, return right, end of 19
// cur: end of 17
// cur: end of 10
// cur: 35, go left, go right
// cur: 22, go left, go right,
// cur: null, go right 
// cur: 30, go left, go right
// cur: null, return right
// cur: null, end of 30
// cur: end of 22
// cur: 42, go left, go right
// cur: null, return right
// cur: null, end of 42

// 20, 10, 5, 17, 15, 19, 35, 22, 30, 42

func preorderTraversal(cur *Node) {
	if (cur == nil) {
		return
	}
	fmt.Println(cur.key)
	preorderTraversal(cur.left)
	preorderTraversal(cur.right)
}

func findNodeByKey(cur *Node, data int64) *Node{
	if (cur == nil) {
		return cur
	}

	if (cur.key == data) {
		return cur
	}

	var leftRightNode *Node

	if (data < cur.key) {
		leftRightNode = findNodeByKey(cur.left, data)
	} else if (data > cur.key) {
		leftRightNode = findNodeByKey(cur.right, data)
	}

	return leftRightNode
}

func insertNode(cur *Node, data int64) *Node {

	if (cur == nil) {
		// newNode := Node{key: data, left: nil, right: nil}
		// fmt.Println(newNode.data)
		// fmt.Println(newNode.left)
		// fmt.Println(newNode.right)
		// fmt.Println("===================")

		// fmt.Println("*cur before")
		// fmt.Println(cur)
		// cur = &Node{key: data}
		// cur = &Node{ key: data, left: nil, right: nil }
		// cur.key = data
		// cur.left = nil
		// cur.right = nil
		// cur = &Node{ key: data, left: nil, right: nil }
		// cur.key = data
		// cur = &newNode
		// fmt.Println("*cur after")
		// fmt.Println(cur)
		return &Node{ key: data, left: nil, right: nil }
	}

	if (data < cur.key) {
		cur.left = insertNode(cur.left, data)
	} else if (data > cur.key) {
		cur.right = insertNode(cur.right, data)
	}

	return cur
}

func trim(cur *Node, data int64) *Node {
   // Base Case
   if (root == NULL)
      return NULL;
 
   // First fix the left and right subtrees of root
   root->left =  removeOutsideRange(root->left, min, max);
   root->right =  removeOutsideRange(root->right, min, max);
 
   // Now fix the root.  There are 2 possible cases 
   // for root 1.a) Root's key is smaller than min
   // value (root is not in range)
   if (root->key < min)
   {
       node *rChild = root->right;
       delete root;
       return rChild;
   }
   // 1.b) Root's key is greater than max value
   // (root is not in range)
   if (root->key > max)
   {
       node *lChild = root->left;
       delete root;
       return lChild;
   }
   // 2. Root is in range
   return root;
}

func deleteNode() {
	
}

func change(p *string) {
	// var s string = "changed"s
	// p = &s
	*p = "can"
	// *val = "abc"
}

func Implement1() {
	var s string = "Hello"
	var p *string = &s
	fmt.Println("result before")
	fmt.Println(p)
	change(p)
	fmt.Println("result ")
	fmt.Println(*p)
	fmt.Println(s)
	// var p2 *int = nil
}

func Implement() {
	// var root *Node = &Node{key: 0, left: nil, right: nil}
	var root *Node
	root = insertNode(root, 20)	
	root = insertNode(root, 10)
	root = insertNode(root, 5)
	root = insertNode(root, 17)
	root = insertNode(root, 15)
	root = insertNode(root, 19)
	root = insertNode(root, 35)
	root = insertNode(root, 22)
	root = insertNode(root, 42)
	root = insertNode(root, 30)

	preorderTraversal(root)

	node := findNodeByKey(root, 19)

	fmt.Println("This node")
	fmt.Println(node)

	// var s string = "Hello"
	// var p *string = nil
	// fmt.Println("result before")
	// fmt.Println(p)
	// change(p)
	// fmt.Println("result ")
	// fmt.Println(p)
}	
