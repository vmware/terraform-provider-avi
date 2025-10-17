<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_techsupportprofile"
sidebar_current: "docs-avi-resource-techsupportprofile"
description: |-
  Creates and manages Avi TechSupportProfile.
---

# avi_techsupportprofile

The TechSupportProfile resource allows the creation and management of Avi TechSupportProfile

## Example Usage

```hcl
resource "avi_techsupportprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `archive_rules` - (Optional) Define the policy for techsupport archive rules. These rules allow you to specify files that should be collected in the techsupport bundle, even if they exceed the default file size threshold. E.g. To ensure a 450mb file, such as /var/sample.log, is collected with every invocation, configure and add its path to the techsupportprofile. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `event_params` - (Optional) Specify this params to set threshold for event files. User provided parameters will take precedence over the profile parameters. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `file_size_threshold` - (Optional) Max file size threshold to archive in techsupport collection. Files above this threshold will not be collected and an warning will be flagged. Allowed values are 128-512. Field introduced in 31.2.1. Unit is mb. Allowed with any value in enterprise, enterprise with cloud services edition.
* `max_disk_size_percent` - (Optional) Max disk size in percent of total disk size reserved for the techsupport. The value is in percentage to make it agnostic of controller flavors. E.g. Small [disk=5 gb, ts space available = 500mb] large [ disk= 100gb, ts space available= 10gb] xl [disk=1tb, ts space available=100gb]. Allowed values are 10-25. Field introduced in 31.2.1. Unit is percent. Allowed with any value in enterprise, enterprise with cloud services edition.
* `min_free_disk_required` - (Optional) Min free disk required for the techsupport invocation. The value is in percentage to make it agnostic of controller flavors. E.g. Small [disk=5 gb, ts space available = 250mb] large [ disk= 100gb, ts space available= 5gb] xl [disk=1tb, ts space available=50gb]. Allowed values are 5-10. Field introduced in 31.2.1. Unit is percent. Allowed with any value in enterprise, enterprise with cloud services edition.
* `no_of_techsupport_retentions` - (Optional) Number of techsupport to retain from techsupport cleanup policy. Allowed values are 1-5. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `simultaneous_invocations` - (Optional) Number of simultaneous techsupport invocation allowed. Allowed values are 1-2. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `task_timeout` - (Optional) Generic timeout for techsupport task collection. This can be used for task, script executions etc. Tweak the timeout value in cases of timeout observation in the logs. Field introduced in 31.2.1. Unit is sec. Allowed with any value in enterprise, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid identifier for the techsupport profile. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.

