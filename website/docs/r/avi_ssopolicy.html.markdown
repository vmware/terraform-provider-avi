############################################################################
# ========================================================================
# Copyright 2021 VMware, Inc.  All rights reserved. VMware Confidential
# ========================================================================
###

<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_ssopolicy"
sidebar_current: "docs-avi-resource-ssopolicy"
description: |-
  Creates and manages Avi SSOPolicy.
---

# avi_ssopolicy

The SSOPolicy resource allows the creation and management of Avi SSOPolicy

## Example Usage

```hcl
resource "avi_ssopolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the sso policy. Field introduced in 18.2.3.
* `authentication_policy` - (Optional) Authentication policy settings. Field introduced in 18.2.1.
* `authorization_policy` - (Optional) Authorization policy settings. Field introduced in 18.2.5.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `tenant_ref` - (Optional) Uuid of the tenant. It is a reference to an object of type tenant. Field introduced in 18.2.3.
* `type` - (Optional) Sso policy type. Enum options - SSO_TYPE_SAML, SSO_TYPE_PINGACCESS, SSO_TYPE_JWT, SSO_TYPE_LDAP, SSO_TYPE_OAUTH. Field introduced in 18.2.5.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the sso policy. Field introduced in 18.2.3.

