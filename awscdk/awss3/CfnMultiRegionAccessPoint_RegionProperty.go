package awss3


// A bucket associated with a specific Region when creating Multi-Region Access Points.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   regionProperty := &RegionProperty{
//   	Bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	BucketAccountId: jsii.String("bucketAccountId"),
//   }
//
type CfnMultiRegionAccessPoint_RegionProperty struct {
	// The name of the associated bucket for the Region.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The AWS account ID that owns the Amazon S3 bucket that's associated with this Multi-Region Access Point.
	BucketAccountId *string `field:"optional" json:"bucketAccountId" yaml:"bucketAccountId"`
}

