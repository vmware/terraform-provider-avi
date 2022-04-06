<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_ipaddrgroup"
sidebar_current: "docs-avi-datasource-ipaddrgroup"
description: |-
  Get information of Avi IpAddrGroup.
---

# avi_ipaddrgroup

This data source is used to to get avi_ipaddrgroup objects.

## Example Usage

```hcl
data "avi_ipaddrgroup" "foo_ipaddrgroup" {
    uuid = "ipaddrgroup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search IpAddrGroup by name.
* `uuid` - (Optional) Search IpAddrGroup by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `addrs` - Configure ip address(es). Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `country_codes` - Populate the ip address ranges from the geo database for this country. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `description` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `ip_ports` - Configure (ip address, port) tuple(s). Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `marathon_app_name` - Populate ip addresses from tasks of this marathon app. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `marathon_service_port` - Task port associated with marathon service port. If marathon app has multiple service ports, this is required. Else, the first task port is used. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `name` - Name of the ip address group. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `prefixes` - Configure ip address prefix(es). Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `ranges` - Configure ip address range(s). Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Uuid of the ip address group. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

