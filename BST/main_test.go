package main

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_bstExecute(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bstExecute(tt.args.root)
		})
	}
}

func Test_plusOne(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plusOne(tt.args.root)
		})
	}
}

func Test_isSameTree(t *testing.T) {
	//请求参数
	tree1 := &TreeNode{&TreeNode{nil, 1, nil}, 2, &TreeNode{nil, 3, nil}}
	tree2 := &TreeNode{&TreeNode{nil, 1, nil}, 2, &TreeNode{nil, 3, nil}}
	tree3 := &TreeNode{&TreeNode{nil, 1, nil}, 3, &TreeNode{nil, 3, nil}}
	//结构 表驱动
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//测试情况案例
		{"same", args{tree1, tree2}, true},
		{"not same", args{tree1, tree3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBST(t *testing.T) {
	//请求参数
	tree1 := &TreeNode{&TreeNode{nil, 1, nil}, 2, &TreeNode{nil, 3, nil}}
	tree2 := &TreeNode{&TreeNode{&TreeNode{nil, 1, nil}, 2, &TreeNode{&TreeNode{nil, 3, nil}, 4, nil}}, 5, &TreeNode{nil, 6, &TreeNode{nil, 7, nil}}}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"tree1 is BST", args{tree1}, true},
		{"tree2 is BST", args{tree2}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBSTTrue(t *testing.T) {
	//请求参数
	tree1 := &TreeNode{&TreeNode{nil, 1, nil}, 2, &TreeNode{nil, 3, nil}}
	tree2 := &TreeNode{&TreeNode{&TreeNode{nil, 1, nil}, 2, &TreeNode{&TreeNode{nil, 3, nil}, 4, nil}}, 5, &TreeNode{nil, 6, &TreeNode{nil, 7, nil}}}
	tree3 := &TreeNode{&TreeNode{nil, 5, nil}, 10, &TreeNode{&TreeNode{nil, 6, nil}, 15, &TreeNode{nil, 20, nil}}}

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"tree1 is BST", args{tree1}, true},
		{"tree2 is BST", args{tree2}, true},
		{"tree3 not BST", args{tree3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBSTTrue(tt.args.root); got != tt.want {
				t.Errorf("isValidBSTTrue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBSTSup(t *testing.T) {
	type args struct {
		root *TreeNode
		min  *TreeNode
		max  *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBSTSup(tt.args.root, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("isValidBSTSup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isInBST(t *testing.T) {
	//请求参数
	tree1 := &TreeNode{&TreeNode{nil, 1, nil}, 2, &TreeNode{nil, 3, nil}}
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"in", args{tree1, 1}, true},
		{"not in", args{tree1, 5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInBST(tt.args.root, tt.args.target); got != tt.want {
				t.Errorf("isInBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isInBSTBer(t *testing.T) {
	//请求参数
	tree1 := &TreeNode{&TreeNode{nil, 1, nil}, 2, &TreeNode{nil, 3, nil}}
	type args struct {
		root   *TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"in", args{tree1, 1}, true},
		{"not in", args{tree1, 5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInBSTBer(tt.args.root, tt.args.target); got != tt.want {
				t.Errorf("isInBSTBer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertBST(t *testing.T) {
	//请求参数
	tree1 := &TreeNode{&TreeNode{nil, 1, nil}, 2, &TreeNode{nil, 3, nil}}
	tree1Need := &TreeNode{&TreeNode{&TreeNode{nil, 0, nil}, 1, nil}, 2, &TreeNode{nil, 3, nil}}
	type args struct {
		root  *TreeNode
		value int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{"insert 0", args{tree1, 0}, tree1Need},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertBST(tt.args.root, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countNodesNor(t *testing.T) {
	//请求参数
	tree1 := &TreeNode{&TreeNode{nil, 1, nil}, 2, &TreeNode{nil, 3, nil}}
	tree2 := &TreeNode{nil, 2, nil}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"3 node tree", args{tree1}, 3},
		{"1 node tree", args{tree2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodesNor(tt.args.root); got != tt.want {
				t.Errorf("countNodesNor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countNodesPef(t *testing.T) {
	//请求参数
	tree1 := &TreeNode{nil, 2, nil}
	tree2 := &TreeNode{&TreeNode{nil, 0, nil}, 0, &TreeNode{nil, 0, nil}}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1 node tree", args{tree1}, 1},
		{"3 node tree", args{tree2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodesPef(tt.args.root); got != tt.want {
				t.Errorf("countNodesPef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countNodesComp(t *testing.T) {
	//请求参数
	tree1 := &TreeNode{nil, 2, nil}
	tree2 := &TreeNode{&TreeNode{nil, 0, nil}, 0, &TreeNode{nil, 0, nil}}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1 node tree", args{tree1}, 1},
		{"3 node tree", args{tree2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodesComp(tt.args.root); got != tt.want {
				t.Errorf("countNodesComp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serialize(t *testing.T) {
	//请求参数
	tree1 := &TreeNode{nil, 2, nil}
	tree2 := &TreeNode{&TreeNode{nil, 1, nil}, 2, &TreeNode{nil, 3, nil}}
	tree3 := &TreeNode{&TreeNode{&TreeNode{nil, 0, nil}, 1, nil}, 2, &TreeNode{nil, 3, nil}}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1 node tree", args{tree1}, "2,#,#,"},
		{"3 node tree", args{tree2}, "2,1,#,#,3,#,#,"},
		{"4 node tree", args{tree3}, "2,1,0,#,#,#,3,#,#,"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := serialize(tt.args.root); got != tt.want {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serializeExe(t *testing.T) {
	type args struct {
		root *TreeNode
		str  string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serializeExe(tt.args.root, tt.args.str)
		})
	}
}

func Test_deserializeEnd(t *testing.T) {
	//请求参数
	tree1 := &TreeNode{nil, 2, nil}
	tree2 := &TreeNode{&TreeNode{nil, 1, nil}, 2, &TreeNode{nil, 3, nil}}
	tree3 := &TreeNode{&TreeNode{&TreeNode{nil, 0, nil}, 1, nil}, 2, &TreeNode{nil, 3, nil}}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1 node tree", args{tree1}, "#,#,2,"},
		{"3 node tree", args{tree2}, "#,#,1,#,#,3,2,"},
		{"4 node tree", args{tree3}, "#,#,0,#,1,#,#,3,2,"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deserializeEnd(tt.args.root); got != tt.want {
				t.Errorf("deserializeEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deserializeEndExe(t *testing.T) {
	type args struct {
		root *TreeNode
		str  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deserializeEndExe(tt.args.root, tt.args.str); got != tt.want {
				t.Errorf("deserializeEndExe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serializeMid(t *testing.T) {
	//请求参数
	tree1 := &TreeNode{nil, 2, nil}
	tree2 := &TreeNode{&TreeNode{nil, 1, nil}, 2, &TreeNode{nil, 3, nil}}
	tree3 := &TreeNode{&TreeNode{&TreeNode{nil, 0, nil}, 1, nil}, 2, &TreeNode{nil, 3, nil}}
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1 node tree", args{tree1}, "#,2,#,"},
		{"3 node tree", args{tree2}, "#,1,#,2,#,3,#,"},
		{"4 node tree", args{tree3}, "#,0,#,1,#,2,#,3,#,"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := serializeMid(tt.args.root); got != tt.want {
				t.Errorf("serializeMid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_serializeMidExe(t *testing.T) {
	type args struct {
		root *TreeNode
		str  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := serializeMidExe(tt.args.root, tt.args.str); got != tt.want {
				t.Errorf("serializeMidExe() = %v, want %v", got, tt.want)
			}
		})
	}
}
