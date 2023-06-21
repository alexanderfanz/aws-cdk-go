package awsquicksight


// The unaggregated field wells of a scatter plot.
//
// Example:
//
//
type CfnAnalysis_ScatterPlotUnaggregatedFieldWellsProperty struct {
	// The size field well of a scatter plot.
	Size interface{} `field:"optional" json:"size" yaml:"size"`
	// The x-axis field well of a scatter plot.
	//
	// The x-axis is a dimension field and cannot be aggregated.
	XAxis interface{} `field:"optional" json:"xAxis" yaml:"xAxis"`
	// The y-axis field well of a scatter plot.
	//
	// The y-axis is a dimension field and cannot be aggregated.
	YAxis interface{} `field:"optional" json:"yAxis" yaml:"yAxis"`
}

