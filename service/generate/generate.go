package generate

import (
	"fmt"
	"strings"
)

type Category struct {
	Name  string    `yaml:"name"`
	ID    string    `yaml:"id"`
	Items []Website `yaml:"items"`
}

type Website struct {
	Name  string   `yaml:"name"`
	Items []Domain `yaml:"items"`
}

type Domain struct {
	Domain  string `yaml:"domain"`
	Comment string `yaml:"comment"`
}

type Availability struct {
	Resolved   bool
	Reachable  bool
	Successful bool
}

func GeneratePage(categories []Category, availability map[string]Availability) string {
	builder := strings.Builder{}

	builder.WriteString("# <a name=\"start\"></a>awesome-победа.рф (<a href=\"https://победапобеда.рф\">победапобеда.рф</a>)\n\nСписок победных сайтов на просторах рунета.")
	builder.WriteString("\n\n---\n\n")

	builder.WriteString(GenerateTOC(categories))
	builder.WriteString("\n\n---\n\n")

	for _, category := range categories {
		builder.WriteString(GenerateCategory(category, availability))
		builder.WriteString("\n\n")
	}

	builder.WriteString("По всем вопросам писать [@xbt573](https://t.me/xbt573)")

	return strings.TrimSpace(builder.String())
}

func GenerateTOC(categories []Category) string {
	builder := strings.Builder{}

	builder.WriteString(`## <a name="toc"></a>Содержание`)

	for _, category := range categories {
		builder.WriteString(fmt.Sprintf("\n- [%v](#%v)", category.Name, category.ID))
	}

	return builder.String()
}

func GenerateCategory(category Category, availability map[string]Availability) string {
	builder := strings.Builder{}

	builder.WriteString(
		fmt.Sprintf(
			"## <a name=\"%v\"></a>%v\n**[`^        назад        ^`](#start)**",
			category.ID,
			category.Name,
		),
	)

	for _, website := range category.Items {
		builder.WriteString("\n" + GenerateWebsite(website, availability))
	}

	return builder.String()
}

func GenerateWebsite(website Website, availability map[string]Availability) string {
	builder := strings.Builder{}

	if len(website.Items) == 1 {
		builder.WriteString(
			fmt.Sprintf(
				"- %v — [%v](http://%v)",
				website.Name,
				website.Items[0].Domain, website.Items[0].Domain,
			),
		)

		if status := GenerateStatus(availability[website.Items[0].Domain]); status != "" {
			builder.WriteString(" (" + status + ")")
		}

		if website.Items[0].Comment != "" {
			builder.WriteString(" (" + website.Items[0].Comment + ")")
		}
	} else if len(website.Items) > 1 {
		builder.WriteString(
			fmt.Sprintf(
				"- %v:",
				website.Name,
			),
		)

		for _, domain := range website.Items {
			builder.WriteString(
				fmt.Sprintf(
					"\n    - [%v](http://%v)",
					domain.Domain, domain.Domain,
				),
			)

			if status := GenerateStatus(availability[domain.Domain]); status != "" {
				builder.WriteString(" (" + status + ")")
			}

			if domain.Comment != "" {
				builder.WriteString(" (" + domain.Comment + ")")
			}
		}
	}

	return builder.String()
}

func GenerateStatus(availability Availability) string {
	switch {
	case !availability.Resolved:
		return "not resolved"
	case !availability.Reachable:
		return "not reachable"
	case !availability.Successful:
		return "non-successful exit code"
	default:
		return ""
	}
}
