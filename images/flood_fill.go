package images

import "alGorithms/internal/bitmap"

// [easy] FloodFill recolors every pixel 4-connected to (x,y) that
// shares the original color of (x,y).
//
// Background. A raster image is a grid of pixels; two pixels are
// 4-connected when they are adjacent horizontally or vertically (not
// diagonally). A connected component is the set of pixels reachable
// from a start pixel by stepping only through pixels of one color.
// Flood fill recolors exactly that set — it is the paint-bucket tool.
//
// The traversal is the same graph walk as BFS/DFS on a maze: each
// pixel is a node, each horizontal/vertical adjacency is an edge, and
// only nodes matching the seed color are enterable. The seed color is
// sampled before any pixel is written, otherwise the fill would chase
// its own new color. Each pixel is colored when it is enqueued, not
// when it is dequeued, so no pixel enters the queue twice.
//
// Definitions:
//  - 4-connected: neighbors of (x,y) are (x±1,y) and (x,y±1) only.
//  - component: the seed pixel plus everything reachable from it
//     through pixels equal to the seed color.
//  - out-of-range: coordinates outside the buffer are not pixels;
//     neighbor steps past an edge are simply skipped.
//
// Steps:
//  1. Sample the seed color at (x,y). If the target color equals it,
//     there is nothing to do.
//  2. Enqueue the seed. Repeat while the queue is not empty: pop a
//     pixel, color it, and enqueue each in-range 4-neighbor whose
//     color is still the seed color.
//
// Notes:
//   - The buffer is mutated in place and also returned for chaining.
//   - The test checks that only the component changed: every pixel
//     of the original color reachable from the seed is recolored and
//     every pixel not reachable from the seed keeps its color.
//
// Expected: O(W*H) time, O(W*H) worst-case queue
// Input: mutated in place
func FloodFill(img *bitmap.RGB, x, y int, nr, ng, nb uint8) *bitmap.RGB {
	// TODO: implement
	return nil
}
