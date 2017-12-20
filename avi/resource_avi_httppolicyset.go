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

func ResourceHTTPPolicySetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_config_cksum": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"http_request_policy": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHTTPRequestPolicySchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"http_response_policy": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHTTPResponsePolicySchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"http_security_policy": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHTTPSecurityPolicySchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"is_internal_policy": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
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

func resourceAviHTTPPolicySet() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviHTTPPolicySetCreate,
		Read:   ResourceAviHTTPPolicySetRead,
		Update: resourceAviHTTPPolicySetUpdate,
		Delete: resourceAviHTTPPolicySetDelete,
		Schema: ResourceHTTPPolicySetSchema(),
	}
}

func ResourceAviHTTPPolicySetRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHTTPPolicySetSchema()
	log.Printf("[INFO] ResourceAviHTTPPolicySetRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/httppolicyset/" + uuid.(string)
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
	log.Printf("ResourceAviHTTPPolicySetRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviHTTPPolicySetRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviHTTPPolicySetRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviHTTPPolicySetRead Updated %v\n", d)
	return nil
}

func resourceAviHTTPPolicySetCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHTTPPolicySetSchema()
	err := ApiCreateOrUpdate(d, meta, "httppolicyset", s)
	log.Printf("[DEBUG] created object %v: %v", "httppolicyset", d)
	if err == nil {
		err = ResourceAviHTTPPolicySetRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "httppolicyset", d)
	return err
}

func resourceAviHTTPPolicySetUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHTTPPolicySetSchema()
	err := ApiCreateOrUpdate(d, meta, "httppolicyset", s)
	if err == nil {
		err = ResourceAviHTTPPolicySetRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "httppolicyset", d)
	return err
}

func resourceAviHTTPPolicySetDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "httppolicyset"
	log.Println("[INFO] ResourceAviHTTPPolicySetRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviHTTPPolicySetDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
