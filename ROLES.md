# Prometheus Community Roles

This document describes the Prometheus contributor ladder. It defines the roles within the project, together with the responsibilities, requirements, and privileges associated with each role. In doing so, it also implicitly describes how contributors grow their participation in the project over time.

Being at a “higher” level on the ladder includes the “lower” levels; e.g., if you are a Maintainer, you are also considered a Member. If you are a Member, you are also considered a Contributor and have the privileges and responsibilities of all the lower levels.

The contributor ladder covers roles that are assessed based on a contributor's demonstrated participation and the documented expectations for that role. It does not cover elected governance roles such as the Steering Committee; those are documented separately in [GOVERNANCE.md](./GOVERNANCE.md), since they are filled through the project's election process rather than by advancement through the ladder.

### Contributor

Description: A Contributor contributes directly to the project and adds value to it. Contributions do not need to be code. People at the Contributor level may be new contributors or contribute only occasionally.

* Qualifying Contributions
  * Report, triage, work on, or resolve issues via comments
  * Submit, review, or test PRs
  * Contribute to the documentation
  * Participate in dev summits or workgroup meetings
  * Answer questions from other community members
  * Help to test new releases
  * Run or help run events
  * Promote the project in public
  * Help run the project infrastructure
* Responsibilities
  * Follow the CNCF CoC
  * Follow the project contributing guide
* Privileges
  * Invitations to contributor events

### Member

Description: A Member is an established contributor who regularly participates in the project. Members have privileges in both project repositories and elections, and as such are expected to act in the interests of the whole project.

A Member must meet the responsibilities and has the requirements of a Contributor, plus:

* Responsibilities include
  * Continuing to contribute regularly.
    * This can be demonstrated by having at least 20 GitHub contributions a year in [DevStats](https://devstats.cncf.io/).
    * Or non-code contributions like attending and contributing to dev summit and workgroup meetings, organizing and running activities within the Prometheus community, etc.

* Requirements
  * Must meet the requirements of a Contributor and have at least 20 qualifying contributions in the past year
  * Must have at least 2 Maintainers as sponsors. At least one sponsor must have a different employer from the applying Contributor
  * Must enable 2FA on their GitHub account

* Privileges
  * Entitled to vote in the Steering Committee Election
  * Entitled to stand in the Steering Committee Election
  * Can recommend other contributors to become Org Members
  * May be assigned Issues and Reviews
  * May give commands to CI/CD automation
  * Can be added to repository teams
  * Added as a member of the GitHub org

The process for a Contributor to become a Member is as follows:

1. Open a GitHub Issue against the [prometheus/governance](https://github.com/prometheus/governance)
  1. Ensure your sponsors are @mentioned on the issue
  2. Assign label `membership-application`
  3. Complete every item on the checklist
  4. Make sure that the list of contributions included is representative of your work on the project.
2. Sponsors reply via a comment on the GitHub Issue with explicit confirmation
3. Once Sponsors have responded positively, the request will be reviewed by the Steering Committee or its delegates to make sure certain contributions are valid and non-trivial.
4. If the threshold is met, the Contributor is added to [https://github.com/orgs/prometheus/people](https://github.com/orgs/prometheus/people) and [https://github.com/orgs/prometheus-community/people](https://github.com/orgs/prometheus-community/people) and becomes a Member.

### Maintainer

Description: Maintainers are well-established contributors who review and approve code contributions. Maintainers are responsible for the long-term health of the project, including reviewing PRs, stability, adoption, security, etc.

Maintainer status is scoped to a repository or a part of the codebase.

A Maintainer must meet the responsibilities and requirements of an Organization Member, plus:

* Responsibilities include
  * Demonstrate sound technical judgment
  * Responsible for project quality control via code reviews
  * Responsible for the technical quality control of the documentation related to their project
  * Expected to be responsive to review requests
  * Mentoring new contributors
  * Determining strategy and policy for the project
  * Participating in and leading community meetings
  * May merge PRs affecting the code they are Maintainers for
* Requirements
  * Consistently contributing to the codebase
  * Reviewer for or author of substantial PRs to the codebase, with the definition of substantial subject to the maintainer’s discretion (e.g., refactors/adds new functionality rather than one-line pulls).
  * Can exercise judgment for the good of the project, independent of their employer, friends, or team
* Additional privileges
  * Can manage CI and package registry secrets
  * Represent the project in public as a Maintainer
  * Take part in the Maintainer decision-making

Process of becoming a maintainer:

1. Any current Maintainer may nominate a current Member to become a new Maintainer by opening a PR, adding them to the `MAINTAINERS.md` file.
  * If the codebase or project has no current maintainers, then the Steering Committee can also nominate new maintainers.
1. The nominee will add a comment to the PR testifying that they agree to all requirements of becoming a Maintainer.
2. None of the other maintainers of the relevant repository or part of the code base must object.

## Stepping Down/Emeritus Process

If and when contributors’ commitment levels change, contributors can consider stepping down (moving down the contributor ladder) or completely stepping away from the project.

Contact the Maintainers about changing to Emeritus status and reducing your contributor level.
