/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
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
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"certificate_management_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
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
		"key": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Sensitive:        true,
			DiffSuppressFunc: suppressSensitiveFieldDiffs,
		},
		"key_base64": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
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
	objType := "sslkeyandcertificate"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviSSLKeyAndCertificateDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
