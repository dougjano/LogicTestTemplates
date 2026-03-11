package main

type Node struct {
	value      int
	leftChild  *Node
	rightChild *Node
}

func BuildTree(arr []int) *Node {
	root := Node{}

	for i, v := range arr {
		if i == 0 {
			root = Node{
				value: v,
			}
		} else {
			if v != -1 {
				loop := true
				searchPlace := &root

				for loop {
					if v < searchPlace.value {
						if searchPlace.leftChild == nil {
							searchPlace.leftChild = &Node{
								value: v,
							}
							loop = false
						} else {
							searchPlace = searchPlace.leftChild
						}

					} else if v > searchPlace.value {
						if searchPlace.rightChild == nil {
							searchPlace.rightChild = &Node{
								value: v,
							}
							loop = false
						} else {
							searchPlace = searchPlace.rightChild
						}
					}
				}
			}
		}
	}

	return &root
}

func TreeTrim(arr []int, root *Node) *Node {

	searchPlace := root

	for _, v := range arr {

		loop := true

		for loop {

			if arr[0] == searchPlace.value || arr[1] == searchPlace.value {
				searchPlace.leftChild = nil
				searchPlace.rightChild = nil
			}

			if searchPlace.leftChild != nil && arr[0] < searchPlace.leftChild.value && arr[1] < searchPlace.leftChild.value {
				searchPlace.leftChild = nil
			}

			if searchPlace.rightChild != nil && arr[0] > searchPlace.rightChild.value && arr[1] > searchPlace.rightChild.value {
				searchPlace.rightChild = nil
			}

			if v < searchPlace.value {
				if searchPlace.leftChild != nil {
					searchPlace = searchPlace.leftChild
				} else {
					loop = false
				}
			}

			if v > searchPlace.value {
				if searchPlace.rightChild != nil {
					searchPlace = searchPlace.rightChild
				} else {
					loop = false
				}
			}
		}
	}

	return root
}
