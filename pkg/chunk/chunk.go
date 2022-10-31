package chunk

import (
	"fmt"
	"io"

	"github.com/x0f5c3/gosync/pkg/compress"
)

type Chunks struct {
	chnks                                          []*Chunk
	length, chunkSize, low, currOffset, currLength int
	next                                           *Chunk
	compressed                                     bool
	ToCompress                                     bool
	compressLevel                                  int
}

func NewChunks(length int, chunkSize int, compressed bool, toCompress bool, compressLevel ...int) *Chunks {
	cl := func() int {
		if len(compressLevel) > 0 {
			return compressLevel[0]
		}
		return -2
	}()
	return &Chunks{length: length, chunkSize: chunkSize, compressed: compressed, ToCompress: toCompress, compressLevel: cl}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (c *Chunks) Next() bool {
	if c.low > c.length {
		c.next = nil
		return false
	} else {
		prevLow := c.low
		c.low += min(c.chunkSize, c.length-c.low+1)
	}
}

func (c *Chunks) Collect() *Chunks {
	var res []*Chunk
}

type Chunk struct {
	Buf           []byte
	Offset        int
	Hi            int
	Pos           int
	CompressLevel int
	Compressed    bool
	ToCompress    bool
}

func (n *Chunk) Read(reader io.ReaderAt) error {
	bufLen := n.Hi - n.Offset + 1
	buf := make([]byte, bufLen)
	cnt, err := reader.ReadAt(buf, int64(n.Offset))
	if err != nil {
		return err
	}
	if cnt != bufLen {
		return fmt.Errorf("read %v bytes, but expected %v", cnt, bufLen)
	}
	if n.ToCompress {
		n.Buf = compress.CompressWithOption(buf, n.CompressLevel)
	} else {
		n.Buf = buf
	}
	return nil
}

func (n *Chunk) Compress() {
	if n.Compressed {
		return
	}
	n.Buf = compress.CompressWithOption(n.Buf, n.CompressLevel)

}
