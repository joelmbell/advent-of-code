package datastructures

//func TestNewTree(t *testing.T) {
//	tree := NewTree("/")
//	tree.Root.Add("a")
//	tree.Root.Add("14848514 b.txt")
//
//	if len(tree.Root.Children) != 2 {
//		t.Errorf("Tree should have two children.")
//	}
//}
//
//func TestTreeNode_FindChild(t *testing.T) {
//	tree := NewTree("/")
//	tree.Root.Add("a").Add("b").Add("c")
//	tree.Root.Add("x").Add("y").Add("z")
//	branch1Node := tree.Root.FindChild("a").FindChild("b").FindChild("c")
//	if branch1Node == nil || branch1Node.Value != "c" {
//		t.Errorf("did not find correct child: %v", branch1Node)
//	}
//
//	branch2Node := tree.Root.FindChild("x").FindChild("y").FindChild("z")
//	if branch2Node == nil || branch2Node.Value != "z" {
//		t.Errorf("did not find correct child: %v", branch2Node)
//	}
//}
