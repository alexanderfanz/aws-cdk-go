package awsmediapackage


// Parameters for Common Media Application Format (CMAF) packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cmafPackageProperty := &CmafPackageProperty{
//   	Encryption: &CmafEncryptionProperty{
//   		SpekeKeyProvider: &SpekeKeyProviderProperty{
//   			ResourceId: jsii.String("resourceId"),
//   			RoleArn: jsii.String("roleArn"),
//   			SystemIds: []*string{
//   				jsii.String("systemIds"),
//   			},
//   			Url: jsii.String("url"),
//
//   			// the properties below are optional
//   			CertificateArn: jsii.String("certificateArn"),
//   			EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   			},
//   		},
//
//   		// the properties below are optional
//   		ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   		EncryptionMethod: jsii.String("encryptionMethod"),
//   		KeyRotationIntervalSeconds: jsii.Number(123),
//   	},
//   	HlsManifests: []interface{}{
//   		&HlsManifestProperty{
//   			Id: jsii.String("id"),
//
//   			// the properties below are optional
//   			AdMarkers: jsii.String("adMarkers"),
//   			AdsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   			AdTriggers: []*string{
//   				jsii.String("adTriggers"),
//   			},
//   			IncludeIframeOnlyStream: jsii.Boolean(false),
//   			ManifestName: jsii.String("manifestName"),
//   			PlaylistType: jsii.String("playlistType"),
//   			PlaylistWindowSeconds: jsii.Number(123),
//   			ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   			Url: jsii.String("url"),
//   		},
//   	},
//   	SegmentDurationSeconds: jsii.Number(123),
//   	SegmentPrefix: jsii.String("segmentPrefix"),
//   	StreamSelection: &StreamSelectionProperty{
//   		MaxVideoBitsPerSecond: jsii.Number(123),
//   		MinVideoBitsPerSecond: jsii.Number(123),
//   		StreamOrder: jsii.String("streamOrder"),
//   	},
//   }
//
type CfnOriginEndpoint_CmafPackageProperty struct {
	// Parameters for encrypting content.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// A list of HLS manifest configurations that are available from this endpoint.
	HlsManifests interface{} `field:"optional" json:"hlsManifests" yaml:"hlsManifests"`
	// Duration (in seconds) of each segment.
	//
	// Actual segments are rounded to the nearest multiple of the source segment duration.
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// An optional custom string that is prepended to the name of each segment.
	//
	// If not specified, the segment prefix defaults to the ChannelId.
	SegmentPrefix *string `field:"optional" json:"segmentPrefix" yaml:"segmentPrefix"`
	// Limitations for outputs from the endpoint, based on the video bitrate.
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
}
