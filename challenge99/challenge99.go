package challenge99

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const File = "0099_base_exp.txt"

func loadFile() (all [][2]int) {
	f, err := os.Open(File)
	if err != nil {
		log.Fatalf("failed to open file: %s", File)
	}
	b := bufio.NewScanner(f)
	b.Split(bufio.ScanLines)
	for b.Scan() {
		var ns [2]int
		fmt.Sscanf(b.Text(), "%d,%d", &ns[0], &ns[1])
		all = append(all, ns)
	}
	return all
}

func Challenge99() {
	var winner int
	var maxFloat float64
	numbers := loadFile()
	for index, line := range numbers {
		lnm := math.Log10(float64(line[0])) * float64(line[1])
		if lnm > maxFloat {
			maxFloat = lnm
			winner = index + 1
		}
	}
	fmt.Printf("challenge 99 solution is: %d\n", winner)
}
