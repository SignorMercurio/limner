package transformer

import (
	"testing"

	"github.com/SignorMercurio/limner/testutil"
)

func TestYamlTransformerFromJson(t *testing.T) {
	tests := []struct {
		name string
		src  string
		dst  string
	}{
		{
			name: "curl -s https://api.github.com/users/SignorMercurio",
			src: `{
  "login": "SignorMercurio",
  "id": 32540679,
  "node_id": "MDQ6VXNlcjMyNTQwNjc5",
  "avatar_url": "https://avatars.githubusercontent.com/u/32540679?v=4",
  "gravatar_id": "",
  "url": "https://api.github.com/users/SignorMercurio",
  "html_url": "https://github.com/SignorMercurio",
  "followers_url": "https://api.github.com/users/SignorMercurio/followers",
  "following_url": "https://api.github.com/users/SignorMercurio/following{/other_user}",
  "gists_url": "https://api.github.com/users/SignorMercurio/gists{/gist_id}",
  "starred_url": "https://api.github.com/users/SignorMercurio/starred{/owner}{/repo}",
  "subscriptions_url": "https://api.github.com/users/SignorMercurio/subscriptions",
  "organizations_url": "https://api.github.com/users/SignorMercurio/orgs",
  "repos_url": "https://api.github.com/users/SignorMercurio/repos",
  "events_url": "https://api.github.com/users/SignorMercurio/events{/privacy}",
  "received_events_url": "https://api.github.com/users/SignorMercurio/received_events",
  "type": "User",
  "site_admin": false,
  "name": "Mercurio",
  "company": null,
  "blog": "https://blog.sigmerc.top/",
  "location": null,
  "email": null,
  "hireable": null,
  "bio": "MSc",
  "twitter_username": null,
  "public_repos": 33,
  "public_gists": 0,
  "followers": 28,
  "following": 12,
  "created_at": "2017-10-05T10:41:16Z",
  "updated_at": "2021-10-14T21:35:16Z"
}
`,
			dst: `avatar_url: https://avatars.githubusercontent.com/u/32540679?v=4
bio: MSc
blog: https://blog.sigmerc.top/
company: null
created_at: "2017-10-05T10:41:16Z"
email: null
events_url: https://api.github.com/users/SignorMercurio/events{/privacy}
followers: 28
followers_url: https://api.github.com/users/SignorMercurio/followers
following: 12
following_url: https://api.github.com/users/SignorMercurio/following{/other_user}
gists_url: https://api.github.com/users/SignorMercurio/gists{/gist_id}
gravatar_id: ""
hireable: null
html_url: https://github.com/SignorMercurio
id: 3.2540679e+07
location: null
login: SignorMercurio
name: Mercurio
node_id: MDQ6VXNlcjMyNTQwNjc5
organizations_url: https://api.github.com/users/SignorMercurio/orgs
public_gists: 0
public_repos: 33
received_events_url: https://api.github.com/users/SignorMercurio/received_events
repos_url: https://api.github.com/users/SignorMercurio/repos
site_admin: false
starred_url: https://api.github.com/users/SignorMercurio/starred{/owner}{/repo}
subscriptions_url: https://api.github.com/users/SignorMercurio/subscriptions
twitter_username: null
type: User
updated_at: "2021-10-14T21:35:16Z"
url: https://api.github.com/users/SignorMercurio
`,
		},
		{
			name: "unmarshal json failed",
			src: `{
a: b,
}`,
			dst: `{
a: b,
}`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			trans := &FormatTransformer{OutType: "yaml"}
			w, _ := trans.Transform([]byte(tt.src))
			testutil.MustEqual(t, []byte(tt.dst), w)

			trans = &FormatTransformer{InType: "json", OutType: "yaml"}
			w, _ = trans.Transform([]byte(tt.src))
			testutil.MustEqual(t, []byte(tt.dst), w)
		})
	}
}
