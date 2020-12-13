# scud-golang

This project is a template of Golang serverless RESTful application for Amazon Web Services. This is a hybrid of pure Golang code (application logic), TypeScript AWS CDK (Infrastructure as a Code) and few libraries to flavour development experience. 


## Getting started

**Let's have a look on the repository structure**

The structure resembles the standard package layout proposed by [in this blog post](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1) with a few simple rules:

1. the root package contains core types to describe domain of your application. A simple types that has no dependency to technology.

2. Use sub-packages to isolate dependencies to external technologies so that this is a bridge between your domain and technology adaptation. 

3. Main package builds a lambda code and ties everything together. The application might produce multiple binaries (aka multiple lambda functions)

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
  ├─ cloud                       // aws cdk application and IaC 
  |    └─ ...                    // It orchestrate everything
  |
  └─ .github
        └─ ...                   // CI/CD with GitHub Actions

```

### Dependencies 

The application requires 3rd party libraries for dev and opts. Fetch them with the following commands:

```bash
go get -d ./...
npm -C cloud install
```

### Test

Test the Golang application only and its IaC

```bash
go test ./...
npm -C cloud run test
npm -C cloud run lint
```

### Build

Build entire application (both Golang and its AWS Infrastructure). It should compile Golang code, assemble binaries for AWS Lambda and produce AWS CloudFormation template

```bash
cd cloud
npm run cdk synth
```

### Deploy

Deploy an application to AWS account, it requires a valid AWS credentials either access keys or assumed roles.

```bash
cd cloud
npm run cdk deploy
```

In few seconds, the application becomes available at

```
curl https://xxxxxxxxxx.execute-api.eu-west-1.amazonaws.com/api/scud
```

### Destroy

Destroy the application and remove all its AWS resource

```bash
cd cloud
npm run cdk destroy
```

### CI/CD


