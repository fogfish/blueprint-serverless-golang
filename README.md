<p align="center">
  <img src="./doc/logo.gif" width="50%"  />
  <h3 align="center">Blueprint: Serverless Golang</h3>
  <p align="center"><strong>Skeleton of Golang serverless application built with AWS CDK</strong></p>

  <p align="center">
    <!-- Documentation -->
    <a href="https://pkg.go.dev/github.com/fogfish/blueprint-serverless-golang">
      <img src="https://pkg.go.dev/badge/github.com/fogfish/blueprint-serverless-golang" />
    </a>
    <!-- Build Status  -->
    <a href="https://github.com/fogfish/blueprint-serverless-golang/actions/">
      <img src="https://github.com/fogfish/blueprint-serverless-golang/workflows/build/badge.svg" />
    </a>
    <!-- GitHub -->
    <a href="http://github.com/fogfish/blueprint-serverless-golang">
      <img src="https://img.shields.io/github/last-commit/fogfish/blueprint-serverless-golang.svg" />
    </a>
    <!-- Coverage -->
    <a href="https://coveralls.io/github/fogfish/blueprint-serverless-golang?branch=main">
      <img src="https://coveralls.io/repos/github/fogfish/blueprint-serverless-golang/badge.svg?branch=main" />
    </a>
    <!-- Go Card -->
    <a href="https://goreportcard.com/report/github.com/fogfish/blueprint-serverless-golang">
      <img src="https://goreportcard.com/badge/github.com/fogfish/blueprint-serverless-golang" />
    </a>
    <!-- Maintainability -->
    <a href="https://codeclimate.com/github/fogfish/blueprint-serverless-golang/maintainability">
      <img src="https://api.codeclimate.com/v1/badges/1b00f59c9634d9d479cf/maintainability" />
    </a>
  </p>
</p>

--- 

This project crafts a fully functional blueprint of Golang serverless RESTful application for Amazon Web Services. The blueprint is a hybrid solution, composed of pure "application logic" and Infrastructure as a Code implemented on top of AWS CDK, both developed with Golang. 


## Inspiration

[AWS CDK](https://aws.amazon.com/cdk) is amazing technology to automate the development and operation of application into one process and one codebase.

However, seeding of new repository for development of Golang serverless application requires a boilerplate code. This blueprint helps you to focus on the application development than waste a time with establish **project layout**, **configure AWS CDK**, **setting up CI/CD** and figuring out how to **testing the application**. All these issues are resolved within this blueprint.


## Installation

The blueprint is fully functional application (Pet Store) that delivers a skeleton for Golang serverless development with AWS CDK. Clone the repository and follow [Getting started](#getting-started) instructions to evaluate its applicability for your purposes. It should take less than 5 minutes to build and deploy this blueprint to AWS.

```
go get github.com/fogfish/blueprint-serverless-golang
```

See [Getting Started](#getting-started) and [Customize Blueprint](#customize-blueprint) chapters for details.


### Install from GitHub

[**Use this template**](https://github.com/fogfish/blueprint-serverless-golang/generate)

Create a new GitHub repository from this blueprint.


### Upgrade the template

Use `git` features to update the blueprint from upstream

```bash
git remote add blueprint https://github.com/fogfish/blueprint-serverless-golang
git fetch blueprint
git merge blueprint/main --allow-unrelated-histories --squash
```

## Requirements

Before Getting started, you have to ensure

* [Golang](https://golang.org/dl/) development environment v1.16 or later
* [assay-it](https://assay.it) utility for testing cloud apps in production 
* [AWS CDK](https://docs.aws.amazon.com/cdk/latest/guide/work-with.html#work-with-prerequisites) for deployment of serverless application using infrastructure as a code
* [GitHub](https://github.com) account for managing source code and running CI/CD pipelines as [GitHub Actions](https://docs.github.com/en/actions)  
* Account on [Amazon Web Services](https://aws.amazon.com) for running the application in production 


## Getting started

**Let's have a look on the repository structure**

The structure resembles the mixture of [Standard package layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1) and [Hexagonal architecture](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3). The proposed structure is better version of Hexagonal architecture that follows Golang best practices:

1. the root is aws cdk application
    
2. Sub-packages to isolate dependencies to external technologies so that they act as bridge between your domain and technology adaptation. `internal` holds sub-packages internal to applications. `pkg` are sharable clients

4. `cmd` contains main packages that build lambda functions and ties everything together.

```
github.com/.../the-beautiful-app
├─ petshop.go                  // aws cdk main application  
|
├─ internal/awspetshop         // IaC, aws cdk application
|
├─ internal/core               // the root defines domain types, unit test 
|  |                           // "algebra" of your application. contains core
|  |                           // types to describe domain of your application.
|  |                           // It contains simple types that has no dependency
|  |                           // to technology but their implements core logic
|  |                           // and use-cases.
|  |
|  └─ storage.go               // defines capability requires to store core
|                              // objects at the external storage, hex-arch
|                              // use "port" concept to depict it          
|
├─ internal/storage            // sub-package for dependency/technology ...
|                              // it follows the standard package layout to 
|                              // adapt domain/implementation/dependency.
|                              // in this example storage implements in-memory
|                              // database for all domain objects.  
|
├─ internal/services           // entry point to the core, implement app logic
|  └─ pets                     // entire logic about pets domain
|     ├─ fetcher.go            // fetch and enrich pets objects 
|     └─ creator.go            // create pets objects
|
├─ internal/mock               // shared mock
|
├─ internal/http               // public REST API exposed by application.
|  ├─ petshop.go               // collection of petshop endpoints impl. by app
|  |                           // endpoints consumer services using ports    
|  | 
|  └─ suites                   // testing suites for api endpoint(s)
|
├─ cmd                         // executables of the project
|  ├─ lambda                   // aws lambda's are main packages
|  |  ├─ petshop               // each lambda stays at own executable
|  |  |  └─ main.go            // single lambda pattern is not recommended
|  | ...
|  └─ server                   // run application as standalone server 
|     └─ main.go
|
├─ pkg/api                     // public domain objects used by application
|                              // client library
|
└─ .github                     // CI/CD with GitHub Actions
    └─ ...                   
```

### Development workflows

**unit testing**

Test the Golang application and its cloud infrastructure

```bash
go test ./...
```

**local testing**

Run application locally

```bash
go run cmd/server/main.go
assay-it test --target http://127.1:8080
```

**build**

Build entire application (both Golang and its AWS infrastructure). It should compile Golang code, assemble binaries for AWS Lambda and produce AWS CloudFormation template

```bash
cdk synth
```

**deploy**

Deploy an application to AWS account, it requires a valid AWS credentials either access keys or assumed roles.

```bash
cdk deploy
```

In few seconds, the application becomes available at

```bash
curl https://xxxxxxxxxx.execute-api.eu-west-1.amazonaws.com/api
```

The write path of api is protected by AWS IAM, request has to be signed.
Either use example client `cmd/petshop-cli` or curl directly

```bash
curl $BLUEPRINT/petshop/pets \
  -XGET \
  -H "Accept: application/json" \
  --aws-sigv4 "aws:amz:eu-west-1:execute-api" \
  --user "$AWS_ACCESS_KEY_ID":"$AWS_SECRET_ACCESS_KEY"
```

See [all available endpoints](./http/petshop.go). 


**test in production**

```bash
assay-it test --target https://xxxxxxxxxx.execute-api.eu-west-1.amazonaws.com/api
```


**destroy**

Destroy the application and remove all its resource from AWS account

```bash
cdk destroy
```


## Continuos Delivery 

Continuos Integration and Delivery is implemented using GitHub Actions. It consists of multiple [.github/workflows](.github/workflows).

`AWS_ACCESS_KEY` and `AWS_SECRET_ACCESS_KEY` are required to enable deployment by GitHub Actions. Store these credentials to secret key vault at your fork settings (Your Fork > Settings > Secrets).

### Check quality of Pull Request

The quality checks are executed every time a new change is proposed via Pull Request:
* **checks** (`check-code.yml`) evaluates a quality of source code and reviews proposed changes (pull requests) using static code analysis.
* **tests** (`check-test.yml`) the quality of software assets with scope on unit tests only and measures the test coverage.
* **spawns** (`check-spawn.yml`) a sandbox(ed) deployment of the application to target AWS account for continuous integrations (optionally executed if pull request is marked with `[@] deploy` label);
* **cleans** (`check-clean.yml`) sandbox environment after Pull Request is either merged or closed.

### Check quality of `main` branch

The quality checks are executed every time a pull request is merged into pipeline:
* **tests** (`check-test.yml`) the quality of software assets with scope on unit tests only and measures the test coverage.
* **builds** (`build.yml`) validates quality of `main` branch once Pull Request is merge by deploying changes to the development environment at target AWS account;

### Release of `main` branch

The quality checks are executed every time a new release is created:
* **carries** (`carry.yml`) "immutable" application snapshot to production environment when GitHub release is published;


## Customize Blueprint

- [ ] rebuild go.mod and go.sum for your application
- [ ] add RESTful api endpoints to [http](http) package
- [ ] add Lambda functions to [aws/lambda](aws/lambda) package
- [ ] set the name of your stack at [cloud/blueprint.go](cloud/src/blueprint.go) and enhance the infrastructure
```go
stackID := fmt.Sprintf("blueprint-golang-%s", vsn(app))
stack := awscdk.NewStack(app, jsii.String(stackID), config)
```
- [ ] update the target stack name at CI/CD workflows [check-spawn.yml](.github/workflows/check-spawn.yml), [build.yml](.github/workflows/build.yml), [carry.yml](.github/workflows/carry.yml) and [check-clean.yml](.github/workflows/check-clean.yml)
```yaml
strategy:
      matrix:
        stack: [blueprint-golang]
```
- [ ] setup access to AWS account for CI/CD
- [ ] integrate api testing 
- [ ] tune CI/CD pipeline according to purpose of your application either removing or commenting out blocks


## How To Contribute

The blueprint is [MIT](LICENSE) licensed and accepts contributions via GitHub pull requests:

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Added some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request


### commit message

The commit message helps us to write a good release note, speed-up review process. The message should address two question what changed and why. The project follows the template defined by chapter [Contributing to a Project](http://git-scm.com/book/ch5-2.html) of Git book.

### bugs

If you experience any issues with the library, please let us know via [GitHub issues](https://github.com/fogfish/blueprint-serverless-golang/issue). We appreciate detailed and accurate reports that help us to identity and replicate the issue. 

## License

[![See LICENSE](https://img.shields.io/github/license/fogfish/blueprint-serverless-golang.svg?style=for-the-badge)](LICENSE)
