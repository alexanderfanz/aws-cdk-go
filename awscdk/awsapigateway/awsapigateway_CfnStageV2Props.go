package awsapigateway


// Properties for defining a `AWS::ApiGatewayV2::Stage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var routeSettings interface{}
//   var stageVariables interface{}
//   var tags interface{}
//
//   cfnStageV2Props := &CfnStageV2Props{
//   	ApiId: jsii.String("apiId"),
//   	StageName: jsii.String("stageName"),
//
//   	// the properties below are optional
//   	AccessLogSettings: &AccessLogSettingsProperty{
//   		DestinationArn: jsii.String("destinationArn"),
//   		Format: jsii.String("format"),
//   	},
//   	AutoDeploy: jsii.Boolean(false),
//   	ClientCertificateId: jsii.String("clientCertificateId"),
//   	DefaultRouteSettings: &RouteSettingsProperty{
//   		DataTraceEnabled: jsii.Boolean(false),
//   		DetailedMetricsEnabled: jsii.Boolean(false),
//   		LoggingLevel: jsii.String("loggingLevel"),
//   		ThrottlingBurstLimit: jsii.Number(123),
//   		ThrottlingRateLimit: jsii.Number(123),
//   	},
//   	DeploymentId: jsii.String("deploymentId"),
//   	Description: jsii.String("description"),
//   	RouteSettings: routeSettings,
//   	StageVariables: stageVariables,
//   	Tags: tags,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html
//
// Deprecated: moved to package aws-apigatewayv2.
type CfnStageV2Props struct {
	// `AWS::ApiGatewayV2::Stage.ApiId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-apiid
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// `AWS::ApiGatewayV2::Stage.StageName`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-stagename
	//
	// Deprecated: moved to package aws-apigatewayv2.
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// `AWS::ApiGatewayV2::Stage.AccessLogSettings`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-accesslogsettings
	//
	// Deprecated: moved to package aws-apigatewayv2.
	AccessLogSettings interface{} `field:"optional" json:"accessLogSettings" yaml:"accessLogSettings"`
	// `AWS::ApiGatewayV2::Stage.AutoDeploy`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-autodeploy
	//
	// Deprecated: moved to package aws-apigatewayv2.
	AutoDeploy interface{} `field:"optional" json:"autoDeploy" yaml:"autoDeploy"`
	// `AWS::ApiGatewayV2::Stage.ClientCertificateId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-clientcertificateid
	//
	// Deprecated: moved to package aws-apigatewayv2.
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// `AWS::ApiGatewayV2::Stage.DefaultRouteSettings`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-defaultroutesettings
	//
	// Deprecated: moved to package aws-apigatewayv2.
	DefaultRouteSettings interface{} `field:"optional" json:"defaultRouteSettings" yaml:"defaultRouteSettings"`
	// `AWS::ApiGatewayV2::Stage.DeploymentId`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-deploymentid
	//
	// Deprecated: moved to package aws-apigatewayv2.
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// `AWS::ApiGatewayV2::Stage.Description`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-description
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::ApiGatewayV2::Stage.RouteSettings`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-routesettings
	//
	// Deprecated: moved to package aws-apigatewayv2.
	RouteSettings interface{} `field:"optional" json:"routeSettings" yaml:"routeSettings"`
	// `AWS::ApiGatewayV2::Stage.StageVariables`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-stagevariables
	//
	// Deprecated: moved to package aws-apigatewayv2.
	StageVariables interface{} `field:"optional" json:"stageVariables" yaml:"stageVariables"`
	// `AWS::ApiGatewayV2::Stage.Tags`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-stage.html#cfn-apigatewayv2-stage-tags
	//
	// Deprecated: moved to package aws-apigatewayv2.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

