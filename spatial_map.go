package main

import (
	"github.com/faiface/pixel"
)

var SpaceMap *SpatialMap

type SpatialMap struct {
	ObjectMap map[pixel.Vec][]*Object
	Size      int
}

func NewSpatialMap(size int) {
	SpaceMap = &SpatialMap{
		ObjectMap: make(map[pixel.Vec][]*Object),
		Size:      size,
	}
}

func (s *SpatialMap) Hash(point pixel.Vec) pixel.Vec {
	return pixel.V(point.X/float64(s.Size), point.Y/float64(s.Size))
}

func (s *SpatialMap) InsertPoint(point pixel.Vec, obj *Object) {
	p := s.Hash(point)
	objList, ok := s.ObjectMap[p]
	if !ok {
		s.ObjectMap[p] = []*Object{obj}
	} else {
		s.ObjectMap[p] = append(objList, obj)
	}
}

func (s *SpatialMap) InsertBox(box pixel.Rect, obj *Object) {
	pMin := s.Hash(box.Min)
	pMax := s.Hash(box.Max)

	for i := pMin.X; i < pMax.X+1; i++ {
		for j := pMin.Y; j < pMax.Y; j++ {
			point := pixel.V(i, j)
			objList, ok := s.ObjectMap[point]
			if !ok {
				s.ObjectMap[point] = []*Object{obj}
			} else {
				s.ObjectMap[point] = append(objList, obj)
			}
		}
	}
}

func (s *SpatialMap) Search(box pixel.Rect) (objList []*Object, err error) {
	pMin := s.Hash(box.Min)
	pMax := s.Hash(box.Max)

	for i := pMin.X; i < pMax.X+1; i++ {
		for j := pMin.Y; j < pMax.Y; j++ {
			point := pixel.V(i, j)
			objs, ok := s.ObjectMap[point]
			if ok {
				objList = append(objList, objs...)
			}
		}
	}

	return
}
