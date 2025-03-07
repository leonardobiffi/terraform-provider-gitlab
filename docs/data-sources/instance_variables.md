---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitlab_instance_variables Data Source - terraform-provider-gitlab"
subcategory: ""
description: |-
  The gitlab_instance_variables data source allows to retrieve all instance-level CI/CD variables.
  Upstream API: GitLab REST API docs https://docs.gitlab.com/ee/api/instance_level_ci_variables.html
---

# gitlab_instance_variables (Data Source)

The `gitlab_instance_variables` data source allows to retrieve all instance-level CI/CD variables.

**Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/instance_level_ci_variables.html)



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) The ID of this resource.

### Read-Only

- `variables` (List of Object) The list of variables returned by the search (see [below for nested schema](#nestedatt--variables))

<a id="nestedatt--variables"></a>
### Nested Schema for `variables`

Read-Only:

- `key` (String)
- `masked` (Boolean)
- `protected` (Boolean)
- `value` (String)
- `variable_type` (String)


