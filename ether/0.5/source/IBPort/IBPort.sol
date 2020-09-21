pragma solidity ^0.5;

import "../Token/Token.sol";
import "../interfaces/ISubscriberBytes.sol";

contract IBPort is ISubscriberBytes {
    enum Status {
        None,
        New,
        Rejected,
        Success,
        Returned
    }

    struct UnwrapRequest {
        bytes32 receiver;
        uint amount;
    }

    event RequestCreated(uint, address, bytes32, uint);

    address public nebula;
    Token public tokenAddress;

    uint public requestPosition = 1;
    mapping(uint => UnwrapRequest) public unwrapRequests;
    mapping(uint => Status) public swapStatus;

    constructor(address _nebula, address _tokenAddress) public {
        nebula = _nebula;
        tokenAddress = Token(_tokenAddress);
    }

    function deserializeUint(bytes memory b, uint startPos, uint len) internal pure returns (uint) {
        uint v = 0;
        for (uint p = startPos + len - 1; p >= startPos; p--) {
            v = v * 256 + uint(uint8(b[p]));
        }
        return v;
    }

    function deserializeAddress(bytes memory b, uint startPos) internal pure returns (address) {
        return address(uint160(deserializeUint(b, startPos, 20)));
    }

    function deserializeStatus(bytes memory b, uint pos) internal pure returns (Status) {
        uint d = uint(uint8(b[pos]));
        if (d == 0) return Status.None;
        if (d == 1) return Status.New;
        if (d == 2) return Status.Rejected;
        if (d == 3) return Status.Success;
        if (d == 4) return Status.Returned;
        revert("invalid status");
    }

    function attachValue(bytes calldata value) external {
        require(msg.sender == nebula, "access denied");
        for (uint pos = 0; pos < value.length; ) {
            bytes1 action = value[pos]; pos++;

            if (action == bytes1("m")) {
                uint swapId = deserializeUint(value, pos, 32); pos += 32;
                uint amount = deserializeUint(value, pos, 32); pos += 32;
                address receiver = deserializeAddress(value, pos); pos += 20;
                mint(swapId, amount, receiver);
                continue;
            }

            if (action == bytes1("c")) {
                uint swapId = deserializeUint(value, pos, 32); pos += 32;
                Status newStatus = deserializeStatus(value, pos); pos += 1;
                changeStatus(swapId, newStatus);
                continue;
            }
            revert("invalid data");
        }
    }

    function mint(uint swapId, uint amount, address receiver) internal {
        require(swapStatus[swapId] == Status.None, "invalid request status");
        Token(tokenAddress).mint(receiver, amount);
        swapStatus[swapId] = Status.Success;
    }

    function changeStatus(uint swapId, Status newStatus) internal {
        require(swapStatus[swapId] == Status.New, "invalid request status");
        swapStatus[swapId] = newStatus;
    }


    function createTransferUnwrapRequest(uint amount, bytes32 receiver) public {
        unwrapRequests[requestPosition] = UnwrapRequest(receiver, amount);
        swapStatus[requestPosition] = Status.New;
        tokenAddress.burnFrom(msg.sender, amount);
        emit RequestCreated(requestPosition, msg.sender, receiver, amount);
    }
}

