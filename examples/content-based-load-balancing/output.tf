# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

output "application_public_ip" {
  value = google_compute_global_forwarding_rule.default.ip_address
}
