package simulateHashmap

type KV struct {
	Key   string
	Value string
}

type LinkNode struct {
	Data     KV
	NextNode *LinkNode
}

func CreateLink() *LinkNode {
	var linkNode = &LinkNode{KV{"", ""}, nil}
	return linkNode
}

func (link *LinkNode) AddNode(data KV) int {
	var count = 0
	tail := link
	for {
		count += 1
		if tail.NextNode == nil {
			break
		} else {
			tail = tail.NextNode
		}
	}

	var newNode = &LinkNode{data, nil}
	tail.NextNode = newNode

	return count + 1
}
