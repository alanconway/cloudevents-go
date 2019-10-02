/*

Package binding defines interfaces for protocol bindings.

NOTE: Most applications that emit or consume events can use the client
package. This package is for implementing new protocol bindings and
intermediaries; processes that forward events between protocols, rather than
emitting or consuming events themselves.

Protocol Bindings

A protocol binding implements at least Message, Sender and Receiver, and usually
Encoder.

Receiver: receives protocol messages and wraps them to implement the Message interface.

Message: converts to protocol-neutral cloudevents.Event or structured event
data. It also provides methods to manage acknowledgment for reliable
delivery across bindings.

Sender: converts arbitrary Message implementations to a protocol-specific form
and sends them.

Message and ExactlyOnceMessage provide methods to allow acknowledgments to
propagate when a reliable messages is forwarded from a Receiver to a Sender.
QoS 0 (unreliable), 1 (at-least-once) and 2 (exactly-once) are supported.

Intermediaries

Intermediaries can forward Messages from a Receiver to a Sender without
knowledge of the underlying protocols. The Message interface allows structured
messages to be forwarded without decoding and re-encoding. It also allows any
Message to be fully decoded and examined if needed.

*/
package binding
