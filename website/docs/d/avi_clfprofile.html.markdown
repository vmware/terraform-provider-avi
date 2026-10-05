<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_clfprofile"
sidebar_current: "docs-avi-datasource-clfprofile"
description: |-
  Get information of Avi ClfProfile.
---

# avi_clfprofile

This data source is used to to get avi_clfprofile objects.

## Example Usage

```hcl
data "avi_clfprofile" "foo_clfprofile" {
    uuid = "clfprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ClfProfile by name.
* `uuid` - (Optional) Search ClfProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `clf_pools` - List of pools associated with this profile. Each pool must have a unique priority value. The controller rejects profiles where two pools share the same priority (http 400). Field introduced in 32.1.5. Maximum of 8 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning and config-push attributes. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `description` - Human-readable description for this clf profile. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `enabled` - Enable or disable log delivery for this profile without disturbing pool state, health monitors, or virtualservice/datascriptset bindings. When false, datascript log forwarding is a silent no-op for every vs attached via this profile; delivery resumes immediately when set back to true. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - The name of the clf profile. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `replicate` - When false (default), log records are routed to the highest-priority pool that has at least one up member (priority-based failover). When true, log records are replicated to all pools regardless of priority. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - Reference to the tenant that owns this clf profile. It is a reference to an object of type tenant. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the clf profile. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `vrf_context_ref` - Virtual routing context for this profile's collector pools. Only virtual services in the same virtual routing context can use this profile. Cannot be changed once set. It is a reference to an object of type vrfcontext. Field introduced in 32.1.5. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

