package images

import "alGorithms/internal/bitmap"

// [medium] BresenhamLine draws a line from (x0,y0) to (x1,y1) by
// coloring the grid pixels closest to the ideal line.
//
// Background. A raster has no lines, only pixels, so drawing a line
// means choosing one pixel per step along the axis of greatest travel
// such that the chosen pixels stay as close as possible to the ideal
// segment. The obvious approach evaluates y = y0 + (x-x0)*slope and
// rounds — but that needs a float add and a rounding per pixel, and
// floats drift. Bresenham's insight is that the error term is an
// integer: after each x step the ideal y moved by dy/dx, so keep the
// accumulated error and whenever it passes half a pixel, step y by
// one and pull the error back. Everything stays in integers, and the
// integer arithmetic can be arranged so no multiplication is needed
// after setup — only adds and sign checks.
//
// The general form handles all octets with one loop: dx and dy can
// each be negative, and the step in y is the sign of dy. The axis of
// greatest travel (|dx| >= |dy| or not) decides whether x steps every
// iteration or y does.
//
// Definitions:
//  - major axis: the direction (x or y) in which the line moves one
//     pixel per step. The minor axis moves at most one pixel, and
//     only when the accumulated error says so.
//  - error term: the accumulated difference between the ideal line
//     position and the drawn position, in minor-axis pixels, scaled
//     by 2*|major length| to stay integral.
//  - 8-connected: consecutive drawn pixels differ by at most 1 in
//     each coordinate — the line has no gaps.
//
// Steps:
//  1. Compute dx = x1-x0, dy = y1-y0 and the integer error term.
//  2. Step along the major axis one pixel at a time; at each step,
//     move the minor axis by one when the error term crosses its
//     threshold and subtract the threshold from the error.
//  3. Color each visited pixel, including both endpoints, and stop
//     after the endpoint is colored.
//
// Notes:
//   - The test checks both endpoints are drawn, every drawn pixel is
//     within half a pixel of the ideal segment, consecutive pixels
//     are 8-connected, and lines in all octets (shallow, steep,
//     ascending, descending) round-trip through symmetry: drawing
//     A-to-B and B-to-A colors the same set of pixels.
//
// Expected: O(max(|dx|,|dy|)) time, O(1) space
// Input: mutated in place
func BresenhamLine(img *bitmap.RGB, x0, y0, x1, y1 int, r, g, b uint8) {
	// TODO: implement
}
