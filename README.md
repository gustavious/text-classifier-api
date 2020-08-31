# Text Classifier API
This is a GOlang simple RESTful API to classify text. Given a small chunk of data, the naive bayes algorithm implementation is able to classify a message in a defined set of categories.


## Learnings
Perhaps the most important part of this learning task. As I've been for a long time working on node, python and rust(back) and flutter, react(front) is pretty neat to try something new.

First, go feels like a very straightforward and scripting-like language. Compiled and typed out of the box. It was cool that you can decide to use a pointer or clone the actual value to pass it as a parameter. The package manager is not as robust as npm but is very lightweight. Also, you don't need a bunch of dependencies to get started. It has an easy to read syntax and it's paradighm seems more in the imperative side with notions of OOP. Also, I'm curious about how to handle concurrency and how the modules work here. Overall, it seems very performant and easy to learn.

This task was initially planned as a neuronal network(perceptron, maybe a feed forward network). Bbecause of the dataset size, and the kind of problem it feels way more suitable to be a NLP classifier. For this purpose I decided to go with a naive bayes algorithm to classify text in categories based on the number of appearences of keywords in a message. This algo still can be optimized a lot but works fine for the purpose of the task. 

### Requirements
    - dep
    - chi
    - Go @ 1.15
  
## Run locally
1. clone the api and cd into the directory `cd tex-classifier-api`
2. run `dep ensure -update`
3. run `make run`

## Api
- `POST /message` Sent a new message to the server to be processed
    
### Requirements
    - dep
    - chi
    - Go @ 1.15
    
#### What to improve:
- Unit test and integration test
- A CI/CI pipeline to integrate (e.g. github actions)
- Dockerize the api
- A DB as persistence mechanism (e.g. mysql or postgres)
- Use Go modules to decouple a bit more the components
- As we start to add more and more functionalities to our api; it becomes necessary to divide layers (models, services, controllers, persistence, types)
- A linter and guidelines
- An ORM or query builder
- Serverless integration (lambdas or cloud run)
- Train the model with more data