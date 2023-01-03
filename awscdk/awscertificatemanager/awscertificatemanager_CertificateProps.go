package awscertificatemanager


// Properties for your certificate.
//
// Example:
//   exampleCom := route53.NewHostedZone(this, jsii.String("ExampleCom"), &hostedZoneProps{
//   	zoneName: jsii.String("example.com"),
//   })
//   exampleNet := route53.NewHostedZone(this, jsii.String("ExampleNet"), &hostedZoneProps{
//   	zoneName: jsii.String("example.net"),
//   })
//
//   cert := acm.NewCertificate(this, jsii.String("Certificate"), &certificateProps{
//   	domainName: jsii.String("test.example.com"),
//   	subjectAlternativeNames: []*string{
//   		jsii.String("cool.example.com"),
//   		jsii.String("test.example.net"),
//   	},
//   	validation: acm.certificateValidation.fromDnsMultiZone(map[string]iHostedZone{
//   		"test.example.com": exampleCom,
//   		"cool.example.com": exampleCom,
//   		"test.example.net": exampleNet,
//   	}),
//   })
//
type CertificateProps struct {
	// Fully-qualified domain name to request a certificate for.
	//
	// May contain wildcards, such as ``*.domain.com``.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The Certifcate name.
	//
	// Since the Certifcate resource doesn't support providing a physical name, the value provided here will be recorded in the `Name` tag.
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// Alternative domain names on your certificate.
	//
	// Use this to register alternative domain names that represent the same site.
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// Enable or disable transparency logging for this certificate.
	//
	// Once a certificate has been logged, it cannot be removed from the log.
	// Opting out at that point will have no effect. If you opt out of logging
	// when you request a certificate and then choose later to opt back in,
	// your certificate will not be logged until it is renewed.
	// If you want the certificate to be logged immediately, we recommend that you issue a new one.
	// See: https://docs.aws.amazon.com/acm/latest/userguide/acm-bestpractices.html#best-practices-transparency
	//
	TransparencyLoggingEnabled *bool `field:"optional" json:"transparencyLoggingEnabled" yaml:"transparencyLoggingEnabled"`
	// How to validate this certificate.
	Validation CertificateValidation `field:"optional" json:"validation" yaml:"validation"`
}

