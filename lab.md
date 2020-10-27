# Message Broker in Go

We are going to use a variant of the publish-subscribe model which uses a
message broker for distributing work between many contributors.

The typical publish-subscribe model sends messages to all subscribers signed up
to a topic. This works for distributing events that all subscribers should see
or act on. For the case where we want just one subscriber to see an event, we
want something more like a [worker pool](https://gobyexample.com/worker-pools)
that accepts new machines joining the pool remotely and uses 'take' semantics to
avoid workers duplicating work.

## Part 1: The Multiply Factory

Look at the skeleton code you've been given. There are two complete components:

1. A `broker` that handles creating channels, publishing `stubs.Pair` events to
those channels, and subscriptions to those channels. When a subscriber sends a
`stubs.Subscription` to a running instance of the broker, it will get callbacks
for work whenever it is available.

2. A `miner` that generates `stubs.Pair` events and publishes them to a
`multiply` channel on the broker.

You should be able to launch the broker with `go run broker/broker.go` and the
miner with `go run miner/miner.go` -- nothing visible should happen except that
the broker will print out a notification that the `multiply` channel has been
created.

Your first task is to complete the `Factory` in `factory/factory.go`. This
process should be a worker that:

- sets up an RPC *server* that registers a `Multiply` procedure. This procedure
  should accept a `stubs.Pair` and respond with a `stubs.JobReport`. You may
also want to print out the operation, so you can see what the instance is doing
while it's running. 
- connects to the broker as an RPC *client*.
- sends a `stubs.Subscription` request to the `multiply` channel, containing its
  own `ip:port` string and the correct string for the broker to use to call its
`Multiply` procedure.

You'll know if this is working correctly because once the factory is subscribed
the broker will start sending it work and printing out the results of jobs.

You should be able to: 

1.  Stop and restart the factory.

2. run a second instance of the factory, and have it also process work.  *Note:
you will need to tell the second factory to use a different port for its RPC
server*.


## Part 2: The Multiply-Divide Pipeline

Your `Multiply` function is producing a lot of results, but the work is only
using one channel on the broker at the moment. Modify your `factory.go` code so
that for every *two* Multiply results produced, a `stubs.Pair` is published to a
new `divide` channel.

To accomplish this, you'll need to modify what your `Multiply` method is doing,
and you'll likely want another goroutine to monitor a [buffered
channel](https://gobyexample.com/channel-buffering). 

Also create a `Divide` procedure for your Factory, and subscribe your instance
to the `divide` channel you created.

You'll know the code is working when the broker reports results for division
operations. Again, you should be able run multiple instances of your Factory (on
different ports), and stop and start each of them.


## Part 3: Triplets

The entire pipeline currently operates on `stubs.Pair`s. Modify it so that it
instead works with Triplets of three integers. As well as the `factory`, you
will need to edit the `miner` and `broker` code to accomplish this, and make a
decision about what Divide means for three integer arguments. 
