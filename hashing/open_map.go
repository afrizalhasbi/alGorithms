package hashing

// [medium] OpenMap is a string->int map using open addressing with linear probing.
//
// Background. Open addressing is the other main collision policy
// (the first is chaining, see chain_map). There are no chains and no
// extra objects: every entry lives directly in one flat array of
// slots. When a key hashes to a slot that is already taken, the
// insert does not give up, it probes: it checks the next slot, then
// the next, until a free slot is found. The entry goes there, which
// may be far from where it hashed. A lookup repeats the same walk,
// comparing keys at each occupied slot, until it finds the key or
// reaches an empty slot.
//
// Reaching an empty slot is what ends a lookup with "absent". That
// makes the empty slot a load-bearing part of the structure: it
// stops the probe. It also explains the grow threshold. With
// chaining, a load factor of one just means chains of length one.
// Here, a full-ish array means an occupied slot clump can form, and
// any key hashing into the clump probes its whole length; clumps
// merge and grow as the table fills. Above about seven tenths full,
// probe lengths blow up. Below that, they stay short.
//
// Resizing works like in chain_map: when the load factor crosses the
// threshold, allocate a double-length slot array and insert every
// live entry again. Entries land in different slots than before,
// since the slot count changed.
//
// Steps:
//  1. Construction: allocate a flat slot array, each slot holding a
//     key, a value, and an occupied flag, all empty. Length is a
//     power of two, since slots are addressed by the low bits of the
//     hash.
//  2. Put: start at the slot the key hashes to and walk forward one
//     slot at a time. A slot holding the same key ends the walk with
//     its value replaced. An empty slot ends the walk with the entry
//     written there and the entry count raised by one.
//  3. Get: the same walk. A slot holding the key ends it with the
//     value; an empty slot ends it with "absent".
//  4. After an insert, if entries have reached about seven tenths of
//     the slot count, double the slot array and re-insert every live
//     entry.
//
// Notes:
//   - Deletion is the hard part of open addressing: clearing a slot
//     to empty would cut probe walks short and lose entries behind
//     it. The standard fixes (tombstones, backward shift) are out of
//     scope; the tests never delete.
//   - Compare the resize threshold (about 0.7) with chain_map's
//     (about 1.0) and the clumping explanation above.
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
