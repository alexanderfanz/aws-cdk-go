package awsdlm


// *[Snapshot and AMI policies only]* Specifies a schedule for a snapshot or AMI lifecycle policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleProperty := &ScheduleProperty{
//   	ArchiveRule: &ArchiveRuleProperty{
//   		RetainRule: &ArchiveRetainRuleProperty{
//   			RetentionArchiveTier: &RetentionArchiveTierProperty{
//   				Count: jsii.Number(123),
//   				Interval: jsii.Number(123),
//   				IntervalUnit: jsii.String("intervalUnit"),
//   			},
//   		},
//   	},
//   	CopyTags: jsii.Boolean(false),
//   	CreateRule: &CreateRuleProperty{
//   		CronExpression: jsii.String("cronExpression"),
//   		Interval: jsii.Number(123),
//   		IntervalUnit: jsii.String("intervalUnit"),
//   		Location: jsii.String("location"),
//   		Times: []*string{
//   			jsii.String("times"),
//   		},
//   	},
//   	CrossRegionCopyRules: []interface{}{
//   		&CrossRegionCopyRuleProperty{
//   			Encrypted: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			CmkArn: jsii.String("cmkArn"),
//   			CopyTags: jsii.Boolean(false),
//   			DeprecateRule: &CrossRegionCopyDeprecateRuleProperty{
//   				Interval: jsii.Number(123),
//   				IntervalUnit: jsii.String("intervalUnit"),
//   			},
//   			RetainRule: &CrossRegionCopyRetainRuleProperty{
//   				Interval: jsii.Number(123),
//   				IntervalUnit: jsii.String("intervalUnit"),
//   			},
//   			Target: jsii.String("target"),
//   			TargetRegion: jsii.String("targetRegion"),
//   		},
//   	},
//   	DeprecateRule: &DeprecateRuleProperty{
//   		Count: jsii.Number(123),
//   		Interval: jsii.Number(123),
//   		IntervalUnit: jsii.String("intervalUnit"),
//   	},
//   	FastRestoreRule: &FastRestoreRuleProperty{
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		Count: jsii.Number(123),
//   		Interval: jsii.Number(123),
//   		IntervalUnit: jsii.String("intervalUnit"),
//   	},
//   	Name: jsii.String("name"),
//   	RetainRule: &RetainRuleProperty{
//   		Count: jsii.Number(123),
//   		Interval: jsii.Number(123),
//   		IntervalUnit: jsii.String("intervalUnit"),
//   	},
//   	ShareRules: []interface{}{
//   		&ShareRuleProperty{
//   			TargetAccounts: []*string{
//   				jsii.String("targetAccounts"),
//   			},
//   			UnshareInterval: jsii.Number(123),
//   			UnshareIntervalUnit: jsii.String("unshareIntervalUnit"),
//   		},
//   	},
//   	TagsToAdd: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VariableTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnLifecyclePolicy_ScheduleProperty struct {
	// *[Snapshot policies that target volumes only]* The snapshot archiving rule for the schedule.
	//
	// When you specify an archiving rule, snapshots are automatically moved from the standard tier to the archive tier once the schedule's retention threshold is met. Snapshots are then retained in the archive tier for the archive retention period that you specify.
	//
	// For more information about using snapshot archiving, see [Considerations for snapshot lifecycle policies](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/snapshot-ami-policy.html#dlm-archive) .
	ArchiveRule interface{} `field:"optional" json:"archiveRule" yaml:"archiveRule"`
	// Copy all user-defined tags on a source volume to snapshots of the volume created by this policy.
	CopyTags interface{} `field:"optional" json:"copyTags" yaml:"copyTags"`
	// The creation rule.
	CreateRule interface{} `field:"optional" json:"createRule" yaml:"createRule"`
	// Specifies a rule for copying snapshots or AMIs across regions.
	//
	// > You can't specify cross-Region copy rules for policies that create snapshots on an Outpost. If the policy creates snapshots in a Region, then snapshots can be copied to up to three Regions or Outposts.
	CrossRegionCopyRules interface{} `field:"optional" json:"crossRegionCopyRules" yaml:"crossRegionCopyRules"`
	// *[AMI policies only]* The AMI deprecation rule for the schedule.
	DeprecateRule interface{} `field:"optional" json:"deprecateRule" yaml:"deprecateRule"`
	// *[Snapshot policies only]* The rule for enabling fast snapshot restore.
	FastRestoreRule interface{} `field:"optional" json:"fastRestoreRule" yaml:"fastRestoreRule"`
	// The name of the schedule.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The retention rule for snapshots or AMIs created by the policy.
	RetainRule interface{} `field:"optional" json:"retainRule" yaml:"retainRule"`
	// *[Snapshot policies only]* The rule for sharing snapshots with other AWS accounts .
	ShareRules interface{} `field:"optional" json:"shareRules" yaml:"shareRules"`
	// The tags to apply to policy-created resources.
	//
	// These user-defined tags are in addition to the AWS -added lifecycle tags.
	TagsToAdd interface{} `field:"optional" json:"tagsToAdd" yaml:"tagsToAdd"`
	// *[AMI policies and snapshot policies that target instances only]* A collection of key/value pairs with values determined dynamically when the policy is executed.
	//
	// Keys may be any valid Amazon EC2 tag key. Values must be in one of the two following formats: `$(instance-id)` or `$(timestamp)` . Variable tags are only valid for EBS Snapshot Management – Instance policies.
	VariableTags interface{} `field:"optional" json:"variableTags" yaml:"variableTags"`
}

