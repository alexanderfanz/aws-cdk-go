package awsroute53

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Construction properties for a TxtRecord.
//
// Example:
//   var myZone hostedZone
//
//
//   route53.NewTxtRecord(this, jsii.String("TXTRecord"), &TxtRecordProps{
//   	Zone: myZone,
//   	RecordName: jsii.String("_foo"),
//   	 // If the name ends with a ".", it will be used as-is;
//   	// if it ends with a "." followed by the zone name, a trailing "." will be added automatically;
//   	// otherwise, a ".", the zone name, and a trailing "." will be added automatically.
//   	// Defaults to zone root if not specified.
//   	Values: []*string{
//   		jsii.String("Bar!"),
//   		jsii.String("Baz?"),
//   	},
//   	Ttl: awscdk.Duration_Minutes(jsii.Number(90)),
//   })
//
type TxtRecordProps struct {
	// The hosted zone in which to define the new record.
	Zone IHostedZone `field:"required" json:"zone" yaml:"zone"`
	// A comment to add on the record.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Whether to delete the same record set in the hosted zone if it already exists (dangerous!).
	//
	// This allows to deploy a new record set while minimizing the downtime because the
	// new record set will be created immediately after the existing one is deleted. It
	// also avoids "manual" actions to delete existing record sets.
	//
	// > **N.B.:** this feature is dangerous, use with caution! It can only be used safely when
	// > `deleteExisting` is set to `true` as soon as the resource is added to the stack. Changing
	// > an existing Record Set's `deleteExisting` property from `false -> true` after deployment
	// > will delete the record!
	DeleteExisting *bool `field:"optional" json:"deleteExisting" yaml:"deleteExisting"`
	// The domain name for this record.
	RecordName *string `field:"optional" json:"recordName" yaml:"recordName"`
	// The resource record cache time to live (TTL).
	Ttl awscdk.Duration `field:"optional" json:"ttl" yaml:"ttl"`
	// The text values.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

