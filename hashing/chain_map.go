package hashing

// [medium] ChainMap is a string->int map built with separate chaining.
//
// Bucket array of linked buckets; collisions append to the bucket's chain.
// Steps:
//  1. NewChainMap(sizeHint) allocates a bucket slice; bucket count is a power
//     of two >= sizeHint. Use FNV1a(key) & (len(buckets)-1) to index.
//  2. Put(key, val): hash to a bucket, walk its chain; update the entry if
//     the key is there, otherwise append a new entry.
//  3. Get(key): hash to a bucket, walk the chain, compare keys with ==.
//  4. Track the number of entries. When load factor (entries/buckets) exceeds
//     ~1.0, rehash: allocate double the buckets and reinsert everything.
//
// Notes:
//   - With the load factor capped, chain lengths stay O(1) on average —
//     that is the entire reason the map is fast. Skipping the rehash makes
//     it degrade to O(n) per op, and the complexity test will catch it.
//   - Chains make deletion easy; that is the trade you are exploring here.
//
// Expected: O(1) average per op (amortized with rehashing)
// Input: not mutated
type ChainMap struct {
	// TODO: fields
}

func NewChainMap(sizeHint int) *ChainMap {
	// TODO: implement
	return nil
}

func (m *ChainMap) Put(key string, val int) {
	// TODO: implement
}

func (m *ChainMap) Get(key string) (int, bool) {
	// TODO: implement
	return 0, false
}
