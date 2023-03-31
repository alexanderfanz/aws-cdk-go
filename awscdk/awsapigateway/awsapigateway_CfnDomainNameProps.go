package awsapigateway

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDomainName`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDomainNameProps := &cfnDomainNameProps{
//   	certificateArn: jsii.String("certificateArn"),
//   	domainName: jsii.String("domainName"),
//   	endpointConfiguration: &endpointConfigurationProperty{
//   		types: []*string{
//   			jsii.String("types"),
//   		},
//   	},
//   	mutualTlsAuthentication: &mutualTlsAuthenticationProperty{
//   		truststoreUri: jsii.String("truststoreUri"),
//   		truststoreVersion: jsii.String("truststoreVersion"),
//   	},
//   	ownershipVerificationCertificateArn: jsii.String("ownershipVerificationCertificateArn"),
//   	regionalCertificateArn: jsii.String("regionalCertificateArn"),
//   	securityPolicy: jsii.String("securityPolicy"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDomainNameProps struct {
	// The reference to an AWS -managed certificate that will be used by edge-optimized endpoint for this domain name.
	//
	// AWS Certificate Manager is the only supported source.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The custom domain name as an API host name, for example, `my-api.example.com` .
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The endpoint configuration of this DomainName showing the endpoint types of the domain name.
	EndpointConfiguration interface{} `field:"optional" json:"endpointConfiguration" yaml:"endpointConfiguration"`
	// The mutual TLS authentication configuration for a custom domain name.
	//
	// If specified, API Gateway performs two-way authentication between the client and the server. Clients must present a trusted certificate to access your API.
	MutualTlsAuthentication interface{} `field:"optional" json:"mutualTlsAuthentication" yaml:"mutualTlsAuthentication"`
	// The ARN of the public certificate issued by ACM to validate ownership of your custom domain.
	//
	// Only required when configuring mutual TLS and using an ACM imported or private CA certificate ARN as the RegionalCertificateArn.
	OwnershipVerificationCertificateArn *string `field:"optional" json:"ownershipVerificationCertificateArn" yaml:"ownershipVerificationCertificateArn"`
	// The reference to an AWS -managed certificate that will be used for validating the regional domain name.
	//
	// AWS Certificate Manager is the only supported source.
	RegionalCertificateArn *string `field:"optional" json:"regionalCertificateArn" yaml:"regionalCertificateArn"`
	// The Transport Layer Security (TLS) version + cipher suite for this DomainName.
	//
	// The valid values are `TLS_1_0` and `TLS_1_2` .
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// The collection of tags.
	//
	// Each tag element is associated with a given resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

