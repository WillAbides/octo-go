// Code generated by octo-go; DO NOT EDIT.

package internal

const (
	AttrRedirectOnly EndpointAttribute = iota
	AttrBoolean
	AttrBodyUploader
	AttrJSONRequestBody
	AttrExplicitURL
	AttrNoResponseBody
)

func init() {
	operationAttributes = operationAttributesGen
}

var operationAttributesGen = map[string][]EndpointAttribute{

	"actions/create-or-update-org-secret":                      {AttrJSONRequestBody},
	"actions/create-or-update-repo-secret":                     {AttrJSONRequestBody},
	"actions/create-workflow-dispatch":                         {AttrJSONRequestBody},
	"actions/download-artifact":                                {AttrRedirectOnly},
	"actions/download-job-logs-for-workflow-run":               {AttrRedirectOnly},
	"actions/download-workflow-run-logs":                       {AttrRedirectOnly},
	"actions/set-selected-repos-for-org-secret":                {AttrJSONRequestBody},
	"activity/mark-notifications-as-read":                      {AttrJSONRequestBody},
	"activity/mark-repo-notifications-as-read":                 {AttrJSONRequestBody},
	"activity/set-repo-subscription":                           {AttrJSONRequestBody},
	"activity/set-thread-subscription":                         {AttrJSONRequestBody},
	"apps/check-token":                                         {AttrJSONRequestBody},
	"apps/create-content-attachment":                           {AttrJSONRequestBody},
	"apps/create-installation-access-token":                    {AttrJSONRequestBody},
	"apps/delete-authorization":                                {AttrJSONRequestBody},
	"apps/delete-installation":                                 {AttrBoolean},
	"apps/delete-token":                                        {AttrJSONRequestBody},
	"apps/reset-token":                                         {AttrJSONRequestBody},
	"apps/suspend-installation":                                {AttrBoolean},
	"apps/unsuspend-installation":                              {AttrBoolean},
	"checks/create":                                            {AttrJSONRequestBody},
	"checks/create-suite":                                      {AttrJSONRequestBody},
	"checks/set-suites-preferences":                            {AttrJSONRequestBody},
	"checks/update":                                            {AttrJSONRequestBody},
	"gists/create":                                             {AttrJSONRequestBody},
	"gists/create-comment":                                     {AttrJSONRequestBody},
	"gists/update":                                             {AttrJSONRequestBody},
	"gists/update-comment":                                     {AttrJSONRequestBody},
	"git/create-blob":                                          {AttrJSONRequestBody},
	"git/create-commit":                                        {AttrJSONRequestBody},
	"git/create-ref":                                           {AttrJSONRequestBody},
	"git/create-tag":                                           {AttrJSONRequestBody},
	"git/create-tree":                                          {AttrJSONRequestBody},
	"git/update-ref":                                           {AttrJSONRequestBody},
	"interactions/set-restrictions-for-org":                    {AttrJSONRequestBody},
	"interactions/set-restrictions-for-repo":                   {AttrJSONRequestBody},
	"issues/add-assignees":                                     {AttrJSONRequestBody},
	"issues/add-labels":                                        {AttrJSONRequestBody},
	"issues/check-user-can-be-assigned":                        {AttrBoolean},
	"issues/create":                                            {AttrJSONRequestBody},
	"issues/create-comment":                                    {AttrJSONRequestBody},
	"issues/create-label":                                      {AttrJSONRequestBody},
	"issues/create-milestone":                                  {AttrJSONRequestBody},
	"issues/delete-milestone":                                  {AttrBoolean},
	"issues/lock":                                              {AttrJSONRequestBody},
	"issues/remove-assignees":                                  {AttrJSONRequestBody},
	"issues/set-labels":                                        {AttrJSONRequestBody},
	"issues/update":                                            {AttrJSONRequestBody},
	"issues/update-comment":                                    {AttrJSONRequestBody},
	"issues/update-label":                                      {AttrJSONRequestBody},
	"issues/update-milestone":                                  {AttrJSONRequestBody},
	"markdown/render":                                          {AttrJSONRequestBody},
	"markdown/render-raw":                                      {AttrBodyUploader},
	"migrations/delete-archive-for-org":                        {AttrBoolean},
	"migrations/map-commit-author":                             {AttrJSONRequestBody},
	"migrations/set-lfs-preference":                            {AttrJSONRequestBody},
	"migrations/start-for-authenticated-user":                  {AttrJSONRequestBody},
	"migrations/start-for-org":                                 {AttrJSONRequestBody},
	"migrations/start-import":                                  {AttrJSONRequestBody},
	"migrations/unlock-repo-for-org":                           {AttrBoolean},
	"migrations/update-import":                                 {AttrJSONRequestBody},
	"oauth-authorizations/create-authorization":                {AttrJSONRequestBody},
	"oauth-authorizations/get-or-create-authorization-for-app": {AttrJSONRequestBody},
	"oauth-authorizations/get-or-create-authorization-for-app-and-fingerprint": {AttrJSONRequestBody},
	"oauth-authorizations/update-authorization":                                {AttrJSONRequestBody},
	"orgs/check-blocked-user":                                                  {AttrBoolean},
	"orgs/check-public-membership-for-user":                                    {AttrBoolean},
	"orgs/create-invitation":                                                   {AttrJSONRequestBody},
	"orgs/create-webhook":                                                      {AttrJSONRequestBody},
	"orgs/delete-webhook":                                                      {AttrBoolean},
	"orgs/ping-webhook":                                                        {AttrBoolean},
	"orgs/remove-saml-sso-authorization":                                       {AttrBoolean},
	"orgs/set-membership-for-user":                                             {AttrJSONRequestBody},
	"orgs/update":                                                              {AttrJSONRequestBody},
	"orgs/update-membership-for-authenticated-user":                            {AttrJSONRequestBody},
	"orgs/update-webhook":                                                      {AttrJSONRequestBody},
	"projects/add-collaborator":                                                {AttrJSONRequestBody},
	"projects/create-card":                                                     {AttrJSONRequestBody},
	"projects/create-column":                                                   {AttrJSONRequestBody},
	"projects/create-for-authenticated-user":                                   {AttrJSONRequestBody},
	"projects/create-for-org":                                                  {AttrJSONRequestBody},
	"projects/create-for-repo":                                                 {AttrJSONRequestBody},
	"projects/move-card":                                                       {AttrJSONRequestBody, AttrNoResponseBody},
	"projects/move-column":                                                     {AttrJSONRequestBody, AttrNoResponseBody},
	"projects/update":                                                          {AttrJSONRequestBody},
	"projects/update-card":                                                     {AttrJSONRequestBody},
	"projects/update-column":                                                   {AttrJSONRequestBody},
	"pulls/check-if-merged":                                                    {AttrBoolean},
	"pulls/create":                                                             {AttrJSONRequestBody},
	"pulls/create-reply-for-review-comment":                                    {AttrJSONRequestBody},
	"pulls/create-review":                                                      {AttrJSONRequestBody},
	"pulls/create-review-comment":                                              {AttrJSONRequestBody},
	"pulls/delete-review-comment":                                              {AttrBoolean},
	"pulls/dismiss-review":                                                     {AttrJSONRequestBody},
	"pulls/merge":                                                              {AttrJSONRequestBody},
	"pulls/remove-requested-reviewers":                                         {AttrJSONRequestBody},
	"pulls/request-reviewers":                                                  {AttrJSONRequestBody},
	"pulls/submit-review":                                                      {AttrJSONRequestBody},
	"pulls/update":                                                             {AttrJSONRequestBody},
	"pulls/update-branch":                                                      {AttrJSONRequestBody},
	"pulls/update-review":                                                      {AttrJSONRequestBody},
	"pulls/update-review-comment":                                              {AttrJSONRequestBody},
	"reactions/create-for-commit-comment":                                      {AttrJSONRequestBody},
	"reactions/create-for-issue":                                               {AttrJSONRequestBody},
	"reactions/create-for-issue-comment":                                       {AttrJSONRequestBody},
	"reactions/create-for-pull-request-review-comment":                         {AttrJSONRequestBody},
	"reactions/create-for-team-discussion-comment-in-org":                      {AttrJSONRequestBody},
	"reactions/create-for-team-discussion-comment-legacy":                      {AttrJSONRequestBody},
	"reactions/create-for-team-discussion-in-org":                              {AttrJSONRequestBody},
	"reactions/create-for-team-discussion-legacy":                              {AttrJSONRequestBody},
	"repos/add-app-access-restrictions":                                        {AttrJSONRequestBody},
	"repos/add-collaborator":                                                   {AttrJSONRequestBody},
	"repos/add-status-check-contexts":                                          {AttrJSONRequestBody},
	"repos/add-team-access-restrictions":                                       {AttrJSONRequestBody},
	"repos/add-user-access-restrictions":                                       {AttrJSONRequestBody},
	"repos/check-collaborator":                                                 {AttrBoolean},
	"repos/check-vulnerability-alerts":                                         {AttrBoolean},
	"repos/create-commit-comment":                                              {AttrJSONRequestBody},
	"repos/create-commit-status":                                               {AttrJSONRequestBody},
	"repos/create-deploy-key":                                                  {AttrJSONRequestBody},
	"repos/create-deployment":                                                  {AttrJSONRequestBody},
	"repos/create-deployment-status":                                           {AttrJSONRequestBody},
	"repos/create-dispatch-event":                                              {AttrJSONRequestBody},
	"repos/create-for-authenticated-user":                                      {AttrJSONRequestBody},
	"repos/create-fork":                                                        {AttrJSONRequestBody},
	"repos/create-in-org":                                                      {AttrJSONRequestBody},
	"repos/create-or-update-file-contents":                                     {AttrJSONRequestBody},
	"repos/create-pages-site":                                                  {AttrJSONRequestBody},
	"repos/create-release":                                                     {AttrJSONRequestBody},
	"repos/create-using-template":                                              {AttrJSONRequestBody},
	"repos/create-webhook":                                                     {AttrJSONRequestBody},
	"repos/delete-admin-branch-protection":                                     {AttrBoolean},
	"repos/delete-commit-comment":                                              {AttrBoolean},
	"repos/delete-commit-signature-protection":                                 {AttrBoolean},
	"repos/delete-file":                                                        {AttrJSONRequestBody},
	"repos/delete-pull-request-review-protection":                              {AttrBoolean},
	"repos/delete-webhook":                                                     {AttrBoolean},
	"repos/download-tarball-archive":                                           {AttrRedirectOnly},
	"repos/download-zipball-archive":                                           {AttrRedirectOnly},
	"repos/merge":                                                              {AttrJSONRequestBody},
	"repos/ping-webhook":                                                       {AttrBoolean},
	"repos/remove-app-access-restrictions":                                     {AttrJSONRequestBody},
	"repos/remove-status-check-contexts":                                       {AttrJSONRequestBody},
	"repos/remove-team-access-restrictions":                                    {AttrJSONRequestBody},
	"repos/remove-user-access-restrictions":                                    {AttrJSONRequestBody},
	"repos/replace-all-topics":                                                 {AttrJSONRequestBody},
	"repos/set-app-access-restrictions":                                        {AttrJSONRequestBody},
	"repos/set-status-check-contexts":                                          {AttrJSONRequestBody},
	"repos/set-team-access-restrictions":                                       {AttrJSONRequestBody},
	"repos/set-user-access-restrictions":                                       {AttrJSONRequestBody},
	"repos/test-push-webhook":                                                  {AttrBoolean},
	"repos/transfer":                                                           {AttrJSONRequestBody},
	"repos/update":                                                             {AttrJSONRequestBody},
	"repos/update-branch-protection":                                           {AttrJSONRequestBody},
	"repos/update-commit-comment":                                              {AttrJSONRequestBody},
	"repos/update-information-about-pages-site":                                {AttrJSONRequestBody},
	"repos/update-invitation":                                                  {AttrJSONRequestBody},
	"repos/update-pull-request-review-protection":                              {AttrJSONRequestBody},
	"repos/update-release":                                                     {AttrJSONRequestBody},
	"repos/update-release-asset":                                               {AttrJSONRequestBody},
	"repos/update-status-check-protection":                                     {AttrJSONRequestBody},
	"repos/update-webhook":                                                     {AttrJSONRequestBody},
	"repos/upload-release-asset":                                               {AttrBodyUploader, AttrExplicitURL},
	"scim/provision-and-invite-user":                                           {AttrJSONRequestBody},
	"scim/set-information-for-provisioned-user":                                {AttrJSONRequestBody},
	"scim/update-attribute-for-user":                                           {AttrJSONRequestBody},
	"teams/add-or-update-membership-for-user-in-org":                           {AttrJSONRequestBody},
	"teams/add-or-update-membership-for-user-legacy":                           {AttrJSONRequestBody},
	"teams/add-or-update-project-permissions-in-org":                           {AttrJSONRequestBody},
	"teams/add-or-update-project-permissions-legacy":                           {AttrJSONRequestBody},
	"teams/add-or-update-repo-permissions-in-org":                              {AttrJSONRequestBody},
	"teams/add-or-update-repo-permissions-legacy":                              {AttrJSONRequestBody},
	"teams/create":                                                             {AttrJSONRequestBody},
	"teams/create-discussion-comment-in-org":                                   {AttrJSONRequestBody},
	"teams/create-discussion-comment-legacy":                                   {AttrJSONRequestBody},
	"teams/create-discussion-in-org":                                           {AttrJSONRequestBody},
	"teams/create-discussion-legacy":                                           {AttrJSONRequestBody},
	"teams/create-or-update-idp-group-connections-in-org":                      {AttrJSONRequestBody},
	"teams/create-or-update-idp-group-connections-legacy":                      {AttrJSONRequestBody},
	"teams/get-member-legacy":                                                  {AttrBoolean},
	"teams/remove-member-legacy":                                               {AttrBoolean},
	"teams/update-discussion-comment-in-org":                                   {AttrJSONRequestBody},
	"teams/update-discussion-comment-legacy":                                   {AttrJSONRequestBody},
	"teams/update-discussion-in-org":                                           {AttrJSONRequestBody},
	"teams/update-discussion-legacy":                                           {AttrJSONRequestBody},
	"teams/update-in-org":                                                      {AttrJSONRequestBody},
	"teams/update-legacy":                                                      {AttrJSONRequestBody},
	"users/add-email-for-authenticated":                                        {AttrJSONRequestBody},
	"users/check-following-for-user":                                           {AttrBoolean},
	"users/create-gpg-key-for-authenticated":                                   {AttrJSONRequestBody},
	"users/create-public-ssh-key-for-authenticated":                            {AttrJSONRequestBody},
	"users/delete-email-for-authenticated":                                     {AttrJSONRequestBody},
	"users/set-primary-email-visibility-for-authenticated":                     {AttrJSONRequestBody},
	"users/update-authenticated":                                               {AttrJSONRequestBody},
}
