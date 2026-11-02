package numerics

// [easy] NearEqual reports whether a and b agree to within a relative
// tolerance.
//
// Background. Comparing floats with equality is almost always wrong, but
// the naive fix — compare the absolute difference against a small constant
// — is also wrong, for a different reason. Absolute tolerance works at
// human scales and fails at machine scales: 1e-12 apart is a rounding
// error for quantities near 1e6, and a chasm for quantities near 1e-12.
// What is (usually) wanted is relative tolerance: are the two values
// within the same few ulps?
//
// The standard formulation: the values are near when the difference is at
// most the tolerance times the larger of the two magnitudes. Pure relative
// comparison still has two edge cases worth handling deliberately. First,
// near zero the relative test is meaningless: relative to 1e-300, even
// 1e-299 is a factor of ten away, so if the difference is smaller than an
// absolute floor it counts as near regardless. Second, infinities of the
// same sign should compare near (they are the same value), while NaN
// compares near nothing, not even itself.
//
// Steps:
//  1. If both are the same exact value, including identical infinities,
//     they are near. If either is NaN, they are not.
//  2. Compute the absolute difference and compare it against the tolerance
//     scaled by the larger magnitude of a and b.
//  3. Also compare the absolute difference against an absolute floor so
//     values near zero are not held to a relative standard.
//
// Convention (matches the tests): the tolerance is relative (a fraction,
// not an ulp count) and the floor is fixed by the caller. Either argument
// being NaN makes the result false; equal infinities make it true.
//
// Expected: O(1)
func NearEqual(a, b, rtol, absFloor float64) bool {
	// TODO: implement
	return false
}
