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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFilestoreSnapshot_filestoreSnapshotBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckFilestoreSnapshotDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFilestoreSnapshot_filestoreSnapshotBasicExample(context),
			},
			{
				ResourceName:            "google_filestore_snapshot.snapshot",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "instance"},
			},
		},
	})
}

func testAccFilestoreSnapshot_filestoreSnapshotBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_filestore_snapshot" "snapshot" {
  name     = "tf-test-test-snapshot%{random_suffix}"
  instance = google_filestore_instance.instance.name
  location = "us-central1"
}

resource "google_filestore_instance" "instance" {
  name     = "tf-test-test-instance-for-snapshot%{random_suffix}"
  location = "us-central1"
  tier     = "ENTERPRISE"

  file_shares {
    capacity_gb = 1024
    name        = "share1"
  }

  networks {
    network = "default"
    modes   = ["MODE_IPV4"]
  }
}
`, context)
}

func TestAccFilestoreSnapshot_filestoreSnapshotFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": RandString(t, 10),
	}

	VcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    TestAccProviders,
		CheckDestroy: testAccCheckFilestoreSnapshotDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFilestoreSnapshot_filestoreSnapshotFullExample(context),
			},
			{
				ResourceName:            "google_filestore_snapshot.snapshot",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"name", "location", "instance"},
			},
		},
	})
}

func testAccFilestoreSnapshot_filestoreSnapshotFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_filestore_snapshot" "snapshot" {
  name     = "tf-test-test-snapshot%{random_suffix}"
  instance = google_filestore_instance.instance.name
  location = "us-central1"

  description = "Snapshot of tf-test-test-instance-for-snapshot%{random_suffix}"

  labels = {
    my_label = "value"
  }
}

resource "google_filestore_instance" "instance" {
  name     = "tf-test-test-instance-for-snapshot%{random_suffix}"
  location = "us-central1"
  tier     = "ENTERPRISE"

  file_shares {
    capacity_gb = 1024
    name        = "share1"
  }

  networks {
    network = "default"
    modes   = ["MODE_IPV4"]
  }
}
`, context)
}

func testAccCheckFilestoreSnapshotDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_filestore_snapshot" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := GoogleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{FilestoreBasePath}}projects/{{project}}/locations/{{location}}/instances/{{instance}}/snapshots/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = SendRequest(config, "GET", billingProject, url, config.UserAgent, nil, isNotFilestoreQuotaError)
			if err == nil {
				return fmt.Errorf("FilestoreSnapshot still exists at %s", url)
			}
		}

		return nil
	}
}
