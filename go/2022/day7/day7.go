package day7

import (
	ds "aoc/datastructures"
	fs "aoc/filesystem"
	"fmt"
	"strconv"
	"strings"
)

func parseCommand(input string) (fs.Command, []string) {
	cmdString := input[2:]
	args := strings.Split(cmdString, " ")

	cmd, err := fs.NewCommand(args[0])
	if err != nil {
		fmt.Printf("unknown command: %v for cmdString %v", cmd, cmdString)
		panic(err)
	}

	return cmd, args[1:]
}

func parseInput(input []string) *fs.Filesystem {

	filesystem := fs.NewFilesystem()

	for _, item := range input {
		switch {
		// Is command
		case strings.Index(item, "$") == 0:
			cmd, args := parseCommand(item)
			filesystem.Execute(cmd, args)
		// Create Folder
		case strings.Index(item, "dir") == 0:
			dirName := item[4:]
			filesystem.CreateFolder(dirName)
		// Create File
		default:
			splitIdx := strings.Index(item, " ")
			size, _ := strconv.Atoi(item[0:splitIdx])
			name := item[splitIdx+1:]
			filesystem.CreateFile(name, size)
		}
	}
	return filesystem
}

func Part1(input []string) int {
	filesystem := parseInput(input)
	filesystem.MoveToRoot()

	//filesystem.Print()

	total := 0
	filesystem.Traverse(func(node *ds.TreeNode[fs.File], level int) {
		if !node.Value.IsFolder {
			return
		}
		size := filesystem.Size(".")
		if size > 100000 {
			return
		}
		total += size
	})

	return total
}

func Part2(input []string) int {

	filesystem := parseInput(input)
	filesystem.MoveToRoot()
	filesystem.Print()

	totalDiskSpace := 70000000
	updateSize := 30000000
	usedDiskSpace := filesystem.Size(".")

	availableSpace := totalDiskSpace - usedDiskSpace
	minNeededSpace := updateSize - availableSpace

	fmt.Printf("min needed: %v", minNeededSpace)

	optionsToDelete := make(map[int]string)
	filesystem.Traverse(func(node *ds.TreeNode[fs.File], level int) {
		if !node.Value.IsFolder {
			return
		}
		size := filesystem.Size(".")
		if size >= minNeededSpace {
			optionsToDelete[size] = node.Value.Name
		}
	})

	min := usedDiskSpace
	for k, _ := range optionsToDelete {
		if k < min {
			min = k
		}
	}

	return min
}
