<p align="center">
  <img src="./doc/logo.gif" height="240" />
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

However, seeding of new repository for development of Golang serverless application requires a boilerplate code. This blueprint helps you to focus on the application development than waste a time with establish project layout, configure AWS CDK, setting up CI/CD and figuring out how to testing the application. All these issues are resolved within this blueprint.

## Installation

The blueprint is fully functional application that delivers a skeleton for Golang serverless development with AWS CDK. Clone the repository and follow [Getting started](#getting-started) instructions to evaluate its applicability for your purposes. It should take less than 5 minutes to build and deploy the template in AWS.

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
* [AWS CDK](https://docs.aws.amazon.com/cdk/latest/guide/work-with.html#work-with-prerequisites)
* Access to AWS Account


## Getting started

**Let's have a look on the repository structure**

The structure resembles the standard package layout proposed in [this blog post](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1):

1. the root package contains core types to describe domain of your application. It contains simple types that has no dependency to technology but their implements core logic and use-cases.

2. Use sub-packages to isolate dependencies to external technologies so that they act as bridge between your domain and technology adaptation. 

3. Main packages build lambda functions and ties everything together.

```
github.com/.../the-beautiful-app  
  ├─ stub.go                     // domain types, unit test
  ├─ ...                         // "algebra" of your application
  |
  ├─ http                        // RESTful API and HTTP protocol
  |    ├─ api.go                 // api endpoint(s), unit tests,
  |    └─ ...                    // other endpoints
  |
  ├─ cmd                         // executables of the project
  |    └─ lambda                 // aws lambda's are main packages
  |         ├─ scud              // each lambda stays at own executable
  |         |    └─ main.go
  |         └─ ...
  |
  ├─ cloud                       // IaC, aws cdk application
  |    └─ ...                    // orchestrate ops model
  |
  ├─ .github                     // CI/CD with GitHub Actions
  |      └─ ...                   
  |
  └─ suite                       // API testing suite 
       ├─ api.go                 // (disabled in this release)
       └─ ... 
```

### Development workflows

**dependencies** 

The application requires 3rd party libraries for dev and opts. Fetch them with the following commands:

```bash
go get -d ./...
```

**unit testing**

Test the Golang application and its cloud infrastructure

```bash
go test ./...
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

```
curl https://xxxxxxxxxx.execute-api.eu-west-1.amazonaws.com/api/scud
```

**destroy**

Destroy the application and remove all its resource from AWS account

```bash
cdk destroy
```


## Continuos Delivery 

Continuos Integration and Delivery is implemented using GitHub Actions. It consists of multiple [.github/workflows](.github/workflows).

### Check quality of Pull Request

### Check quality of `main` branch

* **check** (`check.yml`) the quality of software assets with scope on unit tests only. Checks are executed in parallel for application logic and infrastructure every time a new change is proposed via Pull Request.

* **tests** (`tests.yml`) unit + cov

* **spawns** (`spawn.yml`) a sandbox(ed) deployment of the application to target AWS account for continuous integrations;

* **builds** (`build.yml`) validates quality of `main` branch once Pull Request is merge. Upon the quality check completion, the pipeline deploys changes to the development environment at target AWS account;

* **carries** (`carry.yml`) "immutable" application snapshot to production environment when GitHub release is published;

* **cleans** (`clean.yml`) sandbox environment after Pull Request is either merged or closed.

`AWS_ACCESS_KEY` and `AWS_SECRET_ACCESS_KEY` are required to enable deployment by GitHub Actions. Store these credentials to secret key vault at your fork settings (Your Fork > Settings > Secrets).


## Customize Blueprint

- [ ] rebuild go.mod and go.sum for your application
- [ ] add RESTful api endpoints to [http](http) package
- [ ] add Lambda functions to [aws/lambda](aws/lambda) package
- [ ] set the name of your stack at [cloud/blueprint.go](cloud/src/blueprint.go) and enhance the infrastructure
```go
stackID := fmt.Sprintf("blueprint-golang-%s", vsn(app))
stack := awscdk.NewStack(app, jsii.String(stackID), config)
```
- [ ] update the target stack name at CI/CD workflows [spawn.yml](.github/workflows/spawn.yml), [build.yml](.github/workflows/build.yml) and [carry.yml](.github/workflows/carry.yml)
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
