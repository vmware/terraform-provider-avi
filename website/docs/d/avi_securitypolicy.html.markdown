############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_securitypolicy"
sidebar_current: "docs-avi-datasource-securitypolicy"
description: |-
  Get information of Avi SecurityPolicy.
---

# avi_securitypolicy

This data source is used to to get avi_securitypolicy objects.

## Example Usage

```hcl
data "avi_securitypolicy" "foo_securitypolicy" {
    uuid = "securitypolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SecurityPolicy by name.
* `uuid` - (Optional) Search SecurityPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - Security policy is used to specify various configuration information used to perform distributed denial of service (ddos) attacks detection and mitigation.
* `dns_attacks` - Attacks utilizing the dns protocol operations.
* `dns_policy_index` - Index of the dns policy to use for the mitigation rules applied to the dns attacks.
* `labels` - Key value pairs for granular object access control.
* `name` - The name of the security policy.
* `network_security_policy_index` - Index of the network security policy to use for the mitigation rules applied to the attacks.
* `oper_mode` - Mode of dealing with the attacks - perform detection only, or detect and mitigate the attacks.
* `tcp_attacks` - Attacks utilizing the tcp protocol operations.
* `tenant_ref` - Tenancy of the security policy.
* `udp_attacks` - Attacks utilizing the udp protocol operations.
* `uuid` - The uuid of the security policy.

