package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnHoursOfOperation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnHoursOfOperationProps := &CfnHoursOfOperationProps{
//   	Config: []interface{}{
//   		&HoursOfOperationConfigProperty{
//   			Day: jsii.String("day"),
//   			EndTime: &HoursOfOperationTimeSliceProperty{
//   				Hours: jsii.Number(123),
//   				Minutes: jsii.Number(123),
//   			},
//   			StartTime: &HoursOfOperationTimeSliceProperty{
//   				Hours: jsii.Number(123),
//   				Minutes: jsii.Number(123),
//   			},
//   		},
//   	},
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//   	TimeZone: jsii.String("timeZone"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnHoursOfOperationProps struct {
	// Configuration information for the hours of operation.
	Config interface{} `field:"required" json:"config" yaml:"config"`
	// The Amazon Resource Name (ARN) for the instance.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name for the hours of operation.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The time zone for the hours of operation.
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// The description for the hours of operation.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

