# Event Processor
This is an event pipeline application
where several producers can post their events to be validated and persisted
so that another application can send the event to the end client.

Event contract validations will be described in a configuration file.

# Architecture
```mermaid
---
title: Integration Archteture
---
flowchart LR
    EP1(Event Producer) ---> B
    EP2(Event Producer) ---> B
    EP3(Event Producer) ---> B
    EP4(Event Producer) ---> B
    B[(MessageBroker)]   --> R

    subgraph EPs[Event Processor]
        direction LR
        R[Receiver] --> V
        V[Validator] --> P
        P[Persist]
    end

    P --> DB
    DB[(DataBase)] --> S
    S(Sender) --> C1
    S(Sender) --> C2
    S(Sender) --> C3
    C1([Client])
    C2([Client])
    C3([Client])
```

```mermaid
---
title: Event Processor activity diagram
---
flowchart TD
    R[Receiver]    --> UC
    UC[SaveEventUC] <--> V
    V{isValid}
    C[Contract]    -.usedBy.-> V
    
    V --invalid--> RV[Reject Event]
    V --valid--> P[Persist]
    RV[Reject Event] --> DLQ(events.dlq)
```



# @TODO
 * [ ] Criar um consumer para os eventos
 * [ ] Criar um leitor de contrato de configuração
 * [ ] Criar um validator que utiliza o contrato para rejeitar os eventos inválidos
 * [ ] Utilizar um banco para persistir os eventos e ser futuramente lido pelo sender
 * [ ] Criar um producer parametrizável para colocar eventos no pipeline
 * [ ] Criar testes de aceitação
 * [ ] Criar testes de carga
 * [ ] Otimizar o fluxo
 * [ ] Criar um container docker para o app
 * [ ] Mover para o LocalStack
 * [ ] Usar terraform para declarar os recursos