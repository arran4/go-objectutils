# Agent guidance

## Conversion and error API
- Preserve the README's documented Get*, MustGet*, Get*OrDefault and Get*Ptr behavioural distinctions; do not turn a failure-returning operation into a panic or a silent fallback.
- Preserve caller-relevant error types, including MissingFieldError and InvalidTypeError, and keep field/index context when errors travel through nested arrays/maps. Prefer errors.Is/errors.As tests to fragile exact-string matching where identity is the contract.
- Exercise numeric bounds, overflows, dates/time zones, nil/missing versus present-zero and nested structures using readable table tests and fuzzing when it adds value.
- Maintain public exported API compatibility and do not add unrelated wrappers or giant interfaces just for test mocks.

## Documentation and follow-through
Issue #16 already covers array index errors, #18 covers numeric/date/fuzz boundaries and #19 covers docs/API synchronization; use those issues rather than inventing duplicates.
A coding agent should compare implementation against documented contracts and report only new file/line-backed findings in issue #24. Update README/API docs when public behaviour changes; run relevant Go tests and examples.
## Engineering reference and follow-up

- Adopt only relevant guidance from https://arran4.github.io/blog/post/2026/055-suggested-coding-standards-for-repositories-and-agents/ and https://arran4.github.io/blog/post/2026/020-optional-dependency-injection-via-type-switched-variadic-args/ .
- This documentation is not evidence of repository-wide compliance. A coding-agent pass must report checked files, observed gaps, tests/CI and existing issue overlap to https://github.com/arran4/go-objectutils/issues/24 before any broad implementation.
- Do not merge implementation automatically; request review and follow repository-specific instructions.
