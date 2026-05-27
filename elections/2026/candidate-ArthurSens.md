-------------------------------------------------------------
name: Arthur Sens
ID: ArthurSens
info:
  employer: Grafana Labs
  affiliation: Co-lead of Prometheus Special Interest Group in OpenTelemetry.
-------------------------------------------------------------

## Eligibility

- [x] I have read the
  [Governance document](https://github.com/prometheus/governance/blob/main/GOVERNANCE.md)
  and the
  [Code of Conduct](https://github.com/prometheus/docs/blob/main/code-of-conduct.md),
  and agree to uphold them if elected.

## Teams

- Prometheus organization member; member of the
  [prometheus/default-maintainers](https://github.com/orgs/prometheus/teams/default-maintainers)
  GitHub team.
- [Maintainer for Prometheus Server OTLP endpoint](https://github.com/prometheus/prometheus/blob/main/MAINTAINERS.md).
- Codeowner of a few utility libraries around Prometheus and OpenTelemetry:
  - [prometheus/otlptranslator](https://github.com/prometheus/prometheus/blob/main/CODEOWNERS)
  - [prometheus/opentelemetry-collector-bridge](https://github.com/prometheus/opentelemetry-collector-bridge/blob/main/.github/CODEOWNERS)
- Co-lead of Prometheus' UX/Design working group: [prometheus-community/ux-research](https://github.com/prometheus-community/ux-research/blob/main/.github/CODEOWNERS)

## What I have done

My Prometheus journey started in 2017, while I was working as a sysadmin intern in my first job. My job was to roll out Prometheus across a few thousand virtual machines and introduce good monitoring practices across the organization. Prometheus was also my first contact with open source: I closely watched strangers discuss the future of the project through GitHub issues before I ever felt ready to participate myself.

During the Covid pandemic, while recovering from burnout, I stepped out of the shadows thanks to an [LFX Mentorship](https://mentorship.lfx.linuxfoundation.org/project/de7ca1c2-2d22-4919-bef8-6cca50a54426) with the KubeVirt community. Not long after contributing to KubeVirt, I started making small documentation fixes to Prometheus and eventually became one of Bartek's GSoC mentees, working on the Created Timestamp feature for the Prometheus server. That was when my involvement with Prometheus became serious, going beyond the traditional end-user role.

As I transitioned from a sysadmin career, I learned Go by helping maintain different parts of the Prometheus community. My most notable **emeritus** positions are maintainer of [Prometheus Operator](https://github.com/prometheus-operator/prometheus-operator) and [Prometheus Client Go](https://github.com/prometheus/client_golang).

Among my current responsibilities, I am probably best known for cross-community work between Prometheus and OpenTelemetry. I have co-authored utility libraries to bring coherence to OTLP-to-Prometheus metric translation [[1](https://github.com/prometheus/otlptranslator)] across several moving parts, and to make Prometheus exporters available to the OpenTelemetry Collector [[2](https://github.com/prometheus/opentelemetry-collector-bridge)] so the two communities can avoid duplicated effort.

As a co-lead of the [Prometheus Special Interest Group](https://github.com/open-telemetry/community/blob/main/community-members.md#prometheus-interoperability) within OpenTelemetry, I have worked on the [Prometheus compatibility specification](https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/compatibility/prometheus_and_openmetrics.md) to ensure OTel SDKs follow Prometheus protocols for pull-based monitoring. I also maintain the Prometheus-related components in the OpenTelemetry Collector [[3](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusreceiver)] [[4](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/prometheusremotewritereceiver)] [[5](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/prometheusexporter)] [[6](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/prometheusremotewriteexporter)] and have helped grow the community around Prometheus interoperability.

I have spoken at international conferences about Prometheus several times and from several perspectives: as an end user [[7](https://youtu.be/bVICOulH5IY?si=45LTFLcjFwXOMLvo)], as a core developer of the project [[8](https://youtu.be/pRFUFJIAluk?si=nb_eGMN48xqfa8Zm)] [[9](https://youtu.be/t7mQHaHs_gw?si=0Z4ZPKYV7fj9Gt6u)] [[10](https://youtu.be/JFS0lSfHtMI?si=gcozcTAjgocjNlmz)], and as a community member interested in community growth [[11](https://youtu.be/rUvwnbhn2ZQ?si=qEkmEfnISYk0mSh4)].

With all that said, my proudest achievements are the mentorships I have led since graduating from my own mentorship. I bootstrapped the GSoC mentorships for the Prometheus Operator organization and have also mentored on behalf of the Prometheus organization. I have mentored mentees as well as mentors, helping introduce new long-term contributors, both junior and senior, to areas where Prometheus needs more contributors, such as technical writing, UX research, design, and software engineering. Several of them later joined observability vendors that build on top of Prometheus or end-user companies that use Prometheus internally.

Last but not least: ever since my first job in 2017, I have been known as "the Prometheus guy" in every role I have had.

## What I'll do

When we say "Prometheus", the first thing that comes to mind is often the server, but Prometheus' success has always depended on the ecosystem around it. A non-exhaustive list includes hundreds of Prometheus exporters, Alertmanager, installers such as Helm packages, Ansible playbooks, Chef cookbooks and Kubernetes operators, UX research, mixins, our official website, SDKs for several languages, push and pull protocols, social media, developer advocacy, and documentation for all of the above.

The surrounding ecosystem is just as important to Prometheus' success as the development of the server, but it is not as set up for success as it could be without clearer ownership. My goal as a Steering Committee member is to help scale our organization in ways that promote maintainership for the less visible but essential parts of the ecosystem. I want us to substantially grow the number of active Prometheus organization members and maintainers, aiming for at least 100 by 2027 and at least 200 by 2028.

To make that growth sustainable, I would like to work with the rest of the community to introduce subgroups that can own certain parts of the ecosystem and help them thrive with more autonomy (e.g. Security SIG, Exporter SIG, and SDK SIG). I also want to scale our mentorship programs by training current maintainers to become mentors and making sure mentorships are designed with clear paths for long-term involvement in the Prometheus community and related subgroups.

As the organization grows, we also need governance and access control that can scale with it. One initial project I would like to help drive is moving more of our governance from Markdown documents into "governance as code": clearly defining roles such as Steering Committee member, maintainer, and organization member in code, making GitHub organization administration auditable, and improving control over read, write, and admin permissions so they reflect people's current responsibilities.

Above all, I want to serve the community: to listen actively to our members' needs, help them drive their projects to completion, and remain open to the ideas and priorities that emerge during the term. My involvement with the Prometheus community is the highlight of my Engineering career, and I want more people to proudly say the same thing.

## Resources about me

- GitHub: [@ArthurSens](https://github.com/ArthurSens)
- Social / contact:
  - [LinkedIn](https://www.linkedin.com/in/arthursilvasens/)
  - [Twitter/X](https://x.com/ArthurSilvaSens)
