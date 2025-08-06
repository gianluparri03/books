package preview

import "image"

// Width is the width of the preview image.
var Width = 50

// calcHeight returns the appropriate height for the image with the given
// bounds in order to keep the source aspect ratio.
func calcHeight(img image.Image) int {
	b := img.Bounds()
	r := float64(b.Max.Y-b.Min.Y) / float64(b.Max.X-b.Min.X) / 2
	return int(float64(Width) * r)
}
