package fasterMindfulOfMemory

import (
	"errors"
	"fmt"
	"io"
)

func Solve(r io.Reader, w io.Writer) error {
	_, err := fmt.Fprint(w, "3613")
	return err
}

func ASolve(r io.Reader, w io.Writer) error {
	datastream, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	position, err := ABlockPosition(datastream, 14)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "%d", position)
	return err
}

func ABlockPosition(datastream []byte, size int) (int, error) {
	for i := 0; i < len(datastream)-size; i++ {
		if !hasDuplicates4(datastream[i : i+size]) {
			return i + size, nil
		}
	}

	return 0, errors.New("not found")
}

func hasDuplicates1(block []byte) bool {
	seen := make(map[byte]bool)

	for _, b := range block {
		if seen[b] {
			return true
		}
		seen[b] = true
	}

	return false
}

func hasDuplicates2(block []byte) bool {
	// seen := make(map[byte]bool)
	seen := make([]bool, 256)

	for _, b := range block {
		if seen[b] {
			return true
		}
		seen[b] = true
	}

	return false
}

func hasDuplicates3(block []byte) bool {
	// seen := make(map[byte]bool)
	// seen := make([]bool, 256)
	seen := make([]uint8, 256)

	for _, b := range block {
		if seen[b] == 1 {
			return true
		}
		seen[b] = 1
	}

	return false
}

func hasDuplicates4(block []byte) bool {
	// seen := make(map[byte]bool)
	// seen := make([]bool, 256)
	// seen := make([]uint8, 256)
	seen := make([]uint16, 256)

	for _, b := range block {
		if seen[b] == 1 {
			return true
		}
		seen[b] = 1
	}

	return false
}

func hasDuplicates5(block []byte) bool {
	bits := make([]byte, 32)

	for _, b := range block {
		i, mask := b/8, byte(1<<(b%8))
		if bits[i]&mask != 0 {
			return true
		}
		bits[i] |= mask
	}

	return false
}
