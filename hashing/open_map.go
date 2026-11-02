package hashing

// [medium] OpenMap is a string->int map using open addressing with linear probing.
//
// All entries live in one flat slice; collisions walk to the next free slot.
// Steps:
//  1. Construction: allocate one flat array of slots, each holding a
//     key, a value, and an occupied flag — all empty at first. Make
//     the slot count a power of two, since slots are addressed by
//     keeping the low bits of the key's hash.
//  2. Putting: start at the slot the key hashes to. Walk forward one
//     slot at a time until you reach a slot that either holds the same
//     key (overwrite its value in place; nothing else changes) or is
//     empty (write the entry there and count one more entry). The walk
//     always succeeds: the load-factor rule below guarantees there is
//     always an empty slot to find.
//  3. Getting: start at the same slot and walk forward, comparing
//     keys. Finding the key ends the search with its value. Reaching
//     an empty slot ends the search with "absent" — an occupied slot
//     with a different key just means the walk continues.
//  4. Growing: after an insert, if entries have risen to about seven
//     tenths of the slot count, move to a doubled slot array and
//     reinsert every live entry at its position in the new array.
//     Linear probing degrades sharply beyond that point because
//     occupied slots clump into long runs; the headroom keeps the
//     runs short.
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
