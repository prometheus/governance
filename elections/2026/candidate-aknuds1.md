---
name: Arve Knudsen
ID: aknuds1
info:
  employer: Grafana Labs
  affiliation: None
---

## Eligibility

- [x] I have read the
  [Governance document](https://github.com/prometheus/governance/blob/main/GOVERNANCE.md)
  and the
  [Code of Conduct](https://github.com/prometheus/docs/blob/main/CODE_OF_CONDUCT.md),
  and agree to uphold them if elected.

## Teams

- `prometheus/prometheus` - general maintainer.
- `prometheus/otlptranslator` - maintainer.

## What I have done

I have been active in open source software development since the early 2000s.
Currently employed as a Staff Backend Engineer at Grafana Labs, where I've been working on Grafana Mimir since 2021.
To improve on Mimir's Prometheus foundation, I began contributing to Prometheus in March 2023.
I now have more than 200 merged commits to `prometheus/prometheus` and substantial work across `prometheus/common`, `prometheus/client_golang`, `prometheus/otlptranslator`, and `prometheus/docs`.
In March 2026 I moved from `otlptranslator` part maintainer to general maintainer of `prometheus/prometheus` ([#18211](https://github.com/prometheus/prometheus/pull/18211)).
Being a long time Grafana Mimir team member, I have deep operational empathy for downstream projects running Prometheus at massive scale.

My work has mainly focused on making Prometheus more reliable/performant, easier to operate, and better integrated with the OpenTelemetry (AKA OTel) ecosystem.
Selected areas of work:

- **TSDB performance and correctness.**
  I have worked on the storage engine's memory management and read performance, including faster head-chunk lookups, more efficient mmap handling, smoother WAL replay, and pool-object lifecycle fixes.
  These are the kinds of changes that are not always visible to users, but matter a lot for the performance and operational reliability of the server
  ([#18541],
  [#18302],
  [#18272],
  [#18306],
  [#17895],
  [#17879]).

- **OTel interoperability.**
  I have spent much of the last two years working across the Prometheus and OTel boundary.
  As the main OTel developer for Grafana Mimir, I drove the OTLP ingest endpoint to General Availability.
  On the Prometheus side, I was a primary driver in stabilizing the OTLP receiver for General Availability - fixing semantic edge cases, handling translation, and greatly improving performance.
  In the OTel ecosystem, I have contributed directly to the `opentelemetry-collector-contrib` repository to fix data-loss bugs in the Prometheus Remote Write receiver, stop response leaks in the target allocator, and stabilize integration tests.
  This cross-community work is critical for ensuring clear behaviour and compatibility for both projects (Prometheus PRs:
  [prometheus/proposals#38],
  [#16730],
  [#17345],
  [#17344],
  [#16776],
  [#16737],
  [#17400],
  [#17860],
  [#17917];
  OTel PRs: [otel#45506],
  [otel#45151],
  [otel#44921],
  [otel#32148]).

- **PromQL `info()` function.**
  I co-authored the experimental [`info()` PromQL function](https://prometheus.io/blog/2025/12/16/introducing-info-function/) with Björn Rabenstein, which in particular makes including OTel resource attributes in queries much easier.
  I also wrote user-facing documentation and a blog post to make the feature easier to understand and adopt
  ([prometheus/proposals#37],
  [#14495],
  [#17817],
  [#17932],
  [#18352],
  [#18354],
  [#18367],
  [prometheus/docs#2777],
  [prometheus/docs#2776],
  [prometheus/docs#2756]).

- **Cross-cutting API and validation work.**
  I threaded `context.Context` through core read-path APIs so queries can be effectively canceled.
  I also helped consolidate metric- and label-name validation around shared `prometheus/common` types, reducing divergence between repositories
  ([#12660],
  [#12834],
  [#12667],
  [#12666],
  [#12665],
  [#16928],
  [#16806],
  [prometheus/common#806],
  [prometheus/common#799],
  [prometheus/client_golang#1859]).

- **Project health and release work.**
  I have put a lot of time into the maintenance work that keeps the project moving: fixing various flaky tests, rolling out linters in `prometheus/common`, helping with the Renovate migration, cutting `prometheus/client_golang` v1.23.1 and v1.23.2, and improving the PR template with a concrete release-notes example
  ([#17985],
  [#17965],
  [#17938],
  [#17934],
  [#17933],
  [#18085],
  [#18726],
  [prometheus/common#818],
  [prometheus/common#819],
  [prometheus/common#820],
  [prometheus/common#821],
  [prometheus/common#822],
  [prometheus/common#823],
  [#16654],
  [#17776],
  [#17777],
  [prometheus/client_golang#1867],
  [prometheus/client_golang#1870],
  [#17721]).

## What I'll do

I see the most important challenges ahead, in the years to come, for the Prometheus project to be: Keeping the project healthy and active, and sufficiently equipped to adapt with the evolving observability ecosystem.
My goal is to use my experience in the trenches to keep the project running smoothly, ensure maintainers' concerns are consistently recognized, and to maintain Prometheus' relevance within the ecosystem.
The key areas I would focus on are:

- **Keeping Prometheus relevant and adaptable.**
  The observability ecosystem is moving fast, especially with the rise of OTel.
  While the SC doesn't dictate technical architecture, it does shape how we collaborate outside our own borders.
  I want to ensure we have the governance structures, cross-project liaisons, and community focus needed to adapt to industry shifts.

- **Strengthening cross-ecosystem collaboration.**
  The path for OTel integration spans `prometheus/prometheus`, `prometheus/otlptranslator`, `prometheus/common`, and OTel's own repositories.
  Having worked extensively across this boundary, I want to foster improved interoperability between OTel and Prometheus and joint design discussions with downstream communities and the wider observability ecosystem.

- **Reducing maintainer toil through automation.**
  Since joining the Prometheus project, a lot of my time has been spent on making the development workflow run smoother and less frustrating - I've for example helped migrate dependency updates to Renovate, and consistently fix flaky tests (especially making timing-dependent tests deterministic).
  I'd push the project to keep investing in this layer: Dependency hygiene, CI health, release-process automation, and where it actually helps, AI-assisted triage and review tooling.
  My goal is to reduce maintainer toil as much as possible, so they can focus on the interesting bits.

- **Ensuring fair stewardship of all Prometheus project resources.**
  I will draw on my extensive experience in the OSS and more general software engineering universe to ensure continued fair and balanced stewardship of the Prometheus project and its associated resources.

## Resources about me

- GitHub: [@aknuds1](https://github.com/aknuds1)
- Website: [arveknudsen.com](https://arveknudsen.com/)
- Talks / writing:
  - [PromCon 2025 - OpenTelemetry developments in Prometheus over the last year](https://promcon.io/2025-munich/talks/opentelemetry-developments-in-prometheus-over-the-last-year/)
  - [PromCon 2024 - The Future of Metadata in Prometheus: Enhancing Storage and Usability](https://promcon.io/2024-berlin/talks/the-future-of-metadata-in-prometheus-enhancing-storage-and-usability/)
  - [PromCon 2024 - Practical OpenTelemetry with Prometheus 3.0](https://promcon.io/2024-berlin/talks/practical-opentelemetry-with-prometheus-30/)
  - [Introducing the Experimental info() Function](https://prometheus.io/blog/2025/12/16/introducing-info-function/)
- Social / contact:
  - [BlueSky](https://bsky.app/profile/aknudsen.bsky.social)
  - [LinkedIn](https://www.linkedin.com/in/arve-knudsen-b338741b/)

<!-- PR References -->
[#12660]: https://github.com/prometheus/prometheus/pull/12660
[#12665]: https://github.com/prometheus/prometheus/pull/12665
[#12666]: https://github.com/prometheus/prometheus/pull/12666
[#12667]: https://github.com/prometheus/prometheus/pull/12667
[#12834]: https://github.com/prometheus/prometheus/pull/12834
[#14495]: https://github.com/prometheus/prometheus/pull/14495
[#16654]: https://github.com/prometheus/prometheus/pull/16654
[#16730]: https://github.com/prometheus/prometheus/pull/16730
[#16737]: https://github.com/prometheus/prometheus/pull/16737
[#16776]: https://github.com/prometheus/prometheus/pull/16776
[#16806]: https://github.com/prometheus/prometheus/pull/16806
[#16928]: https://github.com/prometheus/prometheus/pull/16928
[#17344]: https://github.com/prometheus/prometheus/pull/17344
[#17345]: https://github.com/prometheus/prometheus/pull/17345
[#17400]: https://github.com/prometheus/prometheus/pull/17400
[#17721]: https://github.com/prometheus/prometheus/pull/17721
[#17776]: https://github.com/prometheus/prometheus/pull/17776
[#17777]: https://github.com/prometheus/prometheus/pull/17777
[#17817]: https://github.com/prometheus/prometheus/pull/17817
[#17860]: https://github.com/prometheus/prometheus/pull/17860
[#17879]: https://github.com/prometheus/prometheus/pull/17879
[#17895]: https://github.com/prometheus/prometheus/pull/17895
[#17917]: https://github.com/prometheus/prometheus/pull/17917
[#17932]: https://github.com/prometheus/prometheus/pull/17932
[#17933]: https://github.com/prometheus/prometheus/pull/17933
[#17934]: https://github.com/prometheus/prometheus/pull/17934
[#17938]: https://github.com/prometheus/prometheus/pull/17938
[#17965]: https://github.com/prometheus/prometheus/pull/17965
[#17985]: https://github.com/prometheus/prometheus/pull/17985
[#18085]: https://github.com/prometheus/prometheus/pull/18085
[#18272]: https://github.com/prometheus/prometheus/pull/18272
[#18302]: https://github.com/prometheus/prometheus/pull/18302
[#18306]: https://github.com/prometheus/prometheus/pull/18306
[#18352]: https://github.com/prometheus/prometheus/pull/18352
[#18354]: https://github.com/prometheus/prometheus/pull/18354
[#18367]: https://github.com/prometheus/prometheus/pull/18367
[#18541]: https://github.com/prometheus/prometheus/pull/18541
[#18726]: https://github.com/prometheus/prometheus/pull/18726
[otel#32148]: https://github.com/open-telemetry/opentelemetry-collector-contrib/pull/32148
[otel#44921]: https://github.com/open-telemetry/opentelemetry-collector-contrib/pull/44921
[otel#45151]: https://github.com/open-telemetry/opentelemetry-collector-contrib/pull/45151
[otel#45506]: https://github.com/open-telemetry/opentelemetry-collector-contrib/pull/45506
[prometheus/client_golang#1859]: https://github.com/prometheus/client_golang/pull/1859
[prometheus/client_golang#1867]: https://github.com/prometheus/client_golang/pull/1867
[prometheus/client_golang#1870]: https://github.com/prometheus/client_golang/pull/1870
[prometheus/common#799]: https://github.com/prometheus/common/pull/799
[prometheus/common#806]: https://github.com/prometheus/common/pull/806
[prometheus/common#818]: https://github.com/prometheus/common/pull/818
[prometheus/common#819]: https://github.com/prometheus/common/pull/819
[prometheus/common#820]: https://github.com/prometheus/common/pull/820
[prometheus/common#821]: https://github.com/prometheus/common/pull/821
[prometheus/common#822]: https://github.com/prometheus/common/pull/822
[prometheus/common#823]: https://github.com/prometheus/common/pull/823
[prometheus/docs#2756]: https://github.com/prometheus/docs/pull/2756
[prometheus/docs#2776]: https://github.com/prometheus/docs/pull/2776
[prometheus/docs#2777]: https://github.com/prometheus/docs/pull/2777
[prometheus/proposals#37]: https://github.com/prometheus/proposals/pull/37
[prometheus/proposals#38]: https://github.com/prometheus/proposals/pull/38
