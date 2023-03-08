package awsapigateway


// The `MethodSetting` property type configures settings for all methods in a stage.
//
// The `MethodSettings` property of the [Amazon API Gateway Deployment StageDescription](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-deployment-stagedescription.html) property type contains a list of `MethodSetting` property types.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   methodSettingProperty := &MethodSettingProperty{
//   	CacheDataEncrypted: jsii.Boolean(false),
//   	CacheTtlInSeconds: jsii.Number(123),
//   	CachingEnabled: jsii.Boolean(false),
//   	DataTraceEnabled: jsii.Boolean(false),
//   	HttpMethod: jsii.String("httpMethod"),
//   	LoggingLevel: jsii.String("loggingLevel"),
//   	MetricsEnabled: jsii.Boolean(false),
//   	ResourcePath: jsii.String("resourcePath"),
//   	ThrottlingBurstLimit: jsii.Number(123),
//   	ThrottlingRateLimit: jsii.Number(123),
//   }
//
type CfnDeployment_MethodSettingProperty struct {
	// Specifies whether the cached responses are encrypted.
	CacheDataEncrypted interface{} `field:"optional" json:"cacheDataEncrypted" yaml:"cacheDataEncrypted"`
	// Specifies the time to live (TTL), in seconds, for cached responses.
	//
	// The higher the TTL, the longer the response will be cached.
	CacheTtlInSeconds *float64 `field:"optional" json:"cacheTtlInSeconds" yaml:"cacheTtlInSeconds"`
	// Specifies whether responses should be cached and returned for requests.
	//
	// A cache cluster must be enabled on the stage for responses to be cached.
	CachingEnabled interface{} `field:"optional" json:"cachingEnabled" yaml:"cachingEnabled"`
	// Specifies whether data trace logging is enabled for this method, which affects the log entries pushed to Amazon CloudWatch Logs.
	DataTraceEnabled interface{} `field:"optional" json:"dataTraceEnabled" yaml:"dataTraceEnabled"`
	// The HTTP method.
	HttpMethod *string `field:"optional" json:"httpMethod" yaml:"httpMethod"`
	// Specifies the logging level for this method, which affects the log entries pushed to Amazon CloudWatch Logs.
	//
	// Valid values are `OFF` , `ERROR` , and `INFO` . Choose `ERROR` to write only error-level entries to CloudWatch Logs, or choose `INFO` to include all `ERROR` events as well as extra informational events.
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// Specifies whether Amazon CloudWatch metrics are enabled for this method.
	MetricsEnabled interface{} `field:"optional" json:"metricsEnabled" yaml:"metricsEnabled"`
	// The resource path for this method.
	//
	// Forward slashes ( `/` ) are encoded as `~1` and the initial slash must include a forward slash. For example, the path value `/resource/subresource` must be encoded as `/~1resource~1subresource` . To specify the root path, use only a slash ( `/` ).
	ResourcePath *string `field:"optional" json:"resourcePath" yaml:"resourcePath"`
	// Specifies the throttling burst limit.
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// Specifies the throttling rate limit.
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
}

