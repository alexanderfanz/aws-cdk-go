package awsssmincidents


// The `IncidentTemplate` property type specifies details used to create an incident when using this response plan.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   incidentTemplateProperty := &IncidentTemplateProperty{
//   	Impact: jsii.Number(123),
//   	Title: jsii.String("title"),
//
//   	// the properties below are optional
//   	DedupeString: jsii.String("dedupeString"),
//   	IncidentTags: []interface{}{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	NotificationTargets: []interface{}{
//   		&NotificationTargetItemProperty{
//   			SnsTopicArn: jsii.String("snsTopicArn"),
//   		},
//   	},
//   	Summary: jsii.String("summary"),
//   }
//
type CfnResponsePlan_IncidentTemplateProperty struct {
	// Defines the impact to the customers. Providing an impact overwrites the impact provided by a response plan.
	//
	// **Possible impacts:** - `1` - Critical impact, this typically relates to full application failure that impacts many to all customers.
	// - `2` - High impact, partial application failure with impact to many customers.
	// - `3` - Medium impact, the application is providing reduced service to customers.
	// - `4` - Low impact, customer might aren't impacted by the problem yet.
	// - `5` - No impact, customers aren't currently impacted but urgent action is needed to avoid impact.
	Impact *float64 `field:"required" json:"impact" yaml:"impact"`
	// The title of the incident is a brief and easily recognizable.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Used to create only one incident record for an incident.
	DedupeString *string `field:"optional" json:"dedupeString" yaml:"dedupeString"`
	// Tags to assign to the template.
	//
	// When the `StartIncident` API action is called, Incident Manager assigns the tags specified in the template to the incident.
	IncidentTags interface{} `field:"optional" json:"incidentTags" yaml:"incidentTags"`
	// The SNS targets that AWS Chatbot uses to notify the chat channel of updates to an incident.
	//
	// You can also make updates to the incident through the chat channel using the SNS topics.
	NotificationTargets interface{} `field:"optional" json:"notificationTargets" yaml:"notificationTargets"`
	// The summary describes what has happened during the incident.
	Summary *string `field:"optional" json:"summary" yaml:"summary"`
}
