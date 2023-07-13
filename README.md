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
📦xrate
 ┣ 📂.github
 ┃ ┗ 📂workflows
 ┃   ┣ 📜go.yml
 ┃   ┗ 📜golangci.yml
 ┣ 📂api
 ┃ ┣ 📜api.go
 ┃ ┣ 📜config.go
 ┃ ┗ 📜routes.go
 ┣ 📂cmd
 ┃ ┗ 📜main.go
 ┣ 📂doc
 ┃ ┗ 📜openapi.yaml
 ┣ 📂internal
 ┃ ┣ 📂rate
 ┃ ┃ ┣ 📜config.go
 ┃ ┃ ┣ 📂curxrt
 ┃ ┃ ┃ ┣ 📜alphavantage.go
 ┃ ┃ ┃ ┣ 📜coinapi.go
 ┃ ┃ ┃ ┣ 📜coinyep.go
 ┃ ┃ ┃ ┣ 📜curxrt.go
 ┃ ┃ ┃ ┣ 📜ninjas.go
 ┃ ┃ ┃ ┗ 📜xratehost.go
 ┃ ┃ ┣ 📜event.go
 ┃ ┃ ┣ 📜handler.go
 ┃ ┃ ┗ 📜rate.go
 ┃ ┗ 📂subs
 ┃   ┣ 📜config.go
 ┃   ┣ 📜event.go
 ┃   ┣ 📜handler.go
 ┃   ┣ 📜repo.go
 ┃   ┣ 📜repo_test.go
 ┃   ┣ 📜sender.go
 ┃   ┗ 📜subs.go
 ┣ 📂log
 ┃ ┗ 📜sys.log
 ┣ 📂sys
 ┃ ┣ 📂env
 ┃ ┃ ┣ 📜env.go
 ┃ ┃ ┗ 📜env_test.go
 ┃ ┣ 📂event
 ┃ ┃ ┗ 📜event.go
 ┃ ┣ 📂filestore
 ┃ ┃ ┣ 📜filestore.go
 ┃ ┃ ┗ 📜filestore_test.go
 ┃ ┣ 📂logger
 ┃ ┃ ┗ 📜logger.go
 ┃ ┗ 📂web
 ┃   ┣ 📜errors.go
 ┃   ┣ 📜middlewares.go
 ┃   ┣ 📜middlewares_test.go
 ┃   ┣ 📜params.go
 ┃   ┣ 📜request.go
 ┃   ┣ 📜respond.go
 ┃   ┗ 📜web.go
 ┣ 📂test
 ┃ ┣ 📂mock
 ┃ ┃ ┣ 📜email_repository.go
 ┃ ┃ ┣ 📜email_sender.go
 ┃ ┃ ┣ 📜getter.go
 ┃ ┃ ┗ 📜subscriber.go
 ┃ ┣ 📜Dockerfile
 ┃ ┗ 📜postman.json
 ┣ 📜.gitignore
 ┣ 📜.golangci.yml
 ┣ 📜Dockerfile
 ┣ 📜Makefile
 ┣ 📜README.md
 ┣ 📜docker-compose.yml
 ┣ 📜go.mod
 ┗ 📜go.sum

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
## Entities 
--TODO: Finish
```mermaid
classDiagram
    class App {
        <<struct>>
        sig chan os.Signal
        log *logger.Logger
        web *web.Web
    }
    class Route {
        <<type>>
    }
    class ConfigAggregate {
        <<struct>>
        Api Config
        Rate rate.Config
        Subscription subs.Config
    }
    class Config {
        <<struct>>
        Name string
        Path string
        Version string
        Origin string
    }
    class Handler {
        <<struct>>
        rate ExchangeRateService
    }
    class ExchangeRateService {
        <<interface>>
        GetExchangeRate(ctx context.Context, currency CurrencyPair) (*ExchangeRate, error)
    }
    class Web {
        <<struct>>
        mux *httprouter.Router
        mws []Middleware
        sig chan os.Signal
    }
    class Middleware {
        <<type>>
    }
    class SubscriptionService {
        <<interface>>
        Subscribe(context.Context, Subscriber) error
        SendEmails(context.Context) error
    }
    class Response {
        <<struct>>
        Message string
    }
    class Subscriber {
        <<struct>>
        Address *mail.Address
        Topic Topic
    }
    class RateConfig {
        <<struct>>
        Provider struct
        Client struct
    }
    class ProviderConfig {
        <<struct>>
        Name string
        Endpoint string
        Header string
        Key string
    }
    class SubsConfig {
        <<struct>>
        Sender SenderConfig
        Repo RepoConfig
    }
    class SenderConfig {
        <<struct>>
        Address string
        Key string
    }
    class RepoConfig {
        <<struct>>
        Data string
    }
    class Storer {
        <<interface>>
        Store(Subscriber) error
        FetchAll() ([]Subscriber, error)
    }
    class Repo {
        <<struct>>
        Storer
    }
    class Logger {
        <<struct>>
        *zap.SugaredLogger
    }
    App o-- Route
    App --> ConfigAggregate
    App --> Web
    ConfigAggregate o-- Config
    ConfigAggregate o-- RateConfig
    ConfigAggregate o-- SubsConfig
    Handler o-- ExchangeRateService
    Web -- Middleware
    SubscriptionService -- Subscriber
    SubscriptionService -- Response
    RateConfig o-- ProviderConfig
    SubsConfig o-- SenderConfig
    SubsConfig o-- RepoConfig
    Repo o-- Storer
    App --> Logger
```