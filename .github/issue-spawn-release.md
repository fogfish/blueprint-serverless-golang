# Spawn release to live environment

The delivery of microservice to production is automated using git tags (GitHub Releases). A provisioning of a new tag caused a new immutable deployment of microservice(s) to the live environment.

The risk of occasional outage is always there due to regression in software quality. The immutable deployment is the solution for availability and fault tolerance. CI/CD never changes anything at running systems. Immutable deployment making changes by rebuilding microservices, it deploys a new copy in parallel stack. The rollback of defective software happens in a matter of seconds. Another advantage is the ability to split traffic between parallel deployments for quality assurance.

