package octo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go"
)

func TestIssuesAddLabels(t *testing.T) {
	t.Run("as_app", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name(), appInstallationAuth(t))
		_, err := client.IssuesAddLabels(ctx, &octo.IssuesAddLabelsReq{
			Owner:       "WillAbides",
			Repo:        "octo-go",
			IssueNumber: 12,
			RequestBody: octo.IssuesAddLabelsReqBody{
				Labels: []string{"testlabel2", "testlabel"},
			},
		})
		require.NoError(t, err)
	})
}

func TestIssuesCheckAssignee(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name(), patAuth())

	t.Run("true", func(t *testing.T) {
		result, err := client.IssuesCheckAssignee(ctx, &octo.IssuesCheckAssigneeReq{
			Owner:    "WillAbides",
			Repo:     "octo-go",
			Assignee: "WillAbides",
		})
		require.NoError(t, err)
		require.True(t, result.Data)
	})

	t.Run("false", func(t *testing.T) {
		result, err := client.IssuesCheckAssignee(ctx, &octo.IssuesCheckAssigneeReq{
			Owner:    "WillAbides",
			Repo:     "octo-go",
			Assignee: "defunkt",
		})
		require.NoError(t, err)
		require.False(t, result.Data)
	})
}

func TestIssuesListComments(t *testing.T) {
	t.Run("paging", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name(), patAuth())
		var commentIDs []int64
		req := &octo.IssuesListCommentsReq{
			Owner:       "golang",
			Repo:        "go",
			IssueNumber: 1,
			PerPage:     octo.Int64(4),
		}
		ok := true
		for ok {
			resp, err := client.IssuesListComments(ctx, req)
			require.NoError(t, err)
			if resp.Data != nil {
				for _, r := range *resp.Data {
					commentIDs = append(commentIDs, r.Id)
				}
			}
			ok = req.Rel(octo.RelNext, resp)
		}
		require.Len(t, commentIDs, 12)
	})
}
