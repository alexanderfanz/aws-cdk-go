package awscodepipeline

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// An output artifact of an action.
//
// Artifacts can be used as input by some actions.
//
// Example:
//   // later:
//   var project pipelineProject
//   lambdaInvokeAction := codepipeline_actions.NewLambdaInvokeAction(&LambdaInvokeActionProps{
//   	ActionName: jsii.String("Lambda"),
//   	Lambda: lambda.NewFunction(this, jsii.String("Func"), &FunctionProps{
//   		Runtime: lambda.Runtime_NODEJS_14_X(),
//   		Handler: jsii.String("index.handler"),
//   		Code: lambda.Code_FromInline(jsii.String(`
//   		        const AWS = require('aws-sdk');
//
//   		        exports.handler = async function(event, context) {
//   		            const codepipeline = new AWS.CodePipeline();
//   		            await codepipeline.putJobSuccessResult({
//   		                jobId: event['CodePipeline.job'].id,
//   		                outputVariables: {
//   		                    MY_VAR: "some value",
//   		                },
//   		            }).promise();
//   		        }
//   		    `)),
//   	}),
//   	VariablesNamespace: jsii.String("MyNamespace"),
//   })
//   sourceOutput := codepipeline.NewArtifact()
//   codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   	ActionName: jsii.String("CodeBuild"),
//   	Project: Project,
//   	Input: sourceOutput,
//   	EnvironmentVariables: map[string]buildEnvironmentVariable{
//   		"MyVar": &buildEnvironmentVariable{
//   			"value": lambdaInvokeAction.variable(jsii.String("MY_VAR")),
//   		},
//   	},
//   })
//
type Artifact interface {
	ArtifactName() *string
	// The artifact attribute for the name of the S3 bucket where the artifact is stored.
	BucketName() *string
	// The artifact attribute for The name of the .zip file that contains the artifact that is generated by AWS CodePipeline, such as 1ABCyZZ.zip.
	ObjectKey() *string
	// Returns the location of the .zip file in S3 that this Artifact represents. Used by Lambda's `CfnParametersCode` when being deployed in a CodePipeline.
	S3Location() *awss3.Location
	// The artifact attribute of the Amazon Simple Storage Service (Amazon S3) URL of the artifact, such as https://s3-us-west-2.amazonaws.com/artifactstorebucket-yivczw8jma0c/test/TemplateSo/1ABCyZZ.zip.
	Url() *string
	// Returns an ArtifactPath for a file within this artifact.
	//
	// CfnOutput is in the form "<artifact-name>::<file-name>".
	AtPath(fileName *string) ArtifactPath
	// Retrieve the metadata stored in this artifact under the given key.
	//
	// If there is no metadata stored under the given key,
	// null will be returned.
	GetMetadata(key *string) interface{}
	// Returns a token for a value inside a JSON file within this artifact.
	GetParam(jsonFile *string, keyName *string) *string
	// Add arbitrary extra payload to the artifact under a given key.
	//
	// This can be used by CodePipeline actions to communicate data between themselves.
	// If metadata was already present under the given key,
	// it will be overwritten with the new value.
	SetMetadata(key *string, value interface{})
	ToString() *string
}

// The jsii proxy struct for Artifact
type jsiiProxy_Artifact struct {
	_ byte // padding
}

func (j *jsiiProxy_Artifact) ArtifactName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Artifact) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Artifact) ObjectKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Artifact) S3Location() *awss3.Location {
	var returns *awss3.Location
	_jsii_.Get(
		j,
		"s3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Artifact) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


func NewArtifact(artifactName *string) Artifact {
	_init_.Initialize()

	j := jsiiProxy_Artifact{}

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.Artifact",
		[]interface{}{artifactName},
		&j,
	)

	return &j
}

func NewArtifact_Override(a Artifact, artifactName *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_codepipeline.Artifact",
		[]interface{}{artifactName},
		a,
	)
}

// A static factory method used to create instances of the Artifact class.
//
// Mainly meant to be used from `decdk`.
func Artifact_Artifact(name *string) Artifact {
	_init_.Initialize()

	if err := validateArtifact_ArtifactParameters(name); err != nil {
		panic(err)
	}
	var returns Artifact

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_codepipeline.Artifact",
		"artifact",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Artifact) AtPath(fileName *string) ArtifactPath {
	if err := a.validateAtPathParameters(fileName); err != nil {
		panic(err)
	}
	var returns ArtifactPath

	_jsii_.Invoke(
		a,
		"atPath",
		[]interface{}{fileName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Artifact) GetMetadata(key *string) interface{} {
	if err := a.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Artifact) GetParam(jsonFile *string, keyName *string) *string {
	if err := a.validateGetParamParameters(jsonFile, keyName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getParam",
		[]interface{}{jsonFile, keyName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_Artifact) SetMetadata(key *string, value interface{}) {
	if err := a.validateSetMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"setMetadata",
		[]interface{}{key, value},
	)
}

func (a *jsiiProxy_Artifact) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
