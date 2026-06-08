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

I've started using and learning Prometheus in my job in 2016 and since then I fell in love with the project community, principles and direction. In 2017, I co-founded OSS Thanos project that used Prometheus pieces to horizontally scale metric storage and consumption. My goal was to bring core Prometheus concepts and APIs for high scale users, and ensure Prometheus stays relevant and useful for the wider CNCF ecosystem. I'd like to think that this work brought more users and organizations to the ecosystem, helped vendors to notice the importance of Prometheus APIs and improved Prometheus project on the way (e.g. [addition of streaming remote read](https://prometheus.io/blog/#2019-10-10-remote-read-meets-streaming), numerous TSDB and PromQL engine improvements and optimizations). After that, since I joined Google in 2023 my focus shifted to Prometheus to ensure foundational parts of Prometheus ecosystem stay healthy and useful. Highlight of the things I helped with:
* [TSDB Appender improvements](https://github.com/prometheus/prometheus/issues/17632)
* [Native Histograms with custom buckets](https://docs.google.com/document/d/1mpcSWH1B82q-BtJza-eJ8xMLlKt6EJ9oFGH325vtY1Q/edit?tab=t.0#heading=h.ueg7q07wymku).
* [Created/Start timestamp concept](https://www.youtube.com/watch?v=nWf0BfQ5EEA&t=1s) and delta type metrics.
* Prombench improvements.
* Innovations like [semconv adoption and PromQL normalization](https://www.youtube.com/watch?v=Rw4c7lmdyFs), piped PromQL flavors, [native histogram translations](https://github.com/prometheus/prometheus/issues/16948).

[In 2019 I joined the Prometheus team](https://github.com/prometheus/docs/pull/1374) and contributed [480 commits to the main repo](https://github.com/prometheus/prometheus/commits?author=bwplotka), released 5 versions of Prometheus ([2.x](https://github.com/prometheus/prometheus/blob/release-2.55/RELEASE.md),[3.x](https://github.com/prometheus/prometheus/blob/main/RELEASE.md)) and [12 releases of client_golang](https://groups.google.com/g/prometheus-announce/search?q=client_golang%20p%C5%82otka). For the last 3 years alone, I've averaged [~200 contributions per month](https://prometheus.devstats.cncf.io/d/48/users-statistics-by-repository-group?orgId=1&var-period=m&var-metric=contributions&var-repogroup_name=All&var-users=All&from=now-3y&to=now-30d) and I have plans to maintain this year too.

I think the most important part of Prometheus are its APIs. My passion was to ensure they stay clean, well-defined, discoverable and relevant. For those reasons I helped with the [Remote Write 1.0](https://groups.google.com/g/prometheus-announce/c/6xnGh7rcZKU/m/z7xFzHqOAwAJ), [Remote Write 2.0](https://prometheus.io/docs/specs/prw/remote_write_spec_2_0/), [PrometheusProto](https://github.com/prometheus/docs/pull/2836) and [OpenMetrics 2.0](https://groups.google.com/g/prometheus-announce/c/Io9AdjtwRCQ/m/NtedgUrOAgAJ) specifications.

Between 2020-2025 I've been a tech lead for the former CNCF TAG Observability where I represented Prometheus and ensured priority of the metrics pillar e.g. guiding CNCF community with the [Observability Whitepaper](https://github.com/cncf/tag-observability/blob/main/whitepaper.md).

However, probably the most grateful I am with my work as a mentor. I did over 22 of [CNCF mentorships in programs like LFX and GSoC](https://github.com/search?q=repo%3Acncf%2Fmentoring+bwplotka+&type=pullrequests) across Prometheus and Thanos which resulted in new features, but more importantly, active contributors and even maintainers in the ecosystem (Arthur, Saswata, Ben, Harshitha, Prem, among others! ♥️). We even had virtual [mentees meetups](https://www.youtube.com/watch?v=s9L0XukF6jQ) at one point.

Finally, I try to help as much as possible on community side with numerous KubeCon talks, Contributor workshops (e.g. [KubeCon26](https://docs.google.com/presentation/d/1OGG5WhHmccl24kv-Ed2EadxeNACfk5fkSYjEBjYfqXU/edit?slide=id.g3d1a0cc3f6a_0_365#slide=id.g3d1a0cc3f6a_0_365), [KubeCon25](https://docs.google.com/presentation/d/1Kdjq99CUpzmH5cJTrlNAaNdy5gyWr2wMfqf3CjZ9zps/edit?slide=id.p#slide=id.p)) and Prometheus DevSummit (in-person ones in Red Hat and Google offices and shepherding virtual ones). I've also helped organize PromCon 2025 and PromCon 2026 in Google office.

## What I'll do


* The main priority for me is **project health**. We see an interesting cycle where economic situation and LLM put a lot of stress on OSS communities, both on the amount of content being produced that has to be verified and handled, as well as, human developers productivity impact, given new tools and inceased responsibilities. My plan is to work with others to ensure we:
  * Built appropriate tooling, agents, skills and harness for community to cope with the changes in software development.
  * We expand our automatic e2e testing and benchmarking strategy, building on top of what we have with Prombench suite and expanding that to other sub-projects (e.g. Alertmanager). 
  * Promote the project needs and opportunities via various channels (e.g. CNCF, KubeCon, blog, social media).
  * Fostering mentoring new generation of engineers and maintainers.
  * Help with innovations, while ensuring high quality of the core product. Innovation is important to stay relevant and bring talents to help us maintain the project over time.

* **Ecosystem health** is another aspect that I'd love to improve. I would love to see improved collaboration with 3rd party software like exporters, OpenTelemetry, Cortex, Mimir, Thanos and so on. How can we ensure we work together better and help each other? 
  * That might mean more stable coding APIs, encourage shared libraries (e.g. like Thanos PromQL), and championing discussions and tighter collaboration between those projects and Prometheus. 
  * Ensuring new features like Start Timestamp, Delta, OpenTelemetry semconv, etc are not harming long term storage solutions.
  * Ensuring adoption and progress for the new and upcoming versions of Prometheus protocols (e.g. Remote Write and OpenMetrics)
  * Making sure we have a healthy place for promoting and discovery of (e.g. time for community SDKs, exporters marketplace?)
  * Ensuring there's a way to vet 3rd party software on Prometheus compatibility (e.g. compliance suite).

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
