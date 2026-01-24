# go-scm

[![Go Doc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://pkg.go.dev/github.com/open-beagle/go-scm/scm)
[![Go Report Card](https://goreportcard.com/badge/github.com/open-beagle/go-scm)](https://goreportcard.com/report/github.com/open-beagle/go-scm)
[![Version](https://img.shields.io/badge/version-v1.24.1--beagle-blue.svg)](https://github.com/open-beagle/go-scm/releases)

Package scm provides a unified interface to multiple source code management systems including GitHub, GitHub Enterprise, Bitbucket, Bitbucket Server, Gitee, Gitea, Gogs, GitLab, Azure DevOps and Stash.

## Features

- **Unified API**: Single interface for multiple SCM providers
- **Comprehensive**: Support for repositories, issues, pull requests, commits, webhooks and more
- **Flexible Authentication**: OAuth1, OAuth2, token-based authentication
- **Well-tested**: Production-ready code
- **Extensible**: Easy to add new providers

## Supported Providers

- GitHub (github.com)
- GitHub Enterprise
- GitLab (gitlab.com)
- GitLab Self-Hosted
- Bitbucket Cloud
- Bitbucket Server (Stash)
- Azure DevOps
- Gitea
- Gitee
- Gogs
- Beagle (Custom)

## Installation

```bash
go get github.com/open-beagle/go-scm
```

## Quick Start

### GitHub

```go
package main

import (
    "context"
    "github.com/open-beagle/go-scm/scm"
    "github.com/open-beagle/go-scm/scm/driver/github"
)

func main() {
    // Create a GitHub client
    client := github.NewDefault()

    // Get repository information
    repo, _, err := client.Repositories.Find(context.Background(), "octocat/Hello-World")
    if err != nil {
        panic(err)
    }
    println(repo.Name)
}
```

### GitHub Enterprise

```go
import (
    "github.com/open-beagle/go-scm/scm/driver/github"
)

func main() {
    client, err := github.New("https://github.company.com/api/v3")
    if err != nil {
        panic(err)
    }
}
```

### GitLab

```go
import (
    "github.com/open-beagle/go-scm/scm/driver/gitlab"
)

func main() {
    client := gitlab.NewDefault()
    // or for self-hosted
    // client, err := gitlab.New("https://gitlab.company.com")
}
```

### Bitbucket Cloud

```go
import (
    "github.com/open-beagle/go-scm/scm/driver/bitbucket"
)

func main() {
    client := bitbucket.NewDefault()
}
```

### Bitbucket Server (Stash)

```go
import (
    "github.com/open-beagle/go-scm/scm/driver/stash"
)

func main() {
    client, err := stash.New("https://stash.company.com")
}
```

### Gitea

```go
import (
    "github.com/open-beagle/go-scm/scm/driver/gitea"
)

func main() {
    client, err := gitea.New("https://gitea.company.com")
}
```

### Gitee

```go
import (
    "github.com/open-beagle/go-scm/scm/driver/gitee"
)

func main() {
    client, err := gitee.New("https://gitee.com")
}
```

### Azure DevOps

```go
import (
    "github.com/open-beagle/go-scm/scm/driver/azure"
)

func main() {
    client := azure.NewDefault()
}
```

## Authentication

The scm client does not directly handle authentication. Instead, provide an `http.Client` with a transport that handles authentication. This library includes OAuth1, OAuth2, and token-based authentication implementations.

### OAuth2 Token Authentication

```go
package main

import (
    "net/http"
    "github.com/open-beagle/go-scm/scm"
    "github.com/open-beagle/go-scm/scm/driver/github"
    "github.com/open-beagle/go-scm/scm/transport/oauth2"
)

func main() {
    client := github.NewDefault()

    // Inject OAuth2 token
    client.Client = &http.Client{
        Transport: &oauth2.Transport{
            Source: oauth2.StaticTokenSource(
                &scm.Token{
                    Token: "your-oauth2-token",
                },
            ),
        },
    }
}
```

### Private Token Authentication (GitLab)

```go
import (
    "net/http"
    "github.com/open-beagle/go-scm/scm/transport"
)

func main() {
    client := gitlab.NewDefault()

    // Inject private token via header
    client.Client = &http.Client{
        Transport: &transport.PrivateToken{
            Token: "your-private-token",
        },
    }
}
```

### Basic Authentication

```go
import (
    "net/http"
    "github.com/open-beagle/go-scm/scm/transport"
)

func main() {
    client := github.NewDefault()

    client.Client = &http.Client{
        Transport: &transport.BasicAuth{
            Username: "your-username",
            Password: "your-password",
        },
    }
}
```

## Usage Examples

### Get Repository

```go
repo, response, err := client.Repositories.Find(ctx, "owner/repo")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Repository: %s\n", repo.FullName)
```

### List Repositories

```go
opts := scm.ListOptions{
    Page: 1,
    Size: 30,
}

repos, response, err := client.Repositories.List(ctx, opts)
```

### Get Issue

```go
issue, response, err := client.Issues.Find(ctx, "owner/repo", 1)
```

### List Issues

```go
opts := scm.IssueListOptions{
    Page:   1,
    Size:   30,
    Open:   true,
    Closed: false,
}

issues, response, err := client.Issues.List(ctx, "owner/repo", opts)
```

### Create Issue Comment

```go
input := &scm.CommentInput{
    Body: "This is a comment",
}

comment, response, err := client.Issues.CreateComment(ctx, "owner/repo", 1, input)
```

### Get Pull Request

```go
pr, response, err := client.PullRequests.Find(ctx, "owner/repo", 1)
```

### List Pull Requests

```go
opts := scm.PullRequestListOptions{
    Page:   1,
    Size:   30,
    Open:   true,
    Closed: false,
}

prs, response, err := client.PullRequests.List(ctx, "owner/repo", opts)
```

### Get File Contents

```go
content, response, err := client.Contents.Find(ctx, "owner/repo", "path/to/file.txt", "main")
```

### Create/Update File

```go
params := &scm.ContentParams{
    Message: "Update file",
    Data:    []byte("file content"),
    Branch:  "main",
}

response, err := client.Contents.Create(ctx, "owner/repo", "path/to/file.txt", params)
```

### List Commits

```go
opts := scm.CommitListOptions{
    Page: 1,
    Size: 30,
}

commits, response, err := client.Git.ListCommits(ctx, "owner/repo", opts)
```

### Create Webhook

```go
input := &scm.HookInput{
    Name:   "web",
    Target: "https://example.com/webhook",
    Secret: "webhook-secret",
    Events: scm.HookEvents{
        Push:        true,
        PullRequest: true,
        Issue:       true,
    },
}

hook, response, err := client.Repositories.CreateHook(ctx, "owner/repo", input)
```

### Parse Webhook

```go
func handleWebhook(w http.ResponseWriter, r *http.Request) {
    webhook, err := client.Webhooks.Parse(r, func(webhook scm.Webhook) (string, error) {
        return "webhook-secret", nil
    })

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    switch webhook.(type) {
    case *scm.PushHook:
        // Handle push event
    case *scm.PullRequestHook:
        // Handle pull request event
    case *scm.IssueHook:
        // Handle issue event
    }
}
```

## API Coverage

The library provides access to:

- **Repositories**: Create, read, update, delete repositories
- **Contents**: Read, create, update, delete files
- **Git**: Commits, branches, tags, trees, references
- **Issues**: Create, read, update, close issues and comments
- **Pull Requests**: Create, read, update, merge pull requests and comments
- **Reviews**: Create, read, update pull request reviews
- **Organizations**: List and manage organizations
- **Users**: Get user information
- **Webhooks**: Create, list, delete webhooks and parse webhook payloads
- **Releases**: Create, list, update releases
- **Milestones**: Create, list, update milestones

## Provider API Documentation

- [Azure DevOps](https://docs.microsoft.com/en-us/rest/api/azure/devops/git/)
- [Bitbucket Cloud](https://developer.atlassian.com/cloud/bitbucket/rest/intro/)
- [Bitbucket Server/Stash](https://docs.atlassian.com/bitbucket-server/rest/)
- [Gitea](https://docs.gitea.io/en-us/api-usage/)
- [Gitee](https://gitee.com/api/v5/swagger)
- [GitHub](https://docs.github.com/en/rest)
- [GitLab](https://docs.gitlab.com/ee/api/)
- [Gogs](https://github.com/gogs/docs-api)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the BSD-style license. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

This project is a fork of the original [drone/go-scm](https://github.com/drone/go-scm) project, maintained and enhanced by the Open Beagle team.
