#MP1

By Socially Distancing Group

#How to Run
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

Each process sends messages to and receives messages from other processes in the network.

For example, to send a message "Hello" to from process 1 to process 2,
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
Similarly, a similar output will be shown on the terminal window
with different process IDs and system time when process 2 (or 3)
sends a message to process 3 (or 1).

#Architecture and Design

###Config
Package config provides methods that allows the process to read  configuration file and
extract necessary data and display them to the window.

###Packet
There is a struct type named Packet that holds the message taken from user input.
It is then encoded using gob and sent through TCP network connection.
```
type Packet struct {
    Message string
}
```

###Sender
Sender package includes methods that allow process to send the message input
from the user to the destination process. It prompts user input, and Sender
dials the according port to connect via TCP when the user types in the message
with the process number to send it to.

In network layer, sender calls Unicast_send method to connect to the destination
process via TCP. First it creates a Packet then encode the struct using gob.
Once the struct is encoded and sent, sender closes connection and asks for
user input again.

###Receiver
Receiver package includes methods that allow process to constantly listen to
incoming dials and connect them with the process via TCP. When the connection
establishes, receiver decodes the packet struct sent from the other end.
Unicast_receive method is called

###Process
process.go includes the main function that first extract data from the configuration file.
This is done to find the number of goroutines to start in a for-loop to listen to all
possible ports.

After that, the code consists of two parts which

1. Process listens(Receiver) to other processes, trying to receive messages from them.

2. Process sends(sender) over the message input from the user to another process. 

#Notes