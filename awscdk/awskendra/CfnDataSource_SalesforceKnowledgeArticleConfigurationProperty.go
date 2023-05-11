package awskendra


// Provides the configuration information for the knowledge article types that Amazon Kendra indexes.
//
// Amazon Kendra indexes standard knowledge articles and the standard fields of knowledge articles, or the custom fields of custom knowledge articles, but not both.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   salesforceKnowledgeArticleConfigurationProperty := &SalesforceKnowledgeArticleConfigurationProperty{
//   	IncludedStates: []*string{
//   		jsii.String("includedStates"),
//   	},
//
//   	// the properties below are optional
//   	CustomKnowledgeArticleTypeConfigurations: []interface{}{
//   		&SalesforceCustomKnowledgeArticleTypeConfigurationProperty{
//   			DocumentDataFieldName: jsii.String("documentDataFieldName"),
//   			Name: jsii.String("name"),
//
//   			// the properties below are optional
//   			DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   			FieldMappings: []interface{}{
//   				&DataSourceToIndexFieldMappingProperty{
//   					DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   					IndexFieldName: jsii.String("indexFieldName"),
//
//   					// the properties below are optional
//   					DateFieldFormat: jsii.String("dateFieldFormat"),
//   				},
//   			},
//   		},
//   	},
//   	StandardKnowledgeArticleTypeConfiguration: &SalesforceStandardKnowledgeArticleTypeConfigurationProperty{
//   		DocumentDataFieldName: jsii.String("documentDataFieldName"),
//
//   		// the properties below are optional
//   		DocumentTitleFieldName: jsii.String("documentTitleFieldName"),
//   		FieldMappings: []interface{}{
//   			&DataSourceToIndexFieldMappingProperty{
//   				DataSourceFieldName: jsii.String("dataSourceFieldName"),
//   				IndexFieldName: jsii.String("indexFieldName"),
//
//   				// the properties below are optional
//   				DateFieldFormat: jsii.String("dateFieldFormat"),
//   			},
//   		},
//   	},
//   }
//
type CfnDataSource_SalesforceKnowledgeArticleConfigurationProperty struct {
	// Specifies the document states that should be included when Amazon Kendra indexes knowledge articles.
	//
	// You must specify at least one state.
	IncludedStates *[]*string `field:"required" json:"includedStates" yaml:"includedStates"`
	// Configuration information for custom Salesforce knowledge articles.
	CustomKnowledgeArticleTypeConfigurations interface{} `field:"optional" json:"customKnowledgeArticleTypeConfigurations" yaml:"customKnowledgeArticleTypeConfigurations"`
	// Configuration information for standard Salesforce knowledge articles.
	StandardKnowledgeArticleTypeConfiguration interface{} `field:"optional" json:"standardKnowledgeArticleTypeConfiguration" yaml:"standardKnowledgeArticleTypeConfiguration"`
}
