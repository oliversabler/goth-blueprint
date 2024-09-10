package services

import "time"

type Post struct {
	Title   string
	Author  string
	URI     string
	Content string
	Date    time.Time
}

func NewPost(title string, author string, uri string, content string) Post {
	return Post{
		Title:   title,
		Author:  author,
		URI:     uri,
		Content: content,
		Date:    time.Now().UTC(),
	}
}

func GetPosts() []Post {
	return []Post{
		NewPost(
			"Welcome to GoTH",
			"Me",
			"/b/welcome",
			"Write a blogpost here.",
		),
		NewPost(
			"These are just examples",
			"Me",
			"/b/examples",
			"A second post, how nice.",
		),
		NewPost(
			"A database maybe?",
			"Me",
			"/b/database",
			"Replace this with fetching blog posts from a database, if you want.",
		),
	}
}
