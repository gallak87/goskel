package countbox

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'gridSearch' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING_ARRAY G
 *  2. STRING_ARRAY P
 */
/*
func gridSearch(G []string, P []string, print bool) string {
    // Write your code here

    pRow, pCol, firstMatchGCol := 0, 0, 0
    gHeight, gWidth := len(G), len(G[0])
    pHeight, pWidth := len(P), len(P[0])
    isMatching := false

    fmt.Printf("////// NEW TEST G[%d,%d] P[%d,%d]\n", gWidth, gHeight, pWidth, pHeight)

    // use pointers to walk row/col for both G and P
    for gRow := 0; gRow < gHeight; gRow++ {
        yStr := G[gRow]
        pStr := P[pRow]
        // break early
        if !isMatching && strings.Index(yStr, pStr) < 0 {
            continue
        }
        for gCol := 0; gCol < gWidth; gCol++ {
            if isMatching && gCol < firstMatchGCol {
                // while matching, skip any x chars up front
                gCol = firstMatchGCol-1
                continue
            }
            if isMatching && pCol == pWidth {
                break
            }
            // stay within pWidth bounds
            if (pCol < pWidth) {
                if (print) {
                    fmt.Printf("[%c | %c]\n", yStr[gCol], pStr[pCol])
                    fmt.Printf("G (%d,%d) ---- P (%d,%d)\n", gRow, gCol, pRow, pCol)
                }
                if yStr[gCol] == pStr[pCol] {
                    // x column match, if we are within pWidth we are matching
                    if !isMatching && pCol == 0 {
                        // one-time store last x index from G array at point of first match
                        firstMatchGCol = gCol
                    }
                    isMatching = true
                    pCol++
                } else {
                    // x column mismatch, since we didn't reach pWidth we stopped matching
                    pRow, pCol, firstMatchGCol, isMatching = 0, 0, 0, false
                }
            }
        }
        // if we're matching and x column reached width, reset counters
        if isMatching && pCol == pWidth {
            pRow++
            pCol = 0
            if (print) {
                fmt.Printf("-- new row --\n")
                fmt.Printf("isMatching %s\n", strconv.FormatBool(isMatching))
                fmt.Printf("pRow %d pHeight %d\n", pRow, pHeight)
            }
            if pRow == pHeight {
                return "YES"
            }
        } else {
            // edge case, if later row is not matching x column, but exists
            // need to reset with gCol++


            // reset both query pointers and first match pointer
            pRow, pCol, firstMatchGCol, isMatching = 0, 0, 0, false
        }
    }

    return "NO"
}
*/

func checkRow(gSlice string, pStr string) bool {
	g, _ := strconv.ParseInt(gSlice, 10, 64)
	p, _ := strconv.ParseInt(pStr, 10, 64)
	return g == p
}

func checkWindow(G []string, P []string, gRowStart, gColStart int) bool {
	pRow := 0
	// gHeight, gWidth := len(G), len(G[0])
	pHeight, pWidth := len(P), len(P[0])
	searchHeight := gRowStart + pHeight
	for gRow := gRowStart; gRow < searchHeight; gRow++ {
		gStr := G[gRow]
		pStr := P[pRow]
		gSlice := gStr[gColStart : gColStart+pWidth]
		result := checkRow(gSlice, pStr)
		fmt.Printf("%s -> %s\n", gSlice, pStr)
		if !result {
			return false
		}

		pRow++
		// fmt.Printf("pRow %d pHeight %d pWidth %d\n", pRow, pHeight, pWidth)
		// fmt.Printf("gRow %d gHeight %d gWidth %d\n", gRow, gHeight, gWidth)
	}
	return true
}

func gridSearch(G []string, P []string, print bool) string {
	pRow := 0
	gHeight, gWidth := len(G), len(G[0])
	pHeight, pWidth := len(P), len(P[0])

	fmt.Printf("////// NEW TEST G[%d,%d] P[%d,%d]\n", gWidth, gHeight, pWidth, pHeight)

	// Write your code here
	for gRow := 0; gRow < gHeight; gRow++ {
		yStr := G[gRow]
		pStr := P[pRow]
		xStart := strings.Index(yStr, pStr)
		// break early
		if xStart < 0 {
			continue
		}
		// found in this row, check window
		if checkWindow(G, P, gRow, xStart) {
			return "YES"
		}
	}

	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		RTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		R := int32(RTemp)

		var G []string

		for i := 0; i < int(R); i++ {
			GItem := readLine(reader)
			G = append(G, GItem)
		}

		secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		rTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
		checkError(err)
		r := int32(rTemp)

		var P []string

		for i := 0; i < int(r); i++ {
			PItem := readLine(reader)
			P = append(P, PItem)
		}

		result := gridSearch(G, P, tItr > 0)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
