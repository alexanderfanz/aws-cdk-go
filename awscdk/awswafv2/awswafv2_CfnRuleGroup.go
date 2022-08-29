package awswafv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A CloudFormation `AWS::WAFv2::RuleGroup`.
//
// > This is the latest version of *AWS WAF* , named AWS WAF V2, released in November, 2019. For information, including how to migrate your AWS WAF resources from the prior release, see the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
//
// Use an `RuleGroup` to define a collection of rules for inspecting and controlling web requests. You use a rule group in an `WebACL` by providing its Amazon Resource Name (ARN) to the rule statement `RuleGroupReferenceStatement` , when you add rules to the web ACL.
//
// When you create a rule group, you define an immutable capacity limit. If you update a rule group, you must stay within the capacity. This allows others to reuse the rule group with confidence in its capacity requirements.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var all interface{}
//   var allow interface{}
//   var allQueryArguments interface{}
//   var block interface{}
//   var captcha interface{}
//   var count interface{}
//   var method interface{}
//   var queryString interface{}
//   var singleHeader interface{}
//   var singleQueryArgument interface{}
//   var statementProperty_ statementProperty
//   var uriPath interface{}
//
//   cfnRuleGroup := awscdk.Aws_wafv2.NewCfnRuleGroup(this, jsii.String("MyCfnRuleGroup"), &cfnRuleGroupProps{
//   	capacity: jsii.Number(123),
//   	scope: jsii.String("scope"),
//   	visibilityConfig: &visibilityConfigProperty{
//   		cloudWatchMetricsEnabled: jsii.Boolean(false),
//   		metricName: jsii.String("metricName"),
//   		sampledRequestsEnabled: jsii.Boolean(false),
//   	},
//
//   	// the properties below are optional
//   	customResponseBodies: map[string]interface{}{
//   		"customResponseBodiesKey": &CustomResponseBodyProperty{
//   			"content": jsii.String("content"),
//   			"contentType": jsii.String("contentType"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	name: jsii.String("name"),
//   	rules: []interface{}{
//   		&ruleProperty{
//   			name: jsii.String("name"),
//   			priority: jsii.Number(123),
//   			statement: &statementProperty{
//   				andStatement: &andStatementProperty{
//   					statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				byteMatchStatement: &byteMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					positionalConstraint: jsii.String("positionalConstraint"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					searchString: jsii.String("searchString"),
//   					searchStringBase64: jsii.String("searchStringBase64"),
//   				},
//   				geoMatchStatement: &geoMatchStatementProperty{
//   					countryCodes: []*string{
//   						jsii.String("countryCodes"),
//   					},
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   				},
//   				ipSetReferenceStatement: map[string]interface{}{
//   					"arn": jsii.String("arn"),
//
//   					// the properties below are optional
//   					"ipSetForwardedIpConfig": map[string]*string{
//   						"fallbackBehavior": jsii.String("fallbackBehavior"),
//   						"headerName": jsii.String("headerName"),
//   						"position": jsii.String("position"),
//   					},
//   				},
//   				labelMatchStatement: &labelMatchStatementProperty{
//   					key: jsii.String("key"),
//   					scope: jsii.String("scope"),
//   				},
//   				notStatement: &notStatementProperty{
//   					statement: statementProperty_,
//   				},
//   				orStatement: &orStatementProperty{
//   					statements: []interface{}{
//   						statementProperty_,
//   					},
//   				},
//   				rateBasedStatement: &rateBasedStatementProperty{
//   					aggregateKeyType: jsii.String("aggregateKeyType"),
//   					limit: jsii.Number(123),
//
//   					// the properties below are optional
//   					forwardedIpConfig: &forwardedIPConfigurationProperty{
//   						fallbackBehavior: jsii.String("fallbackBehavior"),
//   						headerName: jsii.String("headerName"),
//   					},
//   					scopeDownStatement: statementProperty_,
//   				},
//   				regexMatchStatement: &regexMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					regexString: jsii.String("regexString"),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				regexPatternSetReferenceStatement: &regexPatternSetReferenceStatementProperty{
//   					arn: jsii.String("arn"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				sizeConstraintStatement: &sizeConstraintStatementProperty{
//   					comparisonOperator: jsii.String("comparisonOperator"),
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					size: jsii.Number(123),
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				sqliMatchStatement: &sqliMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					sensitivityLevel: jsii.String("sensitivityLevel"),
//   				},
//   				xssMatchStatement: &xssMatchStatementProperty{
//   					fieldToMatch: &fieldToMatchProperty{
//   						allQueryArguments: allQueryArguments,
//   						body: &bodyProperty{
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						cookies: &cookiesProperty{
//   							matchPattern: &cookieMatchPatternProperty{
//   								all: all,
//   								excludedCookies: []*string{
//   									jsii.String("excludedCookies"),
//   								},
//   								includedCookies: []*string{
//   									jsii.String("includedCookies"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						headers: &headersProperty{
//   							matchPattern: &headerMatchPatternProperty{
//   								all: all,
//   								excludedHeaders: []*string{
//   									jsii.String("excludedHeaders"),
//   								},
//   								includedHeaders: []*string{
//   									jsii.String("includedHeaders"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						jsonBody: &jsonBodyProperty{
//   							matchPattern: &jsonMatchPatternProperty{
//   								all: all,
//   								includedPaths: []*string{
//   									jsii.String("includedPaths"),
//   								},
//   							},
//   							matchScope: jsii.String("matchScope"),
//
//   							// the properties below are optional
//   							invalidFallbackBehavior: jsii.String("invalidFallbackBehavior"),
//   							oversizeHandling: jsii.String("oversizeHandling"),
//   						},
//   						method: method,
//   						queryString: queryString,
//   						singleHeader: singleHeader,
//   						singleQueryArgument: singleQueryArgument,
//   						uriPath: uriPath,
//   					},
//   					textTransformations: []interface{}{
//   						&textTransformationProperty{
//   							priority: jsii.Number(123),
//   							type: jsii.String("type"),
//   						},
//   					},
//   				},
//   			},
//   			visibilityConfig: &visibilityConfigProperty{
//   				cloudWatchMetricsEnabled: jsii.Boolean(false),
//   				metricName: jsii.String("metricName"),
//   				sampledRequestsEnabled: jsii.Boolean(false),
//   			},
//
//   			// the properties below are optional
//   			action: &ruleActionProperty{
//   				allow: allow,
//   				block: block,
//   				captcha: captcha,
//   				count: count,
//   			},
//   			captchaConfig: &captchaConfigProperty{
//   				immunityTimeProperty: &immunityTimePropertyProperty{
//   					immunityTime: jsii.Number(123),
//   				},
//   			},
//   			ruleLabels: []interface{}{
//   				&labelProperty{
//   					name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   })
//
type CfnRuleGroup interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// The Amazon Resource Name (ARN) of the rule group.
	AttrArn() *string
	// Labels that rules in this rule group add to matching requests.
	//
	// These labels are defined in the `RuleLabels` for a `Rule` .
	AttrAvailableLabels() awscdk.IResolvable
	// Labels that rules in this rule group match against.
	//
	// Each of these labels is defined in a `LabelMatchStatement` specification, in the rule statement.
	AttrConsumedLabels() awscdk.IResolvable
	// The ID of the rule group.
	AttrId() *string
	// The label namespace prefix for this rule group.
	//
	// All labels added by rules in this rule group have this prefix.
	//
	// The syntax for the label namespace prefix for a rule group is the following: `awswaf:<account ID>:rule group:<rule group name>:`
	//
	// When a rule with a label matches a web request, AWS WAF adds the fully qualified label to the request. A fully qualified label is made up of the label namespace from the rule group or web ACL where the rule is defined and the label from the rule, separated by a colon.
	AttrLabelNamespace() *string
	// The web ACL capacity units (WCUs) required for this rule group.
	//
	// When you create your own rule group, you define this, and you cannot change it after creation. When you add or modify the rules in a rule group, AWS WAF enforces this limit.
	//
	// AWS WAF uses WCUs to calculate and control the operating resources that are used to run your rules, rule groups, and web ACLs. AWS WAF calculates capacity differently for each rule type, to reflect the relative cost of each rule. Simple rules that cost little to run use fewer WCUs than more complex rules that use more processing power. Rule group capacity is fixed at creation, which helps users plan their web ACL WCU usage when they use a rule group. The WCU limit for web ACLs is 1,500.
	Capacity() *float64
	SetCapacity(val *float64)
	// Options for this resource, such as condition, update policy etc.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	CfnResourceType() *string
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	CreationStack() *[]*string
	// A map of custom response keys and content bodies.
	//
	// When you create a rule with a block action, you can send a custom response to the web request. You define these for the rule group, and then use them in the rules that you define in the rule group.
	//
	// For information about customizing web requests and responses, see [Customizing web requests and responses in AWS WAF](https://docs.aws.amazon.com/waf/latest/developerguide/waf-custom-request-response.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	//
	// For information about the limits on count and size for custom request and response settings, see [AWS WAF quotas](https://docs.aws.amazon.com/waf/latest/developerguide/limits.html) in the [AWS WAF Developer Guide](https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html) .
	CustomResponseBodies() interface{}
	SetCustomResponseBodies(val interface{})
	// A description of the rule group that helps with identification.
	Description() *string
	SetDescription(val *string)
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
	// The name of the rule group.
	//
	// You cannot change the name of a rule group after you create it.
	Name() *string
	SetName(val *string)
	// The tree node.
	Node() constructs.Node
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	Ref() *string
	// The rule statements used to identify the web requests that you want to allow, block, or count.
	//
	// Each rule includes one top-level statement that AWS WAF uses to identify matching web requests, and parameters that govern how AWS WAF handles them.
	Rules() interface{}
	SetRules(val interface{})
	// Specifies whether this is for an Amazon CloudFront distribution or for a regional application.
	//
	// A regional application can be an Application Load Balancer (ALB), an Amazon API Gateway REST API, or an AWS AppSync GraphQL API. Valid Values are `CLOUDFRONT` and `REGIONAL` .
	//
	// > For `CLOUDFRONT` , you must create your WAFv2 resources in the US East (N. Virginia) Region, `us-east-1` .
	Scope() *string
	SetScope(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	Stack() awscdk.Stack
	// Key:value pairs associated with an AWS resource.
	//
	// The key:value pair can be anything you define. Typically, the tag key represents a category (such as "environment") and the tag value represents a specific value within that category (such as "test," "development," or "production"). You can add up to 50 tags to each AWS resource.
	//
	// > To modify tags on existing resources, use the AWS WAF APIs or command line interface. With AWS CloudFormation , you can only add tags to AWS WAF resources during resource creation.
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
	// Defines and enables Amazon CloudWatch metrics and web request sample collection.
	VisibilityConfig() interface{}
	SetVisibilityConfig(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
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
	//    "GlobalSecondaryIndexes": [
	//      {
	//        "Projection": {
	//          "NonKeyAttributes": [ "myattribute" ]
	//          ...
	//        }
	//        ...
	//      },
	//      {
	//        "ProjectionType": "INCLUDE"
	//        ...
	//      },
	//    ]
	//    ...
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
	GetAtt(attributeName *string) awscdk.Reference
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
	// Overrides the auto-generated logical ID with a specific ID.
	OverrideLogicalId(newLogicalId *string)
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
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

// The jsii proxy struct for CfnRuleGroup
type jsiiProxy_CfnRuleGroup struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnRuleGroup) AttrArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) AttrAvailableLabels() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrAvailableLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) AttrConsumedLabels() awscdk.IResolvable {
	var returns awscdk.IResolvable
	_jsii_.Get(
		j,
		"attrConsumedLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) AttrLabelNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrLabelNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Capacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) CustomResponseBodies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customResponseBodies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Rules() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) UpdatedProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnRuleGroup) VisibilityConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibilityConfig",
		&returns,
	)
	return returns
}


// Create a new `AWS::WAFv2::RuleGroup`.
func NewCfnRuleGroup(scope constructs.Construct, id *string, props *CfnRuleGroupProps) CfnRuleGroup {
	_init_.Initialize()

	j := jsiiProxy_CfnRuleGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_wafv2.CfnRuleGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::WAFv2::RuleGroup`.
func NewCfnRuleGroup_Override(c CfnRuleGroup, scope constructs.Construct, id *string, props *CfnRuleGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_wafv2.CfnRuleGroup",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetCapacity(val *float64) {
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetCustomResponseBodies(val interface{}) {
	_jsii_.Set(
		j,
		"customResponseBodies",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetRules(val interface{}) {
	_jsii_.Set(
		j,
		"rules",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetScope(val *string) {
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_CfnRuleGroup) SetVisibilityConfig(val interface{}) {
	_jsii_.Set(
		j,
		"visibilityConfig",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
func CfnRuleGroup_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnRuleGroup",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
func CfnRuleGroup_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnRuleGroup",
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
func CfnRuleGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_wafv2.CfnRuleGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnRuleGroup_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_wafv2.CfnRuleGroup",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnRuleGroup) AddDeletionOverride(path *string) {
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnRuleGroup) AddDependsOn(target awscdk.CfnResource) {
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnRuleGroup) AddMetadata(key *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnRuleGroup) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnRuleGroup) AddPropertyDeletionOverride(propertyPath *string) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnRuleGroup) AddPropertyOverride(propertyPath *string, value interface{}) {
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnRuleGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnRuleGroup) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleGroup) GetMetadata(key *string) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleGroup) Inspect(inspector awscdk.TreeInspector) {
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnRuleGroup) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnRuleGroup) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleGroup) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnRuleGroup) ValidateProperties(_properties interface{}) {
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}
