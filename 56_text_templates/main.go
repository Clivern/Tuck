// Text Templates
//
// Go offers built-in support for creating dynamic content or showing customized
// output to the user with the text/template package. A sibling package named
// html/template provides the same API but has additional security features and
// should be used for generating HTML.
//
// Run: go run ./56_text_templates
package main

import (
	htmltemplate "html/template"
	"os"
	"strings"
	"text/template"
)

func main() {
	// We can create a new template and parse its body from a string. Templates are
	// a mix of static text and "actions" enclosed in {{...}} that are used to
	// dynamically insert content.
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// Alternatively, we can use the template.Must function to panic in case Parse
	// returns an error. This is especially useful for templates initialized in the
	// global scope.
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// By "executing" the template we generate its text with specific values for its
	// actions. The {{.}} action is replaced by the value passed as a parameter to
	// Execute.
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// Helper function we'll use below.
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// If the data is a struct we can use the {{.FieldName}} action to access its
	// fields. The fields should be exported to be accessible when a template is
	// executing.
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// The same applies to maps; with maps there is no restriction on the case of key
	// names.
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// if/else provide conditional execution for templates. A value is considered
	// false if it's the default value of a type, such as 0, an empty string, nil
	// pointer, etc. This sample demonstrates another feature of templates: using -
	// in actions to trim whitespace.
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// range blocks let us loop through slices, arrays, maps or channels. Inside the
	// range block {{.}} is set to the current item of the iteration.
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})

	// Funcs adds custom functions, and | pipes a value into a function.
	t5 := template.Must(template.New("t5").
		Funcs(template.FuncMap{"upper": strings.ToUpper}).
		Parse("{{.Name | upper}} has {{len .Langs}} languages\n"))
	t5.Execute(os.Stdout, map[string]any{"Name": "tuck", "Langs": []string{"go", "sql"}})

	// text/template inserts a value as it is. html/template uses the same actions
	// and escapes the value so the result is safe to serve as HTML.
	raw := template.Must(template.New("raw").Parse("text: {{.}}\n"))
	raw.Execute(os.Stdout, "<b>")
	page := htmltemplate.Must(htmltemplate.New("page").Parse("html: {{.}}\n"))
	page.Execute(os.Stdout, "<b>")
}
