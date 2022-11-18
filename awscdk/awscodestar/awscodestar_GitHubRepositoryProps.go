package awscodestar

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
)

// Construction properties of {@link GitHubRepository}.
//
// Example:
//   import codestar "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//
//   codestar.NewGitHubRepository(this, jsii.String("GitHubRepo"), &gitHubRepositoryProps{
//   	owner: jsii.String("aws"),
//   	repositoryName: jsii.String("aws-cdk"),
//   	accessToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token"), &secretsManagerSecretOptions{
//   		jsonField: jsii.String("token"),
//   	}),
//   	contentsBucket: s3.bucket.fromBucketName(this, jsii.String("Bucket"), jsii.String("bucket-name")),
//   	contentsKey: jsii.String("import.zip"),
//   })
//
// Experimental.
type GitHubRepositoryProps struct {
	// The GitHub user's personal access token for the GitHub repository.
	// Experimental.
	AccessToken awscdk.SecretValue `field:"required" json:"accessToken" yaml:"accessToken"`
	// The name of the Amazon S3 bucket that contains the ZIP file with the content to be committed to the new repository.
	// Experimental.
	ContentsBucket awss3.IBucket `field:"required" json:"contentsBucket" yaml:"contentsBucket"`
	// The S3 object key or file name for the ZIP file.
	// Experimental.
	ContentsKey *string `field:"required" json:"contentsKey" yaml:"contentsKey"`
	// The GitHub user name for the owner of the GitHub repository to be created.
	//
	// If this
	// repository should be owned by a GitHub organization, provide its name.
	// Experimental.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The name of the repository you want to create in GitHub with AWS CloudFormation stack creation.
	// Experimental.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// The object version of the ZIP file, if versioning is enabled for the Amazon S3 bucket.
	// Experimental.
	ContentsS3Version *string `field:"optional" json:"contentsS3Version" yaml:"contentsS3Version"`
	// A comment or description about the new repository.
	//
	// This description is displayed in GitHub after the repository
	// is created.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether to enable issues for the GitHub repository.
	//
	// You can use GitHub issues to track information
	// and bugs for your repository.
	// Experimental.
	EnableIssues *bool `field:"optional" json:"enableIssues" yaml:"enableIssues"`
	// Indicates whether the GitHub repository is a private repository.
	//
	// If so, you choose who can see and commit to
	// this repository.
	// Experimental.
	Visibility RepositoryVisibility `field:"optional" json:"visibility" yaml:"visibility"`
}

