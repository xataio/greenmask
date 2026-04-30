package generators

import (
	"encoding/binary"
	"math/rand"
	"sync"
)

type RandomBytes struct {
	mu    sync.Mutex
	r     *rand.Rand
	size  int
	iters int
}

func NewRandomBytes(seed int64, size int) *RandomBytes {
	iters := size / 8
	if size%8 > 0 {
		iters += 1
	}
	return &RandomBytes{
		r:     rand.New(rand.NewSource(seed)),
		size:  size,
		iters: iters,
	}
}

func (br *RandomBytes) Generate(data []byte) ([]byte, error) {
	res := make([]byte, 0, br.size)
	buf := make([]byte, 8)
	br.mu.Lock()
	defer br.mu.Unlock()
	for i := 0; i < br.iters; i++ {
		binary.LittleEndian.PutUint64(buf, br.r.Uint64())
		res = append(res, buf...)
	}
	return res[:br.size], nil
}

func (br *RandomBytes) Size() int {
	return br.size
}
