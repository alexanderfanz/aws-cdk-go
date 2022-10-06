package awsec2


// Options when executing a file.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var instance instance
//
//
//   asset := awscdk.NewAsset(this, jsii.String("Asset"), &assetProps{
//   	path: jsii.String("./configure.sh"),
//   })
//
//   localPath := instance.userData.addS3DownloadCommand(&s3DownloadOptions{
//   	bucket: asset.bucket,
//   	bucketKey: asset.s3ObjectKey,
//   	region: jsii.String("us-east-1"),
//   })
//   instance.userData.addExecuteFileCommand(&executeFileOptions{
//   	filePath: localPath,
//   	arguments: jsii.String("--verbose -y"),
//   })
//   asset.grantRead(instance.role)
//
// Experimental.
type ExecuteFileOptions struct {
	// The path to the file.
	// Experimental.
	FilePath *string `field:"required" json:"filePath" yaml:"filePath"`
	// The arguments to be passed to the file.
	// Experimental.
	Arguments *string `field:"optional" json:"arguments" yaml:"arguments"`
}

