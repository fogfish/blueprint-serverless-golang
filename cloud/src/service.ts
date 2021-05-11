import * as lambda from '@aws-cdk/aws-lambda'
import * as api from '@aws-cdk/aws-apigateway'
import * as scud from 'aws-scud'
import * as path from 'path'

//
const MyFun = (): lambda.FunctionProps =>
  scud.handler.Go({
    sourceCodePackage: path.join(__dirname, '../..'),
    sourceCodeLambda: 'aws/lambda/scud',
  })

//
const Gateway = (): api.RestApiProps => scud.Gateway({
  restApiName: 'scud',
})

export default scud.mkService(Gateway)
  .addResource('scud', scud.aws.Lambda(MyFun))
