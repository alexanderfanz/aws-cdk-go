package awsbackup


// Specifies an object containing properties used to schedule a task to back up a selection of resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   backupRuleResourceTypeProperty := &BackupRuleResourceTypeProperty{
//   	RuleName: jsii.String("ruleName"),
//   	TargetBackupVault: jsii.String("targetBackupVault"),
//
//   	// the properties below are optional
//   	CompletionWindowMinutes: jsii.Number(123),
//   	CopyActions: []interface{}{
//   		&CopyActionResourceTypeProperty{
//   			DestinationBackupVaultArn: jsii.String("destinationBackupVaultArn"),
//
//   			// the properties below are optional
//   			Lifecycle: &LifecycleResourceTypeProperty{
//   				DeleteAfterDays: jsii.Number(123),
//   				MoveToColdStorageAfterDays: jsii.Number(123),
//   			},
//   		},
//   	},
//   	EnableContinuousBackup: jsii.Boolean(false),
//   	Lifecycle: &LifecycleResourceTypeProperty{
//   		DeleteAfterDays: jsii.Number(123),
//   		MoveToColdStorageAfterDays: jsii.Number(123),
//   	},
//   	RecoveryPointTags: map[string]*string{
//   		"recoveryPointTagsKey": jsii.String("recoveryPointTags"),
//   	},
//   	ScheduleExpression: jsii.String("scheduleExpression"),
//   	StartWindowMinutes: jsii.Number(123),
//   }
//
type CfnBackupPlan_BackupRuleResourceTypeProperty struct {
	// A display name for a backup rule.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// The name of a logical container where backups are stored.
	//
	// Backup vaults are identified by names that are unique to the account used to create them and the AWS Region where they are created. They consist of letters, numbers, and hyphens.
	TargetBackupVault *string `field:"required" json:"targetBackupVault" yaml:"targetBackupVault"`
	// A value in minutes after a backup job is successfully started before it must be completed or it is canceled by AWS Backup .
	CompletionWindowMinutes *float64 `field:"optional" json:"completionWindowMinutes" yaml:"completionWindowMinutes"`
	// An array of CopyAction objects, which contains the details of the copy operation.
	CopyActions interface{} `field:"optional" json:"copyActions" yaml:"copyActions"`
	// Enables continuous backup and point-in-time restores (PITR).
	EnableContinuousBackup interface{} `field:"optional" json:"enableContinuousBackup" yaml:"enableContinuousBackup"`
	// The lifecycle defines when a protected resource is transitioned to cold storage and when it expires.
	//
	// AWS Backup transitions and expires backups automatically according to the lifecycle that you define.
	Lifecycle interface{} `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// To help organize your resources, you can assign your own metadata to the resources that you create.
	//
	// Each tag is a key-value pair.
	RecoveryPointTags interface{} `field:"optional" json:"recoveryPointTags" yaml:"recoveryPointTags"`
	// A CRON expression specifying when AWS Backup initiates a backup job.
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
	// An optional value that specifies a period of time in minutes after a backup is scheduled before a job is canceled if it doesn't start successfully.
	//
	// If this value is included, it must be at least 60 minutes to avoid errors.
	StartWindowMinutes *float64 `field:"optional" json:"startWindowMinutes" yaml:"startWindowMinutes"`
}

