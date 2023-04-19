# A tech talk about kafka (Vccorp - 01/04/2023) 

## Batching config

**linger.ms**
The producer groups together any records that arrive in between request transmissions into a single batched request.


Too many small I/O operations can occur both between the client and the server and in the server's own persistent operations.

## How do clients connect to a Kafka Cluster (bootstrap server)

A client that wants to send or receive messages from the Kafka cluster may connect to any broker in the cluster. Every broker in the cluster has metadata about all the other brokers and will help the client connect to them as well, and therefore any broker in the cluster is also called a bootstrap server.

The bootstrap server will return metadata to the client that consists of a list of all the brokers in the cluster. Then, when required, the client will know which exact broker to connect to to send or receive data, and accurately find which brokers contain the relevant topic-partition.

In practice, it is common for the Kafka client to reference at least two bootstrap servers its connection URL, in the case one of them not being available, the other one should still respond to the connection request.

## Kafka Topic Replication Factor

Data Replication helps prevent data loss by writing the same data to more than one broker.

In Kafka, replication means that data is written down not just to one broker, but many.
The replication factor is a topic setting and is specifies at topic creation time.
- A replication factor of 1 means no replication. 
- A replication factor of 3 is a commonly used replication factor as it provides the right balance between broker loss and replication overhead.

## What are Kafka Partitions Leader and Replicas?

For a given topic-partition, one Kafka broker is designated by the cluster to be responsible for sending and receiving data to clients. That broker is known as the leader of the topic partition. Any other broker that is storing replicated data to that partition is referred to as a replicas.

Therefore, each partition has one leader and multiple replicas.

## What are In-Sync Replicas (ISR)?

An ISR is a replica that up to date with the leader broker for a partition. Any replica that is not up to date with the leader is out of sync.

## Kafka producers acks setting

Kafka producers only write data to the current leader broker for a partition.
Kafka producers must also specify a level of acks to specify if the message must be written to a minimum number of replicas before being considered a successful write.

- acks = 0
When acks = 0 producers consider messages as "written successfully" the moment the message was sent without waiting for the broker to accept it at all.
If the broker goes offline  or an exception happens, we won't know and will lose data. This is useful for data where it's okay to potentially lose messages, such as metrics collection, and produces the highest throughput settings because the network overhead is minimized.

- acks = 1
When acks = 1 producers consider messages as "written successfully" when the message was acks by only the leader. 
Leader response is requested, but replication is not a guarantee as it happens in the background. If an ack is not received, the producer may retry the request. If the leader broker goes offline unexpectedly but replicas have'nt replicated the data yet, we have a data loss

- acks = all

When acks=all, producers consider messages as written successfully when the message is accepted by all in-sync replicas (IRS)

The leader replicas for a partition check to see if there are enough in-sync replicas for safely writing the message (controlled by the broker setting min.insync.replicas). The request will be stored in a buffer until the leader observes that the follower replicas replicated the message, at which point is successful acknowledgement is sent back to the client. 

The min.insync.replicas can be configured at the topic level and the broker-level. The data is considered committed when it is written to all in-sync replicas - min.insync.replicas. A value of 2 implies that at least 2 brokers that are ISR (including leader) must respond that they have the data.


## Minimum In-Sync Replicas

## Kafka Topic Durability & Availability

**Durability** Replicas factor of N, lose up N-1 still recover

**Availability**

- Reads: all replicas is ISR. The topic will be available for reads.
- Writes

## Kafka Consumers Replicas Fetching

Kafka consumers read by default from the partition leader.

Since Kafka 2.4, it is possible to configure consumers to read from in-sync replicas instead (usually the closest).

## Zookeeper with Kafka

Zookeeper is used to track cluster state, membership and leadership

How do the Kafka brokers and clients keep track of all the Kafka brokers if there is more than one? 
Zookeeper is used for metadata management in the Kafka world. 
- Zookeeper keeps track of which brokers are part of the Kafka cluster
- Zookeeper is used by kafka brokers to determine which broker is the leader of a given partition and topic and perform leader elections
- Zookeeper stores configuration for topics and permissions
- Zookeeper sends notifications to Kafka in case of changes (new topic, broker dies, broker come up, delete topic...)
- Zookeeper does NOT store consumer offsets with Kafka clients.

## Kafka KRaft Mode

## Kafka Topic Partitions and Segments

## Delivery Semantics for Kafka Consumers

A consumer reading from a Kafka partition may choose when to commit offsets. That strategy impacts the behaviors if messages are skipped or read twice upon a consumer restart.

**A Most Once Delivery**

Offsets are committed as soon as a message batch is receive after calling poll(). If the subsequent processing fails, the message will be lost. It will not be read again as the offsets of those messages have been committed already. This may be suitable for systems that can afford to lose data.

**A Least Once Delivery**

## Sequential I/O access

## Efficient data transfer through zero copy

Many Web applications serve a significant amount of static content, which amounts to reading data off of a disk and writing the exact same data back to the response socket. It's somewhat inefficient: the kernel reads the data off of disk and pushes it across the kernel-user boundary to be written out to the socket. In effect, the application serves as an inefficient intermediary that gets the data from the disk file to the socket.

Applications that use zero copy request that the kernel copy the data directly from the disk file to the socket, without going through the application. Zero copy greatly improves application performance and reduces the number of the context switches between kernel and user mode.


