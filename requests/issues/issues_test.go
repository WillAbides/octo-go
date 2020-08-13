package issues_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/willabides/octo-go"
	"github.com/willabides/octo-go/internal/testutil"
	"github.com/willabides/octo-go/requests"
	"github.com/willabides/octo-go/requests/issues"
)

func vcrClient(t *testing.T, cas string, opts ...requests.Option) issues.Client {
	return testutil.VCRClient(t, cas, opts...).Issues()
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name(), testutil.PATAuth())
	req := &issues.GetReq{
		Owner:       "golang",
		Repo:        "go",
		IssueNumber: 1,
	}
	httpReq, err := req.HTTPRequest(ctx, client...)
	require.NoError(t, err)
	opts := requests.BuildOptions(client...)
	httpClient := opts.HttpClient()
	httpResp, err := httpClient.Do(httpReq)
	require.NoError(t, err)
	resp := new(issues.GetResponse)
	err = resp.Load(httpResp)
	require.NoError(t, err)
	require.Equal(t, int64(1), resp.Data.Number)
}

func TestAddLabels(t *testing.T) {
	t.Run("as_app", func(t *testing.T) {
		ctx := context.Background()
		vc := vcrClient(t, t.Name())
		authClient := octo.NewClient(vc...)
		authClient = append(authClient, testutil.AppAuth(t))
		client := octo.NewClient(vc...)
		client = append(client, testutil.AppInstallationAuth(authClient))

		_, err := client.Issues().AddLabels(ctx, &issues.AddLabelsReq{
			Owner:       "WillAbides",
			Repo:        "octo-go",
			IssueNumber: 12,
			RequestBody: issues.AddLabelsReqBody{
				Labels: []string{"testlabel2", "testlabel"},
			},
		})
		require.NoError(t, err)
	})
}

func TestCheckUserCanBeAssigned(t *testing.T) {
	ctx := context.Background()
	client := vcrClient(t, t.Name(), testutil.PATAuth())

	t.Run("true", func(t *testing.T) {
		result, err := client.CheckUserCanBeAssigned(ctx, &issues.CheckUserCanBeAssignedReq{
			Owner:    "WillAbides",
			Repo:     "octo-go",
			Assignee: "WillAbides",
		})
		require.NoError(t, err)
		require.True(t, result.Data)
	})

	t.Run("false", func(t *testing.T) {
		result, err := client.CheckUserCanBeAssigned(ctx, &issues.CheckUserCanBeAssignedReq{
			Owner:    "WillAbides",
			Repo:     "octo-go",
			Assignee: "defunkt",
		})
		require.NoError(t, err)
		require.False(t, result.Data)
	})
}

func TestListComments(t *testing.T) {
	four := int64(4)
	t.Run("paging", func(t *testing.T) {
		ctx := context.Background()
		client := vcrClient(t, t.Name(), testutil.PATAuth())
		var commentIDs []int64
		req := &issues.ListCommentsReq{
			Owner:       "golang",
			Repo:        "go",
			IssueNumber: 1,
			PerPage:     &four,
		}
		ok := true
		for ok {
			resp, err := client.ListComments(ctx, req)
			require.NoError(t, err)
			if resp.Data != nil {
				for _, r := range resp.Data {
					commentIDs = append(commentIDs, r.Id)
				}
			}
			ok = req.Rel("next", resp)
		}
		require.Len(t, commentIDs, 12)
	})
}
