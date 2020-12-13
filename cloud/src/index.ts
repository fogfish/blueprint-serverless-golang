import * as cdk from '@aws-cdk/core'
import * as pure from 'aws-cdk-pure'
import service from './service'

// ----------------------------------------------------------------------------
//
// Config
//
// ----------------------------------------------------------------------------
const app = new cdk.App()
const config = {
  env: {
    account: process.env.CDK_DEFAULT_ACCOUNT,
    region: process.env.CDK_DEFAULT_REGION,
  }
}

// ----------------------------------------------------------------------------
//
// Stack
//
// ----------------------------------------------------------------------------
const stack = new cdk.Stack(app, `scud`, { ...config })

pure.join(stack, service)
