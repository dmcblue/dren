package models

import (
	// "dmcblue/dren/models"
	hextype "dmcblue/dren/models/hex-type"
	"dmcblue/dren/utils"
	// "fmt"
	// "fmt"
)

// "fmt"
// "strings"

// "github.com/fatih/color"


type Hex struct {
	Type hextype.HexType
}
type HexMap [][]Hex


func CreateHexMap (width int, height int) HexMap {
	hexMap := make(HexMap, width)
	for x, _ := range hexMap {
		// zero is default value
		hexMap[x] = make([]Hex, height)
		for y, _ := range hexMap[x] {
			hexMap[x][y] = Hex{
				Type: hextype.Plain,
			}
		}
	}

	return hexMap
}

func (hexMap HexMap) Height() int {
	return len(hexMap[0])
}

func (hexMap HexMap) Width() int {
	return len(hexMap)
}

/*
    0
   ___
5 /   \ 1
 /     \ 
 \     /
4 \___/ 2
    3
*/
/* even-q
       ___
      /   \
  ___/ 1,0 \___
 /   \     /   \
/ 0,0 \___/ 2,0 \
\     /   \     /
 \___/ 1,1 \___/
     \     /
  0,3 \___/ 1,2
*/
func (hexMap HexMap) Move(position Point, move int) [2]int {
	nX := position[0]
	nY := position[1]
	switch move {
		case 0:
			// nX
			nY -= 1
		case 1:
			nX += 1
			nY -= 1
		case 2:
			nX += 1
		case 3:
			// nX
			nY += 1
		case 4: 
			// nX
			nX -= 1
		case 5:
			nX -= 1
			nY -= 1
	}
	
	if nX < 0 || nX >= hexMap.Width() { nX = position[0] }
	if nY < 0 || nY >= hexMap.Height() { nY = position[1] }

	return [2]int{nX, nY}
}

func (hexMap HexMap) GetEdges2() [4][]Point {
	// fmt.Println("GetEdges2")
	w := hexMap.Width()
	h := hexMap.Height()
	maxRank := min(w/2, h/2)
	edges := [4][]Point{
		make([]Point, 0),
		make([]Point, 0),
		make([]Point, 0),
		make([]Point, 0),
	}
	mapCopy := make([][]int, w)
	for x, _ := range mapCopy {
		mapCopy[x] = make([]int, h)
		for y, _ := range mapCopy[x] {
			mapCopy[x][y] = int(hexMap[x][y].Type)
		}
	}

	edgeStruct := CreatePointSet()
	edge := &edgeStruct
	queueStruct := CreatePointSet()
	queue := &queueStruct

	foundEdge := false
	i := 0
	for !foundEdge && i < maxRank {
		if mapCopy[i][i] == int(hextype.None) {
			i++
		} else {
			foundEdge = true
		}
	}
	mappings := [1]Point{
		{1, 1},
		// {-1, 1},
		// {-1, -1},
		// {1, -1},
	}
	starts := [1]Point{
		{0, 0},
		// {hexMap.Width() - 1, 1},
		// {hexMap.Width() - 1, hexMap.Height() - 1},
		// {1, hexMap.Height() - 1},
	}
	var current Point
	for i, mapping := range mappings {
		foundEdge := false
		// i := 0
		current = Point{starts[i][0], starts[i][1]}
		for !foundEdge /* && i < maxRank */ {
			if mapCopy[current[0]][current[1]] == int(hextype.None) {
				// i++
				current[0] = current[0] + mapping[0]
				current[1] = current[1] + mapping[1]
			} else {
				foundEdge = true
			}
		}
		queue.Add(current)
		edge.Add(current)
	}

	for queue.Size() > 0 {
		current = queue.Pop()
		neighbors := hexMap.neighbors(current)
		len_neighbors := len(neighbors)
		// fmt.Println("loop: current", current, neighbors)

		if len_neighbors > 0 && len_neighbors < 6 {
			edge.Add(current)

			for _, neighbor := range neighbors {
				if !queue.Has(neighbor) {
					if mapCopy[neighbor[0]][neighbor[1]] != int(hextype.None) {
						// if len(hexMap.neighbors(neighbor)) > 0 {
							queue.Add(neighbor)
						// }
					}
				}
			}
		}
	}
	// fmt.Println("GetEdges2 B", len(edge.Set()))

	// pick edges
	mX := hexMap.Width()/2
	mY := hexMap.Height()/2
	for _, edgePoint := range edge.Set() {
		octant := 0
		v := Point{edgePoint[0] - mX, edgePoint[1] - mY}
		x := abs(v[0])
		y := abs(v[1])

		if v[0] < 0 {
			if v[1] < 0 {
				if x < y {
					octant = 1
				} else {
					octant = 0
				}
			} else {
				if x < y {
					octant = 6
				} else {
					octant = 7
				}
			}
		} else {
			if v[1] < 0 {
				if x < y {
					octant = 2
				} else {
					octant = 3
				}
			} else {
				if x < y {
					octant = 5
				} else {
					octant = 4
				}
			}
		}

		if octant == 1 || octant == 2 {
			// n
			edges[0] = append(edges[0], edgePoint)
		} else if octant == 3 || octant == 4 {
			// e
			edges[1] = append(edges[1], edgePoint)
		} else if octant == 5 || octant == 6 {
			// s
			edges[2] = append(edges[2], edgePoint)
		} else if octant == 7 || octant == 0 {
			// w
			edges[3] = append(edges[3], edgePoint)
		}
	}

	return edges
}

func abs(x int) int {
	return utils.TernaryInt(x < 0, -x, x)
}

func (hexMap HexMap) neighbors(point Point) []Point {
	moves := [6]int{0, 1, 2, 3, 4, 5}
	// neighbors := make([]Point, 0)
	neighborsStruct := CreatePointSet()
	neighbors := &neighborsStruct
	for _, move := range moves {
		newPoint := hexMap.Move(point, move)
		if !PointEquals(newPoint, point) {
			if hexMap[newPoint[0]][newPoint[1]].Type != hextype.None {
				// neighbors = append(neighbors, newPoint)
				neighbors.Add(newPoint)
			}
		}
	}

	return neighbors.Set()
}

func (hexMap HexMap) onEdge(point Point) bool {
	return point[0] == 0 || point[0] == hexMap.Width() - 1 ||
		point[1] == 0 || point[1] == hexMap.Height() - 1
}






// func (hexMap HexMap) GetEdges() [4][]Point {
// 	w := hexMap.Width()
// 	h := hexMap.Height()
// 	maxRank := min(w/2, h/2)
// 	edges := [4][]Point{
// 		make([]Point, w),
// 		make([]Point, h),
// 		make([]Point, w),
// 		make([]Point, h),
// 	}

// 	noEmpties := false
// 	rank := 0

// 	for rank < maxRank && !noEmpties {
// 		noEmpties = true
// 		horz := [2]int{rank, w - 1 - rank}
// 		vert := [2]int{rank, h - 1 - rank}

// 		// north
// 		for i := horz[0]; i < horz[1]; i++ {
// 			if hexMap[i][rank].Type == hextype.None {
// 				noEmpties = false
// 			} else {
// 				edges[0] = append(edges[0], Point{i, rank})
// 			}
// 		}
// 		// east
// 		for i := vert[0]; i < vert[1]; i++ {
// 			if hexMap[w - 1 - rank][i].Type == hextype.None {
// 				noEmpties = false
// 			} else {
// 				edges[1] = append(edges[1], Point{w - 1 - rank, i})
// 			}
// 		}
// 		// south
// 		for i := horz[0]; i < horz[1]; i++ {
// 			if hexMap[i][h - 1 - rank].Type == hextype.None {
// 				noEmpties = false
// 			} else {
// 				edges[2] = append(edges[2], Point{i, h - 1 - rank})
// 			}
// 		}
// 		// west
// 		for i := vert[0]; i < vert[1]; i++ {
// 			if hexMap[rank][i].Type == hextype.None {
// 				noEmpties = false
// 			} else {
// 				edges[1] = append(edges[1], Point{rank, i})
// 			}
// 		}

// 		rank++
// 	}
	

// 	return edges
// }

/*
    0
   ___
5 /   \ 1
 /     \ 
 \     /
4 \___/ 2
    3
*/
// func (hexMap HexMap) GetEdge(direction int) []Point {
// 	w := hexMap.Width()
// 	h := hexMap.Height()

// 	// r := makeRange(0, utils.TernaryInt(direction%2 == 0, w, h))
// 	// d := utils.TernaryInt(direction%2 == 0, w, h)
// 	// queue := make([]Point, d)
// 	// for i := 0; i < d; i++ {
// 	// 	point := Point{0, 0}
// 	// 	switch direction {
// 	// 		case 0:
// 	// 			point = Point{i, 0}
// 	// 		case 1:
// 	// 			point = Point{h - 1, i}
// 	// 		case 2:
// 	// 			point = Point{i, w -1}
// 	// 		case 3:
// 	// 			point = Point{0, i}
// 	// 	}
// 	// 	queue = append(queue, point)
// 	// }

// 	// for len(queue) > 0 {
// 	// 	point, queue = queue[0], queue[1:]
// 	// 	current := Point{point[0], point[1]}

// 	// }



// 	points := make([][2]int, utils.TernaryInt(direction%2 == 0, w, h))
// 	// current := [2]int{0, 0}
// 	// moves := [6]int{5, 0, 1, 2, 3, 4}
// 	// switch direction {
// 	// 	case 1: // East
// 	// 		current = [2]int{w - 1, 0}
// 	// 		moves = [6]int{1, 2, 3, 4, 5, 0}
// 	// 	case 2: // South
// 	// 		current = [2]int{0, h - 1}
// 	// 		moves = [6]int{4, 3, 2, 1, 0, 5}
// 	// 	case 3:
// 	// 		moves = [6]int{5, 4, 3, 2, 1, 0}
// 	// }
// 	// points = append(points, [2]int{current[0], current[1]})
// 	// ended := false
// 	// for !ended {
// 	// 	moveIndex := 0
// 	// 	move := hexMap.Move(current, moves[moveIndex])
// 	// 	for moveIndex < 6 && PointEquals(current, move) {
// 	// 		moveIndex++
// 	// 		move = hexMap.Move(current, moves[moveIndex])
// 	// 	}
// 	// 	current = move
// 	// 	if moveIndex == 6 {
// 	// 		ended = true
// 	// 	}
// 	// 	points = append(points, [2]int{current[0], current[1]})
// 	// }

// 	return points
// }

// https://stackoverflow.com/a/39868255
func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}