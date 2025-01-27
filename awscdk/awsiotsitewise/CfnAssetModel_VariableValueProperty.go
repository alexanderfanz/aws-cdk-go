package awsiotsitewise


// Identifies a property value used in an expression.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   variableValueProperty := &VariableValueProperty{
//   	PropertyLogicalId: jsii.String("propertyLogicalId"),
//
//   	// the properties below are optional
//   	HierarchyLogicalId: jsii.String("hierarchyLogicalId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-variablevalue.html
//
type CfnAssetModel_VariableValueProperty struct {
	// The `LogicalID` of the property to use as the variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-variablevalue.html#cfn-iotsitewise-assetmodel-variablevalue-propertylogicalid
	//
	PropertyLogicalId *string `field:"required" json:"propertyLogicalId" yaml:"propertyLogicalId"`
	// The `LogicalID` of the hierarchy to query for the `PropertyLogicalID` .
	//
	// You use a `hierarchyLogicalID` instead of a model ID because you can have several hierarchies using the same model and therefore the same property. For example, you might have separately grouped assets that come from the same asset model. For more information, see [Defining relationships between assets](https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html) in the *AWS IoT SiteWise User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotsitewise-assetmodel-variablevalue.html#cfn-iotsitewise-assetmodel-variablevalue-hierarchylogicalid
	//
	HierarchyLogicalId *string `field:"optional" json:"hierarchyLogicalId" yaml:"hierarchyLogicalId"`
}

