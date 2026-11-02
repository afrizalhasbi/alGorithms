package hashing

// [medium] OpenMap is a string->int map using open addressing with linear probing.
//
// All entries live in one flat slice; collisions walk to the next free slot.
// Steps:
//  1. NewOpenMap(sizeHint) allocates a slice of {key, val, occupied} slots,
//     power-of-two length. Index with FNV1a(key) & (len-1).
//  2. Put(key, val): probe forward from the hash index until an empty (or
//     deleted, if you implement deletion) slot; write there. If the key is
//     already present, update in place.
//  3. Get(key): probe from the hash index, comparing keys, until you find it
//     or hit an empty slot (absent).
//  4. Track entries; when load factor exceeds ~0.7, grow: double the slice
//     and reinsert every live entry. Linear probing degrades badly above 0.7
//     because probes form clusters.
//
// Notes:
//   - Tombstones vs backward-shift deletion is the classic rabbit hole;
//     for this exercise you may omit deletion entirely.
//   - Compare the resize threshold (0.7) with ChainMap's (~1.0) and think
//     about why probing needs more headroom.
//
// Expected: O(1) average per op (amortized with resizing)
// Input: not mutated
type OpenMap struct {
	// TODO: fields
}

func NewOpenMap(sizeHint int) *OpenMap {
	// TODO: implement
	return nil
}

func (m *OpenMap) Put(key string, val int) {
	// TODO: implement
}

func (m *OpenMap) Get(key string) (int, bool) {
	// TODO: implement
	return 0, false
}
