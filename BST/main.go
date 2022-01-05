package main

import (
	"math"
	"strconv"
)

type TreeNode struct {
	Left  *TreeNode
	Value int
	Right *TreeNode
}

//二叉树通用框架
func bstExecute(root *TreeNode) {
	//root需要做什么?
	if root == nil {
		return
	}
	root.Value += 1

	//其他的不用root操心,抛给递归
	bstExecute(root.Left)
	bstExecute(root.Right)
}

//给二叉树所有值+1
func plusOne(root *TreeNode) {
	if root == nil {
		return
	}
	root.Value += 1
	plusOne(root.Left)
	plusOne(root.Right)
}

//二叉树 检查是否两棵树相同
func isSameTree(root1 *TreeNode, root2 *TreeNode) bool {
	//每个节点对比
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil && root2 != nil {
		return false
	}
	if root1.Value != root2.Value {
		return false
	}

	//递归框架
	return isSameTree(root1.Left, root2.Left) &&
		isSameTree(root1.Right, root2.Right)
}

//二叉树 判断是否是binary search tree - BST
//粗略版,有bug
func isValidBST(root *TreeNode) bool {
	//每个节点做的事
	//如果节点为空,返回
	if root == nil {
		return true
	}
	if root.Left != nil && root.Value <= root.Left.Value {
		return false
	}
	if root.Right != nil && root.Value >= root.Right.Value {
		return false
	}

	//遍历二叉树框架
	return isValidBST(root.Left) && isValidBST(root.Right)
}

//正确的二叉树判断方法
func isValidBSTTrue(root *TreeNode) bool {
	return isValidBSTSup(root, nil, nil)
}

func isValidBSTSup(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Value <= min.Value {
		return false
	}
	if max != nil && root.Value >= max.Value {
		return false
	}
	//二叉树递归框架
	return isValidBSTSup(root.Left, min, root) && isValidBSTSup(root.Right, root, max)
}

//在二叉搜索树中找寻值是否存在
func isInBST(root *TreeNode, target int) bool {
	//找到树的最下层节点,依旧没有就是没有这个数
	if root == nil {
		return false
	}
	if root.Value == target {
		return true
	}
	//二叉树迭代框架
	return isInBST(root.Left, target) || isInBST(root.Right, target)
}

func isInBSTBer(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}
	if root.Value == target {
		return true
	}
	//使用特性优化的二叉树递归框架
	if root.Value < target {
		return isInBSTBer(root.Right, target)
	} else {
		return isInBSTBer(root.Left, target)
	}
}

//二叉搜索树插入新值
func insertBST(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{nil, value, nil}
	}
	//如果没有,就插入新节点
	//如果已经存在,就直接返回已有的
	if root.Value == value {
		return root
	}
	//遍历BST的框架
	if root.Value < value {
		root.Right = insertBST(root.Right, value)
	}
	if root.Value > value {
		root.Left = insertBST(root.Left, value)
	}

	return root
}

//二叉树节点统计
func countNodesNor(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodesNor(root.Left) + countNodesNor(root.Right)
}

func countNodesPef(root *TreeNode) int {
	//得到树的高度
	h := float64(0)
	for {
		if root == nil {
			break
		}
		root = root.Left
		h++
	}
	//高度2^高度 -1 就是节点总数
	return int(math.Pow(2, h)) - 1
}

func countNodesComp(root *TreeNode) int {
	//计算左右子树的高度
	l := root
	r := root
	hl := float64(0)
	hr := float64(0)
	for {
		if l == nil {
			break
		}
		l = l.Left
		hl++
	}
	for {
		if r == nil {
			break
		}
		r = r.Right
		hr++
	}
	//如果左右相同,则是perfect binary tree
	if hl == hr {
		return int(math.Pow(2, hl) - 1)
	}
	//左右不相同,是普通二叉树
	return 1 + countNodesPef(root.Left) + countNodesPef(root.Right)
}

func serialize(root *TreeNode) string {
	str := ""
	return serializeExe(root, str)
}

func serializeExe(root *TreeNode, str string) string {
	//前序遍历操作
	if root == nil {
		return str + "#,"
	}
	str += strconv.Itoa(root.Value) + ","

	//二叉树遍历框架
	str = serializeExe(root.Left, str)
	str = serializeExe(root.Right, str)
	return str
}

func deserializeEnd(root *TreeNode) string {
	return deserializeEndExe(root, "")
}
func deserializeEndExe(root *TreeNode, str string) string {
	if root == nil {
		return str + "#,"
	}
	//二叉树遍历框架
	str = deserializeEndExe(root.Left, str)
	str = deserializeEndExe(root.Right, str)

	//后续遍历节点操作
	str += strconv.Itoa(root.Value) + ","
	return str
}

func serializeMid(root *TreeNode) string {
	return serializeMidExe(root, "")
}
func serializeMidExe(root *TreeNode, str string) string {
	if root == nil {
		return str + "#,"
	}
	//遍历框架
	str = serializeMidExe(root.Left, str)
	str += strconv.Itoa(root.Value) + ","
	str = serializeMidExe(root.Right, str)
	return str
}

func main() {
}
