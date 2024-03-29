package distinctcolors

import (
	"fmt"
)

/*
type Node struct {
	index int32
	color int32
}

type Graph struct {
	nodes map[int32]*Node
	edges map[int32][]*Node
}

var graph Graph

func (g Graph) addEdge(from *Node, edge *Node) {
	if _, ok := g.nodes[from.index]; !ok {
		g.nodes[from.index] = from
	}
	g.edges[from.index] = append(g.edges[from.index], edge)
}

var pathFound bool
var pathCount int32
var rowCount int32
var visitedNodes map[int32]bool
var visitedColors map[int32]bool

func dfsCount(start, stop int32) {
	// mark node visited
	visitedNodes[start] = true
	currentColor := graph.nodes[start].color

	if start == stop {
		pathFound = true
		pathCount = 1
		// mark color visited
		visitedColors[currentColor] = true
		return
	}

	for _, edge := range graph.edges[start] {
		if _, visited := visitedNodes[edge.index]; !visited {
			dfsCount(edge.index, stop)
			if pathFound {
				// only count unique colors in this path
				if _, colorVisited := visitedColors[currentColor]; !colorVisited {
					pathCount += 1
				}
				// mark color visited
				visitedColors[currentColor] = true
				return
			}
		}
	}
}

func reset(size int32) {
	pathFound = false
	pathCount = 0
	visitedNodes = make(map[int32]bool, size)
	visitedColors = make(map[int32]bool, size)
}

// Complete the solve function below.
func solve(c []int32, tree [][]int32) []int32 {
	size := int32(len(c))
	graph = Graph{
		nodes: make(map[int32]*Node, size),
		edges: make(map[int32][]*Node, size),
	}

	for _, treeRow := range tree {
		key1 := treeRow[0] - 1
		key2 := treeRow[1] - 1

		node1 := &Node{index: key1, color: c[key1]}
		node2 := &Node{index: key2, color: c[key2]}

		graph.addEdge(node1, node2)
		graph.addEdge(node2, node1)
	}
	counts := []int32{}
	for start := range c {
		rowCount = 0
		for stop := range c {
			reset(size)
			dfsCount(int32(start), int32(stop))
			rowCount += pathCount
		}
		counts = append(counts, rowCount)
	}
	return counts
}

*/

type GraphMatrix struct {
	colors []int32
	edges  map[int32]map[int32]bool
}

var graphMatrix GraphMatrix

func (g *GraphMatrix) addEdge(from, fromColor, to, toColor, size int32) {
	g.colors[from] = fromColor
	g.colors[to] = toColor
	if len(g.edges[from]) < 1 {
		g.edges[from] = make(map[int32]bool, size)
	}
	g.edges[from][to] = true
}

var pathMemo map[string]int32

var pathFound bool
var pathDistinctColors map[int32]bool
var visitedNodes map[int32]bool

func matrixCount(start, stop, depth int32) {

	// TODO: memoize path subsets for faster lookup
	// if _, memo := pathMemo[strings.Join(start,stop)]; memo {

	// }

	visitedNodes[start] = true
	if start == stop {
		pathFound = true
		pathDistinctColors[graphMatrix.colors[start]] = true
		return
	}
	for next, edge := range graphMatrix.edges[start] {
		if edge {
			if _, visited := visitedNodes[next]; !visited {
				matrixCount(next, stop, depth+1)
				if pathFound {
					pathDistinctColors[graphMatrix.colors[next]] = true
				}
			}
		}
	}
}

// Complete the solve function below.
func solve(c []int32, tree [][]int32) []int32 {
	size := int32(len(c))
	graphMatrix = GraphMatrix{
		colors: make([]int32, size),
		edges:  make(map[int32]map[int32]bool, size),
	}

	for _, treeRow := range tree {
		key1 := treeRow[0] - 1
		key2 := treeRow[1] - 1

		graphMatrix.addEdge(key1, c[key1], key2, c[key2], size)
		graphMatrix.addEdge(key2, c[key2], key1, c[key1], size)
	}
	counts := []int32{}
	for start := range c {
		rowCount := int32(0)
		for stop := range c {
			pathFound = false
			pathDistinctColors = make(map[int32]bool)
			visitedNodes = make(map[int32]bool)
			matrixCount(int32(start), int32(stop), int32(0))
			rowCount += int32(len(pathDistinctColors))
		}
		counts = append(counts, rowCount)
	}
	return counts
}

func main() {
	// c := []int32{1, 2, 3, 2, 3}
	// tree := [][]int32{{1, 2}, {2, 3}, {2, 4}, {1, 5}}

	c := []int32{
		77, 89, 91, 34, 2, 26, 62, 19, 64, 16, 43, 7, 43, 72, 47, 39, 62, 47, 23, 78, 62, 4, 49, 35, 52, 28, 27, 3, 72, 24, 87, 49, 64, 29, 34, 65, 54, 96, 35, 70, 63, 77, 76, 58, 1, 74, 96, 14, 72, 19, 91, 34, 22, 91, 20, 26, 18, 46, 28, 42, 21, 66, 42, 36, 95, 75, 52, 100, 70, 39, 21, 33, 67, 49, 90, 67, 74, 37, 81, 46, 7, 71, 31, 29, 62, 2, 6, 31, 100, 85, 72, 72, 51, 13, 8, 97, 40, 11, 48, 9, 49, 21, 93, 16, 69, 34, 34, 42, 23, 14, 39, 29, 37, 21, 9, 50, 23, 14, 80, 74, 99, 4, 45, 1, 68, 52, 49, 7, 15, 48, 68, 15, 68, 12, 82, 88, 98, 16, 30, 20, 81, 20, 48, 17, 93, 9, 66, 67, 22, 98, 40, 72, 1, 84, 72, 20, 88, 72, 27, 2, 20, 46, 16, 87, 9, 50, 75, 6, 65, 56, 77, 45, 27, 77, 14, 19, 85, 79, 85, 6, 76, 24, 30, 28, 8, 1, 48, 95, 25, 26, 48, 44, 23, 15, 82, 31, 64, 8, 37, 28, 63, 13, 25, 42, 89, 38, 60, 25, 68, 45, 31, 96, 20, 60, 23, 27, 12, 22, 73, 36, 99, 72, 31, 21, 87, 13, 52, 2, 72, 40, 82, 35, 4, 6, 76, 45, 95, 87, 69, 62, 31, 51, 57, 51, 10, 32, 29, 74, 53, 54, 61, 52, 25, 92, 72, 63, 56, 75, 17, 27, 14, 98, 13, 18, 55, 40, 14, 49, 27, 82, 10, 9, 33, 19, 11, 94, 50, 92, 19, 54, 45, 80, 5, 21, 23, 29, 36, 30, 3, 52, 8, 69, 49, 21, 38, 55, 12, 51, 3, 38, 32, 64, 99, 16, 82, 61, 62, 83, 52, 80, 37, 48, 59, 93, 21, 33, 21, 56, 62, 24, 59, 70, 92, 59, 42, 29, 13, 5, 31, 15, 43, 14, 78, 93, 30, 12, 53, 91, 46, 5, 70, 34, 4, 81, 27, 24, 65, 47, 79, 79, 22, 37, 100, 65, 95, 41, 45, 7, 45, 75, 73, 39, 89, 3, 83, 70, 66, 36, 12, 11, 92, 33, 97, 47, 13, 23, 71, 30, 21, 1, 8, 43, 38, 7, 59, 84, 47, 4, 43, 43, 30, 67, 34, 70, 69, 16, 39, 34, 3, 2, 97, 46, 35, 93, 93, 47, 67, 63, 76, 87, 63, 35, 81, 52, 41, 40, 88, 39, 95, 82, 82, 76, 48, 67, 46, 17, 82, 36, 2, 37, 38, 50, 82, 72, 94, 26, 70, 60, 88, 98, 99, 3, 32, 79, 6, 25, 70, 93, 63, 16, 74, 96, 92, 22, 62, 89, 90, 96, 76, 43, 32, 65, 93, 65, 88, 86, 91, 10, 98, 30, 59, 96, 84, 42, 74, 90, 66, 96, 82, 29, 11, 8, 24, 54, 81, 38, 42, 70, 33, 70, 64, 16, 34, 56, 32, 74, 94, 22, 83, 91, 4, 41, 86, 87, 82, 11, 76, 100, 6, 10, 28, 69, 17, 3, 22, 49, 40, 16, 70, 24, 85, 33, 91, 70, 89, 23, 43, 34, 44, 77, 24, 99, 17, 61, 86, 51, 23, 13, 50, 29, 74, 29, 49, 42, 83, 70, 90, 75, 37, 59, 98, 21, 92, 89, 43, 32, 11, 85, 65, 6, 14, 40, 57, 30, 52, 94, 32, 74, 6, 33, 54, 32, 61, 2, 73, 44, 24, 63, 18, 60, 73, 67, 33, 64, 55, 75, 95, 17, 11, 11, 75, 24, 2, 31, 6, 53, 76, 37, 79, 33, 70, 32, 64, 82, 86, 37, 25, 9, 51, 94, 68, 23, 61, 100, 39, 67, 26, 33, 36, 37, 96, 62, 12, 49, 92, 17, 2, 67, 54, 80, 99, 75, 63, 15, 56, 48, 51, 33, 8, 53, 26, 28, 75, 38, 79, 13, 5, 5, 98, 92, 93, 45, 53, 4, 93, 96, 21, 46, 14, 26, 77, 64, 100, 40, 78, 7, 39, 80, 39, 47, 32, 17, 74, 59, 54, 52, 23, 10, 8, 72, 53, 100, 68, 5, 4, 61, 100, 24, 58, 65, 1, 35, 29, 100, 26, 6, 58, 64, 38, 49, 62, 21, 65, 87, 79, 70, 91, 54, 80, 98, 77, 84, 50, 45, 89, 53, 57, 40, 28, 14, 5, 28, 48, 85, 79, 25, 42, 36, 89, 79, 84, 2, 100, 48, 89, 30, 70, 31, 35, 1, 80, 12, 84, 29, 56, 24, 81, 12, 64, 8, 77, 20, 87, 77, 4, 65, 1, 97, 1, 89, 76, 36, 43, 27, 36, 31, 8, 5, 13, 43, 5, 92, 54, 88, 21, 61, 12, 1, 24, 27, 61, 100, 46, 99, 28, 1, 64, 29, 49, 16, 17, 24, 51, 11, 2, 86, 93, 10, 90, 5, 4, 46, 97, 9, 34, 17, 21, 97, 69, 44, 23, 81, 95, 20, 80, 23, 72, 95, 51, 20, 10, 19, 96, 12, 82, 49, 98, 74, 10, 39, 79, 13, 37, 75, 73, 22, 91, 93, 18, 11, 88, 92, 92, 35, 63, 23, 9, 86, 17, 11, 57, 26, 81, 52, 37, 62, 53, 86, 36, 62, 77, 14, 27, 13, 40, 51, 34, 82, 96, 51, 92, 83, 94, 83, 69, 56, 5, 77, 93, 73, 87, 49, 98, 68, 53, 87, 81, 5, 72, 16, 18, 48, 81, 96, 12, 20, 47, 45, 1, 42, 47, 93, 76, 92, 27, 45, 47, 84, 21, 91, 56, 60, 40, 6, 27, 92, 92, 59, 96, 15, 75, 13, 15, 55, 61, 26, 27, 7, 23, 27, 100, 69, 19, 75, 61, 98, 19, 59, 81, 40, 50, 88, 51, 41, 93, 77, 32, 36, 87, 27, 51, 61, 91, 17, 68, 51, 42, 94, 9, 16, 20, 8, 85, 91, 83, 97, 40, 53, 7, 72, 92, 8, 59, 94, 48, 4, 22, 79, 91, 9, 57, 93, 69, 100, 9, 88, 2, 3, 81, 11, 18, 53, 70, 54, 43, 4, 2, 34, 57, 9,
	}
	tree := [][]int32{
		{71, 511},
		{71, 795},
		{795, 731},
		{731, 346},
		{795, 128},
		{795, 3},
		{346, 732},
		{731, 264},
		{128, 204},
		{731, 318},
		{731, 802},
		{795, 277},
		{204, 921},
		{346, 947},
		{802, 47},
		{802, 257},
		{731, 502},
		{732, 390},
		{921, 151},
		{511, 811},
		{318, 850},
		{264, 215},
		{947, 168},
		{128, 388},
		{850, 375},
		{257, 660},
		{850, 254},
		{731, 400},
		{731, 656},
		{811, 607},
		{3, 640},
		{204, 222},
		{390, 85},
		{400, 48},
		{318, 750},
		{204, 986},
		{48, 819},
		{850, 796},
		{151, 766},
		{128, 777},
		{660, 49},
		{656, 994},
		{48, 281},
		{795, 714},
		{47, 897},
		{796, 243},
		{656, 797},
		{921, 801},
		{994, 91},
		{388, 432},
		{797, 673},
		{801, 235},
		{921, 18},
		{3, 698},
		{47, 191},
		{91, 325},
		{994, 879},
		{731, 868},
		{947, 472},
		{986, 481},
		{388, 643},
		{346, 932},
		{257, 242},
		{731, 397},
		{432, 201},
		{879, 572},
		{660, 910},
		{168, 869},
		{795, 330},
		{731, 394},
		{390, 514},
		{802, 707},
		{572, 176},
		{151, 784},
		{879, 741},
		{796, 981},
		{981, 421},
		{346, 828},
		{796, 182},
		{511, 143},
		{795, 492},
		{18, 316},
		{869, 139},
		{572, 905},
		{168, 617},
		{819, 676},
		{222, 900},
		{797, 422},
		{802, 86},
		{85, 70},
		{257, 496},
		{732, 561},
		{191, 763},
		{346, 465},
		{397, 227},
		{492, 642},
		{204, 528},
		{797, 568},
		{643, 509},
		{568, 545},
		{85, 407},
		{168, 927},
		{143, 954},
		{215, 39},
		{168, 381},
		{502, 170},
		{86, 526},
		{660, 297},
		{750, 886},
		{330, 517},
		{732, 476},
		{869, 671},
		{215, 9},
		{346, 172},
		{325, 530},
		{476, 885},
		{86, 735},
		{714, 232},
		{994, 56},
		{400, 184},
		{927, 383},
		{643, 891},
		{346, 148},
		{572, 798},
		{660, 966},
		{3, 535},
		{735, 513},
		{232, 626},
		{617, 813},
		{741, 924},
		{400, 144},
		{85, 770},
		{375, 532},
		{18, 175},
		{325, 935},
		{850, 137},
		{763, 14},
		{151, 521},
		{400, 164},
		{981, 374},
		{257, 256},
		{885, 40},
		{732, 46},
		{676, 843},
		{422, 889},
		{770, 413},
		{905, 393},
		{521, 109},
		{643, 571},
		{46, 79},
		{148, 659},
		{235, 160},
		{660, 810},
		{797, 226},
		{281, 367},
		{49, 248},
		{535, 372},
		{954, 614},
		{966, 533},
		{797, 583},
		{465, 76},
		{770, 125},
		{784, 365},
		{850, 92},
		{170, 518},
		{70, 89},
		{732, 54},
		{318, 341},
		{9, 620},
		{514, 562},
		{394, 730},
		{227, 253},
		{175, 541},
		{763, 762},
		{897, 327},
		{175, 911},
		{232, 163},
		{481, 890},
		{144, 352},
		{125, 259},
		{910, 475},
		{511, 82},
		{47, 313},
		{421, 718},
		{184, 444},
		{421, 761},
		{932, 101},
		{671, 956},
		{256, 16},
		{583, 83},
		{659, 13},
		{730, 317},
		{514, 458},
		{313, 121},
		{184, 687},
		{517, 976},
		{868, 776},
		{511, 785},
		{889, 455},
		{383, 814},
		{13, 115},
		{868, 75},
		{891, 867},
		{801, 613},
		{170, 621},
		{770, 579},
		{49, 699},
		{475, 280},
		{911, 706},
		{352, 948},
		{828, 252},
		{75, 506},
		{607, 24},
		{796, 380},
		{795, 631},
		{9, 609},
		{798, 173},
		{413, 477},
		{698, 7},
		{889, 740},
		{388, 939},
		{397, 681},
		{330, 159},
		{631, 946},
		{784, 593},
		{660, 404},
		{521, 157},
		{455, 214},
		{182, 634},
		{867, 686},
		{346, 224},
		{477, 161},
		{170, 369},
		{477, 493},
		{317, 283},
		{621, 124},
		{614, 677},
		{375, 898},
		{9, 127},
		{375, 998},
		{617, 765},
		{39, 755},
		{994, 844},
		{750, 751},
		{513, 892},
		{956, 906},
		{530, 296},
		{252, 611},
		{513, 414},
		{24, 255},
		{352, 486},
		{631, 278},
		{868, 856},
		{121, 181},
		{397, 169},
		{481, 834},
		{801, 596},
		{76, 456},
		{617, 705},
		{735, 153},
		{784, 309},
		{472, 345},
		{511, 637},
		{296, 565},
		{986, 772},
		{296, 494},
		{611, 733},
		{486, 149},
		{413, 334},
		{341, 781},
		{596, 576},
		{528, 108},
		{151, 10},
		{513, 862},
		{13, 635},
		{7, 807},
		{169, 489},
		{517, 749},
		{677, 273},
		{576, 337},
		{201, 138},
		{766, 799},
		{191, 1000},
		{181, 602},
		{981, 167},
		{40, 231},
		{802, 701},
		{677, 59},
		{718, 303},
		{3, 306},
		{777, 437},
		{533, 692},
		{571, 570},
		{167, 592},
		{421, 542},
		{562, 638},
		{935, 723},
		{214, 424},
		{316, 23},
		{494, 610},
		{642, 559},
		{254, 384},
		{889, 134},
		{384, 899},
		{413, 668},
		{731, 683},
		{784, 639},
		{570, 315},
		{900, 429},
		{761, 221},
		{921, 323},
		{184, 490},
		{296, 679},
		{138, 187},
		{109, 860},
		{656, 758},
		{947, 938},
		{699, 803},
		{810, 974},
		{472, 135},
		{974, 471},
		{966, 778},
		{659, 443},
		{224, 399},
		{777, 261},
		{309, 591},
		{862, 180},
		{493, 680},
		{692, 580},
		{576, 721},
		{562, 95},
		{634, 970},
		{621, 150},
		{807, 63},
		{860, 199},
		{614, 342},
		{784, 524},
		{443, 282},
		{810, 742},
		{323, 516},
		{494, 934},
		{424, 53},
		{680, 197},
		{167, 874},
		{456, 738},
		{583, 664},
		{313, 324},
		{778, 599},
		{583, 600},
		{899, 734},
		{706, 51},
		{766, 822},
		{924, 743},
		{297, 789},
		{721, 405},
		{48, 268},
		{621, 294},
		{677, 907},
		{701, 839},
		{405, 426},
		{303, 69},
		{740, 852},
		{443, 333},
		{541, 988},
		{517, 716},
		{761, 806},
		{399, 300},
		{63, 212},
		{869, 4},
		{935, 971},
		{634, 560},
		{151, 842},
		{148, 563},
		{259, 430},
		{59, 920},
		{394, 728},
		{637, 328},
		{807, 218},
		{755, 780},
		{124, 462},
		{799, 105},
		{394, 855},
		{668, 461},
		{76, 894},
		{47, 425},
		{47, 395},
		{542, 632},
		{798, 923},
		{346, 531},
		{89, 833},
		{413, 32},
		{281, 287},
		{889, 34},
		{780, 131},
		{149, 754},
		{173, 379},
		{138, 336},
		{600, 586},
		{728, 284},
		{948, 136},
		{868, 319},
		{296, 401},
		{255, 550},
		{278, 930},
		{530, 909},
		{325, 952},
		{511, 417},
		{325, 154},
		{939, 993},
		{197, 816},
		{924, 72},
		{281, 598},
		{383, 968},
		{367, 831},
		{524, 208},
		{874, 12},
		{692, 307},
		{517, 371},
		{639, 933},
		{212, 290},
		{294, 832},
		{844, 817},
		{524, 332},
		{127, 207},
		{109, 501},
		{252, 603},
		{899, 275},
		{741, 271},
		{850, 98},
		{513, 177},
		{570, 786},
		{75, 113},
		{679, 756},
		{341, 122},
		{802, 996},
		{92, 62},
		{819, 821},
		{993, 829},
		{521, 488},
		{444, 459},
		{252, 962},
		{894, 406},
		{511, 578},
		{345, 783},
		{10, 903},
		{417, 649},
		{115, 228},
		{51, 247},
		{924, 162},
		{319, 818},
		{443, 356},
		{475, 790},
		{47, 33},
		{599, 552},
		{856, 764},
		{591, 965},
		{786, 887},
		{109, 335},
		{417, 995},
		{224, 87},
		{404, 250},
		{795, 245},
		{867, 566},
		{821, 655},
		{756, 961},
		{388, 403},
		{131, 824},
		{86, 469},
		{69, 760},
		{154, 376},
		{169, 25},
		{889, 820},
		{300, 584},
		{280, 779},
		{476, 722},
		{796, 451},
		{723, 99},
		{563, 302},
		{517, 979},
		{108, 359},
		{218, 378},
		{939, 258},
		{228, 990},
		{563, 618},
		{513, 859},
		{212, 338},
		{770, 84},
		{952, 213},
		{101, 344},
		{692, 919},
		{698, 230},
		{208, 114},
		{762, 320},
		{16, 537},
		{315, 194},
		{1000, 206},
		{493, 538},
		{150, 577},
		{300, 544},
		{359, 704},
		{735, 274},
		{890, 43},
		{723, 357},
		{232, 244},
		{444, 507},
		{832, 484},
		{502, 975},
		{207, 588},
		{325, 449},
		{621, 90},
		{143, 36},
		{191, 301},
		{613, 982},
		{172, 644},
		{14, 155},
		{899, 60},
		{169, 575},
		{306, 864},
		{819, 391},
		{137, 746},
		{426, 195},
		{885, 66},
		{897, 312},
		{228, 536},
		{762, 791},
		{182, 953},
		{900, 675},
		{795, 298},
		{393, 523},
		{749, 350},
		{391, 915},
		{277, 299},
		{643, 386},
		{244, 627},
		{842, 80},
		{297, 793},
		{400, 104},
		{869, 922},
		{639, 845},
		{394, 146},
		{320, 669},
		{303, 782},
		{559, 416},
		{796, 102},
		{426, 431},
		{934, 132},
		{814, 748},
		{146, 382},
		{707, 805},
		{379, 650},
		{536, 646},
		{121, 183},
		{565, 22},
		{383, 377},
		{461, 846},
		{403, 872},
		{621, 710},
		{127, 884},
		{83, 689},
		{915, 629},
		{124, 408},
		{979, 73},
		{422, 525},
		{507, 989},
		{526, 499},
		{383, 582},
		{735, 434},
		{235, 123},
		{602, 311},
		{537, 487},
		{56, 329},
		{43, 44},
		{782, 31},
		{104, 505},
		{344, 220},
		{799, 185},
		{802, 788},
		{986, 837},
		{151, 931},
		{172, 787},
		{406, 835},
		{701, 712},
		{371, 858},
		{565, 321},
		{421, 774},
		{966, 203},
		{352, 58},
		{34, 251},
		{324, 279},
		{673, 653},
		{383, 636},
		{743, 865},
		{437, 949},
		{159, 223},
		{325, 314},
		{571, 605},
		{105, 548},
		{592, 326},
		{588, 615},
		{220, 468},
		{214, 804},
		{761, 950},
		{242, 958},
		{772, 480},
		{254, 960},
		{756, 665},
		{150, 15},
		{516, 997},
		{151, 470},
		{40, 5},
		{488, 520},
		{810, 848},
		{817, 908},
		{311, 515},
		{378, 826},
		{596, 50},
		{528, 744},
		{213, 929},
		{728, 392},
		{772, 691},
		{162, 861},
		{365, 438},
		{643, 549},
		{774, 423},
		{455, 211},
		{10, 166},
		{160, 363},
		{392, 574},
		{603, 926},
		{766, 193},
		{430, 360},
		{367, 260},
		{201, 959},
		{552, 551},
		{791, 463},
		{520, 460},
		{958, 729},
		{583, 641},
		{552, 351},
		{592, 964},
		{778, 936},
		{683, 241},
		{194, 711},
		{222, 21},
		{496, 130},
		{307, 719},
		{852, 269},
		{689, 554},
		{5, 836},
		{692, 343},
		{665, 362},
		{346, 52},
		{993, 100},
		{49, 219},
		{721, 918},
		{600, 491},
		{733, 695},
		{908, 236},
		{732, 771},
		{537, 585},
		{550, 133},
		{634, 812},
		{989, 697},
		{660, 825},
		{489, 27},
		{149, 902},
		{82, 752},
		{496, 595},
		{776, 547},
		{487, 418},
		{50, 883},
		{444, 693},
		{515, 625},
		{139, 370},
		{890, 658},
		{813, 737},
		{63, 453},
		{363, 847},
		{762, 19},
		{51, 955},
		{948, 205},
		{236, 608},
		{269, 619},
		{258, 652},
		{53, 857},
		{193, 757},
		{284, 540},
		{256, 963},
		{315, 775},
		{215, 295},
		{649, 838},
		{921, 882},
		{929, 450},
		{675, 8},
		{837, 984},
		{847, 612},
		{723, 174},
		{456, 977},
		{964, 808},
		{403, 398},
		{351, 823},
		{710, 726},
		{488, 202},
		{989, 999},
		{306, 239},
		{639, 913},
		{934, 68},
		{318, 464},
		{679, 353},
		{139, 597},
		{426, 240},
		{525, 720},
		{591, 171},
		{505, 233},
		{294, 870},
		{669, 893},
		{649, 368},
		{187, 557},
		{760, 111},
		{208, 830},
		{784, 645},
		{989, 700},
		{392, 937},
		{542, 495},
		{659, 482},
		{426, 672},
		{844, 142},
		{526, 445},
		{918, 447},
		{33, 225},
		{16, 980},
		{369, 373},
		{903, 57},
		{36, 690},
		{620, 389},
		{757, 479},
		{627, 292},
		{317, 322},
		{18, 355},
		{316, 35},
		{375, 192},
		{788, 651},
		{70, 119},
		{18, 773},
		{691, 440},
		{618, 768},
		{422, 103},
		{429, 6},
		{908, 145},
		{36, 340},
		{378, 41},
		{275, 606},
		{870, 410},
		{54, 569},
		{414, 539},
		{855, 276},
		{212, 81},
		{823, 196},
		{32, 992},
		{934, 415},
		{278, 37},
		{796, 217},
		{92, 293},
		{242, 983},
		{651, 564},
		{243, 747},
		{479, 61},
		{746, 543},
		{133, 96},
		{751, 310},
		{822, 11},
		{73, 339},
		{834, 28},
		{948, 759},
		{516, 186},
		{559, 97},
		{85, 366},
		{785, 849},
		{194, 270},
		{819, 286},
		{620, 702},
		{544, 67},
		{102, 573},
		{855, 703},
		{520, 118},
		{703, 411},
		{332, 684},
		{330, 348},
		{393, 26},
		{963, 503},
		{244, 661},
		{564, 265},
		{774, 210},
		{91, 246},
		{182, 694},
		{341, 266},
		{367, 285},
		{70, 873},
		{889, 466},
		{366, 485},
		{966, 483},
		{795, 713},
		{455, 727},
		{741, 234},
		{887, 272},
		{571, 45},
		{90, 141},
		{903, 925},
		{50, 670},
		{695, 308},
		{492, 928},
		{325, 590},
		{607, 709},
		{722, 519},
		{852, 442},
		{653, 587},
		{661, 120},
		{864, 474},
		{204, 944},
		{277, 717},
		{953, 841},
		{867, 654},
		{8, 288},
		{393, 110},
		{381, 942},
		{855, 685},
		{540, 792},
		{776, 498},
		{432, 412},
		{619, 209},
		{286, 454},
		{498, 914},
		{758, 354},
		{336, 77},
		{164, 866},
		{321, 529},
		{56, 623},
		{79, 708},
		{825, 88},
		{261, 875},
		{561, 794},
		{964, 126},
		{646, 724},
		{706, 851},
		{894, 917},
		{541, 555},
		{443, 446},
		{771, 662},
		{502, 901},
		{70, 74},
		{503, 116},
		{675, 420},
		{166, 558},
		{155, 815},
		{97, 29},
		{392, 978},
		{569, 753},
		{161, 409},
		{593, 152},
		{999, 604},
		{587, 827},
		{260, 385},
		{938, 630},
		{525, 985},
		{71, 880},
		{434, 361},
		{902, 876},
		{820, 188},
		{995, 969},
		{15, 200},
		{847, 863},
		{429, 428},
		{555, 940},
		{496, 291},
		{929, 129},
		{204, 190},
		{407, 973},
		{652, 364},
		{426, 809},
		{321, 1},
		{768, 65},
		{960, 871},
		{373, 238},
		{718, 467},
		{386, 435},
		{432, 396},
		{576, 624},
		{859, 508},
		{489, 769},
		{961, 896},
		{75, 881},
		{835, 663},
		{223, 216},
		{554, 904},
		{173, 267},
		{760, 189},
		{807, 745},
		{511, 427},
		{992, 616},
		{571, 237},
		{245, 878},
		{899, 38},
		{868, 64},
		{468, 912},
		{294, 991},
		{77, 546},
		{555, 510},
		{222, 767},
		{756, 682},
		{164, 581},
		{704, 106},
		{533, 78},
		{686, 534},
		{955, 178},
		{518, 512},
		{571, 433},
		{994, 158},
		{939, 800},
		{320, 941},
		{38, 55},
		{584, 943},
		{956, 895},
		{77, 553},
		{62, 304},
		{640, 42},
		{950, 594},
		{183, 165},
		{89, 497},
		{127, 448},
		{698, 305},
		{865, 387},
		{865, 688},
		{459, 504},
		{788, 987},
		{432, 622},
		{796, 441},
		{490, 854},
		{499, 666},
		{99, 436},
		{109, 739},
		{415, 478},
		{930, 967},
		{881, 331},
		{726, 945},
		{677, 452},
		{737, 402},
		{91, 972},
		{565, 628},
		{996, 20},
		{494, 347},
		{940, 30},
		{602, 262},
		{12, 957},
		{761, 112},
		{529, 674},
		{85, 349},
		{495, 2},
		{501, 725},
		{250, 678},
		{611, 263},
		{13, 473},
		{834, 888},
		{65, 696},
		{355, 156},
		{149, 601},
		{955, 877},
		{1000, 229},
		{583, 439},
		{850, 249},
		{16, 522},
		{821, 358},
		{461, 715},
		{770, 289},
		{523, 589},
		{789, 457},
		{626, 556},
		{868, 916},
		{312, 840},
		{680, 567},
		{646, 117},
		{670, 951},
		{30, 179},
		{659, 140},
		{781, 736},
		{793, 657},
		{974, 107},
		{816, 647},
		{184, 147},
		{427, 667},
		{789, 419},
		{577, 500},
		{640, 93},
		{153, 648},
		{347, 853},
		{582, 527},
		{275, 633},
		{24, 94},
		{241, 17},
		{984, 198},
	}

	result := solve(c, tree)

	for _, resultItem := range result {
		fmt.Println(resultItem)
	}
}
