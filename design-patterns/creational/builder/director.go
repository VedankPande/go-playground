package main

// Director struct that encapsulates a builder
type Director struct {
	builder iBuilder
}

// get house from builder
func (d *Director) buildHouse() House {
	return d.builder.getHouse()
}
