{
  _config+:: {
    local c = self,
    dashboardNamePrefix: 'Gitea',
    dashboardTags: ['proxgit'],
    dashboardPeriod: 'now-1h',
    dashboardTimezone: 'default',
    dashboardRefresh: '1m',

    // please see https://docs.proxgit.com/administration/config-cheat-sheet#metrics-metrics
    // Show issue by repository metrics with format proxgit_issues_by_repository{repository="org/repo"} 5.
    // Requires Gitea 1.16.0 with ENABLED_ISSUE_BY_REPOSITORY set to true.
    showIssuesByRepository: true,
    // Show graphs for issue by label metrics with format proxgit_issues_by_label{label="bug"} 2.
    // Requires Gitea 1.16.0 with ENABLED_ISSUE_BY_LABEL set to true.
    showIssuesByLabel: true,

    // Requires Gitea 1.16.0.
    showIssuesOpenClose: true,

    // add or remove metrics from dashboard
    proxgitStatMetrics:
      [
        {
          name: 'proxgit_organizations',
          description: 'Organizations',
        },
        {
          name: 'proxgit_teams',
          description: 'Teams',
        },
        {
          name: 'proxgit_users',
          description: 'Users',
        },
        {
          name: 'proxgit_repositories',
          description: 'Repositories',
        },
        {
          name: 'proxgit_milestones',
          description: 'Milestones',
        },
        {
          name: 'proxgit_stars',
          description: 'Stars',
        },
        {
          name: 'proxgit_releases',
          description: 'Releases',
        },
      ]
      +
      if c.showIssuesOpenClose then
        [
          {
            name: 'proxgit_issues_open',
            description: 'Issues opened',
          },
          {
            name: 'proxgit_issues_closed',
            description: 'Issues closed',
          },
        ] else
        [
          {
            name: 'proxgit_issues',
            description: 'Issues',
          },
        ],
    //set this for using label colors on graphs
    issueLabels: [
      {
        label: 'bug',
        color: '#ee0701',
      },
      {
        label: 'duplicate',
        color: '#cccccc',
      },
      {
        label: 'invalid',
        color: '#e6e6e6',
      },
      {
        label: 'enhancement',
        color: '#84b6eb',
      },
      {
        label: 'help wanted',
        color: '#128a0c',
      },
      {
        label: 'question',
        color: '#cc317c',
      },
    ],
  },
}
