package awscdkbatchalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props for EksJobDefinition.
//
// Example:
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   jobDefn := batch.NewEksJobDefinition(this, jsii.String("eksf2"), &EksJobDefinitionProps{
//   	Container: batch.NewEksContainerDefinition(this, jsii.String("container"), &EksContainerDefinitionProps{
//   		Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   		Volumes: []eksVolume{
//   			batch.*eksVolume_EmptyDir(&EmptyDirVolumeOptions{
//   				Name: jsii.String("myEmptyDirVolume"),
//   				MountPath: jsii.String("/mount/path"),
//   				Medium: batch.EmptyDirMediumType_MEMORY,
//   				Readonly: jsii.Boolean(true),
//   				SizeLimit: cdk.Size_Mebibytes(jsii.Number(2048)),
//   			}),
//   		},
//   	}),
//   })
//
// Experimental.
type EksJobDefinitionProps struct {
	// The name of this job definition.
	// Experimental.
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// The default parameters passed to the container These parameters can be referenced in the `command` that you give to the container.
	// See: https://docs.aws.amazon.com/batch/latest/userguide/job_definition_parameters.html#parameters
	//
	// Experimental.
	Parameters *map[string]interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// The number of times to retry a job.
	//
	// The job is retried on failure the same number of attempts as the value.
	// Experimental.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// Defines the retry behavior for this job.
	// Experimental.
	RetryStrategies *[]RetryStrategy `field:"optional" json:"retryStrategies" yaml:"retryStrategies"`
	// The priority of this Job.
	//
	// Only used in Fairshare Scheduling
	// to decide which job to run first when there are multiple jobs
	// with the same share identifier.
	// Experimental.
	SchedulingPriority *float64 `field:"optional" json:"schedulingPriority" yaml:"schedulingPriority"`
	// The timeout time for jobs that are submitted with this job definition.
	//
	// After the amount of time you specify passes,
	// Batch terminates your jobs if they aren't finished.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The container this Job Definition will run.
	// Experimental.
	Container EksContainerDefinition `field:"required" json:"container" yaml:"container"`
	// The DNS Policy of the pod used by this Job Definition.
	// See: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy
	//
	// Experimental.
	DnsPolicy DnsPolicy `field:"optional" json:"dnsPolicy" yaml:"dnsPolicy"`
	// The name of the service account that's used to run the container.
	//
	// service accounts are Kubernetes method of identification and authentication,
	// roughly analogous to IAM users.
	// See: https://docs.aws.amazon.com/eks/latest/userguide/associate-service-account-role.html
	//
	// Experimental.
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// If specified, the Pod used by this Job Definition will use the host's network IP address.
	//
	// Otherwise, the Kubernetes pod networking model is enabled.
	// Most AWS Batch workloads are egress-only and don't require the overhead of IP allocation for each pod for incoming connections.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/#pod-networking
	//
	// Experimental.
	UseHostNetwork *bool `field:"optional" json:"useHostNetwork" yaml:"useHostNetwork"`
}

