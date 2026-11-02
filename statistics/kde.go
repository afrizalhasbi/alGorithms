package statistics

// [medium] KDE1D evaluates a kernel density estimate of the data at each
// query point, with a Gaussian kernel and a bandwidth given by the caller.
//
// Background. A histogram is a density estimate with a cliff for a
// kernel: every point contributes everything to exactly one bin and
// nothing to its neighbors, so the estimate jumps when the bins shift.
// A kernel density estimate replaces each data point with a smooth bump —
// here, a Gaussian centered on the point with standard deviation equal to
// the bandwidth — and sums the bumps. The estimate at any query point is
// the sum of the contributions of all data points, normalized so the
// estimate integrates to one: each bump is a scaled Gaussian, and the
// scaling makes each data point contribute one over (count times
// bandwidth) times the standard normal density at the standardized
// distance.
//
// The bandwidth is the deep parameter: too small and the estimate is a
// spiky re-rendering of the data, too large and it is a single lump that
// has forgotten the data. The exercise fixes the bandwidth as an input —
// choosing one is a statistics question, not a programming one — and the
// tests check the exact formula against a reference computed by an
// independent implementation.
//
// The cost is one kernel evaluation per (data point, query point) pair.
// Nothing about the estimate requires sorting, binning, or any structure;
// it is the honest, direct form, and the complexity test checks that the
// cost grows like the product of the two list lengths.
//
// Steps:
//  1. Validate: bandwidth positive, at least one data point, at least one
//     query point.
//  2. For each query point: sum over data points of the standard normal
//     density at (query minus point) over bandwidth, times one over
//     (count times bandwidth).
//  3. Return one density value per query point, in order.
//
// Convention (matches the tests): the fixture stores the data, the
// bandwidth, the query grid, and reference densities; agreement is
// expected to a relative tolerance of about one part in a billion.
//
// Expected: O(len(xs) times len(at))
// Input: not mutated
func KDE1D(xs []float64, bandwidth float64, at []float64) []float64 {
	// TODO: implement
	return nil
}
