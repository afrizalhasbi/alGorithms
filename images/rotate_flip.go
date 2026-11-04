package images

import "alGorithms/internal/bitmap"

// [easy] Rotate90 returns the image rotated 90° clockwise: pixel
// (x,y) of a WxH image moves to (H-1-y, x) in an HxW image.
//
// Background. A 90° rotation swaps the axes of the grid and reverses
// one of them. Walking a row of the source top to bottom produces a
// column of the destination right to left. The index map is: source
// (x,y) lands at destination (H-1-y, x), and the destination is H
// wide and W tall. Widths and heights swap; a non-square image
// changes shape, so four rotations are the identity, not two.
//
// FlipH mirrors horizontally: (x,y) moves to (W-1-x, y). It is its
// own inverse — two flips are the identity — and it does not change
// dimensions.
//
// Definitions:
//  - clockwise: the top row of the source becomes the right column
//     of the destination.
//  - involution: an operation equal to its own inverse.
//
// Steps:
//  1. Rotate90: allocate an HxW destination, then for every source
//     pixel write dst.Set(H-1-y, x, color).
//  2. FlipH: allocate a WxH destination, then for every source pixel
//     write dst.Set(W-1-x, y, color).
//
// Notes:
//   - The test checks the exact corner mappings, the dimension swap,
//     that rotating four times is the identity, and that flipping
//     twice is the identity.
//
// Expected: O(W*H) time and space for both
// Input: not mutated
func Rotate90(src *bitmap.RGB) *bitmap.RGB {
	// TODO: implement
	return nil
}

// FlipH returns the image mirrored horizontally, as documented above.
func FlipH(src *bitmap.RGB) *bitmap.RGB {
	// TODO: implement
	return nil
}
