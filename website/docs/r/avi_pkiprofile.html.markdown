<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_pkiprofile"
sidebar_current: "docs-avi-resource-pkiprofile"
description: |-
  Creates and manages Avi PKIProfile.
---

# avi_pkiprofile

The PKIProfile resource allows the creation and management of Avi PKIProfile

## Example Usage

```hcl
resource "avi_pkiprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the pki profile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `allow_pki_errors` - (Optional) Exempt errors during certificate verification. Enum options - ALLOW_EXPIRED_CRL, ALLOW_ALL_ERRORS. Field introduced in 30.1.1. Maximum of 1 items allowed. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `ca_certs` - (Optional) List of certificate authorities (root and intermediate) trusted that is used for certificate validation. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `created_by` - (Optional) Creator name. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `crl_check` - (Optional) When enabled, avi will verify via crl checks that certificates in the trust chain have not been revoked. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `crls` - (Optional) Certificate revocation lists. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `ignore_peer_chain` - (Optional) When enabled, avi will not trust intermediate and root certs presented by a client. Instead, only the chain certs configured in the certificate authority section will be used to verify trust of the client's cert. Allowed in enterprise edition with any value, essentials edition(allowed values- true), basic edition(allowed values- true), enterprise with cloud services edition. Special default for essentials edition is true, basic edition is true, enterprise is false.
* `is_federated` - (Optional) This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster and its associated service-engines. If the field is set to true, then the object is replicated across the federation. Field introduced in 17.1.3. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `validate_only_leaf_crl` - (Optional) When enabled, avi will only validate the revocation status of the leaf certificate using crl. To enable validation for the entire chain, disable this option and provide all the relevant crls. Allowed in enterprise edition with any value, essentials edition(allowed values- true), basic edition(allowed values- true), enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

