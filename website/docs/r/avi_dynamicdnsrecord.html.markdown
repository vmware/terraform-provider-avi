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
page_title: "Avi: avi_dynamicdnsrecord"
sidebar_current: "docs-avi-resource-dynamicdnsrecord"
description: |-
  Creates and manages Avi DynamicDnsRecord.
---

# avi_dynamicdnsrecord

The DynamicDnsRecord resource allows the creation and management of Avi DynamicDnsRecord

## Example Usage

```hcl
resource "avi_dynamicdnsrecord" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `algorithm` - (Optional) Specifies the algorithm to pick the ip address(es) to be returned,when multiple entries are configured. This does not apply if num_records_in_response is 0. Default is round-robin. Enum options - DNS_RECORD_RESPONSE_ROUND_ROBIN, DNS_RECORD_RESPONSE_CONSISTENT_HASH. Field introduced in 20.1.3.
* `cname` - (Optional) Canonical name in cname record. Field introduced in 20.1.3.
* `delegated` - (Optional) Configured fqdns are delegated domains (i.e. They represent a zone cut). Field introduced in 20.1.3.
* `description` - (Optional) Details of dns record. Field introduced in 20.1.3.
* `dns_vs_uuid` - (Optional) Uuid of the dns vs. Field introduced in 20.1.3.
* `fqdn` - (Optional) Fully qualified domain name. Field introduced in 20.1.3.
* `ip6_address` - (Optional) Ipv6 address in aaaa record. Field introduced in 20.1.3. Maximum of 4 items allowed.
* `ip_address` - (Optional) Ip address in a record. Field introduced in 20.1.3. Maximum of 4 items allowed.
* `metadata` - (Optional) Internal metadata for the dns record. Field introduced in 20.1.3.
* `mx_records` - (Optional) Mx record. Field introduced in 20.1.3. Maximum of 4 items allowed.
* `name` - (Optional) Dynamicdnsrecord name, needed for a top level uuid protobuf, for display in shell. Field introduced in 20.1.3.
* `ns` - (Optional) Name server information in ns record. Field introduced in 20.1.3. Maximum of 13 items allowed.
* `num_records_in_response` - (Optional) Specifies the number of records returned by the dns service.enter 0 to return all records. Default is 0. Allowed values are 0-20. Special values are 0- 'return all records'. Field introduced in 20.1.3.
* `service_locators` - (Optional) Service locator info in srv record. Field introduced in 20.1.3. Maximum of 4 items allowed.
* `tenant_ref` - (Optional) Tenant_uuid from dns vs's tenant_uuid. It is a reference to an object of type tenant. Field introduced in 20.1.3.
* `ttl` - (Optional) Time to live for this dns record. Field introduced in 20.1.3.
* `txt_records` - (Optional) Text record. Field introduced in 20.1.3. Maximum of 4 items allowed.
* `type` - (Optional) Dns record type. Enum options - DNS_RECORD_OTHER, DNS_RECORD_A, DNS_RECORD_NS, DNS_RECORD_CNAME, DNS_RECORD_SOA, DNS_RECORD_PTR, DNS_RECORD_HINFO, DNS_RECORD_MX, DNS_RECORD_TXT, DNS_RECORD_RP, DNS_RECORD_DNSKEY, DNS_RECORD_AAAA, DNS_RECORD_SRV, DNS_RECORD_OPT, DNS_RECORD_RRSIG, DNS_RECORD_AXFR, DNS_RECORD_ANY. Field introduced in 20.1.3.
* `wildcard_match` - (Optional) Enable wild-card match of fqdn  if an exact match is not found in the dns table, the longest match is chosen by wild-carding the fqdn in the dns request. Default is false. Field introduced in 20.1.3.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the dns record. Field introduced in 20.1.3.

