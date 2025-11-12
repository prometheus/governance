# Prometheus Governance 2.0

This document describes the rules and governance of the project. It is meant to be followed by all the developers of the project and the Prometheus community. This governance is an open, living document, and will continue to evolve as the community and project change.

It outlines the roles and responsibilities of the Steering Committee, the decision-making process, and the election procedures for committee members. It also details the contributor ladder, including the roles of Contributor, Member, and Maintainer, and the requirements and privileges associated with each role.

The Prometheus Steering Committee is the governing body of the Prometheus project, providing decision-making and oversight pertaining to the project bylaws, sub-organizations[^1], and financial planning. The Steering Committee also defines the project values and structure.

## Values

The Prometheus developers and community are expected to follow the values defined in the [CNCF charter](https://www.cncf.io/about/charter/), including the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md). Furthermore, the Prometheus community strives for kindness, giving feedback effectively, and building a welcoming environment. The Prometheus developers default to deciding by consensus and only resort to conflict resolution by a majority vote if consensus cannot be reached.

## Roles

We have a contributor ladder that includes the roles of Contributor, Member, Maintainer and Steering Committee. The steering committee responsibilities and elections are outlined below. The rest of the roles are detailed in [ROLES.md](./ROLES.md).

## Charter

The following responsibilities and powers belong to the Steering Committee.

* Define, evolve, and promote the vision, values, and mission of the project.
* Define and evolve project governance structures and policies, including project roles and how collaborators become members, approvers, leads, and/or administrators. This includes policy for the creation and administration of working groups and committees.
* Steward, control access, delegate access, and establish processes regarding all Prometheus project resources and has the final say in the disposition of those resources.
* Coordinate with the CNCF regarding usage of the Prometheus brand and how that brand can be used in relation to other efforts or vendors.
* Receive and handle reports about code of conduct violations and maintain confidentiality.
* Act as the final escalation point and decide for any disputes, issues, clarifications, or escalations within the project scope.
* Have “owners” role for the `prometheus` and `prometheus-community` github organizations.
* Ultimate responsibility for all key project functions, unless delegated to a team. For example, vulnerability and security report handling, or github administration.

## Committee Meetings

The Steering Committee meets monthly, or as-needed.  
Meetings are held online, and are public by default (anyone can attend).

Given the private nature of some of these discussions (e.g. privacy, private emails to the committee, code of conduct violations, escalations, disputes between members, security reports, etc.) some meetings are held in private.

Public meeting recordings and notes are public and private meeting notes are private.

Questions and proposals for changes to governance or community procedures are posted as [issues in the governance repo](https://github.com/prometheus/governance/issues), and the Steering Committee invites your feedback there. See [Getting in touch](#getting-in-touch) for other options.

## Committee members

Seats on the Steering Committee are held by an individual, not by their employer.

The current membership of the committee is currently (listed alphabetically by first name):

| \\&nbsp;                                                         | Member           | Organization |  
| -------------------------------------------------------------- | ---------------- | ------------ |

## Decision process

The Steering Committee defaults to deciding by consensus.

Decisions requiring a vote include: issuing written policy, amending existing  
written policy, creating, removing, or modifying a working group, all spending,  
hiring, and contracting, official responses to publicly raised issues, or any  
other decisions that at least two of the members present decide require a vote.

Votes happen by simple majority.

## Getting in touch

There are two ways to raise issues to the steering committee for decision:

1. The Steering Committee can be reached privately by emailing [prometheus-steering@googlegroups.com](mailto:prometheus-steering@googlegroups.com). This is a private discussion list to which all members of the committee have access.
2. The Steering Committee can be reached publicly by opening an issue on [the governance repository](https://github.com/prometheus/governance/issues) and indicating that you would like attention from the steering committee.

## Composition

The steering committee has 7 seats. These seats are open to any project member. See [Candidate Eligibility](#candidate-eligibility) for a definition.

Steering Committee members serve for 2 year terms, staggered to preserve continuity. Every year, either 4 or 3 contributor seats are elected.

## Election Procedure

### Timeline

Steering Committee elections are held annually. Six weeks or more before the election, the Steering Committee will appoint [Election Officer(s)](#election-officers).  
Four weeks or more before the election, the Election Officer(s) will issue a call for nominations, publish the list of voters. 

Seven days before the election, the call for nominations will be closed. The election will be open for voting for no less than two weeks and no more than four. The results of the election will be announced within one week of closing the election. New Steering Committee members will take office is May or June of each year.

### Election Officers

Six weeks or more before the election, the Steering Committee will appoint three Election Officer(s) to administer the election. Election Officers will be Organization Members in good standing who are eligible to vote, are not running for Steering Committee in that election, who are not currently part of the Steering Committee and can make a public promise of impartiality. 

They will be responsible for:

* Making all announcements associated with the election
* Preparing and distributing electronic ballots
* Assisting candidates in preparing and sharing statements
* Tallying voting results according to the rules in this governance

### Eligibility to Vote

Anyone who is an Organization Member as defined in [ROLES.md](#roles.md).

The electoral roll of all eligible voters will be captured at [elections](./elections/) and the voters’ guide will be captured at [voters guide](./elections/).

### Candidate Eligibility

Community members must be eligible to vote to stand for election. Candidates should self-nominate. There are no term limits for Steering Committee members.

### Voting Procedure

Elections will be held using a ranked choice voting system which supports voter privacy.  Specific voter tooling will be selected by the Election Officers before the election starts, with this selection to be approved by the Steering Committee. The top vote-getters will be elected to the open seats, with the exceptions for company representation discussed below.

### Limitations on Company Representation

No more than 2 seats may be held by employees of the same organization (or  
conglomerate, in the case of companies owning each other). If an election results in greater than 2 employees of the same organization being selected, the lowest vote getters from any particular employer will be removed until representation on the committee is down to 2.

If employers change because of job changes, acquisitions, or other events, in a way that would yield too many seats being held by employees of the same organization, and the next regular election is more than 10 weeks away, sufficient members of the committee must resign until only 2 employees of the same employer are left. If it is impossible to find sufficient members to resign, all employees of that organization will be removed and new special elections held. 

In the event of a question of company membership (for example evaluating independence of corporate subsidiaries) a majority of all non-involved Steering Committee members will decide.

## Vacancies

In the event of a resignation or other loss of an elected Steering Committee member, a new vote for that seat is started by Election Officers.

A special election for that position will be held as soon as possible, unless the regular Steering Committee election is less than 10 weeks away. Eligible voters from the most recent election will vote in the special election. Eligibility will not be redetermined at the time of the special election. Any replacement Steering Committee member will serve out the remainder of the term of the person they are replacing, regardless of the length of that remainder.

## Changes to the Governance

Changes to this Governance may be proposed via a Pull Request on the Governance itself.  
Amendments are accepted with 2/3rd majority consent of the committee as per the [Decision Process](#decision-process) outlined above.

Proposals and amendments to the charter are available for at least a period of one week for comments and questions before a vote will occur.

## Authority, Facilitation, and Decision Making

Ideally, most decisions will be made at the lowest possible level within the project: within individual working groups. When this is not possible, the Steering Committee can help facilitate a conversation to work through the contended issue. When facilitation by the Steering Committee does not resolve the contention, the Steering Committee may have to decide.

For example, Technical governance is expected to be performed by the Maintainers of the various project areas and subprojects, as defined in the appropriate OWNERS / MAINTAINERS files.

Note that if the committee is called to resolve contended decisions regularly, it is a symptom of a larger problem in the community that will need to be addressed.

## Inactivity

It is important for members to be and stay active to set an example and show commitment to the project. Inactivity is harmful to the project as it may lead to unexpected delays, contributor attrition, and a loss of trust in the project.

* Inactivity is measured by
  * Periods of no contributions for longer than 6 months
  * Not having 20 contributions in the preceding 12-month window for more than 6 months
* Consequences of being inactive include
  * Involuntary removal or demotion
  * Being asked to move to Emeritus status
  * Not counting against quorum in votes

## Involuntary Removal or Demotion

Involuntary removal/demotion of a contributor happens when responsibilities and requirements aren't being met. This may include repeated patterns of inactivity, extended periods of inactivity, a period of failing to meet the requirements of your role, and/or a violation of the Code of Conduct. This process is important because it protects the community and its deliverables while also opening up opportunities for new contributors to step in.

Inactive members can be automatically removed.

Involuntary removal or demotion of Maintainers or Steering Committee members is handled through a vote by a majority of the current Steering Committee. If this pertains to a Steering Committee member, that member is not eligible to discuss or vote on the matter.

## Stepping Down/Emeritus Process

If and when contributors' commitment levels change, contributors can consider stepping down (moving down the contributor ladder) vs moving to emeritus status (completely stepping away from the project).

Contact the Maintainers about changing to Emeritus status, or reducing your contributor level.

## Contact

For inquiries, please reach out to:  
* [the governance repository](https://github.com/prometheus/governance/issues): Governance Repo Issues  
* [prometheus-steering@googlegroups.com](mailto:prometheus-steering@googlegroups.com): Steering Committee Email

## Migrating from the old governance

All current Prometheus team members become members as per new governance. 

The individual projects already define who is a maintainer for which specific project in their corresponding MAINTAINERS.md. People listed in these files are therefore going to be maintainers. 

The existing Prometheus team is going to hold votes, as described in the new governance, to elect a Bootstrap steering committee of 7 members, where the highest voted 3 members will hold a 2 year and the next 4 members will hold a 1-year term.  


[^1]:	[github.com/prometheus](http://github.com/prometheus), [github.com/prometheus-community](http://github.com/prometheus-community), [github.com/promcon](http://github.com/promcon), [github.com/openmetrics](http://github.com/openmetrics)