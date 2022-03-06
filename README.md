# Go + Cassandra example for timestream data

Cassandra might be one database that you haven't heard of because
you were too deep on the realm of SQL (PostgreSQL, MySQL) and MongoDB for the NoSQL part.
It is a NoSQL database, which data model uses a wide-column store (or columnar - as others
might say about it), optimized for high-write and low-read operations, and has out-of-the-box
feature for distributed database, so you can run the database within a few nodes and make them
work with each other.

There are small amount of references about Go + Cassandra, but considering what Cassandra
can do and my current project has a requirement that would be done best by using Cassandra,
I'm going to give it a try.

Run a Cassandra instance easily [via Docker](https://hub.docker.com/_/cassandra) and create
a new keyspace with:

```
CREATE KEYSPACE "timestream" WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1 };
```

Run the application with:
```
go run .
```

And do operation with the supported handlers. See `main.go` and `handler.go` file.

Feel free to explore.

[MIT License](./LICENSE)