---
name: Kemal Akkoyun
ID: kakkoyun
info:
  employer: Datadog
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
- `prometheus/client_golang` - maintainer.
- `prometheus/promu` - maintainer.
- `prometheus/test-infra` - maintainer.

## What I have done

I have contributed around Prometheus since 2020, mostly on the instrumentation and tooling
side rather than the server itself. I want to be honest about that scope up front: the people
who carry the TSDB, PromQL, and the core server are not me, and I have a lot of respect for
that work. My lane has been the libraries and tools that sit around it.

I maintain [`prometheus/client_golang`](https://github.com/prometheus/client_golang), the Go
instrumentation library, together with Bartłomiej Płotka. I took it over
from Björn Rabenstein in 2021
([client_golang#873](https://github.com/prometheus/client_golang/pull/873)). It is the library
most Go services use to expose metrics, so a lot of people meet Prometheus through it. The work
is mostly unglamorous: reviewing contributions, cutting releases, keeping the API stable and the
dependencies honest, and helping newcomers get their first change merged.

I also help keep two pieces of tooling alive that everyone relies on and no one thinks about:
[`prometheus/promu`](https://github.com/prometheus/promu), the build and release tool used across
our repositories, and [`prometheus/test-infra`](https://github.com/prometheus/test-infra), which
runs prombench. It is not heavy or constant work. It is mostly keeping the lights on - merging the occasional
fix, bumping dependencies, and making sure these tools do not quietly rot. That is part of why
it matters. They are the bits everything depends on and few people volunteer for, and I have
been glad to be the one who watches them.

The rest of my background is in the wider ecosystem. I was a maintainer of Thanos, though I am
emeritus now and no longer active there. I helped build Parca for continuous profiling and
maintained Observatorium. These days I work at Datadog on Go instrumentation, eBPF, and
OpenTelemetry. That path through metrics, profiling, and tracing is why the seams between
Prometheus and the rest of the observability world interest me.

On the community side, I have spoken at PromCon and other conferences since 2020, and I have
mentored for Prometheus and Thanos through CNCF LFX and Google Summer of Code. Getting a new
contributor to their first merged pull request is the part I care about most.

## What I'll do

I will be modest about promises. The Steering Committee's job is governance, project health, and
cross-team coordination, and, within the charter, helping shape the project's broader vision and
direction. It is not about setting technical direction inside teams. Here is what I would bring.

**Industry perspective.** I have spent years in observability - metrics, profiling, and tracing -
across Red Hat, Polar Signals, and now Datadog. I have a broad view of where the field is heading and where comparable projects have stumbled.
I would bring that context to the committee's work of keeping Prometheus' priorities aligned with
how the wider observability world is moving, without second-guessing maintainers on technical calls.

**OpenTelemetry expertise.** I work on OpenTelemetry day to day. I am not proposing to make the
OTel and Prometheus boundary my single cause, but it is an area where the project regularly needs
people who genuinely understand both sides, and I am one of them. I would lend that expertise where
it helps: reviewing interoperability questions, helping carry OTel-related work, and being an
informed voice when those topics come up.

**The instrumentation surface.** The client libraries are how most people first touch Prometheus,
and several are kept alive by one or two volunteers. I would push for that layer to have enough
maintainers and enough consistency across languages to stay dependable. It is the part I know best.

**Bringing in contributors.** Prometheus depends on a steady supply of people, and mentorship is
how most of us arrived. I would keep investing in programs like LFX and Google Summer of Code, and
in lowering the bar for first-time and non-code contributions, so the less visible corners of the
project find owners.

I would also add a perspective from outside the handful of companies most maintainers work for.
Seats are held by individuals, not employers, and a committee is healthier for a range of vantage
points. I should be honest about one thing: Datadog has not historically been a significant
contributor to Prometheus. I hope to help change that — not by acting as a company proxy, but
because companies that benefit from Prometheus should put more back into it, and I am in a
position to push for that from inside Datadog. I will
represent the project, not a company. I will not claim to be the strongest candidate on the
ballot; I would be glad to serve, and just as glad to support whoever the community chooses.

## Resources about me

- GitHub: [@kakkoyun](https://github.com/kakkoyun)
- Website: [kakkoyun.me](https://kakkoyun.me/)
- Talks / writing:
  - [PromCon EU 2022 - Best Practices and Pitfalls of Instrumenting Your Cloud-Native Application](https://promcon.io/2022-munich/speakers/kemal-akkoyun/)
  - [PromCon 2020 (Online)](https://promcon.io/2020-online/speakers/kemal-akkoyun/)
  - [The Zen of Prometheus](https://kakkoyun.me/talks/the-zen-of-prometheus/)
  - [Blog and newsletter](https://kakkoyun.me/posts/)
- Social / contact:
  - [Bluesky](https://bsky.app/profile/kakkoyun.me)
  - [LinkedIn](https://www.linkedin.com/in/kakkoyun/)
  - [Twitter/X](https://x.com/kakkoyun_me)
