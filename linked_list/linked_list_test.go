package my_structs

import (
	"testing"

	node "github.com/kolaabaka/my_structs/node"
)

func TestLinkedList(t *testing.T) {
	var nFirst, nSecond, nThird, nFourth node.Node[int]
	linkedList := CreateLinkedList(&nFirst, &nSecond)
	if linkedList == nil || linkedList.LinkedListLength() != 2 {
		t.Error("crating linkedList doesn`t work")
	}

	linkedList.PushLast(&nThird)
	bufNoda, err := linkedList.Get(2)
	if err != nil || bufNoda != &nThird || linkedList.lastNode != &nThird {
		t.Error("pushLast doesn`t work")
	}

	linkedList.PushFirst(&nFourth)
	bufNoda, err = linkedList.Get(0)
	if err != nil || bufNoda != &nFourth || linkedList.firstNode != &nFourth {
		t.Error("pushFirst doesn`t work")
	}
}
