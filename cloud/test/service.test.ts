import * as assert from '@aws-cdk/assert'
import * as cdk from '@aws-cdk/core'
import * as pure from 'aws-cdk-pure'
import service from '../src/service'

//
//
it('creates application RESTful API', () => {
  const stack = new cdk.Stack()
  pure.join(stack, service)

  const requires: {[key: string]: number} = {
    'AWS::ApiGateway::RestApi': 1,
    'AWS::ApiGateway::Deployment': 1,
    'AWS::ApiGateway::Stage': 1,
    'AWS::ApiGateway::Method': 5,
    'AWS::IAM::Role': 3,
    'AWS::Lambda::Function': 2,
    'Custom::LogRetention': 1,
  }

  Object.keys(requires).forEach(
    key => assert.expect(stack).to(
      assert.countResources(key, requires[key])
    )
  )
})