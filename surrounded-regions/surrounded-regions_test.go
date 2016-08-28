package leetcode

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	xb := byte('X')
	ob := byte('O')
	fmt.Println(xb)
	fmt.Println(ob)

	//	data := [][]byte{
	//		[]byte{xb, xb, xb, },
	//		[]byte{xb, ob, xb, },
	//		[]byte{xb, xb, xb, },
	//	}

	//	data := [][]byte{
	//		[]byte("OOO"), []byte("OOO"), []byte("OOO"),
	//	}
	//	data := [][]byte{
	//		[]byte("XOXOXO"),
	//		[]byte("OXOXOX"),
	//		[]byte("XOXOXO"),
	//		[]byte("OXOXOX"),
	//	}

	data := [][]byte{
		[]byte("XXX"),
		[]byte("XOX"),
		[]byte("XXX"),
	}

	t.Log(data)
	solve(data)
	t.Log(data)
}
