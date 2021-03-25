package awscodepipelineactions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.Action",
		reflect.TypeOf((*Action)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_Action{}
			_jsii_.InitJsiiProxy(&j.Type__awscodepipelineIAction)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.AlexaSkillDeployAction",
		reflect.TypeOf((*AlexaSkillDeployAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_AlexaSkillDeployAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.AlexaSkillDeployActionProps",
		reflect.TypeOf((*AlexaSkillDeployActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.BaseJenkinsProvider",
		reflect.TypeOf((*BaseJenkinsProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "providerName", GoGetter: "ProviderName"},
			_jsii_.MemberProperty{JsiiProperty: "serverUrl", GoGetter: "ServerUrl"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_BaseJenkinsProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IJenkinsProvider)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.BitBucketSourceAction",
		reflect.TypeOf((*BitBucketSourceAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_BitBucketSourceAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.BitBucketSourceActionProps",
		reflect.TypeOf((*BitBucketSourceActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.CacheControl",
		reflect.TypeOf((*CacheControl)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			return &jsiiProxy_CacheControl{}
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.CloudFormationCreateReplaceChangeSetAction",
		reflect.TypeOf((*CloudFormationCreateReplaceChangeSetAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "addToDeploymentRolePolicy", GoMethod: "AddToDeploymentRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentRole", GoGetter: "DeploymentRole"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_CloudFormationCreateReplaceChangeSetAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.CloudFormationCreateReplaceChangeSetActionProps",
		reflect.TypeOf((*CloudFormationCreateReplaceChangeSetActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.CloudFormationCreateUpdateStackAction",
		reflect.TypeOf((*CloudFormationCreateUpdateStackAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "addToDeploymentRolePolicy", GoMethod: "AddToDeploymentRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentRole", GoGetter: "DeploymentRole"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_CloudFormationCreateUpdateStackAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.CloudFormationCreateUpdateStackActionProps",
		reflect.TypeOf((*CloudFormationCreateUpdateStackActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.CloudFormationDeleteStackAction",
		reflect.TypeOf((*CloudFormationDeleteStackAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "addToDeploymentRolePolicy", GoMethod: "AddToDeploymentRolePolicy"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberProperty{JsiiProperty: "deploymentRole", GoGetter: "DeploymentRole"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_CloudFormationDeleteStackAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.CloudFormationDeleteStackActionProps",
		reflect.TypeOf((*CloudFormationDeleteStackActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.CloudFormationExecuteChangeSetAction",
		reflect.TypeOf((*CloudFormationExecuteChangeSetAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_CloudFormationExecuteChangeSetAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.CloudFormationExecuteChangeSetActionProps",
		reflect.TypeOf((*CloudFormationExecuteChangeSetActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.CodeBuildAction",
		reflect.TypeOf((*CodeBuildAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variable", GoMethod: "Variable"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_CodeBuildAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.CodeBuildActionProps",
		reflect.TypeOf((*CodeBuildActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_codepipeline_actions.CodeBuildActionType",
		reflect.TypeOf((*CodeBuildActionType)(nil)).Elem(),
		map[string]interface{}{
			"BUILD": CodeBuildActionType_BUILD,
			"TEST": CodeBuildActionType_TEST,
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.CodeCommitSourceAction",
		reflect.TypeOf((*CodeCommitSourceAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
			_jsii_.MemberProperty{JsiiProperty: "variables", GoGetter: "Variables"},
		},
		func() interface{} {
			j := jsiiProxy_CodeCommitSourceAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.CodeCommitSourceActionProps",
		reflect.TypeOf((*CodeCommitSourceActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.CodeCommitSourceVariables",
		reflect.TypeOf((*CodeCommitSourceVariables)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_codepipeline_actions.CodeCommitTrigger",
		reflect.TypeOf((*CodeCommitTrigger)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CodeCommitTrigger_NONE,
			"POLL": CodeCommitTrigger_POLL,
			"EVENTS": CodeCommitTrigger_EVENTS,
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.CodeDeployEcsContainerImageInput",
		reflect.TypeOf((*CodeDeployEcsContainerImageInput)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.CodeDeployEcsDeployAction",
		reflect.TypeOf((*CodeDeployEcsDeployAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_CodeDeployEcsDeployAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.CodeDeployEcsDeployActionProps",
		reflect.TypeOf((*CodeDeployEcsDeployActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.CodeDeployServerDeployAction",
		reflect.TypeOf((*CodeDeployServerDeployAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_CodeDeployServerDeployAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.CodeDeployServerDeployActionProps",
		reflect.TypeOf((*CodeDeployServerDeployActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.EcrSourceAction",
		reflect.TypeOf((*EcrSourceAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
			_jsii_.MemberProperty{JsiiProperty: "variables", GoGetter: "Variables"},
		},
		func() interface{} {
			j := jsiiProxy_EcrSourceAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.EcrSourceActionProps",
		reflect.TypeOf((*EcrSourceActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.EcrSourceVariables",
		reflect.TypeOf((*EcrSourceVariables)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.EcsDeployAction",
		reflect.TypeOf((*EcsDeployAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_EcsDeployAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.EcsDeployActionProps",
		reflect.TypeOf((*EcsDeployActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.GitHubSourceAction",
		reflect.TypeOf((*GitHubSourceAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
			_jsii_.MemberProperty{JsiiProperty: "variables", GoGetter: "Variables"},
		},
		func() interface{} {
			j := jsiiProxy_GitHubSourceAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.GitHubSourceActionProps",
		reflect.TypeOf((*GitHubSourceActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.GitHubSourceVariables",
		reflect.TypeOf((*GitHubSourceVariables)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_codepipeline_actions.GitHubTrigger",
		reflect.TypeOf((*GitHubTrigger)(nil)).Elem(),
		map[string]interface{}{
			"NONE": GitHubTrigger_NONE,
			"POLL": GitHubTrigger_POLL,
			"WEBHOOK": GitHubTrigger_WEBHOOK,
		},
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_codepipeline_actions.IJenkinsProvider",
		reflect.TypeOf((*IJenkinsProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "providerName", GoGetter: "ProviderName"},
			_jsii_.MemberProperty{JsiiProperty: "serverUrl", GoGetter: "ServerUrl"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_IJenkinsProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.JenkinsAction",
		reflect.TypeOf((*JenkinsAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_JenkinsAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.JenkinsActionProps",
		reflect.TypeOf((*JenkinsActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_codepipeline_actions.JenkinsActionType",
		reflect.TypeOf((*JenkinsActionType)(nil)).Elem(),
		map[string]interface{}{
			"BUILD": JenkinsActionType_BUILD,
			"TEST": JenkinsActionType_TEST,
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.JenkinsProvider",
		reflect.TypeOf((*JenkinsProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "providerName", GoGetter: "ProviderName"},
			_jsii_.MemberProperty{JsiiProperty: "serverUrl", GoGetter: "ServerUrl"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_JenkinsProvider{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_BaseJenkinsProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.JenkinsProviderAttributes",
		reflect.TypeOf((*JenkinsProviderAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.JenkinsProviderProps",
		reflect.TypeOf((*JenkinsProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.LambdaInvokeAction",
		reflect.TypeOf((*LambdaInvokeAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variable", GoMethod: "Variable"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaInvokeAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.LambdaInvokeActionProps",
		reflect.TypeOf((*LambdaInvokeActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.ManualApprovalAction",
		reflect.TypeOf((*ManualApprovalAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberProperty{JsiiProperty: "notificationTopic", GoGetter: "NotificationTopic"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_ManualApprovalAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.ManualApprovalActionProps",
		reflect.TypeOf((*ManualApprovalActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.S3DeployAction",
		reflect.TypeOf((*S3DeployAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_S3DeployAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.S3DeployActionProps",
		reflect.TypeOf((*S3DeployActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.S3SourceAction",
		reflect.TypeOf((*S3SourceAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
			_jsii_.MemberProperty{JsiiProperty: "variables", GoGetter: "Variables"},
		},
		func() interface{} {
			j := jsiiProxy_S3SourceAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.S3SourceActionProps",
		reflect.TypeOf((*S3SourceActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.S3SourceVariables",
		reflect.TypeOf((*S3SourceVariables)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_codepipeline_actions.S3Trigger",
		reflect.TypeOf((*S3Trigger)(nil)).Elem(),
		map[string]interface{}{
			"NONE": S3Trigger_NONE,
			"POLL": S3Trigger_POLL,
			"EVENTS": S3Trigger_EVENTS,
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.ServiceCatalogDeployAction",
		reflect.TypeOf((*ServiceCatalogDeployAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_ServiceCatalogDeployAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.ServiceCatalogDeployActionProps",
		reflect.TypeOf((*ServiceCatalogDeployActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.StateMachineInput",
		reflect.TypeOf((*StateMachineInput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "input", GoGetter: "Input"},
			_jsii_.MemberProperty{JsiiProperty: "inputArtifact", GoGetter: "InputArtifact"},
			_jsii_.MemberProperty{JsiiProperty: "inputType", GoGetter: "InputType"},
		},
		func() interface{} {
			return &jsiiProxy_StateMachineInput{}
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_codepipeline_actions.StepFunctionInvokeAction",
		reflect.TypeOf((*StepFunctionInvokeAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actionProperties", GoGetter: "ActionProperties"},
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
			_jsii_.MemberMethod{JsiiMethod: "bound", GoMethod: "Bound"},
			_jsii_.MemberMethod{JsiiMethod: "onStateChange", GoMethod: "OnStateChange"},
			_jsii_.MemberMethod{JsiiMethod: "variableExpression", GoMethod: "VariableExpression"},
		},
		func() interface{} {
			j := jsiiProxy_StepFunctionInvokeAction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Action)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_codepipeline_actions.StepFunctionsInvokeActionProps",
		reflect.TypeOf((*StepFunctionsInvokeActionProps)(nil)).Elem(),
	)
}
