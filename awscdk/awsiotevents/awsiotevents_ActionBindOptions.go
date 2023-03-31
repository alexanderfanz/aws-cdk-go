package awsiotevents

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
)

// Options when binding a Action to a detector model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role role
//
//   actionBindOptions := &actionBindOptions{
//   	role: role,
//   }
//
// Experimental.
type ActionBindOptions struct {
	// The IAM role assumed by IoT Events to perform the action.
	// Experimental.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
}

