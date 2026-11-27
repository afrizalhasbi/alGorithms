package trees

// [medium] Trie is a prefix tree over lowercase words. It supports insert,
// exact search, prefix search, and counting words with a given prefix.
//
// Background. A trie stores words by their characters rather than by
// hashing them. Each edge is labeled with one character, and the path from
// the root spells a prefix. Words that share a prefix share the path, so
// the cost of an operation is the length of the word, independent of how
// many words are stored. A node marks whether a word ends there. A prefix
// query succeeds when its path exists. Counting words with a prefix needs
// one more number per node: how many inserted words pass through it,
// including words that end exactly at it.
//
// Definitions:
//   - node: one position in the trie. The root represents the empty
//     prefix.
//   - child: the node reached by appending one lowercase letter to the
//     current prefix. A node has at most 26 children.
//   - terminal count: the number of inserted words that end exactly at a
//     node. Inserting the same word twice raises it twice.
//   - prefix count: the number of inserted words whose path passes through
//     a node, including words that end there. The root's prefix count is
//     the total number of inserted words.
//   - empty prefix: the root itself. It is a prefix of every word.
//
// Steps:
//  1. Construction: create a root node with terminal count and prefix
//     count zero.
//  2. Insert(word): start at the root and increase its prefix count by
//     one. For each character of the word in order, follow the child for
//     that character, creating it with zero counts when it does not
//     exist, and increase that child's prefix count by one. After the
//     last character, increase the terminal count of the final node by
//     one.
//  3. Search(word): start at the root and follow the child for each
//     character. If a child is missing, return false. After the last
//     character, return whether the final node's terminal count is
//     greater than zero.
//  4. StartsWith(prefix): follow the path for the prefix. Return false if
//     a child is missing, and true otherwise. The empty prefix always
//     exists.
//  5. CountPrefix(prefix): follow the path for the prefix and return the
//     final node's prefix count. A missing child means zero. The empty
//     prefix returns the total number of inserted words.
//
// Notes:
//   - Words are over the lowercase letters a-z. Other bytes are outside
//     the contract.
//   - Duplicate words are counted separately by CountPrefix and by
//     Search, which only reports presence.
//
// Expected: O(L) per operation for a word of length L, O(total characters)
// space
// Input: mutated in place
type Trie struct {
	// TODO: fields
}

// NewTrie returns an empty trie.
func NewTrie() *Trie {
	// TODO: implement
	return nil
}

// Insert adds word to the trie.
func (t *Trie) Insert(word string) {
	// TODO: implement
}

// Search reports whether word was inserted.
func (t *Trie) Search(word string) bool {
	// TODO: implement
	return false
}

// StartsWith reports whether any inserted word begins with prefix.
func (t *Trie) StartsWith(prefix string) bool {
	// TODO: implement
	return false
}

// CountPrefix returns the number of inserted words that begin with
// prefix.
func (t *Trie) CountPrefix(prefix string) int {
	// TODO: implement
	return 0
}
