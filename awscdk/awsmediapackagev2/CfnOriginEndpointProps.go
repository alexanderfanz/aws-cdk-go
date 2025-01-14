package awsmediapackagev2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnOriginEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOriginEndpointProps := &CfnOriginEndpointProps{
//   	ContainerType: jsii.String("containerType"),
//
//   	// the properties below are optional
//   	ChannelGroupName: jsii.String("channelGroupName"),
//   	ChannelName: jsii.String("channelName"),
//   	Description: jsii.String("description"),
//   	HlsManifests: []interface{}{
//   		&HlsManifestConfigurationProperty{
//   			ManifestName: jsii.String("manifestName"),
//
//   			// the properties below are optional
//   			ChildManifestName: jsii.String("childManifestName"),
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				End: jsii.String("end"),
//   				ManifestFilter: jsii.String("manifestFilter"),
//   				Start: jsii.String("start"),
//   				TimeDelaySeconds: jsii.Number(123),
//   			},
//   			ManifestWindowSeconds: jsii.Number(123),
//   			ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   			ScteHls: &ScteHlsProperty{
//   				AdMarkerHls: jsii.String("adMarkerHls"),
//   			},
//   			Url: jsii.String("url"),
//   		},
//   	},
//   	LowLatencyHlsManifests: []interface{}{
//   		&LowLatencyHlsManifestConfigurationProperty{
//   			ManifestName: jsii.String("manifestName"),
//
//   			// the properties below are optional
//   			ChildManifestName: jsii.String("childManifestName"),
//   			FilterConfiguration: &FilterConfigurationProperty{
//   				End: jsii.String("end"),
//   				ManifestFilter: jsii.String("manifestFilter"),
//   				Start: jsii.String("start"),
//   				TimeDelaySeconds: jsii.Number(123),
//   			},
//   			ManifestWindowSeconds: jsii.Number(123),
//   			ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   			ScteHls: &ScteHlsProperty{
//   				AdMarkerHls: jsii.String("adMarkerHls"),
//   			},
//   			Url: jsii.String("url"),
//   		},
//   	},
//   	OriginEndpointName: jsii.String("originEndpointName"),
//   	Segment: &SegmentProperty{
//   		Encryption: &EncryptionProperty{
//   			EncryptionMethod: &EncryptionMethodProperty{
//   				CmafEncryptionMethod: jsii.String("cmafEncryptionMethod"),
//   				TsEncryptionMethod: jsii.String("tsEncryptionMethod"),
//   			},
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				DrmSystems: []*string{
//   					jsii.String("drmSystems"),
//   				},
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   					PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   				ResourceId: jsii.String("resourceId"),
//   				RoleArn: jsii.String("roleArn"),
//   				Url: jsii.String("url"),
//   			},
//
//   			// the properties below are optional
//   			ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   			KeyRotationIntervalSeconds: jsii.Number(123),
//   		},
//   		IncludeIframeOnlyStreams: jsii.Boolean(false),
//   		Scte: &ScteProperty{
//   			ScteFilter: []*string{
//   				jsii.String("scteFilter"),
//   			},
//   		},
//   		SegmentDurationSeconds: jsii.Number(123),
//   		SegmentName: jsii.String("segmentName"),
//   		TsIncludeDvbSubtitles: jsii.Boolean(false),
//   		TsUseAudioRenditionGroup: jsii.Boolean(false),
//   	},
//   	StartoverWindowSeconds: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html
//
type CfnOriginEndpointProps struct {
	// The container type associated with the origin endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-containertype
	//
	ContainerType *string `field:"required" json:"containerType" yaml:"containerType"`
	// The name of the channel group associated with the origin endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-channelgroupname
	//
	ChannelGroupName *string `field:"optional" json:"channelGroupName" yaml:"channelGroupName"`
	// The channel name associated with the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-channelname
	//
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// The description associated with the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The HLS manfiests associated with the origin endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-hlsmanifests
	//
	HlsManifests interface{} `field:"optional" json:"hlsManifests" yaml:"hlsManifests"`
	// The low-latency HLS (LL-HLS) manifests associated with the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-lowlatencyhlsmanifests
	//
	LowLatencyHlsManifests interface{} `field:"optional" json:"lowLatencyHlsManifests" yaml:"lowLatencyHlsManifests"`
	// The name of the origin endpoint associated with the origin endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-originendpointname
	//
	OriginEndpointName *string `field:"optional" json:"originEndpointName" yaml:"originEndpointName"`
	// The segment associated with the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-segment
	//
	Segment interface{} `field:"optional" json:"segment" yaml:"segment"`
	// The size of the window (in seconds) to specify a window of the live stream that's available for on-demand viewing.
	//
	// Viewers can start-over or catch-up on content that falls within the window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-startoverwindowseconds
	//
	StartoverWindowSeconds *float64 `field:"optional" json:"startoverWindowSeconds" yaml:"startoverWindowSeconds"`
	// The tags associated with the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediapackagev2-originendpoint.html#cfn-mediapackagev2-originendpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

