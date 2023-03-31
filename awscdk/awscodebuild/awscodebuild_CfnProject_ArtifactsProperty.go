package awscodebuild


// `Artifacts` is a property of the [AWS::CodeBuild::Project](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codebuild-project.html) resource that specifies output settings for artifacts generated by an AWS CodeBuild build.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   artifactsProperty := &artifactsProperty{
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	artifactIdentifier: jsii.String("artifactIdentifier"),
//   	encryptionDisabled: jsii.Boolean(false),
//   	location: jsii.String("location"),
//   	name: jsii.String("name"),
//   	namespaceType: jsii.String("namespaceType"),
//   	overrideArtifactName: jsii.Boolean(false),
//   	packaging: jsii.String("packaging"),
//   	path: jsii.String("path"),
//   }
//
type CfnProject_ArtifactsProperty struct {
	// The type of build output artifact. Valid values include:.
	//
	// - `CODEPIPELINE` : The build project has build output generated through CodePipeline.
	//
	// > The `CODEPIPELINE` type is not supported for `secondaryArtifacts` .
	// - `NO_ARTIFACTS` : The build project does not produce any build output.
	// - `S3` : The build project stores build output in Amazon S3.
	Type *string `field:"required" json:"type" yaml:"type"`
	// An identifier for this artifact definition.
	ArtifactIdentifier *string `field:"optional" json:"artifactIdentifier" yaml:"artifactIdentifier"`
	// Set to true if you do not want your output artifacts encrypted.
	//
	// This option is valid only if your artifacts type is Amazon Simple Storage Service (Amazon S3). If this is set with another artifacts type, an `invalidInputException` is thrown.
	EncryptionDisabled interface{} `field:"optional" json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// Information about the build output artifact location:.
	//
	// - If `type` is set to `CODEPIPELINE` , AWS CodePipeline ignores this value if specified. This is because CodePipeline manages its build output locations instead of CodeBuild .
	// - If `type` is set to `NO_ARTIFACTS` , this value is ignored if specified, because no build output is produced.
	// - If `type` is set to `S3` , this is the name of the output bucket.
	//
	// If you specify `CODEPIPELINE` or `NO_ARTIFACTS` for the `Type` property, don't specify this property. For all of the other types, you must specify this property.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Along with `path` and `namespaceType` , the pattern that AWS CodeBuild uses to name and store the output artifact:.
	//
	// - If `type` is set to `CODEPIPELINE` , AWS CodePipeline ignores this value if specified. This is because CodePipeline manages its build output names instead of AWS CodeBuild .
	// - If `type` is set to `NO_ARTIFACTS` , this value is ignored if specified, because no build output is produced.
	// - If `type` is set to `S3` , this is the name of the output artifact object. If you set the name to be a forward slash ("/"), the artifact is stored in the root of the output bucket.
	//
	// For example:
	//
	// - If `path` is set to `MyArtifacts` , `namespaceType` is set to `BUILD_ID` , and `name` is set to `MyArtifact.zip` , then the output artifact is stored in `MyArtifacts/ *build-ID* /MyArtifact.zip` .
	// - If `path` is empty, `namespaceType` is set to `NONE` , and `name` is set to " `/` ", the output artifact is stored in the root of the output bucket.
	// - If `path` is set to `MyArtifacts` , `namespaceType` is set to `BUILD_ID` , and `name` is set to " `/` ", the output artifact is stored in `MyArtifacts/ *build-ID*` .
	//
	// If you specify `CODEPIPELINE` or `NO_ARTIFACTS` for the `Type` property, don't specify this property. For all of the other types, you must specify this property.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Along with `path` and `name` , the pattern that AWS CodeBuild uses to determine the name and location to store the output artifact:  - If `type` is set to `CODEPIPELINE` , CodePipeline ignores this value if specified.
	//
	// This is because CodePipeline manages its build output names instead of AWS CodeBuild .
	// - If `type` is set to `NO_ARTIFACTS` , this value is ignored if specified, because no build output is produced.
	// - If `type` is set to `S3` , valid values include:
	//
	// - `BUILD_ID` : Include the build ID in the location of the build output artifact.
	// - `NONE` : Do not include the build ID. This is the default if `namespaceType` is not specified.
	//
	// For example, if `path` is set to `MyArtifacts` , `namespaceType` is set to `BUILD_ID` , and `name` is set to `MyArtifact.zip` , the output artifact is stored in `MyArtifacts/<build-ID>/MyArtifact.zip` .
	NamespaceType *string `field:"optional" json:"namespaceType" yaml:"namespaceType"`
	// If set to true a name specified in the buildspec file overrides the artifact name.
	//
	// The name specified in a buildspec file is calculated at build time and uses the Shell command language. For example, you can append a date and time to your artifact name so that it is always unique.
	OverrideArtifactName interface{} `field:"optional" json:"overrideArtifactName" yaml:"overrideArtifactName"`
	// The type of build output artifact to create:.
	//
	// - If `type` is set to `CODEPIPELINE` , CodePipeline ignores this value if specified. This is because CodePipeline manages its build output artifacts instead of AWS CodeBuild .
	// - If `type` is set to `NO_ARTIFACTS` , this value is ignored if specified, because no build output is produced.
	// - If `type` is set to `S3` , valid values include:
	//
	// - `NONE` : AWS CodeBuild creates in the output bucket a folder that contains the build output. This is the default if `packaging` is not specified.
	// - `ZIP` : AWS CodeBuild creates in the output bucket a ZIP file that contains the build output.
	Packaging *string `field:"optional" json:"packaging" yaml:"packaging"`
	// Along with `namespaceType` and `name` , the pattern that AWS CodeBuild uses to name and store the output artifact:.
	//
	// - If `type` is set to `CODEPIPELINE` , CodePipeline ignores this value if specified. This is because CodePipeline manages its build output names instead of AWS CodeBuild .
	// - If `type` is set to `NO_ARTIFACTS` , this value is ignored if specified, because no build output is produced.
	// - If `type` is set to `S3` , this is the path to the output artifact. If `path` is not specified, `path` is not used.
	//
	// For example, if `path` is set to `MyArtifacts` , `namespaceType` is set to `NONE` , and `name` is set to `MyArtifact.zip` , the output artifact is stored in the output bucket at `MyArtifacts/MyArtifact.zip` .
	Path *string `field:"optional" json:"path" yaml:"path"`
}

