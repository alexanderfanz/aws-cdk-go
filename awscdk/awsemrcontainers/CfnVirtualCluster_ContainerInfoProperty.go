package awsemrcontainers


// The information about the container used for a job run or a managed endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerInfoProperty := &ContainerInfoProperty{
//   	EksInfo: &EksInfoProperty{
//   		Namespace: jsii.String("namespace"),
//   	},
//   }
//
type CfnVirtualCluster_ContainerInfoProperty struct {
	// The information about the Amazon EKS cluster.
	EksInfo interface{} `field:"required" json:"eksInfo" yaml:"eksInfo"`
}

