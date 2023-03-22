package awsapprunner

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Attributes for the App Runner VPC Connector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//
//   vpcConnectorAttributes := &vpcConnectorAttributes{
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	vpcConnectorArn: jsii.String("vpcConnectorArn"),
//   	vpcConnectorName: jsii.String("vpcConnectorName"),
//   	vpcConnectorRevision: jsii.Number(123),
//   }
//
// Experimental.
type VpcConnectorAttributes struct {
	// The security groups associated with the VPC connector.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// The ARN of the VPC connector.
	// Experimental.
	VpcConnectorArn *string `field:"required" json:"vpcConnectorArn" yaml:"vpcConnectorArn"`
	// The name of the VPC connector.
	// Experimental.
	VpcConnectorName *string `field:"required" json:"vpcConnectorName" yaml:"vpcConnectorName"`
	// The revision of the VPC connector.
	// Experimental.
	VpcConnectorRevision *float64 `field:"required" json:"vpcConnectorRevision" yaml:"vpcConnectorRevision"`
}

