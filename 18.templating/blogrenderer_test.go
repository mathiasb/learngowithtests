package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	blogrenderer "github.com/mathiasb/learngowithtests/18.templating"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "Hello, world!",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post to html", func(t *testing.T) {
		buf := bytes.Buffer{}
		blogrenderer, _ := blogrenderer.NewPostRenderer()

		if err := blogrenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello world"}, {Title: "Hello world 2"}}
		postR, err := blogrenderer.NewPostRenderer()

		if err != nil {
			t.Fatal("could not create a renderer")
		}

		if err := postR.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<ol><li><a href="/post/hello-world">Hello world</a></li><li><a href="/post/hello-world-2">Hello world 2</a></li></ol>`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	blogrenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		b.Fatal("could not create a renderer")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blogrenderer.Render(io.Discard, aPost)
	}
}
