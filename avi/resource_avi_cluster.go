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

func ResourceClusterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"nodes": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceClusterNodeSchema(),
		},
		"rejoin_nodes_automatically": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
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
		"virtual_ip": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceIpAddrSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
	}
}

func resourceAviCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviClusterCreate,
		Read:   ResourceAviClusterRead,
		Update: resourceAviClusterUpdate,
		Delete: resourceAviClusterDelete,
		Schema: ResourceClusterSchema(),
	}
}

func ResourceAviClusterRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClusterSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/cluster/" + uuid.(string)
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

func resourceAviClusterCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClusterSchema()
	err := ApiCreateOrUpdate(d, meta, "cluster", s)
	if err == nil {
		err = ResourceAviClusterRead(d, meta)
	}
	return err
}

func resourceAviClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClusterSchema()
	err := ApiCreateOrUpdate(d, meta, "cluster", s)
	if err == nil {
		err = ResourceAviClusterRead(d, meta)
	}
	return err
}

func resourceAviClusterDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "cluster"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviClusterDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
