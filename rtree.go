package main

import (
	"github.com/dhconnelly/rtreego"
)

const (
	Dim = 2
	Min = 25
	Max = 50
)

var RTree *rtreego.Rtree

func NewRTree(obj ...*Object) {
	RTree = rtreego.NewTree(Dim, Min, Max)
	if obj != nil {
		AddToRTree(obj...)
	}
}

func AddToRTree(obj ...*Object) {
	for _, o := range obj {
		RTree.Insert(o)
	}
}
