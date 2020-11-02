package views

import "html/template"

func NewView(Layout string, files ...string) *View {
	files = append(files,
		"views/layouts/bootstrap.gohtml",
		"views/layouts/header.gohtml",
		"views/layouts/footer.gohtml",
		"views/layouts/navbar.gohtml",
	)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   Layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}
