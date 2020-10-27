package relation

import (
	"image"
	"math"

	"goshape/pkg/geom"
	"goshape/pkg/goshape"
)

type RelationsManager struct {
	relations   []goshape.Relation
	roll goshape.Roll
}
var _ goshape.RelationsManager = &RelationsManager{}

func NewRelationsManager(roll goshape.Roll) goshape.RelationsManager {
	return RelationsManager{roll: roll}
}

func (r RelationsManager) ApplyRelations(shape geom.Shape, fixedPoints ...int) geom.Shape {
	for _, relation := range r.relations {
		shape = relation.Repair(shape, fixedPoints)
	}
	return shape
}

func (r RelationsManager) Draw(img *image.RGBA, shape geom.Shape) {
	for _, relation := range r.relations {
		relation.Draw(img, shape)
	}
}

func (r RelationsManager) AddRelation(toAdd goshape.Relation) goshape.RelationsManager{
	for _, rel := range r.relations {
		if toAdd.Collide(rel.GetPoints()...){
			return r
		}
	}
	r.relations = append(r.relations, toAdd)
	return r
}
func (r RelationsManager) DeleteRelation(shape geom.Shape, loc geom.Point) goshape.RelationsManager {
	closest := math.MaxFloat64
	closestIndex := -1
	for i, relation := range r.relations {
		rp := relation.GetMiddle(shape)
		d := geom.VecBetweenPoints(loc, rp).SqNorm()
		if d < closest {
			closestIndex = i
			closest = d
		}
	}
	if closestIndex == -1 {
		return r
	}
	r.relations = append(r.relations[:closestIndex], r.relations[closestIndex + 1:]...)
	return r
}

func (r RelationsManager) RemoveCollisions(pts ...int) goshape.RelationsManager{
	for i, relation := range r.relations {
		if relation.Collide(pts...){
			r.relations = append(r.relations[:i], r.relations[i + 1:]...)
		}
	}
	return r
}

func (r RelationsManager) IncrementOver(index int, roll goshape.Roll) goshape.RelationsManager{
	nr := r.relations
	for i, relation := range r.relations {
		nr[i] = relation.IncrementOver(index, roll)
	}
	return &RelationsManager{relations: nr, roll: roll}
}
