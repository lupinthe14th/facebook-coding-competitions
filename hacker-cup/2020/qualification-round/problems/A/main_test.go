package main

import (
	"os"
	"path/filepath"
)

func Example_main() {
	fd, _ := os.Open(filepath.Join("testdata", "travel_restrictions_sample_input.txt"))
	orgStdin := os.Stdin
	os.Stdin = fd
	defer func() {
		os.Stdin = orgStdin
		fd.Close()
	}()

	main()

	// Output:
	// Case #1:
	// YY
	// YY
	// Case #2:
	// YY
	// NY
	// Case #3:
	// YN
	// NY
	// Case #4:
	// YNNNN
	// YYNNN
	// NNYYN
	// NNNYN
	// NNNYY
	// Case #5:
	// YYYNNNNNNN
	// NYYNNNNNNN
	// NNYNNNNNNN
	// NNYYNNNNNN
	// NNYYYNNNNN
	// NNNNNYNNNN
	// NNNNNNYYYN
	// NNNNNNYYYN
	// NNNNNNNNYN
	// NNNNNNNNYY

}
