pragma solidity >=0.5.16 <=0.6.6;


contract Random {
    event NewRandom(bytes32 random, uint256 roundEnd);

    bytes32 public random;
    uint256 roundEnd;
    uint256 timeoutNewRound;

    constructor(uint256 newTimeoutNewRound) public {
        timeoutNewRound = newTimeoutNewRound;
    }

    function newRandom() public {
        require(block.number > roundEnd, "wait new round");
        roundEnd = block.number + timeoutNewRound;
        emit NewRandom(blockhash(block.number), roundEnd);
    }
}