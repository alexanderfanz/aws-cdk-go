package awscognito


// The time units you use when you set the duration of ID, access, and refresh tokens.
//
// The default unit for RefreshToken is days, and the default for ID and access tokens is hours.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tokenValidityUnitsProperty := &tokenValidityUnitsProperty{
//   	accessToken: jsii.String("accessToken"),
//   	idToken: jsii.String("idToken"),
//   	refreshToken: jsii.String("refreshToken"),
//   }
//
type CfnUserPoolClient_TokenValidityUnitsProperty struct {
	// A time unit of `seconds` , `minutes` , `hours` , or `days` for the value that you set in the `AccessTokenValidity` parameter.
	//
	// The default `AccessTokenValidity` time unit is hours. `AccessTokenValidity` duration can range from five minutes to one day.
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// A time unit of `seconds` , `minutes` , `hours` , or `days` for the value that you set in the `IdTokenValidity` parameter.
	//
	// The default `IdTokenValidity` time unit is hours. `IdTokenValidity` duration can range from five minutes to one day.
	IdToken *string `field:"optional" json:"idToken" yaml:"idToken"`
	// A time unit of `seconds` , `minutes` , `hours` , or `days` for the value that you set in the `RefreshTokenValidity` parameter.
	//
	// The default `RefreshTokenValidity` time unit is days. `RefreshTokenValidity` duration can range from 60 minutes to 10 years.
	RefreshToken *string `field:"optional" json:"refreshToken" yaml:"refreshToken"`
}

