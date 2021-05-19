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

func ResourceSSLProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"accepted_ciphers": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "AES:3DES:RC4",
		},
		"accepted_versions": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     ResourceSSLVersionSchema(),
		},
		"cipher_enums": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"ciphersuites": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "TLS_AES_256_GCM_SHA384:TLS_CHACHA20_POLY1305_SHA256:TLS_AES_128_GCM_SHA256",
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"dhparam": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enable_early_data": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"enable_ssl_session_reuse": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
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
		"prefer_client_cipher_ordering": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"send_close_notify": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"ssl_rating": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSSLRatingSchema(),
		},
		"ssl_session_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  86400,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceTagSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SSL_PROFILE_TYPE_APPLICATION",
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviSSLProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSSLProfileCreate,
		Read:   ResourceAviSSLProfileRead,
		Update: resourceAviSSLProfileUpdate,
		Delete: resourceAviSSLProfileDelete,
		Schema: ResourceSSLProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSSLProfileImporter,
		},
	}
}

func ResourceSSLProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSSLProfileSchema()
	return ResourceImporter(d, m, "sslprofile", s)
}

func ResourceAviSSLProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSLProfileSchema()
	err := APIRead(d, meta, "sslprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSSLProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSLProfileSchema()
	err := APICreateOrUpdate(d, meta, "sslprofile", s)
	if err == nil {
		err = ResourceAviSSLProfileRead(d, meta)
	}
	return err
}

func resourceAviSSLProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSLProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "sslprofile", s)
	if err == nil {
		err = ResourceAviSSLProfileRead(d, meta)
	}
	return err
}

func resourceAviSSLProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "sslprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviSSLProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
