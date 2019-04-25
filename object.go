package main

//type ObjectCollection map[bool][]Object
//
//type Object struct {
//	Position pixel.Vec
//	Bounds   pixel.Rect
//	Sprite   *pixel.Sprite
//}
//
//func NewObject(w, h float64, pos pixel.Vec, sprite *pixel.Sprite) Object {
//	return Object{
//		Position: pos,
//		Bounds:   pixel.R(0, 0, w, h).Moved(pos.Sub(pixel.V(w/2, h/2))),
//		Sprite:   sprite,
//	}
//}
//
//func (o *Object) Draw(target pixel.Target, scale float64) {
//	o.Sprite.Draw(target, pixel.IM.Scaled(pixel.ZV, scale).Moved(o.Position.Scaled(scale)))
//}
//
//func (obj ObjectCollection) Draw(target pixel.Target, scale float64) {
//	for _, obs := range obj {
//		for _, o := range obs {
//			o.Draw(target, scale)
//		}
//	}
//}
