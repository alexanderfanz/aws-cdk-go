package awscognito


// OpenID Connect endpoints.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   oidcEndpoints := &oidcEndpoints{
//   	authorization: jsii.String("authorization"),
//   	jwksUri: jsii.String("jwksUri"),
//   	token: jsii.String("token"),
//   	userInfo: jsii.String("userInfo"),
//   }
//
// Experimental.
type OidcEndpoints struct {
	// Authorization endpoint.
	// Experimental.
	Authorization *string `field:"required" json:"authorization" yaml:"authorization"`
	// Jwks_uri endpoint.
	// Experimental.
	JwksUri *string `field:"required" json:"jwksUri" yaml:"jwksUri"`
	// Token endpoint.
	// Experimental.
	Token *string `field:"required" json:"token" yaml:"token"`
	// UserInfo endpoint.
	// Experimental.
	UserInfo *string `field:"required" json:"userInfo" yaml:"userInfo"`
}

