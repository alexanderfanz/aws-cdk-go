package awscdkcognitoidentitypoolalpha


// External Identity Providers To Connect to User Pools and Identity Pools.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cognito_identitypool_alpha "github.com/aws/aws-cdk-go/awscdkcognitoidentitypoolalpha"
//
//   identityPoolProviders := &IdentityPoolProviders{
//   	Amazon: &IdentityPoolAmazonLoginProvider{
//   		AppId: jsii.String("appId"),
//   	},
//   	Apple: &IdentityPoolAppleLoginProvider{
//   		ServicesId: jsii.String("servicesId"),
//   	},
//   	Digits: &IdentityPoolDigitsLoginProvider{
//   		ConsumerKey: jsii.String("consumerKey"),
//   		ConsumerSecret: jsii.String("consumerSecret"),
//   	},
//   	Facebook: &IdentityPoolFacebookLoginProvider{
//   		AppId: jsii.String("appId"),
//   	},
//   	Google: &IdentityPoolGoogleLoginProvider{
//   		ClientId: jsii.String("clientId"),
//   	},
//   	Twitter: &IdentityPoolTwitterLoginProvider{
//   		ConsumerKey: jsii.String("consumerKey"),
//   		ConsumerSecret: jsii.String("consumerSecret"),
//   	},
//   }
//
// Experimental.
type IdentityPoolProviders struct {
	// App Id for Amazon Identity Federation.
	// Default: -  No Amazon Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Experimental.
	Amazon *IdentityPoolAmazonLoginProvider `field:"optional" json:"amazon" yaml:"amazon"`
	// Services Id for Apple Identity Federation.
	// Default: - No Apple Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Experimental.
	Apple *IdentityPoolAppleLoginProvider `field:"optional" json:"apple" yaml:"apple"`
	// Consumer Key and Secret for Digits Identity Federation.
	// Default: - No Digits Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Experimental.
	Digits *IdentityPoolDigitsLoginProvider `field:"optional" json:"digits" yaml:"digits"`
	// App Id for Facebook Identity Federation.
	// Default: - No Facebook Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Experimental.
	Facebook *IdentityPoolFacebookLoginProvider `field:"optional" json:"facebook" yaml:"facebook"`
	// Client Id for Google Identity Federation.
	// Default: - No Google Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Experimental.
	Google *IdentityPoolGoogleLoginProvider `field:"optional" json:"google" yaml:"google"`
	// Consumer Key and Secret for Twitter Identity Federation.
	// Default: - No Twitter Authentication Provider used without OpenIdConnect or a User Pool.
	//
	// Experimental.
	Twitter *IdentityPoolTwitterLoginProvider `field:"optional" json:"twitter" yaml:"twitter"`
}

