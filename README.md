#Go Blockchain

This project's aim is to have a better overview of what is a blockchain and how it's work.

   - [x] Simple Blockchain
   - [x] Make it work on a TCP Server
   - [x] Build a simple cli to use the server
   - [x] Build a simple client cli
   - [x] Use protobuf
   - [x] Implement a Peer to Peer architecture
   - [ ] Add a peers sharing process
   - [ ] Implement a node discovery algorithm
   - [ ] Add persistency of the blockchain
   - [ ] Implement a Proof of Work or a Proof of Stake validation
   - [ ] Implement a consensus mechanism
   - [ ] Tests with lot of nodes
   
   
## Build Project

To build the project simply run
```sh
user$ make build
```

## Run the programm
To simulate a first node open a terminal window and run, note the ip and the port 
```sh
node1$ gochain start
10:28AM INF Listening for peers. address=tcp://0.0.0.0:8555
*** > Enter a message: 
```

To simulate a second node open a terminal window. Use a different port
```sh
node2$ gochain start -p <PORT> --peers <ADDRESS>
```
example :
 ```sh
 node2$ gochain start -p 8556 --peers tcp://0.0.0.0:8555
 ```
 
 You should now see on node1 window, enter your messsage
 ```sh
2019/01/02 10:31:08 *** New connection from tcp://0.0.0.0:8556.
10:31AM INF Connected to peer(s). peers=["tcp://0.0.0.0:8556"]
*** > Enter a message: Test
```

Once you enter your message, you should go to node2 window and see the blockchain with this 
new message contained in a block was broadcast:
```sh
2019/01/02 10:32:16 [UPDATE BLOCKCHAIN] From: tcp://0.0.0.0:8555
blocks:<timestamp:1546453731 data:"Genesis Block" hash:"\363H/3\275Si_\373\342n\023(\252)h<Y\237i\263W\303\301\273\242o\356P\313\217\251" > blocks:<timestamp:1546453936 prevBlockHash:"\363H/3\275Si_\373\342n\023(\252)h<Y\237i\263W\303\301\273\242o\356P\313\217\251" data:"Test\n" hash:"'c\3063\202\233\352\024\337\215\264\004\003\037\304\370\340\271b\200\310\365\037W\350]ZbqG\372]" >
```

