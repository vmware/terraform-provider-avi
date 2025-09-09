<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_applicationinsightsstate"
sidebar_current: "docs-avi-datasource-applicationinsightsstate"
description: |-
  Get information of Avi ApplicationInsightsState.
---

# avi_applicationinsightsstate

This data source is used to to get avi_applicationinsightsstate objects.

## Example Usage

```hcl
data "avi_applicationinsightsstate" "foo_applicationinsightsstate" {
    uuid = "applicationinsightsstate-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ApplicationInsightsState by name.
* `uuid` - (Optional) Search ApplicationInsightsState by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `application_insights_uuid` - Uuid of the application insights policy. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `application_sampling_runtime` - Runtime application sampling configuration to control rate and volume of data ingestion for application insights. Controller updates the configuration based on the application traffic and the associated serviceengine load. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - The name of the application insights state configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - Details of the tenant for the application insights state. It is a reference to an object of type tenant. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `uuid` - Uuid of the applicationinsightsstate. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.

