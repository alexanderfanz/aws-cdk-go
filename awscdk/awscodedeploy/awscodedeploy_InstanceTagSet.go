package awscodedeploy

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a set of instance tag groups.
//
// An instance will match a set if it matches all of the groups in the set -
// in other words, sets follow 'and' semantics.
// You can have a maximum of 3 tag groups inside a set.
//
// Example:
//   import autoscaling "github.com/aws/aws-cdk-go/awscdk"
//   import cloudwatch "github.com/aws/aws-cdk-go/awscdk"
//
//   var application serverApplication
//   var asg autoScalingGroup
//   var alarm alarm
//
//   deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("CodeDeployDeploymentGroup"), &serverDeploymentGroupProps{
//   	application: application,
//   	deploymentGroupName: jsii.String("MyDeploymentGroup"),
//   	autoScalingGroups: []iAutoScalingGroup{
//   		asg,
//   	},
//   	// adds User Data that installs the CodeDeploy agent on your auto-scaling groups hosts
//   	// default: true
//   	installAgent: jsii.Boolean(true),
//   	// adds EC2 instances matching tags
//   	ec2InstanceTags: codedeploy.NewInstanceTagSet(map[string][]*string{
//   		// any instance with tags satisfying
//   		// key1=v1 or key1=v2 or key2 (any value) or value v3 (any key)
//   		// will match this group
//   		"key1": []*string{
//   			jsii.String("v1"),
//   			jsii.String("v2"),
//   		},
//   		"key2": []*string{
//   		},
//   		"": []*string{
//   			jsii.String("v3"),
//   		},
//   	}),
//   	// adds on-premise instances matching tags
//   	onPremiseInstanceTags: codedeploy.NewInstanceTagSet(map[string][]*string{
//   		"key1": []*string{
//   			jsii.String("v1"),
//   			jsii.String("v2"),
//   		},
//   	}, map[string][]*string{
//   		"key2": []*string{
//   			jsii.String("v3"),
//   		},
//   	}),
//   	// CloudWatch alarms
//   	alarms: []iAlarm{
//   		alarm,
//   	},
//   	// whether to ignore failure to fetch the status of alarms from CloudWatch
//   	// default: false
//   	ignorePollAlarmsFailure: jsii.Boolean(false),
//   	// auto-rollback configuration
//   	autoRollback: &autoRollbackConfig{
//   		failedDeployment: jsii.Boolean(true),
//   		 // default: true
//   		stoppedDeployment: jsii.Boolean(true),
//   		 // default: false
//   		deploymentInAlarm: jsii.Boolean(true),
//   	},
//   })
//
// Experimental.
type InstanceTagSet interface {
	// Experimental.
	InstanceTagGroups() *[]*map[string]*[]*string
}

// The jsii proxy struct for InstanceTagSet
type jsiiProxy_InstanceTagSet struct {
	_ byte // padding
}

func (j *jsiiProxy_InstanceTagSet) InstanceTagGroups() *[]*map[string]*[]*string {
	var returns *[]*map[string]*[]*string
	_jsii_.Get(
		j,
		"instanceTagGroups",
		&returns,
	)
	return returns
}


// Experimental.
func NewInstanceTagSet(instanceTagGroups ...*map[string]*[]*string) InstanceTagSet {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range instanceTagGroups {
		args = append(args, a)
	}

	j := jsiiProxy_InstanceTagSet{}

	_jsii_.Create(
		"monocdk.aws_codedeploy.InstanceTagSet",
		args,
		&j,
	)

	return &j
}

// Experimental.
func NewInstanceTagSet_Override(i InstanceTagSet, instanceTagGroups ...*map[string]*[]*string) {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range instanceTagGroups {
		args = append(args, a)
	}

	_jsii_.Create(
		"monocdk.aws_codedeploy.InstanceTagSet",
		args,
		i,
	)
}

