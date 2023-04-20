package awsamplifyuibuilder

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsamplifyuibuilder/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// A CloudFormation `AWS::AmplifyUIBuilder::Component`.
//
// The AWS::AmplifyUIBuilder::Component resource specifies a component within an Amplify app. A component is a user interface (UI) element that you can customize. Use `ComponentChild` to configure an instance of a `Component` . A `ComponentChild` instance inherits the configuration of the main `Component` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var componentChildProperty_ componentChildProperty
//   var componentPropertyProperty_ componentPropertyProperty
//   var overrides interface{}
//   var predicateProperty_ predicateProperty
//
//   cfnComponent := awscdk.Aws_amplifyuibuilder.NewCfnComponent(this, jsii.String("MyCfnComponent"), &CfnComponentProps{
//   	BindingProperties: map[string]interface{}{
//   		"bindingPropertiesKey": &ComponentBindingPropertiesValueProperty{
//   			"bindingProperties": &ComponentBindingPropertiesValuePropertiesProperty{
//   				"bucket": jsii.String("bucket"),
//   				"defaultValue": jsii.String("defaultValue"),
//   				"field": jsii.String("field"),
//   				"key": jsii.String("key"),
//   				"model": jsii.String("model"),
//   				"predicates": []interface{}{
//   					&predicateProperty{
//   						"and": []interface{}{
//   							predicateProperty_,
//   						},
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operator": jsii.String("operator"),
//   						"or": []interface{}{
//   							predicateProperty_,
//   						},
//   					},
//   				},
//   				"userAttribute": jsii.String("userAttribute"),
//   			},
//   			"defaultValue": jsii.String("defaultValue"),
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	ComponentType: jsii.String("componentType"),
//   	Name: jsii.String("name"),
//   	Overrides: overrides,
//   	Properties: map[string]interface{}{
//   		"propertiesKey": &componentPropertyProperty{
//   			"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   				"property": jsii.String("property"),
//
//   				// the properties below are optional
//   				"field": jsii.String("field"),
//   			},
//   			"bindings": map[string]interface{}{
//   				"bindingsKey": &FormBindingElementProperty{
//   					"element": jsii.String("element"),
//   					"property": jsii.String("property"),
//   				},
//   			},
//   			"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   				"property": jsii.String("property"),
//
//   				// the properties below are optional
//   				"field": jsii.String("field"),
//   			},
//   			"componentName": jsii.String("componentName"),
//   			"concat": []interface{}{
//   				componentPropertyProperty_,
//   			},
//   			"condition": &ComponentConditionPropertyProperty{
//   				"else": componentPropertyProperty_,
//   				"field": jsii.String("field"),
//   				"operand": jsii.String("operand"),
//   				"operandType": jsii.String("operandType"),
//   				"operator": jsii.String("operator"),
//   				"property": jsii.String("property"),
//   				"then": componentPropertyProperty_,
//   			},
//   			"configured": jsii.Boolean(false),
//   			"defaultValue": jsii.String("defaultValue"),
//   			"event": jsii.String("event"),
//   			"importedValue": jsii.String("importedValue"),
//   			"model": jsii.String("model"),
//   			"property": jsii.String("property"),
//   			"type": jsii.String("type"),
//   			"userAttribute": jsii.String("userAttribute"),
//   			"value": jsii.String("value"),
//   		},
//   	},
//   	Variants: []interface{}{
//   		&ComponentVariantProperty{
//   			Overrides: overrides,
//   			VariantValues: map[string]*string{
//   				"variantValuesKey": jsii.String("variantValues"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	AppId: jsii.String("appId"),
//   	Children: []interface{}{
//   		&componentChildProperty{
//   			ComponentType: jsii.String("componentType"),
//   			Name: jsii.String("name"),
//   			Properties: map[string]interface{}{
//   				"propertiesKey": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Children: []interface{}{
//   				componentChildProperty_,
//   			},
//   			Events: map[string]interface{}{
//   				"eventsKey": &ComponentEventProperty{
//   					"action": jsii.String("action"),
//   					"parameters": &ActionParametersProperty{
//   						"anchor": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   						"fields": map[string]interface{}{
//   							"fieldsKey": &componentPropertyProperty{
//   								"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   									"property": jsii.String("property"),
//
//   									// the properties below are optional
//   									"field": jsii.String("field"),
//   								},
//   								"bindings": map[string]interface{}{
//   									"bindingsKey": &FormBindingElementProperty{
//   										"element": jsii.String("element"),
//   										"property": jsii.String("property"),
//   									},
//   								},
//   								"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   									"property": jsii.String("property"),
//
//   									// the properties below are optional
//   									"field": jsii.String("field"),
//   								},
//   								"componentName": jsii.String("componentName"),
//   								"concat": []interface{}{
//   									componentPropertyProperty_,
//   								},
//   								"condition": &ComponentConditionPropertyProperty{
//   									"else": componentPropertyProperty_,
//   									"field": jsii.String("field"),
//   									"operand": jsii.String("operand"),
//   									"operandType": jsii.String("operandType"),
//   									"operator": jsii.String("operator"),
//   									"property": jsii.String("property"),
//   									"then": componentPropertyProperty_,
//   								},
//   								"configured": jsii.Boolean(false),
//   								"defaultValue": jsii.String("defaultValue"),
//   								"event": jsii.String("event"),
//   								"importedValue": jsii.String("importedValue"),
//   								"model": jsii.String("model"),
//   								"property": jsii.String("property"),
//   								"type": jsii.String("type"),
//   								"userAttribute": jsii.String("userAttribute"),
//   								"value": jsii.String("value"),
//   							},
//   						},
//   						"global": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   						"id": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   						"model": jsii.String("model"),
//   						"state": &MutationActionSetStateParameterProperty{
//   							"componentName": jsii.String("componentName"),
//   							"property": jsii.String("property"),
//   							"set": &componentPropertyProperty{
//   								"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   									"property": jsii.String("property"),
//
//   									// the properties below are optional
//   									"field": jsii.String("field"),
//   								},
//   								"bindings": map[string]interface{}{
//   									"bindingsKey": &FormBindingElementProperty{
//   										"element": jsii.String("element"),
//   										"property": jsii.String("property"),
//   									},
//   								},
//   								"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   									"property": jsii.String("property"),
//
//   									// the properties below are optional
//   									"field": jsii.String("field"),
//   								},
//   								"componentName": jsii.String("componentName"),
//   								"concat": []interface{}{
//   									componentPropertyProperty_,
//   								},
//   								"condition": &ComponentConditionPropertyProperty{
//   									"else": componentPropertyProperty_,
//   									"field": jsii.String("field"),
//   									"operand": jsii.String("operand"),
//   									"operandType": jsii.String("operandType"),
//   									"operator": jsii.String("operator"),
//   									"property": jsii.String("property"),
//   									"then": componentPropertyProperty_,
//   								},
//   								"configured": jsii.Boolean(false),
//   								"defaultValue": jsii.String("defaultValue"),
//   								"event": jsii.String("event"),
//   								"importedValue": jsii.String("importedValue"),
//   								"model": jsii.String("model"),
//   								"property": jsii.String("property"),
//   								"type": jsii.String("type"),
//   								"userAttribute": jsii.String("userAttribute"),
//   								"value": jsii.String("value"),
//   							},
//   						},
//   						"target": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   						"type": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   						"url": &componentPropertyProperty{
//   							"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"bindings": map[string]interface{}{
//   								"bindingsKey": &FormBindingElementProperty{
//   									"element": jsii.String("element"),
//   									"property": jsii.String("property"),
//   								},
//   							},
//   							"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   								"property": jsii.String("property"),
//
//   								// the properties below are optional
//   								"field": jsii.String("field"),
//   							},
//   							"componentName": jsii.String("componentName"),
//   							"concat": []interface{}{
//   								componentPropertyProperty_,
//   							},
//   							"condition": &ComponentConditionPropertyProperty{
//   								"else": componentPropertyProperty_,
//   								"field": jsii.String("field"),
//   								"operand": jsii.String("operand"),
//   								"operandType": jsii.String("operandType"),
//   								"operator": jsii.String("operator"),
//   								"property": jsii.String("property"),
//   								"then": componentPropertyProperty_,
//   							},
//   							"configured": jsii.Boolean(false),
//   							"defaultValue": jsii.String("defaultValue"),
//   							"event": jsii.String("event"),
//   							"importedValue": jsii.String("importedValue"),
//   							"model": jsii.String("model"),
//   							"property": jsii.String("property"),
//   							"type": jsii.String("type"),
//   							"userAttribute": jsii.String("userAttribute"),
//   							"value": jsii.String("value"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   	CollectionProperties: map[string]interface{}{
//   		"collectionPropertiesKey": &ComponentDataConfigurationProperty{
//   			"model": jsii.String("model"),
//
//   			// the properties below are optional
//   			"identifiers": []*string{
//   				jsii.String("identifiers"),
//   			},
//   			"predicate": &predicateProperty{
//   				"and": []interface{}{
//   					predicateProperty_,
//   				},
//   				"field": jsii.String("field"),
//   				"operand": jsii.String("operand"),
//   				"operator": jsii.String("operator"),
//   				"or": []interface{}{
//   					predicateProperty_,
//   				},
//   			},
//   			"sort": []interface{}{
//   				&SortPropertyProperty{
//   					"direction": jsii.String("direction"),
//   					"field": jsii.String("field"),
//   				},
//   			},
//   		},
//   	},
//   	EnvironmentName: jsii.String("environmentName"),
//   	Events: map[string]interface{}{
//   		"eventsKey": &ComponentEventProperty{
//   			"action": jsii.String("action"),
//   			"parameters": &ActionParametersProperty{
//   				"anchor": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"fields": map[string]interface{}{
//   					"fieldsKey": &componentPropertyProperty{
//   						"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   							"property": jsii.String("property"),
//
//   							// the properties below are optional
//   							"field": jsii.String("field"),
//   						},
//   						"bindings": map[string]interface{}{
//   							"bindingsKey": &FormBindingElementProperty{
//   								"element": jsii.String("element"),
//   								"property": jsii.String("property"),
//   							},
//   						},
//   						"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   							"property": jsii.String("property"),
//
//   							// the properties below are optional
//   							"field": jsii.String("field"),
//   						},
//   						"componentName": jsii.String("componentName"),
//   						"concat": []interface{}{
//   							componentPropertyProperty_,
//   						},
//   						"condition": &ComponentConditionPropertyProperty{
//   							"else": componentPropertyProperty_,
//   							"field": jsii.String("field"),
//   							"operand": jsii.String("operand"),
//   							"operandType": jsii.String("operandType"),
//   							"operator": jsii.String("operator"),
//   							"property": jsii.String("property"),
//   							"then": componentPropertyProperty_,
//   						},
//   						"configured": jsii.Boolean(false),
//   						"defaultValue": jsii.String("defaultValue"),
//   						"event": jsii.String("event"),
//   						"importedValue": jsii.String("importedValue"),
//   						"model": jsii.String("model"),
//   						"property": jsii.String("property"),
//   						"type": jsii.String("type"),
//   						"userAttribute": jsii.String("userAttribute"),
//   						"value": jsii.String("value"),
//   					},
//   				},
//   				"global": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"id": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"model": jsii.String("model"),
//   				"state": &MutationActionSetStateParameterProperty{
//   					"componentName": jsii.String("componentName"),
//   					"property": jsii.String("property"),
//   					"set": &componentPropertyProperty{
//   						"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   							"property": jsii.String("property"),
//
//   							// the properties below are optional
//   							"field": jsii.String("field"),
//   						},
//   						"bindings": map[string]interface{}{
//   							"bindingsKey": &FormBindingElementProperty{
//   								"element": jsii.String("element"),
//   								"property": jsii.String("property"),
//   							},
//   						},
//   						"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   							"property": jsii.String("property"),
//
//   							// the properties below are optional
//   							"field": jsii.String("field"),
//   						},
//   						"componentName": jsii.String("componentName"),
//   						"concat": []interface{}{
//   							componentPropertyProperty_,
//   						},
//   						"condition": &ComponentConditionPropertyProperty{
//   							"else": componentPropertyProperty_,
//   							"field": jsii.String("field"),
//   							"operand": jsii.String("operand"),
//   							"operandType": jsii.String("operandType"),
//   							"operator": jsii.String("operator"),
//   							"property": jsii.String("property"),
//   							"then": componentPropertyProperty_,
//   						},
//   						"configured": jsii.Boolean(false),
//   						"defaultValue": jsii.String("defaultValue"),
//   						"event": jsii.String("event"),
//   						"importedValue": jsii.String("importedValue"),
//   						"model": jsii.String("model"),
//   						"property": jsii.String("property"),
//   						"type": jsii.String("type"),
//   						"userAttribute": jsii.String("userAttribute"),
//   						"value": jsii.String("value"),
//   					},
//   				},
//   				"target": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"type": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   				"url": &componentPropertyProperty{
//   					"bindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"bindings": map[string]interface{}{
//   						"bindingsKey": &FormBindingElementProperty{
//   							"element": jsii.String("element"),
//   							"property": jsii.String("property"),
//   						},
//   					},
//   					"collectionBindingProperties": &ComponentPropertyBindingPropertiesProperty{
//   						"property": jsii.String("property"),
//
//   						// the properties below are optional
//   						"field": jsii.String("field"),
//   					},
//   					"componentName": jsii.String("componentName"),
//   					"concat": []interface{}{
//   						componentPropertyProperty_,
//   					},
//   					"condition": &ComponentConditionPropertyProperty{
//   						"else": componentPropertyProperty_,
//   						"field": jsii.String("field"),
//   						"operand": jsii.String("operand"),
//   						"operandType": jsii.String("operandType"),
//   						"operator": jsii.String("operator"),
//   						"property": jsii.String("property"),
//   						"then": componentPropertyProperty_,
//   					},
//   					"configured": jsii.Boolean(false),
//   					"defaultValue": jsii.String("defaultValue"),
//   					"event": jsii.String("event"),
//   					"importedValue": jsii.String("importedValue"),
//   					"model": jsii.String("model"),
//   					"property": jsii.String("property"),
//   					"type": jsii.String("type"),
//   					"userAttribute": jsii.String("userAttribute"),
//   					"value": jsii.String("value"),
//   				},
//   			},
//   		},
//   	},
//   	SchemaVersion: jsii.String("schemaVersion"),
//   	SourceId: jsii.String("sourceId"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   })
//
type CfnComponent interface {
	awscdk.CfnResource
	awscdk.IInspectable
	// `AWS::AmplifyUIBuilder::Component.AppId`.
	AppId() *string
	SetAppId(val *string)
	// The unique ID of the component.
	AttrId() *string
	// The information to connect a component's properties to data at runtime.
	//
	// You can't specify `tags` as a valid property for `bindingProperties` .
	BindingProperties() interface{}
	SetBindingProperties(val interface{})
	// Options for this resource, such as condition, update policy etc.
	// Experimental.
	CfnOptions() awscdk.ICfnResourceOptions
	CfnProperties() *map[string]interface{}
	// AWS resource type.
	// Experimental.
	CfnResourceType() *string
	// A list of the component's `ComponentChild` instances.
	Children() interface{}
	SetChildren(val interface{})
	// The data binding configuration for the component's properties.
	//
	// Use this for a collection component. You can't specify `tags` as a valid property for `collectionProperties` .
	CollectionProperties() interface{}
	SetCollectionProperties(val interface{})
	// The type of the component.
	//
	// This can be an Amplify custom UI component or another custom component.
	ComponentType() *string
	SetComponentType(val *string)
	// Returns: the stack trace of the point where this Resource was created from, sourced
	// from the +metadata+ entry typed +aws:cdk:logicalId+, and with the bottom-most
	// node +internal+ entries filtered.
	// Experimental.
	CreationStack() *[]*string
	// `AWS::AmplifyUIBuilder::Component.EnvironmentName`.
	EnvironmentName() *string
	SetEnvironmentName(val *string)
	// Describes the events that can be raised on the component.
	//
	// Use for the workflow feature in Amplify Studio that allows you to bind events and actions to components.
	Events() interface{}
	SetEvents(val interface{})
	// The logical ID for this CloudFormation stack element.
	//
	// The logical ID of the element
	// is calculated from the path of the resource node in the construct tree.
	//
	// To override this value, use `overrideLogicalId(newLogicalId)`.
	//
	// Returns: the logical ID as a stringified token. This value will only get
	// resolved during synthesis.
	// Experimental.
	LogicalId() *string
	// The name of the component.
	Name() *string
	SetName(val *string)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Describes the component's properties that can be overriden in a customized instance of the component.
	//
	// You can't specify `tags` as a valid property for `overrides` .
	Overrides() interface{}
	SetOverrides(val interface{})
	// Describes the component's properties.
	//
	// You can't specify `tags` as a valid property for `properties` .
	Properties() interface{}
	SetProperties(val interface{})
	// Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.
	//
	// If, by any chance, the intrinsic reference of a resource is not a string, you could
	// coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.
	// Experimental.
	Ref() *string
	// `AWS::AmplifyUIBuilder::Component.SchemaVersion`.
	SchemaVersion() *string
	SetSchemaVersion(val *string)
	// The unique ID of the component in its original source system, such as Figma.
	SourceId() *string
	SetSourceId(val *string)
	// The stack in which this element is defined.
	//
	// CfnElements must be defined within a stack scope (directly or indirectly).
	// Experimental.
	Stack() awscdk.Stack
	// One or more key-value pairs to use when tagging the component.
	Tags() awscdk.TagManager
	// Return properties modified after initiation.
	//
	// Resources that expose mutable properties should override this function to
	// collect and return the properties object for this resource.
	// Experimental.
	UpdatedProperites() *map[string]interface{}
	// A list of the component's variants.
	//
	// A variant is a unique style configuration of a main component.
	Variants() interface{}
	SetVariants(val interface{})
	// Syntactic sugar for `addOverride(path, undefined)`.
	// Experimental.
	AddDeletionOverride(path *string)
	// Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.
	//
	// This can be used for resources across stacks (or nested stack) boundaries
	// and the dependency will automatically be transferred to the relevant scope.
	// Experimental.
	AddDependsOn(target awscdk.CfnResource)
	// Add a value to the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
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
	// Experimental.
	AddOverride(path *string, value interface{})
	// Adds an override that deletes the value of a property from the resource definition.
	// Experimental.
	AddPropertyDeletionOverride(propertyPath *string)
	// Adds an override to a resource property.
	//
	// Syntactic sugar for `addOverride("Properties.<...>", value)`.
	// Experimental.
	AddPropertyOverride(propertyPath *string, value interface{})
	// Sets the deletion policy of the resource based on the removal policy specified.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	// Experimental.
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions)
	// Returns a token for an runtime attribute of this resource.
	//
	// Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
	// in case there is no generated attribute.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Retrieve a value value from the CloudFormation Resource Metadata.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html
	//
	// Note that this is a different set of metadata from CDK node metadata; this
	// metadata ends up in the stack template under the resource, whereas CDK
	// node metadata ends up in the Cloud Assembly.
	//
	// Experimental.
	GetMetadata(key *string) interface{}
	// Examines the CloudFormation resource and discloses attributes.
	Inspect(inspector awscdk.TreeInspector)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	RenderProperties(props *map[string]interface{}) *map[string]interface{}
	// Can be overridden by subclasses to determine if this resource will be rendered into the cloudformation template.
	//
	// Returns: `true` if the resource should be included or `false` is the resource
	// should be omitted.
	// Experimental.
	ShouldSynthesize() *bool
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	//
	// Returns: a string representation of this resource.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Experimental.
	ValidateProperties(_properties interface{})
}

// The jsii proxy struct for CfnComponent
type jsiiProxy_CfnComponent struct {
	internal.Type__awscdkCfnResource
	internal.Type__awscdkIInspectable
}

func (j *jsiiProxy_CfnComponent) AppId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) AttrId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attrId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) BindingProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bindingProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CfnOptions() awscdk.ICfnResourceOptions {
	var returns awscdk.ICfnResourceOptions
	_jsii_.Get(
		j,
		"cfnOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CfnProperties() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"cfnProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CfnResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Children() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"children",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CollectionProperties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"collectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) ComponentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"componentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) EnvironmentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Events() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"events",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) LogicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Overrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Properties() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) SchemaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) SourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) UpdatedProperites() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"updatedProperites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponent) Variants() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"variants",
		&returns,
	)
	return returns
}


// Create a new `AWS::AmplifyUIBuilder::Component`.
func NewCfnComponent(scope awscdk.Construct, id *string, props *CfnComponentProps) CfnComponent {
	_init_.Initialize()

	if err := validateNewCfnComponentParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnComponent{}

	_jsii_.Create(
		"monocdk.aws_amplifyuibuilder.CfnComponent",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Create a new `AWS::AmplifyUIBuilder::Component`.
func NewCfnComponent_Override(c CfnComponent, scope awscdk.Construct, id *string, props *CfnComponentProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplifyuibuilder.CfnComponent",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CfnComponent)SetAppId(val *string) {
	_jsii_.Set(
		j,
		"appId",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetBindingProperties(val interface{}) {
	if err := j.validateSetBindingPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindingProperties",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetChildren(val interface{}) {
	if err := j.validateSetChildrenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"children",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetCollectionProperties(val interface{}) {
	if err := j.validateSetCollectionPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collectionProperties",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetComponentType(val *string) {
	if err := j.validateSetComponentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"componentType",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetEnvironmentName(val *string) {
	_jsii_.Set(
		j,
		"environmentName",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetEvents(val interface{}) {
	if err := j.validateSetEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"events",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetOverrides(val interface{}) {
	if err := j.validateSetOverridesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrides",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetProperties(val interface{}) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetSchemaVersion(val *string) {
	_jsii_.Set(
		j,
		"schemaVersion",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetSourceId(val *string) {
	_jsii_.Set(
		j,
		"sourceId",
		val,
	)
}

func (j *jsiiProxy_CfnComponent)SetVariants(val interface{}) {
	if err := j.validateSetVariantsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"variants",
		val,
	)
}

// Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).
//
// Uses duck-typing instead of `instanceof` to allow stack elements from different
// versions of this library to be included in the same stack.
//
// Returns: The construct as a stack element or undefined if it is not a stack element.
// Experimental.
func CfnComponent_IsCfnElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnComponent_IsCfnElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplifyuibuilder.CfnComponent",
		"isCfnElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Check whether the given construct is a CfnResource.
// Experimental.
func CfnComponent_IsCfnResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateCfnComponent_IsCfnResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplifyuibuilder.CfnComponent",
		"isCfnResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CfnComponent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnComponent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_amplifyuibuilder.CfnComponent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnComponent_CFN_RESOURCE_TYPE_NAME() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_amplifyuibuilder.CfnComponent",
		"CFN_RESOURCE_TYPE_NAME",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnComponent) AddDeletionOverride(path *string) {
	if err := c.validateAddDeletionOverrideParameters(path); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDeletionOverride",
		[]interface{}{path},
	)
}

func (c *jsiiProxy_CfnComponent) AddDependsOn(target awscdk.CfnResource) {
	if err := c.validateAddDependsOnParameters(target); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addDependsOn",
		[]interface{}{target},
	)
}

func (c *jsiiProxy_CfnComponent) AddMetadata(key *string, value interface{}) {
	if err := c.validateAddMetadataParameters(key, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMetadata",
		[]interface{}{key, value},
	)
}

func (c *jsiiProxy_CfnComponent) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CfnComponent) AddPropertyDeletionOverride(propertyPath *string) {
	if err := c.validateAddPropertyDeletionOverrideParameters(propertyPath); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyDeletionOverride",
		[]interface{}{propertyPath},
	)
}

func (c *jsiiProxy_CfnComponent) AddPropertyOverride(propertyPath *string, value interface{}) {
	if err := c.validateAddPropertyOverrideParameters(propertyPath, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPropertyOverride",
		[]interface{}{propertyPath, value},
	)
}

func (c *jsiiProxy_CfnComponent) ApplyRemovalPolicy(policy awscdk.RemovalPolicy, options *awscdk.RemovalPolicyOptions) {
	if err := c.validateApplyRemovalPolicyParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyRemovalPolicy",
		[]interface{}{policy, options},
	)
}

func (c *jsiiProxy_CfnComponent) GetAtt(attributeName *string) awscdk.Reference {
	if err := c.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		c,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) GetMetadata(key *string) interface{} {
	if err := c.validateGetMetadataParameters(key); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"getMetadata",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) Inspect(inspector awscdk.TreeInspector) {
	if err := c.validateInspectParameters(inspector); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"inspect",
		[]interface{}{inspector},
	)
}

func (c *jsiiProxy_CfnComponent) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnComponent) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnComponent) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CfnComponent) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CfnComponent) RenderProperties(props *map[string]interface{}) *map[string]interface{} {
	if err := c.validateRenderPropertiesParameters(props); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"renderProperties",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) ShouldSynthesize() *bool {
	var returns *bool

	_jsii_.Invoke(
		c,
		"shouldSynthesize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) Synthesize(session awscdk.ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CfnComponent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnComponent) ValidateProperties(_properties interface{}) {
	if err := c.validateValidatePropertiesParameters(_properties); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"validateProperties",
		[]interface{}{_properties},
	)
}

