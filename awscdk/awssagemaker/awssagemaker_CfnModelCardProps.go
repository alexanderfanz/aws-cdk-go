package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnModelCard`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var value interface{}
//
//   cfnModelCardProps := &CfnModelCardProps{
//   	Content: &ContentProperty{
//   		AdditionalInformation: &AdditionalInformationProperty{
//   			CaveatsAndRecommendations: jsii.String("caveatsAndRecommendations"),
//   			CustomDetails: map[string]*string{
//   				"customDetailsKey": jsii.String("customDetails"),
//   			},
//   			EthicalConsiderations: jsii.String("ethicalConsiderations"),
//   		},
//   		BusinessDetails: &BusinessDetailsProperty{
//   			BusinessProblem: jsii.String("businessProblem"),
//   			BusinessStakeholders: jsii.String("businessStakeholders"),
//   			LineOfBusiness: jsii.String("lineOfBusiness"),
//   		},
//   		EvaluationDetails: []interface{}{
//   			&EvaluationDetailProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				Datasets: []*string{
//   					jsii.String("datasets"),
//   				},
//   				EvaluationJobArn: jsii.String("evaluationJobArn"),
//   				EvaluationObservation: jsii.String("evaluationObservation"),
//   				Metadata: map[string]*string{
//   					"metadataKey": jsii.String("metadata"),
//   				},
//   				MetricGroups: []interface{}{
//   					&MetricGroupProperty{
//   						MetricData: []interface{}{
//   							&MetricDataItemsProperty{
//   								Name: jsii.String("name"),
//   								Type: jsii.String("type"),
//   								Value: value,
//
//   								// the properties below are optional
//   								Notes: jsii.String("notes"),
//   								XAxisName: []*string{
//   									jsii.String("xAxisName"),
//   								},
//   								YAxisName: []*string{
//   									jsii.String("yAxisName"),
//   								},
//   							},
//   						},
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   		IntendedUses: &IntendedUsesProperty{
//   			ExplanationsForRiskRating: jsii.String("explanationsForRiskRating"),
//   			FactorsAffectingModelEfficiency: jsii.String("factorsAffectingModelEfficiency"),
//   			IntendedUses: jsii.String("intendedUses"),
//   			PurposeOfModel: jsii.String("purposeOfModel"),
//   			RiskRating: jsii.String("riskRating"),
//   		},
//   		ModelOverview: &ModelOverviewProperty{
//   			AlgorithmType: jsii.String("algorithmType"),
//   			InferenceEnvironment: &InferenceEnvironmentProperty{
//   				ContainerImage: []*string{
//   					jsii.String("containerImage"),
//   				},
//   			},
//   			ModelArtifact: []*string{
//   				jsii.String("modelArtifact"),
//   			},
//   			ModelCreator: jsii.String("modelCreator"),
//   			ModelDescription: jsii.String("modelDescription"),
//   			ModelId: jsii.String("modelId"),
//   			ModelName: jsii.String("modelName"),
//   			ModelOwner: jsii.String("modelOwner"),
//   			ModelVersion: jsii.Number(123),
//   			ProblemType: jsii.String("problemType"),
//   		},
//   		TrainingDetails: &TrainingDetailsProperty{
//   			ObjectiveFunction: &ObjectiveFunctionProperty{
//   				Function: &FunctionProperty{
//   					Condition: jsii.String("condition"),
//   					Facet: jsii.String("facet"),
//   					Function: jsii.String("function"),
//   				},
//   				Notes: jsii.String("notes"),
//   			},
//   			TrainingJobDetails: &TrainingJobDetailsProperty{
//   				HyperParameters: []interface{}{
//   					&TrainingHyperParameterProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				TrainingArn: jsii.String("trainingArn"),
//   				TrainingDatasets: []*string{
//   					jsii.String("trainingDatasets"),
//   				},
//   				TrainingEnvironment: &TrainingEnvironmentProperty{
//   					ContainerImage: []*string{
//   						jsii.String("containerImage"),
//   					},
//   				},
//   				TrainingMetrics: []interface{}{
//   					&TrainingMetricProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.Number(123),
//
//   						// the properties below are optional
//   						Notes: jsii.String("notes"),
//   					},
//   				},
//   				UserProvidedHyperParameters: []interface{}{
//   					&TrainingHyperParameterProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				UserProvidedTrainingMetrics: []interface{}{
//   					&TrainingMetricProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.Number(123),
//
//   						// the properties below are optional
//   						Notes: jsii.String("notes"),
//   					},
//   				},
//   			},
//   			TrainingObservations: jsii.String("trainingObservations"),
//   		},
//   	},
//   	ModelCardName: jsii.String("modelCardName"),
//   	ModelCardStatus: jsii.String("modelCardStatus"),
//
//   	// the properties below are optional
//   	CreatedBy: &UserContextProperty{
//   		DomainId: jsii.String("domainId"),
//   		UserProfileArn: jsii.String("userProfileArn"),
//   		UserProfileName: jsii.String("userProfileName"),
//   	},
//   	LastModifiedBy: &UserContextProperty{
//   		DomainId: jsii.String("domainId"),
//   		UserProfileArn: jsii.String("userProfileArn"),
//   		UserProfileName: jsii.String("userProfileName"),
//   	},
//   	SecurityConfig: &SecurityConfigProperty{
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnModelCardProps struct {
	// The content of the model card.
	//
	// Content uses the [model card JSON schema](https://docs.aws.amazon.com/sagemaker/latest/dg/model-cards.html#model-cards-json-schema) .
	Content interface{} `field:"required" json:"content" yaml:"content"`
	// The unique name of the model card.
	ModelCardName *string `field:"required" json:"modelCardName" yaml:"modelCardName"`
	// The approval status of the model card within your organization.
	//
	// Different organizations might have different criteria for model card review and approval.
	//
	// - `Draft` : The model card is a work in progress.
	// - `PendingReview` : The model card is pending review.
	// - `Approved` : The model card is approved.
	// - `Archived` : The model card is archived. No more updates should be made to the model card, but it can still be exported.
	ModelCardStatus *string `field:"required" json:"modelCardStatus" yaml:"modelCardStatus"`
	// Information about the user who created or modified one or more of the following:.
	//
	// - Experiment
	// - Trial
	// - Trial component
	// - Lineage group
	// - Project
	// - Model Card.
	CreatedBy interface{} `field:"optional" json:"createdBy" yaml:"createdBy"`
	// `AWS::SageMaker::ModelCard.LastModifiedBy`.
	LastModifiedBy interface{} `field:"optional" json:"lastModifiedBy" yaml:"lastModifiedBy"`
	// The security configuration used to protect model card data.
	SecurityConfig interface{} `field:"optional" json:"securityConfig" yaml:"securityConfig"`
	// Key-value pairs used to manage metadata for the model card.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

