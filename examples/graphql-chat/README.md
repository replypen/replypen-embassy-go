# GraphQL-shaped embedded-chat example

This is the smallest dependency-free version of a common Go backend shape. `POST /graphql` accepts
`query ReplypenChatToken { replypenChatToken { token project baseUrl } }`; it is intentionally not a
general GraphQL server. Replace that small adapter with your existing GraphQL framework, keeping the
resolver body and response fields.

The example reads the external user id from an HttpOnly demo session cookie, mints a 2-hour token,
loads the page-mode widget under an explicit CSP, refreshes the token at half-life, and lets the
widget's `auth-expired` recovery reload the host page for a fresh token. The action and health mounts
are always present; without action secrets they return `ACTION_PLANE_DISABLED`.

## Run

```sh
export REPLYPEN_CHAT_SECRET='received out of band from the project operator'
export REPLYPEN_CHAT_PROJECT='your-project-slug'
export REPLYPEN_CHAT_PRINCIPAL_KIND='your_configured_principal_kind'
export REPLYPEN_CHAT_BASE_URL='https://app.replypen.com' # optional; this is the default
export APP_ORIGIN='http://localhost:8080'
go run ./examples/graphql-chat
```

Open <http://localhost:8080>. `/login` creates a fixed demo session; real applications use their
existing authenticated session and must never accept `external_id`, principal kind, tenant, or origin
from browser input.

To enable actions later, add `REPLYPEN_ACTION_SECRET` and `REPLYPEN_FETCH_URL`. Keep the chat and
action secrets different.

## Verify ladder

1. `curl -i http://localhost:8080/healthz` shows `actions:false` for chat-only boot.
2. Sign in once, then POST the query with the returned cookie; inspect only the JWT header/claims,
   never paste the token or signing secret into logs.
3. Open the page and confirm the widget renders. DevTools must show no CSP error.
4. Reuse the same freshly minted token for a second session open: expect `TOKEN_REPLAYED`.
5. Change `APP_ORIGIN` without changing the browser origin: expect `ORIGIN_MISMATCH`.
6. Run the complete public ladder in the
   [integrator verification guide](https://github.com/replypen/replypen-embassy/blob/main/docs/integrator/verify.md),
   including message SSE and a foreign-origin refusal.
7. When actions are configured, probe `/replypen/action`, signed `/replypen/action/health`, then a
   signed `dry_run:true` invocation before enabling execution.

For a stable error code, its self-fix, and the escalation fields, use the
[error catalogue](https://github.com/replypen/replypen-embassy/blob/main/docs/integrator/errors.md).
