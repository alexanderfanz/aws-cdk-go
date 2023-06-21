package awssecurityhub


// Properties for defining a `CfnAutomationRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var id interface{}
//   var updatedBy interface{}
//
//   cfnAutomationRuleProps := &CfnAutomationRuleProps{
//   	Actions: []interface{}{
//   		&AutomationRulesActionProperty{
//   			FindingFieldsUpdate: &AutomationRulesFindingFieldsUpdateProperty{
//   				Confidence: jsii.Number(123),
//   				Criticality: jsii.Number(123),
//   				Note: &NoteUpdateProperty{
//   					Text: jsii.String("text"),
//   					UpdatedBy: updatedBy,
//   				},
//   				RelatedFindings: []interface{}{
//   					&RelatedFindingProperty{
//   						Id: id,
//   						ProductArn: jsii.String("productArn"),
//   					},
//   				},
//   				Severity: &SeverityUpdateProperty{
//   					Label: jsii.String("label"),
//   					Normalized: jsii.Number(123),
//   					Product: jsii.Number(123),
//   				},
//   				Types: []*string{
//   					jsii.String("types"),
//   				},
//   				UserDefinedFields: map[string]*string{
//   					"userDefinedFieldsKey": jsii.String("userDefinedFields"),
//   				},
//   				VerificationState: jsii.String("verificationState"),
//   				Workflow: &WorkflowUpdateProperty{
//   					Status: jsii.String("status"),
//   				},
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Criteria: &AutomationRulesFindingFiltersProperty{
//   		AwsAccountId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		CompanyName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ComplianceAssociatedStandardsId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ComplianceSecurityControlId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ComplianceStatus: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Confidence: []interface{}{
//   			&NumberFilterProperty{
//   				Eq: jsii.Number(123),
//   				Gte: jsii.Number(123),
//   				Lte: jsii.Number(123),
//   			},
//   		},
//   		CreatedAt: []interface{}{
//   			&DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   		Criticality: []interface{}{
//   			&NumberFilterProperty{
//   				Eq: jsii.Number(123),
//   				Gte: jsii.Number(123),
//   				Lte: jsii.Number(123),
//   			},
//   		},
//   		Description: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		FirstObservedAt: []interface{}{
//   			&DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   		GeneratorId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Id: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		LastObservedAt: []interface{}{
//   			&DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   		NoteText: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		NoteUpdatedAt: []interface{}{
//   			&DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   		NoteUpdatedBy: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ProductArn: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ProductName: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		RecordState: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		RelatedFindingsId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		RelatedFindingsProductArn: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceDetailsOther: []interface{}{
//   			&MapFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceId: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourcePartition: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceRegion: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceTags: []interface{}{
//   			&MapFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		ResourceType: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		SeverityLabel: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		SourceUrl: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Title: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Type: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		UpdatedAt: []interface{}{
//   			&DateFilterProperty{
//   				DateRange: &DateRangeProperty{
//   					Unit: jsii.String("unit"),
//   					Value: jsii.Number(123),
//   				},
//   				End: jsii.String("end"),
//   				Start: jsii.String("start"),
//   			},
//   		},
//   		UserDefinedFields: []interface{}{
//   			&MapFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		VerificationState: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		WorkflowStatus: []interface{}{
//   			&StringFilterProperty{
//   				Comparison: jsii.String("comparison"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	IsTerminal: jsii.Boolean(false),
//   	RuleName: jsii.String("ruleName"),
//   	RuleOrder: jsii.Number(123),
//   	RuleStatus: jsii.String("ruleStatus"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnAutomationRuleProps struct {
	// One or more actions to update finding fields if a finding matches the defined criteria of the rule.
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// A set of [AWS Security Finding Format](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-findings-format.html) finding field attributes and corresponding expected values that Security Hub uses to filter findings. If a finding matches the conditions specified in this parameter, Security Hub applies the rule action to the finding.
	Criteria interface{} `field:"optional" json:"criteria" yaml:"criteria"`
	// A description of the rule.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether a rule is the last to be applied with respect to a finding that matches the rule criteria.
	//
	// This is useful when a finding matches the criteria for multiple rules, and each rule has different actions. If the value of this field is set to `true` for a rule, Security Hub applies the rule action to a finding that matches the rule criteria and won't evaluate other rules for the finding. The default value of this field is `false` .
	IsTerminal interface{} `field:"optional" json:"isTerminal" yaml:"isTerminal"`
	// The name of the rule.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// An integer ranging from 1 to 1000 that represents the order in which the rule action is applied to findings.
	//
	// Security Hub applies rules with lower values for this parameter first.
	RuleOrder *float64 `field:"optional" json:"ruleOrder" yaml:"ruleOrder"`
	// Whether the rule is active after it is created.
	//
	// If this parameter is equal to `ENABLED` , Security Hub will apply the rule to findings and finding updates after the rule is created.
	RuleStatus *string `field:"optional" json:"ruleStatus" yaml:"ruleStatus"`
	// User-defined tags that help you label the purpose of a rule.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

