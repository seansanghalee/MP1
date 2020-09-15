# MP1

By Socially Distancing Group

# How to Run
First, you need to edit config.txt to provide the minimum and maximum delay time
(in milliseconds), as well as IDs, IP Addresses, and Port #'s for all processes
that will be used in the network.

An example of a configuration file looks like this

```
1000 3000
1 127.0.0.1 123
2 127.0.0.1 456
3 127.0.0.1 789
```
indicating, in this case, that the minimum delay time is 1 second, the maximum delay time is 3 seconds,
and there will be three processes connected to each other.

Then, you will open three terminal windows which will run process.go individually.

```
// first terminal window
go run process.go 1:127.0.0.1:123
```
```
// second terminal window
go run process.go 2:127.0.0.1:456
```
```
// third terminal window
go run process.go 3:127.0.0.1:789
```

Each process sends/receives messages to/from other processes in the network.

For example, to send a message "Hello" from process 1 to process 2,
simply type on process 1's terminal window
```
send 2 Hello
```
which also prints out the time the message was sent.
```
send 2 Hello
sent [Hello] to process 2, system time is XXX
```

Process 2 will receive "Hello" and print out
```
received [Hello] from process 1. system time is XXX
```
A similar output with varying ID's, messages and time will be
shown on the terminal window for each sending and receiving process.

# Architecture and Design

![Diagram](diagram.png)

### Config
Package config provides methods that allows the process to read the
configuration file and extract necessary data. Also, it has a method
to display them to the terminal window in a readable format.

### Packet
There is a struct of type Packet that stores the sender's process ID and the message taken from user input.
It is then encoded using gob and sent through TCP network connection.
```
type Packet struct {
    Source, Message string
}
```

### Sender
Sender package includes methods that allow process to send the message
to the destination process. It prompts user input, and Sender dials the
according port when the user types in the message with the process number to send it to.

In network layer, Sender first creates a packet, then starts a goroutine
that calls Unicast_send method to connect to the destination process via TCP.
Sender encodes the packet using gob, and Sender closes the connection once
the struct is encoded and sent.

### Receiver
Receiver package includes methods that allow process to constantly listen to
incoming dials and connect them with the dialing process using TCP. When the connection
establishes, receiver decodes the Packet struct sent from the other end.

Finally, Receiver calls Unicast_receive method to print out the message and where it
came from.

### Process
process.go includes the main function that first extracts data from the configuration file.
This is done to find the number of goroutines to start in a for-loop that will listen to all
available ports.

After that, the code consists of two parts which

1. Process listens to other processes, trying to receive messages from them. (Receiver)

2. Process sends over the message input from the user to another process. (Sender)

# Notes
In this program, process is continually listening to the port. However, it is not continually
dialing to all possible ports. Only when the user types in the instruction, the process
establishes connection with the destination port. For the prior case, the idea of trying to access
the already established TCP channel from the outside was hard to grasp. Thus, the group chose the
to implement the second model which also seemed more efficient.

Another thing the group did not implement in this program is the Queue Professor Tseng demonstrated
in the example slide. There exists a message queue that stores source and message, but the group
didn't find it necessary to build a centralized data structure that stores all the messages
before printing them out.
