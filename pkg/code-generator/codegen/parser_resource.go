package codegen

import (
	"sort"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/pseudomuto/protokit"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/log"
)

const (
	// solo-kit types
	// required fields
	metadataTypeName = ".core.solo.io.Metadata"
	statusTypeName   = ".core.solo.io.Status"

	// magic comments
	shortNameDeclaration      = "@solo-kit:resource.short_name="
	pluralNameDeclaration     = "@solo-kit:resource.plural_name="
	resourceGroupsDeclaration = "@solo-kit:resource.resource_groups="
)

func getResources(project *Project, messages []*protokit.Descriptor) ([]*Resource, []*ResourceGroup, error) {
	resourcesByGroup := make(map[string][]*Resource)
	var resources []*Resource
	for _, msg := range messages {
		resource, groups, err := describeResource(msg)
		if err != nil {
			return nil, nil, err
		}
		if resource == nil {
			// message is not a resource
			continue
		}
		resource.Project = project
		for _, group := range groups {
			resourcesByGroup[group] = append(resourcesByGroup[group], resource)
		}
		resources = append(resources, resource)
	}

	var resourceGroups []*ResourceGroup

	for group, resources := range resourcesByGroup {
		log.Printf("group: %v", group)
		rg := &ResourceGroup{
			Name:      group,
			GoName:    goName(group),
			Project:   project,
			Resources: resources,
		}
		for _, res := range resources {
			res.Project = project
			res.ResourceGroups = append(res.ResourceGroups, rg)
		}
		resourceGroups = append(resourceGroups, rg)
	}
	for _, res := range resources {
		// sort for stability
		sort.SliceStable(res.ResourceGroups, func(i, j int) bool {
			return res.ResourceGroups[i].Name < res.ResourceGroups[j].Name
		})
	}
	return resources, resourceGroups, nil
}

func describeResource(msg *protokit.Descriptor) (*Resource, []string, error) {
	// not a solo kit resource, or you messed up!
	if !hasField(msg, "metadata", metadataTypeName) {
		return nil, nil, nil
	}

	comments := strings.Split(msg.GetComments().Leading, "\n")

	name := msg.GetName()
	// required flags
	shortName, ok := getCommentValue(comments, shortNameDeclaration)
	if !ok {
		return nil, nil, errors.Errorf("must provide %s", shortNameDeclaration)
	}
	pluralName, ok := getCommentValue(comments, pluralNameDeclaration)
	if !ok {
		return nil, nil, errors.Errorf("must provide %s", pluralNameDeclaration)
	}
	// always make it upper camel
	pluralName = strcase.ToCamel(pluralName)

	// optional flags
	joinedResourceGroups, _ := getCommentValue(comments, resourceGroupsDeclaration)
	resourceGroups := strings.Split(joinedResourceGroups, ",")
	if resourceGroups[0] == "" {
		resourceGroups = nil
	}

	hasStatus := hasField(msg, "status", statusTypeName)

	fields := collectFields(msg)

	return &Resource{
		Name:       name,
		ShortName:  shortName,
		PluralName: pluralName,
		HasStatus:  hasStatus,
		Fields:     fields,
	}, resourceGroups, nil
}
