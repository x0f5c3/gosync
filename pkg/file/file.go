package file

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"github.com/cespare/xxhash/v2"
	"github.com/kalafut/imohash"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/md4"
)

type Algo uint

const (
	MD4        Algo = 1 + iota // import golang.org/x/crypto/md4
	MD5                        // import crypto/md5
	SHA1                       // import crypto/sha1
	SHA224                     // import crypto/sha256
	SHA256                     // import crypto/sha256
	SHA384                     // import crypto/sha512
	SHA512                     // import crypto/sha512
	Sha512224                  // import crypto/sha512
	Sha512256                  // import crypto/sha512
	Blake2b256                 // import golang.org/x/crypto/blake2b
	Blake2b384                 // import golang.org/x/crypto/blake2b
	Blake2b512                 // import golang.org/x/crypto/blake2b
	XXHASH                     // import github.com/cespare/xxhash/v2
	IMOHASH                    // import github.com/kalafut/imohash
)

func (a Algo) Sum(data []byte) string {
	switch a {
	case MD4:
		h := md4.New()
		_, err := h.Write(data)
		if err != nil {
			return ""
		}
		return fmt.Sprintf("%x", h.Sum(nil))
	case MD5:
		return fmt.Sprintf("%x", md5.Sum(data))
	case SHA1:
		return fmt.Sprintf("%x", sha1.Sum(data))
	case SHA224:
		return fmt.Sprintf("%x", sha256.Sum224(data))
	case SHA256:
		return fmt.Sprintf("%x", sha256.Sum256(data))
	case SHA384:
		return fmt.Sprintf("%x", sha512.Sum384(data))
	case SHA512:
		return fmt.Sprintf("%x", sha512.Sum512(data))
	case Sha512224:
		return fmt.Sprintf("%x", sha512.Sum512_224(data))
	case Sha512256:
		return fmt.Sprintf("%x", sha512.Sum512_256(data))
	case Blake2b256:
		return fmt.Sprintf("%x", blake2b.Sum256(data))
	case Blake2b384:
		return fmt.Sprintf("%x", blake2b.Sum384(data))
	case Blake2b512:
		return fmt.Sprintf("%x", blake2b.Sum512(data))
	case XXHASH:
		return fmt.Sprintf("%x", xxhash.Sum64(data))
	case IMOHASH:
		return fmt.Sprintf("%x", imohash.Sum(data))
	default:
		return ""
	}
}

type CheckSum struct {
	Ref  string
	Algo Algo
}

func NewSum(algo Algo, ref ...string) *CheckSum {
	refVar := func() string {
		if len(ref) > 0 {
			return ref[0]
		}
		return ""
	}()
	return &CheckSum{
		Ref:  refVar,
		Algo: algo,
	}
}

func (c *CheckSum) Check(data []byte) bool {
	return c.Ref == c.Algo.Sum(data)
}

type File struct {
	Src string
	Dst string
	Sum *CheckSum
}

func NewFile(src, dst string, algo Algo, ref ...string) *File {
	return &File{
		Src: src,
		Dst: dst,
		Sum: NewSum(algo, ref...),
	}
}
