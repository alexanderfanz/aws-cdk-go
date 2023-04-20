package awsrds


// Properties for SQL Server Express Edition instance engines.
//
// Used in {@link DatabaseInstanceEngine.sqlServerEx}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sqlServerEngineVersion sqlServerEngineVersion
//
//   sqlServerExInstanceEngineProps := &SqlServerExInstanceEngineProps{
//   	Version: sqlServerEngineVersion,
//   }
//
// Experimental.
type SqlServerExInstanceEngineProps struct {
	// The exact version of the engine to use.
	// Experimental.
	Version SqlServerEngineVersion `field:"required" json:"version" yaml:"version"`
}

