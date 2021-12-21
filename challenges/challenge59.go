package challenges

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func factory() func() (bool, [3]int) {
	var count int
	return func() (bool, [3]int) {
		if count >= 26*26*26 {
			return true, [3]int{}
		}
		zero := count % 26
		one := (count / 26) % 26
		two := count / (26 * 26)
		count++
		return false, [3]int{'a' + zero, 'a' + one, 'a' + two}
	}
}
func Challenge59() {

	f, err := os.Open("p059_cipher.txt")
	if err != nil {
		fmt.Printf("%d\n", err)
		os.Exit(-1)
	}
	defer f.Close()

	var keyTriplet [3]int
	var finished bool
	counter := factory()
	for {
		finished, keyTriplet = counter()
		if finished {
			break
		}
		f.Seek(0, 0)
		r := bufio.NewReader(f)
		var encNum int
		var unencBuffer [1455]int
		for count := 0; true; count++ {
			word, err := r.ReadString(',')
			word = strings.TrimRight(word, ",")
			fmt.Sscanf(word, "%d", &encNum)
			unencBuffer[count] = encNum ^ keyTriplet[count%3]
			if err == io.EOF {
				break
			}
		}
		containsAnd := false
		containsThe := false
		containsFor := false
		for k, v := range unencBuffer {
			if v == 'a' && unencBuffer[k+1] == 'n' && unencBuffer[k+2] == 'd' {
				containsAnd = true
			}
			if v == 'f' && unencBuffer[k+1] == 'o' && unencBuffer[k+2] == 'r' {
				containsFor = true
			}
			if v == 't' && unencBuffer[k+1] == 'h' && unencBuffer[k+2] == 'e' {
				containsThe = true
			}
		}
		if containsThe && containsAnd && containsFor {
			fmt.Printf("Contains the words 'and', 'for' and 'the' so may be onto a winner with the password: '%c%c%c'\n\n", keyTriplet[0], keyTriplet[1], keyTriplet[2])
			for _, v := range unencBuffer {
				fmt.Printf("%c", v)
			}
		}
	}
}
