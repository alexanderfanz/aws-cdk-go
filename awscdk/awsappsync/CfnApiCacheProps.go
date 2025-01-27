package awsappsync


// Properties for defining a `CfnApiCache`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnApiCacheProps := &CfnApiCacheProps{
//   	ApiCachingBehavior: jsii.String("apiCachingBehavior"),
//   	ApiId: jsii.String("apiId"),
//   	Ttl: jsii.Number(123),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	AtRestEncryptionEnabled: jsii.Boolean(false),
//   	TransitEncryptionEnabled: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apicache.html
//
type CfnApiCacheProps struct {
	// Caching behavior.
	//
	// - *FULL_REQUEST_CACHING* : All requests are fully cached.
	// - *PER_RESOLVER_CACHING* : Individual resolvers that you specify are cached.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apicache.html#cfn-appsync-apicache-apicachingbehavior
	//
	ApiCachingBehavior *string `field:"required" json:"apiCachingBehavior" yaml:"apiCachingBehavior"`
	// The GraphQL API ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apicache.html#cfn-appsync-apicache-apiid
	//
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// TTL in seconds for cache entries.
	//
	// Valid values are 1–3,600 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apicache.html#cfn-appsync-apicache-ttl
	//
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
	// The cache instance type. Valid values are.
	//
	// - `SMALL`
	// - `MEDIUM`
	// - `LARGE`
	// - `XLARGE`
	// - `LARGE_2X`
	// - `LARGE_4X`
	// - `LARGE_8X` (not available in all regions)
	// - `LARGE_12X`
	//
	// Historically, instance types were identified by an EC2-style value. As of July 2020, this is deprecated, and the generic identifiers above should be used.
	//
	// The following legacy instance types are available, but their use is discouraged:
	//
	// - *T2_SMALL* : A t2.small instance type.
	// - *T2_MEDIUM* : A t2.medium instance type.
	// - *R4_LARGE* : A r4.large instance type.
	// - *R4_XLARGE* : A r4.xlarge instance type.
	// - *R4_2XLARGE* : A r4.2xlarge instance type.
	// - *R4_4XLARGE* : A r4.4xlarge instance type.
	// - *R4_8XLARGE* : A r4.8xlarge instance type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apicache.html#cfn-appsync-apicache-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// At-rest encryption flag for cache.
	//
	// You cannot update this setting after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apicache.html#cfn-appsync-apicache-atrestencryptionenabled
	//
	AtRestEncryptionEnabled interface{} `field:"optional" json:"atRestEncryptionEnabled" yaml:"atRestEncryptionEnabled"`
	// Transit encryption flag when connecting to cache.
	//
	// You cannot update this setting after creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-apicache.html#cfn-appsync-apicache-transitencryptionenabled
	//
	TransitEncryptionEnabled interface{} `field:"optional" json:"transitEncryptionEnabled" yaml:"transitEncryptionEnabled"`
}

