# gobarebones

W.I.P.

A stripped down and embeddable implementation of a DataDog agent in golang.

Five objectives:
    - Extremely simple setup
    - Few dependencies, all of them must be from std lib
    - Push metrics
    - Push logs
    - Sending in chunks using a concurrent goroutine

TODO:
    - Logs
    - Finish the agent
    - Testing
