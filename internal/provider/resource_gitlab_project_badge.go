package provider

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	gitlab "github.com/xanzy/go-gitlab"
)

var _ = registerResource("gitlab_project_badge", func() *schema.Resource {
	return &schema.Resource{
		Description: `The ` + "`gitlab_project_badge`" + ` resource allows to mange the lifecycle of project badges.

**Upstream API**: [GitLab REST API docs](https://docs.gitlab.com/ee/user/project/badges.html#project-badges)`,

		CreateContext: resourceGitlabProjectBadgeCreate,
		ReadContext:   resourceGitlabProjectBadgeRead,
		UpdateContext: resourceGitlabProjectBadgeUpdate,
		DeleteContext: resourceGitlabProjectBadgeDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"project": {
				Description: "The id of the project to add the badge to.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"link_url": {
				Description: "The url linked with the badge.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"image_url": {
				Description: "The image url which will be presented on project overview.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"rendered_link_url": {
				Description: "The link_url argument rendered (in case of use of placeholders).",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"rendered_image_url": {
				Description: "The image_url argument rendered (in case of use of placeholders).",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
})

func resourceGitlabProjectBadgeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)
	projectID := d.Get("project").(string)
	options := &gitlab.AddProjectBadgeOptions{
		LinkURL:  gitlab.String(d.Get("link_url").(string)),
		ImageURL: gitlab.String(d.Get("image_url").(string)),
	}

	log.Printf("[DEBUG] create gitlab project badge %q / %q", *options.LinkURL, *options.ImageURL)

	badge, _, err := client.ProjectBadges.AddProjectBadge(projectID, options, gitlab.WithContext(ctx))
	if err != nil {
		return diag.FromErr(err)
	}

	badgeID := strconv.Itoa(badge.ID)

	d.SetId(buildTwoPartID(&projectID, &badgeID))

	return resourceGitlabProjectBadgeRead(ctx, d, meta)
}

func resourceGitlabProjectBadgeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)
	ids := strings.Split(d.Id(), ":")
	projectID := ids[0]
	badgeID, err := strconv.Atoi(ids[1])
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] read gitlab project badge %s/%d", projectID, badgeID)

	badge, _, err := client.ProjectBadges.GetProjectBadge(projectID, badgeID, gitlab.WithContext(ctx))
	if err != nil {
		if is404(err) {
			log.Printf("[DEBUG] project badge %d in project %s doesn't exist anymore, removing from state", badgeID, projectID)
			d.SetId("")
			return nil
		}
		return diag.FromErr(err)
	}

	resourceGitlabProjectBadgeSetToState(d, badge, &projectID)
	return nil
}

func resourceGitlabProjectBadgeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)
	ids := strings.Split(d.Id(), ":")
	projectID := ids[0]
	badgeID, err := strconv.Atoi(ids[1])
	if err != nil {
		return diag.FromErr(err)
	}

	options := &gitlab.EditProjectBadgeOptions{
		LinkURL:  gitlab.String(d.Get("link_url").(string)),
		ImageURL: gitlab.String(d.Get("image_url").(string)),
	}

	log.Printf("[DEBUG] update gitlab project badge %s/%d", projectID, badgeID)

	_, _, err = client.ProjectBadges.EditProjectBadge(projectID, badgeID, options, gitlab.WithContext(ctx))
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceGitlabProjectBadgeRead(ctx, d, meta)
}

func resourceGitlabProjectBadgeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*gitlab.Client)
	ids := strings.Split(d.Id(), ":")
	projectID := ids[0]
	badgeID, err := strconv.Atoi(ids[1])
	if err != nil {
		return diag.FromErr(err)
	}

	log.Printf("[DEBUG] Delete gitlab project badge %s/%d", projectID, badgeID)

	_, err = client.ProjectBadges.DeleteProjectBadge(projectID, badgeID, gitlab.WithContext(ctx))
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceGitlabProjectBadgeSetToState(d *schema.ResourceData, badge *gitlab.ProjectBadge, projectID *string) {
	d.Set("link_url", badge.LinkURL)
	d.Set("image_url", badge.ImageURL)
	d.Set("rendered_link_url", badge.RenderedLinkURL)
	d.Set("rendered_image_url", badge.RenderedImageURL)
	d.Set("project", projectID)
}
