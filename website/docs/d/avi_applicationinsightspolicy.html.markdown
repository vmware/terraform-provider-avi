<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_applicationinsightspolicy"
sidebar_current: "docs-avi-datasource-applicationinsightspolicy"
description: |-
  Get information of Avi ApplicationInsightsPolicy.
---

# avi_applicationinsightspolicy

This data source is used to to get avi_applicationinsightspolicy objects.

## Example Usage

```hcl
data "avi_applicationinsightspolicy" "foo_applicationinsightspolicy" {
    uuid = "applicationinsightspolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ApplicationInsightsPolicy by name.
* `uuid` - (Optional) Search ApplicationInsightsPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `application_insights_params` - Application insights parameters to filter application learning from clients. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `application_sampling_config` - Application sampling configuration to control rate and volume of data ingestion for application insights that the serviceengines are expected to send to the controller. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `description` - Details of the application insights configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `enable_application_insights` - Enable application insights, formerly called learning for this virtual service. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `enable_application_sampling` - Enable application sampling. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `name` - The name of the application insights configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - Details of the tenant for the application insights configuration. It is a reference to an object of type tenant. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `uuid` - Uuid of the application insights configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.

