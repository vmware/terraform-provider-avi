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

func ResourceWafPolicyPSMGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"enable": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"hit_action": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "WAF_ACTION_NO_OP",
		},
		"is_learning_group": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"locations": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafPSMLocationSchema(),
		},
		"miss_action": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "WAF_ACTION_NO_OP",
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviWafPolicyPSMGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWafPolicyPSMGroupCreate,
		Read:   ResourceAviWafPolicyPSMGroupRead,
		Update: resourceAviWafPolicyPSMGroupUpdate,
		Delete: resourceAviWafPolicyPSMGroupDelete,
		Schema: ResourceWafPolicyPSMGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceWafPolicyPSMGroupImporter,
		},
	}
}

func ResourceWafPolicyPSMGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceWafPolicyPSMGroupSchema()
	return ResourceImporter(d, m, "wafpolicypsmgroup", s)
}

func ResourceAviWafPolicyPSMGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicyPSMGroupSchema()
	err := ApiRead(d, meta, "wafpolicypsmgroup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWafPolicyPSMGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicyPSMGroupSchema()
	err := ApiCreateOrUpdate(d, meta, "wafpolicypsmgroup", s)
	if err == nil {
		err = ResourceAviWafPolicyPSMGroupRead(d, meta)
	}
	return err
}

func resourceAviWafPolicyPSMGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicyPSMGroupSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "wafpolicypsmgroup", s)
	if err == nil {
		err = ResourceAviWafPolicyPSMGroupRead(d, meta)
	}
	return err
}

func resourceAviWafPolicyPSMGroupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "wafpolicypsmgroup"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviWafPolicyPSMGroupDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
