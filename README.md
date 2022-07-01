This project aims to crate a structure to handle an EMR - Electronic Medical Register.

In order to achieve that, the project is designed under DDD guidelines and TDD methodology (BTW I am learning both of them).

So far I have only design the domain layer of Family structure, that is the basis to build the relations among family members, which in turn will allow the mother or the father to handle, manage the EMR of their children.

There is a blueprint of the technologies that will be used in the project. For now, the project is agnostic to any infrastructure or technology. However, It looks like to be a monolith rather than multiple microservices.