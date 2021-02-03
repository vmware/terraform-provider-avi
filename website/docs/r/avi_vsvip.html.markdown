############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_vsvip"
sidebar_current: "docs-avi-resource-vsvip"
description: |-
  Creates and manages Avi VsVip.
---

# avi_vsvip

The VsVip resource allows the creation and management of Avi VsVip

## Example Usage

```hcl
resource "avi_vsvip" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the vsvip object. Field introduced in 17.1.1.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud. Field introduced in 17.1.1.
* `dns_info` - (Optional) Service discovery specific data including fully qualified domain name, type and time-to-live of the dns record. Field introduced in 17.1.1. Maximum of 1000 items allowed. Allowed in basic edition, essentials edition, enterprise edition.
* `east_west_placement` - (Optional) Force placement on all service engines in the service engine group (container clouds only). Field introduced in 17.1.1. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `ipam_selector` - (Optional) Determines the set of ipam networks to use for this vsvip. Selector type must be selector_ipam and only one label is supported. Field introduced in 20.1.3.
* `labels` - (Optional) Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 17.1.1.
* `tier1_lr` - (Optional) This sets the placement scope of virtualservice to given tier1 logical router in nsx-t. Field introduced in 20.1.1.
* `use_standard_alb` - (Optional) This overrides the cloud level default and needs to match the se group value in which it will be used if the se group use_standard_alb value is set. This is only used when fip is used for vs on azure cloud. Field introduced in 18.2.3. Allowed in basic edition, essentials edition, enterprise edition.
* `vip` - (Optional) List of virtual service ips and other shareable entities. Field introduced in 17.1.1.
* `vrf_context_ref` - (Optional) Virtual routing context that the virtual service is bound to. This is used to provide the isolation of the set of networks the application is attached to. It is a reference to an object of type vrfcontext. Field introduced in 17.1.1.
* `vsvip_cloud_config_cksum` - (Optional) Checksum of cloud configuration for vsvip. Internally set by cloud connector. Field introduced in 17.2.9, 18.1.2.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the vsvip object. Field introduced in 17.1.1.

