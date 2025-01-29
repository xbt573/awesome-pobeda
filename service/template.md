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
<% category["items"].each do |item| -%>
- <%= item[0] %> — [<%= item[1] %>](https://<%= item[1] %>) <%= '(может быть недоступен)' unless item[2] %>
<% end -%>

<% end -%>
