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

func ResourceApplicationProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"dns_service_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceDnsServiceApplicationProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"dos_rl_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceDosRateLimitProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"http_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHTTPApplicationProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"preserve_client_ip": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"tcp_app_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceTCPApplicationProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"type": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviApplicationProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviApplicationProfileCreate,
		Read:   ResourceAviApplicationProfileRead,
		Update: resourceAviApplicationProfileUpdate,
		Delete: resourceAviApplicationProfileDelete,
		Schema: ResourceApplicationProfileSchema(),
	}
}

func ResourceAviApplicationProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationProfileSchema()
	err := ApiRead(d, meta, "applicationprofile", s)
	return err
}

func resourceAviApplicationProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "applicationprofile", s)
	if err == nil {
		err = ResourceAviApplicationProfileRead(d, meta)
	}
	return err
}

func resourceAviApplicationProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "applicationprofile", s)
	if err == nil {
		err = ResourceAviApplicationProfileRead(d, meta)
	}
	return err
}

func resourceAviApplicationProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "applicationprofile"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviApplicationProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
