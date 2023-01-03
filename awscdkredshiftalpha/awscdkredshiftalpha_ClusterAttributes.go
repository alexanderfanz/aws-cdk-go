// The CDK Construct Library for AWS::Redshift
package awscdkredshiftalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties that describe an existing cluster instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import redshift_alpha "github.com/aws/aws-cdk-go/awscdkredshiftalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//
//   clusterAttributes := &clusterAttributes{
//   	clusterEndpointAddress: jsii.String("clusterEndpointAddress"),
//   	clusterEndpointPort: jsii.Number(123),
//   	clusterName: jsii.String("clusterName"),
//
//   	// the properties below are optional
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   }
//
// Experimental.
type ClusterAttributes struct {
	// Cluster endpoint address.
	// Experimental.
	ClusterEndpointAddress *string `field:"required" json:"clusterEndpointAddress" yaml:"clusterEndpointAddress"`
	// Cluster endpoint port.
	// Experimental.
	ClusterEndpointPort *float64 `field:"required" json:"clusterEndpointPort" yaml:"clusterEndpointPort"`
	// Identifier for the cluster.
	// Experimental.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The security groups of the redshift cluster.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

