package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleTag  = "Title: "
	descTag   = "Description: "
	tagsTag   = "Tags: "
	bodyDelim = "---"
)

func NewPostFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := GetPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPost(fileSystem fs.FS, name string) (post Post, err error) {
	postFile, err := fileSystem.Open(name)
	post = Post{}
	if err != nil {
		return
	}
	defer postFile.Close()

	return NewPost(postFile)
}

func NewPost(postFile io.Reader) (post Post, err error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	post = Post{
		Title:       readMetaLine(titleTag),
		Description: readMetaLine(descTag),
		Tags:        strings.Split(readMetaLine(tagsTag), ", "),
		Body:        readBody(scanner)}
	return
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	body := strings.TrimSuffix(buf.String(), "\n")
	return body
}
