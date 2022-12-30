package awss3


// The ObjectOwnership of the bucket.
//
// Example:
//   s3.NewBucket(this, jsii.String("MyBucket"), &bucketProps{
//   	objectOwnership: s3.objectOwnership_OBJECT_WRITER,
//   })
//
// See: https://docs.aws.amazon.com/AmazonS3/latest/dev/about-object-ownership.html
//
// Experimental.
type ObjectOwnership string

const (
	// ACLs are disabled, and the bucket owner automatically owns and has full control over every object in the bucket.
	//
	// ACLs no longer affect permissions to data in the S3 bucket.
	// The bucket uses policies to define access control.
	// Experimental.
	ObjectOwnership_BUCKET_OWNER_ENFORCED ObjectOwnership = "BUCKET_OWNER_ENFORCED"
	// Objects uploaded to the bucket change ownership to the bucket owner .
	// Experimental.
	ObjectOwnership_BUCKET_OWNER_PREFERRED ObjectOwnership = "BUCKET_OWNER_PREFERRED"
	// The uploading account will own the object.
	// Experimental.
	ObjectOwnership_OBJECT_WRITER ObjectOwnership = "OBJECT_WRITER"
)

