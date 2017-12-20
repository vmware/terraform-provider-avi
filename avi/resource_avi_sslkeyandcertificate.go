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

func ResourceSSLKeyAndCertificateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ca_certs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCertificateAuthoritySchema(),
		},
		"certificate": &schema.Schema{
			Type:     schema.TypeSet,
			Required: true, Elem: ResourceSSLCertificateSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"certificate_management_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"dynamic_params": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomParamsSchema(),
		},
		"enckey_base64": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"enckey_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"hardwaresecuritymodulegroup_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"key": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"key_params": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceSSLKeyParamsSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"status": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SSL_CERTIFICATE_FINISHED",
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"type": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SSL_CERTIFICATE_TYPE_VIRTUALSERVICE",
		},
		"uuid": &schema.Schema{
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
	}
}

func ResourceAviSSLKeyAndCertificateRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSLKeyAndCertificateSchema()
	log.Printf("[INFO] ResourceAviSSLKeyAndCertificateRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/sslkeyandcertificate/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	// no need to set the ID
	log.Printf("ResourceAviSSLKeyAndCertificateRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviSSLKeyAndCertificateRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviSSLKeyAndCertificateRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviSSLKeyAndCertificateRead Updated %v\n", d)
	return nil
}

func resourceAviSSLKeyAndCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSLKeyAndCertificateSchema()
	err := ApiCreateOrUpdate(d, meta, "sslkeyandcertificate", s)
	log.Printf("[DEBUG] created object %v: %v", "sslkeyandcertificate", d)
	if err == nil {
		err = ResourceAviSSLKeyAndCertificateRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "sslkeyandcertificate", d)
	return err
}

func resourceAviSSLKeyAndCertificateUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSLKeyAndCertificateSchema()
	err := ApiCreateOrUpdate(d, meta, "sslkeyandcertificate", s)
	if err == nil {
		err = ResourceAviSSLKeyAndCertificateRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "sslkeyandcertificate", d)
	return err
}

func resourceAviSSLKeyAndCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "sslkeyandcertificate"
	log.Println("[INFO] ResourceAviSSLKeyAndCertificateRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviSSLKeyAndCertificateDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
