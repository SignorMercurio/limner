package printer

import (
	"bytes"
	"testing"

	"github.com/SignorMercurio/limner/testutil"
)

func TestJsonPrinter(t *testing.T) {
	tests := []struct {
		name          string
		src           string
		dst           string
		onlyForceType bool
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
  "followers": 27,
  "following": 12,
  "created_at": "2017-10-05T10:41:16Z",
  "updated_at": "2021-10-11T19:48:46Z"
}
`,
			dst: `{
    [31mavatar_url[0m: "[32mhttps://avatars.githubusercontent.com/u/32540679?v=4[0m",
    [31mbio[0m: "[32mMSc[0m",
    [31mblog[0m: "[32mhttps://blog.sigmerc.top/[0m",
    [31mcompany[0m: [36mnull[0m,
    [31mcreated_at[0m: "[32m2017-10-05T10:41:16Z[0m",
    [31memail[0m: [36mnull[0m,
    [31mevents_url[0m: "[32mhttps://api.github.com/users/SignorMercurio/events{/privacy}[0m",
    [31mfollowers[0m: [33m27[0m,
    [31mfollowers_url[0m: "[32mhttps://api.github.com/users/SignorMercurio/followers[0m",
    [31mfollowing[0m: [33m12[0m,
    [31mfollowing_url[0m: "[32mhttps://api.github.com/users/SignorMercurio/following{/other_user}[0m",
    [31mgists_url[0m: "[32mhttps://api.github.com/users/SignorMercurio/gists{/gist_id}[0m",
    [31mgravatar_id[0m: "[32m[0m",
    [31mhireable[0m: [36mnull[0m,
    [31mhtml_url[0m: "[32mhttps://github.com/SignorMercurio[0m",
    [31mid[0m: [33m32540679[0m,
    [31mlocation[0m: [36mnull[0m,
    [31mlogin[0m: "[32mSignorMercurio[0m",
    [31mname[0m: "[32mMercurio[0m",
    [31mnode_id[0m: "[32mMDQ6VXNlcjMyNTQwNjc5[0m",
    [31morganizations_url[0m: "[32mhttps://api.github.com/users/SignorMercurio/orgs[0m",
    [31mpublic_gists[0m: [33m0[0m,
    [31mpublic_repos[0m: [33m33[0m,
    [31mreceived_events_url[0m: "[32mhttps://api.github.com/users/SignorMercurio/received_events[0m",
    [31mrepos_url[0m: "[32mhttps://api.github.com/users/SignorMercurio/repos[0m",
    [31msite_admin[0m: [33mfalse[0m,
    [31mstarred_url[0m: "[32mhttps://api.github.com/users/SignorMercurio/starred{/owner}{/repo}[0m",
    [31msubscriptions_url[0m: "[32mhttps://api.github.com/users/SignorMercurio/subscriptions[0m",
    [31mtwitter_username[0m: [36mnull[0m,
    [31mtype[0m: "[32mUser[0m",
    [31mupdated_at[0m: "[32m2021-10-11T19:48:46Z[0m",
    [31murl[0m: "[32mhttps://api.github.com/users/SignorMercurio[0m"
}
`,
		},
		{
			name: "strange cases",
			src: `{
  "array": [
    "item1",
    "item2"
  ]
}`,
			dst: `{
    [31marray[0m: [
        "[32mitem1[0m",
        "[32mitem2[0m"
    ]
}
`,
		},
		{
			name: "unmarshal failed",
			src:  `hello`,
			dst: `[33mFailed to unmarshal json, using default printer[0m
[32mhello[0m`,
			onlyForceType: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var w bytes.Buffer

			if !tt.onlyForceType {
				p := &ColorPrinter{}
				p.Print(tt.src, &w)
				testutil.MustEqual(t, tt.dst, w.String())
				w.Reset()
			}

			p := &ColorPrinter{Type: "json"}
			p.Print(tt.src, &w)
			testutil.MustEqual(t, tt.dst, w.String())
		})
	}
}
