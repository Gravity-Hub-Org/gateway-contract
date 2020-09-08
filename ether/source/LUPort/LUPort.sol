pragma solidity >=0.5.16 <=0.6.6;

import "./Token.sol";
import "../interfaces/ISubscription.sol";

contract LUPort is ISubscription {
    address public nebula;
    Token public tokenAddress;

    constructor(address _nebula, address _tokenAddress) public {
        nebula = _nebula;
        tokenAddress = Token(_tokenAddress);
    }

    function attachData(bytes calldata data) external {
    }
    function createTransferUnwrapRequest(uint amount, address receiver) public {
    }
}
