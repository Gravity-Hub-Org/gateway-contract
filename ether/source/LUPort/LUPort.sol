pragma solidity ^0.7;

import "../Token/Token.sol";
import "../interfaces/ISubscriberBytes.sol";

contract LUPort is ISubscriberBytes {
    address public nebula;
    Token public tokenAddress;

    constructor(address _nebula, address _tokenAddress) {
        nebula = _nebula;
        tokenAddress = Token(_tokenAddress);
    }

    function attachValue(bytes calldata value) external override {
    }

    function createTransferUnwrapRequest(uint amount, address receiver) public {
    }
}
