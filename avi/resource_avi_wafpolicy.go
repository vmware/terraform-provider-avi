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

func ResourceWafPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"crs_groups": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafRuleGroupSchema(),
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"mode": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "WAF_MODE_DETECTION_ONLY",
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"paranoia_level": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "WAF_PARANOIA_LEVEL_LOW",
		},
		"post_crs_groups": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafRuleGroupSchema(),
		},
		"pre_crs_groups": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafRuleGroupSchema(),
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
		"waf_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func resourceAviWafPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWafPolicyCreate,
		Read:   ResourceAviWafPolicyRead,
		Update: resourceAviWafPolicyUpdate,
		Delete: resourceAviWafPolicyDelete,
		Schema: ResourceWafPolicySchema(),
	}
}

func ResourceAviWafPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicySchema()
	log.Printf("[INFO] ResourceAviWafPolicyRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/wafpolicy/" + uuid.(string)
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
	log.Printf("ResourceAviWafPolicyRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviWafPolicyRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviWafPolicyRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviWafPolicyRead Updated %v\n", d)
	return nil
}

func resourceAviWafPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicySchema()
	err := ApiCreateOrUpdate(d, meta, "wafpolicy", s)
	log.Printf("[DEBUG] created object %v: %v", "wafpolicy", d)
	if err == nil {
		err = ResourceAviWafPolicyRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "wafpolicy", d)
	return err
}

func resourceAviWafPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicySchema()
	err := ApiCreateOrUpdate(d, meta, "wafpolicy", s)
	if err == nil {
		err = ResourceAviWafPolicyRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "wafpolicy", d)
	return err
}

func resourceAviWafPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "wafpolicy"
	log.Println("[INFO] ResourceAviWafPolicyRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviWafPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
