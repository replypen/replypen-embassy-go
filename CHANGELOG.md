# Changelog

## Unreleased

**0.8.0 — analysis credentials.** Conforms to hub `9aa0bfc` (decision 29).

- `AnalysisRequest.Credentials` sends the optional per-run `credentials` map: the chat claim's rules
  (now one validator, `internal/credentials`, shared with `chat`) plus "requires a `Principal`", else
  `CREDENTIALS_INVALID` before sending. Omitted when empty. Needs a host that knows the field (an
  older host answers `400`). Re-vendored fixtures incl. `analysis/trigger_with_credentials.json`.
- Credential values containing NUL are refused (chat token too: `CHAT_CREDENTIALS_INVALID`), matching the host.

**0.7.0 — renamed to ReplyPen.** Conforms to hub `450e9d8`.

- Module path is now `github.com/replypen/replypen-embassy-go`; the old path stays resolvable at its
  existing tags (v0.1.0–v0.6.0). Moving = change the `require` and every import.
- Chat tokens are minted with `aud` `replypen:chat:<project>` (`chat.Audience`); the host accepts the
  legacy `rootcause:chat:<project>` too. Tests asserting the old string need updating.
- Config reads `REPLYPEN_*` env first, then the legacy `ROOTCAUSE_*` names; refusal hints name
  `REPLYPEN_*`. Action scripts may import either module path.

Conforms to hub `b3a9dad`: widget tags use loader `?v=5` for per-message page Markdown and queued
snapshots. Re-vendored all fixtures, including unsigned page URL normalization cases.

## 0.6.0

Conforms to hub `2be70da` (decision 23): loader contract `?v=4` in the generated widget tag (hosted
loader gained the persistent Turbo mode; generated tags behave as before).

## 0.5.0

Conforms to hub `6d2c818` (decision 22: action-run id as chat-context locator).

- `ActionAPI.ActionRunID()` and virtual env `RC_ACTION_RUN_ID` expose the invocation's host-stamped
  `action_run_id`, invocation-scoped; empty/unset when absent, never inherited from process env.
  A malformed value refuses as signed `400 invalid_request` before resolution (dry run included).
  Breaking only for code that implements `ActionAPI` itself.
- `AnalysisRequest.ContextRefs` (`[]ContextRef{Kind, ID}`) sends `context_refs` after `session_id`.
  Client-side: at most one, `kind` `action_run`, canonical lowercase UUID id; otherwise
  `ANALYSIS_REQUEST_INVALID` before sending.
