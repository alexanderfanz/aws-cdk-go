package awsstepfunctions


// The `TagsEntry` property specifies *tags* to identify a state machine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tagsEntryProperty := &TagsEntryProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
type CfnStateMachine_TagsEntryProperty struct {
	// The `key` for a key-value pair in a tag entry.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The `value` for a key-value pair in a tag entry.
	Value *string `field:"required" json:"value" yaml:"value"`
}

