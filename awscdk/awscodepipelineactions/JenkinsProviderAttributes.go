package awscodepipelineactions


// Properties for importing an existing Jenkins provider.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   jenkinsProvider := codepipeline_actions.JenkinsProvider_FromJenkinsProviderAttributes(this, jsii.String("JenkinsProvider"), &JenkinsProviderAttributes{
//   	ProviderName: jsii.String("MyJenkinsProvider"),
//   	ServerUrl: jsii.String("http://my-jenkins.com:8080"),
//   	Version: jsii.String("2"),
//   })
//
type JenkinsProviderAttributes struct {
	// The name of the Jenkins provider that you set in the AWS CodePipeline plugin configuration of your Jenkins project.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   "MyJenkinsProvider"
	//
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// The base URL of your Jenkins server.
	//
	// Example:
	//   // Example automatically generated from non-compiling source. May contain errors.
	//   "http://myjenkins.com:8080"
	//
	ServerUrl *string `field:"required" json:"serverUrl" yaml:"serverUrl"`
	// The version of your provider.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

