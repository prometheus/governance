-------------------------------------------------------------
name: Bartłomiej Płotka
ID: bwplotka
info:
  employer: Google
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
- Maintainer of:
- https://github.com/prometheus/prometheus
- https://github.com/prometheus/client_golang
- https://github.com/prometheus/test-infra
- https://github.com/prometheus/prom2json
- https://github.com/prometheus/proposals
- https://github.com/prometheus/compliance
- Co-lead of [OpenMetrics 2.0 Working Group](https://docs.google.com/document/d/1FCD-38Xz1-9b3ExgHOeDTQUKUatzgj5KbCND9t-abZY/edit?tab=t.lvx6fags1fga#heading=h.uaaplxxbz60u)

## What I have done

I started using and learning Prometheus in my job in 2016, and since then, I have fallen in love with the project's community, principles, and direction. In 2017, I co-founded the open-source Thanos project, which uses Prometheus components to horizontally scale metric storage and consumption. My goal was to bring core Prometheus concepts and APIs to high-scale users and ensure Prometheus remains relevant and useful for the wider CNCF ecosystem. I'd like to think that this work brought more users and organizations into the ecosystem, helped vendors recognize the importance of Prometheus APIs, and improved the Prometheus project along the way (e.g., through the [addition of streaming remote read](https://prometheus.io/blog/#2019-10-10-remote-read-meets-streaming), and numerous TSDB and PromQL engine improvements and optimizations). After joining Google in 2023, my focus shifted directly to Prometheus to ensure the foundational parts of the Prometheus ecosystem stay healthy and useful. Key highlights of my contributions include:
* [TSDB Appender improvements](https://github.com/prometheus/prometheus/issues/17632)
* [Native Histograms with custom buckets](https://docs.google.com/document/d/1mpcSWH1B82q-BtJza-eJ8xMLlKt6EJ9oFGH325vtY1Q/edit?tab=t.0#heading=h.ueg7q07wymku)
* The [Created/Start timestamp concept](https://www.youtube.com/watch?v=nWf0BfQ5EEA&t=1s) and delta-type metrics.
* Prombench improvements.
* Innovations such as [semconv adoption and PromQL normalization](https://www.youtube.com/watch?v=Rw4c7lmdyFs), piped PromQL flavors, and [native histogram translations](https://github.com/prometheus/prometheus/issues/16948).

[In 2019, I joined the Prometheus team](https://github.com/prometheus/docs/pull/1374) and contributed [480 commits to the main repo](https://github.com/prometheus/prometheus/commits?author=bwplotka), released 5 versions of Prometheus ([2.x](https://github.com/prometheus/prometheus/blob/release-2.55/RELEASE.md) and [3.x](https://github.com/prometheus/prometheus/blob/main/RELEASE.md)), and managed [12 releases of client_golang](https://groups.google.com/g/prometheus-announce/search?q=client_golang%20p%C5%82otka). Over the last 3 years alone, I have averaged [~200 contributions per month](https://prometheus.devstats.cncf.io/d/48/users-statistics-by-repository-group?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-users=All&from=now-3y&to=now-30d), and I plan to maintain this level of activity this year, too.

I believe that the most important parts of Prometheus are its APIs. My passion has been ensuring they remain clean, well-defined, discoverable, and relevant. To this end, I helped author and drive the [Remote Write 1.0](https://groups.google.com/g/prometheus-announce/c/6xnGh7rcZKU/m/z7xFzHqOAwAJ), [Remote Write 2.0](https://prometheus.io/docs/specs/prw/remote_write_spec_2_0/), [PrometheusProto](https://github.com/prometheus/docs/pull/2836), and [OpenMetrics 2.0](https://groups.google.com/g/prometheus-announce/c/Io9AdjtwRCQ/m/NtedgUrOAgAJ) specifications.

Between 2020 and 2025, I served as a Tech Lead for the former CNCF TAG Observability, where I represented Prometheus and championed the importance of the metrics pillar, which included guiding the CNCF community by co-authoring the [Observability Whitepaper](https://github.com/cncf/tag-observability/blob/main/whitepaper.md).

However, I am probably most proud of my work as a mentor. I have mentored over 22 students through [CNCF mentorship programs like LFX and GSoC](https://github.com/search?q=repo%3Acncf%2Fmentoring+bwplotka+&type=pullrequests) across the Prometheus and Thanos projects. This work has resulted in key new features, and more importantly, has helped foster active contributors and even maintainers in the ecosystem (including Arthur, Saswata, Ben, Harshitha, and Prem, among others! ♥️). We even organized virtual [mentees meetups](https://www.youtube.com/watch?v=s9L0XukF6jQ) at one point.

Finally, I actively support the community by presenting numerous KubeCon talks, facilitating Contributor Workshops (such as [KubeCon 2026](https://docs.google.com/presentation/d/1OGG5WhHmccl24kv-Ed2EadxeNACfk5fkSYjEBjYfqXU/edit?slide=id.g3d1a0cc3f6a_0_365#slide=id.g3d1a0cc3f6a_0_365) and [KubeCon 2025](https://docs.google.com/presentation/d/1Kdjq99CUpzmH5cJTrlNAaNdy5gyWr2wMfqf3CjZ9zps/edit?slide=id.p#slide=id.p)), and organizing Prometheus DevSummits (both in-person events at the Red Hat and Google offices, and shepherding virtual ones). I also helped organize PromCon 2025 and PromCon 2026 at the Google office.

## What I'll do

* My primary priority is **project health**. We are witnessing a challenging cycle where the current economic climate and the rise of LLMs place significant stress on OSS communities-- the increased amount content that's being produced has to be verified and handled, human developers productivity got impact, given new tools and inceased responsibilities. My plan is to work with others to ensure we:This impacts both the increased volume of contributions that must be verified and handled, and developer productivity under new tools and increased responsibilities. To address this, I plan to work with others to:
  * Build appropriate tooling, agents, skills, and testing harnesses to help the community adapt to these changes in software development.
  * Expand our automatic E2E testing and benchmarking strategy, building on top of the existing Prombench suite and extending it to other sub-projects (e.g., Alertmanager).
  * Promote project needs and opportunities across various channels (e.g., CNCF, KubeCon, blogs, and social media).
  * Foster and mentor the next generation of engineers and maintainers.
  * Support innovation while maintaining the high quality of the core product. Innovation is vital to keeping the project relevant and attracting new talent to help sustain it over the long term.

* Improving **ecosystem health** is another key focus of mine. I want to strengthen collaboration with third-party software and projects like exporters, OpenTelemetry, Cortex, Mimir, and Thanos. How can we work together more effectively and support one another?
  * Drive the creation of more stable coding APIs, encourage shared libraries (similar to Thanos PromQL), and champion discussions and tighter collaboration between these projects and Prometheus.
  * Ensure new features (e.g., Start Timestamp, Delta, and OpenTelemetry semconv) integrate seamlessly without harming long-term storage solutions.
  * Facilitate the adoption and progress of new and upcoming versions of Prometheus protocols (e.g., Remote Write and OpenMetrics).
  * Create healthy avenues for promoting and discovering community contributions (e.g., exploring ideas like community SDK registries or an exporters marketplace).
  * Ensure there are robust ways to vet third-party software for Prometheus compatibility (e.g., expanding the compliance test suites).
  * Foster interoperatibility and health of a pull-based metric collection model, given OpenTelemetry popularity.

## Resources about me

- GitHub: [@bwplotka](https://github.com/bwplotka)
- Website: https://bwplotka.dev/
- Talks:
  - [KubeCon: "What We Learned After Mentoring 30+ Mentees"](https://docs.google.com/presentation/d/1lZkSIBWaGr5h9aLpZo26A81U-lD-JBc0KdVMjUDMo_c/edit?slide=id.p#slide=id.p)
  - [KubeCon: "Prometheus v3 One Year in: OpenMetrics 2.0 and more!](https://docs.google.com/presentation/d/1kvlf1xz47bgpSvFvxy6C_v-ZM50wiLJM3EuUpE-_ce0/edit?slide=id.p#slide=id.p)
  - [PromCon: "Everything you need to know about OpenMetrics 2.0!"](https://www.youtube.com/watch?v=KTdhXHY-Hqo)
  - [PromCon: "LT: Scrape Trolley Dilemma"](https://docs.google.com/presentation/d/1jKrUklPdAor9292HrPWtJkIa6ruUhOGo9IFO7fNj-DE/edit?slide=id.p#slide=id.p)
  - [KubeCon: "How To Rename Metrics Without Impacting Somebody’s Observability"](https://www.youtube.com/watch?v=Rw4c7lmdyFs)
  - [Older talks](https://www.bwplotka.dev/about/#public-talks)
- Writing:
  - https://prometheus.io/blog/#2026-02-14-modernizing-prometheus-composite-samples
  - https://prometheus.io/blog/#2025-09-22-promcon-2025
  - https://prometheus.io/blog/#2021-11-16-agent
  - https://prometheus.io/blog/#2019-10-10-remote-read-meets-streaming
- Social / contact:
  - https://www.linkedin.com/in/bwplotka/
  - https://bsky.app/profile/bwplotka.dev
  - http://x.com/bwplotka
  - https://bwplotka.dev/
