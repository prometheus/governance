# Prometheus Governance

This document describes the rules and governance of the Prometheus project. It is meant to be followed by all the developers of the project and the Prometheus community. This governance is an open, living document, and will continue to evolve as the community and project change.

It outlines the Steering Committee's roles and responsibilities, the decision-making process, and the election procedures for committee members. It also details the contributor ladder, including the roles of Contributor, Member, and Maintainer, along with the requirements and privileges associated with each.

The Prometheus Steering Committee is the governing body of the Prometheus project, providing decision-making and oversight pertaining to the project by-laws, GitHub organizations[^1], and financial planning. The Steering Committee also defines the project values and structure.

## Values

The Prometheus developers and community are expected to follow the values defined in the [CNCF charter](https://www.cncf.io/about/charter/), including the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md). Furthermore, the Prometheus community strives for kindness, effective feedback, and a welcoming environment. The Prometheus developers default to decision by consensus and only fall back to voting if consensus cannot be reached.

## Roles

We have a contributor ladder that includes the roles of Contributor, Member, Maintainer, and Steering Committee. The Steering Committee responsibilities and elections are outlined below. The remaining roles are detailed in [ROLES.md](./ROLES.md).

## Charter

The Steering Committee has the following responsibilities and powers.

* Define, evolve, and promote the vision, values, and mission of the project.
* Define and evolve project governance structures and policies, including project roles and how collaborators adopt these roles. This includes policy for the creation and administration of working groups and committees.
* Steward, control access, delegate access, and establish processes regarding all Prometheus project resources, having the final say in the disposition of those resources.
* Coordinate with the LF and the CNCF regarding usage of the Prometheus trademark and how that trademark can be used in relation to and with other efforts, projects, or vendors.
__TODO__: Discuss what happens if SC members is violating CoC.
* Receive and handle reports about Code of Conduct violations while maintaining confidentiality.
* Act as the final escalation point and decide for any disputes, issues, clarifications, or escalations within the project scope.
* Have "owners" role for the project's GitHub organizations[^1].
* Ultimate responsibility for all key project functions, unless delegated to a team. For example, vulnerability and security report handling, or GitHub administration.

## Committee Meetings

The Steering Committee meets monthly, or as-needed.  
Meetings are held online and are public by default (anyone can attend).

Given the private nature of some of these discussions (e.g., privacy, private emails to the committee, code of conduct violations, escalations, disputes between members, security reports, etc.), some meetings are held in private.

Questions and proposals for changes to governance or community procedures are posted as [issues in the governance repo](https://github.com/prometheus/governance/issues), and the Steering Committee invites your feedback there. See [Getting in touch](#getting-in-touch) for other options.

## Committee members

Seats on the Steering Committee are held by individuals, not by or through their respective employers.

The current membership of the Steering Committee, listed alphabetically by first name:

| \\&nbsp;                                                         | Member           | Organization |  
| -------------------------------------------------------------- | ---------------- | ------------ |

## Decision process

The Steering Committee defaults to deciding by consensus.

Decisions requiring a vote include:
* Issuing written policy
* Amending existing written policy
* Creating, removing, or modifying a working group
* All spending, hiring, and contracting
* Official responses to publicly raised issues
* Any other decisions that two or more members present agree should be put to a vote

To pass a decision, it needs approving votes from more than half of the current Steering Committee members.

__TODO__: Explain how to handle multi-proposal votes.

## Getting in touch

There are two ways to raise issues to the Steering Committee for decision:

__TODO__: Figure out if we want to keep the archive or a mailing-alias.

1. The Steering Committee can be reached privately by emailing [prometheus-steering@googlegroups.com](mailto:prometheus-steering@googlegroups.com). This group is a private discussion list to which all committee members have access.
2. The Steering Committee can be reached publicly by opening an issue on [the governance repository](https://github.com/prometheus/governance/issues) and indicating that you would like attention from the Steering Committee.

## Composition

The Steering Committee has 7 seats. These seats are open to any project member. See [Candidate Eligibility](#candidate-eligibility) for a definition.

Steering Committee members serve for 2-year terms, staggered to preserve continuity. Not accounting for backfill, this means that either 3 or 4 Steering Committee members are elected per year.

## Election Procedure

### Timeline

Steering Committee elections are held annually, aiming to seat new members in __TODO__. Six weeks or more before the election, the Steering Committee will appoint [Election Officer(s)](#election-officers).
Four weeks or more before the election, the Election Officer(s) will issue a call for nominations and publish the list of eligible voters.

Seven days before the election, the call for nominations will be closed. The election will be open for voting for no less than two weeks and no more than four weeks. The election results will be announced within one week of the election closing.

### Election Officers

Six weeks or more before the election, the Steering Committee will appoint three Election Officers to administer the election. Election Officers will be Organization Members in good standing who are eligible to vote, are not running for the Steering Committee in that election, are not currently part of the Steering Committee, and can make a public promise of impartiality. 

They will be responsible for:

* Making all announcements associated with the election
* Preparing and distributing electronic ballots
* Assisting candidates in preparing and sharing statements
* Tallying voting results according to the rules in this governance

### Eligibility to Vote

Anyone who is an Organization Member as defined in [ROLES.md](./ROLES.md) may vote.

The electoral roll of all eligible voters will be captured at [elections](./elections/) and the voters' guide will be captured at [voters guide](./elections/).

### Candidate Eligibility

Organization members must be eligible to vote to stand for election. Candidates should self-nominate. There are no term limits for Steering Committee members.

### Voting Procedure

Elections will be held using a ranked choice voting system which supports voter privacy. Specific voter tooling will be selected by the Election Officers before the election starts, with this selection to be approved by the Steering Committee. The top vote-getters will be elected to the open seats, with the exceptions for company representation discussed below.

### Limitations on Company Representation

No more than 2 seats may be held by employees of the same organization (or  
conglomerate, in the case of companies owning each other). If an election results in more than 2 employees of the same organization being selected, the lowest vote-getters from any particular employer will be removed until representation on the committee is down to 2.

If affiliations change due to job changes, acquisitions, or other events and would result in more than 2 affiliated individuals serving on the Steering Committee, and the next regular election is more than 10 weeks away, a sufficient number of Steering Committee members must resign. If this does not occur within two weeks, all affiliated individuals from that set are removed and backfilled through normal backfill mechanisms.

In the event of a question of company membership (for example, evaluating the independence of corporate subsidiaries), a majority of all non-involved Steering Committee members will decide.

## Backfilling

In the event of a resignation or other loss of an elected Steering Committee member, Election Officers begin a new vote for that seat.

A special election for that position will be held as soon as possible, unless the regular Steering Committee election is less than 10 weeks away. Eligible voters from the most recent election will vote in the special election. Eligibility will not be redetermined at the time of the special election. Any replacement Steering Committee member will serve out the remainder of the term of the person they are replacing, regardless of the length of that remainder.

## Changes to the Governance

Changes to this governance may be proposed via a Pull Request on the governance itself.

Amendments are accepted if 2/3 of all current committee members cast an approving vote, as per the [Decision Process](#decision-process) outlined above.

Proposals and amendments to the charter are available for at least 2 weeks for comments and questions before a vote occurs.

## Authority, Facilitation, and Decision-Making

Ideally, most decisions will be made at the lowest possible level within the project: within individual working groups. When this is not possible, the Steering Committee can help facilitate a conversation to work through the contested issue. When the Steering Committee's facilitation does not resolve the contention, the Steering Committee may have to decide.

For example, technical governance is expected to be performed by the Maintainers of the various project areas and subprojects, as defined in the appropriate OWNERS / MAINTAINERS files.

Note that if the Steering Committee is regularly called to resolve contested decisions, it is a symptom of a larger problem in the community that will need to be addressed.

## Inactivity

Members need to be and stay active to set an example and show commitment to the project. Inactivity is harmful to the project, as it may lead to unexpected delays, contributor attrition, and a loss of trust.

* Inactivity is measured by
  * Periods of no contributions for longer than 6 months
  * Not having 20 contributions in the preceding 12-month window for more than 6 months

* Consequences of being inactive include
  * Involuntary removal or demotion
  * Being asked to move to Emeritus status

For maintainers of mature or stable projects that require little to no maintenance, the Steering Committee can grant exceptions.

## Involuntary Removal or Demotion

Involuntary removal/demotion of a contributor happens when responsibilities and requirements aren't being met. This may include repeated patterns of inactivity, extended periods of inactivity, a period of failing to meet the requirements of the assigned role, and/or a violation of the Code of Conduct. This process is important because it protects the community and its deliverables while also opening up opportunities for new contributors to step in.

Involuntary removal or demotion of Maintainers or Steering Committee members is handled by a majority vote of the current Steering Committee. If this pertains to a Steering Committee member, that member is not eligible to discuss or vote on the matter.

## Contact

For inquiries, please reach out to:  
* [the governance repository](https://github.com/prometheus/governance/issues): Governance Repo Issues  
* [prometheus-steering@googlegroups.com](mailto:prometheus-steering@googlegroups.com): Steering Committee Email

## Migrating from the old governance

All current Prometheus team members become members as per the new governance. 

The individual projects already define who a maintainer is for which specific project in their corresponding MAINTAINERS.md. People listed in these files will therefore be maintainers. This includes all the GitHub organizations[^1].

The initial list of eligible voters might include members and maintainers who are inactive according to [our requirements](#inactivity). This is OK, and they can still participate in the vote for the Bootstrap Steering Committee.

The existing Prometheus team will hold votes, as described in the new governance, to elect a Bootstrap Steering Committee of 7 members, with the 3 highest-voted members serving 2-year terms and the next 4 serving 1-year terms.  


[^1]:	[github.com/prometheus](http://github.com/prometheus), [github.com/prometheus-community](http://github.com/prometheus-community), [github.com/promcon](http://github.com/promcon)

