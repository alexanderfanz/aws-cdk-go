package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The versions for the MySQL instance engines (those returned by {@link DatabaseInstanceEngine.mysql}).
//
// Example:
//   var vpc vpc
//
//   role := iam.NewRole(this, jsii.String("RDSDirectoryServicesRole"), &roleProps{
//   	assumedBy: iam.NewServicePrincipal(jsii.String("rds.amazonaws.com")),
//   	managedPolicies: []iManagedPolicy{
//   		iam.managedPolicy.fromAwsManagedPolicyName(jsii.String("service-role/AmazonRDSDirectoryServiceAccess")),
//   	},
//   })
//   instance := rds.NewDatabaseInstance(this, jsii.String("Instance"), &databaseInstanceProps{
//   	engine: rds.databaseInstanceEngine.mysql(&mySqlInstanceEngineProps{
//   		version: rds.mysqlEngineVersion_VER_8_0_19(),
//   	}),
//   	vpc: vpc,
//   	domain: jsii.String("d-????????"),
//   	 // The ID of the domain for the instance to join.
//   	domainRole: role,
//   })
//
// Experimental.
type MysqlEngineVersion interface {
	// The full version string, for example, "10.5.28".
	// Experimental.
	MysqlFullVersion() *string
	// The major version of the engine, for example, "10.5".
	// Experimental.
	MysqlMajorVersion() *string
}

// The jsii proxy struct for MysqlEngineVersion
type jsiiProxy_MysqlEngineVersion struct {
	_ byte // padding
}

func (j *jsiiProxy_MysqlEngineVersion) MysqlFullVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlFullVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MysqlEngineVersion) MysqlMajorVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mysqlMajorVersion",
		&returns,
	)
	return returns
}


// Create a new MysqlEngineVersion with an arbitrary version.
// Experimental.
func MysqlEngineVersion_Of(mysqlFullVersion *string, mysqlMajorVersion *string) MysqlEngineVersion {
	_init_.Initialize()

	if err := validateMysqlEngineVersion_OfParameters(mysqlFullVersion, mysqlMajorVersion); err != nil {
		panic(err)
	}
	var returns MysqlEngineVersion

	_jsii_.StaticInvoke(
		"monocdk.aws_rds.MysqlEngineVersion",
		"of",
		[]interface{}{mysqlFullVersion, mysqlMajorVersion},
		&returns,
	)

	return returns
}

func MysqlEngineVersion_VER_5_5() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_5",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_5_46() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_5_46",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_5_53() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_5_53",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_5_57() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_5_57",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_5_59() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_5_59",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_5_61() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_5_61",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_6",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_34() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_6_34",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_35() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_6_35",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_37() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_6_37",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_39() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_6_39",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_40() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_6_40",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_41() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_6_41",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_43() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_6_43",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_44() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_6_44",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_46() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_6_46",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_48() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_6_48",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_49() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_6_49",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_6_51() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_6_51",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_16() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_16",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_17() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_17",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_19() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_19",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_21() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_21",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_22() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_22",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_23() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_23",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_24() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_24",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_25() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_25",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_26() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_26",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_28() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_28",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_30() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_30",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_31() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_31",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_33() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_33",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_34() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_34",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_35() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_35",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_36() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_36",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_5_7_37() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_5_7_37",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_8_0",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_11() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_8_0_11",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_13() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_8_0_13",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_15() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_8_0_15",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_16() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_8_0_16",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_17() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_8_0_17",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_19() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_8_0_19",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_20() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_8_0_20",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_21() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_8_0_21",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_23() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_8_0_23",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_25() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_8_0_25",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_26() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_8_0_26",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_27() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_8_0_27",
		&returns,
	)
	return returns
}

func MysqlEngineVersion_VER_8_0_28() MysqlEngineVersion {
	_init_.Initialize()
	var returns MysqlEngineVersion
	_jsii_.StaticGet(
		"monocdk.aws_rds.MysqlEngineVersion",
		"VER_8_0_28",
		&returns,
	)
	return returns
}

