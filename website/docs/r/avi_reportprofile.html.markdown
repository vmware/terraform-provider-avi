<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_reportprofile"
sidebar_current: "docs-avi-resource-reportprofile"
description: |-
  Creates and manages Avi ReportProfile.
---

# avi_reportprofile

The ReportProfile resource allows the creation and management of Avi ReportProfile

## Example Usage

```hcl
resource "avi_reportprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `collection_rules` - (Optional) Collection rules for the report. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `max_concurrent_reports` - (Optional) Maximum number of concurrent reports allowed to be generated. Allowed values are 1-10. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `remote_controller` - (Optional) Remote controller request to enable report generation for remote controller. If enabled, the report generation will be done for the remote controller. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid identifier for the reportprofile object. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.

