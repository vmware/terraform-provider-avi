/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
)

func ResourceSSLProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"accepted_ciphers": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "AES:3DES:RC4",
		},
		"accepted_versions": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceSSLVersionSchema(),
		},
		"cipher_enums": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"dhparam": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"enable_ssl_session_reuse": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"prefer_client_cipher_ordering": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"send_close_notify": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"ssl_rating": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceSSLRatingSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"ssl_session_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  86400,
		},
		"tags": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceTagSchema(),
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"uuid": &schema.Schema{
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
	}
}

func ResourceAviSSLProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSLProfileSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/sslprofile/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	if _, err := ApiDataToSchema(obj, d, s); err == nil {
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	return nil
}

func resourceAviSSLProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSLProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "sslprofile", s)
	if err == nil {
		err = ResourceAviSSLProfileRead(d, meta)
	}
	return err
}

func resourceAviSSLProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSLProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "sslprofile", s)
	if err == nil {
		err = ResourceAviSSLProfileRead(d, meta)
	}
	return err
}

func resourceAviSSLProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "sslprofile"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviSSLProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
