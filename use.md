go build
--------------------------------------------------------

NODE 3000
./blockchain_go 3000 createwallet
1NpmjossotfKnGvJvCLgo69id4Tjz9bV4z

./blockchain_go 3000 createblockchain -address 1NpmjossotfKnGvJvCLgo69id4Tjz9bV4z
8fb60d28eef979f8391f20ea5946be8caf54b54638b6a9543cc47783b55cb5fa

cp blockchain_3000.db blockchain_genesis.db 


---------------------------------------------------------


NODE 3001

./blockchain_go 3001 createwallet
1LkHexUdC6x9kuKP2JzRwHmgJLwVYXpAXZ
197QcFY64EqtYfw8euhmGELUfVm3L7yFCc
1Eed1QTXfU9ytNpsMYvAuApRaq6A1vZKuh
-----------------------------------------------

发币
./blockchain_go 3000 send -from 1NpmjossotfKnGvJvCLgo69id4Tjz9bV4z -to 1LkHexUdC6x9kuKP2JzRwHmgJLwVYXpAXZ -amount 10 -mine
./blockchain_go 3000 send -from 1NpmjossotfKnGvJvCLgo69id4Tjz9bV4z -to 197QcFY64EqtYfw8euhmGELUfVm3L7yFCc -amount 10 -mine

./blockchain_go 3000 startnode -miner 1NpmjossotfKnGvJvCLgo69id4Tjz9bV4z




cp blockchain_genesis.db blockchain_3001.db
./blockchain_go 3001 startnode

./blockchain_go 3001 getbalance -address 1LkHexUdC6x9kuKP2JzRwHmgJLwVYXpAXZ
./blockchain_go 3001 getbalance -address 197QcFY64EqtYfw8euhmGELUfVm3L7yFCc
./blockchain_go 3001 getbalance -address 1Eed1QTXfU9ytNpsMYvAuApRaq6A1vZKuh

./blockchain_go 3000 getbalance -address 1NpmjossotfKnGvJvCLgo69id4Tjz9bV4z


cp blockchain_genesis.db blockchain_3002.db
./blockchain_go 3002 createwallet
15rcpPjtDSnJVDovDaX8tfEv93hqUQxHXZ
./blockchain_go 3002 startnode -miner 15rcpPjtDSnJVDovDaX8tfEv93hqUQxHXZ

./blockchain_go 3001 send -from 1LkHexUdC6x9kuKP2JzRwHmgJLwVYXpAXZ -to 1Eed1QTXfU9ytNpsMYvAuApRaq6A1vZKuh -amount 1


./blockchain_go 3000 getbalance -address 15rcpPjtDSnJVDovDaX8tfEv93hqUQxHXZ