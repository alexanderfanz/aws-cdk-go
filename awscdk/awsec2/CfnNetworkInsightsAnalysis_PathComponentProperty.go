package awsec2


// Describes a path component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pathComponentProperty := &PathComponentProperty{
//   	AclRule: &AnalysisAclRuleProperty{
//   		Cidr: jsii.String("cidr"),
//   		Egress: jsii.Boolean(false),
//   		PortRange: &PortRangeProperty{
//   			From: jsii.Number(123),
//   			To: jsii.Number(123),
//   		},
//   		Protocol: jsii.String("protocol"),
//   		RuleAction: jsii.String("ruleAction"),
//   		RuleNumber: jsii.Number(123),
//   	},
//   	AdditionalDetails: []interface{}{
//   		&AdditionalDetailProperty{
//   			AdditionalDetailType: jsii.String("additionalDetailType"),
//   			Component: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   		},
//   	},
//   	Component: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	DestinationVpc: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	ElasticLoadBalancerListener: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	Explanations: []interface{}{
//   		&ExplanationProperty{
//   			Acl: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			AclRule: &AnalysisAclRuleProperty{
//   				Cidr: jsii.String("cidr"),
//   				Egress: jsii.Boolean(false),
//   				PortRange: &PortRangeProperty{
//   					From: jsii.Number(123),
//   					To: jsii.Number(123),
//   				},
//   				Protocol: jsii.String("protocol"),
//   				RuleAction: jsii.String("ruleAction"),
//   				RuleNumber: jsii.Number(123),
//   			},
//   			Address: jsii.String("address"),
//   			Addresses: []*string{
//   				jsii.String("addresses"),
//   			},
//   			AttachedTo: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			AvailabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			Cidrs: []*string{
//   				jsii.String("cidrs"),
//   			},
//   			ClassicLoadBalancerListener: &AnalysisLoadBalancerListenerProperty{
//   				InstancePort: jsii.Number(123),
//   				LoadBalancerPort: jsii.Number(123),
//   			},
//   			Component: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			ComponentAccount: jsii.String("componentAccount"),
//   			ComponentRegion: jsii.String("componentRegion"),
//   			CustomerGateway: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			Destination: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			DestinationVpc: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			Direction: jsii.String("direction"),
//   			ElasticLoadBalancerListener: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			ExplanationCode: jsii.String("explanationCode"),
//   			IngressRouteTable: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			InternetGateway: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			LoadBalancerArn: jsii.String("loadBalancerArn"),
//   			LoadBalancerListenerPort: jsii.Number(123),
//   			LoadBalancerTarget: &AnalysisLoadBalancerTargetProperty{
//   				Address: jsii.String("address"),
//   				AvailabilityZone: jsii.String("availabilityZone"),
//   				Instance: &AnalysisComponentProperty{
//   					Arn: jsii.String("arn"),
//   					Id: jsii.String("id"),
//   				},
//   				Port: jsii.Number(123),
//   			},
//   			LoadBalancerTargetGroup: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			LoadBalancerTargetGroups: []interface{}{
//   				&AnalysisComponentProperty{
//   					Arn: jsii.String("arn"),
//   					Id: jsii.String("id"),
//   				},
//   			},
//   			LoadBalancerTargetPort: jsii.Number(123),
//   			MissingComponent: jsii.String("missingComponent"),
//   			NatGateway: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			NetworkInterface: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			PacketField: jsii.String("packetField"),
//   			Port: jsii.Number(123),
//   			PortRanges: []interface{}{
//   				&PortRangeProperty{
//   					From: jsii.Number(123),
//   					To: jsii.Number(123),
//   				},
//   			},
//   			PrefixList: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			Protocols: []*string{
//   				jsii.String("protocols"),
//   			},
//   			RouteTable: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			RouteTableRoute: &AnalysisRouteTableRouteProperty{
//   				DestinationCidr: jsii.String("destinationCidr"),
//   				DestinationPrefixListId: jsii.String("destinationPrefixListId"),
//   				EgressOnlyInternetGatewayId: jsii.String("egressOnlyInternetGatewayId"),
//   				GatewayId: jsii.String("gatewayId"),
//   				InstanceId: jsii.String("instanceId"),
//   				NatGatewayId: jsii.String("natGatewayId"),
//   				NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   				Origin: jsii.String("origin"),
//   				State: jsii.String("state"),
//   				TransitGatewayId: jsii.String("transitGatewayId"),
//   				VpcPeeringConnectionId: jsii.String("vpcPeeringConnectionId"),
//   			},
//   			SecurityGroup: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			SecurityGroupRule: &AnalysisSecurityGroupRuleProperty{
//   				Cidr: jsii.String("cidr"),
//   				Direction: jsii.String("direction"),
//   				PortRange: &PortRangeProperty{
//   					From: jsii.Number(123),
//   					To: jsii.Number(123),
//   				},
//   				PrefixListId: jsii.String("prefixListId"),
//   				Protocol: jsii.String("protocol"),
//   				SecurityGroupId: jsii.String("securityGroupId"),
//   			},
//   			SecurityGroups: []interface{}{
//   				&AnalysisComponentProperty{
//   					Arn: jsii.String("arn"),
//   					Id: jsii.String("id"),
//   				},
//   			},
//   			SourceVpc: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			State: jsii.String("state"),
//   			Subnet: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			SubnetRouteTable: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			TransitGateway: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			TransitGatewayAttachment: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			TransitGatewayRouteTable: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			TransitGatewayRouteTableRoute: &TransitGatewayRouteTableRouteProperty{
//   				AttachmentId: jsii.String("attachmentId"),
//   				DestinationCidr: jsii.String("destinationCidr"),
//   				PrefixListId: jsii.String("prefixListId"),
//   				ResourceId: jsii.String("resourceId"),
//   				ResourceType: jsii.String("resourceType"),
//   				RouteOrigin: jsii.String("routeOrigin"),
//   				State: jsii.String("state"),
//   			},
//   			Vpc: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			VpcEndpoint: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			VpcPeeringConnection: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			VpnConnection: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   			VpnGateway: &AnalysisComponentProperty{
//   				Arn: jsii.String("arn"),
//   				Id: jsii.String("id"),
//   			},
//   		},
//   	},
//   	InboundHeader: &AnalysisPacketHeaderProperty{
//   		DestinationAddresses: []*string{
//   			jsii.String("destinationAddresses"),
//   		},
//   		DestinationPortRanges: []interface{}{
//   			&PortRangeProperty{
//   				From: jsii.Number(123),
//   				To: jsii.Number(123),
//   			},
//   		},
//   		Protocol: jsii.String("protocol"),
//   		SourceAddresses: []*string{
//   			jsii.String("sourceAddresses"),
//   		},
//   		SourcePortRanges: []interface{}{
//   			&PortRangeProperty{
//   				From: jsii.Number(123),
//   				To: jsii.Number(123),
//   			},
//   		},
//   	},
//   	OutboundHeader: &AnalysisPacketHeaderProperty{
//   		DestinationAddresses: []*string{
//   			jsii.String("destinationAddresses"),
//   		},
//   		DestinationPortRanges: []interface{}{
//   			&PortRangeProperty{
//   				From: jsii.Number(123),
//   				To: jsii.Number(123),
//   			},
//   		},
//   		Protocol: jsii.String("protocol"),
//   		SourceAddresses: []*string{
//   			jsii.String("sourceAddresses"),
//   		},
//   		SourcePortRanges: []interface{}{
//   			&PortRangeProperty{
//   				From: jsii.Number(123),
//   				To: jsii.Number(123),
//   			},
//   		},
//   	},
//   	RouteTableRoute: &AnalysisRouteTableRouteProperty{
//   		DestinationCidr: jsii.String("destinationCidr"),
//   		DestinationPrefixListId: jsii.String("destinationPrefixListId"),
//   		EgressOnlyInternetGatewayId: jsii.String("egressOnlyInternetGatewayId"),
//   		GatewayId: jsii.String("gatewayId"),
//   		InstanceId: jsii.String("instanceId"),
//   		NatGatewayId: jsii.String("natGatewayId"),
//   		NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   		Origin: jsii.String("origin"),
//   		State: jsii.String("state"),
//   		TransitGatewayId: jsii.String("transitGatewayId"),
//   		VpcPeeringConnectionId: jsii.String("vpcPeeringConnectionId"),
//   	},
//   	SecurityGroupRule: &AnalysisSecurityGroupRuleProperty{
//   		Cidr: jsii.String("cidr"),
//   		Direction: jsii.String("direction"),
//   		PortRange: &PortRangeProperty{
//   			From: jsii.Number(123),
//   			To: jsii.Number(123),
//   		},
//   		PrefixListId: jsii.String("prefixListId"),
//   		Protocol: jsii.String("protocol"),
//   		SecurityGroupId: jsii.String("securityGroupId"),
//   	},
//   	SequenceNumber: jsii.Number(123),
//   	SourceVpc: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	Subnet: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	TransitGateway: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   	TransitGatewayRouteTableRoute: &TransitGatewayRouteTableRouteProperty{
//   		AttachmentId: jsii.String("attachmentId"),
//   		DestinationCidr: jsii.String("destinationCidr"),
//   		PrefixListId: jsii.String("prefixListId"),
//   		ResourceId: jsii.String("resourceId"),
//   		ResourceType: jsii.String("resourceType"),
//   		RouteOrigin: jsii.String("routeOrigin"),
//   		State: jsii.String("state"),
//   	},
//   	Vpc: &AnalysisComponentProperty{
//   		Arn: jsii.String("arn"),
//   		Id: jsii.String("id"),
//   	},
//   }
//
type CfnNetworkInsightsAnalysis_PathComponentProperty struct {
	// The network ACL rule.
	AclRule interface{} `field:"optional" json:"aclRule" yaml:"aclRule"`
	// The additional details.
	AdditionalDetails interface{} `field:"optional" json:"additionalDetails" yaml:"additionalDetails"`
	// The component.
	Component interface{} `field:"optional" json:"component" yaml:"component"`
	// The destination VPC.
	DestinationVpc interface{} `field:"optional" json:"destinationVpc" yaml:"destinationVpc"`
	// The load balancer listener.
	ElasticLoadBalancerListener interface{} `field:"optional" json:"elasticLoadBalancerListener" yaml:"elasticLoadBalancerListener"`
	// The explanation codes.
	Explanations interface{} `field:"optional" json:"explanations" yaml:"explanations"`
	// The inbound header.
	InboundHeader interface{} `field:"optional" json:"inboundHeader" yaml:"inboundHeader"`
	// The outbound header.
	OutboundHeader interface{} `field:"optional" json:"outboundHeader" yaml:"outboundHeader"`
	// The route table route.
	RouteTableRoute interface{} `field:"optional" json:"routeTableRoute" yaml:"routeTableRoute"`
	// The security group rule.
	SecurityGroupRule interface{} `field:"optional" json:"securityGroupRule" yaml:"securityGroupRule"`
	// The sequence number.
	SequenceNumber *float64 `field:"optional" json:"sequenceNumber" yaml:"sequenceNumber"`
	// The source VPC.
	SourceVpc interface{} `field:"optional" json:"sourceVpc" yaml:"sourceVpc"`
	// The subnet.
	Subnet interface{} `field:"optional" json:"subnet" yaml:"subnet"`
	// The transit gateway.
	TransitGateway interface{} `field:"optional" json:"transitGateway" yaml:"transitGateway"`
	// The route in a transit gateway route table.
	TransitGatewayRouteTableRoute interface{} `field:"optional" json:"transitGatewayRouteTableRoute" yaml:"transitGatewayRouteTableRoute"`
	// The component VPC.
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}

