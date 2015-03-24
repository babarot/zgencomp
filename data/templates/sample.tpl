#compdef {{.Command}}

{{if .Properties.Author}}# Copyright (c) {{dateYear}} {{.Properties.Author}}{{end}}
{{if .Properties.License}}# License: {{.Properties.License}}{{end}}

function _{{.Command}}() {
    local context curcontext=$curcontext state line
    typeset -A opt_args
    local ret=1

    _arguments -C \
        '{{.Properties.Help.Option | dealWithExclusion}}{{.Properties.Help.Option | dealWithOption}}[{{.Properties.Help.Description | dealWithDescription}}]' \
        '{{.Properties.Version.Option | dealWithExclusion}}{{.Properties.Version.Option | dealWithOption}}[{{.Properties.Version.Description | dealWithDescription}}]' \{{range .Options.Switch}}{{if .Option | whetherOptionIsEnabled}}
        '{{.| dealWithSwitchExclusion}}{{.| dealWithSwitchOption}}[{{.Description | dealWithDescription}}]' \{{end}}{{end}}{{range .Options.Flag}}{{if .Option | whetherOptionIsEnabled}}
        '{{.| dealWithFlagExclusion}}{{.| dealWithFlagOption}}[{{.Description | dealWithDescription}}]:{{.| setFlagMessage}}:{{.| setAction}}' \{{end}}{{end}}
        '{{if not .Arguments.After_arg}}(-){{end}}*:arguments:{{.Arguments.Type | setAction}}' \
        && ret=0

    case $state in
        {{range .Options.Flag}}{{if .Option | whetherOptionIsEnabled}}{{if .Argument.Type | whetherTypeIsFunc}}{{.| setAction | helperTrimArrowInType}})
          # TODO
          ;;
        {{end}}{{end}}{{end}}{{if .Arguments.Type | whetherTypeIsFunc}}{{.Arguments.Type | setAction | helperTrimArrowInType}})
          # TODO
          ;;{{end}}
    esac

    return ret
}

_{{.Command}} "$@"
