---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitlab_group_variables Data Source - terraform-provider-gitlab"
subcategory: ""
description: |-
  The gitlab_group_variables data source allows to retrieve all group-level CI/CD variables.
  Upstream API: GitLab REST API docs https://docs.gitlab.com/ee/api/group_level_variables.html
---

# gitlab_group_variables (Data Source)

The `gitlab_group_variables` data source allows to retrieve all group-level CI/CD variables.

**Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/group_level_variables.html)



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group` (String) The name or id of the group.

### Optional

- `environment_scope` (String) The environment scope of the variable. Defaults to all environment (`*`).
- `id` (String) The ID of this resource.

### Read-Only

- `variables` (List of Object) The list of variables returned by the search (see [below for nested schema](#nestedatt--variables))

<a id="nestedatt--variables"></a>
### Nested Schema for `variables`

Read-Only:

- `environment_scope` (String)
- `group` (String)
- `key` (String)
- `masked` (Boolean)
- `protected` (Boolean)
- `value` (String)
- `variable_type` (String)


