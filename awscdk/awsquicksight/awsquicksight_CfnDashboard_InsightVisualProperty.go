package awsquicksight


// An insight visual.
//
// For more information, see [Working with insights](https://docs.aws.amazon.com/quicksight/latest/user/computational-insights.html) in the *Amazon QuickSight User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   insightVisualProperty := &InsightVisualProperty{
//   	DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	VisualId: jsii.String("visualId"),
//
//   	// the properties below are optional
//   	Actions: []interface{}{
//   		&VisualCustomActionProperty{
//   			ActionOperations: []interface{}{
//   				&VisualCustomActionOperationProperty{
//   					FilterOperation: &CustomActionFilterOperationProperty{
//   						SelectedFieldsConfiguration: &FilterOperationSelectedFieldsConfigurationProperty{
//   							SelectedFieldOptions: jsii.String("selectedFieldOptions"),
//   							SelectedFields: []*string{
//   								jsii.String("selectedFields"),
//   							},
//   						},
//   						TargetVisualsConfiguration: &FilterOperationTargetVisualsConfigurationProperty{
//   							SameSheetTargetVisualConfiguration: &SameSheetTargetVisualConfigurationProperty{
//   								TargetVisualOptions: jsii.String("targetVisualOptions"),
//   								TargetVisuals: []*string{
//   									jsii.String("targetVisuals"),
//   								},
//   							},
//   						},
//   					},
//   					NavigationOperation: &CustomActionNavigationOperationProperty{
//   						LocalNavigationConfiguration: &LocalNavigationConfigurationProperty{
//   							TargetSheetId: jsii.String("targetSheetId"),
//   						},
//   					},
//   					SetParametersOperation: &CustomActionSetParametersOperationProperty{
//   						ParameterValueConfigurations: []interface{}{
//   							&SetParameterValueConfigurationProperty{
//   								DestinationParameterName: jsii.String("destinationParameterName"),
//   								Value: &DestinationParameterValueConfigurationProperty{
//   									CustomValuesConfiguration: &CustomValuesConfigurationProperty{
//   										CustomValues: &CustomParameterValuesProperty{
//   											DateTimeValues: []*string{
//   												jsii.String("dateTimeValues"),
//   											},
//   											DecimalValues: []interface{}{
//   												jsii.Number(123),
//   											},
//   											IntegerValues: []interface{}{
//   												jsii.Number(123),
//   											},
//   											StringValues: []*string{
//   												jsii.String("stringValues"),
//   											},
//   										},
//
//   										// the properties below are optional
//   										IncludeNullValue: jsii.Boolean(false),
//   									},
//   									SelectAllValueOptions: jsii.String("selectAllValueOptions"),
//   									SourceField: jsii.String("sourceField"),
//   									SourceParameterName: jsii.String("sourceParameterName"),
//   								},
//   							},
//   						},
//   					},
//   					UrlOperation: &CustomActionURLOperationProperty{
//   						UrlTarget: jsii.String("urlTarget"),
//   						UrlTemplate: jsii.String("urlTemplate"),
//   					},
//   				},
//   			},
//   			CustomActionId: jsii.String("customActionId"),
//   			Name: jsii.String("name"),
//   			Trigger: jsii.String("trigger"),
//
//   			// the properties below are optional
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	InsightConfiguration: &InsightConfigurationProperty{
//   		Computations: []interface{}{
//   			&ComputationProperty{
//   				Forecast: &ForecastComputationProperty{
//   					ComputationId: jsii.String("computationId"),
//   					Time: &DimensionFieldProperty{
//   						CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						DateDimensionField: &DateDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							DateGranularity: jsii.String("dateGranularity"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						NumericalDimensionField: &NumericalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					CustomSeasonalityValue: jsii.Number(123),
//   					LowerBoundary: jsii.Number(123),
//   					Name: jsii.String("name"),
//   					PeriodsBackward: jsii.Number(123),
//   					PeriodsForward: jsii.Number(123),
//   					PredictionInterval: jsii.Number(123),
//   					Seasonality: jsii.String("seasonality"),
//   					UpperBoundary: jsii.Number(123),
//   					Value: &MeasureFieldProperty{
//   						CalculatedMeasureField: &CalculatedMeasureFieldProperty{
//   							Expression: jsii.String("expression"),
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						CategoricalMeasureField: &CategoricalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						DateMeasureField: &DateMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						NumericalMeasureField: &NumericalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: &NumericalAggregationFunctionProperty{
//   								PercentileAggregation: &PercentileAggregationProperty{
//   									PercentileValue: jsii.Number(123),
//   								},
//   								SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   							},
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   				GrowthRate: &GrowthRateComputationProperty{
//   					ComputationId: jsii.String("computationId"),
//   					Time: &DimensionFieldProperty{
//   						CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						DateDimensionField: &DateDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							DateGranularity: jsii.String("dateGranularity"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						NumericalDimensionField: &NumericalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Name: jsii.String("name"),
//   					PeriodSize: jsii.Number(123),
//   					Value: &MeasureFieldProperty{
//   						CalculatedMeasureField: &CalculatedMeasureFieldProperty{
//   							Expression: jsii.String("expression"),
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						CategoricalMeasureField: &CategoricalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						DateMeasureField: &DateMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						NumericalMeasureField: &NumericalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: &NumericalAggregationFunctionProperty{
//   								PercentileAggregation: &PercentileAggregationProperty{
//   									PercentileValue: jsii.Number(123),
//   								},
//   								SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   							},
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   				MaximumMinimum: &MaximumMinimumComputationProperty{
//   					ComputationId: jsii.String("computationId"),
//   					Time: &DimensionFieldProperty{
//   						CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						DateDimensionField: &DateDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							DateGranularity: jsii.String("dateGranularity"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						NumericalDimensionField: &NumericalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   					},
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Name: jsii.String("name"),
//   					Value: &MeasureFieldProperty{
//   						CalculatedMeasureField: &CalculatedMeasureFieldProperty{
//   							Expression: jsii.String("expression"),
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						CategoricalMeasureField: &CategoricalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						DateMeasureField: &DateMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						NumericalMeasureField: &NumericalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: &NumericalAggregationFunctionProperty{
//   								PercentileAggregation: &PercentileAggregationProperty{
//   									PercentileValue: jsii.Number(123),
//   								},
//   								SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   							},
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   				MetricComparison: &MetricComparisonComputationProperty{
//   					ComputationId: jsii.String("computationId"),
//   					FromValue: &MeasureFieldProperty{
//   						CalculatedMeasureField: &CalculatedMeasureFieldProperty{
//   							Expression: jsii.String("expression"),
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						CategoricalMeasureField: &CategoricalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						DateMeasureField: &DateMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						NumericalMeasureField: &NumericalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: &NumericalAggregationFunctionProperty{
//   								PercentileAggregation: &PercentileAggregationProperty{
//   									PercentileValue: jsii.Number(123),
//   								},
//   								SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   							},
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   					TargetValue: &MeasureFieldProperty{
//   						CalculatedMeasureField: &CalculatedMeasureFieldProperty{
//   							Expression: jsii.String("expression"),
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						CategoricalMeasureField: &CategoricalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						DateMeasureField: &DateMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						NumericalMeasureField: &NumericalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: &NumericalAggregationFunctionProperty{
//   								PercentileAggregation: &PercentileAggregationProperty{
//   									PercentileValue: jsii.Number(123),
//   								},
//   								SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   							},
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   					Time: &DimensionFieldProperty{
//   						CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						DateDimensionField: &DateDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							DateGranularity: jsii.String("dateGranularity"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						NumericalDimensionField: &NumericalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Name: jsii.String("name"),
//   				},
//   				PeriodOverPeriod: &PeriodOverPeriodComputationProperty{
//   					ComputationId: jsii.String("computationId"),
//   					Time: &DimensionFieldProperty{
//   						CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						DateDimensionField: &DateDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							DateGranularity: jsii.String("dateGranularity"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						NumericalDimensionField: &NumericalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Name: jsii.String("name"),
//   					Value: &MeasureFieldProperty{
//   						CalculatedMeasureField: &CalculatedMeasureFieldProperty{
//   							Expression: jsii.String("expression"),
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						CategoricalMeasureField: &CategoricalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						DateMeasureField: &DateMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						NumericalMeasureField: &NumericalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: &NumericalAggregationFunctionProperty{
//   								PercentileAggregation: &PercentileAggregationProperty{
//   									PercentileValue: jsii.Number(123),
//   								},
//   								SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   							},
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   				PeriodToDate: &PeriodToDateComputationProperty{
//   					ComputationId: jsii.String("computationId"),
//   					Time: &DimensionFieldProperty{
//   						CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						DateDimensionField: &DateDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							DateGranularity: jsii.String("dateGranularity"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						NumericalDimensionField: &NumericalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   					},
//
//   					// the properties below are optional
//   					Name: jsii.String("name"),
//   					PeriodTimeGranularity: jsii.String("periodTimeGranularity"),
//   					Value: &MeasureFieldProperty{
//   						CalculatedMeasureField: &CalculatedMeasureFieldProperty{
//   							Expression: jsii.String("expression"),
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						CategoricalMeasureField: &CategoricalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						DateMeasureField: &DateMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						NumericalMeasureField: &NumericalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: &NumericalAggregationFunctionProperty{
//   								PercentileAggregation: &PercentileAggregationProperty{
//   									PercentileValue: jsii.Number(123),
//   								},
//   								SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   							},
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   				TopBottomMovers: &TopBottomMoversComputationProperty{
//   					Category: &DimensionFieldProperty{
//   						CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						DateDimensionField: &DateDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							DateGranularity: jsii.String("dateGranularity"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						NumericalDimensionField: &NumericalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   					},
//   					ComputationId: jsii.String("computationId"),
//   					Time: &DimensionFieldProperty{
//   						CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						DateDimensionField: &DateDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							DateGranularity: jsii.String("dateGranularity"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						NumericalDimensionField: &NumericalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   					},
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					MoverSize: jsii.Number(123),
//   					Name: jsii.String("name"),
//   					SortOrder: jsii.String("sortOrder"),
//   					Value: &MeasureFieldProperty{
//   						CalculatedMeasureField: &CalculatedMeasureFieldProperty{
//   							Expression: jsii.String("expression"),
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						CategoricalMeasureField: &CategoricalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						DateMeasureField: &DateMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						NumericalMeasureField: &NumericalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: &NumericalAggregationFunctionProperty{
//   								PercentileAggregation: &PercentileAggregationProperty{
//   									PercentileValue: jsii.Number(123),
//   								},
//   								SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   							},
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   				TopBottomRanked: &TopBottomRankedComputationProperty{
//   					Category: &DimensionFieldProperty{
//   						CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						DateDimensionField: &DateDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							DateGranularity: jsii.String("dateGranularity"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						NumericalDimensionField: &NumericalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   					},
//   					ComputationId: jsii.String("computationId"),
//   					Type: jsii.String("type"),
//
//   					// the properties below are optional
//   					Name: jsii.String("name"),
//   					ResultSize: jsii.Number(123),
//   					Value: &MeasureFieldProperty{
//   						CalculatedMeasureField: &CalculatedMeasureFieldProperty{
//   							Expression: jsii.String("expression"),
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						CategoricalMeasureField: &CategoricalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						DateMeasureField: &DateMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						NumericalMeasureField: &NumericalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: &NumericalAggregationFunctionProperty{
//   								PercentileAggregation: &PercentileAggregationProperty{
//   									PercentileValue: jsii.Number(123),
//   								},
//   								SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   							},
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   				},
//   				TotalAggregation: &TotalAggregationComputationProperty{
//   					ComputationId: jsii.String("computationId"),
//   					Value: &MeasureFieldProperty{
//   						CalculatedMeasureField: &CalculatedMeasureFieldProperty{
//   							Expression: jsii.String("expression"),
//   							FieldId: jsii.String("fieldId"),
//   						},
//   						CategoricalMeasureField: &CategoricalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						DateMeasureField: &DateMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: jsii.String("aggregationFunction"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   						NumericalMeasureField: &NumericalMeasureFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							AggregationFunction: &NumericalAggregationFunctionProperty{
//   								PercentileAggregation: &PercentileAggregationProperty{
//   									PercentileValue: jsii.Number(123),
//   								},
//   								SimpleNumericalAggregation: jsii.String("simpleNumericalAggregation"),
//   							},
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   						},
//   					},
//
//   					// the properties below are optional
//   					Name: jsii.String("name"),
//   				},
//   				UniqueValues: &UniqueValuesComputationProperty{
//   					Category: &DimensionFieldProperty{
//   						CategoricalDimensionField: &CategoricalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &StringFormatConfigurationProperty{
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						DateDimensionField: &DateDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							DateGranularity: jsii.String("dateGranularity"),
//   							FormatConfiguration: &DateTimeFormatConfigurationProperty{
//   								DateTimeFormat: jsii.String("dateTimeFormat"),
//   								NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   									NullString: jsii.String("nullString"),
//   								},
//   								NumericFormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   						NumericalDimensionField: &NumericalDimensionFieldProperty{
//   							Column: &ColumnIdentifierProperty{
//   								ColumnName: jsii.String("columnName"),
//   								DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   							},
//   							FieldId: jsii.String("fieldId"),
//
//   							// the properties below are optional
//   							FormatConfiguration: &NumberFormatConfigurationProperty{
//   								FormatConfiguration: &NumericFormatConfigurationProperty{
//   									CurrencyDisplayFormatConfiguration: &CurrencyDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   										Symbol: jsii.String("symbol"),
//   									},
//   									NumberDisplayFormatConfiguration: &NumberDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										NumberScale: jsii.String("numberScale"),
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   									PercentageDisplayFormatConfiguration: &PercentageDisplayFormatConfigurationProperty{
//   										DecimalPlacesConfiguration: &DecimalPlacesConfigurationProperty{
//   											DecimalPlaces: jsii.Number(123),
//   										},
//   										NegativeValueConfiguration: &NegativeValueConfigurationProperty{
//   											DisplayMode: jsii.String("displayMode"),
//   										},
//   										NullValueFormatConfiguration: &NullValueFormatConfigurationProperty{
//   											NullString: jsii.String("nullString"),
//   										},
//   										Prefix: jsii.String("prefix"),
//   										SeparatorConfiguration: &NumericSeparatorConfigurationProperty{
//   											DecimalSeparator: jsii.String("decimalSeparator"),
//   											ThousandsSeparator: &ThousandSeparatorOptionsProperty{
//   												Symbol: jsii.String("symbol"),
//   												Visibility: jsii.String("visibility"),
//   											},
//   										},
//   										Suffix: jsii.String("suffix"),
//   									},
//   								},
//   							},
//   							HierarchyId: jsii.String("hierarchyId"),
//   						},
//   					},
//   					ComputationId: jsii.String("computationId"),
//
//   					// the properties below are optional
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   		CustomNarrative: &CustomNarrativeOptionsProperty{
//   			Narrative: jsii.String("narrative"),
//   		},
//   	},
//   	Subtitle: &VisualSubtitleLabelOptionsProperty{
//   		FormatText: &LongFormatTextProperty{
//   			PlainText: jsii.String("plainText"),
//   			RichText: jsii.String("richText"),
//   		},
//   		Visibility: jsii.String("visibility"),
//   	},
//   	Title: &VisualTitleLabelOptionsProperty{
//   		FormatText: &ShortFormatTextProperty{
//   			PlainText: jsii.String("plainText"),
//   			RichText: jsii.String("richText"),
//   		},
//   		Visibility: jsii.String("visibility"),
//   	},
//   }
//
type CfnDashboard_InsightVisualProperty struct {
	// The dataset that is used in the insight visual.
	DataSetIdentifier *string `field:"required" json:"dataSetIdentifier" yaml:"dataSetIdentifier"`
	// The unique identifier of a visual.
	//
	// This identifier must be unique within the context of a dashboard, template, or analysis. Two dashboards, analyses, or templates can have visuals with the same identifiers.
	VisualId *string `field:"required" json:"visualId" yaml:"visualId"`
	// The list of custom actions that are configured for a visual.
	Actions interface{} `field:"optional" json:"actions" yaml:"actions"`
	// The configuration of an insight visual.
	InsightConfiguration interface{} `field:"optional" json:"insightConfiguration" yaml:"insightConfiguration"`
	// The subtitle that is displayed on the visual.
	Subtitle interface{} `field:"optional" json:"subtitle" yaml:"subtitle"`
	// The title that is displayed on the visual.
	Title interface{} `field:"optional" json:"title" yaml:"title"`
}
