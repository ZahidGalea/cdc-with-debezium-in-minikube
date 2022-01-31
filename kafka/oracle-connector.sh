


{
    "name": "inventory-connector",
    "config": {
        "connector.class" : "io.debezium.connector.oracle.OracleConnector",
        "database.hostname" : "<ORACLE_IP_ADDRESS>",
        "database.port" : "1521",
        "database.user" : "c##dbzuser",
        "database.password" : "dbz",
        "database.dbname" : "ORCLCDB",
        "database.server.name" : "server1",
        "tasks.max" : "1",
        "database.pdb.name" : "ORCLPDB1",
        "database.history.kafka.bootstrap.servers" : "kafka:9092",
        "database.history.kafka.topic": "schema-changes.inventory"
    }
}