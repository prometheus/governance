# Prometheus 2026 Bootstrap Steering Committee election

This document lays out the 2026 Bootstrap Steering Committee election process
per the Prometheus [Governance document](../../GOVERNANCE.md). This is the
**bootstrap election** described in
[Migrating from the old governance](../../GOVERNANCE.md#migrating-from-the-old-governance):
it elects all 7 seats at once, after which the new governance comes into
effect.

The 3 highest-voted candidates will serve **2-year terms**; the next 4 will
serve **1-year terms** to establish the staggered cycle.

Election schedule:

* **18 May 2026** — call for nominations opens; voters list published; this
  document announced
* **8 June 2026 23:59 UTC** — final deadline to submit nominations
* **12 June 2026** — official nominees list published; deadline to request a
  voter eligibility review
* **15 June 2026 to 29 June 2026** — voting takes place
* **6 July 2026** — results announced; new Steering Committee seated;
  governance comes into effect

We highly encourage participation in this election cycle to ensure that the
community is well-represented by the Steering Committee.

# TL;DR

* If you are willing to nominate yourself: read the
  [Governance document](../../GOVERNANCE.md) and confirm you are ready for the
  commitment. Submit a candidate page based on the
  [nomination template](./nomination-template.md) before 8 June 2026 23:59
  UTC.
* If you are an active community member: confirm that you are on the
  [voters list](./voters-list/write-access.csv), or request a review via
  [this form](https://forms.gle/tasSbJ9cQdsfj5GV9) before 12 June 2026 23:59
  UTC if you believe you should be eligible but are not listed.
* Vote between 15 June 2026 and 29 June 2026 23:59 UTC using the link the
  Election Officers will send to eligible voters.
* Keep being awesome and contributing to the project!

# Vacancies

There is no sitting Steering Committee — the new governance comes into effect
once this bootstrap election concludes. **All 7 seats are up for election in
this cycle.**

Per the
[migration section of the Governance document](../../GOVERNANCE.md#migrating-from-the-old-governance):

* The 3 highest-voted candidates will serve **2-year terms** (until the
  regular 2028 election).
* The next 4 highest-voted candidates will serve **1-year terms** (until the
  regular 2027 election), establishing the staggered cycle.

# Voting process

Anyone can track the 2026 election process via the election announcements
issue (TBD — to be opened in
[prometheus/governance](https://github.com/prometheus/governance/issues) by
the Election Officers; questions and progress updates will live there).

Per the migration section of the Governance document, the Election Officers
for the bootstrap are picked by the existing Prometheus team members (rather
than by a Steering Committee, since none exists yet). For this election they
are:

* **Arianna Vespri** ([@vesari](https://github.com/vesari))
* **Goutham Veeramachaneni** ([@gouthamve](https://github.com/gouthamve))
* **Joe Adams** ([@sysadmind](https://github.com/sysadmind))

All three are Organization Members in good standing, are eligible to vote,
are not running in this election, and have made a public promise of
impartiality.

We strive for transparency in the election process and hold the community's
interests as our priority. All documents and assets related to the 2026
election process will be public, and notes for any meetings held as part of
this process will be recorded and distributed.

Per the [Governance document](../../GOVERNANCE.md#voting-procedure),
elections are held using a ranked **Condorcet** voting method. For this
election the Election Officers have selected
[CIVS](https://civs1.civs.us/) (Condorcet Internet Voting Service), using
**Condorcet with Instant Runoff Voting (IRV) as the tie-breaker**. CIVS
supports ranked Condorcet ballots and provides standard tie-breaking via IRV;
voting is private (no individual's vote is revealed).

# Nominations

Per [the Governance document](../../GOVERNANCE.md#candidate-eligibility),
Organization Members eligible to vote may stand for election. **Candidates
must self-nominate.** There are no term limits.

To nominate yourself, copy [`nomination-template.md`](./nomination-template.md)
to `candidate-<githubid>.md` in this directory, fill it out, and submit it as
a Pull Request to the
[prometheus/governance](https://github.com/prometheus/governance) repository
before **8 June 2026 23:59 UTC**.

The template asks for:

* Full name and GitHub ID
* Employer and any other affiliations that materially direct your Prometheus
  work (this matters for the company-cap rule below)
* An eligibility checkbox confirming you have read the Governance document
  and the Code of Conduct
* The Prometheus teams and sub-projects you are a member of
* A narrative of what you have done and what you would push on as a Steering
  Committee member
* Links so voters can learn more about you

Nominees should also send a preferred contact email address to the Election
Officers — it will be kept private and used only for candidate communications
as the election process proceeds. Contact via the email addresses listed on
the Election Officers' GitHub profiles, or by tagging all three on the
nomination PR.

# Voter Eligibility

All [Organization Members](../../ROLES.md#member) are eligible to vote. The
voter roll for this election is captured under [`voters-list/`](./voters-list/)
in this directory:

* [`write-access.csv`](./voters-list/write-access.csv) — canonical voter
  list, derived from accounts with write access across the
  [prometheus](https://github.com/prometheus) and
  [prometheus-community](https://github.com/prometheus-community) GitHub
  organizations.
* [`write-access-raw.csv`](./voters-list/write-access-raw.csv) — full
  per-org / per-repo provenance behind the canonical list.
* [`manifest.json`](./voters-list/manifest.json) — record of every API call
  used to assemble the list, with timestamps and counts.
* [`main.go`](./voters-list/main.go) (with `go.mod` / `go.sum`) — the
  `collect-write-access` tool used to generate the lists, so anyone can
  reproduce the result.

Per the migration section of the Governance document, eligibility for the
bootstrap election may include members and maintainers who are inactive
under the regular activity requirements; that is intentional, and they may
still participate in this vote.

If you believe you should be eligible to vote but are not on the list,
please submit an eligibility review request via this form:
[https://forms.gle/tasSbJ9cQdsfj5GV9](https://forms.gle/tasSbJ9cQdsfj5GV9),
before **12 June 2026 23:59 UTC**. Eligibility review requests are private;
only the Election Officers will have access to this information.

# Vote

Voting will be conducted via [CIVS](https://civs1.civs.us/) using ranked
Condorcet with IRV tie-breaking. Eligible voters listed in
[`write-access.csv`](./voters-list/write-access.csv) will receive a unique
voting link by email at the address associated with their GitHub account (or
the address provided through the eligibility-review process).

Each voter ranks the candidates in order of preference. CIVS computes the
Condorcet winner(s) and uses IRV to break any cycles or ties.

The Election Officers will accept late requests to re-send the voting link.
We encourage pre-registration via the eligibility-review process to minimize
the effort required to run this election.

Voting is private: nobody — including the Election Officers — will know any
individual's vote.

# Results

Voting closed at **29 June 2026 23:59 UTC**. The following candidates were
elected to the inaugural Steering Committee:

| Member               | Organization     | Term ends |
| -------------------- | ---------------- | --------- |
| Bartłomiej Płotka    | Google           | July 2028 |
| Ben Kochie           | Reddit, Inc.     | July 2028 |
| Bryan Boreham        | Grafana Labs     | July 2028 |
| Arthur Sens          | Grafana Labs     | July 2027 |
| Jan Fajerski         | Red Hat          | July 2027 |
| Kemal Akkoyun        | Datadog          | July 2027 |
| Solomon Jacobs       | Checkmk GmbH     | July 2027 |

Per the migration section of the Governance document:

* The **top 3** by Condorcet ranking serve **2-year terms** (through the
  regular 2028 election cycle): Bartłomiej Płotka, Ben Kochie, and Bryan
  Boreham.
* The **next 4** serve **1-year terms** (through the regular 2027 election
  cycle): Arthur Sens, Jan Fajerski, Kemal Akkoyun, and Solomon Jacobs.

The company-representation cap was applied per the
[Governance document](../../GOVERNANCE.md#limitations-on-company-representation):
no more than 2 seats may be held by employees or contractors of the same
organization (or conglomerate).

With these results announced, the new Steering Committee is seated, the
[Committee members](../../GOVERNANCE.md#committee-members) table in
GOVERNANCE.md has been filled in, and the new governance comes into effect.

# Schedule

| Date                       | Activity                                                                                                                                       |
|----------------------------|------------------------------------------------------------------------------------------------------------------------------------------------|
| 18 May 2026                | This document announced; call for nominations opens; voters list published; announced via [prometheus-developers](https://groups.google.com/g/prometheus-developers) and [prometheus-users](https://groups.google.com/g/prometheus-users) |
| 8 June 2026 23:59 UTC      | End of call for nominations                                                                                                                    |
| 12 June 2026               | List of nominees finalized and announced; deadline to request a voter eligibility review                                                       |
| 15 June 2026               | Voting period begins (CIVS links sent)                                                                                                         |
| 29 June 2026 23:59 UTC     | Voting ends                                                                                                                                    |
| 6 July 2026                | Results announced; new Steering Committee seated; governance comes into effect                                                                 |
