package main

import (
	"os"
	"path/filepath"
)

func Example_main() {
	fd, _ := os.Open(filepath.Join("testdata", "alchemy_sample_input.txt"))
	orgStdin := os.Stdin
	os.Stdin = fd
	defer func() {
		os.Stdin = orgStdin
		fd.Close()
	}()

	main()

	// Output:
	// Case #1: Y
	// Case #2: N
	// Case #3: Y
	// Case #4: N
	// Case #5: Y
	// Case #6: N
}
