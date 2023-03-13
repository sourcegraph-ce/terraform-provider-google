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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceEssentialContactsContact() *schema.Resource {
	return &schema.Resource{
		Create: resourceEssentialContactsContactCreate,
		Read:   resourceEssentialContactsContactRead,
		Update: resourceEssentialContactsContactUpdate,
		Delete: resourceEssentialContactsContactDelete,

		Importer: &schema.ResourceImporter{
			State: resourceEssentialContactsContactImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The email address to send notifications to. This does not need to be a Google account.`,
			},
			"language_tag": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The preferred language for notifications, as a ISO 639-1 language code. See Supported languages for a list of supported languages.`,
			},
			"notification_category_subscriptions": {
				Type:        schema.TypeList,
				Required:    true,
				Description: `The categories of notifications that the contact will receive communications for.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"parent": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or projects/{project_id}`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The identifier for the contact. Format: {resourceType}/{resource_id}/contacts/{contact_id}`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceEssentialContactsContactCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	emailProp, err := expandEssentialContactsContactEmail(d.Get("email"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("email"); !isEmptyValue(reflect.ValueOf(emailProp)) && (ok || !reflect.DeepEqual(v, emailProp)) {
		obj["email"] = emailProp
	}
	notificationCategorySubscriptionsProp, err := expandEssentialContactsContactNotificationCategorySubscriptions(d.Get("notification_category_subscriptions"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_category_subscriptions"); !isEmptyValue(reflect.ValueOf(notificationCategorySubscriptionsProp)) && (ok || !reflect.DeepEqual(v, notificationCategorySubscriptionsProp)) {
		obj["notificationCategorySubscriptions"] = notificationCategorySubscriptionsProp
	}
	languageTagProp, err := expandEssentialContactsContactLanguageTag(d.Get("language_tag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("language_tag"); !isEmptyValue(reflect.ValueOf(languageTagProp)) && (ok || !reflect.DeepEqual(v, languageTagProp)) {
		obj["languageTag"] = languageTagProp
	}

	url, err := replaceVars(d, config, "{{EssentialContactsBasePath}}{{parent}}/contacts")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Contact: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Contact: %s", err)
	}
	if err := d.Set("name", flattenEssentialContactsContactName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Contact %q: %#v", d.Id(), res)

	return resourceEssentialContactsContactRead(d, meta)
}

func resourceEssentialContactsContactRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{EssentialContactsBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("EssentialContactsContact %q", d.Id()))
	}

	if err := d.Set("name", flattenEssentialContactsContactName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Contact: %s", err)
	}
	if err := d.Set("email", flattenEssentialContactsContactEmail(res["email"], d, config)); err != nil {
		return fmt.Errorf("Error reading Contact: %s", err)
	}
	if err := d.Set("notification_category_subscriptions", flattenEssentialContactsContactNotificationCategorySubscriptions(res["notificationCategorySubscriptions"], d, config)); err != nil {
		return fmt.Errorf("Error reading Contact: %s", err)
	}
	if err := d.Set("language_tag", flattenEssentialContactsContactLanguageTag(res["languageTag"], d, config)); err != nil {
		return fmt.Errorf("Error reading Contact: %s", err)
	}

	return nil
}

func resourceEssentialContactsContactUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	notificationCategorySubscriptionsProp, err := expandEssentialContactsContactNotificationCategorySubscriptions(d.Get("notification_category_subscriptions"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("notification_category_subscriptions"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, notificationCategorySubscriptionsProp)) {
		obj["notificationCategorySubscriptions"] = notificationCategorySubscriptionsProp
	}
	languageTagProp, err := expandEssentialContactsContactLanguageTag(d.Get("language_tag"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("language_tag"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, languageTagProp)) {
		obj["languageTag"] = languageTagProp
	}

	url, err := replaceVars(d, config, "{{EssentialContactsBasePath}}{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Contact %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("notification_category_subscriptions") {
		updateMask = append(updateMask, "notificationCategorySubscriptions")
	}

	if d.HasChange("language_tag") {
		updateMask = append(updateMask, "languageTag")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Contact %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Contact %q: %#v", d.Id(), res)
	}

	return resourceEssentialContactsContactRead(d, meta)
}

func resourceEssentialContactsContactDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := replaceVars(d, config, "{{EssentialContactsBasePath}}{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Contact %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := SendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Contact")
	}

	log.Printf("[DEBUG] Finished deleting Contact %q: %#v", d.Id(), res)
	return nil
}

func resourceEssentialContactsContactImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"(?P<name>.+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenEssentialContactsContactName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenEssentialContactsContactEmail(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenEssentialContactsContactNotificationCategorySubscriptions(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenEssentialContactsContactLanguageTag(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandEssentialContactsContactEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandEssentialContactsContactNotificationCategorySubscriptions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandEssentialContactsContactLanguageTag(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
