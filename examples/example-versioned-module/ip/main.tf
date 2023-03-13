# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

resource "google_compute_address" "ip_address" {
  # We'll only generate this block if the value of
  # has_labels is 0! Effectively an if statement.
  count = 1 - local.has_labels

  name = var.name
}

resource "google_compute_address" "ip_address_beta" {
  # And this block is only present if we have
  # at least one entry, effectively an elif.
  count = local.has_labels

  name   = var.name
  labels = var.labels
}
