package postmaster

type Package struct {
	Width  float32
	Height float32
	Length float32
	Weight float32
	WeightUnits string `dontMap:"true" json:"weight_units"`
	Type string `dontMap:"true"`
	LabelUrl string `dontMap:"true" json:"label_url"`
	DimensionUnits string `dontMap:"true" json:"dimension_units"`
}
