{{- $config := . -}}
# <a name="start"></a>awesome-победа.рф

Список победных сайтов на просторах рунета.

---

## <a name="toc"></a>Содержание
{{- range $category := .Config.Categories }}
- [{{$category.Name}}](#{{$category.ID}})
{{- end }}

---

{{- range $category := .Config.Categories }}
## <a name="{{$category.ID}}"></a>{{$category.Name}}
**[`^        назад        ^`](#start)**
    {{- range $domain := $category.Items }}
        {{- if eq (len $domain.Items) 1 }}
            {{- $url := index $domain.Items 0 }}
            {{- $availability := index $config.Availability $url }}
            {{- $status := "" }}
            {{- if not $availability.Resolved }}
                {{- $status = "(домен недоступен)" }}
            {{- else if not $availability.Reachable }}
                {{- $status = "(сервер недоступен)" }}
            {{- else if not $availability.Successful }}
                {{- $status = "(неуспешный статус код)" }}
            {{- end}}
- {{$domain.Name}} — [{{index $domain.Items 0}}](https://{{index $domain.Items 0}}) {{$status}}
        {{- else if gt (len $domain.Items) 1 }}
- {{$domain.Name}}:
            {{- range $url := $domain.Items }}
    - [{{index $domain.Items 0}}](https://{{index $domain.Items 0}})
            {{- end }}
        {{- end }}
    {{- end }}
{{ end }}