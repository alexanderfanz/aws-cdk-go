package awseks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
)

// Helm Chart properties.
//
// Example:
//   var cluster cluster
//
//   // option 1: use a construct
//   // option 1: use a construct
//   eks.NewHelmChart(this, jsii.String("MyOCIChart"), &HelmChartProps{
//   	Cluster: Cluster,
//   	Chart: jsii.String("some-chart"),
//   	Repository: jsii.String("oci://${ACCOUNT_ID}.dkr.ecr.${ACCOUNT_REGION}.amazonaws.com/${REPO_NAME}"),
//   	Namespace: jsii.String("oci"),
//   	Version: jsii.String("0.0.1"),
//   })
//
type HelmChartProps struct {
	// The name of the chart.
	//
	// Either this or `chartAsset` must be specified.
	Chart *string `field:"optional" json:"chart" yaml:"chart"`
	// The chart in the form of an asset.
	//
	// Either this or `chart` must be specified.
	ChartAsset awss3assets.Asset `field:"optional" json:"chartAsset" yaml:"chartAsset"`
	// create namespace if not exist.
	CreateNamespace *bool `field:"optional" json:"createNamespace" yaml:"createNamespace"`
	// The Kubernetes namespace scope of the requests.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The name of the release.
	Release *string `field:"optional" json:"release" yaml:"release"`
	// The repository which contains the chart.
	//
	// For example: https://charts.helm.sh/stable/
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
	// if set, no CRDs will be installed.
	SkipCrds *bool `field:"optional" json:"skipCrds" yaml:"skipCrds"`
	// Amount of time to wait for any individual Kubernetes operation.
	//
	// Maximum 15 minutes.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The values to be used by the chart.
	//
	// For nested values use a nested dictionary. For example:
	// values: {
	//  installationCRDs: true,
	//  webhook: { port: 9443 }
	// }.
	Values *map[string]interface{} `field:"optional" json:"values" yaml:"values"`
	// The chart version to install.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Whether or not Helm should wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful.
	Wait *bool `field:"optional" json:"wait" yaml:"wait"`
	// The EKS cluster to apply this configuration to.
	//
	// [disable-awslint:ref-via-interface].
	Cluster ICluster `field:"required" json:"cluster" yaml:"cluster"`
}
