package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n int, bs [2][]byte) []string {
	inb := bs[0]
	outb := bs[1]

	p := make([][]byte, n)
	for i := range p {
		p[i] = make([]byte, n)
		for j := range p[i] {
			p[i][j] = 'N'
		}
	}

	for i := 0; i < n; i++ {
		p[i][i] = 'Y'
		for j := i + 1; j < n; j++ {
			if p[i][j-1] == 'Y' && inb[j] == 'Y' && outb[j-1] == 'Y' {
				p[i][j] = 'Y'
			}
		}
		for j := i - 1; j >= 0; j-- {
			if p[i][j+1] == 'Y' && inb[j] == 'Y' && outb[j+1] == 'Y' {
				p[i][j] = 'Y'
			}
		}
	}
	out := make([]string, n)
	for i := range out {
		out[i] = string(p[i])
	}
	return out
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(r, &t)
	for i := 1; i < t+1; i++ {
		var n int
		fmt.Fscan(r, &n)
		var bs [2][]byte
		for j := 0; j < 2; j++ {
			bs[j] = make([]byte, n)
			var s string
			fmt.Fscan(r, &s)
			for k := 0; k < n; k++ {
				bs[j][k] = s[k]
			}
		}
		fmt.Printf("Case #%v:\n", i)
		for _, v := range solution(n, bs) {
			fmt.Println(v)
		}
	}
}
