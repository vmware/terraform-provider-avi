---
layout: "avi"
page_title: "AVI: avi_sslkeyandcertificate"
sidebar_current: "docs-avi-datasource-sslkeyandcertificate"
description: |-
  Get information of Avi SSLKeyAndCertificate.
---

# avi_sslkeyandcertificate

This data source is used to to get avi_sslkeyandcertificate objects.

## Example Usage

```hcl
data "avi_sslkeyandcertificate" "foo_sslkeyandcertificate" {
    uuid = "sslkeyandcertificate-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SSLKeyAndCertificate by name.
* `uuid` - (Optional) Search SSLKeyAndCertificate by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `ca_certs` - Ca certificates in certificate chain.
* `certificate` - Dict settings for sslkeyandcertificate.
* `certificate_base64` - States if the certificate is base64 encoded.
* `certificate_management_profile_ref` - It is a reference to an object of type certificatemanagementprofile.
* `created_by` - Creator name.
* `dynamic_params` - Dynamic parameters needed for certificate management profile.
* `enable_ocsp_stapling` - Enables ocsp stapling.
* `enckey_base64` - Encrypted private key corresponding to the private key (e.g.
* `enckey_name` - Name of the encrypted private key (e.g.
* `format` - Format of the key/certificate file.
* `hardwaresecuritymodulegroup_ref` - It is a reference to an object of type hardwaresecuritymodulegroup.
* `key` - Private key.
* `key_base64` - States if the private key is base64 encoded.
* `key_params` - Dict settings for sslkeyandcertificate.
* `key_passphrase` - Passphrase used to encrypt the private key.
* `name` - Name of the object.
* `ocsp_config` - Configuration related to ocsp.
* `ocsp_error_status` - Error reported during ocsp status query.
* `ocsp_response_info` - Information related to ocsp response.
* `status` - Enum options - ssl_certificate_finished, ssl_certificate_pending.
* `tenant_ref` - It is a reference to an object of type tenant.
* `type` - Enum options - ssl_certificate_type_virtualservice, ssl_certificate_type_system, ssl_certificate_type_ca.
* `uuid` - Unique object identifier of the object.

