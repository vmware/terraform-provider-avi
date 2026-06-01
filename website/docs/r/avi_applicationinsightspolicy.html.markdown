<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_applicationinsightspolicy"
sidebar_current: "docs-avi-resource-applicationinsightspolicy"
description: |-
  Creates and manages Avi ApplicationInsightsPolicy.
---

# avi_applicationinsightspolicy

The ApplicationInsightsPolicy resource allows the creation and management of Avi ApplicationInsightsPolicy

## Example Usage

```hcl
resource "avi_applicationinsightspolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the application insights configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `application_insights_params` - (Optional) Application insights parameters to filter application learning from clients. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `application_sampling_config` - (Optional) Application sampling configuration to control rate and volume of data ingestion for application insights that the serviceengines are expected to send to the controller. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `description` - (Optional) Details of the application insights configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `enable_application_insights` - (Optional) Enable application insights, formerly called learning for this virtual service. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tenant_ref` - (Optional) Details of the tenant for the application insights configuration. It is a reference to an object of type tenant. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the application insights configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

