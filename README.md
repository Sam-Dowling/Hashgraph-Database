# Distributed Fault-Tolerant Data Structure

A distributed database runnning on a Hashgraph network with the [Swirlds&#169; consensus algorithm.](http://www.swirlds.com/downloads/SWIRLDS-TR-2016-01.pdf)

## What is this and what are it's advantages?

Distributed unpermissioned system such as Blockchain require a proof-of-work to prevent malicious nodes from disrupting the natural flow of the Blockchain.

A proof of work is a piece of data which is difficult (costly, time-consuming) to produce but easy for others to verify and which satisfies certain requirements.

This system negates the need for proof-of-work but keeps all the benefits and also introduces the concept of Fairness and [Byzantine fault tolerance](https://en.wikipedia.org/wiki/Byzantine_fault_tolerance).

### What's a Hashgraph?

A Hashgraph is a data structure that records who gossiped to whom, and in what order.
Gossping; in the context is where one node selected another node at random and they both synchronize their hasgraphs with each other.

### What's a Consensus Algorithm?

For the system to work, every node must have the same exact copy of the Hashgraph this is acheived through the use of a Consensus Algorithm. Traditionally this would involve each node voting or a central server that would monitor each node. The Swirlds&#169; consensus algorithm enables virtual voting.

`Each node knows what every other node knows and will know how each node will vote.`

### Installing

This project uses Golang. Golang was chosen for it's high-performance and it's ease of parallelization.

```
This project was written and tested on Go 1.8
```

## Author

* **Sam Dowling** - *Software Development Student* - [Institute of Technology Tralee](http://ittralee.ie)

This project was undertaken as part of my 4<sup>th</sup> year of Computing with Software Development 

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Peter Given - FYP Supervisor 
* [Swirlds&#169;](http://www.swirlds.com/) - Consensus Algorithm
