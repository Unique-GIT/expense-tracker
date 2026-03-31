package types


type Label struct {
	name string
}

type Record struct {
	Name string
	Label Label
	Cost float32
}

type AnalyseConfig struct {
	days int
	months int
	years int
}

type ReportRow struct {
	Label Label
	Cost  float32
	CostPercentage float32
}

type Report []ReportRow
