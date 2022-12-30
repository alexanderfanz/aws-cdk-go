package awscodeguruprofiler


// Properties for creating a new Profiling Group.
//
// Example:
//   profilingGroup := codeguruprofiler.NewProfilingGroup(this, jsii.String("MyProfilingGroup"), &profilingGroupProps{
//   	computePlatform: codeguruprofiler.computePlatform_AWS_LAMBDA,
//   })
//
// Experimental.
type ProfilingGroupProps struct {
	// The compute platform of the profiling group.
	// Experimental.
	ComputePlatform ComputePlatform `field:"optional" json:"computePlatform" yaml:"computePlatform"`
	// A name for the profiling group.
	// Experimental.
	ProfilingGroupName *string `field:"optional" json:"profilingGroupName" yaml:"profilingGroupName"`
}

