# gRPC trace test

The gRPC trace on unimplemented service calls will show up as *Active* state, and
with be in that state for the entire life of the server.

Whether it's due to the client using a new proto file that contains a new `rpc` 
declaration on an already existing service, or a service that's defined, but 
not implemented by the server.

![Screenshot of trace][screenshot]

[screenshot]: https://raw.githubusercontent.com/wari/grpc-trace-test/master/screenshot.png
