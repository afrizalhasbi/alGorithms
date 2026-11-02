package hashing

// [medium] ChainMap is a string->int map built with separate chaining.
//
// Bucket array of linked buckets; collisions append to the bucket's chain.
// Steps:
//  1. Construction: pick a number of buckets that is a power of two
//     and comfortably holds the expected size; remember how many
//     entries the map holds. Hash each key and reduce it to a bucket
//     by keeping only the low bits — which is exactly why the bucket
//     count should be a power of two.
//  2. Putting: hash to a bucket, then walk that bucket's chain of
//     entries. The walk ends when a matching key is found (overwrite
//     its value; the entry count does not change) or when the chain
//     runs out without a match (append a new entry at the end; the
//     entry count goes up by one).
//  3. Getting: hash to a bucket and walk the chain the same way. A
//     matching key ends the walk with its value; a chain that runs
//     out ends the walk with "absent".
//  4. Rehashing: after an insert, compare the entry count against
//     the bucket count. Once entries outnumber buckets, replace the
//     bucket array with one of double the length and re-walk every
//     old chain, inserting each entry into its new bucket. Skipping
//     this lets chains grow without bound and every operation
//     degrades toward a linear scan.
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
