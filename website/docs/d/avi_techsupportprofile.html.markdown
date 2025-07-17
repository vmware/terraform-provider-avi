<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_techsupportprofile"
sidebar_current: "docs-avi-datasource-techsupportprofile"
description: |-
  Get information of Avi TechSupportProfile.
---

# avi_techsupportprofile

This data source is used to to get avi_techsupportprofile objects.

## Example Usage

```hcl
data "avi_techsupportprofile" "foo_techsupportprofile" {
    uuid = "techsupportprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search TechSupportProfile by name.
* `uuid` - (Optional) Search TechSupportProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `archive_rules` - Defined policy for tech-support archive rules.these are predefined files which are exception for default file size thresholduser can add file path with custom threshold in allowed limits to be collected in bundlee.g. A file /var/sample.log is with size 450mb needs to be collected for each invocationuser should configure and add path in techsupportprofile. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `collect_customer_files` - A list of user-specified file paths for collectionthat are not part of the predefined yaml configuration. This is useful forcollecting logs from third-party applications or other custom files.e.g. A file located at /var/sample.log which is not a part of pre-define yamluser should configure this path as source in collect_customer_files so that subsequent collectioncollect this file, once user no longer needed this file they can remove from techsupportprofile. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `event_params` - Specify this params to set threshold for event files.user provided parameters will take precedence over the profile parameters. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `file_size_threshold` - Max file size threshold to archive in tech-support collectionfiles above this threshold will not be collected and an warning will be flagged. Allowed values are 128-512. Field introduced in 31.2.1. Unit is mb. Allowed with any value in enterprise, enterprise with cloud services edition.
* `max_disk_size_percent` - Max disk size in percent of total disk size reserved for the tech-support.the value is in percentage to make it agnostic of controller flavors.e.g. Small [disk=5 gb, ts space available = 500mb]large [ disk= 100gb, ts space available= 10gb]xl [disk=1tb, ts space available=100gb]. Allowed values are 10-25. Field introduced in 31.2.1. Unit is percent. Allowed with any value in enterprise, enterprise with cloud services edition.
* `min_free_disk_required` - Min free disk required for the tech-support invocation.the value is in percentage to make it agnostic of controller flavors.e.g. Small [disk=5 gb, ts space available = 250mb]large [ disk= 100gb, ts space available= 5gb]xl [disk=1tb, ts space available=50gb]. Allowed values are 5-10. Field introduced in 31.2.1. Unit is percent. Allowed with any value in enterprise, enterprise with cloud services edition.
* `no_of_techsupport_retentions` - Number of techsupport to retain from techsupport cleanup policy. Allowed values are 1-5. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `simultaneous_invocations` - Number of simultaneous tech-support invocation allowed. Allowed values are 1-2. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `task_timeout` - Generic timeout for tech-support task collection.this can be used for task, script executions etc.tweak the timeout value in cases of timeout observation in the logs. Field introduced in 31.2.1. Unit is sec. Allowed with any value in enterprise, enterprise with cloud services edition.
* `uuid` - Uuid identifier for the tech-support profile. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.

