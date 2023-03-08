package awsiottwinmaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataTypeProperty_ dataTypeProperty
//   var dataValueProperty_ dataValueProperty
//   var relationshipValue interface{}
//
//   dataTypeProperty := &dataTypeProperty{
//   	AllowedValues: []interface{}{
//   		&dataValueProperty{
//   			BooleanValue: jsii.Boolean(false),
//   			DoubleValue: jsii.Number(123),
//   			Expression: jsii.String("expression"),
//   			IntegerValue: jsii.Number(123),
//   			ListValue: []interface{}{
//   				dataValueProperty_,
//   			},
//   			LongValue: jsii.Number(123),
//   			MapValue: map[string]interface{}{
//   				"mapValueKey": dataValueProperty_,
//   			},
//   			RelationshipValue: relationshipValue,
//   			StringValue: jsii.String("stringValue"),
//   		},
//   	},
//   	NestedType: &dataTypeProperty{
//   		AllowedValues: []interface{}{
//   			&dataValueProperty{
//   				BooleanValue: jsii.Boolean(false),
//   				DoubleValue: jsii.Number(123),
//   				Expression: jsii.String("expression"),
//   				IntegerValue: jsii.Number(123),
//   				ListValue: []interface{}{
//   					dataValueProperty_,
//   				},
//   				LongValue: jsii.Number(123),
//   				MapValue: map[string]interface{}{
//   					"mapValueKey": dataValueProperty_,
//   				},
//   				RelationshipValue: relationshipValue,
//   				StringValue: jsii.String("stringValue"),
//   			},
//   		},
//   		NestedType: dataTypeProperty_,
//   		Relationship: &RelationshipProperty{
//   			RelationshipType: jsii.String("relationshipType"),
//   			TargetComponentTypeId: jsii.String("targetComponentTypeId"),
//   		},
//   		Type: jsii.String("type"),
//   		UnitOfMeasure: jsii.String("unitOfMeasure"),
//   	},
//   	Relationship: &RelationshipProperty{
//   		RelationshipType: jsii.String("relationshipType"),
//   		TargetComponentTypeId: jsii.String("targetComponentTypeId"),
//   	},
//   	Type: jsii.String("type"),
//   	UnitOfMeasure: jsii.String("unitOfMeasure"),
//   }
//
type CfnEntity_DataTypeProperty struct {
	// `CfnEntity.DataTypeProperty.AllowedValues`.
	AllowedValues interface{} `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// `CfnEntity.DataTypeProperty.NestedType`.
	NestedType interface{} `field:"optional" json:"nestedType" yaml:"nestedType"`
	// `CfnEntity.DataTypeProperty.Relationship`.
	Relationship interface{} `field:"optional" json:"relationship" yaml:"relationship"`
	// `CfnEntity.DataTypeProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// `CfnEntity.DataTypeProperty.UnitOfMeasure`.
	UnitOfMeasure *string `field:"optional" json:"unitOfMeasure" yaml:"unitOfMeasure"`
}

