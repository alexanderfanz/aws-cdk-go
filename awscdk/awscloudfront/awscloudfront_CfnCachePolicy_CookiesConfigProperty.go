package awscloudfront


// An object that determines whether any cookies in viewer requests (and if so, which cookies) are included in the cache key and automatically included in requests that CloudFront sends to the origin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cookiesConfigProperty := &cookiesConfigProperty{
//   	cookieBehavior: jsii.String("cookieBehavior"),
//
//   	// the properties below are optional
//   	cookies: []*string{
//   		jsii.String("cookies"),
//   	},
//   }
//
type CfnCachePolicy_CookiesConfigProperty struct {
	// Determines whether any cookies in viewer requests are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	//
	// Valid values are:
	//
	// - `none` – Cookies in viewer requests are not included in the cache key and are not automatically included in requests that CloudFront sends to the origin. Even when this field is set to `none` , any cookies that are listed in an `OriginRequestPolicy` *are* included in origin requests.
	// - `whitelist` – The cookies in viewer requests that are listed in the `CookieNames` type are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	// - `allExcept` – All cookies in viewer requests that are **not** listed in the `CookieNames` type are included in the cache key and automatically included in requests that CloudFront sends to the origin.
	// - `all` – All cookies in viewer requests are included in the cache key and are automatically included in requests that CloudFront sends to the origin.
	CookieBehavior *string `field:"required" json:"cookieBehavior" yaml:"cookieBehavior"`
	// Contains a list of cookie names.
	Cookies *[]*string `field:"optional" json:"cookies" yaml:"cookies"`
}
