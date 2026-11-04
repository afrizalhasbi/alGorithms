// Package sampling holds the random-sampling exercises: turning uniform
// draws into draws from a target distribution, by inversion, rejection,
// transformation, and Markov chains.
package sampling

// [medium] InverseDiscrete samples from a discrete distribution with the
// given point probabilities, given a list of uniform draws in [0, 1).
//
// Background. The inverse transform sampler is the master recipe: if U is
// uniform on the zero-one interval and F is a cumulative distribution
// function, then the value of F's inverse at U has distribution F. For a
// discrete distribution F is a step function — a staircase that jumps by
// the point probability at each outcome — and its inverse says: draw a
// uniform, and return the first outcome whose cumulative probability
// exceeds it. Equivalently: the uniform draw chooses a point in the
// interval from zero to one, and the outcome is the step of the staircase
// the point lands on, with each step's width equal to its probability.
// That picture is the whole correctness argument: a uniform point is
// twice as likely to land on a step twice as wide.
//
// The exercise's substance is the edge cases of the staircase walk.
// Draws of exactly zero must return the first outcome with positive
// probability, not fall through a zero-width step. The comparison
// convention (strictly less than the cumulative sum, or less-or-equal)
// decides which outcome a draw exactly at a step boundary receives, and
// while the difference has probability zero in theory, in floating point
// boundary hits happen — the fixture places uniforms exactly on
// boundaries and pins down the convention. And the last cumulative
// entries must be treated as exactly one (see the statistics package's
// CumulativeDistribution), so a draw of just under one cannot walk off
// the end of the list; the fixture includes one such draw.
//
// Definitions:
//  - pmf: the point probability of each outcome, one entry per
//     outcome, in outcome order.
//  - staircase: the cumulative sums of the pmf — step k's height is
//     the pmf up to and including outcome k.
//  - the walk: for each uniform draw, scanning the cumulative sums in
//     order and returning the first index whose cumulative sum is
//     STRICTLY greater than the draw. The strictness is the boundary
//     convention: a draw exactly on a step boundary belongs to the
//     NEXT step.
//  - walk-off-the-end: a draw that reaches the final cumulative sum
//     through rounding. Handled by returning the last index with
//     positive probability; the final cumulative entry is treated as
//     exactly one (see statistics.CumulativeDistribution).
//
// Steps:
//  1. Validate: every probability is non-negative; the list is non-empty.
//  2. For each uniform draw: walk the cumulative sums in order and
//     return the index of the first step whose cumulative sum is
//     strictly greater than the draw.
//  3. If no step qualifies (the draw reached the final cumulative sum
//     through rounding), return the last index with positive probability.
//
// Convention (matches the tests): the fixture stores the probability
// list, the uniform draws (including exact-boundary values and
// near-one values), and the expected outcome indices, which were
// computed with the convention above.
//
// Expected: O(draws times outcomes) in the direct form
// Input: not mutated
func InverseDiscrete(pmf []float64, us []float64) []int {
	// TODO: implement
	return nil
}
