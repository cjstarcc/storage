package consistent

import "stathat.com/c/consistent"

type Hash struct {
	consistent *consistent.Consistent
}

func (h *Hash) Add(key string) {
	h.consistent.Add(key)
}

func (h *Hash) Get(key string) (string, error) {
	return h.consistent.Get(key)
}

func (h *Hash) Remove(key string) {
	h.consistent.Remove(key)
}

func NewHash() *Hash {
	consistent := consistent.New()
	h := &Hash{consistent: consistent}
	return h
}
