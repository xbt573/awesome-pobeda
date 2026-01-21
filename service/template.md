{{- define "status" }}
    {{- if not .Resolved }}(домен не резолвится) {{ end }}
    {{- if and .Resolved (not .Reachable) }}(сервер недоступен) {{ end }}
    {{- if and .Resolved .Reachable (not .Successful) }}(неуспешный статус код) {{ end }}

    {{- if .Comment }}{{ .Comment }}{{ end }}
{{- end -}}

# <a name="start"></a>awesome-победа.рф

Список победных сайтов на просторах рунета.

---

## <a name="toc"></a>Содержание
{{- range .Config.Categories }}
- [{{ .Name }}](#{{ .ID }})
{{- end }}

---
{{ range .Config.Categories }}
## <a name="{{ .ID }}"></a>{{ .Name }}
**[`^        назад        ^`](#start)**
    {{- range .Items}}
        {{- if eq (len .Items) 1 }}
- {{ .Name }} — [{{ (index .Items 0).Domain }}](http://{{ (index .Items 0).Domain }}) {{ template "status" (index $.Availability (index .Items 0).Domain) }}
        {{- else }}
- {{ .Name }}:
            {{- range .Items}}
    - [{{ . }}](http://{{ . }}) {{ template "status" (index $.Availability .) }}
            {{- end }}
        {{- end }}
    {{- end }}
{{ end -}}

--

По вопросам наполнения обращаться к https://t.me/xbt573.
