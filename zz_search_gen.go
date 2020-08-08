// Code generated by octo-go; DO NOT EDIT.

package octo

import (
	"context"
	"fmt"
	components "github.com/willabides/octo-go/components"
	"net/http"
	"net/url"
	"strconv"
)

/*
SearchCode performs requests for "search/code"

Search code.

  GET /search/code

https://developer.github.com/v3/search/#search-code
*/
func SearchCode(ctx context.Context, req *SearchCodeReq, opt ...RequestOption) (*SearchCodeResponse, error) {
	if req == nil {
		req = new(SearchCodeReq)
	}
	resp := &SearchCodeResponse{request: req}
	r, err := doRequest(ctx, req, "search/code", opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = SearchCodeResponseBody{}
	err = r.decodeBody(&resp.Data, "search/code")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
SearchCode performs requests for "search/code"

Search code.

  GET /search/code

https://developer.github.com/v3/search/#search-code
*/
func (c Client) SearchCode(ctx context.Context, req *SearchCodeReq, opt ...RequestOption) (*SearchCodeResponse, error) {
	return SearchCode(ctx, req, append(c, opt...)...)
}

/*
SearchCodeReq is request data for Client.SearchCode

https://developer.github.com/v3/search/#search-code
*/
type SearchCodeReq struct {
	_url string

	/*
	The query contains one or more search keywords and qualifiers. Qualifiers allow
	you to limit your search to specific areas of GitHub. The REST API supports the
	same qualifiers as GitHub.com. To learn more about the format of the query, see
	[Constructing a search
	query](https://developer.github.com/v3/search/#constructing-a-search-query). See
	"[Searching code](https://help.github.com/articles/searching-code/)" for a
	detailed list of qualifiers.
	*/
	Q *string

	/*
	Sorts the results of your query. Can only be `indexed`, which indicates how
	recently a file has been indexed by the GitHub search infrastructure. Default:
	[best match](https://developer.github.com/v3/search/#ranking-search-results)
	*/
	Sort *string

	/*
	Determines whether the first search result returned is the highest number of
	matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored
	unless you provide `sort`.
	*/
	Order *string

	// Results per page (max 100)
	PerPage *int64

	// Page number of the results to fetch.
	Page *int64
}

func (r *SearchCodeReq) url() string {
	return r._url
}

func (r *SearchCodeReq) urlPath() string {
	return fmt.Sprintf("/search/code")
}

func (r *SearchCodeReq) method() string {
	return "GET"
}

func (r *SearchCodeReq) urlQuery() url.Values {
	query := url.Values{}
	if r.Q != nil {
		query.Set("q", *r.Q)
	}
	if r.Sort != nil {
		query.Set("sort", *r.Sort)
	}
	if r.Order != nil {
		query.Set("order", *r.Order)
	}
	if r.PerPage != nil {
		query.Set("per_page", strconv.FormatInt(*r.PerPage, 10))
	}
	if r.Page != nil {
		query.Set("page", strconv.FormatInt(*r.Page, 10))
	}
	return query
}

func (r *SearchCodeReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *SearchCodeReq) body() interface{} {
	return nil
}

func (r *SearchCodeReq) dataStatuses() []int {
	return []int{200}
}

func (r *SearchCodeReq) validStatuses() []int {
	return []int{200, 304}
}

// HTTPRequest builds an *http.Request
func (r *SearchCodeReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, "search/code", opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *SearchCodeReq) Rel(link RelName, resp *SearchCodeResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
SearchCodeResponseBody is a response body for SearchCode

https://developer.github.com/v3/search/#search-code
*/
type SearchCodeResponseBody struct {
	IncompleteResults bool                              `json:"incomplete_results,omitempty"`
	Items             []components.CodeSearchResultItem `json:"items,omitempty"`
	TotalCount        int64                             `json:"total_count,omitempty"`
}

/*
SearchCodeResponse is a response for SearchCode

https://developer.github.com/v3/search/#search-code
*/
type SearchCodeResponse struct {
	response
	request *SearchCodeReq
	Data    SearchCodeResponseBody
}

/*
SearchCommits performs requests for "search/commits"

Search commits.

  GET /search/commits

https://developer.github.com/v3/search/#search-commits
*/
func SearchCommits(ctx context.Context, req *SearchCommitsReq, opt ...RequestOption) (*SearchCommitsResponse, error) {
	if req == nil {
		req = new(SearchCommitsReq)
	}
	resp := &SearchCommitsResponse{request: req}
	r, err := doRequest(ctx, req, "search/commits", opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = SearchCommitsResponseBody{}
	err = r.decodeBody(&resp.Data, "search/commits")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
SearchCommits performs requests for "search/commits"

Search commits.

  GET /search/commits

https://developer.github.com/v3/search/#search-commits
*/
func (c Client) SearchCommits(ctx context.Context, req *SearchCommitsReq, opt ...RequestOption) (*SearchCommitsResponse, error) {
	return SearchCommits(ctx, req, append(c, opt...)...)
}

/*
SearchCommitsReq is request data for Client.SearchCommits

https://developer.github.com/v3/search/#search-commits
*/
type SearchCommitsReq struct {
	_url string

	/*
	The query contains one or more search keywords and qualifiers. Qualifiers allow
	you to limit your search to specific areas of GitHub. The REST API supports the
	same qualifiers as GitHub.com. To learn more about the format of the query, see
	[Constructing a search
	query](https://developer.github.com/v3/search/#constructing-a-search-query). See
	"[Searching commits](https://help.github.com/articles/searching-commits/)" for a
	detailed list of qualifiers.
	*/
	Q *string

	/*
	Sorts the results of your query by `author-date` or `committer-date`. Default:
	[best match](https://developer.github.com/v3/search/#ranking-search-results)
	*/
	Sort *string

	/*
	Determines whether the first search result returned is the highest number of
	matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored
	unless you provide `sort`.
	*/
	Order *string

	// Results per page (max 100)
	PerPage *int64

	// Page number of the results to fetch.
	Page *int64

	/*
	The Commit Search API is currently available for developers to preview. During
	the preview period, the APIs may change without advance notice. Please see the
	[blog post](https://developer.github.com/changes/2017-01-05-commit-search-api/)
	for full details.

	To access the API you must set this to true.
	*/
	CloakPreview bool
}

func (r *SearchCommitsReq) url() string {
	return r._url
}

func (r *SearchCommitsReq) urlPath() string {
	return fmt.Sprintf("/search/commits")
}

func (r *SearchCommitsReq) method() string {
	return "GET"
}

func (r *SearchCommitsReq) urlQuery() url.Values {
	query := url.Values{}
	if r.Q != nil {
		query.Set("q", *r.Q)
	}
	if r.Sort != nil {
		query.Set("sort", *r.Sort)
	}
	if r.Order != nil {
		query.Set("order", *r.Order)
	}
	if r.PerPage != nil {
		query.Set("per_page", strconv.FormatInt(*r.PerPage, 10))
	}
	if r.Page != nil {
		query.Set("page", strconv.FormatInt(*r.Page, 10))
	}
	return query
}

func (r *SearchCommitsReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{"cloak": r.CloakPreview}
	if requiredPreviews {
		previewVals["cloak"] = true
	}
	if allPreviews {
		previewVals["cloak"] = true
	}
	return requestHeaders(headerVals, previewVals)
}

func (r *SearchCommitsReq) body() interface{} {
	return nil
}

func (r *SearchCommitsReq) dataStatuses() []int {
	return []int{200}
}

func (r *SearchCommitsReq) validStatuses() []int {
	return []int{200, 304}
}

// HTTPRequest builds an *http.Request
func (r *SearchCommitsReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, "search/commits", opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *SearchCommitsReq) Rel(link RelName, resp *SearchCommitsResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
SearchCommitsResponseBody is a response body for SearchCommits

https://developer.github.com/v3/search/#search-commits
*/
type SearchCommitsResponseBody struct {
	IncompleteResults bool                                `json:"incomplete_results,omitempty"`
	Items             []components.CommitSearchResultItem `json:"items,omitempty"`
	TotalCount        int64                               `json:"total_count,omitempty"`
}

/*
SearchCommitsResponse is a response for SearchCommits

https://developer.github.com/v3/search/#search-commits
*/
type SearchCommitsResponse struct {
	response
	request *SearchCommitsReq
	Data    SearchCommitsResponseBody
}

/*
SearchIssuesAndPullRequests performs requests for "search/issues-and-pull-requests"

Search issues and pull requests.

  GET /search/issues

https://developer.github.com/v3/search/#search-issues-and-pull-requests
*/
func SearchIssuesAndPullRequests(ctx context.Context, req *SearchIssuesAndPullRequestsReq, opt ...RequestOption) (*SearchIssuesAndPullRequestsResponse, error) {
	if req == nil {
		req = new(SearchIssuesAndPullRequestsReq)
	}
	resp := &SearchIssuesAndPullRequestsResponse{request: req}
	r, err := doRequest(ctx, req, "search/issues-and-pull-requests", opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = SearchIssuesAndPullRequestsResponseBody{}
	err = r.decodeBody(&resp.Data, "search/issues-and-pull-requests")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
SearchIssuesAndPullRequests performs requests for "search/issues-and-pull-requests"

Search issues and pull requests.

  GET /search/issues

https://developer.github.com/v3/search/#search-issues-and-pull-requests
*/
func (c Client) SearchIssuesAndPullRequests(ctx context.Context, req *SearchIssuesAndPullRequestsReq, opt ...RequestOption) (*SearchIssuesAndPullRequestsResponse, error) {
	return SearchIssuesAndPullRequests(ctx, req, append(c, opt...)...)
}

/*
SearchIssuesAndPullRequestsReq is request data for Client.SearchIssuesAndPullRequests

https://developer.github.com/v3/search/#search-issues-and-pull-requests
*/
type SearchIssuesAndPullRequestsReq struct {
	_url string

	/*
	The query contains one or more search keywords and qualifiers. Qualifiers allow
	you to limit your search to specific areas of GitHub. The REST API supports the
	same qualifiers as GitHub.com. To learn more about the format of the query, see
	[Constructing a search
	query](https://developer.github.com/v3/search/#constructing-a-search-query). See
	"[Searching issues and pull
	requests](https://help.github.com/articles/searching-issues-and-pull-requests/)"
	for a detailed list of qualifiers.
	*/
	Q *string

	/*
	Sorts the results of your query by the number of `comments`, `reactions`,
	`reactions-+1`, `reactions--1`, `reactions-smile`, `reactions-thinking_face`,
	`reactions-heart`, `reactions-tada`, or `interactions`. You can also sort
	results by how recently the items were `created` or `updated`, Default: [best
	match](https://developer.github.com/v3/search/#ranking-search-results)
	*/
	Sort *string

	/*
	Determines whether the first search result returned is the highest number of
	matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored
	unless you provide `sort`.
	*/
	Order *string

	// Results per page (max 100)
	PerPage *int64

	// Page number of the results to fetch.
	Page *int64
}

func (r *SearchIssuesAndPullRequestsReq) url() string {
	return r._url
}

func (r *SearchIssuesAndPullRequestsReq) urlPath() string {
	return fmt.Sprintf("/search/issues")
}

func (r *SearchIssuesAndPullRequestsReq) method() string {
	return "GET"
}

func (r *SearchIssuesAndPullRequestsReq) urlQuery() url.Values {
	query := url.Values{}
	if r.Q != nil {
		query.Set("q", *r.Q)
	}
	if r.Sort != nil {
		query.Set("sort", *r.Sort)
	}
	if r.Order != nil {
		query.Set("order", *r.Order)
	}
	if r.PerPage != nil {
		query.Set("per_page", strconv.FormatInt(*r.PerPage, 10))
	}
	if r.Page != nil {
		query.Set("page", strconv.FormatInt(*r.Page, 10))
	}
	return query
}

func (r *SearchIssuesAndPullRequestsReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *SearchIssuesAndPullRequestsReq) body() interface{} {
	return nil
}

func (r *SearchIssuesAndPullRequestsReq) dataStatuses() []int {
	return []int{200}
}

func (r *SearchIssuesAndPullRequestsReq) validStatuses() []int {
	return []int{200, 304}
}

// HTTPRequest builds an *http.Request
func (r *SearchIssuesAndPullRequestsReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, "search/issues-and-pull-requests", opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *SearchIssuesAndPullRequestsReq) Rel(link RelName, resp *SearchIssuesAndPullRequestsResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
SearchIssuesAndPullRequestsResponseBody is a response body for SearchIssuesAndPullRequests

https://developer.github.com/v3/search/#search-issues-and-pull-requests
*/
type SearchIssuesAndPullRequestsResponseBody struct {
	IncompleteResults bool                               `json:"incomplete_results,omitempty"`
	Items             []components.IssueSearchResultItem `json:"items,omitempty"`
	TotalCount        int64                              `json:"total_count,omitempty"`
}

/*
SearchIssuesAndPullRequestsResponse is a response for SearchIssuesAndPullRequests

https://developer.github.com/v3/search/#search-issues-and-pull-requests
*/
type SearchIssuesAndPullRequestsResponse struct {
	response
	request *SearchIssuesAndPullRequestsReq
	Data    SearchIssuesAndPullRequestsResponseBody
}

/*
SearchLabels performs requests for "search/labels"

Search labels.

  GET /search/labels

https://developer.github.com/v3/search/#search-labels
*/
func SearchLabels(ctx context.Context, req *SearchLabelsReq, opt ...RequestOption) (*SearchLabelsResponse, error) {
	if req == nil {
		req = new(SearchLabelsReq)
	}
	resp := &SearchLabelsResponse{request: req}
	r, err := doRequest(ctx, req, "search/labels", opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = SearchLabelsResponseBody{}
	err = r.decodeBody(&resp.Data, "search/labels")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
SearchLabels performs requests for "search/labels"

Search labels.

  GET /search/labels

https://developer.github.com/v3/search/#search-labels
*/
func (c Client) SearchLabels(ctx context.Context, req *SearchLabelsReq, opt ...RequestOption) (*SearchLabelsResponse, error) {
	return SearchLabels(ctx, req, append(c, opt...)...)
}

/*
SearchLabelsReq is request data for Client.SearchLabels

https://developer.github.com/v3/search/#search-labels
*/
type SearchLabelsReq struct {
	_url string

	// The id of the repository.
	RepositoryId *int64

	/*
	The search keywords. This endpoint does not accept qualifiers in the query. To
	learn more about the format of the query, see [Constructing a search
	query](https://developer.github.com/v3/search/#constructing-a-search-query).
	*/
	Q *string

	/*
	Sorts the results of your query by when the label was `created` or `updated`.
	Default: [best
	match](https://developer.github.com/v3/search/#ranking-search-results)
	*/
	Sort *string

	/*
	Determines whether the first search result returned is the highest number of
	matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored
	unless you provide `sort`.
	*/
	Order *string
}

func (r *SearchLabelsReq) url() string {
	return r._url
}

func (r *SearchLabelsReq) urlPath() string {
	return fmt.Sprintf("/search/labels")
}

func (r *SearchLabelsReq) method() string {
	return "GET"
}

func (r *SearchLabelsReq) urlQuery() url.Values {
	query := url.Values{}
	if r.RepositoryId != nil {
		query.Set("repository_id", strconv.FormatInt(*r.RepositoryId, 10))
	}
	if r.Q != nil {
		query.Set("q", *r.Q)
	}
	if r.Sort != nil {
		query.Set("sort", *r.Sort)
	}
	if r.Order != nil {
		query.Set("order", *r.Order)
	}
	return query
}

func (r *SearchLabelsReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *SearchLabelsReq) body() interface{} {
	return nil
}

func (r *SearchLabelsReq) dataStatuses() []int {
	return []int{200}
}

func (r *SearchLabelsReq) validStatuses() []int {
	return []int{200, 304}
}

// HTTPRequest builds an *http.Request
func (r *SearchLabelsReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, "search/labels", opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *SearchLabelsReq) Rel(link RelName, resp *SearchLabelsResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
SearchLabelsResponseBody is a response body for SearchLabels

https://developer.github.com/v3/search/#search-labels
*/
type SearchLabelsResponseBody struct {
	IncompleteResults bool                               `json:"incomplete_results,omitempty"`
	Items             []components.LabelSearchResultItem `json:"items,omitempty"`
	TotalCount        int64                              `json:"total_count,omitempty"`
}

/*
SearchLabelsResponse is a response for SearchLabels

https://developer.github.com/v3/search/#search-labels
*/
type SearchLabelsResponse struct {
	response
	request *SearchLabelsReq
	Data    SearchLabelsResponseBody
}

/*
SearchRepos performs requests for "search/repos"

Search repositories.

  GET /search/repositories

https://developer.github.com/v3/search/#search-repositories
*/
func SearchRepos(ctx context.Context, req *SearchReposReq, opt ...RequestOption) (*SearchReposResponse, error) {
	if req == nil {
		req = new(SearchReposReq)
	}
	resp := &SearchReposResponse{request: req}
	r, err := doRequest(ctx, req, "search/repos", opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = SearchReposResponseBody{}
	err = r.decodeBody(&resp.Data, "search/repos")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
SearchRepos performs requests for "search/repos"

Search repositories.

  GET /search/repositories

https://developer.github.com/v3/search/#search-repositories
*/
func (c Client) SearchRepos(ctx context.Context, req *SearchReposReq, opt ...RequestOption) (*SearchReposResponse, error) {
	return SearchRepos(ctx, req, append(c, opt...)...)
}

/*
SearchReposReq is request data for Client.SearchRepos

https://developer.github.com/v3/search/#search-repositories
*/
type SearchReposReq struct {
	_url string

	/*
	The query contains one or more search keywords and qualifiers. Qualifiers allow
	you to limit your search to specific areas of GitHub. The REST API supports the
	same qualifiers as GitHub.com. To learn more about the format of the query, see
	[Constructing a search
	query](https://developer.github.com/v3/search/#constructing-a-search-query). See
	"[Searching for
	repositories](https://help.github.com/articles/searching-for-repositories/)" for
	a detailed list of qualifiers.
	*/
	Q *string

	/*
	Sorts the results of your query by number of `stars`, `forks`, or
	`help-wanted-issues` or how recently the items were `updated`. Default: [best
	match](https://developer.github.com/v3/search/#ranking-search-results)
	*/
	Sort *string

	/*
	Determines whether the first search result returned is the highest number of
	matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored
	unless you provide `sort`.
	*/
	Order *string

	// Results per page (max 100)
	PerPage *int64

	// Page number of the results to fetch.
	Page *int64

	/*
	The `topics` property for repositories on GitHub is currently available for
	developers to preview. To view the `topics` property in calls that return
	repository results, you must set this to true.
	*/
	MercyPreview bool
}

func (r *SearchReposReq) url() string {
	return r._url
}

func (r *SearchReposReq) urlPath() string {
	return fmt.Sprintf("/search/repositories")
}

func (r *SearchReposReq) method() string {
	return "GET"
}

func (r *SearchReposReq) urlQuery() url.Values {
	query := url.Values{}
	if r.Q != nil {
		query.Set("q", *r.Q)
	}
	if r.Sort != nil {
		query.Set("sort", *r.Sort)
	}
	if r.Order != nil {
		query.Set("order", *r.Order)
	}
	if r.PerPage != nil {
		query.Set("per_page", strconv.FormatInt(*r.PerPage, 10))
	}
	if r.Page != nil {
		query.Set("page", strconv.FormatInt(*r.Page, 10))
	}
	return query
}

func (r *SearchReposReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{"mercy": r.MercyPreview}
	if allPreviews {
		previewVals["mercy"] = true
	}
	return requestHeaders(headerVals, previewVals)
}

func (r *SearchReposReq) body() interface{} {
	return nil
}

func (r *SearchReposReq) dataStatuses() []int {
	return []int{200}
}

func (r *SearchReposReq) validStatuses() []int {
	return []int{200, 304}
}

// HTTPRequest builds an *http.Request
func (r *SearchReposReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, "search/repos", opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *SearchReposReq) Rel(link RelName, resp *SearchReposResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
SearchReposResponseBody is a response body for SearchRepos

https://developer.github.com/v3/search/#search-repositories
*/
type SearchReposResponseBody struct {
	IncompleteResults bool                              `json:"incomplete_results,omitempty"`
	Items             []components.RepoSearchResultItem `json:"items,omitempty"`
	TotalCount        int64                             `json:"total_count,omitempty"`
}

/*
SearchReposResponse is a response for SearchRepos

https://developer.github.com/v3/search/#search-repositories
*/
type SearchReposResponse struct {
	response
	request *SearchReposReq
	Data    SearchReposResponseBody
}

/*
SearchTopics performs requests for "search/topics"

Search topics.

  GET /search/topics

https://developer.github.com/v3/search/#search-topics
*/
func SearchTopics(ctx context.Context, req *SearchTopicsReq, opt ...RequestOption) (*SearchTopicsResponse, error) {
	if req == nil {
		req = new(SearchTopicsReq)
	}
	resp := &SearchTopicsResponse{request: req}
	r, err := doRequest(ctx, req, "search/topics", opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = SearchTopicsResponseBody{}
	err = r.decodeBody(&resp.Data, "search/topics")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
SearchTopics performs requests for "search/topics"

Search topics.

  GET /search/topics

https://developer.github.com/v3/search/#search-topics
*/
func (c Client) SearchTopics(ctx context.Context, req *SearchTopicsReq, opt ...RequestOption) (*SearchTopicsResponse, error) {
	return SearchTopics(ctx, req, append(c, opt...)...)
}

/*
SearchTopicsReq is request data for Client.SearchTopics

https://developer.github.com/v3/search/#search-topics
*/
type SearchTopicsReq struct {
	_url string

	/*
	The query contains one or more search keywords and qualifiers. Qualifiers allow
	you to limit your search to specific areas of GitHub. The REST API supports the
	same qualifiers as GitHub.com. To learn more about the format of the query, see
	[Constructing a search
	query](https://developer.github.com/v3/search/#constructing-a-search-query).
	*/
	Q *string

	/*
	The `topics` property for repositories on GitHub is currently available for
	developers to preview. To view the `topics` property in calls that return
	repository results, you must set this to true.
	*/
	MercyPreview bool
}

func (r *SearchTopicsReq) url() string {
	return r._url
}

func (r *SearchTopicsReq) urlPath() string {
	return fmt.Sprintf("/search/topics")
}

func (r *SearchTopicsReq) method() string {
	return "GET"
}

func (r *SearchTopicsReq) urlQuery() url.Values {
	query := url.Values{}
	if r.Q != nil {
		query.Set("q", *r.Q)
	}
	return query
}

func (r *SearchTopicsReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{"mercy": r.MercyPreview}
	if requiredPreviews {
		previewVals["mercy"] = true
	}
	if allPreviews {
		previewVals["mercy"] = true
	}
	return requestHeaders(headerVals, previewVals)
}

func (r *SearchTopicsReq) body() interface{} {
	return nil
}

func (r *SearchTopicsReq) dataStatuses() []int {
	return []int{200}
}

func (r *SearchTopicsReq) validStatuses() []int {
	return []int{200, 304}
}

// HTTPRequest builds an *http.Request
func (r *SearchTopicsReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, "search/topics", opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *SearchTopicsReq) Rel(link RelName, resp *SearchTopicsResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
SearchTopicsResponseBody is a response body for SearchTopics

https://developer.github.com/v3/search/#search-topics
*/
type SearchTopicsResponseBody struct {
	IncompleteResults bool                               `json:"incomplete_results,omitempty"`
	Items             []components.TopicSearchResultItem `json:"items,omitempty"`
	TotalCount        int64                              `json:"total_count,omitempty"`
}

/*
SearchTopicsResponse is a response for SearchTopics

https://developer.github.com/v3/search/#search-topics
*/
type SearchTopicsResponse struct {
	response
	request *SearchTopicsReq
	Data    SearchTopicsResponseBody
}

/*
SearchUsers performs requests for "search/users"

Search users.

  GET /search/users

https://developer.github.com/v3/search/#search-users
*/
func SearchUsers(ctx context.Context, req *SearchUsersReq, opt ...RequestOption) (*SearchUsersResponse, error) {
	if req == nil {
		req = new(SearchUsersReq)
	}
	resp := &SearchUsersResponse{request: req}
	r, err := doRequest(ctx, req, "search/users", opt...)
	if r != nil {
		resp.response = *r
	}
	if err != nil {
		return resp, err
	}
	resp.Data = SearchUsersResponseBody{}
	err = r.decodeBody(&resp.Data, "search/users")
	if err != nil {
		return nil, err
	}
	return resp, nil
}

/*
SearchUsers performs requests for "search/users"

Search users.

  GET /search/users

https://developer.github.com/v3/search/#search-users
*/
func (c Client) SearchUsers(ctx context.Context, req *SearchUsersReq, opt ...RequestOption) (*SearchUsersResponse, error) {
	return SearchUsers(ctx, req, append(c, opt...)...)
}

/*
SearchUsersReq is request data for Client.SearchUsers

https://developer.github.com/v3/search/#search-users
*/
type SearchUsersReq struct {
	_url string

	/*
	The query contains one or more search keywords and qualifiers. Qualifiers allow
	you to limit your search to specific areas of GitHub. The REST API supports the
	same qualifiers as GitHub.com. To learn more about the format of the query, see
	[Constructing a search
	query](https://developer.github.com/v3/search/#constructing-a-search-query). See
	"[Searching users](https://help.github.com/articles/searching-users/)" for a
	detailed list of qualifiers.
	*/
	Q *string

	/*
	Sorts the results of your query by number of `followers` or `repositories`, or
	when the person `joined` GitHub. Default: [best
	match](https://developer.github.com/v3/search/#ranking-search-results)
	*/
	Sort *string

	/*
	Determines whether the first search result returned is the highest number of
	matches (`desc`) or lowest number of matches (`asc`). This parameter is ignored
	unless you provide `sort`.
	*/
	Order *string

	// Results per page (max 100)
	PerPage *int64

	// Page number of the results to fetch.
	Page *int64
}

func (r *SearchUsersReq) url() string {
	return r._url
}

func (r *SearchUsersReq) urlPath() string {
	return fmt.Sprintf("/search/users")
}

func (r *SearchUsersReq) method() string {
	return "GET"
}

func (r *SearchUsersReq) urlQuery() url.Values {
	query := url.Values{}
	if r.Q != nil {
		query.Set("q", *r.Q)
	}
	if r.Sort != nil {
		query.Set("sort", *r.Sort)
	}
	if r.Order != nil {
		query.Set("order", *r.Order)
	}
	if r.PerPage != nil {
		query.Set("per_page", strconv.FormatInt(*r.PerPage, 10))
	}
	if r.Page != nil {
		query.Set("page", strconv.FormatInt(*r.Page, 10))
	}
	return query
}

func (r *SearchUsersReq) header(requiredPreviews, allPreviews bool) http.Header {
	headerVals := map[string]*string{"accept": String("application/json")}
	previewVals := map[string]bool{}
	return requestHeaders(headerVals, previewVals)
}

func (r *SearchUsersReq) body() interface{} {
	return nil
}

func (r *SearchUsersReq) dataStatuses() []int {
	return []int{200}
}

func (r *SearchUsersReq) validStatuses() []int {
	return []int{200, 304}
}

// HTTPRequest builds an *http.Request
func (r *SearchUsersReq) HTTPRequest(ctx context.Context, opt ...RequestOption) (*http.Request, error) {
	return buildHTTPRequest(ctx, r, "search/users", opt)
}

/*
Rel updates this request to point to a relative link from resp. Returns false if
the link does not exist. Handy for paging.
*/
func (r *SearchUsersReq) Rel(link RelName, resp *SearchUsersResponse) bool {
	u := resp.RelLink(link)
	if u == "" {
		return false
	}
	r._url = u
	return true
}

/*
SearchUsersResponseBody is a response body for SearchUsers

https://developer.github.com/v3/search/#search-users
*/
type SearchUsersResponseBody struct {
	IncompleteResults bool                              `json:"incomplete_results,omitempty"`
	Items             []components.UserSearchResultItem `json:"items,omitempty"`
	TotalCount        int64                             `json:"total_count,omitempty"`
}

/*
SearchUsersResponse is a response for SearchUsers

https://developer.github.com/v3/search/#search-users
*/
type SearchUsersResponse struct {
	response
	request *SearchUsersReq
	Data    SearchUsersResponseBody
}
