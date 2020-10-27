package goshape

import "goshape/pkg/geom"

type Roll func(int) int

type Repairer interface {
	Repair(shape geom.Shape, fixedPoints []int) geom.Shape
}

type Relation interface {
	Repairer
	Drawer
	IncrementOver(int, Roll) Relation
	GetMiddle(geom.Shape) geom.Point
	Collide(pts ...int) bool
	GetPoints()[]int

}

type RelationsManager interface {
	Drawer
	IncrementOver(int, Roll) RelationsManager
	ApplyRelations(shape geom.Shape, fixedPoints ...int) geom.Shape
	AddRelation(relation Relation) RelationsManager
	DeleteRelation(geom.Shape, geom.Point) RelationsManager
	RemoveCollisions(pts ...int) RelationsManager
}
