<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_systemreport"
sidebar_current: "docs-avi-resource-systemreport"
description: |-
  Creates and manages Avi SystemReport.
---

# avi_systemreport

The SystemReport resource allows the creation and management of Avi SystemReport

## Example Usage

```hcl
resource "avi_systemreport" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `archive_ref` - (Optional) Relative path to the report archive file on filesystem.the archive includes exported system configuration and current object as json. Field introduced in 22.1.6, 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `controller_patch_image_ref` - (Optional) Controller patch image associated with the report. It is a reference to an object of type image. Field introduced in 22.1.6, 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `downloadable` - (Optional) Indicates whether this report is downloadable as an archive. Field introduced in 22.1.6, 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `events` - (Optional) List of events associated with the report. Field introduced in 22.1.6, 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `image_ref` - (Optional) System image associated with the report. It is a reference to an object of type image. Field introduced in 22.1.6, 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - (Optional) Name of the report derived from operation in a readable format. Ex  upgrade_system_1a5c. Field introduced in 22.1.6, 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `readiness_reports` - (Optional) Readiness state of the system. Ex  upgrade pre-check results. Field introduced in 22.1.6, 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `se_patch_image_ref` - (Optional) Se patch image associated with the report. It is a reference to an object of type image. Field introduced in 22.1.6, 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `state` - (Optional) Report state combines all applicable states. Ex  readiness_reports.system_readiness.state. Field introduced in 22.1.6, 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `summary` - (Optional) Summary of the report. Field introduced in 22.1.6, 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant uuid associated with the object. It is a reference to an object of type tenant. Field introduced in 22.1.6, 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid identifier for the report. Field introduced in 22.1.6, 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

