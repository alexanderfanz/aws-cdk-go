package cxapi


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataEntryResult := &metadataEntryResult{
//   	path: jsii.String("path"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	data: jsii.String("data"),
//   	trace: []*string{
//   		jsii.String("trace"),
//   	},
//   }
//
// Experimental.
type MetadataEntryResult struct {
	// The type of the metadata entry.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// The data.
	// Experimental.
	Data interface{} `field:"optional" json:"data" yaml:"data"`
	// A stack trace for when the entry was created.
	// Experimental.
	Trace *[]*string `field:"optional" json:"trace" yaml:"trace"`
	// The path in which this entry was defined.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
}

