package generators

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBytesRandom_Generate(t *testing.T) {
	r := NewRandomBytes(0, 3)
	res, err := r.Generate(nil)
	require.NoError(t, err)
	require.Len(t, res, 3)
	require.Equal(t, []byte{1, 148, 253}, res)
}

func TestRandomBytes_GenerateConcurrent(t *testing.T) {
	r := NewRandomBytes(0, 16)
	var wg sync.WaitGroup
	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				if _, err := r.Generate(nil); err != nil {
					t.Errorf("Generate returned error: %v", err)
					return
				}
			}
		}()
	}
	wg.Wait()
}

func TestInt64Random_GenerateConcurrent(t *testing.T) {
	r, err := NewInt64Random(0)
	require.NoError(t, err)
	var wg sync.WaitGroup
	for i := 0; i < 64; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				if _, err := r.Generate(nil); err != nil {
					t.Errorf("Generate returned error: %v", err)
					return
				}
			}
		}()
	}
	wg.Wait()
}
