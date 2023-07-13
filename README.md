# Genesis Software Engineering School 3.0

## Doc

[openapi.yaml](doc%2Fopenapi.yaml)

## Introduction

The application is divided into several key modules as detailed below:

- **cmd**: Contains the application's entry point.
- **data**: Contains file store, or raw data.
- **docs**: Contains documentation files.
- **internal**: Contains the core application logic divided into `rate`, `subscription`, and `transport` packages.
- **scripts**: Contains auxiliary scripts for various tasks.
- **sys**: Contains system-level packages like `env`, `filestore`, and `logger`.

Each module is responsible for a specific function within the application, allowing for clear separation of concerns and
making the codebase easy to manage and navigate.

## Installation and Setup

```shell
make install
```

```shell
make run
```

```shell
make docker-build
 ``` 

```shell
make docker-run
 ```  

## Module Tree
--- TODO: Update
```
📦gentest
 ┣ 📂cmd
 ┃ ┗ 📜main.go
 ┣ 📂data
 ┣ 📂docs
 ┣ 📂internal
 ┃ ┣ 📂rate
 ┃ ┃ ┣ 📜getter_mock_test.go
 ┃ ┃ ┣ 📜handler.go
 ┃ ┃ ┣ 📜handler_test.go
 ┃ ┃ ┣ 📜rate.go
 ┃ ┃ ┗ 📜rate_test.go
 ┃ ┣ 📂subscription
 ┃ ┃ ┣ 📜handler.go
 ┃ ┃ ┣ 📜handler_test.go
 ┃ ┃ ┣ 📜repository.go
 ┃ ┃ ┣ 📜subscriber_mock_test.go
 ┃ ┃ ┗ 📜subscription.go
 ┃ ┗ 📂transport
 ┃   ┣ 📜http.go
 ┃   ┗ 📜middleware.go
 ┣ 📂scripts
 ┣ 📂sys
 ┃ ┣ 📂env
 ┃ ┃ ┣ 📜env.go
 ┃ ┃ ┗ 📜env_test.go
 ┃ ┣ 📂filestore
 ┃ ┃ ┣ 📜filestore.go
 ┃ ┃ ┗ 📜filestore_test.go
 ┃ ┗ 📂logger
 ┃   ┗ 📜logger.go
 ┣ 📜.env
 ┣ 📜.gitignore
 ┣ 📜.golangci.yml
 ┣ 📜Dockerfile
 ┣ 📜go.mod
 ┣ 📜go.sum
 ┣ 📜Makefile
 ┗ 📜README.md
```

## Architecture

```mermaid
graph LR 
    main((main)) ==> App
    main ==> Env
    main & EventBus & Web & App ==> Logger>Logger]
    App & Handlers & SubscriptionAdapters & RateAdapters --> |uses| Web
    App -->|binds| RateService & SubscriptionService & Infrastructure & RateAdapters & SubscriptionAdapters
    Domain ==> Handlers
    
    RateAdapters -.->|impl| ExchangeRateProvider
    SubscriptionAdapters -.->|impl| EmailSender
    SubscriptionService -.->|impl| SubscriptionServiceInterface
    RateService -.->|impl| RateServiceInterface
    
    subgraph Transport
        App((APP)) -->|binds| RateHandlers[Rate Handlers]
        App -->|binds| SubscriptionHandlers[Subscription Handlers]
        subgraph Handlers
            RateHandlers -->|uses| RateServiceInterface{{RateService}}
            SubscriptionHandlers -->|uses| SubscriptionServiceInterface{{SubscriptionService}}
        end
    end
    subgraph Domain
        subgraph Rate
            subgraph RateCore 
                ExchangeRate(ExchangeRate)
            end
            RateService((SERVICE)) --> ExchangeRate
            RateService -->|uses| RateEvent
            RateService -->|uses| ExchangeRateProvider{{ExchangeRateProvider}}
        end
        subgraph Subscription
            subgraph SubscriptionCore
                Subscriber{Subscriber}
                Message(Message)
                Topic(Topic)
            end
            SubscriptionService((SERVICE)) --> SubscriptionCore
            SubscriptionService -->|uses| Repository{{SubscriberRepository}}
            SubscriptionService -->|uses| EmailSender{{EmailSender}}
            SubscriptionService -->|uses| SubscriptionEvent

        end
    end
    
    subgraph RateAdapters
        A
        B
        C
        D
    end
    
    subgraph SubscriptionAdapters
        EmailClient
    end
   
    subgraph Infrastructure
        Repository -.->|impl| FileStore[(File Store)]
        App((Controller)) -->|uses| Web[[Web]]
        RateHandlers[/Rate Handler/] -->|uses| Web[[Web]]
        SubscriptionHandlers[/Subscription Handler/] -->| uses| Web[[Web]]
        subgraph Event
            SubscriptionEvent{Event} -->|uses| EventBus((Event Bus))
            RateEvent{Event} -->|uses| EventBus((Event Bus))
        end
        subgraph Web
            Middleware
            Tooling
        end
        subgraph Env
        end
    end


    Client[Client] -->|interacts| App(((APP)))

    classDef main fill:#FF2800;
    classDef ultra fill:#FC8000;
    classDef orange fill:#FC8000;
    classDef pink fill:#FF9CD6;
    classDef blue fill:#0000FF;
    classDef yellow fill:#E8BA36;
    classDef func fill:#4EADE5;
    classDef carrot fill:#E35C5C;
    classDef face fill:#700FC1;
    classDef cherry fill:#960039;
    classDef green fill:#39FF14 
    classDef rust fill:#B04721; 
    classDef blues fill:#663DFE; 
    classDef pale fill:#3B898A; 
    classDef pack fill:#C418A2; 
    
    class Web pack;
    class Subscriber,ExchangeRate face;
    class SubscriptionService,RateService cherry;
    class ExchangeRateProvider,Repository,EmailSender,SubscriptionServiceInterface,RateServiceInterface,RateServiceInterface,ExchangeRateServiceInterface func;
    class FileStore pale;
    class EventBus,RateEvent,SubscriptionEvent green;
    class App pack;
    class Message,Topic blues;
    class SubscriptionHandlers,RateHandlers rust;
    class Logger blue;
    class Env yellow;
    class Client pink; 
   
```
