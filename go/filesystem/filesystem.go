package filesystem

import (
	ds "aoc/datastructures"
	"errors"
	"fmt"
)

type File struct {
	Name     string
	IsFolder bool
	Size     int
}

func (f File) Id() string {
	return f.Name
}

type Filesystem struct {
	tree    *ds.Tree[File]
	Current *ds.TreeNode[File]
}

func NewFilesystem() *Filesystem {
	tree := ds.NewTree(File{
		Name:     "/",
		IsFolder: true,
	})
	fs := &Filesystem{
		tree:    tree,
		Current: tree.Root,
	}
	return fs
}

type Command string

const (
	List Command = "ls"
	CD   Command = "cd"
)

func NewCommand(input string) (Command, error) {
	switch input {
	case "ls":
		return List, nil
	case "cd":
		return CD, nil
	default:
		return "", errors.New(fmt.Sprintf("unknown command: %v", input))
	}
}

func (fs *Filesystem) Execute(cmd Command, args []string) interface{} {
	switch cmd {
	case List:
		return fs.List()
	case CD:
		switch args[0] {
		case "/":
			fs.MoveToRoot()
		case "..":
			fs.Backward()
		default:
			fs.Forward(args[0])
		}

		return nil
	}
	return make([]File, 0)
}

func (fs *Filesystem) Forward(name string) (File, error) {

	found := fs.Current.FindChild(File{
		Name:     name,
		IsFolder: true,
	})

	if found == nil {
		return File{}, errors.New(fmt.Sprintf("Unable to find folder named %v within directory %v", name, fs.Current.Value.Name))
	} else {
		fs.Current = found
	}
	return fs.Current.Value, nil
}

func (fs *Filesystem) Backward() File {
	if fs.Current == fs.tree.Root {
		return fs.tree.Root.Value
	}

	fs.Current = fs.Current.Parent
	return fs.Current.Value
}

func (fs *Filesystem) MoveToRoot() {
	fs.Current = fs.tree.Root
}

func (fs *Filesystem) List() []File {
	files := make([]File, 0)

	for _, child := range fs.Current.Children {
		files = append(files, child.Value)
	}

	return files
}

func (fs *Filesystem) CreateFolder(name string) {
	fs.Current.Add(File{
		Name:     name,
		IsFolder: true,
		Size:     0,
	})
}

func (fs *Filesystem) CreateFile(name string, size int) {
	fs.Current.Add(File{
		Name:     name,
		IsFolder: false,
		Size:     size,
	})
}

func (fs *Filesystem) Size(file string) int {
	var nodeToIterate *ds.TreeNode[File]
	if file == "." {
		nodeToIterate = fs.Current
	} else {
		for _, child := range fs.Current.Children {
			if child.Value.Name == file {
				nodeToIterate = child
			}
		}
	}

	size := 0
	if nodeToIterate != nil {
		nodeToIterate.Iterate(func(node *ds.TreeNode[File], level int, isLast bool) {
			size += node.Value.Size
		}, 0)
	}

	return size
}

func (fs *Filesystem) Traverse(iterator func(node *ds.TreeNode[File], level int)) {
	start := fs.Current
	fs.Current.Iterate(func(node *ds.TreeNode[File], level int, isLast bool) {
		fs.Current = node
		iterator(node, level)
	}, 0)
	fs.Current = start
}

func (f File) Description() string {
	desc := "folder"
	if !f.IsFolder {
		desc = "file"
	}
	return fmt.Sprintf("%v: %v, size: %v", desc, f.Name, f.Size)
}

func (fs *Filesystem) Print() {
	fmt.Printf("\n\n")
	fs.Traverse(func(node *ds.TreeNode[File], level int) {
		indents := ""
		for i := 0; i < level; i++ {
			indents += "\t"
		}
		fmt.Printf("%s %v\n", indents, node.Value.Description())
	})
	fmt.Printf("\n\n")
}
