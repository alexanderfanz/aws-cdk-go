package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnLocationObjectStorage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnLocationObjectStorageProps := &cfnLocationObjectStorageProps{
//   	agentArns: []*string{
//   		jsii.String("agentArns"),
//   	},
//   	bucketName: jsii.String("bucketName"),
//   	serverHostname: jsii.String("serverHostname"),
//
//   	// the properties below are optional
//   	accessKey: jsii.String("accessKey"),
//   	secretKey: jsii.String("secretKey"),
//   	serverPort: jsii.Number(123),
//   	serverProtocol: jsii.String("serverProtocol"),
//   	subdirectory: jsii.String("subdirectory"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLocationObjectStorageProps struct {
	// Specifies the Amazon Resource Names (ARNs) of the DataSync agents that can securely connect with your location.
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// Specifies the name of the object storage bucket involved in the transfer.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Specifies the domain name or IP address of the object storage server.
	//
	// A DataSync agent uses this hostname to mount the object storage server in a network.
	ServerHostname *string `field:"required" json:"serverHostname" yaml:"serverHostname"`
	// Specifies the access key (for example, a user name) if credentials are required to authenticate with the object storage server.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Specifies the secret key (for example, a password) if credentials are required to authenticate with the object storage server.
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// Specifies the port that your object storage server accepts inbound network traffic on (for example, port 443).
	ServerPort *float64 `field:"optional" json:"serverPort" yaml:"serverPort"`
	// Specifies the protocol that your object storage server uses to communicate.
	ServerProtocol *string `field:"optional" json:"serverProtocol" yaml:"serverProtocol"`
	// Specifies the object prefix for your object storage server.
	//
	// If this is a source location, DataSync only copies objects with this prefix. If this is a destination location, DataSync writes all objects with this prefix.
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// Specifies the key-value pair that represents a tag that you want to add to the resource.
	//
	// Tags can help you manage, filter, and search for your resources. We recommend creating a name tag for your location.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

