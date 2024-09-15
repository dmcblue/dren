package models

import "fmt"

type Point [2]int

func PointEquals(a Point, b Point) bool {
	return a[0] == b[0] && a[1] == b[1]
}

type UniquePointQueue struct {
	set []Point
	hashMap map[string]bool
}

func CreatePointSet() UniquePointQueue {
	return UniquePointQueue{
		set: make([]Point, 0),
		hashMap: make(map[string]bool),
	}
}

func (upq *UniquePointQueue) Add(point Point) {
	key := upq.makeKey(point)
	// fmt.Println("key", key, upq.set, upq.hashMap[key])
	if !upq.hashMap[key] {
		upq.hashMap[key] = true
		upq.set = append(upq.set, point)
		// fmt.Println("key post", key, upq.set, upq.hashMap[key], upq)
	}
}

func (upq *UniquePointQueue) Has(point Point) bool {
	key := upq.makeKey(point)
	return upq.hashMap[key]
}

func (upq *UniquePointQueue) Size() int {
	return len(upq.set)
}

func (upq *UniquePointQueue) makeKey(point Point) string {
	return fmt.Sprintf("%dx%d", point[0], point[1])
}

func (upq *UniquePointQueue) Pop() (Point) {
	point := upq.set[0]
	set := upq.set[1:]
	upq.set = set

	return point
}

func (upq *UniquePointQueue) Set() []Point {
	return upq.set
}