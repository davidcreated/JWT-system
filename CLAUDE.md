# Global Agent Rules (Master — applies to ALL projects)

These rules apply to **every** repository, existing or new, and to **any** coding
agent (Claude Code especially, Codex, and others). Project-level `CLAUDE.md` /
`AGENTS.md` files add to — never override — these. Where a capability named below
doesn't exist for a given agent (e.g. "plan mode", "subagents"), apply its intent
(write the plan first and confirm before executing; parallelize research where
you can).

---

## Workflow Orchestration

### 1. Plan Mode Default
- Enter plan mode for ANY non-trivial task (3+ steps or architectural decisions).
- If something goes sideways, STOP and re-plan immediately.
- Use plan mode for verification steps, not just building.
- Write detailed specs upfront to reduce ambiguity.

### 2. Subagent Strategy
- Use subagents liberally to keep the main context window clean.
- Offload research, exploration, and parallel analysis to subagents.
- For complex problems, throw more compute at it via subagents.
- One task per subagent for focused execution.

### 3. Self-Improvement Loop
- After ANY correction from the user: update `tasks/lessons.md` with the pattern.
- Write rules for yourself that prevent the same mistake.
- Ruthlessly iterate on these lessons until the mistake rate drops.
- Review lessons at session start for the relevant project.

### 4. Verification Before Done
- Never mark a task complete without proving it works.
- Diff behavior between main and your changes when relevant.
- Ask yourself: "Would a staff engineer approve this?"
- Run tests, check logs, demonstrate correctness.

### 5. Demand Elegance (Balanced)
- For non-trivial changes: pause and ask "is there a more elegant way?"
- If a fix feels hacky: "Knowing everything I know now, implement the elegant solution."
- Skip this for simple, obvious fixes — don't over-engineer.
- Challenge your own work before presenting it.

### 6. Autonomous Bug Fixing
- When given a bug report: just fix it. Don't ask for hand-holding.
- Point at logs, errors, failing tests — then resolve them.
- Zero context switching required from the user.
- Go fix failing CI tests without being told how.

---

## Task Management
1. **Plan First:** Write the plan to `tasks/todo.md` with checkable items.
2. **Verify Plan:** Check in before starting implementation.
3. **Track Progress:** Mark items complete as you go.
4. **Explain Changes:** High-level summary at each step.
5. **Document Results:** Add a review section to `tasks/todo.md`.
6. **Capture Lessons:** Update `tasks/lessons.md` after corrections.

---

## Core Principles
- **Simplicity First:** Make every change as simple as possible. Impact minimal code.
- **No Laziness:** Find root causes. No temporary fixes. Senior-developer standards.
- **Minimal Impact:** Only touch what's necessary. No side effects, no new bugs.

---

## Token & Context Economy (avoid waste at all costs)
Tokens and context window are scarce, paid resources. Treat every read, tool call,
and word of output as costing money — spend the minimum that still does the job right.

**Reading & searching**
- **Locate, then read narrowly.** Use search (grep/glob) to find the exact spot,
  then read only the needed line range. Never read a whole large file to find one
  thing.
- **Never re-read a file you just wrote or edited** — the change already succeeded.
- **Reuse what's already in context.** Don't re-fetch, re-derive, or re-read facts
  established earlier in the session.

**Tool calls**
- **Filter output at the source.** Pipe long command output through `grep`/`head`/
  `tail`/`jq` and print only the relevant lines. Never dump whole logs, JSON blobs,
  or file contents into context.
- **Batch & parallelize.** Run independent reads/searches together in one step;
  don't issue serial round-trips or duplicate calls.
- **Don't re-run expensive commands.** Capture a result once and reuse it.
- **Offload heavy exploration to subagents** so their large intermediate context
  stays out of the main window; take back only the conclusion.

**Output & editing**
- **Be concise.** Answer directly. No filler, no restating the task, no re-explaining
  what the user already knows. Reference files by `path:line` instead of pasting them.
- **Surgical edits over rewrites.** Change the few lines that matter; don't rewrite
  whole files.
- **Targeted debug output only** (e.g. keys + first item), never verbose dumps.

**Stop the bleed**
- **No rabbit holes.** If an approach isn't converging, STOP and re-plan — thrashing
  burns tokens faster than thinking. (See Plan Mode + Self-Improvement Loop.)
- **Match effort to the task.** Trivial change → do it directly; don't spin up
  research, subagents, or long analysis for a one-liner.
