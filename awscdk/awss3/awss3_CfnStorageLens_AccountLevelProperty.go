package awss3


// This resource contains the details of the account-level metrics for Amazon S3 Storage Lens.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accountLevelProperty := &AccountLevelProperty{
//   	BucketLevel: &BucketLevelProperty{
//   		ActivityMetrics: &ActivityMetricsProperty{
//   			IsEnabled: jsii.Boolean(false),
//   		},
//   		AdvancedCostOptimizationMetrics: &AdvancedCostOptimizationMetricsProperty{
//   			IsEnabled: jsii.Boolean(false),
//   		},
//   		AdvancedDataProtectionMetrics: &AdvancedDataProtectionMetricsProperty{
//   			IsEnabled: jsii.Boolean(false),
//   		},
//   		DetailedStatusCodesMetrics: &DetailedStatusCodesMetricsProperty{
//   			IsEnabled: jsii.Boolean(false),
//   		},
//   		PrefixLevel: &PrefixLevelProperty{
//   			StorageMetrics: &PrefixLevelStorageMetricsProperty{
//   				IsEnabled: jsii.Boolean(false),
//   				SelectionCriteria: &SelectionCriteriaProperty{
//   					Delimiter: jsii.String("delimiter"),
//   					MaxDepth: jsii.Number(123),
//   					MinStorageBytesPercentage: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	ActivityMetrics: &ActivityMetricsProperty{
//   		IsEnabled: jsii.Boolean(false),
//   	},
//   	AdvancedCostOptimizationMetrics: &AdvancedCostOptimizationMetricsProperty{
//   		IsEnabled: jsii.Boolean(false),
//   	},
//   	AdvancedDataProtectionMetrics: &AdvancedDataProtectionMetricsProperty{
//   		IsEnabled: jsii.Boolean(false),
//   	},
//   	DetailedStatusCodesMetrics: &DetailedStatusCodesMetricsProperty{
//   		IsEnabled: jsii.Boolean(false),
//   	},
//   }
//
type CfnStorageLens_AccountLevelProperty struct {
	// This property contains the details of the account-level bucket-level configurations for Amazon S3 Storage Lens.
	BucketLevel interface{} `field:"required" json:"bucketLevel" yaml:"bucketLevel"`
	// This property contains the details of account-level activity metrics for S3 Storage Lens.
	ActivityMetrics interface{} `field:"optional" json:"activityMetrics" yaml:"activityMetrics"`
	// This property contains the details of account-level advanced cost optimization metrics for S3 Storage Lens.
	AdvancedCostOptimizationMetrics interface{} `field:"optional" json:"advancedCostOptimizationMetrics" yaml:"advancedCostOptimizationMetrics"`
	// This property contains the details of account-level advanced data protection metrics for S3 Storage Lens.
	AdvancedDataProtectionMetrics interface{} `field:"optional" json:"advancedDataProtectionMetrics" yaml:"advancedDataProtectionMetrics"`
	// This property contains the details of account-level detailed status code metrics for S3 Storage Lens.
	DetailedStatusCodesMetrics interface{} `field:"optional" json:"detailedStatusCodesMetrics" yaml:"detailedStatusCodesMetrics"`
}

