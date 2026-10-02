<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_clfprofile"
sidebar_current: "docs-avi-resource-clfprofile"
description: |-
  Creates and manages Avi ClfProfile.
---

# avi_clfprofile

The ClfProfile resource allows the creation and management of Avi ClfProfile

## Example Usage

```hcl
resource "avi_clfprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the clf profile. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `vrf_context_ref` - (Required) Virtual routing context for this profile's collector pools. Only virtual services in the same virtual routing context can use this profile. Cannot be changed once set. It is a reference to an object of type vrfcontext. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. Changing this value forces the resource to be recreated.
* `clf_pools` - (Optional) List of pools associated with this profile. Each pool must have a unique priority value. The controller rejects profiles where two pools share the same priority (http 400). Field introduced in 32.1.5. Maximum of 8 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `configpb_attributes` - (Optional) Protobuf versioning and config-push attributes. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `description` - (Optional) Human-readable description for this clf profile. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `enabled` - (Optional) Enable or disable log delivery for this profile without disturbing pool state, health monitors, or virtualservice/datascriptset bindings. When false, avi.vs.log_forward() is a silent no-op for every vs attached via this profile; delivery resumes immediately when set back to true. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `replicate` - (Optional) When false (default), log records are routed to the highest-priority pool that has at least one up member (priority-based failover). When true, log records are replicated to all pools regardless of priority. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tenant_ref` - (Optional) Reference to the tenant that owns this clf profile. It is a reference to an object of type tenant. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the clf profile. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

