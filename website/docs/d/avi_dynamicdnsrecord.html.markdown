<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_dynamicdnsrecord"
sidebar_current: "docs-avi-datasource-dynamicdnsrecord"
description: |-
  Get information of Avi DynamicDnsRecord.
---

# avi_dynamicdnsrecord

This data source is used to to get avi_dynamicdnsrecord objects.

## Example Usage

```hcl
data "avi_dynamicdnsrecord" "foo_dynamicdnsrecord" {
    uuid = "dynamicdnsrecord-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search DynamicDnsRecord by name.
* `uuid` - (Optional) Search DynamicDnsRecord by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `algorithm` - Specifies the algorithm to pick the ip address(es) to be returned,when multiple entries are configured. This does not apply if num_records_in_response is 0. Default is round-robin. Enum options - DNS_RECORD_RESPONSE_ROUND_ROBIN, DNS_RECORD_RESPONSE_CONSISTENT_HASH. Field introduced in 20.1.3.
* `cname` - Canonical name in cname record. Field introduced in 20.1.3.
* `delegated` - Configured fqdns are delegated domains (i.e. They represent a zone cut). Field introduced in 20.1.3.
* `description` - Details of dns record. Field introduced in 20.1.3.
* `dns_vs_uuid` - Uuid of the dns vs. Field introduced in 20.1.3.
* `fqdn` - Fully qualified domain name. Field introduced in 20.1.3.
* `ip6_address` - Ipv6 address in aaaa record. Field introduced in 20.1.3. Maximum of 4 items allowed.
* `ip_address` - Ip address in a record. Field introduced in 20.1.3. Maximum of 4 items allowed.
* `metadata` - Internal metadata for the dns record. Field introduced in 20.1.3.
* `mx_records` - Mx record. Field introduced in 20.1.3. Maximum of 4 items allowed.
* `name` - Dynamicdnsrecord name, needed for a top level uuid protobuf, for display in shell. Field introduced in 20.1.3.
* `ns` - Name server information in ns record. Field introduced in 20.1.3. Maximum of 13 items allowed.
* `num_records_in_response` - Specifies the number of records returned by the dns service.enter 0 to return all records. Default is 0. Allowed values are 0-20. Special values are 0- 'return all records'. Field introduced in 20.1.3.
* `service_locators` - Service locator info in srv record. Field introduced in 20.1.3. Maximum of 4 items allowed.
* `tenant_ref` - Tenant_uuid from dns vs's tenant_uuid. It is a reference to an object of type tenant. Field introduced in 20.1.3.
* `ttl` - Time to live for this dns record. Field introduced in 20.1.3.
* `txt_records` - Text record. Field introduced in 20.1.3. Maximum of 4 items allowed.
* `type` - Dns record type. Enum options - DNS_RECORD_OTHER, DNS_RECORD_A, DNS_RECORD_NS, DNS_RECORD_CNAME, DNS_RECORD_SOA, DNS_RECORD_PTR, DNS_RECORD_HINFO, DNS_RECORD_MX, DNS_RECORD_TXT, DNS_RECORD_RP, DNS_RECORD_DNSKEY, DNS_RECORD_AAAA, DNS_RECORD_SRV, DNS_RECORD_OPT, DNS_RECORD_RRSIG, DNS_RECORD_AXFR, DNS_RECORD_ANY. Field introduced in 20.1.3.
* `uuid` - Uuid of the dns record. Field introduced in 20.1.3.
* `wildcard_match` - Enable wild-card match of fqdn  if an exact match is not found in the dns table, the longest match is chosen by wild-carding the fqdn in the dns request. Default is false. Field introduced in 20.1.3.

