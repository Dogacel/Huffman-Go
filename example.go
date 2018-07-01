package main

import (
	"bufio"
	"fmt"
	"os"
)

func printTree(t *Tree) {
	fmt.Printf("val: %c || %d", t.val, t.weight)
	if t.left != nil {
		fmt.Printf("\n Left: ")
		printTree(t.left)
		fmt.Println("END LEFT")
	}
	if t.right != nil {
		fmt.Printf("\n Right: ")
		printTree(t.right)
		fmt.Println("END RIGHT")
	}
}

func dohuf(str string) {
	out := HuffmanEncode(str)

	tree, d := BuildHoffmanTree(str)

	dout := HuffmanDecode(out, tree)

	fmt.Println("================")
	fmt.Println("INPUT: ")
	fmt.Println(str)
	fmt.Println("================")
	fmt.Println("ENCODED OUTPUT: ")
	fmt.Println(out)
	fmt.Println("================")
	fmt.Println("DECODED OUTPUT: ")
	fmt.Println(dout)
	fmt.Println("================")

	if str != dout {
		panic("ASSERTION FAILED, INPUT OUTPUT DOES NOT MATCH")
	}

	dstr := ""
	for i := range d {
		dstr += string(i) + ". "
	}

	firstSize := len(str) * 32
	endSize := len(dstr)*32 + len(out) - 32
	rate := float64(firstSize) / float64(endSize)

	fmt.Println("First size: ", firstSize)
	fmt.Println("Final size: ", endSize)
	fmt.Printf("Compression rate: %.2f%%\n", rate*100.0)
}

func main() {
	str := " "
	inputReader := bufio.NewReader(os.Stdin)
	for str != "" {
		str, _ = inputReader.ReadString('\n')
		dohuf(str)
	}
}
