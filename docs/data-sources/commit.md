---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gdiff_commit Data Source - terraform-provider-gdiff"
subcategory: ""
description: |-
  This data source provides the last commit of a file or folder
---

# gdiff_commit (Data Source)

This data source provides the last commit of a file or folder

## Example Usage

```terraform
data "gdiff_commit" "example" {
  path = "foo"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `path` (String) Absolute path to the file,directory.

### Read-Only

- `hash` (String) The hash of the last commit
- `id` (String) The ID of this resource.


