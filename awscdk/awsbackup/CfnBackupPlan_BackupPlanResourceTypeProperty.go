package awsbackup


// Specifies an object containing properties used to create a backup plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var backupOptions interface{}
//
//   backupPlanResourceTypeProperty := &BackupPlanResourceTypeProperty{
//   	BackupPlanName: jsii.String("backupPlanName"),
//   	BackupPlanRule: []interface{}{
//   		&BackupRuleResourceTypeProperty{
//   			RuleName: jsii.String("ruleName"),
//   			TargetBackupVault: jsii.String("targetBackupVault"),
//
//   			// the properties below are optional
//   			CompletionWindowMinutes: jsii.Number(123),
//   			CopyActions: []interface{}{
//   				&CopyActionResourceTypeProperty{
//   					DestinationBackupVaultArn: jsii.String("destinationBackupVaultArn"),
//
//   					// the properties below are optional
//   					Lifecycle: &LifecycleResourceTypeProperty{
//   						DeleteAfterDays: jsii.Number(123),
//   						MoveToColdStorageAfterDays: jsii.Number(123),
//   					},
//   				},
//   			},
//   			EnableContinuousBackup: jsii.Boolean(false),
//   			Lifecycle: &LifecycleResourceTypeProperty{
//   				DeleteAfterDays: jsii.Number(123),
//   				MoveToColdStorageAfterDays: jsii.Number(123),
//   			},
//   			RecoveryPointTags: map[string]*string{
//   				"recoveryPointTagsKey": jsii.String("recoveryPointTags"),
//   			},
//   			ScheduleExpression: jsii.String("scheduleExpression"),
//   			StartWindowMinutes: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	AdvancedBackupSettings: []interface{}{
//   		&AdvancedBackupSettingResourceTypeProperty{
//   			BackupOptions: backupOptions,
//   			ResourceType: jsii.String("resourceType"),
//   		},
//   	},
//   }
//
type CfnBackupPlan_BackupPlanResourceTypeProperty struct {
	// The display name of a backup plan.
	BackupPlanName *string `field:"required" json:"backupPlanName" yaml:"backupPlanName"`
	// An array of `BackupRule` objects, each of which specifies a scheduled task that is used to back up a selection of resources.
	BackupPlanRule interface{} `field:"required" json:"backupPlanRule" yaml:"backupPlanRule"`
	// A list of backup options for each resource type.
	AdvancedBackupSettings interface{} `field:"optional" json:"advancedBackupSettings" yaml:"advancedBackupSettings"`
}
