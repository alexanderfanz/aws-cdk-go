package awsquicksight


// The control of a parameter that users can interact with in a dashboard or an analysis.
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parameterControlProperty := &ParameterControlProperty{
//   	DateTimePicker: &ParameterDateTimePickerControlProperty{
//   		ParameterControlId: jsii.String("parameterControlId"),
//   		SourceParameterName: jsii.String("sourceParameterName"),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		DisplayOptions: &DateTimePickerControlDisplayOptionsProperty{
//   			DateTimeFormat: jsii.String("dateTimeFormat"),
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   	},
//   	Dropdown: &ParameterDropDownControlProperty{
//   		ParameterControlId: jsii.String("parameterControlId"),
//   		SourceParameterName: jsii.String("sourceParameterName"),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		CascadingControlConfiguration: &CascadingControlConfigurationProperty{
//   			SourceControls: []interface{}{
//   				&CascadingControlSourceProperty{
//   					ColumnToMatch: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					SourceSheetControlId: jsii.String("sourceSheetControlId"),
//   				},
//   			},
//   		},
//   		DisplayOptions: &DropDownControlDisplayOptionsProperty{
//   			SelectAllOptions: &ListControlSelectAllOptionsProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   		SelectableValues: &ParameterSelectableValuesProperty{
//   			LinkToDataSetColumn: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	List: &ParameterListControlProperty{
//   		ParameterControlId: jsii.String("parameterControlId"),
//   		SourceParameterName: jsii.String("sourceParameterName"),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		CascadingControlConfiguration: &CascadingControlConfigurationProperty{
//   			SourceControls: []interface{}{
//   				&CascadingControlSourceProperty{
//   					ColumnToMatch: &ColumnIdentifierProperty{
//   						ColumnName: jsii.String("columnName"),
//   						DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   					},
//   					SourceSheetControlId: jsii.String("sourceSheetControlId"),
//   				},
//   			},
//   		},
//   		DisplayOptions: &ListControlDisplayOptionsProperty{
//   			SearchOptions: &ListControlSearchOptionsProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			SelectAllOptions: &ListControlSelectAllOptionsProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   		SelectableValues: &ParameterSelectableValuesProperty{
//   			LinkToDataSetColumn: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Type: jsii.String("type"),
//   	},
//   	Slider: &ParameterSliderControlProperty{
//   		MaximumValue: jsii.Number(123),
//   		MinimumValue: jsii.Number(123),
//   		ParameterControlId: jsii.String("parameterControlId"),
//   		SourceParameterName: jsii.String("sourceParameterName"),
//   		StepSize: jsii.Number(123),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		DisplayOptions: &SliderControlDisplayOptionsProperty{
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   	},
//   	TextArea: &ParameterTextAreaControlProperty{
//   		ParameterControlId: jsii.String("parameterControlId"),
//   		SourceParameterName: jsii.String("sourceParameterName"),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		Delimiter: jsii.String("delimiter"),
//   		DisplayOptions: &TextAreaControlDisplayOptionsProperty{
//   			PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   	},
//   	TextField: &ParameterTextFieldControlProperty{
//   		ParameterControlId: jsii.String("parameterControlId"),
//   		SourceParameterName: jsii.String("sourceParameterName"),
//   		Title: jsii.String("title"),
//
//   		// the properties below are optional
//   		DisplayOptions: &TextFieldControlDisplayOptionsProperty{
//   			PlaceholderOptions: &TextControlPlaceholderOptionsProperty{
//   				Visibility: jsii.String("visibility"),
//   			},
//   			TitleOptions: &LabelOptionsProperty{
//   				CustomLabel: jsii.String("customLabel"),
//   				FontConfiguration: &FontConfigurationProperty{
//   					FontColor: jsii.String("fontColor"),
//   					FontDecoration: jsii.String("fontDecoration"),
//   					FontSize: &FontSizeProperty{
//   						Relative: jsii.String("relative"),
//   					},
//   					FontStyle: jsii.String("fontStyle"),
//   					FontWeight: &FontWeightProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Visibility: jsii.String("visibility"),
//   			},
//   		},
//   	},
//   }
//
type CfnAnalysis_ParameterControlProperty struct {
	// A control from a date parameter that specifies date and time.
	DateTimePicker interface{} `field:"optional" json:"dateTimePicker" yaml:"dateTimePicker"`
	// A control to display a dropdown list with buttons that are used to select a single value.
	Dropdown interface{} `field:"optional" json:"dropdown" yaml:"dropdown"`
	// A control to display a list with buttons or boxes that are used to select either a single value or multiple values.
	List interface{} `field:"optional" json:"list" yaml:"list"`
	// A control to display a horizontal toggle bar.
	//
	// This is used to change a value by sliding the toggle.
	Slider interface{} `field:"optional" json:"slider" yaml:"slider"`
	// A control to display a text box that is used to enter multiple entries.
	TextArea interface{} `field:"optional" json:"textArea" yaml:"textArea"`
	// A control to display a text box that is used to enter a single entry.
	TextField interface{} `field:"optional" json:"textField" yaml:"textField"`
}
