# Roadmap

## Done

+ Subscription
+ Fix data race on timer management
+ Message dispatch to client using the library (publish message
  received are dispatched)
+ Basic publish support
+ Support reset of timer when sending a packet from client
+ Implement unsubscribe
+ Implement disconnect
+ Make connect synchronous. It is easy in Go to wrap the call to make it asynchronous: http://stackoverflow.com/a/6329459/559289
+ Reconnect with backoff strategy (basic)
+ Central place for errors definitions
+ Ability to set clientID.
+ Handle teardown & reconnect from either sender or receiver.
+ Support persistent session option
+ Use new Marshaller: Preallocate buffer of correct size in marshallers to improve performance (+buffer write can return errors).
+ Rename PDU to a name more in line with MQTT specification.
+ Address go vet + various other linter

## TODO

- errcheck: check that all required errors are handled properly (errcheck)
- TLS support: Use URL scheme to define connection to server: tcp:// tls://
- Ability to set session as persistent. If session is persistent, there is no need to resubscribe on reconnect.
  See: http://www.hivemq.com/blog/mqtt-essentials-part-7-persistent-session-queuing-messages
- We need to setup subscriptions after background reconnect if there was not persistent session
- Manage Packet ID during session.
- Implement store interface and backend to ensure no message loss in client.
- Use context to clean data flow ? (https://www.youtube.com/watch?v=3EW1hZ8DVyw&list=PL2ntRZ1ySWBf-_z-gHCOR2N156Nw930Hm)
- Support timeout on PingResp to trigger reconnect
- Support subscription based on callbacks or on channels
- QOS 1 and 2
- Authentication with username, password. They can be place in URL scheme. tcp://username:password@server 
- Certificate based authentication
- More unit tests
- Example of publish / subscribe sharing Go structures with encoding/gob (RPC like)
- Support command-line option for examples (to pass server, port, username, ...)
- Support for MQTT-SN spec ? http://mqtt.org/new/wp-content/uploads/2009/06/MQTT-SN_spec_v1.2.pdf
