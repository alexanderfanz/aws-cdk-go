package integtests

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.integ_tests.ActualResult",
		reflect.TypeOf((*ActualResult)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "result", GoGetter: "Result"},
		},
		func() interface{} {
			return &jsiiProxy_ActualResult{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.integ_tests.AssertionRequest",
		reflect.TypeOf((*AssertionRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.integ_tests.AssertionResult",
		reflect.TypeOf((*AssertionResult)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.integ_tests.AssertionResultData",
		reflect.TypeOf((*AssertionResultData)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.integ_tests.AssertionType",
		reflect.TypeOf((*AssertionType)(nil)).Elem(),
		map[string]interface{}{
			"EQUALS": AssertionType_EQUALS,
			"OBJECT_LIKE": AssertionType_OBJECT_LIKE,
			"ARRAY_WITH": AssertionType_ARRAY_WITH,
		},
	)
	_jsii_.RegisterClass(
		"monocdk.integ_tests.AssertionsProvider",
		reflect.TypeOf((*AssertionsProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addPolicyStatementFromSdkCall", GoMethod: "AddPolicyStatementFromSdkCall"},
			_jsii_.MemberMethod{JsiiMethod: "encode", GoMethod: "Encode"},
			_jsii_.MemberProperty{JsiiProperty: "handlerRoleArn", GoGetter: "HandlerRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "serviceToken", GoGetter: "ServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_AssertionsProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.integ_tests.AwsApiCall",
		reflect.TypeOf((*AwsApiCall)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assertAtPath", GoMethod: "AssertAtPath"},
			_jsii_.MemberMethod{JsiiMethod: "expect", GoMethod: "Expect"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getAttString", GoMethod: "GetAttString"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_AwsApiCall{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAwsApiCall)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.integ_tests.AwsApiCallOptions",
		reflect.TypeOf((*AwsApiCallOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.integ_tests.AwsApiCallProps",
		reflect.TypeOf((*AwsApiCallProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.integ_tests.AwsApiCallRequest",
		reflect.TypeOf((*AwsApiCallRequest)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.integ_tests.AwsApiCallResult",
		reflect.TypeOf((*AwsApiCallResult)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.integ_tests.EqualsAssertion",
		reflect.TypeOf((*EqualsAssertion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "result", GoGetter: "Result"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_EqualsAssertion{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.integ_tests.EqualsAssertionProps",
		reflect.TypeOf((*EqualsAssertionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.integ_tests.ExpectedResult",
		reflect.TypeOf((*ExpectedResult)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "result", GoGetter: "Result"},
		},
		func() interface{} {
			return &jsiiProxy_ExpectedResult{}
		},
	)
	_jsii_.RegisterInterface(
		"monocdk.integ_tests.IAwsApiCall",
		reflect.TypeOf((*IAwsApiCall)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assertAtPath", GoMethod: "AssertAtPath"},
			_jsii_.MemberMethod{JsiiMethod: "expect", GoMethod: "Expect"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getAttString", GoMethod: "GetAttString"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_IAwsApiCall{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIConstruct)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"monocdk.integ_tests.IDeployAssert",
		reflect.TypeOf((*IDeployAssert)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "awsApiCall", GoMethod: "AwsApiCall"},
			_jsii_.MemberMethod{JsiiMethod: "expect", GoMethod: "Expect"},
			_jsii_.MemberMethod{JsiiMethod: "invokeFunction", GoMethod: "InvokeFunction"},
		},
		func() interface{} {
			return &jsiiProxy_IDeployAssert{}
		},
	)
	_jsii_.RegisterClass(
		"monocdk.integ_tests.IntegTest",
		reflect.TypeOf((*IntegTest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assertions", GoGetter: "Assertions"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_IntegTest{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.integ_tests.IntegTestCase",
		reflect.TypeOf((*IntegTestCase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "assertions", GoGetter: "Assertions"},
			_jsii_.MemberProperty{JsiiProperty: "manifest", GoGetter: "Manifest"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_IntegTestCase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.integ_tests.IntegTestCaseProps",
		reflect.TypeOf((*IntegTestCaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.integ_tests.IntegTestCaseStack",
		reflect.TypeOf((*IntegTestCaseStack)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "account", GoGetter: "Account"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDockerImageAsset", GoMethod: "AddDockerImageAsset"},
			_jsii_.MemberMethod{JsiiMethod: "addFileAsset", GoMethod: "AddFileAsset"},
			_jsii_.MemberMethod{JsiiMethod: "addTransform", GoMethod: "AddTransform"},
			_jsii_.MemberMethod{JsiiMethod: "allocateLogicalId", GoMethod: "AllocateLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "artifactId", GoGetter: "ArtifactId"},
			_jsii_.MemberProperty{JsiiProperty: "assertions", GoGetter: "Assertions"},
			_jsii_.MemberProperty{JsiiProperty: "availabilityZones", GoGetter: "AvailabilityZones"},
			_jsii_.MemberProperty{JsiiProperty: "bundlingRequired", GoGetter: "BundlingRequired"},
			_jsii_.MemberProperty{JsiiProperty: "dependencies", GoGetter: "Dependencies"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberMethod{JsiiMethod: "exportValue", GoMethod: "ExportValue"},
			_jsii_.MemberMethod{JsiiMethod: "formatArn", GoMethod: "FormatArn"},
			_jsii_.MemberMethod{JsiiMethod: "getLogicalId", GoMethod: "GetLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "nested", GoGetter: "Nested"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackParent", GoGetter: "NestedStackParent"},
			_jsii_.MemberProperty{JsiiProperty: "nestedStackResource", GoGetter: "NestedStackResource"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "notificationArns", GoGetter: "NotificationArns"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "parentStack", GoGetter: "ParentStack"},
			_jsii_.MemberMethod{JsiiMethod: "parseArn", GoMethod: "ParseArn"},
			_jsii_.MemberProperty{JsiiProperty: "partition", GoGetter: "Partition"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberMethod{JsiiMethod: "prepareCrossReference", GoMethod: "PrepareCrossReference"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "regionalFact", GoMethod: "RegionalFact"},
			_jsii_.MemberMethod{JsiiMethod: "renameLogicalId", GoMethod: "RenameLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "reportMissingContext", GoMethod: "ReportMissingContext"},
			_jsii_.MemberMethod{JsiiMethod: "reportMissingContextKey", GoMethod: "ReportMissingContextKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberMethod{JsiiMethod: "splitArn", GoMethod: "SplitArn"},
			_jsii_.MemberProperty{JsiiProperty: "stackId", GoGetter: "StackId"},
			_jsii_.MemberProperty{JsiiProperty: "stackName", GoGetter: "StackName"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "synthesizer", GoGetter: "Synthesizer"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "templateFile", GoGetter: "TemplateFile"},
			_jsii_.MemberProperty{JsiiProperty: "templateOptions", GoGetter: "TemplateOptions"},
			_jsii_.MemberProperty{JsiiProperty: "terminationProtection", GoGetter: "TerminationProtection"},
			_jsii_.MemberMethod{JsiiMethod: "toJsonString", GoMethod: "ToJsonString"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "urlSuffix", GoGetter: "UrlSuffix"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_IntegTestCaseStack{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkStack)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.integ_tests.IntegTestCaseStackProps",
		reflect.TypeOf((*IntegTestCaseStackProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.integ_tests.IntegTestProps",
		reflect.TypeOf((*IntegTestProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.integ_tests.InvocationType",
		reflect.TypeOf((*InvocationType)(nil)).Elem(),
		map[string]interface{}{
			"EVENT": InvocationType_EVENT,
			"REQUEST_RESPONE": InvocationType_REQUEST_RESPONE,
			"DRY_RUN": InvocationType_DRY_RUN,
		},
	)
	_jsii_.RegisterClass(
		"monocdk.integ_tests.LambdaInvokeFunction",
		reflect.TypeOf((*LambdaInvokeFunction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "assertAtPath", GoMethod: "AssertAtPath"},
			_jsii_.MemberMethod{JsiiMethod: "expect", GoMethod: "Expect"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getAttString", GoMethod: "GetAttString"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaInvokeFunction{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AwsApiCall)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.integ_tests.LambdaInvokeFunctionProps",
		reflect.TypeOf((*LambdaInvokeFunctionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.integ_tests.LogType",
		reflect.TypeOf((*LogType)(nil)).Elem(),
		map[string]interface{}{
			"NONE": LogType_NONE,
			"TAIL": LogType_TAIL,
		},
	)
	_jsii_.RegisterClass(
		"monocdk.integ_tests.Match",
		reflect.TypeOf((*Match)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_Match{}
		},
	)
	_jsii_.RegisterEnum(
		"monocdk.integ_tests.Status",
		reflect.TypeOf((*Status)(nil)).Elem(),
		map[string]interface{}{
			"PASS": Status_PASS,
			"FAIL": Status_FAIL,
		},
	)
}