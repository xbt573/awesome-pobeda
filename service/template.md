# <a name="start"></a>awesome-победа.рф

Список победных сайтов на просторах рунета.

---

## <a name="toc"></a>Содержание
<% data.each do |category| -%>
- [<%= category["name"] %>](#<%= category["id"] %>)
<% end -%>

---

<% data.each do |category| -%>
## <a name="<%= category["id"] %>"></a><%= category["name"] %>
**[`^        назад        ^`](#start)**
<% category["items"].each do |key, domains| -%>
<% if domains.size > 1 -%>
- <%= key %>:
<% domains.each do |domain| -%>
    - [<%= domain %>](<%= availability[domain][1] || "https" %>://<%= domain %>)<%= ' (может быть недоступен)' unless availability[domain][0] %>
<% end -%>
<% else -%>
- <%= key %> — [<%= domains.first %>](<%= availability[domains.first][1] || "https" %>://<%= domains.first %>)<%= ' (может быть недоступен)' unless availability[domains.first][0] %>
<% end -%>
<% end -%>
<% end -%>
