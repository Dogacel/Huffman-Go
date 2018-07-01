package main

import (
	"sort"
)

// Tree struct
type Tree struct {
	val    rune
	weight int
	left   *Tree
	right  *Tree
}

// BuildCodebook from Hoffman tree
func BuildCodebook(hoffman *Tree, stream Bitstream, book map[rune]Bitstream) {
	// Check if we have arrived at a rune.
	if hoffman.val != 0xFFFF {
		book[hoffman.val] = stream
		return
	}

	// Traverse the left subtree if exists
	if hoffman.left != nil {
		stream.Append(Zero)
		BuildCodebook(hoffman.left, stream, book)
		stream.Pop()
	}

	// Traverse the right subtree if exists.
	if hoffman.right != nil {
		stream.Append(One)
		BuildCodebook(hoffman.right, stream, book)
		stream.Pop()
	}
}

// BuildHoffmanTree from given string
func BuildHoffmanTree(msg string) (Tree, map[rune]Bitstream) {
	runeMap := make(map[rune]int)

	// Count the rune frequency
	for _, currentRune := range msg {
		runeMap[currentRune]++
	}

	// Sort runes alphabetically for consistency
	keys := make([]int, 0)
	for k := range runeMap {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	// Convert each rune to an empty tree in a slice.
	treeList := make([]Tree, 0)
	for _, k := range keys {
		treeList = append(treeList, Tree{rune(k), runeMap[rune(k)], nil, nil})
	}

	treeSize := len(treeList)

	// Iterate over treeList until all trees are connected.
	for treeSize > 1 {

		pos1 := 0
		pos2 := 1

		// Find the positions of the two least occuruing trees
		for i := 1; i < treeSize; i++ {
			if treeList[i].weight <= treeList[pos1].weight {
				pos2 = pos1
				pos1 = i
			} else if treeList[i].weight < treeList[pos2].weight {
				pos2 = i
			}
		}

		// Copy the trees.
		tr1 := treeList[pos1]
		tr2 := treeList[pos2]
		// Connect two trees in an upper tree.
		upperTree := Tree{0xFFFF, tr1.weight + tr2.weight, &tr1, &tr2} // 0xFFFF is not a valid UTF-8 character

		// Remove old trees from the list and add the new upperTree to the list carefully.
		if pos2 == treeSize-1 {
			treeList[pos1] = upperTree
			treeSize--
		} else if pos1 == treeSize-1 {
			treeList[pos2] = upperTree
			treeSize--
		} else {
			treeList[pos1] = upperTree
			treeList[pos2] = treeList[treeSize-1]
			treeSize--
		}

	}

	// Build the codebook for decoding
	codebook := make(map[rune]Bitstream)
	BuildCodebook(&treeList[0], Bitstream{}, codebook)

	return treeList[0], codebook

}

// HuffmanEncode given string
func HuffmanEncode(in string) (out Bitstream) {
	_, d := BuildHoffmanTree(in)
	for _, val := range in {
		// Concatenate the bitstrings of all runes one by one.
		out.Appends(d[val])
	}

	return out
}

// HuffmanDecode given string from the given huffman tree
func HuffmanDecode(in Bitstream, huffman Tree) (out string) {

	current := huffman

	for in.BitCount > 0 {

		// If we hit a rune, append it to the out string and return back to the root node.
		if current.val != 0xFFFF {
			out += string(current.val)
			current = huffman
		}

		val := in.Pop()

		if val == Zero {
			// Go to the left subtree if next rune is 0
			current = *current.left
		} else if val == One {
			// Go to the right subtree if next rune is 1
			current = *current.right
		}

	}

	// Decode the latest rune.
	if current.val != 0xFFFF {
		out += string(current.val)
		current = huffman
	}

	return out
}
