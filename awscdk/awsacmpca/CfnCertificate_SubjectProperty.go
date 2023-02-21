package awsacmpca


// Contains information about the certificate subject.
//
// The `Subject` field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The `Subject` must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subjectProperty := &SubjectProperty{
//   	CommonName: jsii.String("commonName"),
//   	Country: jsii.String("country"),
//   	CustomAttributes: []interface{}{
//   		&CustomAttributeProperty{
//   			ObjectIdentifier: jsii.String("objectIdentifier"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	DistinguishedNameQualifier: jsii.String("distinguishedNameQualifier"),
//   	GenerationQualifier: jsii.String("generationQualifier"),
//   	GivenName: jsii.String("givenName"),
//   	Initials: jsii.String("initials"),
//   	Locality: jsii.String("locality"),
//   	Organization: jsii.String("organization"),
//   	OrganizationalUnit: jsii.String("organizationalUnit"),
//   	Pseudonym: jsii.String("pseudonym"),
//   	SerialNumber: jsii.String("serialNumber"),
//   	State: jsii.String("state"),
//   	Surname: jsii.String("surname"),
//   	Title: jsii.String("title"),
//   }
//
type CfnCertificate_SubjectProperty struct {
	// For CA and end-entity certificates in a private PKI, the common name (CN) can be any string within the length limit.
	//
	// Note: In publicly trusted certificates, the common name must be a fully qualified domain name (FQDN) associated with the certificate subject.
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// Two-digit code that specifies the country in which the certificate subject located.
	Country *string `field:"optional" json:"country" yaml:"country"`
	// Contains a sequence of one or more X.500 relative distinguished names (RDNs), each of which consists of an object identifier (OID) and a value. For more information, see NIST’s definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier) .
	//
	// > Custom attributes cannot be used in combination with standard attributes.
	CustomAttributes interface{} `field:"optional" json:"customAttributes" yaml:"customAttributes"`
	// Disambiguating information for the certificate subject.
	DistinguishedNameQualifier *string `field:"optional" json:"distinguishedNameQualifier" yaml:"distinguishedNameQualifier"`
	// Typically a qualifier appended to the name of an individual.
	//
	// Examples include Jr. for junior, Sr. for senior, and III for third.
	GenerationQualifier *string `field:"optional" json:"generationQualifier" yaml:"generationQualifier"`
	// First name.
	GivenName *string `field:"optional" json:"givenName" yaml:"givenName"`
	// Concatenation that typically contains the first letter of the *GivenName* , the first letter of the middle name if one exists, and the first letter of the *Surname* .
	Initials *string `field:"optional" json:"initials" yaml:"initials"`
	// The locality (such as a city or town) in which the certificate subject is located.
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// Legal name of the organization with which the certificate subject is affiliated.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// A subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated.
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// Typically a shortened version of a longer *GivenName* .
	//
	// For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.
	Pseudonym *string `field:"optional" json:"pseudonym" yaml:"pseudonym"`
	// The certificate serial number.
	SerialNumber *string `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// State in which the subject of the certificate is located.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Family name.
	//
	// In the US and the UK, for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first.
	Surname *string `field:"optional" json:"surname" yaml:"surname"`
	// A title such as Mr.
	//
	// or Ms., which is pre-pended to the name to refer formally to the certificate subject.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

