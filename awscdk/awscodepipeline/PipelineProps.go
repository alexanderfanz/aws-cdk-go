package awscodepipeline

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Example:
//   var project pipelineProject
//
//   repository := codecommit.NewRepository(this, jsii.String("MyRepository"), &RepositoryProps{
//   	RepositoryName: jsii.String("MyRepository"),
//   })
//   project := codebuild.NewPipelineProject(this, jsii.String("MyProject"))
//
//   sourceOutput := codepipeline.NewArtifact()
//   sourceAction := codepipeline_actions.NewCodeCommitSourceAction(&CodeCommitSourceActionProps{
//   	ActionName: jsii.String("CodeCommit"),
//   	Repository: Repository,
//   	Output: sourceOutput,
//   })
//   buildAction := codepipeline_actions.NewCodeBuildAction(&CodeBuildActionProps{
//   	ActionName: jsii.String("CodeBuild"),
//   	Project: Project,
//   	Input: sourceOutput,
//   	Outputs: []artifact{
//   		codepipeline.NewArtifact(),
//   	},
//   	 // optional
//   	ExecuteBatchBuild: jsii.Boolean(true),
//   	 // optional, defaults to false
//   	CombineBatchBuildArtifacts: jsii.Boolean(true),
//   })
//
//   codepipeline.NewPipeline(this, jsii.String("MyPipeline"), &PipelineProps{
//   	Stages: []stageProps{
//   		&stageProps{
//   			StageName: jsii.String("Source"),
//   			Actions: []iAction{
//   				sourceAction,
//   			},
//   		},
//   		&stageProps{
//   			StageName: jsii.String("Build"),
//   			Actions: []*iAction{
//   				buildAction,
//   			},
//   		},
//   	},
//   })
//
type PipelineProps struct {
	// The S3 bucket used by this Pipeline to store artifacts.
	// Default: - A new S3 bucket will be created.
	//
	ArtifactBucket awss3.IBucket `field:"optional" json:"artifactBucket" yaml:"artifactBucket"`
	// Create KMS keys for cross-account deployments.
	//
	// This controls whether the pipeline is enabled for cross-account deployments.
	//
	// By default cross-account deployments are enabled, but this feature requires
	// that KMS Customer Master Keys are created which have a cost of $1/month.
	//
	// If you do not need cross-account deployments, you can set this to `false` to
	// not create those keys and save on that cost (the artifact bucket will be
	// encrypted with an AWS-managed key). However, cross-account deployments will
	// no longer be possible.
	// Default: true.
	//
	CrossAccountKeys *bool `field:"optional" json:"crossAccountKeys" yaml:"crossAccountKeys"`
	// A map of region to S3 bucket name used for cross-region CodePipeline.
	//
	// For every Action that you specify targeting a different region than the Pipeline itself,
	// if you don't provide an explicit Bucket for that region using this property,
	// the construct will automatically create a Stack containing an S3 Bucket in that region.
	// Default: - None.
	//
	CrossRegionReplicationBuckets *map[string]awss3.IBucket `field:"optional" json:"crossRegionReplicationBuckets" yaml:"crossRegionReplicationBuckets"`
	// Enable KMS key rotation for the generated KMS keys.
	//
	// By default KMS key rotation is disabled, but will add an additional $1/month
	// for each year the key exists when enabled.
	// Default: - false (key rotation is disabled).
	//
	EnableKeyRotation *bool `field:"optional" json:"enableKeyRotation" yaml:"enableKeyRotation"`
	// Name of the pipeline.
	// Default: - AWS CloudFormation generates an ID and uses that for the pipeline name.
	//
	PipelineName *string `field:"optional" json:"pipelineName" yaml:"pipelineName"`
	// Indicates whether to rerun the AWS CodePipeline pipeline after you update it.
	// Default: false.
	//
	RestartExecutionOnUpdate *bool `field:"optional" json:"restartExecutionOnUpdate" yaml:"restartExecutionOnUpdate"`
	// Reuse the same cross region support stack for all pipelines in the App.
	// Default: - true (Use the same support stack for all pipelines in App).
	//
	ReuseCrossRegionSupportStacks *bool `field:"optional" json:"reuseCrossRegionSupportStacks" yaml:"reuseCrossRegionSupportStacks"`
	// The IAM role to be assumed by this Pipeline.
	// Default: a new IAM role will be created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The list of Stages, in order, to create this Pipeline with.
	//
	// You can always add more Stages later by calling `Pipeline#addStage`.
	// Default: - None.
	//
	Stages *[]*StageProps `field:"optional" json:"stages" yaml:"stages"`
}

