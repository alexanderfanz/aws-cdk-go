package awsomics


// The store's file parsing options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var schema interface{}
//
//   storeOptionsProperty := &storeOptionsProperty{
//   	tsvStoreOptions: &tsvStoreOptionsProperty{
//   		annotationType: jsii.String("annotationType"),
//   		formatToHeader: map[string]*string{
//   			"formatToHeaderKey": jsii.String("formatToHeader"),
//   		},
//   		schema: schema,
//   	},
//   }
//
type CfnAnnotationStore_StoreOptionsProperty struct {
	// Formatting options for a TSV file.
	TsvStoreOptions interface{} `field:"required" json:"tsvStoreOptions" yaml:"tsvStoreOptions"`
}
