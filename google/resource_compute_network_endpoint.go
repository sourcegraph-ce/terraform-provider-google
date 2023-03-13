// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceComputeNetworkEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeNetworkEndpointCreate,
		Read:   resourceComputeNetworkEndpointRead,
		Delete: resourceComputeNetworkEndpointDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeNetworkEndpointImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `IPv4 address of network endpoint. The IP address must belong
to a VM in GCE (either the primary IP or as part of an aliased IP
range).`,
			},
			"network_endpoint_group": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareResourceNames,
				Description:      `The network endpoint group this endpoint is part of.`,
			},
			"instance": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `The name for a specific VM instance that the IP address belongs to.
This is required for network endpoints of type GCE_VM_IP_PORT.
The instance must be in the same zone of network endpoint group.`,
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: `Port number of network endpoint.`,
			},
			"zone": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `Zone where the containing network endpoint group is located.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeNetworkEndpointCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	instanceProp, err := expandNestedComputeNetworkEndpointInstance(d.Get("instance"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("instance"); !isEmptyValue(reflect.ValueOf(instanceProp)) && (ok || !reflect.DeepEqual(v, instanceProp)) {
		obj["instance"] = instanceProp
	}
	portProp, err := expandNestedComputeNetworkEndpointPort(d.Get("port"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("port"); !isEmptyValue(reflect.ValueOf(portProp)) && (ok || !reflect.DeepEqual(v, portProp)) {
		obj["port"] = portProp
	}
	ipAddressProp, err := expandNestedComputeNetworkEndpointIpAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("ip_address"); !isEmptyValue(reflect.ValueOf(ipAddressProp)) && (ok || !reflect.DeepEqual(v, ipAddressProp)) {
		obj["ipAddress"] = ipAddressProp
	}

	obj, err = resourceComputeNetworkEndpointEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	lockName, err := replaceVars(d, config, "networkEndpoint/{{project}}/{{zone}}/{{network_endpoint_group}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/attachNetworkEndpoints")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new NetworkEndpoint: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkEndpoint: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating NetworkEndpoint: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{project}}/{{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = ComputeOperationWaitTime(
		config, res, project, "Creating NetworkEndpoint", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create NetworkEndpoint: %s", err)
	}

	log.Printf("[DEBUG] Finished creating NetworkEndpoint %q: %#v", d.Id(), res)

	return resourceComputeNetworkEndpointRead(d, meta)
}

func resourceComputeNetworkEndpointRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/listNetworkEndpoints")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkEndpoint: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "POST", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeNetworkEndpoint %q", d.Id()))
	}

	res, err = flattenNestedComputeNetworkEndpoint(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Object isn't there any more - remove it from the state.
		log.Printf("[DEBUG] Removing ComputeNetworkEndpoint because it couldn't be matched.")
		d.SetId("")
		return nil
	}

	res, err = resourceComputeNetworkEndpointDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing ComputeNetworkEndpoint because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading NetworkEndpoint: %s", err)
	}

	if err := d.Set("instance", flattenNestedComputeNetworkEndpointInstance(res["instance"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpoint: %s", err)
	}
	if err := d.Set("port", flattenNestedComputeNetworkEndpointPort(res["port"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpoint: %s", err)
	}
	if err := d.Set("ip_address", flattenNestedComputeNetworkEndpointIpAddress(res["ipAddress"], d, config)); err != nil {
		return fmt.Errorf("Error reading NetworkEndpoint: %s", err)
	}

	return nil
}

func resourceComputeNetworkEndpointDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for NetworkEndpoint: %s", err)
	}
	billingProject = project

	lockName, err := replaceVars(d, config, "networkEndpoint/{{project}}/{{zone}}/{{network_endpoint_group}}")
	if err != nil {
		return err
	}
	mutexKV.Lock(lockName)
	defer mutexKV.Unlock(lockName)

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/detachNetworkEndpoints")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	toDelete := make(map[string]interface{})
	instanceProp, err := expandNestedComputeNetworkEndpointInstance(d.Get("instance"), d, config)
	if err != nil {
		return err
	}
	if instanceProp != "" {
		toDelete["instance"] = instanceProp
	}

	portProp, err := expandNestedComputeNetworkEndpointPort(d.Get("port"), d, config)
	if err != nil {
		return err
	}
	if portProp != 0 {
		toDelete["port"] = portProp
	}

	ipAddressProp, err := expandNestedComputeNetworkEndpointIpAddress(d.Get("ip_address"), d, config)
	if err != nil {
		return err
	}
	toDelete["ipAddress"] = ipAddressProp

	obj = map[string]interface{}{
		"networkEndpoints": []map[string]interface{}{toDelete},
	}
	log.Printf("[DEBUG] Deleting NetworkEndpoint %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "NetworkEndpoint")
	}

	err = ComputeOperationWaitTime(
		config, res, project, "Deleting NetworkEndpoint", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting NetworkEndpoint %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeNetworkEndpointImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	// instance is optional, so use * instead of + when reading the import id
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/zones/(?P<zone>[^/]+)/networkEndpointGroups/(?P<network_endpoint_group>[^/]+)/(?P<instance>[^/]*)/(?P<ip_address>[^/]+)/(?P<port>[^/]+)",
		"(?P<project>[^/]+)/(?P<zone>[^/]+)/(?P<network_endpoint_group>[^/]+)/(?P<instance>[^/]*)/(?P<ip_address>[^/]+)/(?P<port>[^/]+)",
		"(?P<zone>[^/]+)/(?P<network_endpoint_group>[^/]+)/(?P<instance>[^/]*)/(?P<ip_address>[^/]+)/(?P<port>[^/]+)",
		"(?P<network_endpoint_group>[^/]+)/(?P<instance>[^/]*)/(?P<ip_address>[^/]+)/(?P<port>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{project}}/{{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNestedComputeNetworkEndpointInstance(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func flattenNestedComputeNetworkEndpointPort(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles int given in float64 format
	if floatVal, ok := v.(float64); ok {
		return int(floatVal)
	}
	return v
}

func flattenNestedComputeNetworkEndpointIpAddress(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandNestedComputeNetworkEndpointInstance(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return GetResourceNameFromSelfLink(v.(string)), nil
}

func expandNestedComputeNetworkEndpointPort(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNestedComputeNetworkEndpointIpAddress(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func resourceComputeNetworkEndpointEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// Network Endpoint Group is a URL parameter only, so replace self-link/path with resource name only.
	if err := d.Set("network_endpoint_group", GetResourceNameFromSelfLink(d.Get("network_endpoint_group").(string))); err != nil {
		return nil, fmt.Errorf("Error setting network_endpoint_group: %s", err)
	}

	wrappedReq := map[string]interface{}{
		"networkEndpoints": []interface{}{obj},
	}
	return wrappedReq, nil
}

func flattenNestedComputeNetworkEndpoint(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	var v interface{}
	var ok bool

	v, ok = res["items"]
	if !ok || v == nil {
		return nil, nil
	}

	switch v.(type) {
	case []interface{}:
		break
	case map[string]interface{}:
		// Construct list out of single nested resource
		v = []interface{}{v}
	default:
		return nil, fmt.Errorf("expected list or map for value items. Actual value: %v", v)
	}

	_, item, err := resourceComputeNetworkEndpointFindNestedObjectInList(d, meta, v.([]interface{}))
	if err != nil {
		return nil, err
	}
	return item, nil
}

func resourceComputeNetworkEndpointFindNestedObjectInList(d *schema.ResourceData, meta interface{}, items []interface{}) (index int, item map[string]interface{}, err error) {
	expectedInstance, err := expandNestedComputeNetworkEndpointInstance(d.Get("instance"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedInstance := flattenNestedComputeNetworkEndpointInstance(expectedInstance, d, meta.(*Config))
	expectedIpAddress, err := expandNestedComputeNetworkEndpointIpAddress(d.Get("ip_address"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedIpAddress := flattenNestedComputeNetworkEndpointIpAddress(expectedIpAddress, d, meta.(*Config))
	expectedPort, err := expandNestedComputeNetworkEndpointPort(d.Get("port"), d, meta.(*Config))
	if err != nil {
		return -1, nil, err
	}
	expectedFlattenedPort := flattenNestedComputeNetworkEndpointPort(expectedPort, d, meta.(*Config))

	// Search list for this resource.
	for idx, itemRaw := range items {
		if itemRaw == nil {
			continue
		}
		item := itemRaw.(map[string]interface{})

		// Decode list item before comparing.
		item, err := resourceComputeNetworkEndpointDecoder(d, meta, item)
		if err != nil {
			return -1, nil, err
		}

		itemInstance := flattenNestedComputeNetworkEndpointInstance(item["instance"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemInstance)) && isEmptyValue(reflect.ValueOf(expectedFlattenedInstance))) && !reflect.DeepEqual(itemInstance, expectedFlattenedInstance) {
			log.Printf("[DEBUG] Skipping item with instance= %#v, looking for %#v)", itemInstance, expectedFlattenedInstance)
			continue
		}
		itemIpAddress := flattenNestedComputeNetworkEndpointIpAddress(item["ipAddress"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemIpAddress)) && isEmptyValue(reflect.ValueOf(expectedFlattenedIpAddress))) && !reflect.DeepEqual(itemIpAddress, expectedFlattenedIpAddress) {
			log.Printf("[DEBUG] Skipping item with ipAddress= %#v, looking for %#v)", itemIpAddress, expectedFlattenedIpAddress)
			continue
		}
		itemPort := flattenNestedComputeNetworkEndpointPort(item["port"], d, meta.(*Config))
		// isEmptyValue check so that if one is nil and the other is "", that's considered a match
		if !(isEmptyValue(reflect.ValueOf(itemPort)) && isEmptyValue(reflect.ValueOf(expectedFlattenedPort))) && !reflect.DeepEqual(itemPort, expectedFlattenedPort) {
			log.Printf("[DEBUG] Skipping item with port= %#v, looking for %#v)", itemPort, expectedFlattenedPort)
			continue
		}
		log.Printf("[DEBUG] Found item for resource %q: %#v)", d.Id(), item)
		return idx, item, nil
	}
	return -1, nil, nil
}
func resourceComputeNetworkEndpointDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	v, ok := res["networkEndpoint"]
	if !ok || v == nil {
		return res, nil
	}

	return v.(map[string]interface{}), nil
}
