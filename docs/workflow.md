# Contribution Workflow

The Fabric Smart Client project uses a contribution workflow heavily inspired by the workflow developed by the [Hiero community](https://github.com/hiero-ledger/hiero-sdk-python), presented at LFDT Maintainer Days ([recording](https://www.youtube.com/watch?v=I87WCpiXOOs)).

The process is designed to:
- Reduce duplicate work and assignment conflicts.
- Give external contributors a clear path to claim and deliver work.
- Keep issues and PRs moving without requiring constant manual intervention from maintainers.

---

## Issue Lifecycle

### Issue Types

Every issue should have a type:

| Type      | Meaning                           |
|-----------|-----------------------------------|
| `Bug`     | Something is broken               |
| `Feature` | New feature or improvement        |
| `Task`    | A specific, bounded piece of work |

And one of the following status labels:

| Label                     | Meaning                                                            |
|---------------------------|--------------------------------------------------------------------|
| `status: awaiting triage` | New issue that needs to be reviewed and categorized by maintainers |
| `status: ready for dev`   | Fully defined and ready for a contributor to pick up               |
| `status: in progress`     | A contributor is actively working on this issue                    |

An issue without a `status: ready for dev` label is **not ready for contribution**.

### Creating an Issue

Anyone may open an issue using the `bug`, `feature`, or `task` template. New issues start untriaged.

Maintainers and core contributors review new issues and apply `status: ready for dev` once the issue is well-defined, scoped, and accepted.

#### Parent and Child Issues

For larger efforts, a **parent issue** may be opened to capture the overall goal, with individual **child issues** that are well scoped and actionable.
Each child issue is the unit of assignment, `status: ready for dev` labeling, and PR linking.
Note that dependencies between children can be expressed using the `Marked as blocked by` relationship.

### Claiming an Issue

Contributors must claim an issue **before** opening a PR:

1. Comment `/assign` on the issue.
2. The bot checks:
    - The issue carries `status: ready for dev`.
    - The contributor has **no more than two open assigned issues** (limit across all issues in the repository).
3. If both conditions are met, the bot assigns the contributor and confirms in a comment.
4. To release an issue voluntarily, comment `/unassign`.

Maintainers may assign any contributor directly — this bypasses all bot eligibility checks.

### Issue Inactivity

Once assigned, the bot monitors activity (comments, linked PR events):

| Threshold                                | Action                                                                                               |
|------------------------------------------|------------------------------------------------------------------------------------------------------|
| 5 days of no activity                    | Bot posts a reminder tagging the assignee                                                            |
| 7 days of no activity                    | Bot unassigns the contributor with an explanatory comment; issue becomes available for re-assignment |
| 9 days assigned with no linked PR opened | Bot posts a reminder asking the contributor to open a PR or comment `/working`                       |

A contributor can reset the inactivity timer at any time by:
- Commenting `/working` — signals continued active progress.
- Commenting `/blocked` — signals the contributor is waiting on feedback from maintainers. The inactivity timer is paused until a maintainer responds or the contributor comments `/working` again.

The 7-day unassignment window is intentionally short to keep the queue moving. For issues of higher complexity, a maintainer may manually extend the window or re-assign as appropriate.

## Pull Request (PR) Lifecycle

### Linking a PR to an Issue

Every PR opened by a contributor must:
1. Reference an open issue carrying `ready` or `good-first-issue`, using a closing keyword in the PR description (e.g. `Fixes #123`).
2. Have the PR author **assigned to the linked issue**.

PRs by maintainers, core contributors, and `dependabot` are exempt from both requirements.

If either condition is not met for a contributor's PR, the bot posts a warning comment. If the PR remains non-compliant for **12 hours**, it is automatically closed with an explanation and a pointer to the contribution process docs.

### PR Labels

| Label                    | Meaning                                                                               |
|--------------------------|---------------------------------------------------------------------------------------|
| `status: needs review`   | The pull request is ready for maintainer review                                       |
| `status: needs revision` | The pull request requires changes from the author before it can be reviewed or merged |

### PR Inactivity

| Threshold                                                   | Action                                        |
|-------------------------------------------------------------|-----------------------------------------------|
| 5 days of no activity (commits, review responses, comments) | Bot posts a reminder tagging the author       |
| 7 days of no activity                                       | Bot closes the PR with an explanatory comment |

When a PR is auto-closed, the linked issue is **not** automatically unassigned. The contributor may open a new PR against the same issue within their original assignment window.

### PR Checks

Standard automated checks run on every PR:

- DCO sign-off
- Unit tests
- Integration tests
- Linter / static analysis

The full set of checks is defined in the repository's CI configuration.

## Roles

| Role                 | Definition                                          |
|----------------------|-----------------------------------------------------|
| **Maintainer**       | Has `admin` or `write` permission on the repository |
| **Core contributor** | Has `triage` permission on the repository           |
| **Contributor**      | Everyone else                                       |

Maintainers and core contributors are exempt from all assignment and PR-linking rules described below. PRs from `dependabot` are also exempt. Maintainers may directly assign any contributor to any issue at any time, bypassing eligibility checks.

---

## References

- [Hiero SDK Python — GitHub Workflows](https://github.com/hiero-ledger/hiero-sdk-python/tree/main/.github/workflows)
- [Hiero SDK C++ — GitHub Workflows](https://github.com/hiero-ledger/hiero-sdk-cpp/tree/main/.github/workflows)
- [LFDT Maintainer Days — Contribution Workflow Presentation](https://www.youtube.com/watch?v=I87WCpiXOOs)
- [Original proposal issue — fabric-x#130](https://github.com/hyperledger/fabric-x/issues/130)
