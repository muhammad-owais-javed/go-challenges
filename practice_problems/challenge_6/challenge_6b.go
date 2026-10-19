package main

import (
	"fmt"
	"strings"
)


type Node struct {
	Tag      string
	Children []*Node
}


func main(){

	root := &Node{
		Tag: "html",
		Children: []*Node{
			{
				Tag: "body",
				Children: []*Node{
					{
						Tag: "",
						Children: []*Node{
							{Tag: "p"},
							{Tag: "span"},
						},
					},
					{
						Tag: "section",
						Children: []*Node{
							{
								Tag: "div",
								Children: []*Node{
									{Tag: "p"},
								},
							},
						},
					},
					{
						Tag: "p",
						Children: []*Node{
							{
								Tag: "div",
								Children: []*Node{
									{Tag: "p"},
								},
							},
						},
					},
				},
			},
		},
	}


	fmt.Println(FindPaths(root, "p"))

}


func FindPaths(root *Node, tag string) []string {
    
	var results []string

	if root == nil {
		return nil
	}

	var startingPath []string

	walkTree(root, tag, startingPath, &results)

	return results
}

func walkTree(node *Node, targetTag string, currentPath []string, results *[]string) {
	// TODO: Step 2 - Add the current node's tag to the path (if it's not empty)

	if node.Tag != "" {
		currentPath = append(currentPath, node.Tag)
	}

	// TODO: Step 3 - Check if this node matches the targetTag, and save the path
	if node.Tag == targetTag {

		fullPathString := strings.Join(currentPath, ">")
		*results = append(*results, fullPathString)

	}

	// TODO: Step 4 - Loop through the node's children and call walkTree() on them!

	for _, child := range node.Children{
		walkTree(child, targetTag, currentPath, results)
	}

}