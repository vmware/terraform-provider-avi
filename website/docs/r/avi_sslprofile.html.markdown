<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_sslprofile"
sidebar_current: "docs-avi-resource-sslprofile"
description: |-
  Creates and manages Avi SSLProfile.
---

# avi_sslprofile

The SSLProfile resource allows the creation and management of Avi SSLProfile

## Example Usage

```hcl
resource "avi_sslprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `accepted_versions` - (Required) Set of versions accepted by the server. Minimum of 1 items required. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `name` - (Required) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `accepted_ciphers` - (Optional) Ciphers suites represented as defined by https //www.openssl.org/docs/apps/ciphers.html. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `cipher_enums` - (Optional) Enum options - tls_ecdhe_ecdsa_with_aes_128_gcm_sha256, tls_ecdhe_ecdsa_with_aes_256_gcm_sha384, tls_ecdhe_rsa_with_aes_128_gcm_sha256, tls_ecdhe_rsa_with_aes_256_gcm_sha384, tls_ecdhe_ecdsa_with_aes_128_cbc_sha256, tls_ecdhe_ecdsa_with_aes_256_cbc_sha384, tls_ecdhe_rsa_with_aes_128_cbc_sha256, tls_ecdhe_rsa_with_aes_256_cbc_sha384, tls_rsa_with_aes_128_gcm_sha256, tls_rsa_with_aes_256_gcm_sha384, tls_rsa_with_aes_128_cbc_sha256, tls_rsa_with_aes_256_cbc_sha256, tls_ecdhe_ecdsa_with_aes_128_cbc_sha, tls_ecdhe_ecdsa_with_aes_256_cbc_sha, tls_ecdhe_rsa_with_aes_128_cbc_sha, tls_ecdhe_rsa_with_aes_256_cbc_sha, tls_rsa_with_aes_128_cbc_sha, tls_rsa_with_aes_256_cbc_sha, tls_rsa_with_3des_ede_cbc_sha, tls_aes_256_gcm_sha384... Allowed in enterprise edition with any value, essentials edition(allowed values- tls_ecdhe_ecdsa_with_aes_128_gcm_sha256,tls_ecdhe_ecdsa_with_aes_256_gcm_sha384,tls_ecdhe_rsa_with_aes_128_gcm_sha256,tls_ecdhe_rsa_with_aes_256_gcm_sha384,tls_ecdhe_ecdsa_with_aes_128_cbc_sha256,tls_ecdhe_ecdsa_with_aes_256_cbc_sha384,tls_ecdhe_rsa_with_aes_128_cbc_sha256,tls_ecdhe_rsa_with_aes_256_cbc_sha384,tls_rsa_with_aes_128_gcm_sha256,tls_rsa_with_aes_256_gcm_sha384,tls_rsa_with_aes_128_cbc_sha256,tls_rsa_with_aes_256_cbc_sha256,tls_ecdhe_ecdsa_with_aes_128_cbc_sha,tls_ecdhe_ecdsa_with_aes_256_cbc_sha,tls_ecdhe_rsa_with_aes_128_cbc_sha,tls_ecdhe_rsa_with_aes_256_cbc_sha,tls_rsa_with_aes_128_cbc_sha,tls_rsa_with_aes_256_cbc_sha,tls_rsa_with_3des_ede_cbc_sha), basic edition(allowed values- tls_ecdhe_ecdsa_with_aes_128_gcm_sha256,tls_ecdhe_ecdsa_with_aes_256_gcm_sha384,tls_ecdhe_rsa_with_aes_128_gcm_sha256,tls_ecdhe_rsa_with_aes_256_gcm_sha384,tls_ecdhe_ecdsa_with_aes_128_cbc_sha256,tls_ecdhe_ecdsa_with_aes_256_cbc_sha384,tls_ecdhe_rsa_with_aes_128_cbc_sha256,tls_ecdhe_rsa_with_aes_256_cbc_sha384,tls_rsa_with_aes_128_gcm_sha256,tls_rsa_with_aes_256_gcm_sha384,tls_rsa_with_aes_128_cbc_sha256,tls_rsa_with_aes_256_cbc_sha256,tls_ecdhe_ecdsa_with_aes_128_cbc_sha,tls_ecdhe_ecdsa_with_aes_256_cbc_sha,tls_ecdhe_rsa_with_aes_128_cbc_sha,tls_ecdhe_rsa_with_aes_256_cbc_sha,tls_rsa_with_aes_128_cbc_sha,tls_rsa_with_aes_256_cbc_sha,tls_rsa_with_3des_ede_cbc_sha), enterprise with cloud services edition.
* `ciphersuites` - (Optional) Tls 1.3 ciphers suites represented as defined by u(https //www.openssl.org/docs/manmaster/man1/ciphers.html). Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition. Special default for essentials edition is tls_aes_256_gcm_sha384-tls_aes_128_gcm_sha256, basic edition is tls_aes_256_gcm_sha384-tls_aes_128_gcm_sha256, enterprise is tls_aes_256_gcm_sha384-tls_chacha20_poly1305_sha256-tls_aes_128_gcm_sha256.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `dhparam` - (Optional) Dh parameters used in ssl. At this time, it is not configurable and is set to 2048 bits. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `ec_named_curve` - (Optional) Elliptic curve cryptography namedcurves (tls supported groups)represented as defined by rfc 8422-section 5.1.1 andhttps //www.openssl.org/docs/man1.1.0/man3/ssl_ctx_set1_curves.html. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `enable_early_data` - (Optional) Enable early data processing for tls1.3 connections. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `enable_ssl_session_reuse` - (Optional) Enable ssl session re-use. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `prefer_client_cipher_ordering` - (Optional) Prefer the ssl cipher ordering presented by the client during the ssl handshake over the one specified in the ssl profile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `send_close_notify` - (Optional) Send 'close notify' alert message for a clean shutdown of the ssl connection. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `signature_algorithm` - (Optional) Signature algorithms represented as defined by rfc5246-section 7.4.1.4.1 andhttps //www.openssl.org/docs/man1.1.0/man3/ssl_ctx_set1_client_sigalgs_list.html. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `ssl_rating` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `ssl_session_timeout` - (Optional) The amount of time in seconds before an ssl session expires. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tags` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `type` - (Optional) Ssl profile type. Enum options - SSL_PROFILE_TYPE_APPLICATION, SSL_PROFILE_TYPE_SYSTEM. Field introduced in 17.2.8. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

