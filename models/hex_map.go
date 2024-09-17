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
	Description string
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

func (hexMap HexMap) GetEdges() [4][]Point {
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
		current = Point{starts[i][0], starts[i][1]}
		for !foundEdge {
			if mapCopy[current[0]][current[1]] == int(hextype.None) {
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

		if len_neighbors > 0 && len_neighbors < 6 {
			edge.Add(current)

			for _, neighbor := range neighbors {
				if !queue.Has(neighbor) {
					if mapCopy[neighbor[0]][neighbor[1]] != int(hextype.None) {
						queue.Add(neighbor)
					}
				}
			}
		}
	}

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
	neighborsStruct := CreatePointSet()
	neighbors := &neighborsStruct
	for _, move := range moves {
		newPoint := hexMap.Move(point, move)
		if !PointEquals(newPoint, point) {
			if hexMap[newPoint[0]][newPoint[1]].Type != hextype.None {
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

// https://stackoverflow.com/a/39868255
func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}