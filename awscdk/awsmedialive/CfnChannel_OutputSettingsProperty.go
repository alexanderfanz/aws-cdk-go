package awsmedialive


// The output settings.
//
// The parent of this entity is Output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputSettingsProperty := &OutputSettingsProperty{
//   	ArchiveOutputSettings: &ArchiveOutputSettingsProperty{
//   		ContainerSettings: &ArchiveContainerSettingsProperty{
//   			M2TsSettings: &M2tsSettingsProperty{
//   				AbsentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   				Arib: jsii.String("arib"),
//   				AribCaptionsPid: jsii.String("aribCaptionsPid"),
//   				AribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   				AudioBufferModel: jsii.String("audioBufferModel"),
//   				AudioFramesPerPes: jsii.Number(123),
//   				AudioPids: jsii.String("audioPids"),
//   				AudioStreamType: jsii.String("audioStreamType"),
//   				Bitrate: jsii.Number(123),
//   				BufferModel: jsii.String("bufferModel"),
//   				CcDescriptor: jsii.String("ccDescriptor"),
//   				DvbNitSettings: &DvbNitSettingsProperty{
//   					NetworkId: jsii.Number(123),
//   					NetworkName: jsii.String("networkName"),
//   					RepInterval: jsii.Number(123),
//   				},
//   				DvbSdtSettings: &DvbSdtSettingsProperty{
//   					OutputSdt: jsii.String("outputSdt"),
//   					RepInterval: jsii.Number(123),
//   					ServiceName: jsii.String("serviceName"),
//   					ServiceProviderName: jsii.String("serviceProviderName"),
//   				},
//   				DvbSubPids: jsii.String("dvbSubPids"),
//   				DvbTdtSettings: &DvbTdtSettingsProperty{
//   					RepInterval: jsii.Number(123),
//   				},
//   				DvbTeletextPid: jsii.String("dvbTeletextPid"),
//   				Ebif: jsii.String("ebif"),
//   				EbpAudioInterval: jsii.String("ebpAudioInterval"),
//   				EbpLookaheadMs: jsii.Number(123),
//   				EbpPlacement: jsii.String("ebpPlacement"),
//   				EcmPid: jsii.String("ecmPid"),
//   				EsRateInPes: jsii.String("esRateInPes"),
//   				EtvPlatformPid: jsii.String("etvPlatformPid"),
//   				EtvSignalPid: jsii.String("etvSignalPid"),
//   				FragmentTime: jsii.Number(123),
//   				Klv: jsii.String("klv"),
//   				KlvDataPids: jsii.String("klvDataPids"),
//   				NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   				NullPacketBitrate: jsii.Number(123),
//   				PatInterval: jsii.Number(123),
//   				PcrControl: jsii.String("pcrControl"),
//   				PcrPeriod: jsii.Number(123),
//   				PcrPid: jsii.String("pcrPid"),
//   				PmtInterval: jsii.Number(123),
//   				PmtPid: jsii.String("pmtPid"),
//   				ProgramNum: jsii.Number(123),
//   				RateMode: jsii.String("rateMode"),
//   				Scte27Pids: jsii.String("scte27Pids"),
//   				Scte35Control: jsii.String("scte35Control"),
//   				Scte35Pid: jsii.String("scte35Pid"),
//   				Scte35PrerollPullupMilliseconds: jsii.Number(123),
//   				SegmentationMarkers: jsii.String("segmentationMarkers"),
//   				SegmentationStyle: jsii.String("segmentationStyle"),
//   				SegmentationTime: jsii.Number(123),
//   				TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   				TimedMetadataPid: jsii.String("timedMetadataPid"),
//   				TransportStreamId: jsii.Number(123),
//   				VideoPid: jsii.String("videoPid"),
//   			},
//   			RawSettings: &RawSettingsProperty{
//   			},
//   		},
//   		Extension: jsii.String("extension"),
//   		NameModifier: jsii.String("nameModifier"),
//   	},
//   	FrameCaptureOutputSettings: &FrameCaptureOutputSettingsProperty{
//   		NameModifier: jsii.String("nameModifier"),
//   	},
//   	HlsOutputSettings: &HlsOutputSettingsProperty{
//   		H265PackagingType: jsii.String("h265PackagingType"),
//   		HlsSettings: &HlsSettingsProperty{
//   			AudioOnlyHlsSettings: &AudioOnlyHlsSettingsProperty{
//   				AudioGroupId: jsii.String("audioGroupId"),
//   				AudioOnlyImage: &InputLocationProperty{
//   					PasswordParam: jsii.String("passwordParam"),
//   					Uri: jsii.String("uri"),
//   					Username: jsii.String("username"),
//   				},
//   				AudioTrackType: jsii.String("audioTrackType"),
//   				SegmentType: jsii.String("segmentType"),
//   			},
//   			Fmp4HlsSettings: &Fmp4HlsSettingsProperty{
//   				AudioRenditionSets: jsii.String("audioRenditionSets"),
//   				NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   				TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   			},
//   			FrameCaptureHlsSettings: &FrameCaptureHlsSettingsProperty{
//   			},
//   			StandardHlsSettings: &StandardHlsSettingsProperty{
//   				AudioRenditionSets: jsii.String("audioRenditionSets"),
//   				M3U8Settings: &M3u8SettingsProperty{
//   					AudioFramesPerPes: jsii.Number(123),
//   					AudioPids: jsii.String("audioPids"),
//   					EcmPid: jsii.String("ecmPid"),
//   					NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   					PatInterval: jsii.Number(123),
//   					PcrControl: jsii.String("pcrControl"),
//   					PcrPeriod: jsii.Number(123),
//   					PcrPid: jsii.String("pcrPid"),
//   					PmtInterval: jsii.Number(123),
//   					PmtPid: jsii.String("pmtPid"),
//   					ProgramNum: jsii.Number(123),
//   					Scte35Behavior: jsii.String("scte35Behavior"),
//   					Scte35Pid: jsii.String("scte35Pid"),
//   					TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   					TimedMetadataPid: jsii.String("timedMetadataPid"),
//   					TransportStreamId: jsii.Number(123),
//   					VideoPid: jsii.String("videoPid"),
//   				},
//   			},
//   		},
//   		NameModifier: jsii.String("nameModifier"),
//   		SegmentModifier: jsii.String("segmentModifier"),
//   	},
//   	MediaPackageOutputSettings: &MediaPackageOutputSettingsProperty{
//   	},
//   	MsSmoothOutputSettings: &MsSmoothOutputSettingsProperty{
//   		H265PackagingType: jsii.String("h265PackagingType"),
//   		NameModifier: jsii.String("nameModifier"),
//   	},
//   	MultiplexOutputSettings: &MultiplexOutputSettingsProperty{
//   		Destination: &OutputLocationRefProperty{
//   			DestinationRefId: jsii.String("destinationRefId"),
//   		},
//   	},
//   	RtmpOutputSettings: &RtmpOutputSettingsProperty{
//   		CertificateMode: jsii.String("certificateMode"),
//   		ConnectionRetryInterval: jsii.Number(123),
//   		Destination: &OutputLocationRefProperty{
//   			DestinationRefId: jsii.String("destinationRefId"),
//   		},
//   		NumRetries: jsii.Number(123),
//   	},
//   	UdpOutputSettings: &UdpOutputSettingsProperty{
//   		BufferMsec: jsii.Number(123),
//   		ContainerSettings: &UdpContainerSettingsProperty{
//   			M2TsSettings: &M2tsSettingsProperty{
//   				AbsentInputAudioBehavior: jsii.String("absentInputAudioBehavior"),
//   				Arib: jsii.String("arib"),
//   				AribCaptionsPid: jsii.String("aribCaptionsPid"),
//   				AribCaptionsPidControl: jsii.String("aribCaptionsPidControl"),
//   				AudioBufferModel: jsii.String("audioBufferModel"),
//   				AudioFramesPerPes: jsii.Number(123),
//   				AudioPids: jsii.String("audioPids"),
//   				AudioStreamType: jsii.String("audioStreamType"),
//   				Bitrate: jsii.Number(123),
//   				BufferModel: jsii.String("bufferModel"),
//   				CcDescriptor: jsii.String("ccDescriptor"),
//   				DvbNitSettings: &DvbNitSettingsProperty{
//   					NetworkId: jsii.Number(123),
//   					NetworkName: jsii.String("networkName"),
//   					RepInterval: jsii.Number(123),
//   				},
//   				DvbSdtSettings: &DvbSdtSettingsProperty{
//   					OutputSdt: jsii.String("outputSdt"),
//   					RepInterval: jsii.Number(123),
//   					ServiceName: jsii.String("serviceName"),
//   					ServiceProviderName: jsii.String("serviceProviderName"),
//   				},
//   				DvbSubPids: jsii.String("dvbSubPids"),
//   				DvbTdtSettings: &DvbTdtSettingsProperty{
//   					RepInterval: jsii.Number(123),
//   				},
//   				DvbTeletextPid: jsii.String("dvbTeletextPid"),
//   				Ebif: jsii.String("ebif"),
//   				EbpAudioInterval: jsii.String("ebpAudioInterval"),
//   				EbpLookaheadMs: jsii.Number(123),
//   				EbpPlacement: jsii.String("ebpPlacement"),
//   				EcmPid: jsii.String("ecmPid"),
//   				EsRateInPes: jsii.String("esRateInPes"),
//   				EtvPlatformPid: jsii.String("etvPlatformPid"),
//   				EtvSignalPid: jsii.String("etvSignalPid"),
//   				FragmentTime: jsii.Number(123),
//   				Klv: jsii.String("klv"),
//   				KlvDataPids: jsii.String("klvDataPids"),
//   				NielsenId3Behavior: jsii.String("nielsenId3Behavior"),
//   				NullPacketBitrate: jsii.Number(123),
//   				PatInterval: jsii.Number(123),
//   				PcrControl: jsii.String("pcrControl"),
//   				PcrPeriod: jsii.Number(123),
//   				PcrPid: jsii.String("pcrPid"),
//   				PmtInterval: jsii.Number(123),
//   				PmtPid: jsii.String("pmtPid"),
//   				ProgramNum: jsii.Number(123),
//   				RateMode: jsii.String("rateMode"),
//   				Scte27Pids: jsii.String("scte27Pids"),
//   				Scte35Control: jsii.String("scte35Control"),
//   				Scte35Pid: jsii.String("scte35Pid"),
//   				Scte35PrerollPullupMilliseconds: jsii.Number(123),
//   				SegmentationMarkers: jsii.String("segmentationMarkers"),
//   				SegmentationStyle: jsii.String("segmentationStyle"),
//   				SegmentationTime: jsii.Number(123),
//   				TimedMetadataBehavior: jsii.String("timedMetadataBehavior"),
//   				TimedMetadataPid: jsii.String("timedMetadataPid"),
//   				TransportStreamId: jsii.Number(123),
//   				VideoPid: jsii.String("videoPid"),
//   			},
//   		},
//   		Destination: &OutputLocationRefProperty{
//   			DestinationRefId: jsii.String("destinationRefId"),
//   		},
//   		FecOutputSettings: &FecOutputSettingsProperty{
//   			ColumnDepth: jsii.Number(123),
//   			IncludeFec: jsii.String("includeFec"),
//   			RowLength: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnChannel_OutputSettingsProperty struct {
	// The settings for an archive output.
	ArchiveOutputSettings interface{} `field:"optional" json:"archiveOutputSettings" yaml:"archiveOutputSettings"`
	// The settings for a frame capture output.
	//
	// The parent of this entity is OutputGroupSettings.
	FrameCaptureOutputSettings interface{} `field:"optional" json:"frameCaptureOutputSettings" yaml:"frameCaptureOutputSettings"`
	// The settings for an HLS output.
	//
	// The parent of this entity is OutputGroupSettings.
	HlsOutputSettings interface{} `field:"optional" json:"hlsOutputSettings" yaml:"hlsOutputSettings"`
	// The settings for a MediaPackage output.
	//
	// The parent of this entity is OutputGroupSettings.
	MediaPackageOutputSettings interface{} `field:"optional" json:"mediaPackageOutputSettings" yaml:"mediaPackageOutputSettings"`
	// The settings for a Microsoft Smooth output.
	MsSmoothOutputSettings interface{} `field:"optional" json:"msSmoothOutputSettings" yaml:"msSmoothOutputSettings"`
	// Configuration of a Multiplex output.
	MultiplexOutputSettings interface{} `field:"optional" json:"multiplexOutputSettings" yaml:"multiplexOutputSettings"`
	// The settings for an RTMP output.
	//
	// The parent of this entity is OutputGroupSettings.
	RtmpOutputSettings interface{} `field:"optional" json:"rtmpOutputSettings" yaml:"rtmpOutputSettings"`
	// The settings for a UDP output.
	//
	// The parent of this entity is OutputGroupSettings.
	UdpOutputSettings interface{} `field:"optional" json:"udpOutputSettings" yaml:"udpOutputSettings"`
}
