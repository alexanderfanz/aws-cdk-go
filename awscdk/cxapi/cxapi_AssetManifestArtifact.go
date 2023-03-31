package cxapi

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/cloudassemblyschema"
)

// Asset manifest is a description of a set of assets which need to be built and published.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var cloudAssembly cloudAssembly
//
//   assetManifestArtifact := awscdk.Cx_api.NewAssetManifestArtifact(cloudAssembly, jsii.String("name"), &artifactManifest{
//   	type: awscdk.Cloud_assembly_schema.artifactType_NONE,
//
//   	// the properties below are optional
//   	dependencies: []*string{
//   		jsii.String("dependencies"),
//   	},
//   	displayName: jsii.String("displayName"),
//   	environment: jsii.String("environment"),
//   	metadata: map[string][]metadataEntry{
//   		"metadataKey": []*metadataEntry{
//   			&metadataEntry{
//   				"type": jsii.String("type"),
//
//   				// the properties below are optional
//   				"data": jsii.String("data"),
//   				"trace": []*string{
//   					jsii.String("trace"),
//   				},
//   			},
//   		},
//   	},
//   	properties: &awsCloudFormationStackProperties{
//   		templateFile: jsii.String("templateFile"),
//
//   		// the properties below are optional
//   		assumeRoleArn: jsii.String("assumeRoleArn"),
//   		assumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   		bootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   		cloudFormationExecutionRoleArn: jsii.String("cloudFormationExecutionRoleArn"),
//   		lookupRole: &bootstrapRole{
//   			arn: jsii.String("arn"),
//
//   			// the properties below are optional
//   			assumeRoleExternalId: jsii.String("assumeRoleExternalId"),
//   			bootstrapStackVersionSsmParameter: jsii.String("bootstrapStackVersionSsmParameter"),
//   			requiresBootstrapStackVersion: jsii.Number(123),
//   		},
//   		parameters: map[string]*string{
//   			"parametersKey": jsii.String("parameters"),
//   		},
//   		requiresBootstrapStackVersion: jsii.Number(123),
//   		stackName: jsii.String("stackName"),
//   		stackTemplateAssetObjectUrl: jsii.String("stackTemplateAssetObjectUrl"),
//   		tags: map[string]*string{
//   			"tagsKey": jsii.String("tags"),
//   		},
//   		terminationProtection: jsii.Boolean(false),
//   		validateOnSynth: jsii.Boolean(false),
//   	},
//   })
//
// Experimental.
type AssetManifestArtifact interface {
	CloudArtifact
	// Experimental.
	Assembly() CloudAssembly
	// Name of SSM parameter with bootstrap stack version.
	// Experimental.
	BootstrapStackVersionSsmParameter() *string
	// Returns all the artifacts that this artifact depends on.
	// Experimental.
	Dependencies() *[]CloudArtifact
	// The file name of the asset manifest.
	// Experimental.
	File() *string
	// An identifier that shows where this artifact is located in the tree of nested assemblies, based on their manifests.
	//
	// Defaults to the normal
	// id. Should only be used in user interfaces.
	// Experimental.
	HierarchicalId() *string
	// Experimental.
	Id() *string
	// The artifact's manifest.
	// Experimental.
	Manifest() *cloudassemblyschema.ArtifactManifest
	// The set of messages extracted from the artifact's metadata.
	// Experimental.
	Messages() *[]*SynthesisMessage
	// Version of bootstrap stack required to deploy this stack.
	// Experimental.
	RequiresBootstrapStackVersion() *float64
	// Returns: all the metadata entries of a specific type in this artifact.
	// Experimental.
	FindMetadataByType(type_ *string) *[]*MetadataEntryResult
}

// The jsii proxy struct for AssetManifestArtifact
type jsiiProxy_AssetManifestArtifact struct {
	jsiiProxy_CloudArtifact
}

func (j *jsiiProxy_AssetManifestArtifact) Assembly() CloudAssembly {
	var returns CloudAssembly
	_jsii_.Get(
		j,
		"assembly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetManifestArtifact) BootstrapStackVersionSsmParameter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bootstrapStackVersionSsmParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetManifestArtifact) Dependencies() *[]CloudArtifact {
	var returns *[]CloudArtifact
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetManifestArtifact) File() *string {
	var returns *string
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetManifestArtifact) HierarchicalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hierarchicalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetManifestArtifact) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetManifestArtifact) Manifest() *cloudassemblyschema.ArtifactManifest {
	var returns *cloudassemblyschema.ArtifactManifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetManifestArtifact) Messages() *[]*SynthesisMessage {
	var returns *[]*SynthesisMessage
	_jsii_.Get(
		j,
		"messages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssetManifestArtifact) RequiresBootstrapStackVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiresBootstrapStackVersion",
		&returns,
	)
	return returns
}


// Experimental.
func NewAssetManifestArtifact(assembly CloudAssembly, name *string, artifact *cloudassemblyschema.ArtifactManifest) AssetManifestArtifact {
	_init_.Initialize()

	if err := validateNewAssetManifestArtifactParameters(assembly, name, artifact); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssetManifestArtifact{}

	_jsii_.Create(
		"monocdk.cx_api.AssetManifestArtifact",
		[]interface{}{assembly, name, artifact},
		&j,
	)

	return &j
}

// Experimental.
func NewAssetManifestArtifact_Override(a AssetManifestArtifact, assembly CloudAssembly, name *string, artifact *cloudassemblyschema.ArtifactManifest) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.cx_api.AssetManifestArtifact",
		[]interface{}{assembly, name, artifact},
		a,
	)
}

// Returns a subclass of `CloudArtifact` based on the artifact type defined in the artifact manifest.
//
// Returns: the `CloudArtifact` that matches the artifact type or `undefined` if it's an artifact type that is unrecognized by this module.
// Experimental.
func AssetManifestArtifact_FromManifest(assembly CloudAssembly, id *string, artifact *cloudassemblyschema.ArtifactManifest) CloudArtifact {
	_init_.Initialize()

	if err := validateAssetManifestArtifact_FromManifestParameters(assembly, id, artifact); err != nil {
		panic(err)
	}
	var returns CloudArtifact

	_jsii_.StaticInvoke(
		"monocdk.cx_api.AssetManifestArtifact",
		"fromManifest",
		[]interface{}{assembly, id, artifact},
		&returns,
	)

	return returns
}

// Checks if `art` is an instance of this class.
//
// Use this method instead of `instanceof` to properly detect `AssetManifestArtifact`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `cx-api` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `AssetManifestArtifact` in each copy of the `cx-api` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `cx-api`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
// Experimental.
func AssetManifestArtifact_IsAssetManifestArtifact(art interface{}) *bool {
	_init_.Initialize()

	if err := validateAssetManifestArtifact_IsAssetManifestArtifactParameters(art); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.cx_api.AssetManifestArtifact",
		"isAssetManifestArtifact",
		[]interface{}{art},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssetManifestArtifact) FindMetadataByType(type_ *string) *[]*MetadataEntryResult {
	if err := a.validateFindMetadataByTypeParameters(type_); err != nil {
		panic(err)
	}
	var returns *[]*MetadataEntryResult

	_jsii_.Invoke(
		a,
		"findMetadataByType",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

