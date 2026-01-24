# OpenTelemetry Transport

这个包提供了 OpenTelemetry 追踪支持，用于追踪所有 SCM API 调用。

## 使用方法

### 基本用法

```go
import (
    "github.com/drone/go-scm/scm/driver/gitlab"
    "github.com/drone/go-scm/scm/transport/otel"
    "net/http"
)

// 创建 GitLab 客户端
client, _ := gitlab.New("https://gitlab.com")

// 包装 HTTP Transport 以启用 OpenTelemetry 追踪
client.Client = &http.Client{
    Transport: otel.Transport(http.DefaultTransport),
}
```

### 与 OAuth2 结合使用

```go
import (
    "github.com/drone/go-scm/scm/driver/gitlab"
    "github.com/drone/go-scm/scm/transport/oauth2"
    "github.com/drone/go-scm/scm/transport/otel"
    "net/http"
)

client, _ := gitlab.New("https://gitlab.com")

// 先包装 OAuth2，再包装 OpenTelemetry
client.Client = &http.Client{
    Transport: &oauth2.Transport{
        Source: oauth2.ContextTokenSource(),
        Base:   otel.Transport(http.DefaultTransport),
    },
}
```

## 追踪效果

启用后，所有 SCM API 调用都会在 Jaeger 中显示为独立的 HTTP Span：

```
▼ GET /api/v1/repos/octocat/hello-world    [200ms]
  │
  ├─▶ HTTP GET                             [150ms]
  │   http.url: https://gitlab.com/api/v4/projects/123
  │   http.status_code: 200
  │
  └─▶ SELECT                               [30ms]
      db.system: mysql
```

## 依赖

```bash
go get go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp
```

## 注意事项

- 确保在应用启动时初始化了 OpenTelemetry TracerProvider
- Transport 包装顺序很重要：OAuth2 → OpenTelemetry → Base Transport
- 所有 API 调用都会被追踪，包括 GitLab、GitHub、Gitea、Bitbucket 等
