package awsmedialive


// The encoding information for output captions.
//
// The parent of this entity is EncoderSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   captionDescriptionProperty := &CaptionDescriptionProperty{
//   	CaptionSelectorName: jsii.String("captionSelectorName"),
//   	DestinationSettings: &CaptionDestinationSettingsProperty{
//   		AribDestinationSettings: &AribDestinationSettingsProperty{
//   		},
//   		BurnInDestinationSettings: &BurnInDestinationSettingsProperty{
//   			Alignment: jsii.String("alignment"),
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			BackgroundOpacity: jsii.Number(123),
//   			Font: &InputLocationProperty{
//   				PasswordParam: jsii.String("passwordParam"),
//   				Uri: jsii.String("uri"),
//   				Username: jsii.String("username"),
//   			},
//   			FontColor: jsii.String("fontColor"),
//   			FontOpacity: jsii.Number(123),
//   			FontResolution: jsii.Number(123),
//   			FontSize: jsii.String("fontSize"),
//   			OutlineColor: jsii.String("outlineColor"),
//   			OutlineSize: jsii.Number(123),
//   			ShadowColor: jsii.String("shadowColor"),
//   			ShadowOpacity: jsii.Number(123),
//   			ShadowXOffset: jsii.Number(123),
//   			ShadowYOffset: jsii.Number(123),
//   			TeletextGridControl: jsii.String("teletextGridControl"),
//   			XPosition: jsii.Number(123),
//   			YPosition: jsii.Number(123),
//   		},
//   		DvbSubDestinationSettings: &DvbSubDestinationSettingsProperty{
//   			Alignment: jsii.String("alignment"),
//   			BackgroundColor: jsii.String("backgroundColor"),
//   			BackgroundOpacity: jsii.Number(123),
//   			Font: &InputLocationProperty{
//   				PasswordParam: jsii.String("passwordParam"),
//   				Uri: jsii.String("uri"),
//   				Username: jsii.String("username"),
//   			},
//   			FontColor: jsii.String("fontColor"),
//   			FontOpacity: jsii.Number(123),
//   			FontResolution: jsii.Number(123),
//   			FontSize: jsii.String("fontSize"),
//   			OutlineColor: jsii.String("outlineColor"),
//   			OutlineSize: jsii.Number(123),
//   			ShadowColor: jsii.String("shadowColor"),
//   			ShadowOpacity: jsii.Number(123),
//   			ShadowXOffset: jsii.Number(123),
//   			ShadowYOffset: jsii.Number(123),
//   			TeletextGridControl: jsii.String("teletextGridControl"),
//   			XPosition: jsii.Number(123),
//   			YPosition: jsii.Number(123),
//   		},
//   		EbuTtDDestinationSettings: &EbuTtDDestinationSettingsProperty{
//   			CopyrightHolder: jsii.String("copyrightHolder"),
//   			FillLineGap: jsii.String("fillLineGap"),
//   			FontFamily: jsii.String("fontFamily"),
//   			StyleControl: jsii.String("styleControl"),
//   		},
//   		EmbeddedDestinationSettings: &EmbeddedDestinationSettingsProperty{
//   		},
//   		EmbeddedPlusScte20DestinationSettings: &EmbeddedPlusScte20DestinationSettingsProperty{
//   		},
//   		RtmpCaptionInfoDestinationSettings: &RtmpCaptionInfoDestinationSettingsProperty{
//   		},
//   		Scte20PlusEmbeddedDestinationSettings: &Scte20PlusEmbeddedDestinationSettingsProperty{
//   		},
//   		Scte27DestinationSettings: &Scte27DestinationSettingsProperty{
//   		},
//   		SmpteTtDestinationSettings: &SmpteTtDestinationSettingsProperty{
//   		},
//   		TeletextDestinationSettings: &TeletextDestinationSettingsProperty{
//   		},
//   		TtmlDestinationSettings: &TtmlDestinationSettingsProperty{
//   			StyleControl: jsii.String("styleControl"),
//   		},
//   		WebvttDestinationSettings: &WebvttDestinationSettingsProperty{
//   			StyleControl: jsii.String("styleControl"),
//   		},
//   	},
//   	LanguageCode: jsii.String("languageCode"),
//   	LanguageDescription: jsii.String("languageDescription"),
//   	Name: jsii.String("name"),
//   }
//
type CfnChannel_CaptionDescriptionProperty struct {
	// Specifies which input captions selector to use as a captions source when generating output captions.
	//
	// This field should match a captionSelector name.
	CaptionSelectorName *string `field:"optional" json:"captionSelectorName" yaml:"captionSelectorName"`
	// Additional settings for a captions destination that depend on the destination type.
	DestinationSettings interface{} `field:"optional" json:"destinationSettings" yaml:"destinationSettings"`
	// An ISO 639-2 three-digit code.
	//
	// For more information, see http://www.loc.gov/standards/iso639-2/.
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
	// Human-readable information to indicate the captions that are available for players (for example, English or Spanish).
	LanguageDescription *string `field:"optional" json:"languageDescription" yaml:"languageDescription"`
	// The name of the captions description.
	//
	// The name is used to associate a captions description with an output. Names must be unique within a channel.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

