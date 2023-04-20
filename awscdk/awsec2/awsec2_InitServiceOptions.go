package awsec2


// Options for an InitService.
//
// Example:
//   var myBucket bucket
//
//
//   handle := ec2.NewInitServiceRestartHandle()
//
//   ec2.CloudFormationInit_FromElements(ec2.InitFile_FromString(jsii.String("/etc/nginx/nginx.conf"), jsii.String("..."), &InitFileOptions{
//   	ServiceRestartHandles: []initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.InitSource_FromS3Object(jsii.String("/var/www/html"), myBucket, jsii.String("html.zip"), &InitSourceOptions{
//   	ServiceRestartHandles: []*initServiceRestartHandle{
//   		handle,
//   	},
//   }), ec2.InitService_Enable(jsii.String("nginx"), &InitServiceOptions{
//   	ServiceRestartHandle: handle,
//   }))
//
// Experimental.
type InitServiceOptions struct {
	// Enable or disable this service.
	//
	// Set to true to ensure that the service will be started automatically upon boot.
	//
	// Set to false to ensure that the service will not be started automatically upon boot.
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// Make sure this service is running or not running after cfn-init finishes.
	//
	// Set to true to ensure that the service is running after cfn-init finishes.
	//
	// Set to false to ensure that the service is not running after cfn-init finishes.
	// Experimental.
	EnsureRunning *bool `field:"optional" json:"ensureRunning" yaml:"ensureRunning"`
	// Restart service when the actions registered into the restartHandle have been performed.
	//
	// Register actions into the restartHandle by passing it to `InitFile`, `InitCommand`,
	// `InitPackage` and `InitSource` objects.
	// Experimental.
	ServiceRestartHandle InitServiceRestartHandle `field:"optional" json:"serviceRestartHandle" yaml:"serviceRestartHandle"`
}

