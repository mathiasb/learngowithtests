package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

type postViewModel struct {
	post Post
	html template.HTML
}

func newPostViewModel(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{post: p}
	vm.html = template.HTML(markdown.ToHTML([]byte(p.Body), r.parser, nil))
	return vm
}

func (p Post) ViewTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

var (
	//go:embed "templates/*"
	postTemplate embed.FS
)

type PostRenderer struct {
	templ  *template.Template
	parser *parser.Parser
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, parser: parser}, nil
}

func (r *PostRenderer) Render(writer io.Writer, post Post) error {

	if err := r.templ.ExecuteTemplate(writer, "blog.gohtml", newPostViewModel(post, r)); err != nil {
		return err
	}

	if post.Body != "" {
		body_html := markdown.ToHTML([]byte(post.Body), nil, nil)
		post.Body = string(body_html)
	}

	return nil
}

func (r *PostRenderer) RenderIndex(writer io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(writer, "index.gohtml", posts)
}
