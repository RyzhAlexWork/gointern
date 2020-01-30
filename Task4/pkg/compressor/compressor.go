package compressor

// Compressor ...
type Compressor interface {
	Increase(bar int)
	Reduce(bar int)
	GiveValuePressure() (valuePressure int)
}

type compressor struct {
	bar int
}

// Increase value pressure.
func (c *compressor) Increase(bar int) {
	c.bar += bar
}

// Reduce value pressure.
func (c *compressor) Reduce(bar int) {
	c.bar -= bar
}

// GiveValuePressure return value pressure.
func (c *compressor) GiveValuePressure() (valuePressure int) {
	valuePressure = c.bar
	return
}

// NewCompressor create compressor implementation for interface Compressor.
func NewCompressor() Compressor {
	return &compressor{}
}
