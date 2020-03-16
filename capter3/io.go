package capter3

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	var s string
	if _, err := fmt.Fscanf(f, "%s\n", &s); err == nil {
		fmt.Println(s)
	}

	return nil
}

func CreateFile(fileName string) error {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := fmt.Fprintf(f, "%s\n", "CreateFile"); err != nil {
		return err
	}

	return nil
}

func WriteTo(w io.Writer, lines []string) error {
	for _, line := range lines {
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}

func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
