package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(t, n int, b []byte) string {
	var helper func(bb []byte) []byte
	helper = func(bb []byte) []byte {
		var memo [2]int
		for i := range bb {
			memo[bb[i]-'A']++
		}
		if memo[0] == 2 {
			return []byte{'A'}
		}
		if memo[1] == 2 {
			return []byte{'B'}
		}
		return bb
	}
	i := 3
	for len(b) != 1 {
		bb := make([]byte, 3)
		copy(bb, b[i-3:i])
		tmp := helper(bb)
		if len(tmp) == 1 {
			b = append(b[:i-3], append(tmp, b[i:]...)...)
			n = len(b)
			i = 3
			continue
		}
		if i >= n {
			break
		}
		i++
	}
	if len(b) != 1 {
		return fmt.Sprintf("Case #%v: N", t)
	}
	return fmt.Sprintf("Case #%v: Y", t)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(r, &t)
	for i := 1; i <= t; i++ {
		var n int
		fmt.Fscan(r, &n)
		var s string
		fmt.Fscan(r, &s)
		b := make([]byte, n)
		for j := range s {
			b[j] = s[j]
		}
		fmt.Println(solution(i, n, b))
	}
}
