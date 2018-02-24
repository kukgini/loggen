package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	ExampleWriteTo()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var prefixRunes = []string{"OPP#", "ERR#"}
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	prefix := prefixRunes[rand.Intn(2)]
	return prefix + string(b)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func ExampleWriteTo() {
	var lines []string
	for i := 0; i < 500; i++ {
		s := RandStringRunes(10)
		lines = append(lines, s)
	}
	f, err := os.OpenFile("INPUT/INPUT.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		//return err
	}
	defer f.Close()
	if err := WriteTo(f, lines); err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(f, "Q\n")
}
func WriteTo(w io.Writer, lines []string) error {
	for _, line := range lines {
		if _, err := fmt.Fprintf(w, "%s\n", line); err != nil {
			return err
		}
	}
	return nil
}
