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

* `accepted_versions` - (Required) Set of versions accepted by the server. Minimum of 1 items required.
* `name` - (Required) Name of the object.
* `accepted_ciphers` - (Optional) Ciphers suites represented as defined by https //www.openssl.org/docs/apps/ciphers.html.
* `cipher_enums` - (Optional) Enum options - tls_ecdhe_ecdsa_with_aes_128_gcm_sha256, tls_ecdhe_ecdsa_with_aes_256_gcm_sha384, tls_ecdhe_rsa_with_aes_128_gcm_sha256, tls_ecdhe_rsa_with_aes_256_gcm_sha384, tls_ecdhe_ecdsa_with_aes_128_cbc_sha256, tls_ecdhe_ecdsa_with_aes_256_cbc_sha384, tls_ecdhe_rsa_with_aes_128_cbc_sha256, tls_ecdhe_rsa_with_aes_256_cbc_sha384, tls_rsa_with_aes_128_gcm_sha256, tls_rsa_with_aes_256_gcm_sha384, tls_rsa_with_aes_128_cbc_sha256, tls_rsa_with_aes_256_cbc_sha256, tls_ecdhe_ecdsa_with_aes_128_cbc_sha, tls_ecdhe_ecdsa_with_aes_256_cbc_sha, tls_ecdhe_rsa_with_aes_128_cbc_sha, tls_ecdhe_rsa_with_aes_256_cbc_sha, tls_rsa_with_aes_128_cbc_sha, tls_rsa_with_aes_256_cbc_sha, tls_rsa_with_3des_ede_cbc_sha, tls_aes_256_gcm_sha384... Allowed in basic(allowed values- tls_ecdhe_ecdsa_with_aes_128_gcm_sha256,tls_ecdhe_ecdsa_with_aes_256_gcm_sha384,tls_ecdhe_rsa_with_aes_128_gcm_sha256,tls_ecdhe_rsa_with_aes_256_gcm_sha384,tls_ecdhe_ecdsa_with_aes_128_cbc_sha256,tls_ecdhe_ecdsa_with_aes_256_cbc_sha384,tls_ecdhe_rsa_with_aes_128_cbc_sha256,tls_ecdhe_rsa_with_aes_256_cbc_sha384,tls_rsa_with_aes_128_gcm_sha256,tls_rsa_with_aes_256_gcm_sha384,tls_rsa_with_aes_128_cbc_sha256,tls_rsa_with_aes_256_cbc_sha256,tls_ecdhe_ecdsa_with_aes_128_cbc_sha,tls_ecdhe_ecdsa_with_aes_256_cbc_sha,tls_ecdhe_rsa_with_aes_128_cbc_sha,tls_ecdhe_rsa_with_aes_256_cbc_sha,tls_rsa_with_aes_128_cbc_sha,tls_rsa_with_aes_256_cbc_sha,tls_rsa_with_3des_ede_cbc_sha) edition, essentials(allowed values- tls_ecdhe_ecdsa_with_aes_128_gcm_sha256,tls_ecdhe_ecdsa_with_aes_256_gcm_sha384,tls_ecdhe_rsa_with_aes_128_gcm_sha256,tls_ecdhe_rsa_with_aes_256_gcm_sha384,tls_ecdhe_ecdsa_with_aes_128_cbc_sha256,tls_ecdhe_ecdsa_with_aes_256_cbc_sha384,tls_ecdhe_rsa_with_aes_128_cbc_sha256,tls_ecdhe_rsa_with_aes_256_cbc_sha384,tls_rsa_with_aes_128_gcm_sha256,tls_rsa_with_aes_256_gcm_sha384,tls_rsa_with_aes_128_cbc_sha256,tls_rsa_with_aes_256_cbc_sha256,tls_ecdhe_ecdsa_with_aes_128_cbc_sha,tls_ecdhe_ecdsa_with_aes_256_cbc_sha,tls_ecdhe_rsa_with_aes_128_cbc_sha,tls_ecdhe_rsa_with_aes_256_cbc_sha,tls_rsa_with_aes_128_cbc_sha,tls_rsa_with_aes_256_cbc_sha,tls_rsa_with_3des_ede_cbc_sha) edition, enterprise edition.
* `ciphersuites` - (Optional) Tls 1.3 ciphers suites represented as defined by u(https //www.openssl.org/docs/manmaster/man1/ciphers.html). Field introduced in 18.2.6. Allowed in basic edition, essentials edition, enterprise edition. Special default for basic edition is tls_aes_256_gcm_sha384-tls_aes_128_gcm_sha256, essentials edition is tls_aes_256_gcm_sha384-tls_aes_128_gcm_sha256, enterprise is tls_aes_256_gcm_sha384-tls_chacha20_poly1305_sha256-tls_aes_128_gcm_sha256.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `description` - (Optional) User defined description for the object.
* `dhparam` - (Optional) Dh parameters used in ssl. At this time, it is not configurable and is set to 2048 bits.
* `ec_named_curve` - (Optional) Elliptic curve cryptography namedcurves (tls supported groups)represented as defined by rfc 8422-section 5.1.1 andhttps //www.openssl.org/docs/man1.1.0/man3/ssl_ctx_set1_curves.html. Field introduced in 21.1.1.
* `enable_early_data` - (Optional) Enable early data processing for tls1.3 connections. Field introduced in 18.2.6. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `enable_ssl_session_reuse` - (Optional) Enable ssl session re-use.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5.
* `prefer_client_cipher_ordering` - (Optional) Prefer the ssl cipher ordering presented by the client during the ssl handshake over the one specified in the ssl profile.
* `send_close_notify` - (Optional) Send 'close notify' alert message for a clean shutdown of the ssl connection.
* `signature_algorithm` - (Optional) Signature algorithms represented as defined by rfc5246-section 7.4.1.4.1 andhttps //www.openssl.org/docs/man1.1.0/man3/ssl_ctx_set1_client_sigalgs_list.html. Field introduced in 21.1.1.
* `ssl_rating` - (Optional) Dict settings for sslprofile.
* `ssl_session_timeout` - (Optional) The amount of time in seconds before an ssl session expires. Unit is sec.
* `tags` - (Optional) List of list.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `type` - (Optional) Ssl profile type. Enum options - SSL_PROFILE_TYPE_APPLICATION, SSL_PROFILE_TYPE_SYSTEM. Field introduced in 17.2.8.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

