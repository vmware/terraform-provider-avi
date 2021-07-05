<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `description` - Security policy is used to specify various configuration information used to perform distributed denial of service (ddos) attacks detection and mitigation. Field introduced in 18.2.1.
* `dns_amplification_denyports` - Source ports and port ranges to deny in dns amplification attacks. Field introduced in 21.1.1.
* `dns_attacks` - Attacks utilizing the dns protocol operations. Field introduced in 18.2.1.
* `dns_policy_index` - Index of the dns policy to use for the mitigation rules applied to the dns attacks. Field introduced in 18.2.1.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `name` - The name of the security policy. Field introduced in 18.2.1.
* `network_security_policy_index` - Index of the network security policy to use for the mitigation rules applied to the attacks. Field introduced in 18.2.1.
* `oper_mode` - Mode of dealing with the attacks - perform detection only, or detect and mitigate the attacks. Enum options - DETECTION, MITIGATION. Field introduced in 18.2.1.
* `tcp_attacks` - Attacks utilizing the tcp protocol operations. Field introduced in 18.2.1.
* `tenant_ref` - Tenancy of the security policy. It is a reference to an object of type tenant. Field introduced in 18.2.1.
* `udp_attacks` - Attacks utilizing the udp protocol operations. Field introduced in 18.2.1.
* `uuid` - The uuid of the security policy. Field introduced in 18.2.1.

