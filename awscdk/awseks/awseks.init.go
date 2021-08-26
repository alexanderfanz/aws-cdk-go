package awseks

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.AutoScalingGroupCapacityOptions",
		reflect.TypeOf((*AutoScalingGroupCapacityOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.AutoScalingGroupOptions",
		reflect.TypeOf((*AutoScalingGroupOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.AwsAuth",
		reflect.TypeOf((*AwsAuth)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAccount", GoMethod: "AddAccount"},
			_jsii_.MemberMethod{JsiiMethod: "addMastersRole", GoMethod: "AddMastersRole"},
			_jsii_.MemberMethod{JsiiMethod: "addRoleMapping", GoMethod: "AddRoleMapping"},
			_jsii_.MemberMethod{JsiiMethod: "addUserMapping", GoMethod: "AddUserMapping"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_AwsAuth{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.AwsAuthMapping",
		reflect.TypeOf((*AwsAuthMapping)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.AwsAuthProps",
		reflect.TypeOf((*AwsAuthProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.BootstrapOptions",
		reflect.TypeOf((*BootstrapOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_eks.CapacityType",
		reflect.TypeOf((*CapacityType)(nil)).Elem(),
		map[string]interface{}{
			"SPOT": CapacityType_SPOT,
			"ON_DEMAND": CapacityType_ON_DEMAND,
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.CfnAddon",
		reflect.TypeOf((*CfnAddon)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "addonName", GoGetter: "AddonName"},
			_jsii_.MemberProperty{JsiiProperty: "addonVersion", GoGetter: "AddonVersion"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "resolveConflicts", GoGetter: "ResolveConflicts"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountRoleArn", GoGetter: "ServiceAccountRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAddon{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.CfnAddonProps",
		reflect.TypeOf((*CfnAddonProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.CfnCluster",
		reflect.TypeOf((*CfnCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrCertificateAuthorityData", GoGetter: "AttrCertificateAuthorityData"},
			_jsii_.MemberProperty{JsiiProperty: "attrClusterSecurityGroupId", GoGetter: "AttrClusterSecurityGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "attrEncryptionConfigKeyArn", GoGetter: "AttrEncryptionConfigKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrEndpoint", GoGetter: "AttrEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "attrOpenIdConnectIssuerUrl", GoGetter: "AttrOpenIdConnectIssuerUrl"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfig", GoGetter: "EncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "kubernetesNetworkConfig", GoGetter: "KubernetesNetworkConfig"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "resourcesVpcConfig", GoGetter: "ResourcesVpcConfig"},
			_jsii_.MemberProperty{JsiiProperty: "roleArn", GoGetter: "RoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.CfnCluster.EncryptionConfigProperty",
		reflect.TypeOf((*CfnCluster_EncryptionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.CfnCluster.KubernetesNetworkConfigProperty",
		reflect.TypeOf((*CfnCluster_KubernetesNetworkConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.CfnCluster.ProviderProperty",
		reflect.TypeOf((*CfnCluster_ProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.CfnCluster.ResourcesVpcConfigProperty",
		reflect.TypeOf((*CfnCluster_ResourcesVpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.CfnClusterProps",
		reflect.TypeOf((*CfnClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.CfnFargateProfile",
		reflect.TypeOf((*CfnFargateProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fargateProfileName", GoGetter: "FargateProfileName"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "podExecutionRoleArn", GoGetter: "PodExecutionRoleArn"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "selectors", GoGetter: "Selectors"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnets", GoGetter: "Subnets"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFargateProfile{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.CfnFargateProfile.LabelProperty",
		reflect.TypeOf((*CfnFargateProfile_LabelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.CfnFargateProfile.SelectorProperty",
		reflect.TypeOf((*CfnFargateProfile_SelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.CfnFargateProfileProps",
		reflect.TypeOf((*CfnFargateProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.CfnNodegroup",
		reflect.TypeOf((*CfnNodegroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberProperty{JsiiProperty: "amiType", GoGetter: "AmiType"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "attrClusterName", GoGetter: "AttrClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "attrNodegroupName", GoGetter: "AttrNodegroupName"},
			_jsii_.MemberProperty{JsiiProperty: "capacityType", GoGetter: "CapacityType"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diskSize", GoGetter: "DiskSize"},
			_jsii_.MemberProperty{JsiiProperty: "forceUpdateEnabled", GoGetter: "ForceUpdateEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTypes", GoGetter: "InstanceTypes"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "launchTemplate", GoGetter: "LaunchTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodegroupName", GoGetter: "NodegroupName"},
			_jsii_.MemberProperty{JsiiProperty: "nodeRole", GoGetter: "NodeRole"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberProperty{JsiiProperty: "releaseVersion", GoGetter: "ReleaseVersion"},
			_jsii_.MemberProperty{JsiiProperty: "remoteAccess", GoGetter: "RemoteAccess"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberProperty{JsiiProperty: "scalingConfig", GoGetter: "ScalingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "subnets", GoGetter: "Subnets"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "taints", GoGetter: "Taints"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updateConfig", GoGetter: "UpdateConfig"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNodegroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.CfnNodegroup.LaunchTemplateSpecificationProperty",
		reflect.TypeOf((*CfnNodegroup_LaunchTemplateSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.CfnNodegroup.RemoteAccessProperty",
		reflect.TypeOf((*CfnNodegroup_RemoteAccessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.CfnNodegroup.ScalingConfigProperty",
		reflect.TypeOf((*CfnNodegroup_ScalingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.CfnNodegroup.TaintProperty",
		reflect.TypeOf((*CfnNodegroup_TaintProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.CfnNodegroup.UpdateConfigProperty",
		reflect.TypeOf((*CfnNodegroup_UpdateConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.CfnNodegroupProps",
		reflect.TypeOf((*CfnNodegroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.Cluster",
		reflect.TypeOf((*Cluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAutoScalingGroupCapacity", GoMethod: "AddAutoScalingGroupCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "addCdk8sChart", GoMethod: "AddCdk8sChart"},
			_jsii_.MemberMethod{JsiiMethod: "addFargateProfile", GoMethod: "AddFargateProfile"},
			_jsii_.MemberMethod{JsiiMethod: "addHelmChart", GoMethod: "AddHelmChart"},
			_jsii_.MemberMethod{JsiiMethod: "addManifest", GoMethod: "AddManifest"},
			_jsii_.MemberMethod{JsiiMethod: "addNodegroupCapacity", GoMethod: "AddNodegroupCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "addServiceAccount", GoMethod: "AddServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "adminRole", GoGetter: "AdminRole"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "awsAuth", GoGetter: "AwsAuth"},
			_jsii_.MemberProperty{JsiiProperty: "clusterArn", GoGetter: "ClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterCertificateAuthorityData", GoGetter: "ClusterCertificateAuthorityData"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEncryptionConfigKeyArn", GoGetter: "ClusterEncryptionConfigKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "clusterOpenIdConnectIssuer", GoGetter: "ClusterOpenIdConnectIssuer"},
			_jsii_.MemberProperty{JsiiProperty: "clusterOpenIdConnectIssuerUrl", GoGetter: "ClusterOpenIdConnectIssuerUrl"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSecurityGroup", GoGetter: "ClusterSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSecurityGroupId", GoGetter: "ClusterSecurityGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "connectAutoScalingGroupCapacity", GoMethod: "ConnectAutoScalingGroupCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCapacity", GoGetter: "DefaultCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "defaultNodegroup", GoGetter: "DefaultNodegroup"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getServiceLoadBalancerAddress", GoMethod: "GetServiceLoadBalancerAddress"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlEnvironment", GoGetter: "KubectlEnvironment"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlLayer", GoGetter: "KubectlLayer"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlMemory", GoGetter: "KubectlMemory"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlPrivateSubnets", GoGetter: "KubectlPrivateSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlRole", GoGetter: "KubectlRole"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlSecurityGroup", GoGetter: "KubectlSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectProvider", GoGetter: "OpenIdConnectProvider"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "prune", GoGetter: "Prune"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_Cluster{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICluster)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.ClusterAttributes",
		reflect.TypeOf((*ClusterAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.ClusterOptions",
		reflect.TypeOf((*ClusterOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.ClusterProps",
		reflect.TypeOf((*ClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.CommonClusterOptions",
		reflect.TypeOf((*CommonClusterOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_eks.CoreDnsComputeType",
		reflect.TypeOf((*CoreDnsComputeType)(nil)).Elem(),
		map[string]interface{}{
			"EC2": CoreDnsComputeType_EC2,
			"FARGATE": CoreDnsComputeType_FARGATE,
		},
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_eks.CpuArch",
		reflect.TypeOf((*CpuArch)(nil)).Elem(),
		map[string]interface{}{
			"ARM_64": CpuArch_ARM_64,
			"X86_64": CpuArch_X86_64,
		},
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_eks.DefaultCapacityType",
		reflect.TypeOf((*DefaultCapacityType)(nil)).Elem(),
		map[string]interface{}{
			"NODEGROUP": DefaultCapacityType_NODEGROUP,
			"EC2": DefaultCapacityType_EC2,
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.EksOptimizedImage",
		reflect.TypeOf((*EksOptimizedImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "getImage", GoMethod: "GetImage"},
		},
		func() interface{} {
			j := jsiiProxy_EksOptimizedImage{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IMachineImage)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.EksOptimizedImageProps",
		reflect.TypeOf((*EksOptimizedImageProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.EndpointAccess",
		reflect.TypeOf((*EndpointAccess)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "onlyFrom", GoMethod: "OnlyFrom"},
		},
		func() interface{} {
			return &jsiiProxy_EndpointAccess{}
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.FargateCluster",
		reflect.TypeOf((*FargateCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addAutoScalingGroupCapacity", GoMethod: "AddAutoScalingGroupCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "addCdk8sChart", GoMethod: "AddCdk8sChart"},
			_jsii_.MemberMethod{JsiiMethod: "addFargateProfile", GoMethod: "AddFargateProfile"},
			_jsii_.MemberMethod{JsiiMethod: "addHelmChart", GoMethod: "AddHelmChart"},
			_jsii_.MemberMethod{JsiiMethod: "addManifest", GoMethod: "AddManifest"},
			_jsii_.MemberMethod{JsiiMethod: "addNodegroupCapacity", GoMethod: "AddNodegroupCapacity"},
			_jsii_.MemberMethod{JsiiMethod: "addServiceAccount", GoMethod: "AddServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "adminRole", GoGetter: "AdminRole"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "awsAuth", GoGetter: "AwsAuth"},
			_jsii_.MemberProperty{JsiiProperty: "clusterArn", GoGetter: "ClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterCertificateAuthorityData", GoGetter: "ClusterCertificateAuthorityData"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEncryptionConfigKeyArn", GoGetter: "ClusterEncryptionConfigKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "clusterOpenIdConnectIssuer", GoGetter: "ClusterOpenIdConnectIssuer"},
			_jsii_.MemberProperty{JsiiProperty: "clusterOpenIdConnectIssuerUrl", GoGetter: "ClusterOpenIdConnectIssuerUrl"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSecurityGroup", GoGetter: "ClusterSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSecurityGroupId", GoGetter: "ClusterSecurityGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "connectAutoScalingGroupCapacity", GoMethod: "ConnectAutoScalingGroupCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "defaultCapacity", GoGetter: "DefaultCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "defaultNodegroup", GoGetter: "DefaultNodegroup"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getServiceLoadBalancerAddress", GoMethod: "GetServiceLoadBalancerAddress"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlEnvironment", GoGetter: "KubectlEnvironment"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlLayer", GoGetter: "KubectlLayer"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlMemory", GoGetter: "KubectlMemory"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlPrivateSubnets", GoGetter: "KubectlPrivateSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlRole", GoGetter: "KubectlRole"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlSecurityGroup", GoGetter: "KubectlSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectProvider", GoGetter: "OpenIdConnectProvider"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "prune", GoGetter: "Prune"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_FargateCluster{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_Cluster)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.FargateClusterProps",
		reflect.TypeOf((*FargateClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.FargateProfile",
		reflect.TypeOf((*FargateProfile)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "fargateProfileArn", GoGetter: "FargateProfileArn"},
			_jsii_.MemberProperty{JsiiProperty: "fargateProfileName", GoGetter: "FargateProfileName"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "podExecutionRole", GoGetter: "PodExecutionRole"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_FargateProfile{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.FargateProfileOptions",
		reflect.TypeOf((*FargateProfileOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.FargateProfileProps",
		reflect.TypeOf((*FargateProfileProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.HelmChart",
		reflect.TypeOf((*HelmChart)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_HelmChart{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.HelmChartOptions",
		reflect.TypeOf((*HelmChartOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.HelmChartProps",
		reflect.TypeOf((*HelmChartProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_eks.ICluster",
		reflect.TypeOf((*ICluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addCdk8sChart", GoMethod: "AddCdk8sChart"},
			_jsii_.MemberMethod{JsiiMethod: "addHelmChart", GoMethod: "AddHelmChart"},
			_jsii_.MemberMethod{JsiiMethod: "addManifest", GoMethod: "AddManifest"},
			_jsii_.MemberMethod{JsiiMethod: "addServiceAccount", GoMethod: "AddServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "clusterArn", GoGetter: "ClusterArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterCertificateAuthorityData", GoGetter: "ClusterCertificateAuthorityData"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEncryptionConfigKeyArn", GoGetter: "ClusterEncryptionConfigKeyArn"},
			_jsii_.MemberProperty{JsiiProperty: "clusterEndpoint", GoGetter: "ClusterEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSecurityGroup", GoGetter: "ClusterSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSecurityGroupId", GoGetter: "ClusterSecurityGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlEnvironment", GoGetter: "KubectlEnvironment"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlLayer", GoGetter: "KubectlLayer"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlMemory", GoGetter: "KubectlMemory"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlPrivateSubnets", GoGetter: "KubectlPrivateSubnets"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlRole", GoGetter: "KubectlRole"},
			_jsii_.MemberProperty{JsiiProperty: "kubectlSecurityGroup", GoGetter: "KubectlSecurityGroup"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectProvider", GoGetter: "OpenIdConnectProvider"},
			_jsii_.MemberProperty{JsiiProperty: "prune", GoGetter: "Prune"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_ICluster{}
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"monocdk.aws_eks.INodegroup",
		reflect.TypeOf((*INodegroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodegroupName", GoGetter: "NodegroupName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_INodegroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.KubernetesManifest",
		reflect.TypeOf((*KubernetesManifest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_KubernetesManifest{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.KubernetesManifestOptions",
		reflect.TypeOf((*KubernetesManifestOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.KubernetesManifestProps",
		reflect.TypeOf((*KubernetesManifestProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.KubernetesObjectValue",
		reflect.TypeOf((*KubernetesObjectValue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
		},
		func() interface{} {
			j := jsiiProxy_KubernetesObjectValue{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.KubernetesObjectValueProps",
		reflect.TypeOf((*KubernetesObjectValueProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.KubernetesPatch",
		reflect.TypeOf((*KubernetesPatch)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_KubernetesPatch{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.KubernetesPatchProps",
		reflect.TypeOf((*KubernetesPatchProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.KubernetesVersion",
		reflect.TypeOf((*KubernetesVersion)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_KubernetesVersion{}
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.LaunchTemplateSpec",
		reflect.TypeOf((*LaunchTemplateSpec)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_eks.MachineImageType",
		reflect.TypeOf((*MachineImageType)(nil)).Elem(),
		map[string]interface{}{
			"AMAZON_LINUX_2": MachineImageType_AMAZON_LINUX_2,
			"BOTTLEROCKET": MachineImageType_BOTTLEROCKET,
		},
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_eks.NodeType",
		reflect.TypeOf((*NodeType)(nil)).Elem(),
		map[string]interface{}{
			"STANDARD": NodeType_STANDARD,
			"GPU": NodeType_GPU,
			"INFERENTIA": NodeType_INFERENTIA,
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.Nodegroup",
		reflect.TypeOf((*Nodegroup)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodegroupArn", GoGetter: "NodegroupArn"},
			_jsii_.MemberProperty{JsiiProperty: "nodegroupName", GoGetter: "NodegroupName"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_Nodegroup{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_INodegroup)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_eks.NodegroupAmiType",
		reflect.TypeOf((*NodegroupAmiType)(nil)).Elem(),
		map[string]interface{}{
			"AL2_X86_64": NodegroupAmiType_AL2_X86_64,
			"AL2_X86_64_GPU": NodegroupAmiType_AL2_X86_64_GPU,
			"AL2_ARM_64": NodegroupAmiType_AL2_ARM_64,
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.NodegroupOptions",
		reflect.TypeOf((*NodegroupOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.NodegroupProps",
		reflect.TypeOf((*NodegroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.NodegroupRemoteAccess",
		reflect.TypeOf((*NodegroupRemoteAccess)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.OpenIdConnectProvider",
		reflect.TypeOf((*OpenIdConnectProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectProviderArn", GoGetter: "OpenIdConnectProviderArn"},
			_jsii_.MemberProperty{JsiiProperty: "openIdConnectProviderIssuer", GoGetter: "OpenIdConnectProviderIssuer"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_OpenIdConnectProvider{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamOpenIdConnectProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.OpenIdConnectProviderProps",
		reflect.TypeOf((*OpenIdConnectProviderProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_eks.PatchType",
		reflect.TypeOf((*PatchType)(nil)).Elem(),
		map[string]interface{}{
			"JSON": PatchType_JSON,
			"MERGE": PatchType_MERGE,
			"STRATEGIC": PatchType_STRATEGIC,
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.Selector",
		reflect.TypeOf((*Selector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_eks.ServiceAccount",
		reflect.TypeOf((*ServiceAccount)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addToPolicy", GoMethod: "AddToPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "addToPrincipalPolicy", GoMethod: "AddToPrincipalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "assumeRoleAction", GoGetter: "AssumeRoleAction"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "onPrepare", GoMethod: "OnPrepare"},
			_jsii_.MemberMethod{JsiiMethod: "onSynthesize", GoMethod: "OnSynthesize"},
			_jsii_.MemberMethod{JsiiMethod: "onValidate", GoMethod: "OnValidate"},
			_jsii_.MemberProperty{JsiiProperty: "policyFragment", GoGetter: "PolicyFragment"},
			_jsii_.MemberMethod{JsiiMethod: "prepare", GoMethod: "Prepare"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountName", GoGetter: "ServiceAccountName"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountNamespace", GoGetter: "ServiceAccountNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "synthesize", GoMethod: "Synthesize"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validate", GoMethod: "Validate"},
		},
		func() interface{} {
			j := jsiiProxy_ServiceAccount{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIPrincipal)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.ServiceAccountOptions",
		reflect.TypeOf((*ServiceAccountOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.ServiceAccountProps",
		reflect.TypeOf((*ServiceAccountProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.ServiceLoadBalancerAddressOptions",
		reflect.TypeOf((*ServiceLoadBalancerAddressOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_eks.TaintEffect",
		reflect.TypeOf((*TaintEffect)(nil)).Elem(),
		map[string]interface{}{
			"NO_SCHEDULE": TaintEffect_NO_SCHEDULE,
			"PREFER_NO_SCHEDULE": TaintEffect_PREFER_NO_SCHEDULE,
			"NO_EXECUTE": TaintEffect_NO_EXECUTE,
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_eks.TaintSpec",
		reflect.TypeOf((*TaintSpec)(nil)).Elem(),
	)
}
