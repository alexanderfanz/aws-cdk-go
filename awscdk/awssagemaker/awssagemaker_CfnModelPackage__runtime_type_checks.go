//go:build !no_runtime_type_checking

package awssagemaker

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/constructs-go/constructs/v3"
)

func (c *jsiiProxy_CfnModelPackage) validateAddDeletionOverrideParameters(path *string) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnModelPackage) validateAddDependsOnParameters(target awscdk.CfnResource) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnModelPackage) validateAddMetadataParameters(key *string, value interface{}) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnModelPackage) validateAddOverrideParameters(path *string, value interface{}) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnModelPackage) validateAddPropertyDeletionOverrideParameters(propertyPath *string) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnModelPackage) validateAddPropertyOverrideParameters(propertyPath *string, value interface{}) error {
	if propertyPath == nil {
		return fmt.Errorf("parameter propertyPath is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnModelPackage) validateApplyRemovalPolicyParameters(options *awscdk.RemovalPolicyOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CfnModelPackage) validateGetAttParameters(attributeName *string) error {
	if attributeName == nil {
		return fmt.Errorf("parameter attributeName is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnModelPackage) validateGetMetadataParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnModelPackage) validateInspectParameters(inspector awscdk.TreeInspector) error {
	if inspector == nil {
		return fmt.Errorf("parameter inspector is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnModelPackage) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnModelPackage) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	if newLogicalId == nil {
		return fmt.Errorf("parameter newLogicalId is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnModelPackage) validateRenderPropertiesParameters(props *map[string]interface{}) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnModelPackage) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	if session == nil {
		return fmt.Errorf("parameter session is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CfnModelPackage) validateValidatePropertiesParameters(_properties interface{}) error {
	if _properties == nil {
		return fmt.Errorf("parameter _properties is required, but nil was provided")
	}

	return nil
}

func validateCfnModelPackage_IsCfnElementParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateCfnModelPackage_IsCfnResourceParameters(construct constructs.IConstruct) error {
	if construct == nil {
		return fmt.Errorf("parameter construct is required, but nil was provided")
	}

	return nil
}

func validateCfnModelPackage_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_CfnModelPackage) validateSetAdditionalInferenceSpecificationDefinitionParameters(val interface{}) error {
	switch val.(type) {
	case *CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty:
		val := val.(*CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty:
		val_ := val.(CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnModelPackage) validateSetAdditionalInferenceSpecificationsParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *[]interface{}:
		val := val.(*[]interface{})
		for idx_97dfc6, v := range *val {
			switch v.(type) {
			case *CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty:
				v := v.(*CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty)
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty:
				v_ := v.(CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty)
				v := &v_
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case awscdk.IResolvable:
				// ok
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter val[%#v] must be one of the allowed types: *CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty, awscdk.IResolvable; received %#v (a %T)", idx_97dfc6, v, v)
				}
			}
		}
	case []interface{}:
		val_ := val.([]interface{})
		val := &val_
		for idx_97dfc6, v := range *val {
			switch v.(type) {
			case *CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty:
				v := v.(*CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty)
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty:
				v_ := v.(CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty)
				v := &v_
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case awscdk.IResolvable:
				// ok
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter val[%#v] must be one of the allowed types: *CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty, awscdk.IResolvable; received %#v (a %T)", idx_97dfc6, v, v)
				}
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *[]interface{}; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnModelPackage) validateSetAdditionalInferenceSpecificationsToAddParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *[]interface{}:
		val := val.(*[]interface{})
		for idx_97dfc6, v := range *val {
			switch v.(type) {
			case *CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty:
				v := v.(*CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty)
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty:
				v_ := v.(CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty)
				v := &v_
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case awscdk.IResolvable:
				// ok
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter val[%#v] must be one of the allowed types: *CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty, awscdk.IResolvable; received %#v (a %T)", idx_97dfc6, v, v)
				}
			}
		}
	case []interface{}:
		val_ := val.([]interface{})
		val := &val_
		for idx_97dfc6, v := range *val {
			switch v.(type) {
			case *CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty:
				v := v.(*CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty)
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty:
				v_ := v.(CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty)
				v := &v_
				if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
					return err
				}
			case awscdk.IResolvable:
				// ok
			default:
				if !_jsii_.IsAnonymousProxy(v) {
					return fmt.Errorf("parameter val[%#v] must be one of the allowed types: *CfnModelPackage_AdditionalInferenceSpecificationDefinitionProperty, awscdk.IResolvable; received %#v (a %T)", idx_97dfc6, v, v)
				}
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *[]interface{}; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnModelPackage) validateSetCertifyForMarketplaceParameters(val interface{}) error {
	switch val.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *bool, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnModelPackage) validateSetCreatedByParameters(val interface{}) error {
	switch val.(type) {
	case *CfnModelPackage_UserContextProperty:
		val := val.(*CfnModelPackage_UserContextProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnModelPackage_UserContextProperty:
		val_ := val.(CfnModelPackage_UserContextProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnModelPackage_UserContextProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnModelPackage) validateSetCustomerMetadataPropertiesParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *map[string]*string:
		// ok
	case map[string]*string:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *map[string]*string; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnModelPackage) validateSetDriftCheckBaselinesParameters(val interface{}) error {
	switch val.(type) {
	case *CfnModelPackage_DriftCheckBaselinesProperty:
		val := val.(*CfnModelPackage_DriftCheckBaselinesProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnModelPackage_DriftCheckBaselinesProperty:
		val_ := val.(CfnModelPackage_DriftCheckBaselinesProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnModelPackage_DriftCheckBaselinesProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnModelPackage) validateSetEnvironmentParameters(val interface{}) error {
	switch val.(type) {
	case awscdk.IResolvable:
		// ok
	case *map[string]*string:
		// ok
	case map[string]*string:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: awscdk.IResolvable, *map[string]*string; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnModelPackage) validateSetInferenceSpecificationParameters(val interface{}) error {
	switch val.(type) {
	case *CfnModelPackage_InferenceSpecificationProperty:
		val := val.(*CfnModelPackage_InferenceSpecificationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnModelPackage_InferenceSpecificationProperty:
		val_ := val.(CfnModelPackage_InferenceSpecificationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnModelPackage_InferenceSpecificationProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnModelPackage) validateSetLastModifiedByParameters(val interface{}) error {
	switch val.(type) {
	case *CfnModelPackage_UserContextProperty:
		val := val.(*CfnModelPackage_UserContextProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnModelPackage_UserContextProperty:
		val_ := val.(CfnModelPackage_UserContextProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnModelPackage_UserContextProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnModelPackage) validateSetMetadataPropertiesParameters(val interface{}) error {
	switch val.(type) {
	case *CfnModelPackage_MetadataPropertiesProperty:
		val := val.(*CfnModelPackage_MetadataPropertiesProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnModelPackage_MetadataPropertiesProperty:
		val_ := val.(CfnModelPackage_MetadataPropertiesProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnModelPackage_MetadataPropertiesProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnModelPackage) validateSetModelMetricsParameters(val interface{}) error {
	switch val.(type) {
	case *CfnModelPackage_ModelMetricsProperty:
		val := val.(*CfnModelPackage_ModelMetricsProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnModelPackage_ModelMetricsProperty:
		val_ := val.(CfnModelPackage_ModelMetricsProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnModelPackage_ModelMetricsProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnModelPackage) validateSetModelPackageStatusDetailsParameters(val interface{}) error {
	switch val.(type) {
	case *CfnModelPackage_ModelPackageStatusDetailsProperty:
		val := val.(*CfnModelPackage_ModelPackageStatusDetailsProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnModelPackage_ModelPackageStatusDetailsProperty:
		val_ := val.(CfnModelPackage_ModelPackageStatusDetailsProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnModelPackage_ModelPackageStatusDetailsProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnModelPackage) validateSetModelPackageStatusItemParameters(val interface{}) error {
	switch val.(type) {
	case *CfnModelPackage_ModelPackageStatusItemProperty:
		val := val.(*CfnModelPackage_ModelPackageStatusItemProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnModelPackage_ModelPackageStatusItemProperty:
		val_ := val.(CfnModelPackage_ModelPackageStatusItemProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnModelPackage_ModelPackageStatusItemProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnModelPackage) validateSetSourceAlgorithmSpecificationParameters(val interface{}) error {
	switch val.(type) {
	case *CfnModelPackage_SourceAlgorithmSpecificationProperty:
		val := val.(*CfnModelPackage_SourceAlgorithmSpecificationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnModelPackage_SourceAlgorithmSpecificationProperty:
		val_ := val.(CfnModelPackage_SourceAlgorithmSpecificationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnModelPackage_SourceAlgorithmSpecificationProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_CfnModelPackage) validateSetValidationSpecificationParameters(val interface{}) error {
	switch val.(type) {
	case *CfnModelPackage_ValidationSpecificationProperty:
		val := val.(*CfnModelPackage_ValidationSpecificationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case CfnModelPackage_ValidationSpecificationProperty:
		val_ := val.(CfnModelPackage_ValidationSpecificationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case awscdk.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *CfnModelPackage_ValidationSpecificationProperty, awscdk.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func validateNewCfnModelPackageParameters(scope awscdk.Construct, id *string, props *CfnModelPackageProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}
