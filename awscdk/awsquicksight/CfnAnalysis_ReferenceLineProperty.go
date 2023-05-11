package awsquicksight


// The reference line visual display options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   referenceLineProperty := &ReferenceLineProperty{
//   	DataConfiguration: &ReferenceLineDataConfigurationProperty{
//   		AxisBinding: jsii.String("axisBinding"),
//   		DynamicConfiguration: &ReferenceLineDynamicDataConfigurationProperty{
//   			Calculation: &NumericalAggregationFunctionProperty{
//   				PercentileAggregation: &PercentileAggregationProperty{
//   					PercentileValue: jsii.Number(123),
//   				},
//   				SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   			},
//   			Column: &ColumnIdentifierProperty{
//   				ColumnName: jsii.String("columnName"),
//   				DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   			},
//   			MeasureAggregationFunction: &AggregationFunctionProperty{
//   				CategoricalAggregationFunction: jsii.String("categoricalAggregationFunction"),
//   				DateAggregationFunction: jsii.String("dateAggregationFunction"),
//   				NumericalAggregationFunction: &NumericalAggregationFunctionProperty{
//   					PercentileAggregation: &PercentileAggregationProperty{
//   						PercentileValue: jsii.Number(123),
//   					},
//   					SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   				},
//   			},
//   		},
//   		StaticConfiguration: &ReferenceLineStaticDataConfigurationProperty{
//   			Value: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	LabelConfiguration: &ReferenceLineLabelConfigurationProperty{
//   		CustomLabelConfiguration: &ReferenceLineCustomLabelConfigurationProperty{
//   			CustomLabel: jsii.String("customLabel"),
//   		},
//   		FontColor: jsii.String("fontColor"),
//   		FontConfiguration: &FontConfigurationProperty{
//   			FontColor: jsii.String("fontColor"),
//   			FontDecoration: jsii.String("fontDecoration"),
//   			FontSize: &FontSizeProperty{
//   				Relative: jsii.String("relative"),
//   			},
//   			FontStyle: jsii.String("fontStyle"),
//   			FontWeight: &FontWeightProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		HorizontalPosition: jsii.String("horizontalPosition"),
//   		ValueLabelConfiguration: &ReferenceLineValueLabelConfigurationProperty{
//   			FormatConfiguration: &NumericFormatConfigurationProperty{
//   				CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   					DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   						DecimalPlaces: jsii.Number(123),
//   					},
//   					NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   						DisplayMode: jsii.String("displayMode"),
//   					},
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					NumberScale: jsii.String("numberScale"),
//   					Prefix: jsii.String("prefix"),
//   					SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   						DecimalSeparator: jsii.String("decimalSeparator"),
//   						ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   							Symbol: jsii.String("symbol"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					Suffix: jsii.String("suffix"),
//   					Symbol: jsii.String("symbol"),
//   				},
//   				NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   					DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   						DecimalPlaces: jsii.Number(123),
//   					},
//   					NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   						DisplayMode: jsii.String("displayMode"),
//   					},
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					NumberScale: jsii.String("numberScale"),
//   					Prefix: jsii.String("prefix"),
//   					SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   						DecimalSeparator: jsii.String("decimalSeparator"),
//   						ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   							Symbol: jsii.String("symbol"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					Suffix: jsii.String("suffix"),
//   				},
//   				PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   					DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   						DecimalPlaces: jsii.Number(123),
//   					},
//   					NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   						DisplayMode: jsii.String("displayMode"),
//   					},
//   					NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   						NullString: jsii.String("nullString"),
//   					},
//   					Prefix: jsii.String("prefix"),
//   					SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   						DecimalSeparator: jsii.String("decimalSeparator"),
//   						ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   							Symbol: jsii.String("symbol"),
//   							Visibility: jsii.String("visibility"),
//   						},
//   					},
//   					Suffix: jsii.String("suffix"),
//   				},
//   			},
//   			RelativePosition: jsii.String("relativePosition"),
//   		},
//   		VerticalPosition: jsii.String("verticalPosition"),
//   	},
//   	Status: jsii.String("status"),
//   	StyleConfiguration: &ReferenceLineStyleConfigurationProperty{
//   		Color: jsii.String("color"),
//   		Pattern: jsii.String("pattern"),
//   	},
//   }
//
type CfnAnalysis_ReferenceLineProperty struct {
	// The data configuration of the reference line.
	DataConfiguration interface{} `field:"required" json:"dataConfiguration" yaml:"dataConfiguration"`
	// The label configuration of the reference line.
	LabelConfiguration interface{} `field:"optional" json:"labelConfiguration" yaml:"labelConfiguration"`
	// The status of the reference line. Choose one of the following options:.
	//
	// - `ENABLE`
	// - `DISABLE`.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The style configuration of the reference line.
	StyleConfiguration interface{} `field:"optional" json:"styleConfiguration" yaml:"styleConfiguration"`
}
