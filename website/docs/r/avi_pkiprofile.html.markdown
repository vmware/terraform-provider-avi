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

* `name` - (Required) Name of the pki profile.
* `ca_certs` - (Optional) List of certificate authorities (root and intermediate) trusted that is used for certificate validation.
* `created_by` - (Optional) Creator name.
* `crl_check` - (Optional) When enabled, avi will verify via crl checks that certificates in the trust chain have not been revoked.
* `crls` - (Optional) Certificate revocation lists.
* `ignore_peer_chain` - (Optional) When enabled, avi will not trust intermediate and root certs presented by a client. Instead, only the chain certs configured in the certificate authority section will be used to verify trust of the client's cert.
* `is_federated` - (Optional) This field describes the object's replication scope. If the field is set to false, then the object is visible within the controller-cluster and its associated service-engines. If the field is set to true, then the object is replicated across the federation. Field introduced in 17.1.3.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `validate_only_leaf_crl` - (Optional) When enabled, avi will only validate the revocation status of the leaf certificate using crl. To enable validation for the entire chain, disable this option and provide all the relevant crls.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

