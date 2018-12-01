// сюда писать функцию DirTree

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	maxContentSize = 30
	newLine        = "\n"
	tab            = "\t"
	stickTab       = "│\t"
	stickLine      = "├───"
	angleLine      = "└───"
)

func setStick(cur string, angle bool) string {
	if angle {
		return cur + angleLine
	} else {
		return cur + stickLine
	}
}

func outputFormat(file os.FileInfo) string {

	if !file.IsDir() {
		var size string
		if file.Size() != 0 {
			size = " (" + strconv.FormatInt(file.Size(), 10) + "b" + ")"
		} else {
			size = " (empty)"
		}
		return file.Name() + size + newLine
	} else {
		return file.Name() + newLine
	}
}

func recursiveTree(mainDir string, printFiles bool, stick string, angle bool, newLine bool) string {
	var tree string

	parent, _ := os.Open(mainDir)
	parentDir, _ := parent.Stat()

	if !newLine && (parentDir.IsDir() || printFiles) {
		tree += setStick(stick, angle)
		tree += outputFormat(parentDir)
	}

	if parentDir.IsDir() {
		content, _ := parent.Readdir(maxContentSize)
		dirs := []os.FileInfo{} // slice of os.FileInfo structures

		if !printFiles { // if -f not inputed - only dirs
			for _, dir := range content {
				fmt.Println(dir.Name())
				if dir.IsDir() {
					dirs = append(dirs, dir)
				}
			}
		} else {
			dirs = content
		}
		current := stick
		if newLine {
		} else if angle {
			current += tab
		} else {
			current += stickTab
		}

		for i, dir := range dirs {
			tree += recursiveTree(mainDir+"/"+dir.Name(), printFiles, current, i+1 == len(dirs), false)
		}
	}
	return tree
}

func dirTree(out io.Writer, dir string, printFiles bool) error {
	tree := recursiveTree(dir, printFiles, "", false, true)
	out.Write([]byte(tree))

	return nil
}
