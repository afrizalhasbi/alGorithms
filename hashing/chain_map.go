package hashing

// [medium] ChainMap is a string->int map built with separate chaining.
//
// Background. A hash table stores entries in an array of buckets and
// uses the hash of a key to pick the bucket. Two keys can hash to
// the same bucket (a collision), and any table needs a policy for
// it. Separate chaining is the simplest one: each bucket holds a
// list (the chain) of all entries hashed to it. A lookup hashes,
// finds the bucket, and walks the chain comparing keys. All the
// real work of the table happens inside those chains.
//
// The length of a chain decides how fast a lookup is, and it is
// controlled by the load factor: the number of entries divided by
// the number of buckets. With the load factor held around one, the
// average chain holds about one entry and every operation is
// constant time. If entries are added without growing the table, the
// load factor climbs, chains turn into linked lists, and every
// operation degrades toward scanning them.
//
// The fix is rehashing: when the load factor crosses the threshold,
// allocate a bucket array of double the length and re-insert every
// entry. Each entry lands in a different bucket than before (the
// bucket count changed, so the low bits of the hash pick a different
// bucket), which is why everything must be walked again. Doubling
// keeps the total cost of all rehashes linear in the number of
// inserts, so the amortized cost per operation stays constant.
//
// Structure. There are two levels: a bucket array (one fixed-length
// slice, whose slots are the buckets), and inside each slot a chain
// (a slice of entries, each entry a key/value pair). Chains grow by
// append; the bucket array is only replaced wholesale on rehash. A
// bucket index never comes from the hash directly: the raw hash is a
// huge number, and it is reduced to a valid index with the bitwise
// AND against the bucket count minus one (hash & (n-1)), which keeps
// only the low bits and yields a value in 0..n-1.
//
// One hash per operation. Hashing happens exactly once per Put or
// Get, no matter what. Collisions do not cause a second hash: if two
// keys land in the same bucket, both entries just sit in that
// bucket's chain, and a lookup tells them apart by comparing full
// keys with == as it walks. The hash picks the bucket; the key
// comparison picks the entry.
//
// Steps:
//  1. Construction: allocate a bucket array whose length is a power
//     of two large enough for the size hint, and start the entry
//     count at zero. Each slot starts as an empty chain. A power of
//     two matters because a bucket is picked by keeping only the low
//     bits of the hash (hash & (n-1)), which works exactly when the
//     count is a power of two.
//  2. Put: hash the key to a bucket and walk its chain. If an entry
//     with the same key is found, replace its value; the entry count
//     does not change. If the chain runs out without a match, append
//     a new entry and count one more.
//  3. Get: hash to a bucket and walk the chain the same way. A
//     matching key ends the walk with its value; a chain that runs
//     out ends the walk with "absent".
//  4. After an insert that grew the entry count, if the count has
//     reached the bucket count (load factor about one), rehash:
//     allocate a new bucket array of double the length, then walk
//     every entry in every old chain and insert it into the new
//     array (hash again with the new count to pick its bucket). This
//     is not a copy: entries land in different buckets than before,
//     which is why it cannot be left to the built-in slice growth.
//
// Notes:
//   - The complexity test inserts and looks up unbounded numbers of
//     keys, so skipping the rehash makes the exponent drift above
//     constant and the test fails.
//   - Chains make deletion easy (unlink one entry); that is the
//     trade being explored here, compared with open_map.
//
// Expected: O(1) average per op (amortized with rehashing)
// Input: not mutated
type ChainMap struct {
	bucks [][]entry
	n     int // number of entries
}

type entry struct {
	k string
	v int
}

func NewChainMap(sizeHint int) *ChainMap {
	sizeHint = max(sizeHint, 1)
	// s = size * 2
	bucks := make([][]entry, uint32(sizeHint*2))
	return &ChainMap{bucks: bucks, n: sizeHint}
}

func (m *ChainMap) Put(key string, val int) {
	nbucks := len(m.bucks)
	var i = FNV1a(key) & uint32((nbucks - 1))
	var overwrite = false
	for j := range m.bucks[i] {
		if m.bucks[i][j].k == key {
			m.bucks[i][j].v = val
			overwrite = true
		}
	}

	if !overwrite {
		m.bucks[i] = append(m.bucks[i], entry{k: key, v: val})
		m.n++
	}

	if m.n == len(m.bucks) {
		newMap := NewChainMap(m.n)
		for k := range m.bucks {
			for _, oldEntry := range m.bucks[k] {
				newMap.Put(oldEntry.k, oldEntry.v)
			}
		}
		m.bucks = newMap.bucks
		m.n = newMap.n
	}
}

func (m *ChainMap) Get(key string) (int, bool) {
	nbucks := len(m.bucks)
	var i = FNV1a(key) & uint32((nbucks - 1))
	for j := range m.bucks[i] {
		if m.bucks[i][j].k == key {
			return m.bucks[i][j].v, true
		}
	}

	return 0, false
}
