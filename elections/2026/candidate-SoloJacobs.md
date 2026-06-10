---
name: Solomon Jacobs
ID: SoloJacobs
info:
  employer: Checkmk GmbH
  affiliation: None
---

## Eligibility

- [x] I have read the
  [Governance document](https://github.com/prometheus/governance/blob/main/GOVERNANCE.md)
  and the
  [Code of Conduct](https://github.com/prometheus/docs/blob/main/CODE_OF_CONDUCT.md),
  and agree to uphold them if elected.

## Teams

- `prometheus` - organization member.
- `prometheus/alertmanager` - maintainer.

## What I have done

I'm a software engineer at [Checkmk](https://github.com/Checkmk/checkmk), where most of our customers are system admins.
I've worked in monitoring for ~5 years but only started contributing to Prometheus in October 2025, so I'm fresh blood here.

My first contact with it was back in 2021, when I helped build our [Kubernetes monitoring](https://www.youtube.com/watch?v=fsxn7laB-IU&t=10126s).
I've come to genuinely appreciate the project since.
My employer doesn't sponsor this work.
I contribute because Prometheus aligns with what I value as an engineer.

I'm primarily a maintainer of Alertmanager.
I have about [60 commits in the project](https://github.com/prometheus/alertmanager/commits?author=SoloJacobs).
Some of the more interesting PRs I've landed are:

* serving pre-compressed UI assets and adding caching [#5133](https://github.com/prometheus/alertmanager/pull/5133), [#5113](https://github.com/prometheus/alertmanager/pull/5113)
* allowing builds with Podman [#5092](https://github.com/prometheus/alertmanager/pull/5092)
* moving the build process to `//go:embed` [#5028](https://github.com/prometheus/alertmanager/pull/5028)
* adding a config fuzzer (and fixing [the bugs](https://github.com/prometheus/alertmanager/pull/4979) it found) [#4913](https://github.com/prometheus/alertmanager/pull/4913)

The more valuable part of my work, though, is reviews.
I've also been release shepherd for Alertmanager v0.30, v0.31, and v0.32.

Some more of my open source can be found at [Robotmk](https://github.com/elabit/robotmk), where I am one of two maintainers.

## What I'll do

I think the Steering Committee serves the project best by empowering the people doing the work.
Centralizing too much would only make the committee a bottleneck.
Decisions belong as close to the contributors and teams as possible, with the committee as a trusted backstop they can escalate to whenever they need it.
Clear, documented, mechanical rules keep that backstop fair and free of politics.
With that in mind, here is where I'd focus.

**Connect Alertmanager to the wider ecosystem.**
This is my primary goal: closing the communication gap between the Alertmanager team and the rest of Prometheus.
More broadly, no component should be left isolated.

**Recognize reviews as first-class work.**
I want the work that rarely shows up in commit counts, especially reviewing, to be recognized as a real contribution rather than an afterthought.
One concrete way to give it visibility is a dedicated triage role.
`ROLES.md` counts triage as a contribution, but grants no triage permissions short of full merge rights.
A dedicated role would give newcomers a low-barrier first step toward reviewing.

**Keep reviews moving.**
Slow or stalled review burns out contributors and maintainers alike.
Concretely, we could borrow review habits from [matklad's mechanical habits](https://matklad.github.io/2025/12/06/mechanical-habits.html), though the autonomy of individual projects matters more overall.

**Surface the needs of users coming from legacy monitoring.**
I come from the world of Nagios and its successors, where Prometheus still lags in some areas.
[Colin Douch's OSMC 2025 talk](https://www.slideshare.net/slideshow/osmc-2025-onotify-a-scalable-flexible-alertmanager-by-colin-douch-pdf/284300918) is a good example: acknowledgments, downtimes, and silencing that Nagios handled natively are delegated to external systems, which fragments the on-call experience.
I want frustrations like these surfaced and heard.
One concrete step would be providing the libraries and infrastructure that let users share anonymized data, so projects can reproduce and understand real-world problems.

I'm newer to Prometheus than most on this ballot, but I bring an outside perspective and I'm glad to do the unglamorous work that keeps the project running.

- GitHub: [@SoloJacobs](https://github.com/SoloJacobs)
- Talks / writing:
  - [Master's thesis: "Split Gibbs Sampler"](https://doi.org/10.5281/zenodo.20563670), not monitoring-related, but a good way to get a feel for who I am.
- Social / contact:
  - [LinkedIn](https://www.linkedin.com/in/solomon-jacobs-53a5b2151/)
