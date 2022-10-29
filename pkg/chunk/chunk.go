package chunk

import (
	"fmt"
	"io"

	"github.com/x0f5c3/gosync/pkg/compress"
)

type NewChunk struct {
	Buf           []byte
	Offset        int
	Hi            int
	Pos           int
	CompressLevel int
	Compressed    bool
	ToCompress    bool
}

func (n *NewChunk) Read(reader io.ReaderAt) error {
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

func (n *NewChunk) Compress() {
	if n.Compressed {
		return
	}
	n.Buf = compress.CompressWithOption(n.Buf, n.CompressLevel)

}
