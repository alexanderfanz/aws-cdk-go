package awscodedeploy


// Lambda Deployment config type.
//
// Example:
//   var application lambdaApplication
//   var alias alias
//   config := codedeploy.NewCustomLambdaDeploymentConfig(this, jsii.String("CustomConfig"), &customLambdaDeploymentConfigProps{
//   	type: codedeploy.customLambdaDeploymentConfigType_CANARY,
//   	interval: awscdk.Duration.minutes(jsii.Number(1)),
//   	percentage: jsii.Number(5),
//   })
//   deploymentGroup := codedeploy.NewLambdaDeploymentGroup(this, jsii.String("BlueGreenDeployment"), &lambdaDeploymentGroupProps{
//   	application: application,
//   	alias: alias,
//   	deploymentConfig: config,
//   })
//
type CustomLambdaDeploymentConfigType string

const (
	// Canary deployment type.
	CustomLambdaDeploymentConfigType_CANARY CustomLambdaDeploymentConfigType = "CANARY"
	// Linear deployment type.
	CustomLambdaDeploymentConfigType_LINEAR CustomLambdaDeploymentConfigType = "LINEAR"
)
