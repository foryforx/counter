# Sequence Counter

## Description

Service that generates unique sequence number for every request.

## Requirements

- golang > 1.19
- docker
- shell
- ab tool (optional)

## How to run

```
  make run
```

## How to test

- Run in 2 different terminals

```
  make runner1
```

```
  make runner2
```

## How to test with ab tool

```
  make ab
```

## coverage

```
  make coverage
```

## Future improvements

[ ]Add JWT auth for the api exposed
[ ] Persist the counter value with database to retain the seq number counter
[ ] Add docker compose with file watcher to make development much more easier
[ ] Add pipelines fo CI/CD

## Other queries

Consider a hypothetical scenario where a B2B company is currently deploying a dockerised SaaS solution for many of its Clients. Its a single tenant solution that consists of a UI, Backend Service and NoSQL DB. This solution needs to be deployed, branded and its DNS configured for each Client, similar to a solution such as vercel.com, but instead deploying a bespoke SaaS solution. The company intends to expand its solution to 100+ businesses; each business can have between 50,000 to 10,000,000 users.

### a. As an engineer, what steps would you take to address this challenge? Please walk me through your approach.

KAL>>

1. As-Is:

- Identify the setup, configuration, coverage and scalability of the current solution.
- Run performance tests to identify the current performance of the solution. Identify baseline for current setup.

2. Infrastrcuture:

- Move the infrastructure creation using terraform with each tenants configuration.
- Use cloud CDN for caching the static content.
- Use cloud load balancer for load balancing the traffic.

3. CI/CD:

- Create a CI/CD pipeline to deploy the solution for each tenant.
- Use K8s for deployment and scaling. Use k8s namespace to split based on tenants
- Make sure the CD takes care of the DNS configuration for each tenant. In google cloud, we can use cloud DNS for this.
- Implement pod scale based on CPU and mem usage. for eg: 60% cpu usage, scale up.
  (see possibility of using Kong API gateway for routing the traffic to the right downstream service + other auth features)
- If possible, deploy the UI in S3 and use cloudfront for CDN.
- Execute performance script using K6 and identify if any bottlenecks. Identify and store baseline.

4. Storage

- Use cloud BigTable for storage. (Row key prefixes provide a scalable solution for a "multi-tenancy" use case, a scenario in which you store similar data, using the same data model, on behalf of multiple clients. Using one table for all tenants is the most efficient way to store and access multi-tenant data.)
- Configure snapshots/backup at regular intervals.

5. Monitoring:

- Use Datadog or NewRelic to monitor the performance of the solution. Make sure we instrument the code to capture the metrics.
- Use uptime for api monitoring and site availability.
- Send Node, pod, container, database metrics to datadog to monitor the performance.

6. Security:

- Use JWT auth for the api exposed. Use cloud IAM for access control.

7. Testing:

- Use unit tests, integration tests and e2e tests to make sure the solution is working as expected. Keep everything automated and make the CI/CD pipeline fail if any of the tests fail. Otherwise, it can deploy till UAT env.

8. Logging:

- Forward all logs to datadog to setup alerts, dashboards and to monitor the logs.

9. Alerting:

- Setup alerts for the metrics captured. For eg: if the cpu usage is more than 60% for 5 mins, send an alert to the team.
- Setup Level 1 alerts where it needs immediate attention and its P0
- Setup Level 2 alerts where it needs attention but not immediate and its P1
- Setup Level 3 alerts where it needs attention but not immediate and its P2
- Setup Level 4 alerts where it needs attention but not immediate and its P3

10. Incident management:

- Setup incident management process to handle the alerts.
- Setup incident management process to handle the incidents.
- Setup incident management process to handle the RCA.

11. Documentation:

- Document the solution, architecture, setup, configuration, monitoring, alerting, incident management, RCA, etc.

12. Support:

- Setup support process to handle the issues raised by the tenants.
- Setup support process to handle the issues raised by the users.
- Setup support process to handle the issues raised by the internal team.

13. Onboarding:

- Setup onboarding process to onboard new tenants.
- Setup onboarding process to onboard new users.
- Setup onboarding process to onboard new team members.

14. Training:

- Setup training process to train the tenants.
- Setup training process to train the users.
- Setup training process to train the team members.

### b. Furthermore, please provide a high-level component diagram using cloud-based design principles, utilizing components from GCP and open source tools that would support your proposed solution.

KAL>> diagram in root folder
[diagram](https://raw.githubusercontent.com/foryforx/counter/main/component_dig.jpg?raw=true)

### c. Finally, how would you approach continuous integration and continuous deployment (CI/CD) for this SaaS-based application? What advice would you offer the company to ensure successful implementation and maintenance of the application?

KAL>>
I prefer github for CI/CD.

- Basic idea of successful delivery/deployment/maintenance starts from small code changes which are less risky, well tested in automated way and writen in scalable way. We need ever improving process based on the learnings from the past.
- We can follow trunk based development where we can merge the code to master branch and deploy to all the environments.
  - We can use feature flags to enable/disable the features. But make sure feature flags dont live for more than X months.
- CI/CD pipeline should be automated and should be able to deploy to all the environments.
- Keep everything in descriptive versioned code. For eg: terraform code, k8s code, dockerfile, etc.
- Automate everything and anything.
- Please note that CI/CD is done on the service/repository level. So, we need to have a CI/CD pipeline for each service/repository.

Local development:

- We can use docker compose to setup the local environment. This will help in making the local development much more easier.
- Use file checker to auto restart docker continers
- Use file checker to auto run tests on file change
- Use file checker to auto run lint on file change
- Use file checker to auto run security check on file change
- Have methods to validate that test coverage is not below the baseline.

Merge request:

- All commits should be sqashed and tagged with the ticket number.
- All commits should have a meaningful commit message.
- All commits should have a meaningful PR message.
- All commits should have a meaningful branch name.
- All merge request should be reviewed by atleast 2 team members. [Please define code review guidelines and keep a living document for it]

CI:

- CI will run on all branches to validate every commit before merging to master.
- CI will do below mentioned on every commit in every branch
  - unit tests, integration tests
  - Lint check
  - Security check
  - Build the docker image
  - Push the docker image to Image repository with commit tag
- CI will also run below mentioned in scheduled fashion(like daily)
  - K6 performance test for all API with comparison to baseline and send the report to the team.

CD: [Keep it with manual trigger]

- We can use k8s for deployment and scaling. Use k8s namespace to split based on tenants.
- CD will appear over CI, only in case of master branch or branch with prefix hotfix/
- CD will do below mentioned
  - Apply migrations, if any
  - Pull the commit tagged docker image from image repository
  - Deploy the docker image to k8s cluster
  - Run sanity/synthetic tests to validate the deployment
  - Autotag the branch
- If possible move to canary deployment to validate your changes before deploying to all the environments.

IAC:

- We can use terraform to create the infrastructure.
  - k8s cluster.
  - k8s namespace.
  - Bigtable
  - Cloud CDN
  - Cloud Load balancer
  - Cloud DNS
  - Cloud IAM
  - Cloud Storage

Monitoring:

- Use Datadog or NewRelic to monitor the performance of the solution. Make sure we instrument the code to capture the metrics.
- With monitors triggering communication to members, we can proactively handle the issues.
- Use uptime for api monitoring and site availability.
- Send Node, pod, container, database metrics to datadog to monitor the performance.
- Send logs to datadog to setup alerts, dashboards and to monitor the logs.
