package database

import "testing"

func TestGithubCreate(t *testing.T) {
	DatabaseInitTest()
	GithubWebhookCreate("dingdinglz/dingbot", "114997885")
}
