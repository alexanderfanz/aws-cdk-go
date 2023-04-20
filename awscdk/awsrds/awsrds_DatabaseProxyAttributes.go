package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// Properties that describe an existing DB Proxy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//
//   databaseProxyAttributes := &DatabaseProxyAttributes{
//   	DbProxyArn: jsii.String("dbProxyArn"),
//   	DbProxyName: jsii.String("dbProxyName"),
//   	Endpoint: jsii.String("endpoint"),
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   }
//
// Experimental.
type DatabaseProxyAttributes struct {
	// DB Proxy ARN.
	// Experimental.
	DbProxyArn *string `field:"required" json:"dbProxyArn" yaml:"dbProxyArn"`
	// DB Proxy Name.
	// Experimental.
	DbProxyName *string `field:"required" json:"dbProxyName" yaml:"dbProxyName"`
	// Endpoint.
	// Experimental.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The security groups of the instance.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"required" json:"securityGroups" yaml:"securityGroups"`
}

