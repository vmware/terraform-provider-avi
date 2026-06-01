<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_reportprofile"
sidebar_current: "docs-avi-datasource-reportprofile"
description: |-
  Get information of Avi ReportProfile.
---

# avi_reportprofile

This data source is used to to get avi_reportprofile objects.

## Example Usage

```hcl
data "avi_reportprofile" "foo_reportprofile" {
    uuid = "reportprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ReportProfile by name.
* `uuid` - (Optional) Search ReportProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `collection_rules` - Collection rules for the report. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `max_concurrent_reports` - Maximum number of concurrent reports allowed to be generated. Allowed values are 1-10. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `remote_controller` - Remote controller request to enable report generation for remote controller. If enabled, the report generation will be done for the remote controller. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid identifier for the reportprofile object. Field introduced in 31.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

