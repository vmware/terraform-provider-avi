<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_customipamdnsprofile"
sidebar_current: "docs-avi-datasource-customipamdnsprofile"
description: |-
  Get information of Avi CustomIpamDnsProfile.
---

# avi_customipamdnsprofile

This data source is used to to get avi_customipamdnsprofile objects.

## Example Usage

```hcl
data "avi_customipamdnsprofile" "foo_customipamdnsprofile" {
    uuid = "customipamdnsprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search CustomIpamDnsProfile by name.
* `uuid` - (Optional) Search CustomIpamDnsProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `name` - Name of the custom ipam dns profile. Field introduced in 17.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `script_params` - Parameters that are always passed to the ipam/dns script. Field introduced in 17.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `script_uri` - Script uri of form controller //ipamdnsscripts/<file-name>, file-name must have a .py extension and conform to pep8 naming convention. Field introduced in 17.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 17.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Field introduced in 17.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

