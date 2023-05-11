package awskinesisfirehose

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::KinesisFirehose::DeliveryStream`.
//
// The `AWS::KinesisFirehose::DeliveryStream` resource specifies an Amazon Kinesis Data Firehose (Kinesis Data Firehose) delivery stream that delivers real-time streaming data to an Amazon Simple Storage Service (Amazon S3), Amazon Redshift, or Amazon Elasticsearch Service (Amazon ES) destination. For more information, see [Creating an Amazon Kinesis Data Firehose Delivery Stream](https://docs.aws.amazon.com/firehose/latest/dev/basic-create.html) in the *Amazon Kinesis Data Firehose Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeliveryStream := awscdk.Aws_kinesisfirehose.NewCfnDeliveryStream(this, jsii.String("MyCfnDeliveryStream"), &CfnDeliveryStreamProps{
//   	AmazonOpenSearchServerlessDestinationConfiguration: &AmazonOpenSearchServerlessDestinationConfigurationProperty{
//   		IndexName: jsii.String("indexName"),
//   		RoleArn: jsii.String("roleArn"),
//   		S3Configuration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		BufferingHints: &AmazonOpenSearchServerlessBufferingHintsProperty{
//   			IntervalInSeconds: jsii.Number(123),
//   			SizeInMBs: jsii.Number(123),
//   		},
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		CollectionEndpoint: jsii.String("collectionEndpoint"),
//   		ProcessingConfiguration: &ProcessingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			Processors: []interface{}{
//   				&ProcessorProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Parameters: []interface{}{
//   						&ProcessorParameterProperty{
//   							ParameterName: jsii.String("parameterName"),
//   							ParameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		RetryOptions: &AmazonOpenSearchServerlessRetryOptionsProperty{
//   			DurationInSeconds: jsii.Number(123),
//   		},
//   		S3BackupMode: jsii.String("s3BackupMode"),
//   		VpcConfiguration: &VpcConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	AmazonopensearchserviceDestinationConfiguration: &AmazonopensearchserviceDestinationConfigurationProperty{
//   		IndexName: jsii.String("indexName"),
//   		RoleArn: jsii.String("roleArn"),
//   		S3Configuration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		BufferingHints: &AmazonopensearchserviceBufferingHintsProperty{
//   			IntervalInSeconds: jsii.Number(123),
//   			SizeInMBs: jsii.Number(123),
//   		},
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		ClusterEndpoint: jsii.String("clusterEndpoint"),
//   		DocumentIdOptions: &DocumentIdOptionsProperty{
//   			DefaultDocumentIdFormat: jsii.String("defaultDocumentIdFormat"),
//   		},
//   		DomainArn: jsii.String("domainArn"),
//   		IndexRotationPeriod: jsii.String("indexRotationPeriod"),
//   		ProcessingConfiguration: &ProcessingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			Processors: []interface{}{
//   				&ProcessorProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Parameters: []interface{}{
//   						&ProcessorParameterProperty{
//   							ParameterName: jsii.String("parameterName"),
//   							ParameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		RetryOptions: &AmazonopensearchserviceRetryOptionsProperty{
//   			DurationInSeconds: jsii.Number(123),
//   		},
//   		S3BackupMode: jsii.String("s3BackupMode"),
//   		TypeName: jsii.String("typeName"),
//   		VpcConfiguration: &VpcConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	DeliveryStreamEncryptionConfigurationInput: &DeliveryStreamEncryptionConfigurationInputProperty{
//   		KeyType: jsii.String("keyType"),
//
//   		// the properties below are optional
//   		KeyArn: jsii.String("keyArn"),
//   	},
//   	DeliveryStreamName: jsii.String("deliveryStreamName"),
//   	DeliveryStreamType: jsii.String("deliveryStreamType"),
//   	ElasticsearchDestinationConfiguration: &ElasticsearchDestinationConfigurationProperty{
//   		IndexName: jsii.String("indexName"),
//   		RoleArn: jsii.String("roleArn"),
//   		S3Configuration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		BufferingHints: &ElasticsearchBufferingHintsProperty{
//   			IntervalInSeconds: jsii.Number(123),
//   			SizeInMBs: jsii.Number(123),
//   		},
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		ClusterEndpoint: jsii.String("clusterEndpoint"),
//   		DocumentIdOptions: &DocumentIdOptionsProperty{
//   			DefaultDocumentIdFormat: jsii.String("defaultDocumentIdFormat"),
//   		},
//   		DomainArn: jsii.String("domainArn"),
//   		IndexRotationPeriod: jsii.String("indexRotationPeriod"),
//   		ProcessingConfiguration: &ProcessingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			Processors: []interface{}{
//   				&ProcessorProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Parameters: []interface{}{
//   						&ProcessorParameterProperty{
//   							ParameterName: jsii.String("parameterName"),
//   							ParameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		RetryOptions: &ElasticsearchRetryOptionsProperty{
//   			DurationInSeconds: jsii.Number(123),
//   		},
//   		S3BackupMode: jsii.String("s3BackupMode"),
//   		TypeName: jsii.String("typeName"),
//   		VpcConfiguration: &VpcConfigurationProperty{
//   			RoleArn: jsii.String("roleArn"),
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	ExtendedS3DestinationConfiguration: &ExtendedS3DestinationConfigurationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		BufferingHints: &BufferingHintsProperty{
//   			IntervalInSeconds: jsii.Number(123),
//   			SizeInMBs: jsii.Number(123),
//   		},
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		CompressionFormat: jsii.String("compressionFormat"),
//   		DataFormatConversionConfiguration: &DataFormatConversionConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			InputFormatConfiguration: &InputFormatConfigurationProperty{
//   				Deserializer: &DeserializerProperty{
//   					HiveJsonSerDe: &HiveJsonSerDeProperty{
//   						TimestampFormats: []*string{
//   							jsii.String("timestampFormats"),
//   						},
//   					},
//   					OpenXJsonSerDe: &OpenXJsonSerDeProperty{
//   						CaseInsensitive: jsii.Boolean(false),
//   						ColumnToJsonKeyMappings: map[string]*string{
//   							"columnToJsonKeyMappingsKey": jsii.String("columnToJsonKeyMappings"),
//   						},
//   						ConvertDotsInJsonKeysToUnderscores: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   			OutputFormatConfiguration: &OutputFormatConfigurationProperty{
//   				Serializer: &SerializerProperty{
//   					OrcSerDe: &OrcSerDeProperty{
//   						BlockSizeBytes: jsii.Number(123),
//   						BloomFilterColumns: []*string{
//   							jsii.String("bloomFilterColumns"),
//   						},
//   						BloomFilterFalsePositiveProbability: jsii.Number(123),
//   						Compression: jsii.String("compression"),
//   						DictionaryKeyThreshold: jsii.Number(123),
//   						EnablePadding: jsii.Boolean(false),
//   						FormatVersion: jsii.String("formatVersion"),
//   						PaddingTolerance: jsii.Number(123),
//   						RowIndexStride: jsii.Number(123),
//   						StripeSizeBytes: jsii.Number(123),
//   					},
//   					ParquetSerDe: &ParquetSerDeProperty{
//   						BlockSizeBytes: jsii.Number(123),
//   						Compression: jsii.String("compression"),
//   						EnableDictionaryCompression: jsii.Boolean(false),
//   						MaxPaddingBytes: jsii.Number(123),
//   						PageSizeBytes: jsii.Number(123),
//   						WriterVersion: jsii.String("writerVersion"),
//   					},
//   				},
//   			},
//   			SchemaConfiguration: &SchemaConfigurationProperty{
//   				CatalogId: jsii.String("catalogId"),
//   				DatabaseName: jsii.String("databaseName"),
//   				Region: jsii.String("region"),
//   				RoleArn: jsii.String("roleArn"),
//   				TableName: jsii.String("tableName"),
//   				VersionId: jsii.String("versionId"),
//   			},
//   		},
//   		DynamicPartitioningConfiguration: &DynamicPartitioningConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			RetryOptions: &RetryOptionsProperty{
//   				DurationInSeconds: jsii.Number(123),
//   			},
//   		},
//   		EncryptionConfiguration: &EncryptionConfigurationProperty{
//   			KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   				AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   			},
//   			NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   		},
//   		ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		Prefix: jsii.String("prefix"),
//   		ProcessingConfiguration: &ProcessingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			Processors: []interface{}{
//   				&ProcessorProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Parameters: []interface{}{
//   						&ProcessorParameterProperty{
//   							ParameterName: jsii.String("parameterName"),
//   							ParameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		S3BackupConfiguration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		S3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   	HttpEndpointDestinationConfiguration: &HttpEndpointDestinationConfigurationProperty{
//   		EndpointConfiguration: &HttpEndpointConfigurationProperty{
//   			Url: jsii.String("url"),
//
//   			// the properties below are optional
//   			AccessKey: jsii.String("accessKey"),
//   			Name: jsii.String("name"),
//   		},
//   		S3Configuration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		BufferingHints: &BufferingHintsProperty{
//   			IntervalInSeconds: jsii.Number(123),
//   			SizeInMBs: jsii.Number(123),
//   		},
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		ProcessingConfiguration: &ProcessingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			Processors: []interface{}{
//   				&ProcessorProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Parameters: []interface{}{
//   						&ProcessorParameterProperty{
//   							ParameterName: jsii.String("parameterName"),
//   							ParameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		RequestConfiguration: &HttpEndpointRequestConfigurationProperty{
//   			CommonAttributes: []interface{}{
//   				&HttpEndpointCommonAttributeProperty{
//   					AttributeName: jsii.String("attributeName"),
//   					AttributeValue: jsii.String("attributeValue"),
//   				},
//   			},
//   			ContentEncoding: jsii.String("contentEncoding"),
//   		},
//   		RetryOptions: &RetryOptionsProperty{
//   			DurationInSeconds: jsii.Number(123),
//   		},
//   		RoleArn: jsii.String("roleArn"),
//   		S3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   	KinesisStreamSourceConfiguration: &KinesisStreamSourceConfigurationProperty{
//   		KinesisStreamArn: jsii.String("kinesisStreamArn"),
//   		RoleArn: jsii.String("roleArn"),
//   	},
//   	RedshiftDestinationConfiguration: &RedshiftDestinationConfigurationProperty{
//   		ClusterJdbcurl: jsii.String("clusterJdbcurl"),
//   		CopyCommand: &CopyCommandProperty{
//   			DataTableName: jsii.String("dataTableName"),
//
//   			// the properties below are optional
//   			CopyOptions: jsii.String("copyOptions"),
//   			DataTableColumns: jsii.String("dataTableColumns"),
//   		},
//   		Password: jsii.String("password"),
//   		RoleArn: jsii.String("roleArn"),
//   		S3Configuration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		Username: jsii.String("username"),
//
//   		// the properties below are optional
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		ProcessingConfiguration: &ProcessingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			Processors: []interface{}{
//   				&ProcessorProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Parameters: []interface{}{
//   						&ProcessorParameterProperty{
//   							ParameterName: jsii.String("parameterName"),
//   							ParameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		RetryOptions: &RedshiftRetryOptionsProperty{
//   			DurationInSeconds: jsii.Number(123),
//   		},
//   		S3BackupConfiguration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		S3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   	S3DestinationConfiguration: &S3DestinationConfigurationProperty{
//   		BucketArn: jsii.String("bucketArn"),
//   		RoleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		BufferingHints: &BufferingHintsProperty{
//   			IntervalInSeconds: jsii.Number(123),
//   			SizeInMBs: jsii.Number(123),
//   		},
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		CompressionFormat: jsii.String("compressionFormat"),
//   		EncryptionConfiguration: &EncryptionConfigurationProperty{
//   			KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   				AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   			},
//   			NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   		},
//   		ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   		Prefix: jsii.String("prefix"),
//   	},
//   	SplunkDestinationConfiguration: &SplunkDestinationConfigurationProperty{
//   		HecEndpoint: jsii.String("hecEndpoint"),
//   		HecEndpointType: jsii.String("hecEndpointType"),
//   		HecToken: jsii.String("hecToken"),
//   		S3Configuration: &S3DestinationConfigurationProperty{
//   			BucketArn: jsii.String("bucketArn"),
//   			RoleArn: jsii.String("roleArn"),
//
//   			// the properties below are optional
//   			BufferingHints: &BufferingHintsProperty{
//   				IntervalInSeconds: jsii.Number(123),
//   				SizeInMBs: jsii.Number(123),
//   			},
//   			CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   				Enabled: jsii.Boolean(false),
//   				LogGroupName: jsii.String("logGroupName"),
//   				LogStreamName: jsii.String("logStreamName"),
//   			},
//   			CompressionFormat: jsii.String("compressionFormat"),
//   			EncryptionConfiguration: &EncryptionConfigurationProperty{
//   				KmsEncryptionConfig: &KMSEncryptionConfigProperty{
//   					AwskmsKeyArn: jsii.String("awskmsKeyArn"),
//   				},
//   				NoEncryptionConfig: jsii.String("noEncryptionConfig"),
//   			},
//   			ErrorOutputPrefix: jsii.String("errorOutputPrefix"),
//   			Prefix: jsii.String("prefix"),
//   		},
//
//   		// the properties below are optional
//   		CloudWatchLoggingOptions: &CloudWatchLoggingOptionsProperty{
//   			Enabled: jsii.Boolean(false),
//   			LogGroupName: jsii.String("logGroupName"),
//   			LogStreamName: jsii.String("logStreamName"),
//   		},
//   		HecAcknowledgmentTimeoutInSeconds: jsii.Number(123),
//   		ProcessingConfiguration: &ProcessingConfigurationProperty{
//   			Enabled: jsii.Boolean(false),
//   			Processors: []interface{}{
//   				&ProcessorProperty{
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Parameters: []interface{}{
//   						&ProcessorParameterProperty{
//   							ParameterName: jsii.String("parameterName"),
//   							ParameterValue: jsii.String("parameterValue"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		RetryOptions: &SplunkRetryOptionsProperty{
//   			DurationInSeconds: jsii.Number(123),
//   		},
//   		S3BackupMode: jsii.String("s3BackupMode"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnDeliveryStream interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::KinesisFirehose::DeliveryStream.AmazonOpenSearchServerlessDestinationConfiguration`.
	AmazonOpenSearchServerlessDestinationConfiguration() interface{}
	SetAmazonOpenSearchServerlessDestinationConfiguration(val interface{})
	// The destination in Amazon OpenSearch Service.
	//
	// You can specify only one destination.
	AmazonopensearchserviceDestinationConfiguration() interface{}
	SetAmazonopensearchserviceDestinationConfiguration(val interface{})
	// The Amazon Resource Name (ARN) of the delivery stream, such as `arn:aws:firehose:us-east-2:123456789012:deliverystream/delivery-stream-name` .
	AttrArn() *string
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// Specifies the type and Amazon Resource Name (ARN) of the CMK to use for Server-Side Encryption (SSE).
	DeliveryStreamEncryptionConfigurationInput() interface{}
	SetDeliveryStreamEncryptionConfigurationInput(val interface{})
	// The name of the delivery stream.
	DeliveryStreamName() *string
	SetDeliveryStreamName(val *string)
	// The delivery stream type. This can be one of the following values:.
	//
	// - `DirectPut` : Provider applications access the delivery stream directly.
	// - `KinesisStreamAsSource` : The delivery stream uses a Kinesis data stream as a source.
	DeliveryStreamType() *string
	SetDeliveryStreamType(val *string)
	// An Amazon ES destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon ES destination to an Amazon S3 or Amazon Redshift destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	ElasticsearchDestinationConfiguration() interface{}
	SetElasticsearchDestinationConfiguration(val interface{})
	// An Amazon S3 destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon Extended S3 destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	ExtendedS3DestinationConfiguration() interface{}
	SetExtendedS3DestinationConfiguration(val interface{})
	// Enables configuring Kinesis Firehose to deliver data to any HTTP endpoint destination.
	//
	// You can specify only one destination.
	HttpEndpointDestinationConfiguration() interface{}
	SetHttpEndpointDestinationConfiguration(val interface{})
	// When a Kinesis stream is used as the source for the delivery stream, a [KinesisStreamSourceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisfirehose-deliverystream-kinesisstreamsourceconfiguration.html) containing the Kinesis stream ARN and the role ARN for the source stream.
	KinesisStreamSourceConfiguration() interface{}
	SetKinesisStreamSourceConfiguration(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	LogicalId() *string
	// The tree node.
	Node() constructs.Node
	// An Amazon Redshift destination for the delivery stream.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon Redshift destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	RedshiftDestinationConfiguration() interface{}
	SetRedshiftDestinationConfiguration(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The `S3DestinationConfiguration` property type specifies an Amazon Simple Storage Service (Amazon S3) destination to which Amazon Kinesis Data Firehose (Kinesis Data Firehose) delivers data.
	//
	// Conditional. You must specify only one destination configuration.
	//
	// If you change the delivery stream destination from an Amazon S3 destination to an Amazon ES destination, update requires [some interruptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-some-interrupt) .
	S3DestinationConfiguration() interface{}
	SetS3DestinationConfiguration(val interface{})
	// The configuration of a destination in Splunk for the delivery stream.
	SplunkDestinationConfiguration() interface{}
	SetSplunkDestinationConfiguration(val interface{})
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// A set of tags to assign to the delivery stream.
	//
	// A tag is a key-value pair that you can define and assign to AWS resources. Tags are metadata. For example, you can add friendly names and descriptions or other types of information that can help you distinguish the delivery stream. For more information about tags, see [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html) in the AWS Billing and Cost Management User Guide.
	//
	// You can specify up to 50 tags when creating a delivery stream.
	Tags() awscdk.TagManager
	// Deprecated.
	// Deprecated: use `updatedProperties`
	//
	// Return properties modified after initiation
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperites() *map[string]interface{}
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	UpdatedProperties() *map[string]interface{}
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	AddDependency(target awscdk.CfnResource)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	// Deprecated: use addDependency.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	AddMetadata(key *string, value interface{})
	// Adds an override to the synthesized CloudFormation resource.
	//
	// To add a
	// property override, either use `addPropertyOverride` or prefix `path` with
	// "Properties." (i.e. `Properties.TopicName`).
	//
	// If the override is nested, separate each nested level using a dot (.) in the path parameter.
	// If there is an array as part of the nesting, specify the index in the path.
	//
	// To include a literal `.` in the property name, prefix with a `\`. In most
	// programming languages you will need to write this as `"\\."` because the
	// `\` itself will need to be escaped.
	//
	// For example,
	// ```typescript
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
	// cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
	// ```
	// would add the overrides
	// ```json
	// "Properties": {
	//   "GlobalSecondaryIndexes": [
	//     {
	//       "Projection": {
	//         "NonKeyAttributes": [ "myattribute" ]
	//         ...
	//       }
	//       ...
	//     },
	//     {
	//       "ProjectionType": "INCLUDE"
	//       ...
	//     },
	//   ]
	//   ...
	// }
	// ```
	//
	// The `value` argument to `addOverride` will not be processed or translated
	// in any way. Pass raw JSON values in here with the correct capitalization
	// for CloudFormation. If you pass CDK classes or structs, they will be
	// rendered with lowercased key names, and CloudFormation will reject the
	// template.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
	// cases, a snapshot can be taken of the resource prior to deletion
	// (`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
	// can be found in the following link:.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options
	//
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Retrieves an array of resources this resource depends on.
	//
	// This assembles dependencies on resources across stacks (including nested stacks)
	// automatically.
	ObtainDependencies() *[]interface{}
	// Get a shallow copy of dependencies between this resource and other resources in the same stack.
	ObtainResourceDependencies() *[]awscdk.CfnResource
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	// Indicates that this resource no longer depends on another resource.
	//
	// This can be used for resources across stacks (including nested stacks)
	// and the dependency will automatically be removed from the relevant scope.
	RemoveDependency(target awscdk.CfnResource)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Replaces one dependency with another.
	ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource)
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	ShouldSynthesize() *bool
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	ToString() *string
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnDeliveryStream
type jsiiProxy_CfnDeliveryStream struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnDeliveryStream) AmazonOpenSearchServerlessDestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonOpenSearchServerlessDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) AmazonopensearchserviceDestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonopensearchserviceDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) DeliveryStreamEncryptionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliveryStreamEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) DeliveryStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) DeliveryStreamType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) ElasticsearchDestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) ExtendedS3DestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extendedS3DestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) HttpEndpointDestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpEndpointDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) KinesisStreamSourceConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisStreamSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) RedshiftDestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) S3DestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3DestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) SplunkDestinationConfiguration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splunkDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeliveryStream) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}


// Create a new `AWS::KinesisFirehose::DeliveryStream`.
func NewCfnDeliveryStream(scope constructs.Construct, id *string, props *CfnDeliveryStreamProps) CfnDeliveryStream {
	_init_.Initialize()

	if err := validateNewCfnDeliveryStreamParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDeliveryStream{}

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::KinesisFirehose::DeliveryStream`.
func NewCfnDeliveryStream_Override(c CfnDeliveryStream, scope constructs.Construct, id *string, props *CfnDeliveryStreamProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnDeliveryStream)SetAmazonOpenSearchServerlessDestinationConfiguration(val interface{}) {
	if err := j.validateSetAmazonOpenSearchServerlessDestinationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amazonOpenSearchServerlessDestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream)SetAmazonopensearchserviceDestinationConfiguration(val interface{}) {
	if err := j.validateSetAmazonopensearchserviceDestinationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amazonopensearchserviceDestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream)SetDeliveryStreamEncryptionConfigurationInput(val interface{}) {
	if err := j.validateSetDeliveryStreamEncryptionConfigurationInputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliveryStreamEncryptionConfigurationInput",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream)SetDeliveryStreamName(val *string) {
	_jsii_.Set(
		j,
		"deliveryStreamName",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream)SetDeliveryStreamType(val *string) {
	_jsii_.Set(
		j,
		"deliveryStreamType",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream)SetElasticsearchDestinationConfiguration(val interface{}) {
	if err := j.validateSetElasticsearchDestinationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticsearchDestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream)SetExtendedS3DestinationConfiguration(val interface{}) {
	if err := j.validateSetExtendedS3DestinationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extendedS3DestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream)SetHttpEndpointDestinationConfiguration(val interface{}) {
	if err := j.validateSetHttpEndpointDestinationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpEndpointDestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream)SetKinesisStreamSourceConfiguration(val interface{}) {
	if err := j.validateSetKinesisStreamSourceConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kinesisStreamSourceConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream)SetRedshiftDestinationConfiguration(val interface{}) {
	if err := j.validateSetRedshiftDestinationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redshiftDestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream)SetS3DestinationConfiguration(val interface{}) {
	if err := j.validateSetS3DestinationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3DestinationConfiguration",
		val,
	)
}

func (j *jsiiProxy_CfnDeliveryStream)SetSplunkDestinationConfiguration(val interface{}) {
	if err := j.validateSetSplunkDestinationConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"splunkDestinationConfiguration",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnDeliveryStream_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDeliveryStream_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnDeliveryStream_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnDeliveryStream_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func CfnDeliveryStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDeliveryStream_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeliveryStream_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_kinesisfirehose.CfnDeliveryStream",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDeliveryStream) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) AddDependency(target awscdk.CfnResource) {
	if err := c.validateAddDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) GetAtt(attributeName *string, typeHint awscdk.ResolutionTypeHint) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName, typeHint},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeliveryStream) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeliveryStream) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) ObtainDependencies() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"obtainDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeliveryStream) ObtainResourceDependencies() *[]awscdk.CfnResource {
	var returns *[]awscdk.CfnResource

	_jsii_.Invoke(
		c,
		"obtainResourceDependencies",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeliveryStream) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) RemoveDependency(target awscdk.CfnResource) {
	if err := c.validateRemoveDependencyParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"removeDependency",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeliveryStream) ReplaceDependency(target awscdk.CfnResource, newTarget awscdk.CfnResource) {
	if err := c.validateReplaceDependencyParameters(target, newTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"replaceDependency",
		[]interface{}{target, newTarget},
	)
}

func (c *jsiiProxy_CfnDeliveryStream) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeliveryStream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnDeliveryStream) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}
