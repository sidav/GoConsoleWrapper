package main

func randInRange(from, to int) int { //should be inclusive
	if to < from {
		t := from
		from = to
		to = t
	}
	if from == to {
		return from
	}
	return Random(to-from) + from // TODO: replace routines.random usage with package own implementation
}

type container struct {
	x, y, w, h int
}

type treeNode struct {
	parent, left, right *treeNode
	room                *container
}

type returningMap struct { //this struct is returned from generation routine.
	dmap []rune
}

func (r *container) returnCenter() (int, int) {
	return (r.x + r.w/2), (r.y + r.h/2)
}

func (m *returningMap) init() {
	m.dmap = make([]rune, MAP_W*MAP_H)
	for i := 0; i < len(m.dmap); i++ {
		m.dmap[i] = FLOOR
	}
}

func (m *returningMap) getCell(x, y int) rune {
	return m.dmap[x+MAP_W*y]
}

func (m *returningMap) setCell(cell rune, x, y int) {
	m.dmap[x+MAP_W*y] = cell
}

func getSplitRangeForPercent(wh int, percent int) (int, int) {
	min := wh * percent / 100
	return min, wh - min
}

func (t *treeNode) splitHoriz() { // splits node into "lower" and "upper"
	current_w := t.room.w
	current_h := t.room.h
	current_x := t.room.x
	current_y := t.room.y
	minSplSize, maxSplSize := getSplitRangeForPercent(current_h, SPLIT_MIN_RATIO)
	// Let's try to split the node without breaking min room size constraints
	for try := 0; try < TRIES_FOR_SPLITTING; try++ {
		upper_h := randInRange(minSplSize, maxSplSize)
		lower_h := current_h - upper_h + 1
		if upper_h < MIN_ROOM_H || lower_h < MIN_ROOM_H {
			continue
		} else { // Okay, sizes are acceptable. Let's do the split
			upperNode := treeNode{parent: t, room: &container{x: current_x, y: current_y, w: current_w, h: upper_h}}
			// Most error-probable place:
			lowerNode := treeNode{parent: t, room: &container{x: current_x, y: current_y + upper_h - 1, w: current_w, h: lower_h}}
			// hm... Left is upper and right is lower. Everything is obvious.
			t.left = &upperNode
			t.right = &lowerNode
			return
		}
	}
}

func (t *treeNode) splitVert() { // splits node into left and right
	current_w := t.room.w
	current_h := t.room.h
	current_x := t.room.x
	current_y := t.room.y
	minSplSize, maxSplSize := getSplitRangeForPercent(current_w, SPLIT_MIN_RATIO)
	// Let's try to split the node without breaking min room size constraints
	for try := 0; try < TRIES_FOR_SPLITTING; try++ {
		left_w := randInRange(minSplSize, maxSplSize)
		right_w := current_w - left_w + 1
		if left_w < MIN_ROOM_W || right_w < MIN_ROOM_W {
			continue
		} else { // Okay, sizes are acceptable. Let's do the split
			leftNode := treeNode{parent: t, room: &container{x: current_x, y: current_y, w: left_w, h: current_h}}
			// Most error-probable place:
			rightNode := treeNode{parent: t, room: &container{x: current_x + left_w - 1, y: current_y, w: right_w, h: current_h}}
			t.left = &leftNode
			t.right = &rightNode
			return
		}
	}
}

func (t *treeNode) splitNTimes(n int) {
	if n == 0 {
		return
	}
	toSplitOrNotToSplit := Random(100)
	if toSplitOrNotToSplit < SPLIT_PROBABILITY || t.room.w > MAX_ROOM_W || t.room.h > MAX_ROOM_H {
		horOrVert := Random(100)
		if horOrVert < HORIZ_PROBABILITY {
			t.splitHoriz()
		} else {
			t.splitVert()
		}
		if t.left != nil && t.right != nil { //if split was successful
			t.left.splitNTimes(n - 1)
			t.right.splitNTimes(n - 1)
		}
	}
}

func countOutsizedRooms(node *treeNode) int {
	total := 0
	if node.left == nil {
		if node.room.w > MAX_ROOM_W || node.room.h > MAX_ROOM_H {
			return 1
		}
		return 0
	} else {
		total += countOutsizedRooms(node.left)
		total += countOutsizedRooms(node.right)
	}
	return total
}

/////////////////////////////////////////

const (
	WALL = '#'
	RIVER = '~'
	DOOR = '+'
	FLOOR = ' '
	TRIES_FOR_SPLITTING  = 10
	TRIES_FOR_GENERATION = 1000
	MAX_OUTSIZED_ROOMS   = 5
)

var (
	MAP_W, MAP_H      int
	treeRoot          *treeNode
	SPLIT_PROBABILITY = 70 // in percent.
	SPLIT_MIN_RATIO   = 30 // in percent.
	MIN_ROOM_W        = 4
	MIN_ROOM_H        = 4
	MAX_ROOM_W        = 20 // this and next lines are not guaranteed. Think of them as a recommendations.
	MAX_ROOM_H        = 10 //
	HORIZ_PROBABILITY = 30 // in percent. Horiz splits should occur less frequently than vertical ones because of w > h
)

func GenerateDungeon(width, height, splits, sp_prob, sp_ratio, h_prob int) *returningMap {
	MAP_W = width
	MAP_H = height
	if splits == 0 {
		splits = 6
	}
	if sp_prob != 0 {
		SPLIT_PROBABILITY = sp_prob
	}
	if sp_ratio != 0 {
		SPLIT_MIN_RATIO = sp_ratio
	}
	if h_prob != 0 {
		HORIZ_PROBABILITY = h_prob
	}

	for i := 0; i < TRIES_FOR_GENERATION; i++ {
		// generate parent node
		treeRoot = &treeNode{room: &container{x: 0, y: 0, w: MAP_W, h: MAP_H}}
		// recursively split into rooms
		treeRoot.splitNTimes(splits)
		if countOutsizedRooms(treeRoot) > MAX_OUTSIZED_ROOMS {
			continue
		} else {
			break
		}
	}

	// init returning struct
	result := &returningMap{}
	result.init()

	renderTreeToDungeonMap(treeRoot, result)
	addDoorsForDungeonMap(treeRoot, result)
	addRiverForDungeonMap(result, 4)

	return result
}

func renderTreeToDungeonMap(node *treeNode, dmap *returningMap) {
	// recursively traverse through nodes and draw their containers
	if node.left != nil {
		renderTreeToDungeonMap(node.left, dmap)
		renderTreeToDungeonMap(node.right, dmap)
		return
	}
	for x := node.room.x; x < node.room.x+node.room.w; x++ {
		dmap.setCell(WALL, x, node.room.y)
		dmap.setCell(WALL, x, node.room.y+node.room.h-1)
	}
	for y := node.room.y; y < node.room.y+node.room.h; y++ {
		dmap.setCell(WALL, node.room.x, y)
		dmap.setCell(WALL, node.room.x+node.room.w-1, y)
	}
}

func addRiverForDungeonMap(dmap *returningMap, riverWidth int) {
	x := randInRange(MAP_W / 3, MAP_W * 2 / 3)
	bridgeYCoord := randInRange(1, MAP_H-1)
	bridgeHeight := 2
	for y:=0; y < MAP_H; y++ {
		dmap.setCell(FLOOR, x-1, y)
		dmap.setCell(FLOOR, x+riverWidth, y)
		for cx:=0; cx<riverWidth; cx++ {
			if y >= bridgeYCoord && y < bridgeYCoord + bridgeHeight {
				dmap.setCell(FLOOR, x+cx, y)
			} else {
				dmap.setCell(RIVER, x+cx, y)
			}
		}
		leftOrRight := randInRange(0, 5)
		if leftOrRight == 0 {
			x--
		}
		if leftOrRight == 1 {
			x++
		}
	}
}

// BUGGED! Rooms connectivity still not guaranteed!
func addDoorsForDungeonMap(node *treeNode, dmap *returningMap) {
	if node.left != nil {
		lx, ly := node.left.room.returnCenter()
		rx, ry := node.right.room.returnCenter()

		if ly == ry {
			// ly += randInRange(-MIN_ROOM_H/2, MIN_ROOM_H/2)
			for x := lx; x < rx; x ++ {
				if dmap.getCell(x, ly) == WALL {
					dmap.setCell(DOOR, x, ly)
					x += 3
				}
			}
		}
		if lx == rx {
			// lx += randInRange(-MIN_ROOM_W/2, MIN_ROOM_W/2)
			for y := ly; y < ry; y ++ {
				if dmap.getCell(lx, y) == WALL {
					dmap.setCell(DOOR, lx, y)
					y += 3
				}
			}
		}
		addDoorsForDungeonMap(node.left, dmap)
		addDoorsForDungeonMap(node.right, dmap)
	}
}
