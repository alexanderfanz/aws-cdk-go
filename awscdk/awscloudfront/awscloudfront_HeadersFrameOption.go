package awscloudfront


// Enum representing possible values of the X-Frame-Options HTTP response header.
//
// Example:
//   // Using an existing managed response headers policy
//   var bucketOrigin s3Origin
//
//   cloudfront.NewDistribution(this, jsii.String("myDistManagedPolicy"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: bucketOrigin,
//   		responseHeadersPolicy: cloudfront.responseHeadersPolicy_CORS_ALLOW_ALL_ORIGINS(),
//   	},
//   })
//
//   // Creating a custom response headers policy -- all parameters optional
//   myResponseHeadersPolicy := cloudfront.NewResponseHeadersPolicy(this, jsii.String("ResponseHeadersPolicy"), &responseHeadersPolicyProps{
//   	responseHeadersPolicyName: jsii.String("MyPolicy"),
//   	comment: jsii.String("A default policy"),
//   	corsBehavior: &responseHeadersCorsBehavior{
//   		accessControlAllowCredentials: jsii.Boolean(false),
//   		accessControlAllowHeaders: []*string{
//   			jsii.String("X-Custom-Header-1"),
//   			jsii.String("X-Custom-Header-2"),
//   		},
//   		accessControlAllowMethods: []*string{
//   			jsii.String("GET"),
//   			jsii.String("POST"),
//   		},
//   		accessControlAllowOrigins: []*string{
//   			jsii.String("*"),
//   		},
//   		accessControlExposeHeaders: []*string{
//   			jsii.String("X-Custom-Header-1"),
//   			jsii.String("X-Custom-Header-2"),
//   		},
//   		accessControlMaxAge: awscdk.Duration.seconds(jsii.Number(600)),
//   		originOverride: jsii.Boolean(true),
//   	},
//   	customHeadersBehavior: &responseCustomHeadersBehavior{
//   		customHeaders: []responseCustomHeader{
//   			&responseCustomHeader{
//   				header: jsii.String("X-Amz-Date"),
//   				value: jsii.String("some-value"),
//   				override: jsii.Boolean(true),
//   			},
//   			&responseCustomHeader{
//   				header: jsii.String("X-Amz-Security-Token"),
//   				value: jsii.String("some-value"),
//   				override: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   	securityHeadersBehavior: &responseSecurityHeadersBehavior{
//   		contentSecurityPolicy: &responseHeadersContentSecurityPolicy{
//   			contentSecurityPolicy: jsii.String("default-src https:;"),
//   			override: jsii.Boolean(true),
//   		},
//   		contentTypeOptions: &responseHeadersContentTypeOptions{
//   			override: jsii.Boolean(true),
//   		},
//   		frameOptions: &responseHeadersFrameOptions{
//   			frameOption: cloudfront.headersFrameOption_DENY,
//   			override: jsii.Boolean(true),
//   		},
//   		referrerPolicy: &responseHeadersReferrerPolicy{
//   			referrerPolicy: cloudfront.headersReferrerPolicy_NO_REFERRER,
//   			override: jsii.Boolean(true),
//   		},
//   		strictTransportSecurity: &responseHeadersStrictTransportSecurity{
//   			accessControlMaxAge: awscdk.Duration.seconds(jsii.Number(600)),
//   			includeSubdomains: jsii.Boolean(true),
//   			override: jsii.Boolean(true),
//   		},
//   		xssProtection: &responseHeadersXSSProtection{
//   			protection: jsii.Boolean(true),
//   			modeBlock: jsii.Boolean(true),
//   			reportUri: jsii.String("https://example.com/csp-report"),
//   			override: jsii.Boolean(true),
//   		},
//   	},
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDistCustomPolicy"), &distributionProps{
//   	defaultBehavior: &behaviorOptions{
//   		origin: bucketOrigin,
//   		responseHeadersPolicy: myResponseHeadersPolicy,
//   	},
//   })
//
// Experimental.
type HeadersFrameOption string

const (
	// The page can only be displayed in a frame on the same origin as the page itself.
	// Experimental.
	HeadersFrameOption_DENY HeadersFrameOption = "DENY"
	// The page can only be displayed in a frame on the specified origin.
	// Experimental.
	HeadersFrameOption_SAMEORIGIN HeadersFrameOption = "SAMEORIGIN"
)

