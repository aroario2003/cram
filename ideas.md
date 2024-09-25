# Ideas for the project

## Database Ideas

- keep the database connection alive via unix domain sockets
  - external binary creates socket
  - cli connects to socket in order to get database connection
  - database connection is never restarted unless external binary is killed

pros: 
- increases efficiency

cons: 
- complicated
- possibly hard to maintain or code

## Caching ideas

- cache queries and results into hashtable
- write hashtable to file before exiting
- when restart, read file and recreate hashtable

pros:
- caching makes the program run faster

cons:
- time to read file could be long
- time to write file could be long
- time to create hashtable could be long
	
## General ideas

- make cli and gui rebuild itself with external binary (live reloading)

pros: 
- good for debugging and testing

cons:
- possibly tricky to implement correctly
