package awsec2


// Options for a ClientVpnAuthorizationRule.
//
// Example:
//   endpoint := vpc.addClientVpnEndpoint(jsii.String("Endpoint"), &clientVpnEndpointOptions{
//   	cidr: jsii.String("10.100.0.0/16"),
//   	serverCertificateArn: jsii.String("arn:aws:acm:us-east-1:123456789012:certificate/server-certificate-id"),
//   	userBasedAuthentication: ec2.clientVpnUserBasedAuthentication.federated(samlProvider),
//   	authorizeAllUsersToVpcCidr: jsii.Boolean(false),
//   })
//
//   endpoint.addAuthorizationRule(jsii.String("Rule"), &clientVpnAuthorizationRuleOptions{
//   	cidr: jsii.String("10.0.10.0/32"),
//   	groupId: jsii.String("group-id"),
//   })
//
// Experimental.
type ClientVpnAuthorizationRuleOptions struct {
	// The IPv4 address range, in CIDR notation, of the network for which access is being authorized.
	// Experimental.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// A brief description of the authorization rule.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the group to grant access to, for example, the Active Directory group or identity provider (IdP) group.
	// Experimental.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
}

