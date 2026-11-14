# Fork rules

This is Martin's public fork of [korotovsky/slack-mcp-server](https://github.com/korotovsky/slack-mcp-server). Everything committed here is public.

- Never commit personal data: Slack tokens and cookies (`xoxc`, `xoxd`, `xoxp`, `xoxb`), workspace, channel, user or file IDs and names from a real workspace, message content, email addresses, or local absolute paths (`/Users/...`, home directories). Tests, fixtures and docs use synthetic values only (`T0000000000`, `C0000000000`, `F0000000001`, `example.com`).
- Keep fork changes minimal: add only what upstream lacks, and prefer upstream's implementation once it covers a feature.
- Keep the fork rebased on upstream's default branch as a few clean, linear commits. When syncing, drop commits that upstream now covers.
- Verify with `go build ./...` and `go test -skip Integration ./...` (integration tests need ngrok and an OpenAI key).
