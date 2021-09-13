package tree

type BaseTreeNode struct {
	Val   int
	Left  *BaseTreeNode
	Right *BaseTreeNode
}

// var index int

func NewBaseTreeNode(arr ...int) *BaseTreeNode {
	return newBaseTreeNode(arr)
}

func newBaseTreeNode(arr []int) *BaseTreeNode {
	length := len(arr)
	if length == 0 {
		return nil
	}
	root := &BaseTreeNode{
		Val: arr[0],
	}
	
	queue := make([]*BaseTreeNode, 1, length*2)
	queue[0] = root
	i := 0
	j := 1
	if j < length {
		node := queue[i]
		
		if j < length && arr[j] != 0 {
			node.Left = &BaseTreeNode{
				Val: arr[j],
			}
			queue = append(queue, node.Left)
		}
		j++
		
		if j < length && arr[j] != 0 {
			node.Right = &BaseTreeNode{
				Val: arr[j],
			}
		}
		j++
		i++
	}
	return root
}

// func newBaseTreeNode(arr []int) *BaseTreeNode {
// 	length := len(arr)
// 	if length == 0 {
// 		return nil
// 	}
// 	index = 0
// 	root := AddNode(arr)
// 	return root
// }

// func AddNode(arr []int) *BaseTreeNode {
// 	if greaterEqual(index, len(arr)) || arr[index] == 0 {
// 		index += 1
// 		return nil
// 	}
// 	tree := &BaseTreeNode{Val: arr[index]}
// 	index += 1
// 	tree.Left = AddNode(arr)
// 	tree.Right = AddNode(arr)
// 	return tree
// }

func less(a, b int) bool {
	if a < b {
		return true
	}
	return false
}

func lessEqual(a, b int) bool {
	if a <= b {
		return true
	}
	return false
}

func greater(a, b int) bool {
	if a > b {
		return true
	}
	return false
}

func greaterEqual(a, b int) bool {
	if a >= b {
		return true
	}
	return false
}

// func TestNewBaseTreeNode(t *testing.T) {
// 	tree1 := tree.NewBaseTreeNode([]int{1, 0, 1})
// 	fmt.Println(tree1)
// 	fmt.Println(tree1.Right)
// 	fmt.Println(tree1.Left)
// }
