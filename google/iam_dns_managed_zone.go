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

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var DNSManagedZoneIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"managed_zone": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type DNSManagedZoneIamUpdater struct {
	project     string
	managedZone string
	d           TerraformResourceData
	Config      *Config
}

func DNSManagedZoneIamUpdaterProducer(d TerraformResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	if v, ok := d.GetOk("managed_zone"); ok {
		values["managed_zone"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/managedZones/(?P<managed_zone>[^/]+)", "(?P<project>[^/]+)/(?P<managed_zone>[^/]+)", "(?P<managed_zone>[^/]+)"}, d, config, d.Get("managed_zone").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DNSManagedZoneIamUpdater{
		project:     values["project"],
		managedZone: values["managed_zone"],
		d:           d,
		Config:      config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("managed_zone", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting managed_zone: %s", err)
	}

	return u, nil
}

func DNSManagedZoneIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/managedZones/(?P<managed_zone>[^/]+)", "(?P<project>[^/]+)/(?P<managed_zone>[^/]+)", "(?P<managed_zone>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DNSManagedZoneIamUpdater{
		project:     values["project"],
		managedZone: values["managed_zone"],
		d:           d,
		Config:      config,
	}
	if err := d.Set("managed_zone", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting managed_zone: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *DNSManagedZoneIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyManagedZoneUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return nil, err
	}

	policy, err := SendRequest(u.Config, "POST", project, url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *DNSManagedZoneIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyManagedZoneUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.UserAgent)
	if err != nil {
		return err
	}

	_, err = SendRequestWithTimeout(u.Config, "POST", project, url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *DNSManagedZoneIamUpdater) qualifyManagedZoneUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{DNSBasePath}}%s:%s", fmt.Sprintf("projects/%s/managedZones/%s", u.project, u.managedZone), methodIdentifier)
	url, err := replaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *DNSManagedZoneIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/managedZones/%s", u.project, u.managedZone)
}

func (u *DNSManagedZoneIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-dns-managedzone-%s", u.GetResourceId())
}

func (u *DNSManagedZoneIamUpdater) DescribeResource() string {
	return fmt.Sprintf("dns managedzone %q", u.GetResourceId())
}
