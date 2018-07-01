package main

type Bit bool

const (
	Zero Bit = false
	One  Bit = true
)

type Bitstream struct {
	Bits     []uint64
	BitCount int64
}

func (b *Bitstream) info() (int64, int64) {
	arrPos := int64(b.BitCount % 64)
	arrSize := int64(b.BitCount / 64)

	return arrPos, arrSize
}

func (b *Bitstream) Enlarge() {
	b.Bits = append(b.Bits, make([]uint64, len(b.Bits)+1)...)
}

func (b *Bitstream) Append(a Bit) {
	pos, size := b.info()
	if len(b.Bits) == int(size) {
		b.Enlarge()
	}

	if a {
		b.Bits[size] |= 1 << uint(pos)
	}

	b.BitCount++
}

func (b *Bitstream) Appends(a Bitstream) {
	for i := b.BitCount; i < a.BitCount+b.BitCount; i++ {
		b.Append(a.Pop())
	}
}

func (b *Bitstream) Pop() Bit {
	pos, size := b.info()
	if size != 0 || pos != 0 {
		if pos == 0 {
			if (1<<63)&b.Bits[size-1] == 0 {
				return Zero
			}

		} else {
			if len(b.Bits) == 0 || (1<<uint64(pos-1))&b.Bits[size] == 0 {
				return Zero
			}
		}
		b.BitCount--

		return One
	}

	return Zero
}
