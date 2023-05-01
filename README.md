Challenge 1 --> Basic Screaming arch + Hexagonal

Challenge 2 --> Unit test + error handling

Challenge 3 --> Gin + e2e test + a bit of DDD (string value object)

Challenge 4 --> Added sqlite + a bit of DDD (collections) + search as bounded context

Opinionated implementation of Screaming + Hexagonal Architecture


## Project Structure (per challenge)

```
src
|-- cmd: Entry points of the bounded context
|   |-- Context(adView): this allows an entry point per bounded context
|         enabling a monorepo of multiple apps
│
|-- internal: Diferent Bounded Context in the same repo, owned by different teams
|   |-- Context (adView, search)
|       |-- Module: Usually, an aggregate, but also Shared
|           |-- Application: Use cases
|           |-- Domain: Aggregate root, Value objects, Interfaces
|           |-- Infrastructure: Outbound dependencies (external services, db...)
|
tests
|-- cmd
|   |-- Context
|       |-- Module: Black box acceptance tests
```


## Responsibilities per layer
### Infrastructure
- Controller:
    1. Receive request
    2. Parse params
    3. Call query/command bus or the use case in case of not having them
    4. Map result/domain errors to response

### Application
- Query Handle
    1. Subscribe and receive query
    2. Parse arguments and call use-case
    3. Apply projection in case of necessity

- Command Handle
    1. Subscribe and receive command
    2. Parse arguments and call use-case
    3. Return error if exist

- Use case
    1. Transactional unit
    2. Can use other use cases/query bus
    3. Can use infrastructure through dependency inversion
    4. Orchestrator, try not to have business logic here
    5. Persist
    6. Emit domain events

### Domain
- Aggregate root
    1. Composed by value objects and other entities
    2. Creates and record domain events (but dont emit)
    3. Single transactional unit
- Collection
    1. Compose by the list of an entity
- Value object
    1. Dont emit events
    2. Usually owns the validation logic
- Read Projection
    1. Variants of an aggregate to expose in a use-case

## Communication Patterns
- Controller uses query/command bus to call use cases
- Use cases in the same module can be called directly
- Try to use the query bus for use cases in different modules
- If information from other bounded context is necessary, a projection
  listening events of that bounded context should be created

## Testing Strategy

C/Q = Command/Query

Data Flow Diagram

```
Controller --> C/Q Bus --> C/Q Handler --> Use Case --> Repositories
                                                    --> External Services

                           ------------------------     -----------------
                                   Unit Test             Integration Test
-------------------------------------------------------------------------
                              Acceptance Tests

```

Dependency Diagram

```
Controller <-- C/Q Bus --> C/Q Handler --> Use Case --> Repositories
                                                    --> External Services

```

## Todo:
- Implement event bus
- Create **listingAd** domain entity in _search_ based on **Ad** created on adView context
- Implement query/command bus
- Implement a read projection
