package awsmsk


// Properties for defining a `CfnCluster`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClusterProps := &CfnClusterProps{
//   	BrokerNodeGroupInfo: &BrokerNodeGroupInfoProperty{
//   		ClientSubnets: []*string{
//   			jsii.String("clientSubnets"),
//   		},
//   		InstanceType: jsii.String("instanceType"),
//
//   		// the properties below are optional
//   		BrokerAzDistribution: jsii.String("brokerAzDistribution"),
//   		ConnectivityInfo: &ConnectivityInfoProperty{
//   			PublicAccess: &PublicAccessProperty{
//   				Type: jsii.String("type"),
//   			},
//   			VpcConnectivity: &VpcConnectivityProperty{
//   				ClientAuthentication: &VpcConnectivityClientAuthenticationProperty{
//   					Sasl: &VpcConnectivitySaslProperty{
//   						Iam: &VpcConnectivityIamProperty{
//   							Enabled: jsii.Boolean(false),
//   						},
//   						Scram: &VpcConnectivityScramProperty{
//   							Enabled: jsii.Boolean(false),
//   						},
//   					},
//   					Tls: &VpcConnectivityTlsProperty{
//   						Enabled: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   		},
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		StorageInfo: &StorageInfoProperty{
//   			EbsStorageInfo: &EBSStorageInfoProperty{
//   				ProvisionedThroughput: &ProvisionedThroughputProperty{
//   					Enabled: jsii.Boolean(false),
//   					VolumeThroughput: jsii.Number(123),
//   				},
//   				VolumeSize: jsii.Number(123),
//   			},
//   		},
//   	},
//   	ClusterName: jsii.String("clusterName"),
//   	KafkaVersion: jsii.String("kafkaVersion"),
//   	NumberOfBrokerNodes: jsii.Number(123),
//
//   	// the properties below are optional
//   	ClientAuthentication: &ClientAuthenticationProperty{
//   		Sasl: &SaslProperty{
//   			Iam: &IamProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   			Scram: &ScramProperty{
//   				Enabled: jsii.Boolean(false),
//   			},
//   		},
//   		Tls: &TlsProperty{
//   			CertificateAuthorityArnList: []*string{
//   				jsii.String("certificateAuthorityArnList"),
//   			},
//   			Enabled: jsii.Boolean(false),
//   		},
//   		Unauthenticated: &UnauthenticatedProperty{
//   			Enabled: jsii.Boolean(false),
//   		},
//   	},
//   	ConfigurationInfo: &ConfigurationInfoProperty{
//   		Arn: jsii.String("arn"),
//   		Revision: jsii.Number(123),
//   	},
//   	CurrentVersion: jsii.String("currentVersion"),
//   	EncryptionInfo: &EncryptionInfoProperty{
//   		EncryptionAtRest: &EncryptionAtRestProperty{
//   			DataVolumeKmsKeyId: jsii.String("dataVolumeKmsKeyId"),
//   		},
//   		EncryptionInTransit: &EncryptionInTransitProperty{
//   			ClientBroker: jsii.String("clientBroker"),
//   			InCluster: jsii.Boolean(false),
//   		},
//   	},
//   	EnhancedMonitoring: jsii.String("enhancedMonitoring"),
//   	LoggingInfo: &LoggingInfoProperty{
//   		BrokerLogs: &BrokerLogsProperty{
//   			CloudWatchLogs: &CloudWatchLogsProperty{
//   				Enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				LogGroup: jsii.String("logGroup"),
//   			},
//   			Firehose: &FirehoseProperty{
//   				Enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				DeliveryStream: jsii.String("deliveryStream"),
//   			},
//   			S3: &S3Property{
//   				Enabled: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				Bucket: jsii.String("bucket"),
//   				Prefix: jsii.String("prefix"),
//   			},
//   		},
//   	},
//   	OpenMonitoring: &OpenMonitoringProperty{
//   		Prometheus: &PrometheusProperty{
//   			JmxExporter: &JmxExporterProperty{
//   				EnabledInBroker: jsii.Boolean(false),
//   			},
//   			NodeExporter: &NodeExporterProperty{
//   				EnabledInBroker: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	StorageMode: jsii.String("storageMode"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnClusterProps struct {
	// The setup to be used for brokers in the cluster.
	//
	// AWS CloudFormation may replace the cluster when you update certain `BrokerNodeGroupInfo` properties. To understand the update behavior for your use case, you should review the child properties for [`BrokerNodeGroupInfo`](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-msk-cluster-brokernodegroupinfo.html#aws-properties-msk-cluster-brokernodegroupinfo-properties) .
	BrokerNodeGroupInfo interface{} `field:"required" json:"brokerNodeGroupInfo" yaml:"brokerNodeGroupInfo"`
	// The name of the cluster.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// The version of Apache Kafka.
	//
	// For more information, see [Supported Apache Kafka versions](https://docs.aws.amazon.com/msk/latest/developerguide/supported-kafka-versions.html) in the Amazon MSK Developer Guide.
	KafkaVersion *string `field:"required" json:"kafkaVersion" yaml:"kafkaVersion"`
	// The number of broker nodes you want in the Amazon MSK cluster.
	//
	// You can submit an update to increase the number of broker nodes in a cluster.
	NumberOfBrokerNodes *float64 `field:"required" json:"numberOfBrokerNodes" yaml:"numberOfBrokerNodes"`
	// Includes information related to client authentication.
	ClientAuthentication interface{} `field:"optional" json:"clientAuthentication" yaml:"clientAuthentication"`
	// The Amazon MSK configuration to use for the cluster.
	ConfigurationInfo interface{} `field:"optional" json:"configurationInfo" yaml:"configurationInfo"`
	// The version of the cluster that you want to update.
	CurrentVersion *string `field:"optional" json:"currentVersion" yaml:"currentVersion"`
	// Includes all encryption-related information.
	EncryptionInfo interface{} `field:"optional" json:"encryptionInfo" yaml:"encryptionInfo"`
	// Specifies the level of monitoring for the MSK cluster.
	//
	// The possible values are `DEFAULT` , `PER_BROKER` , and `PER_TOPIC_PER_BROKER` .
	EnhancedMonitoring *string `field:"optional" json:"enhancedMonitoring" yaml:"enhancedMonitoring"`
	// You can configure your Amazon MSK cluster to send broker logs to different destination types.
	//
	// This is a container for the configuration details related to broker logs.
	LoggingInfo interface{} `field:"optional" json:"loggingInfo" yaml:"loggingInfo"`
	// The settings for open monitoring.
	OpenMonitoring interface{} `field:"optional" json:"openMonitoring" yaml:"openMonitoring"`
	// This controls storage mode for supported storage tiers.
	StorageMode *string `field:"optional" json:"storageMode" yaml:"storageMode"`
	// A map of key:value pairs to apply to this resource.
	//
	// Both key and value are of type String.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}
