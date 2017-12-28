package suanfa

import (
	"math"
	"fmt"
)

type tree_data struct {
	data int
}

type tree_node struct {
	num    int
	height int
	data tree_data
	parent *tree_node
	left *tree_node
	right *tree_node
}
//构造函数
func New_tree_node(num int,data tree_data,parent *tree_node) *tree_node {
	node := tree_node{}
	node.num = num
	node.data = data
	node.height = 0
	node.left = nil
	node.right = nil
	node.parent = parent
	return &node

}

//插入方法
func (t *tree_node) Insert(num int,data tree_data) {
	if num > t.num{
		if t.right != nil{
			t.right.Insert(num,data)
			//调整算法，向右子树插入，只能是右侧深度破坏条件
			if t.GetLength(t.left) > t.GetLength(t.right) + 1{
				if num > t.right.num{
					//当待插入的数大于右侧标签，为插入做节点的右子树，单旋转
					t.RightSimpleRotate()
				}else {
					//当插入的数小于右侧标签，为插入右节点的左子树，双旋转
					t.RightDoubleRotate()
				}
			}
		}else {
			//新节点不会导致条件破坏
			t.right = New_tree_node(num,data,t)
		}
	}else if num < t.num{
		if t.left != nil{
			t.left.Insert(num,data)
			//调整算法，向左子树插入，只能是左侧深度破坏条件
			if t.GetLength(t.left) > t.GetLength(t.right) +1{
				if num > t.left.num{
					//当待插入的数小于左侧标签，为插入左节点的左子树，单旋转
					t.LeftSimpleRotate()
				}else {
					//当待插入的数大于左侧标签，为插入左节点的左子树，双旋转
					t.LeftDoubleRotate()
				}
			}else {
				t.data = data
			}
			t.compute_height()
		}
	}
}

//获得深度
func (t *tree_node) GetLength(node *tree_node) int {
	if node == nil{
		return -1
	}else {
		return node.height
	}/*[图片上传中...(simple.png-2406ad)]*/
}

//单向旋转法
func (t *tree_node) LeftSimpleRotate() {
	if t.parent.left == t{
		t.parent.left = t.left
	}else {
		t.right = t.left
	}
	temp := t.left.right
	t.left.right = t
	t.left.parent = t.parent

	t.parent = t.left
	t.left = temp
}

func (t *tree_node) RightSimpleRotate() {
	if t.parent.left == t{
		t.parent.left = t.right
	}else {
		t.parent.right = t.right
	}
	temp := t.right.left
	t.right.left = t
	t.right.parent = t.parent

	t.parent = t.right
	t.right = temp
}

//双旋转
func (t *tree_node) LeftDoubleRotate() {
	t.left.RightSimpleRotate()
	t.LeftSimpleRotate()
}

func (t *tree_node) RightDoubleRotate() {
	t.right.LeftSimpleRotate()
	t.RightSimpleRotate()
}

//计算深度
func (t *tree_node) compute_height() {
	t.height = int(math.Max(float64(t.GetLength(t.left)),float64(t.GetLength(t.right))))+1
}

//遍历
func (t *tree_node) Visit(indent string) {
	fmt.Println(indent,t.num,t.GetLength(t.left),t.GetLength(t.right))
	if t.left != nil{
		t.left.Visit(indent + "   ")
	}
	if t.right != nil{
		t.right.Visit(indent + "  ")
	}
}