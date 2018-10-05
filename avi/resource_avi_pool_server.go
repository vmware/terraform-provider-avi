/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/avinetworks/sdk/go/models"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
)

func ResourceAviPoolServerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"pool_ref": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"ip": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"port": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"type": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "V4",
		},
		"autoscaling_group_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"external_orchestration_id": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"external_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"hostname": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"location": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceGeoLocationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"nw_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"prst_hdr_val": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"rewrite_host_header": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"vm_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviServerCreateOrUpdate,
		Read:   ResourceAviServerRead,
		Update: resourceAviServerCreateOrUpdate,
		Delete: resourceAviServerDelete,
		Schema: ResourceAviPoolServerSchema(),
	}
}

func resourceAviServerCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AviClient)
	err, pUUID, poolObj, pserver := resourceAviServerReadApi(d, meta)
	//added check for err and poolObj.
	if err != nil || poolObj == nil {
		log.Printf("[ERROR] resourceAviServerCreateOrUpdate Error during fetching pool object using pool_ref %v", err)
		return err
	}
	if pserver == nil {
		// not found
		newServer := models.Server{}
		if port, ok := d.GetOk("port"); ok {
			newServer.Port = int32(port.(int))
		}
		pserver = &newServer
	}
	log.Printf("[INFO] resourceAviServerCreateOrUpdate pool %v server %v", pUUID, pserver)
	//set other attributes from server.
	if hostname, ok := d.GetOk("hostname"); ok {
		pserver.Hostname = hostname.(string)
	}
	if AutoscalingGroupName, ok := d.GetOk("autoscaling_group_name"); ok {
		pserver.AutoscalingGroupName = AutoscalingGroupName.(string)
	}
	if AvailabilityZone, ok := d.GetOk("availability_zone"); ok {
		pserver.AvailabilityZone = AvailabilityZone.(string)
	}
	if Description, ok := d.GetOk("description"); ok {
		pserver.Description = Description.(string)
	}
	if DiscoveredNetworks, ok := d.GetOk("discovered_networks"); ok {
		pserver.DiscoveredNetworks = DiscoveredNetworks.([]*models.DiscoveredNetwork)
	}
	if Enabled, ok := d.GetOk("enabled"); ok {
		pserver.Enabled = Enabled.(bool)
	}
	if ExternalOrchestrationID, ok := d.GetOk("external_orchestration_id"); ok {
		pserver.ExternalOrchestrationID = ExternalOrchestrationID.(string)
	}
	if ExternalUUID, ok := d.GetOk("external_uuid"); ok {
		pserver.ExternalUUID = ExternalUUID.(string)
	}
	if Location, ok := d.GetOk("location"); ok {
		pserver.Location = Location.(*models.GeoLocation)
	}
	if MacAddress, ok := d.GetOk("mac_address"); ok {
		pserver.MacAddress = MacAddress.(string)
	}
	if NwRef, ok := d.GetOk("nw_ref"); ok {
		pserver.NwRef = NwRef.(string)
	}
	if PrstHdrVal, ok := d.GetOk("prst_hdr_val"); ok {
		pserver.PrstHdrVal = PrstHdrVal.(string)
	}
	if Ratio, ok := d.GetOk("ratio"); ok {
		pserver.Ratio = Ratio.(int32)
	}
	if ResolveServerByDNS, ok := d.GetOk("resolve_server_by_dns"); ok {
		pserver.ResolveServerByDNS = ResolveServerByDNS.(bool)
	}
	if RewriteHostHeader, ok := d.GetOk("rewrite_host_header"); ok {
		pserver.RewriteHostHeader = RewriteHostHeader.(bool)
	}
	if ServerNode, ok := d.GetOk("server_node"); ok {
		pserver.ServerNode = ServerNode.(string)
	}
	if Static, ok := d.GetOk("static"); ok {
		pserver.Static = Static.(bool)
	}
	if VerifyNetwork, ok := d.GetOk("verify_network"); ok {
		pserver.VerifyNetwork = VerifyNetwork.(bool)
	}
	if VMRef, ok := d.GetOk("vm_ref"); ok {
		pserver.VMRef = VMRef.(string)
	}
	if t, ok := d.GetOk("type"); ok {
		pserver.IP = &models.IPAddr{Type: t.(string), Addr: d.Get("ip").(string)}
	}

	uri := "api/pool/" + pUUID
	var response interface{}
	patchPool := models.Pool{}
	patchPool.Name = poolObj.Name
	patchPool.TenantRef = poolObj.TenantRef
	patchPool.CloudRef = poolObj.CloudRef
	patchPool.Servers = append(patchPool.Servers, pserver)
	err = client.AviSession.Patch(uri, patchPool, "add", response)
	log.Printf("[INFO] resourceAviServerCreateOrUpdate pool %v poolobj %v err %v response %v",
		pUUID, patchPool, err, response)
	if err == nil {
		err = ResourceAviServerRead(d, meta)
	}
	return err
}

func ResourceAviServerRead(d *schema.ResourceData, meta interface{}) error {
	err, pUUID, _, pserver := resourceAviServerReadApi(d, meta)
	if err == nil && pserver != nil {
		//Set id to include port number. if port is not in tf then use 0
		var sUUID string
		portStr := "0"
		ip := d.Get("ip").(string)
		if port, ok := d.GetOk("port"); ok {
			portStr = strconv.Itoa(port.(int))
		}
		sUUID = pUUID + ":" + ip + ":" + portStr
		log.Printf("[INFO] pool %v ip %v port %v", pUUID, ip, portStr)
		d.SetId(sUUID)
		// Fill in the server information
		if pserver.Hostname != "" {
			d.Set("hostname", pserver.Hostname)
		}

		d.Set("enabled", pserver.Enabled)

		if pserver.AutoscalingGroupName != "" {
			d.Set("autoscaling_group_name", pserver.AutoscalingGroupName)
		}
		if pserver.AvailabilityZone != "" {
			d.Set("availability_zone", pserver.AvailabilityZone)
		}
		if pserver.Description != "" {
			d.Set("description", pserver.Description)
		}
		if pserver.DiscoveredNetworks != nil {
			d.Set("discovered_networks", pserver.DiscoveredNetworks)
		}

		if pserver.ExternalUUID != "" {
			d.Set("external_uuid", pserver.ExternalUUID)
		}
		if pserver.ExternalOrchestrationID != "" {
			d.Set("external_orchestration_id", pserver.ExternalOrchestrationID)
		}

		if pserver.Location != nil {
			d.Set("mac_address", pserver.Location)
		}
		if pserver.NwRef != "" {
			d.Set("nw_ref", pserver.NwRef)
		}
		if pserver.PrstHdrVal != "" {
			d.Set("prst_hdr_val", pserver.PrstHdrVal)
		}
		d.Set("ratio", pserver.Ratio)
		d.Set("resolve_server_by_dns", pserver.ResolveServerByDNS)
		d.Set("rewrite_host_header", pserver.RewriteHostHeader)
		if pserver.ServerNode != "" {
			d.Set("server_node", pserver.ServerNode)
		}
		d.Set("static", pserver.Static)
		d.Set("verify_network", pserver.VerifyNetwork)
		if pserver.VMRef != "" {
			d.Set("vm_ref", pserver.VMRef)
		}
		// Add more fields to read.
	}
	return err
}

func resourceAviServerReadApi(d *schema.ResourceData, meta interface{}) (error, string, *models.Pool, *models.Server) {
	client := meta.(*clients.AviClient)
	pUUID := UUIDFromID(d.Get("pool_ref").(string))
	uri := "api/pool/" + pUUID
	var poolObj *models.Pool
	err := client.AviSession.Get(uri, &poolObj)
	if err != nil {
		log.Printf("[ERROR] pool uuid %v not found", pUUID)
		return err, pUUID, nil, nil
	}
	log.Printf("[INFO] found pool %v", poolObj.Name)
	ip := d.Get("ip").(string)
	port := d.Get("port")

	// find the server in the pool object.
	var matchedServer *models.Server = nil
	for i := 0; i < len(poolObj.Servers); i++ {
		sObj := poolObj.Servers[i]
		if sObj.IP.Addr == ip {
			if (port == nil && sObj.Port == 0) || (port != nil && int32(port.(int)) == sObj.Port) {
				matchedServer = sObj
				break
			}
		}
	}
	return nil, pUUID, poolObj, matchedServer
}

func resourceAviServerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AviClient)
	err, pUUID, poolObj, pserver := resourceAviServerReadApi(d, meta)
	log.Printf("[DEBUG] pool %v %v server %v", pUUID, poolObj.Name, d.Id())
	if pserver != nil {
		uri := "api/pool/" + pUUID
		var response interface{}
		var patchPool = make(map[string]interface{})
		var servers = make([]models.Server, 1)
		servers[0] = *pserver
		patchPool["servers"] = servers
		err = client.AviSession.Patch(uri, patchPool, "delete", response)
		log.Printf("[INFO] pool %v server %v deleted err %v", patchPool, d.Id(), err)
	}
	d.SetId("")
	return err
}
