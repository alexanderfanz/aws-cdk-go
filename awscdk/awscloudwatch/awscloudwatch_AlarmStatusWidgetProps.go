package awscloudwatch


// Properties for an Alarm Status Widget.
//
// Example:
//   var dashboard dashboard
//   var errorAlarm alarm
//
//
//   dashboard.AddWidgets(
//   cloudwatch.NewAlarmStatusWidget(&AlarmStatusWidgetProps{
//   	Alarms: []iAlarm{
//   		errorAlarm,
//   	},
//   }))
//
// Experimental.
type AlarmStatusWidgetProps struct {
	// CloudWatch Alarms to show in widget.
	// Experimental.
	Alarms *[]IAlarm `field:"required" json:"alarms" yaml:"alarms"`
	// Height of the widget.
	// Experimental.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Specifies how to sort the alarms in the widget.
	// Experimental.
	SortBy AlarmStatusWidgetSortBy `field:"optional" json:"sortBy" yaml:"sortBy"`
	// Use this field to filter the list of alarms displayed in the widget to only those alarms currently in the specified states.
	//
	// You can specify one or more alarm states in the value for this field.
	// The alarm states that you can specify are ALARM, INSUFFICIENT_DATA, and OK.
	//
	// If you omit this field or specify an empty array, all the alarms specifed in alarms are displayed.
	// Experimental.
	States *[]AlarmState `field:"optional" json:"states" yaml:"states"`
	// The title of the widget.
	// Experimental.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Width of the widget, in a grid of 24 units wide.
	// Experimental.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

