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
// Steps:
//  1. Construction: allocate a bucket array whose length is a power
//     of two large enough for the size hint, and start the entry
//     count at zero. A power of two matters because a bucket is
//     picked by keeping only the low bits of the hash, which works
//     exactly when the count is a power of two.
//  2. Put: hash the key to a bucket and walk its chain. If an entry
//     with the same key is found, replace its value; the entry count
//     does not change. If the chain runs out without a match, append
//     a new entry and count one more.
//  3. Get: hash to a bucket and walk the chain the same way. A
//     matching key ends the walk with its value; a chain that runs
//     out ends the walk with "absent".
//  4. After an insert that grew the entry count, if the count has
//     reached the bucket count (load factor about one), rehash:
//     double the buckets and re-insert every entry.
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
