# Overview

The Azure Service Operator allows you to manage Azure resources using Kubernetes Custom Resource Definitions.

## Developer's Getting Started guide

### Install Operator SDK
Install the [operator-sdk](https://github.com/operator-framework/operator-sdk) using the instruction details on that repository.

### Start the Operator Locally

```
export AZURE_CLIENT_ID=<client_id>
export AZURE_CLIENT_SECRET=<client_secret>
export AZURE_TENANT_ID=<tenant_id>
export AZURE_SUBSCRIPTION_ID=<subscription_id>
export OPERATOR_NAME=azure-service-operator
operator-sdk up local
```
