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

func ResourceCustomIpamDnsProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"script_params": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomParamsSchema(),
		},
		"script_uri": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
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

func resourceAviCustomIpamDnsProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCustomIpamDnsProfileCreate,
		Read:   ResourceAviCustomIpamDnsProfileRead,
		Update: resourceAviCustomIpamDnsProfileUpdate,
		Delete: resourceAviCustomIpamDnsProfileDelete,
		Schema: ResourceCustomIpamDnsProfileSchema(),
	}
}

func ResourceAviCustomIpamDnsProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCustomIpamDnsProfileSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/customipamdnsprofile/" + uuid.(string)
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

func resourceAviCustomIpamDnsProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCustomIpamDnsProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "customipamdnsprofile", s)
	if err == nil {
		err = ResourceAviCustomIpamDnsProfileRead(d, meta)
	}
	return err
}

func resourceAviCustomIpamDnsProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCustomIpamDnsProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "customipamdnsprofile", s)
	if err == nil {
		err = ResourceAviCustomIpamDnsProfileRead(d, meta)
	}
	return err
}

func resourceAviCustomIpamDnsProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "customipamdnsprofile"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviCustomIpamDnsProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
