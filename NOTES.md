# Glossary & Notes of some terminologies learned

### *Zookeper*:

Zookeper helps distributed system such as Hadoo, Kafka, to handle their configuration (sharing config across all nodes),
naming (find a machine in a cluster of 1000 servers), sincronization()
coordination, etc.

https://www.youtube.com/watch?v=gifeThkqHjg

### *Kafka*:

* Brokers: a computer instance or container running
    * Manage partitions
    * handle write and request
    * manage replication of partitions
    * https://www.youtube.com/watch?v=jHnyBSUVcOU

* Topic: Where a meesage arrives?
    * Topic is replicated on all brokers
    * Multi subscriber
    * Logical entity


* Partitionings: Distribution of logs into multiple nodes
    * To wich partition can depend of the key
    * division of a topic in multiple parts
    * If no key, splits into multiple partitions
    * keys work for ordering of the events
    * each partition if replicated across multiple nodes
    * only one partition will be active called Leader, others are followers
    * https://www.youtube.com/watch?v=y9BStKvVzSs
    * https://www.youtube.com/watch?v=q72vMNXoQ2E

* Replication factor: Definition of replication of partitions in multiple brokers
    *

### *Oracle DB*:

* Contanerized databases

https://docs.oracle.com/en/database/oracle/oracle-database/18/xeinw/connecting-oracle-database-xe.html

* Solution to create db on startups.

https://medium.com/@m.emmanuel/how-to-navigate-across-cdb-and-pdbs-in-oracle-18c-xe-781ac217f5b0

## Resources

##### Zookeeper & kafka deployment

https://github.com/d1egoaz/minikube-kafka-cluster

##### Oracle setup un minikube:

https://ronekins.com/2020/03/14/running-oracle-12c-in-kubernetes-with-minikube-and-virtualbox/

#### Debezium kafka guide:

https://www.startdataengineering.com/post/change-data-capture-using-debezium-kafka-and-pg/

### Kafka explanation:

https://www.youtube.com/watch?v=QYbXDp4Vu-8
