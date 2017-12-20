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
	log.Printf("[INFO] ResourceAviApplicationProfileRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/applicationprofile/" + uuid.(string)
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
	log.Printf("ResourceAviApplicationProfileRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviApplicationProfileRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviApplicationProfileRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviApplicationProfileRead Updated %v\n", d)
	return nil
}

func resourceAviApplicationProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "applicationprofile", s)
	log.Printf("[DEBUG] created object %v: %v", "applicationprofile", d)
	if err == nil {
		err = ResourceAviApplicationProfileRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "applicationprofile", d)
	return err
}

func resourceAviApplicationProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "applicationprofile", s)
	if err == nil {
		err = ResourceAviApplicationProfileRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "applicationprofile", d)
	return err
}

func resourceAviApplicationProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "applicationprofile"
	log.Println("[INFO] ResourceAviApplicationProfileRead Avi Client")
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
