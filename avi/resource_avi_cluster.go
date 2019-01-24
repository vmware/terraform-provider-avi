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
	"time"
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
			Computed: true,
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
		Importer: &schema.ResourceImporter{
			State: ResourceClusterImporter,
		},
	}
}

func ResourceClusterImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceClusterSchema()
	return ResourceImporter(d, m, "cluster", s)
}

func ResourceAviClusterRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClusterSchema()
	err := ApiRead(d, meta, "cluster", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviClusterCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClusterSchema()
	err := ApiCreateOrUpdate(d, meta, "cluster", s)
	// Added wait for cluster initialization process as cluster initialization starts after few seconds.
	// This is necessary to store correct state of initialized cluster.
	time.Sleep(90 * time.Second)
	if err == nil {
		err = ResourceAviClusterRead(d, meta)
	}
	return err
}

func resourceAviClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClusterSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "cluster", s)
	// Added wait for cluster initialization process as cluster initialization starts after few seconds.
	// This is necessary to store correct state of initialized cluster.
	time.Sleep(90 * time.Second)
	if err == nil {
		err = ResourceAviClusterRead(d, meta)
	}
	return err
}

func resourceAviClusterDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "cluster"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviClusterDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
