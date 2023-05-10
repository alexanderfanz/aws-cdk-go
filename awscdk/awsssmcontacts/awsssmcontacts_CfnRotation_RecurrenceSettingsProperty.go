package awsssmcontacts


// Information about when an on-call rotation is in effect and how long the rotation period lasts.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   recurrenceSettingsProperty := &RecurrenceSettingsProperty{
//   	NumberOfOnCalls: jsii.Number(123),
//   	RecurrenceMultiplier: jsii.Number(123),
//
//   	// the properties below are optional
//   	DailySettings: []*string{
//   		jsii.String("dailySettings"),
//   	},
//   	MonthlySettings: []interface{}{
//   		&MonthlySettingProperty{
//   			DayOfMonth: jsii.Number(123),
//   			HandOffTime: jsii.String("handOffTime"),
//   		},
//   	},
//   	ShiftCoverages: []interface{}{
//   		&ShiftCoverageProperty{
//   			CoverageTimes: []interface{}{
//   				&CoverageTimeProperty{
//   					EndTime: jsii.String("endTime"),
//   					StartTime: jsii.String("startTime"),
//   				},
//   			},
//   			DayOfWeek: jsii.String("dayOfWeek"),
//   		},
//   	},
//   	WeeklySettings: []interface{}{
//   		&WeeklySettingProperty{
//   			DayOfWeek: jsii.String("dayOfWeek"),
//   			HandOffTime: jsii.String("handOffTime"),
//   		},
//   	},
//   }
//
type CfnRotation_RecurrenceSettingsProperty struct {
	// The number of contacts, or shift team members designated to be on call concurrently during a shift.
	//
	// For example, in an on-call schedule that contains ten contacts, a value of `2` designates that two of them are on call at any given time.
	NumberOfOnCalls *float64 `field:"required" json:"numberOfOnCalls" yaml:"numberOfOnCalls"`
	// The number of days, weeks, or months a single rotation lasts.
	RecurrenceMultiplier *float64 `field:"required" json:"recurrenceMultiplier" yaml:"recurrenceMultiplier"`
	// Information about on-call rotations that recur daily.
	DailySettings *[]*string `field:"optional" json:"dailySettings" yaml:"dailySettings"`
	// Information about on-call rotations that recur monthly.
	MonthlySettings interface{} `field:"optional" json:"monthlySettings" yaml:"monthlySettings"`
	// Information about the days of the week included in on-call rotation coverage.
	ShiftCoverages interface{} `field:"optional" json:"shiftCoverages" yaml:"shiftCoverages"`
	// Information about on-call rotations that recur weekly.
	WeeklySettings interface{} `field:"optional" json:"weeklySettings" yaml:"weeklySettings"`
}
