package awsamplify

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// GitHub source code provider.
//
// Example:
//   amplifyApp := amplify.NewApp(this, jsii.String("MyApp"), &appProps{
//   	sourceCodeProvider: amplify.NewGitHubSourceCodeProvider(&gitHubSourceCodeProviderProps{
//   		owner: jsii.String("<user>"),
//   		repository: jsii.String("<repo>"),
//   		oauthToken: awscdk.SecretValue.secretsManager(jsii.String("my-github-token")),
//   	}),
//   	autoBranchCreation: &autoBranchCreation{
//   		 // Automatically connect branches that match a pattern set
//   		patterns: []*string{
//   			jsii.String("feature/*"),
//   			jsii.String("test/*"),
//   		},
//   	},
//   	autoBranchDeletion: jsii.Boolean(true),
//   })
//
// Experimental.
type GitHubSourceCodeProvider interface {
	ISourceCodeProvider
	// Binds the source code provider to an app.
	// Experimental.
	Bind(_app App) *SourceCodeProviderConfig
}

// The jsii proxy struct for GitHubSourceCodeProvider
type jsiiProxy_GitHubSourceCodeProvider struct {
	jsiiProxy_ISourceCodeProvider
}

// Experimental.
func NewGitHubSourceCodeProvider(props *GitHubSourceCodeProviderProps) GitHubSourceCodeProvider {
	_init_.Initialize()

	if err := validateNewGitHubSourceCodeProviderParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GitHubSourceCodeProvider{}

	_jsii_.Create(
		"monocdk.aws_amplify.GitHubSourceCodeProvider",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewGitHubSourceCodeProvider_Override(g GitHubSourceCodeProvider, props *GitHubSourceCodeProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_amplify.GitHubSourceCodeProvider",
		[]interface{}{props},
		g,
	)
}

func (g *jsiiProxy_GitHubSourceCodeProvider) Bind(_app App) *SourceCodeProviderConfig {
	if err := g.validateBindParameters(_app); err != nil {
		panic(err)
	}
	var returns *SourceCodeProviderConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{_app},
		&returns,
	)

	return returns
}

