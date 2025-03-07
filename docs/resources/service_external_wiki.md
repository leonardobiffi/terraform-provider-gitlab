---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "gitlab_service_external_wiki Resource - terraform-provider-gitlab"
subcategory: ""
description: |-
  The gitlab_service_external_wiki resource allows to manage the lifecycle of a project integration with External Wiki Service.
  Upstream API: GitLab REST API docs https://docs.gitlab.com/ee/api/integrations.html#external-wiki
---

# gitlab_service_external_wiki (Resource)

The `gitlab_service_external_wiki` resource allows to manage the lifecycle of a project integration with External Wiki Service.

**Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/api/integrations.html#external-wiki)

## Example Usage

```terraform
resource "gitlab_project" "awesome_project" {
  name             = "awesome_project"
  description      = "My awesome project."
  visibility_level = "public"
}

resource "gitlab_service_external_wiki" "wiki" {
  project           = gitlab_project.awesome_project.id
  external_wiki_url = "https://MyAwesomeExternalWikiURL.com"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `external_wiki_url` (String) The URL of the external wiki.
- `project` (String) ID of the project you want to activate integration on.

### Optional

- `id` (String) The ID of this resource.

### Read-Only

- `active` (Boolean) Whether the integration is active.
- `created_at` (String) The ISO8601 date/time that this integration was activated at in UTC.
- `slug` (String) The name of the integration in lowercase, shortened to 63 bytes, and with everything except 0-9 and a-z replaced with -. No leading / trailing -. Use in URLs, host names and domain names.
- `title` (String) Title of the integration.
- `updated_at` (String) The ISO8601 date/time that this integration was last updated at in UTC.

## Import

Import is supported using the following syntax:

```shell
# You can import a gitlab_service_external_wiki state using the project ID, e.g.
terraform import gitlab_service_external_wiki.wiki 1
```
