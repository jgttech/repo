# Save (AI-Assisted Commit)

> Stage all unstaged files and generate a conventional commit message.

---

## Execution

<execution>

### Step 1: Stage All Changes

```bash
git add .
```

Stage all untracked and modified files before committing.

### Step 2: Verify Staged Changes

```bash
git diff --staged --quiet
```

**If no staged changes exist → FAIL immediately.** Do NOT proceed.

### Step 3: Gather Context

1. `git diff --staged` — review what's being committed
2. `git log --oneline -5` — match repository commit style
3. Glob `**/WIP*.md` — if found and relevant, read for ongoing work context

### Step 4: Generate Commit Message

**Format:** `<type>[(scope)][!]: <description>` + optional body/footers

**Types:** `feat` | `fix` | `docs` | `style` | `refactor` | `perf` | `test` | `build` | `ci` | `chore`

**Scope** (optional): area affected — e.g., `feat(auth):`, `fix(api):`

**Breaking Changes:** Add `!` after type/scope — e.g., `feat(api)!: change response format`

**Description:** imperative mood, lowercase, no period, max 72 chars. Completes "This commit will..."

**Body** (optional): wrap at 72 cols, explain _what_ and _why_ (not _how_), 2-4 sentences max, blank line after description.

**Footers** (optional):
- `BREAKING CHANGE: <explanation>`
- `Refs: #123`
- `WIP: <filename>: <status>` — only if related to tracked WIP work

### Step 5: Create Commit

Single-line: `git commit -m "<message>"`

With body/footers:

```bash
cat << 'EOF' > /tmp/commit-msg.txt
<type>[(scope)]: <description>

<body>

<footers>
EOF
git commit -F /tmp/commit-msg.txt
rm /tmp/commit-msg.txt
```

### Step 6: Verify

```bash
git log -1 --oneline
```

Confirm the commit hash and message appear. If the commit failed, exit with error.

</execution>

---

## Rules

<rules>

- Match the repository's existing commit style
- Be specific about what changed; explain why if not obvious
- Reference related issues/tickets if known
- Do NOT add AI signatures, co-author lines, or emoji
- Do NOT list every file changed
- Do NOT pad the body or attempt to recover from failures

</rules>

---

## Examples

```
fix(parser): handle ISO-8601 timestamps with timezone offset

The regex rejected valid timestamps with +00:00 format.
```

```
feat(auth)!: require API key for all endpoints

All endpoints now require X-API-Key header. Previously only
write operations required authentication.

BREAKING CHANGE: Unauthenticated read requests will return 401.
```

```
refactor(settings): extract preference validation

Moves validation logic to shared module for reuse across
settings and profile endpoints.

WIP: WIP-settings.md: validation complete, UI pending
```

---
