{{/* vim: set filetype=mustache: */}}

{{- define "gloo.roleKind" -}}
{{- if .Values.global.glooRbac.namespaced -}}
Role
{{- else -}}
ClusterRole
{{- end -}}
{{- end -}}

{{- define "gloo.rbacNameSuffix" -}}
{{- if .Values.global.glooRbac.nameSuffix -}}
-{{ .Values.global.glooRbac.nameSuffix }}
{{- else if not .Values.global.glooRbac.namespaced -}}
-{{ .Release.Namespace }}
{{- end -}}
{{- end -}}

{{/*
Expand the name of a container image
*/}}
{{- define "gloo.image" -}}
{{ .registry }}/{{ .repository }}:{{ .tag }}{{ ternary "-extended" "" (default false .extended) }}
{{- end -}}

{{- define "gloo.pullSecret" -}}
{{- if .pullSecret -}}
imagePullSecrets:
- name: {{ .pullSecret }}
{{ end -}}
{{- end -}}

{{- /*
    gloo.util.merge is a fork of a helm library chart function (https://github.com/helm/charts/blob/master/incubator/common/templates/_util.tpl)

This takes an array of three values:
- the top context
- the yaml block that will be merged in (override)
- the template name of the base (source)

The behaviour is as follows, to align with already existing helm behaviour:
- If no source is found (template is empty), the merged output will be empty
- If no overrides are specified, the source is rendered as is
- If overrides are specified and source is not empty, overrides will be merged in to the source.

Overrides can replace / add to deeply nested dictionaries, but will completely replace lists.
Examples:

┌─────────────────────┬───────────────────────┬────────────────────────┐
│       Source        │       Overrides       │        Result          │
├─────────────────────┼───────────────────────┼────────────────────────┤
│ metadata:           │ metadata:             │ metadata:              │
│   labels:           │   labels:             │   labels:              │
│     app: gloo       │    app:gloo1          │     app: gloo1         │
│     cluster: useast │    author: infra-team │     author: infra-team │
│                     │                       │     cluster: useast    │
├─────────────────────┼───────────────────────┼────────────────────────┤
│ lists:              │ lists:                │ lists:                 │
│   groceries:        │  groceries:           │   groceries:           │
│   - apple           │   - grapes            │   - grapes             │
│   - banana          │                       │                        │
└─────────────────────┴───────────────────────┴────────────────────────┘

*/ -}}
{{- define "gloo.util.merge" -}}
{{- $top := first . -}}
{{- $overrides := (index . 1) -}}
{{- $tpl := fromYaml (include (index . 2) $top) -}}
{{- if or (empty $overrides) (empty $tpl) -}}
{{ include (index . 2) $top }} {{/* render source as is */}}
{{- else -}}
{{- $merged := merge $overrides $tpl -}}
{{- toYaml (merge $overrides $tpl) -}} {{/* render source with overrides as YAML */}}
{{- end -}}
{{- end -}}
