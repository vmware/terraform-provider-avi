// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceSSLKeyAndCertificateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ca_certs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCertificateAuthoritySchema(),
		},
		"certificate": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceSSLCertificateSchema(),
		},
		"certificate_base64": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"certificate_management_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"dynamic_params": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomParamsSchema(),
		},
		"enable_ocsp_stapling": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"enckey_base64": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enckey_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"format": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SSL_PEM",
		},
		"hardwaresecuritymodulegroup_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"import_key_to_hsm": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"is_federated": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"key": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Sensitive:        true,
			DiffSuppressFunc: suppressSensitiveFieldDiffs,
		},
		"key_base64": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"key_params": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSSLKeyParamsSchema(),
		},
		"key_passphrase": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Sensitive:        true,
			DiffSuppressFunc: suppressSensitiveFieldDiffs,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"ocsp_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceOCSPConfigSchema(),
		},
		"ocsp_error_status": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ocsp_responder_url_list_from_certs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"ocsp_response_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceOCSPResponseInfoSchema(),
		},
		"status": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SSL_CERTIFICATE_FINISHED",
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviSSLKeyAndCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSSLKeyAndCertificateCreate,
		Read:   ResourceAviSSLKeyAndCertificateRead,
		Update: resourceAviSSLKeyAndCertificateUpdate,
		Delete: resourceAviSSLKeyAndCertificateDelete,
		Schema: ResourceSSLKeyAndCertificateSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSSLKeyAndCertificateImporter,
		},
	}
}

func ResourceSSLKeyAndCertificateImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSSLKeyAndCertificateSchema()
	return ResourceImporter(d, m, "sslkeyandcertificate", s)
}

func ResourceAviSSLKeyAndCertificateRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSLKeyAndCertificateSchema()
	err := APIRead(d, meta, "sslkeyandcertificate", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSSLKeyAndCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSLKeyAndCertificateSchema()
	err := APICreateOrUpdate(d, meta, "sslkeyandcertificate", s)
	if err == nil {
		err = ResourceAviSSLKeyAndCertificateRead(d, meta)
	}
	return err
}

func resourceAviSSLKeyAndCertificateUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSLKeyAndCertificateSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "sslkeyandcertificate", s)
	if err == nil {
		err = ResourceAviSSLKeyAndCertificateRead(d, meta)
	}
	return err
}

func resourceAviSSLKeyAndCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "sslkeyandcertificate")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
