package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDistribution`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDistributionProps := &CfnDistributionProps{
//   	DistributionConfig: &DistributionConfigProperty{
//   		DefaultCacheBehavior: &DefaultCacheBehaviorProperty{
//   			TargetOriginId: jsii.String("targetOriginId"),
//   			ViewerProtocolPolicy: jsii.String("viewerProtocolPolicy"),
//
//   			// the properties below are optional
//   			AllowedMethods: []*string{
//   				jsii.String("allowedMethods"),
//   			},
//   			CachedMethods: []*string{
//   				jsii.String("cachedMethods"),
//   			},
//   			CachePolicyId: jsii.String("cachePolicyId"),
//   			Compress: jsii.Boolean(false),
//   			DefaultTtl: jsii.Number(123),
//   			FieldLevelEncryptionId: jsii.String("fieldLevelEncryptionId"),
//   			ForwardedValues: &ForwardedValuesProperty{
//   				QueryString: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				Cookies: &CookiesProperty{
//   					Forward: jsii.String("forward"),
//
//   					// the properties below are optional
//   					WhitelistedNames: []*string{
//   						jsii.String("whitelistedNames"),
//   					},
//   				},
//   				Headers: []*string{
//   					jsii.String("headers"),
//   				},
//   				QueryStringCacheKeys: []*string{
//   					jsii.String("queryStringCacheKeys"),
//   				},
//   			},
//   			FunctionAssociations: []interface{}{
//   				&FunctionAssociationProperty{
//   					EventType: jsii.String("eventType"),
//   					FunctionArn: jsii.String("functionArn"),
//   				},
//   			},
//   			LambdaFunctionAssociations: []interface{}{
//   				&LambdaFunctionAssociationProperty{
//   					EventType: jsii.String("eventType"),
//   					IncludeBody: jsii.Boolean(false),
//   					LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   				},
//   			},
//   			MaxTtl: jsii.Number(123),
//   			MinTtl: jsii.Number(123),
//   			OriginRequestPolicyId: jsii.String("originRequestPolicyId"),
//   			RealtimeLogConfigArn: jsii.String("realtimeLogConfigArn"),
//   			ResponseHeadersPolicyId: jsii.String("responseHeadersPolicyId"),
//   			SmoothStreaming: jsii.Boolean(false),
//   			TrustedKeyGroups: []*string{
//   				jsii.String("trustedKeyGroups"),
//   			},
//   			TrustedSigners: []*string{
//   				jsii.String("trustedSigners"),
//   			},
//   		},
//   		Enabled: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		Aliases: []*string{
//   			jsii.String("aliases"),
//   		},
//   		CacheBehaviors: []interface{}{
//   			&CacheBehaviorProperty{
//   				PathPattern: jsii.String("pathPattern"),
//   				TargetOriginId: jsii.String("targetOriginId"),
//   				ViewerProtocolPolicy: jsii.String("viewerProtocolPolicy"),
//
//   				// the properties below are optional
//   				AllowedMethods: []*string{
//   					jsii.String("allowedMethods"),
//   				},
//   				CachedMethods: []*string{
//   					jsii.String("cachedMethods"),
//   				},
//   				CachePolicyId: jsii.String("cachePolicyId"),
//   				Compress: jsii.Boolean(false),
//   				DefaultTtl: jsii.Number(123),
//   				FieldLevelEncryptionId: jsii.String("fieldLevelEncryptionId"),
//   				ForwardedValues: &ForwardedValuesProperty{
//   					QueryString: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					Cookies: &CookiesProperty{
//   						Forward: jsii.String("forward"),
//
//   						// the properties below are optional
//   						WhitelistedNames: []*string{
//   							jsii.String("whitelistedNames"),
//   						},
//   					},
//   					Headers: []*string{
//   						jsii.String("headers"),
//   					},
//   					QueryStringCacheKeys: []*string{
//   						jsii.String("queryStringCacheKeys"),
//   					},
//   				},
//   				FunctionAssociations: []interface{}{
//   					&FunctionAssociationProperty{
//   						EventType: jsii.String("eventType"),
//   						FunctionArn: jsii.String("functionArn"),
//   					},
//   				},
//   				LambdaFunctionAssociations: []interface{}{
//   					&LambdaFunctionAssociationProperty{
//   						EventType: jsii.String("eventType"),
//   						IncludeBody: jsii.Boolean(false),
//   						LambdaFunctionArn: jsii.String("lambdaFunctionArn"),
//   					},
//   				},
//   				MaxTtl: jsii.Number(123),
//   				MinTtl: jsii.Number(123),
//   				OriginRequestPolicyId: jsii.String("originRequestPolicyId"),
//   				RealtimeLogConfigArn: jsii.String("realtimeLogConfigArn"),
//   				ResponseHeadersPolicyId: jsii.String("responseHeadersPolicyId"),
//   				SmoothStreaming: jsii.Boolean(false),
//   				TrustedKeyGroups: []*string{
//   					jsii.String("trustedKeyGroups"),
//   				},
//   				TrustedSigners: []*string{
//   					jsii.String("trustedSigners"),
//   				},
//   			},
//   		},
//   		CnamEs: []*string{
//   			jsii.String("cnamEs"),
//   		},
//   		Comment: jsii.String("comment"),
//   		ContinuousDeploymentPolicyId: jsii.String("continuousDeploymentPolicyId"),
//   		CustomErrorResponses: []interface{}{
//   			&CustomErrorResponseProperty{
//   				ErrorCode: jsii.Number(123),
//
//   				// the properties below are optional
//   				ErrorCachingMinTtl: jsii.Number(123),
//   				ResponseCode: jsii.Number(123),
//   				ResponsePagePath: jsii.String("responsePagePath"),
//   			},
//   		},
//   		CustomOrigin: &LegacyCustomOriginProperty{
//   			DnsName: jsii.String("dnsName"),
//   			OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//   			OriginSslProtocols: []*string{
//   				jsii.String("originSslProtocols"),
//   			},
//
//   			// the properties below are optional
//   			HttpPort: jsii.Number(123),
//   			HttpsPort: jsii.Number(123),
//   		},
//   		DefaultRootObject: jsii.String("defaultRootObject"),
//   		HttpVersion: jsii.String("httpVersion"),
//   		Ipv6Enabled: jsii.Boolean(false),
//   		Logging: &LoggingProperty{
//   			Bucket: jsii.String("bucket"),
//
//   			// the properties below are optional
//   			IncludeCookies: jsii.Boolean(false),
//   			Prefix: jsii.String("prefix"),
//   		},
//   		OriginGroups: &OriginGroupsProperty{
//   			Quantity: jsii.Number(123),
//
//   			// the properties below are optional
//   			Items: []interface{}{
//   				&OriginGroupProperty{
//   					FailoverCriteria: &OriginGroupFailoverCriteriaProperty{
//   						StatusCodes: &StatusCodesProperty{
//   							Items: []interface{}{
//   								jsii.Number(123),
//   							},
//   							Quantity: jsii.Number(123),
//   						},
//   					},
//   					Id: jsii.String("id"),
//   					Members: &OriginGroupMembersProperty{
//   						Items: []interface{}{
//   							&OriginGroupMemberProperty{
//   								OriginId: jsii.String("originId"),
//   							},
//   						},
//   						Quantity: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		Origins: []interface{}{
//   			&OriginProperty{
//   				DomainName: jsii.String("domainName"),
//   				Id: jsii.String("id"),
//
//   				// the properties below are optional
//   				ConnectionAttempts: jsii.Number(123),
//   				ConnectionTimeout: jsii.Number(123),
//   				CustomOriginConfig: &CustomOriginConfigProperty{
//   					OriginProtocolPolicy: jsii.String("originProtocolPolicy"),
//
//   					// the properties below are optional
//   					HttpPort: jsii.Number(123),
//   					HttpsPort: jsii.Number(123),
//   					OriginKeepaliveTimeout: jsii.Number(123),
//   					OriginReadTimeout: jsii.Number(123),
//   					OriginSslProtocols: []*string{
//   						jsii.String("originSslProtocols"),
//   					},
//   				},
//   				OriginAccessControlId: jsii.String("originAccessControlId"),
//   				OriginCustomHeaders: []interface{}{
//   					&OriginCustomHeaderProperty{
//   						HeaderName: jsii.String("headerName"),
//   						HeaderValue: jsii.String("headerValue"),
//   					},
//   				},
//   				OriginPath: jsii.String("originPath"),
//   				OriginShield: &OriginShieldProperty{
//   					Enabled: jsii.Boolean(false),
//   					OriginShieldRegion: jsii.String("originShieldRegion"),
//   				},
//   				S3OriginConfig: &S3OriginConfigProperty{
//   					OriginAccessIdentity: jsii.String("originAccessIdentity"),
//   				},
//   			},
//   		},
//   		PriceClass: jsii.String("priceClass"),
//   		Restrictions: &RestrictionsProperty{
//   			GeoRestriction: &GeoRestrictionProperty{
//   				RestrictionType: jsii.String("restrictionType"),
//
//   				// the properties below are optional
//   				Locations: []*string{
//   					jsii.String("locations"),
//   				},
//   			},
//   		},
//   		S3Origin: &LegacyS3OriginProperty{
//   			DnsName: jsii.String("dnsName"),
//
//   			// the properties below are optional
//   			OriginAccessIdentity: jsii.String("originAccessIdentity"),
//   		},
//   		Staging: jsii.Boolean(false),
//   		ViewerCertificate: &ViewerCertificateProperty{
//   			AcmCertificateArn: jsii.String("acmCertificateArn"),
//   			CloudFrontDefaultCertificate: jsii.Boolean(false),
//   			IamCertificateId: jsii.String("iamCertificateId"),
//   			MinimumProtocolVersion: jsii.String("minimumProtocolVersion"),
//   			SslSupportMethod: jsii.String("sslSupportMethod"),
//   		},
//   		WebAclId: jsii.String("webAclId"),
//   	},
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDistributionProps struct {
	// The distribution's configuration.
	DistributionConfig interface{} `field:"required" json:"distributionConfig" yaml:"distributionConfig"`
	// A complex type that contains zero or more `Tag` elements.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}
