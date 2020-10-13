############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_ipaddrgroup"
sidebar_current: "docs-avi-resource-ipaddrgroup"
description: |-
  Creates and manages Avi IpAddrGroup.
---

# avi_ipaddrgroup

The IpAddrGroup resource allows the creation and management of Avi IpAddrGroup

## Example Usage

```hcl
resource "avi_ipaddrgroup" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the ip address group.
* `addrs` - (Optional) Configure ip address(es).
* `apic_epg_name` - (Optional) Populate ip addresses from members of this cisco apic epg.
* `country_codes` - (Optional) Populate the ip address ranges from the geo database for this country.
* `description` - (Optional) User defined description for the object.
* `ip_ports` - (Optional) Configure (ip address, port) tuple(s).
* `labels` - (Optional) Key value pairs for granular object access control.
* `marathon_app_name` - (Optional) Populate ip addresses from tasks of this marathon app.
* `marathon_service_port` - (Optional) Task port associated with marathon service port.
* `prefixes` - (Optional) Configure ip address prefix(es).
* `ranges` - (Optional) Configure ip address range(s).
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the ip address group.

