package awsmediastore

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnContainer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContainerProps := &CfnContainerProps{
//   	ContainerName: jsii.String("containerName"),
//
//   	// the properties below are optional
//   	AccessLoggingEnabled: jsii.Boolean(false),
//   	CorsPolicy: []interface{}{
//   		&CorsRuleProperty{
//   			AllowedHeaders: []*string{
//   				jsii.String("allowedHeaders"),
//   			},
//   			AllowedMethods: []*string{
//   				jsii.String("allowedMethods"),
//   			},
//   			AllowedOrigins: []*string{
//   				jsii.String("allowedOrigins"),
//   			},
//   			ExposeHeaders: []*string{
//   				jsii.String("exposeHeaders"),
//   			},
//   			MaxAgeSeconds: jsii.Number(123),
//   		},
//   	},
//   	LifecyclePolicy: jsii.String("lifecyclePolicy"),
//   	MetricPolicy: &MetricPolicyProperty{
//   		ContainerLevelMetrics: jsii.String("containerLevelMetrics"),
//
//   		// the properties below are optional
//   		MetricPolicyRules: []interface{}{
//   			&MetricPolicyRuleProperty{
//   				ObjectGroup: jsii.String("objectGroup"),
//   				ObjectGroupName: jsii.String("objectGroupName"),
//   			},
//   		},
//   	},
//   	Policy: jsii.String("policy"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnContainerProps struct {
	// The name for the container.
	//
	// The name must be from 1 to 255 characters. Container names must be unique to your AWS account within a specific region. As an example, you could create a container named `movies` in every region, as long as you don’t have an existing container with that name.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// The state of access logging on the container.
	//
	// This value is `false` by default, indicating that AWS Elemental MediaStore does not send access logs to Amazon CloudWatch Logs. When you enable access logging on the container, MediaStore changes this value to `true` , indicating that the service delivers access logs for objects stored in that container to CloudWatch Logs.
	AccessLoggingEnabled interface{} `field:"optional" json:"accessLoggingEnabled" yaml:"accessLoggingEnabled"`
	// Sets the cross-origin resource sharing (CORS) configuration on a container so that the container can service cross-origin requests.
	//
	// For example, you might want to enable a request whose origin is http://www.example.com to access your AWS Elemental MediaStore container at my.example.container.com by using the browser's XMLHttpRequest capability.
	//
	// To enable CORS on a container, you attach a CORS policy to the container. In the CORS policy, you configure rules that identify origins and the HTTP methods that can be executed on your container. The policy can contain up to 398,000 characters. You can add up to 100 rules to a CORS policy. If more than one rule applies, the service uses the first applicable rule listed.
	//
	// To learn more about CORS, see [Cross-Origin Resource Sharing (CORS) in AWS Elemental MediaStore](https://docs.aws.amazon.com/mediastore/latest/ug/cors-policy.html) .
	CorsPolicy interface{} `field:"optional" json:"corsPolicy" yaml:"corsPolicy"`
	// Writes an object lifecycle policy to a container.
	//
	// If the container already has an object lifecycle policy, the service replaces the existing policy with the new policy. It takes up to 20 minutes for the change to take effect.
	//
	// For information about how to construct an object lifecycle policy, see [Components of an Object Lifecycle Policy](https://docs.aws.amazon.com/mediastore/latest/ug/policies-object-lifecycle-components.html) .
	LifecyclePolicy *string `field:"optional" json:"lifecyclePolicy" yaml:"lifecyclePolicy"`
	// `AWS::MediaStore::Container.MetricPolicy`.
	MetricPolicy interface{} `field:"optional" json:"metricPolicy" yaml:"metricPolicy"`
	// Creates an access policy for the specified container to restrict the users and clients that can access it.
	//
	// For information about the data that is included in an access policy, see the [AWS Identity and Access Management User Guide](https://docs.aws.amazon.com/iam/) .
	//
	// For this release of the REST API, you can create only one policy for a container. If you enter `PutContainerPolicy` twice, the second command modifies the existing policy.
	Policy *string `field:"optional" json:"policy" yaml:"policy"`
	// `AWS::MediaStore::Container.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}
