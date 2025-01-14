package awsscheduler


// The task placement strategy for a task or service.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   placementStrategyProperty := &PlacementStrategyProperty{
//   	Field: jsii.String("field"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-placementstrategy.html
//
type CfnSchedule_PlacementStrategyProperty struct {
	// The field to apply the placement strategy against.
	//
	// For the spread placement strategy, valid values are `instanceId` (or `instanceId` , which has the same effect), or any platform or custom attribute that is applied to a container instance, such as `attribute:ecs.availability-zone` . For the binpack placement strategy, valid values are `cpu` and `memory` . For the random placement strategy, this field is not used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-placementstrategy.html#cfn-scheduler-schedule-placementstrategy-field
	//
	Field *string `field:"optional" json:"field" yaml:"field"`
	// The type of placement strategy.
	//
	// The random placement strategy randomly places tasks on available candidates. The spread placement strategy spreads placement across available candidates evenly based on the field parameter. The binpack strategy places tasks on available candidates that have the least available amount of the resource that is specified with the field parameter. For example, if you binpack on memory, a task is placed on the instance with the least amount of remaining memory (but still enough to run the task).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-placementstrategy.html#cfn-scheduler-schedule-placementstrategy-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

