#compdef {{.Command}}

{{if .Properties.Author}}# Copyright (c) {{dateYear}} {{.Properties.Author}}{{end}}
{{if .Properties.License}}# License: {{.Properties.License}}{{end}}

function _{{.Command}}() {
    local context curcontext=$curcontext state line
    typeset -A opt_args
    local ret=1

    _arguments -C \
      {{if .Properties.Help.Option -}}
        '{{.Properties.Help.Option | dealWithExclusion}}{{.Properties.Help.Option | dealWithOption}}[{{.Properties.Help.Description | dealWithDescription}}]' \
      {{end -}}
      {{if .Properties.Version.Option -}}
        '{{.Properties.Version.Option | dealWithExclusion}}{{.Properties.Version.Option | dealWithOption}}[{{.Properties.Version.Description | dealWithDescription}}]' \
      {{end -}}
      {{range .Options.Switch -}}
        '{{.| dealWithSwitchExclusion}}{{.| dealWithSwitchOption}}[{{.Description | dealWithDescription}}]' \
      {{end -}}
      {{range .Options.Flag -}}
        '{{.| dealWithFlagExclusion}}{{.| dealWithFlagOption}}[{{.Description | dealWithDescription}}]:{{.| setFlagMessage}}:{{.| setAction}}' \
      {{end -}}
      {{if .Arguments -}}
        '{{if not .Arguments.After_arg}}(-){{end}}*:arguments:{{.Arguments.Type | setAction}}' \
      {{end -}}
        && ret=0

  {{if .Arguments -}}
    case $state in
        {{range .Options.Flag}}{{if .Option | whetherOptionIsEnabled}}{{if .Argument.Type | whetherTypeIsFunc}}{{.| setAction | helperTrimArrowInType}})
          # TODO
          ;;
        {{end}}{{end}}{{end}}{{if .Arguments.Type | whetherTypeIsFunc}}{{.Arguments.Type | setAction | helperTrimArrowInType}})
          # TODO
          ;;{{end}}
    esac
  {{end -}}

    return ret
}

_{{.Command}} "$@"
