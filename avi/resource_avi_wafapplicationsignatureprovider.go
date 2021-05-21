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

func ResourceWafApplicationSignatureProviderSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"service_status": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceWebApplicationSignatureServiceStatusSchema(),
		},
		"tenant_ref": {
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

func resourceAviWafApplicationSignatureProvider() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWafApplicationSignatureProviderCreate,
		Read:   ResourceAviWafApplicationSignatureProviderRead,
		Update: resourceAviWafApplicationSignatureProviderUpdate,
		Delete: resourceAviWafApplicationSignatureProviderDelete,
		Schema: ResourceWafApplicationSignatureProviderSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceWafApplicationSignatureProviderImporter,
		},
	}
}

func ResourceWafApplicationSignatureProviderImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceWafApplicationSignatureProviderSchema()
	return ResourceImporter(d, m, "wafapplicationsignatureprovider", s)
}

func ResourceAviWafApplicationSignatureProviderRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafApplicationSignatureProviderSchema()
	err := APIRead(d, meta, "wafapplicationsignatureprovider", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWafApplicationSignatureProviderCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafApplicationSignatureProviderSchema()
	err := APICreateOrUpdate(d, meta, "wafapplicationsignatureprovider", s)
	if err == nil {
		err = ResourceAviWafApplicationSignatureProviderRead(d, meta)
	}
	return err
}

func resourceAviWafApplicationSignatureProviderUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafApplicationSignatureProviderSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "wafapplicationsignatureprovider", s)
	if err == nil {
		err = ResourceAviWafApplicationSignatureProviderRead(d, meta)
	}
	return err
}

func resourceAviWafApplicationSignatureProviderDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "wafapplicationsignatureprovider"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviWafApplicationSignatureProviderDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
