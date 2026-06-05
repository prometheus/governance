-------------------------------------------------------------
name: Jan Fajerski
ID: jan--f
info:
  employer: Red Hat
  affiliation: None
-------------------------------------------------------------

## Eligibility

- [x] I have read the
  [Governance document](https://github.com/prometheus/governance/blob/main/GOVERNANCE.md)
  and the
  [Code of Conduct](https://github.com/prometheus/docs/blob/main/CODE_OF_CONDUCT.md),
  and agree to uphold them if elected.

## Teams

- Prometheus organization member; member of the
  [prometheus/default-maintainers](https://github.com/orgs/prometheus/teams/default-maintainers)
  GitHub team.
- [Prometheus v3 release coordinator](https://github.com/prometheus/prometheus/blob/main/MAINTAINERS.md).
- [Maintainer of Pushgateway](https://github.com/prometheus/pushgateway/blob/master/MAINTAINERS.md).

## What I have done

I first touched Prometheus in 2016 to monitor my own hardware at home.  In 2017 I started working on the [Prometheus manager module for Ceph](https://docs.ceph.com/en/latest/mgr/prometheus/) while working for SUSE. Prometheus was the natural fit for exposing Ceph cluster metrics, and that work led me deeper into the ecosystem. My first contribution to the Prometheus project itself came in 2019 with a small fix to the node exporter, but sustained involvement started in 2021 when I joined Red Hat and began working with Prometheus full-time.

One of my first contributions was replacing vfsgen with Go's native `embed` package for static web assets n Prometheus ([prometheus/common#343](https://github.com/prometheus/common/pull/343), [prometheus/prometheus#9719](https://github.com/prometheus/prometheus/pull/9719)). The respective Alertmanager change was also merged eventually.

The work I am most known for is coordinating the **Prometheus 3.0 release**. I shepherded the release from beta through RC to GA and subsequent patch releases, wrote the [3.0 migration guide](https://github.com/prometheus/prometheus/pull/15099) and a few breaking changes such as [renaming `holt_winters` to `double_exponential_smoothing`](https://github.com/prometheus/prometheus/pull/14930) and [changing the remote-write `enable_http2` default](https://github.com/prometheus/prometheus/pull/15219) to address head-of-line blocking. At [PromCon EU 2024 in Berlin](https://promcon.io/2024-berlin/speakers/jan-fajerski/) I kicked off the conference by releasing the 3.0 beta live on stage.

Beyond releases, I have contributed native histogram support for the Pushgateway ([UI](https://github.com/prometheus/pushgateway/pull/596), [API](https://github.com/prometheus/pushgateway/pull/611), [tests](https://github.com/prometheus/pushgateway/pull/592)) and [prom2json](https://github.com/prometheus/prom2json/pull/169) and worked on build and CI infrastructure improvements across the ecosystem: GitHub Actions migration for Alertmanager, promci improvements, and [podman crossbuild support in promu](https://github.com/prometheus/promu/pull/360).

I care about documentation and contributor experience. Recently I have done housekeeping in [prometheus/docs](https://github.com/prometheus/docs) reducing the number of open PRs to ~20 and added a few things myself (["Report an issue" button](https://github.com/prometheus/docs/pull/2906), [contribution guide](https://github.com/prometheus/docs/pull/2840)).

I have also filed governance and process proposals for [label conventions](https://github.com/prometheus/proposals/pull/65) and [label automation](https://github.com/prometheus/proposals/pull/66) to help keep our issue and PR queues manageable, and I took over the [Pushgateway maintenance](https://github.com/prometheus/pushgateway/pull/801).

## What I'll do

Prometheus is technically excellent. We ship ambitious features and our core development is strong. Where we struggle is the maintenance and review load across all project repos. Some PRs sit open for too long, issue triaging can be slow, knowledge about CI infra is a bottleneck and our documentation degrades as the code evolves.

My priorities as a Steering Committee member:

**Keep the project maintainable.** I want to extend established regular triage and review practices across the organization where required. I'm convinced we can implement low overhead conventions and automations to improve this for both reviewers and contributors.

**Encourage sustained human contribution.** Prometheus can be intimidating to new contributors. I want to improve onboarding through better documentation (the contribution guide I am working on is a start) and clearer issue and PR labeling. The new governance is great but I think we can improve further. I believe in open, transparent governance where people feel safe to disagree, propose ideas, and take ownership. As a Steering Committee member I would focus on creating the conditions for that: clear processes, consistent and transparent decision-making and making sure every contributor feels their work is valued and visible.

**Fill documentation gaps.** The Prometheus documentation is strong when it comes to configuration references. When it comes to How-Tos or tutorial style content the documentation can be much improved. While we have content like this, it is hard to find, sometimes no longer up to date or incomplete. I want to propose a major change to how we structure our documentation and keep it up to date, and also encourage for content that is focused on deployment and integration.

I am not the person to set technical direction — the project has excellent engineers for that. I see my strength in communication and I'll pay attention to the less glamorous work and processes serve the community rather than burden it. Prometheus is a fantastic project with a lovely community and I hope we can grow while maintaining what makes Prometheus and its community great.

## Resources about me

- GitHub: [@jan--f](https://github.com/jan--f)
- Talks:
  - [Prometheus 3.0 Overview — PromCon EU 2024, Berlin](https://www.youtube.com/watch?v=2pC38yWxanM)
  - [Prometheus Version 3 — FOSDEM 2025, Brussels](https://archive.fosdem.org/2025/schedule/event/fosdem-2025-6571-prometheus-version-3/) (with Bryan Boreham)
  - [Downsampling in Prometheus — PromCon EU 2025, Munich](https://www.youtube.com/watch?v=SC_9VvT5xuA)
  - [Monitoring Ceph with Prometheus — Cephalocon 2019, Barcelona](https://www.youtube.com/watch?v=Ag8pWQ_M8i0)
- Social / contact:
  - prometheus@fajerski.name
