pragma solidity >=0.5.16 <=0.6.6;

import "./Token.sol";
import "../interfaces/ISubscription.sol";

contract IBPort is ISubscription {
    enum Status {
        None,
        New,
        Rejected,
        Success,
        Returned
    }

    address public nebula;
    Token public tokenAddress;
    mapping(bytes32 => Status) public swapStatus;

    constructor(address _nebula, address _tokenAddress) public {
        nebula = _nebula;
        tokenAddress = Token(_tokenAddress);
    }

    function data2address(bytes32 d) internal pure returns (address payable) {
        return address(uint160(uint(d)));
    }

    function data2status(bytes32 dd) internal pure returns (Status) {
        uint d = uint(dd);
        if (d == 0) return Status.None;
        if (d == 1) return Status.New;
        if (d == 2) return Status.Rejected;
        if (d == 3) return Status.Success;
        if (d == 4) return Status.Returned;
        revert("invalid status");
    }

    function attachData(bytes32[] calldata data) external {
        require(msg.sender == nebula, "access denied");
        for (uint pos = 0; pos < data.length; ) {
            bytes32 action = data[pos]; pos++;

            if (action == bytes32("mint")) {
                bytes32 swapId = data[pos]; pos++;
                uint amount = uint(data[pos]); pos++;
                address receiver = data2address(data[pos]); pos++;
                mint(swapId, amount, receiver);
                continue;
            }

            if (action == bytes32("changeStatus")) {
                bytes32 swapId = data[pos]; pos++;
                Status newStatus = data2status(data[pos]); pos++;
                changeStatus(swapId, newStatus);
                continue;
            }
            revert("invalid data");
        }
    }

    function mint(bytes32 swapId, uint amount, address receiver) internal {
        require(swapStatus[swapId] == Status.New, "invalid request status");
        require(Token(tokenAddress).mint(receiver, amount), "invalid mint");
        swapStatus[swapId] = Status.Success;
    }

    function changeStatus(bytes32 swapId, Status newStatus) internal {
        require(swapStatus[swapId] == Status.New, "invalid request status");
        swapStatus[swapId] = newStatus;
    }

    // function createTransferUnwrapRequest() public {
    // }
}

