# Amazon Lambda Python Library

This library provides constructs for Python Lambda functions.

To use this module, you will need to have Docker installed.

## Python Function

Define a `PythonFunction`:

```go
lambda.NewPythonFunction(this, jsii.String("MyFunction"), &pythonFunctionProps{
	entry: jsii.String("/path/to/my/function"),
	 // required
	runtime: awscdk.Runtime_PYTHON_3_8(),
	 // required
	index: jsii.String("my_index.py"),
	 // optional, defaults to 'index.py'
	handler: jsii.String("my_exported_func"),
})
```

All other properties of `lambda.Function` are supported, see also the [AWS Lambda construct library](https://github.com/aws/aws-cdk/tree/master/packages/%40aws-cdk/aws-lambda).

## Python Layer

You may create a python-based lambda layer with `PythonLayerVersion`. If `PythonLayerVersion` detects a `requirements.txt`
or `Pipfile` or `poetry.lock` with the associated `pyproject.toml` at the entry path, then `PythonLayerVersion` will include the dependencies inline with your code in the
layer.

Define a `PythonLayerVersion`:

```go
lambda.NewPythonLayerVersion(this, jsii.String("MyLayer"), &pythonLayerVersionProps{
	entry: jsii.String("/path/to/my/layer"),
})
```

A layer can also be used as a part of a `PythonFunction`:

```go
lambda.NewPythonFunction(this, jsii.String("MyFunction"), &pythonFunctionProps{
	entry: jsii.String("/path/to/my/function"),
	runtime: awscdk.Runtime_PYTHON_3_8(),
	layers: []iLayerVersion{
		lambda.NewPythonLayerVersion(this, jsii.String("MyLayer"), &pythonLayerVersionProps{
			entry: jsii.String("/path/to/my/layer"),
		}),
	},
})
```

## Packaging

If `requirements.txt`, `Pipfile` or `poetry.lock` exists at the entry path, the construct will handle installing all required modules in a [Lambda compatible Docker container](https://gallery.ecr.aws/sam/build-python3.7) according to the `runtime` and with the Docker platform based on the target architecture of the Lambda function.

Python bundles are only recreated and published when a file in a source directory has changed.
Therefore (and as a general best-practice), it is highly recommended to commit a lockfile with a
list of all transitive dependencies and their exact versions. This will ensure that when any dependency version is updated, the bundle asset is recreated and uploaded.

To that end, we recommend using [`pipenv`] or [`poetry`] which have lockfile support.

* [`pipenv`](https://pipenv-fork.readthedocs.io/en/latest/basics.html#example-pipfile-lock)
* [`poetry`](https://python-poetry.org/docs/basic-usage/#commit-your-poetrylock-file-to-version-control)

Packaging is executed using the `Packaging` class, which:

1. Infers the packaging type based on the files present.
2. If it sees a `Pipfile` or a `poetry.lock` file, it exports it to a compatible `requirements.txt` file with credentials (if they're available in the source files or in the bundling container).
3. Installs dependencies using `pip`.
4. Copies the dependencies into an asset that is bundled for the Lambda package.

**Lambda with a requirements.txt**

```plaintext
.
├── lambda_function.py # exports a function named 'handler'
├── requirements.txt # has to be present at the entry path
```

**Lambda with a Pipfile**

```plaintext
.
├── lambda_function.py # exports a function named 'handler'
├── Pipfile # has to be present at the entry path
├── Pipfile.lock # your lock file
```

**Lambda with a poetry.lock**

```plaintext
.
├── lambda_function.py # exports a function named 'handler'
├── pyproject.toml # your poetry project definition
├── poetry.lock # your poetry lock file has to be present at the entry path
```

## Custom Bundling

Custom bundling can be performed by passing in additional build arguments that point to index URLs to private repos, or by using an entirely custom Docker images for bundling dependencies. The build args currently supported are:

* `PIP_INDEX_URL`
* `PIP_EXTRA_INDEX_URL`
* `HTTPS_PROXY`

Additional build args for bundling that refer to PyPI indexes can be specified as:

```go
entry := "/path/to/function"
image := awscdk.DockerImage.fromBuild(entry)

lambda.NewPythonFunction(this, jsii.String("function"), &pythonFunctionProps{
	entry: jsii.String(entry),
	runtime: awscdk.Runtime_PYTHON_3_8(),
	bundling: &bundlingOptions{
		buildArgs: map[string]*string{
			"PIP_INDEX_URL": jsii.String("https://your.index.url/simple/"),
			"PIP_EXTRA_INDEX_URL": jsii.String("https://your.extra-index.url/simple/"),
		},
	},
})
```

If using a custom Docker image for bundling, the dependencies are installed with `pip`, `pipenv` or `poetry` by using the `Packaging` class. A different bundling Docker image that is in the same directory as the function can be specified as:

```go
entry := "/path/to/function"
image := awscdk.DockerImage.fromBuild(entry)

lambda.NewPythonFunction(this, jsii.String("function"), &pythonFunctionProps{
	entry: jsii.String(entry),
	runtime: awscdk.Runtime_PYTHON_3_8(),
	bundling: &bundlingOptions{
		image: image,
	},
})
```

## Custom Bundling with Code Artifact

To use a Code Artifact PyPI repo, the `PIP_INDEX_URL` for bundling the function can be customized (requires AWS CLI in the build environment):

```go
import "github.com/aws-samples/dummy/child_process"


entry := "/path/to/function"
image := awscdk.DockerImage.fromBuild(entry)

domain := "my-domain"
domainOwner := "111122223333"
repoName := "my_repo"
region := "us-east-1"
codeArtifactAuthToken := child_process.ExecSync(fmt.Sprintf("aws codeartifact get-authorization-token --domain %v --domain-owner %v --query authorizationToken --output text", domain, domainOwner)).toString().trim()

indexUrl := fmt.Sprintf("https://aws:%v@%v-%v.d.codeartifact.%v.amazonaws.com/pypi/%v/simple/", codeArtifactAuthToken, domain, domainOwner, region, repoName)

lambda.NewPythonFunction(this, jsii.String("function"), &pythonFunctionProps{
	entry: jsii.String(entry),
	runtime: awscdk.Runtime_PYTHON_3_8(),
	bundling: &bundlingOptions{
		environment: map[string]*string{
			"PIP_INDEX_URL": indexUrl,
		},
	},
})
```

The index URL or the token are only used during bundling and thus not included in the final asset. Setting only environment variable for `PIP_INDEX_URL` or `PIP_EXTRA_INDEX_URL` should work for accesing private Python repositories with `pip`, `pipenv` and `poetry` based dependencies.

If you also want to use the Code Artifact repo for building the base Docker image for bundling, use `buildArgs`. However, note that setting custom build args for bundling will force the base bundling image to be rebuilt every time (i.e. skip the Docker cache). Build args can be customized as:

```go
import "github.com/aws-samples/dummy/child_process"


entry := "/path/to/function"
image := awscdk.DockerImage.fromBuild(entry)

domain := "my-domain"
domainOwner := "111122223333"
repoName := "my_repo"
region := "us-east-1"
codeArtifactAuthToken := child_process.ExecSync(fmt.Sprintf("aws codeartifact get-authorization-token --domain %v --domain-owner %v --query authorizationToken --output text", domain, domainOwner)).toString().trim()

indexUrl := fmt.Sprintf("https://aws:%v@%v-%v.d.codeartifact.%v.amazonaws.com/pypi/%v/simple/", codeArtifactAuthToken, domain, domainOwner, region, repoName)

lambda.NewPythonFunction(this, jsii.String("function"), &pythonFunctionProps{
	entry: jsii.String(entry),
	runtime: awscdk.Runtime_PYTHON_3_8(),
	bundling: &bundlingOptions{
		buildArgs: map[string]*string{
			"PIP_INDEX_URL": indexUrl,
		},
	},
})
```