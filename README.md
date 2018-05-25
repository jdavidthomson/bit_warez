# bit_warez
warez build around bitcoin and its networking protocols

The vision for bit_warez is to build warez supported by the bitcoin protocol.

The bitcoin peer-to-peer networking layer is a robust and seemingly three-nines protocol layer.  With that in mind why not build warez ontop of the bitcoin protocols.

The tools will use bitcoin to pay for services.
The tools will use the bitcoin networking layer for discovery of additional nodes.
The tools will run alongside bitcoind on a full node.
The tools will utilize localhost (or configured as you see fit) to reach its initial BTCd instance.
The tools will utilize RPC commands to interact with the full node - wallet passphrases are not necessary to be entered.
The tools are written in golang and are extendable to use python http servers (like tornado) to access python only libraries.

The belief of the author is that most full-nodes are not in use most of the time.  When new block is transmitted it is validated by a full-node.  That means that 5-7 minutes of each cycle the node is not doing anything.  So why not utilze that CPU time for doing money-generating work.  Examples are utilizing Spacy for NLP, doing distributed calculations, and routing packets around the internet (think tor but not tor).
