package sorting

// [hard] TimSort returns xs sorted in ascending order (simplified TimSort).
//
// Real-world sorting strategy: find naturally sorted "runs", extend short
// ones with insertion sort, then merge runs pairwise.
// Steps:
//  1. Walk the slice from the front, cutting it into runs — maximal
//     stretches that are already in ascending order. A stretch counting
//     downward also counts as a run, but only if it is STRICTLY
//     decreasing; reverse it in place to make it ascending. (Allowing
//     equal items in a downward run would break stability once it is
//     reversed.)
//  2. If a run comes out shorter than a fixed minimum length, extend
//     it: sort the following items into it, with binary insertion,
//     until the run reaches the minimum length or the slice ends.
//     Run-finding and extending are one interleaved walk: find a run,
//     extend it to the minimum length, then keep scanning from the
//     extended run's end.
//  3. Merge the runs pairwise, always merging the two leftmost runs
//     into one (taking from the left run on ties, as in merge sort)
//     and placing the merged run where the pair began. Each sweep
//     halves the number of runs; repeat the sweeping until a single
//     run spans the whole slice, at which point the slice is sorted.
//
// Notes:
//   - The hard part is the merging discipline; a simple left-to-right merge
//     of runs is acceptable for this exercise. Real TimSort maintains a stack
//     with size invariants to keep merges balanced.
//   - Already-sorted input is one single run: zero merges, O(n).
//   - Stable.
//
// Expected: best O(n), average/worst O(n log n)
// Input: mutated in place
func TimSort(xs []int) []int {
	// TODO: implement
	return nil
}
