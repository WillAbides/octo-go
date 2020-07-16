package octotest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go"
	"github.com/willabides/octo-go/components"
)

func TestPaging(t *testing.T) {
	req1 := &octo.IssuesListForRepoReq{
		Owner:   "golang",
		Repo:    "go",
		Labels:  octo.String("release-blocker"),
		PerPage: octo.Int64(2),
	}
	req2 := &octo.IssuesListForRepoReq{
		Owner:   "golang",
		Repo:    "go",
		Labels:  octo.String("release-blocker"),
		Page:    octo.Int64(2),
		PerPage: octo.Int64(2),
	}
	req3 := &octo.IssuesListForRepoReq{
		Owner:   "golang",
		Repo:    "go",
		Labels:  octo.String("release-blocker"),
		Page:    octo.Int64(3),
		PerPage: octo.Int64(2),
	}
	res1 := &octo.IssuesListForRepoResponseBody{
		{components.IssueSimple2{Id: 1}},
		{components.IssueSimple2{Id: 2}},
	}
	res2 := &octo.IssuesListForRepoResponseBody{
		{components.IssueSimple2{Id: 3}},
		{components.IssueSimple2{Id: 4}},
	}
	res3 := &octo.IssuesListForRepoResponseBody{
		{components.IssueSimple2{Id: 5}},
		{components.IssueSimple2{Id: 6}},
	}
	server := New(octo.PreserveResponseBody())
	t.Cleanup(server.Finish)
	server.Expect(req1, RelLinkHandler(octo.RelNext, JSONResponder(200, res1), req2, server))
	server.Expect(req2, RelLinkHandler(octo.RelNext, JSONResponder(200, res2), req3, server))
	server.Expect(req3, JSONResponder(200, res3))
	ctx := context.Background()
	client := server.Client()
	ok := true
	req := req1
	var got []int64
	for ok {
		resp, err := client.IssuesListForRepo(ctx, req)
		require.NoError(t, err)
		for _, data := range *resp.Data {
			got = append(got, data.Id)
		}
		ok = req.Rel(octo.RelNext, resp)
	}
	want := []int64{1, 2, 3, 4, 5, 6}
	require.Equal(t, want, got)
}

func TestDistinguishesBodies(t *testing.T) {
	req1 := &octo.ChecksCreateReq{
		Owner: "foo",
		Repo:  "bar",
		RequestBody: octo.ChecksCreateReqBody{
			Name:    octo.String("name 1"),
			HeadSha: octo.String("deadbeef"),
		},
	}
	req2 := &octo.ChecksCreateReq{
		Owner: "foo",
		Repo:  "bar",
		RequestBody: octo.ChecksCreateReqBody{
			Name:    octo.String("name 2"),
			HeadSha: octo.String("deadbeef"),
		},
	}
	respBody1 := &octo.ChecksCreateResponseBody{
		CheckRun: components.CheckRun{
			Conclusion: "conclusion 1",
		},
	}
	respBody2 := &octo.ChecksCreateResponseBody{
		CheckRun: components.CheckRun{
			Conclusion: "conclusion 2",
		},
	}
	ctx := context.Background()
	server := New(octo.PreserveResponseBody())
	t.Cleanup(server.Finish)
	server.Expect(req2, JSONResponder(201, respBody2))
	server.Expect(req1, JSONResponder(201, respBody1))
	client := server.Client()
	got1, err := client.ChecksCreate(ctx, req1)
	require.NoError(t, err)
	require.Equal(t, respBody1, got1.Data)
	got2, err := client.ChecksCreate(ctx, req2)
	require.NoError(t, err)
	require.Equal(t, respBody2, got2.Data)
}
