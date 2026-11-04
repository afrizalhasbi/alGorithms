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
// Definitions:
//  - slot: one element of the flat array: key, value, occupied flag.
//     There is no bucket-of-chain; occupied or not, the slot is the
//     whole story.
//  - probe: the walk from the home slot forward one slot at a time,
//     wrapping around the end of the array if reached. Both Put and
//     Get probe the same way.
//  - home slot: the position the key hashes to: FNV1a(key) & (n-1).
//     The entry may end up far from it.
//  - clump: a maximal run of occupied slots. A probe entering a clump
//     must traverse it; clumps merge as the table fills, which is why
//     the grow threshold is far below full.
//  - load factor: entries divided by slots. Grows at about 0.7.
//
// Steps:
//  1. Construction: allocate a flat slot array, each slot holding a
//     key, a value, and an occupied flag, all empty. Length is a
//     power of two, since slots are addressed by the low bits of the
//     hash: hash & (n-1) turns the raw hash into a valid index in
//     0..n-1.
//  2. Put: start at the slot the key hashes to and walk forward one
//     slot at a time. A slot holding the same key ends the walk with
//     its value replaced. An empty slot ends the walk with the entry
//     written there and the entry count raised by one.
//  3. Get: the same walk, still exactly one hash. Compare full keys
//     with == at each occupied slot (a different key may occupy the
//     slot this key hashed to); an empty slot ends the walk with
//     "absent".
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
