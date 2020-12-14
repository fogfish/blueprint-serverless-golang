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
    <a href="https://coveralls.io/github/fogfish/blueprint-serverless-golang?branch=master">
      <img src="https://coveralls.io/repos/github/fogfish/blueprint-serverless-golang/badge.svg?branch=master" />
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

This project crafts a fully functional template of Golang serverless RESTful application for Amazon Web Services. The template is a hybrid solution, composed of pure "application logic" developed with Golang, TypeScript-based Infrastructure as a Code implemented on top of AWS CDK and couple of open source libraries to flavour development experience. 


## Inspiration

[AWS CDK](https://aws.amazon.com/cdk) is amazing technology to automate the development and operation of application into one process and one codebase.

However, seeding of new repository for development of Golang serverless application requires a boilerplate code. This template helps you to focus on the application development than waste a time with establish project layout, configure AWS CDK, setting up CI/CD and figuring out how to testing the application. All this is resolved within this template.

## Installation

This template is fully functional "Hello World" like application that delivers a skeleton for Golang serverless development with AWS CDK. Clone the repository and follow [Getting started](#getting-started) instructions to evaluate its applicability for your purposes. It should take less than 5 minutes to build and deploy the template.

```
go get github.com/fogfish/blueprint-serverless-golang
```
 
<!-- 

TODO:
 * How To use GitHub template feature
 * How To upgrade the template in existing app
 * How To Customize Template

-->

## Getting started

**Let's have a look on the repository structure**

The structure resembles the standard package layout proposed in [this blog post](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1):

1. the root package contains core types to describe domain of your application. A simple types that has no dependency to technology.

2. Use sub-packages to isolate dependencies to external technologies so that they act as bridge between your domain and technology adaptation. 

3. Main packages build lambda functions and ties everything together.

```
github.com/.../the-beautiful-app  
  ├─ stub.go                     // domain types and its tests
  ├─ ...                         // other domain types
  |
  ├─ http                        // http and rest adaptations
  |    ├─ api.go                 // api endpoint(s) and its tests
  |    └─ ...                    // other endpoints
  |
  ├─ aws                         // adaptations to AWS
  |    ├─ ...                    // any code to deal with AWS APIs
  |    └─ lambda                 // aws lambda's are main packages
  |         ├─ scud              // each lambda stays at own pkg
  |         |    └─ main.go
  |         └─ ...
  |
  ├─ cloud                       // IaC, aws cdk application
  |    └─ ...                    // It orchestrate everything
  |
  ├─ .github
  |      └─ ...                   // CI/CD with GitHub Actions
  |
  └─ suite                        // api testing suite 
       ├─ api.go
       └─ ... 
```

### Development workflows

**dependencies** 

The application requires 3rd party libraries for dev and opts. Fetch them with the following commands:

```bash
# Golang deps
go get -d ./...

# AWS CDK TypeScript deps
npm -C cloud install
```

**unit testing**

Test the Golang application and its cloud infrastructure

```bash
# Golang unit testing
go test ./...

# AWS CDK TypeScript testing
npm -C cloud run test
npm -C cloud run lint
```

**build**

Build entire application (both Golang and its AWS Infrastructure). It should compile Golang code, assemble binaries for AWS Lambda and produce AWS CloudFormation template

```bash
npm -C cloud run -- cdk synth 
```

**deploy**

Deploy an application to AWS account, it requires a valid AWS credentials either access keys or assumed roles.

```bash
npm -C cloud run -- cdk deploy
```

In few seconds, the application becomes available at

```
curl https://xxxxxxxxxx.execute-api.eu-west-1.amazonaws.com/api/scud
```

**destroy**

Destroy the application and remove all its resource from AWS account

```bash
npm -C cloud run -- cdk destroy
```


## Continuos Delivery 

Continuos Integration and Delivery is implemented using GitHub Actions. It consists of multiple [.github/workflows](.github/workflows):

* **checks** (`golang.yml`, `cloud.yml`) the quality of software assets with scope on unit tests only. Checks are executed in parallel for application logic and infrastructure every time a new change is proposed via Pull Request.
* **spawns** (`spawn.yml`) a sandbox(ed) deployment of the application to target AWS account for continuous integrations;
* **builds** (`build.yml`) validates quality of `main` branch once Pull Request is merge. Upon the quality check completion, the pipeline deploys changes to the development environment at target AWS account;
* **carries** (`carry.yml`) "immutable" application snapshot to production environment when GitHub release is published;
* **cleans** (`clean.yml`) sandbox environment after Pull Request is either merged or closed.


## How To Contribute

The library is [MIT](LICENSE) licensed and accepts contributions via GitHub pull requests:

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
