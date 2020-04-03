pragma solidity >=0.5.0 <0.6.3;

import "./Token.sol";

contract Supersymmetry {
    enum Status {
        None,
        New,
        Rejected,
        Success
    }

    enum Type {
        Burn,
        Mint
    }

    struct Request {
        Status status;
        Type rType;
        address owner;
        uint256 height;
        string target;
        uint256 tokenAmount;
    }

    address[5] public admins;
    address public tokenAddress;
    uint8 public bftCoefficent;

    constructor(address[5] memory newAdmins, uint8 newBftCoefficent) public {
        require(newAdmins[0] != address(0x00), "empty address");
        require(newAdmins[1] != address(0x00), "empty address");
        require(newAdmins[2] != address(0x00), "empty address");
        require(newAdmins[3] != address(0x00), "empty address");
        require(newAdmins[4] != address(0x00), "empty address");

        tokenAddress = address(new Token());
        bftCoefficent = newBftCoefficent;
        admins = newAdmins;
    }

    event NewRequest(bytes32 requestHash, address owner, string target);
    event StatusChanged(bytes32 requestHash, Status status);
    event Mint(bytes32 requestHash, address owner, uint amount);
    event Return(bytes32 requestHash, address owner, uint amount);

    mapping(bytes32 => Request) public requests;

    function createBurnRequest(string memory recipient, uint256 amount) public returns (bytes32) {
        require(Token(tokenAddress).transferFrom(msg.sender, address(this), amount), "invalid balance");
        require(amount > 0, "value less or equals 0");

        bytes32 requestHash = sha256(abi.encodePacked(this, msg.sender, recipient, amount, block.number));
        require(requests[requestHash].status == Status.None, "swap exist");

        requests[requestHash] = Request(
            {status: Status.New, rType: Type.Burn, owner: msg.sender, target: recipient, height: block.number, tokenAmount: amount });
        emit NewRequest(requestHash, msg.sender, recipient);
        return requestHash;
    }

    function createMintRequest(string memory sender, uint256 value) public returns (bytes32) {
        address recipient = msg.sender;
        bytes32 requestHash = sha256(abi.encodePacked(this, sender, recipient, value, block.number));
        require(value > 0, "value less or equals 0");
        require(requests[requestHash].status == Status.None, "swap exist");

        requests[requestHash] = Request(
            {status: Status.New, rType: Type.Mint, owner: recipient, height: block.number, target: sender, tokenAmount: value });
        emit NewRequest(requestHash, recipient, sender);
        return requestHash;
    }

    function changeStatus(bytes32 requestHash, uint8[5] memory v, bytes32[5] memory r, bytes32[5] memory s, uint8 intStatus) public {
        require(requests[requestHash].status == Status.New, "status is now new");
        require(intStatus != uint8(Status.None), "invalid status");
        require(intStatus != uint8(Status.New), "invalid status");

        int count = 0;
        for(uint i = 0; i < 5; i++) {
            count += ecrecover(keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n33", abi.encodePacked(requestHash, intStatus))),
                v[i], r[i], s[i]) == admins[i] ? 1 : 0;
        }

        require(count >= bftCoefficent, "admins vote count is less bftCoefficent");
      
        Status status = Status(intStatus);
        if (intStatus == uint8(Status.Success)) {
            if (requests[requestHash].rType == Type.Mint) {
                require(Token(tokenAddress).mint(requests[requestHash].owner, requests[requestHash].tokenAmount), "invalid balance");
            }
        } else if (intStatus == uint8(Status.Rejected)) {
            if (requests[requestHash].rType == Type.Burn) {
                require(Token(tokenAddress).transferFrom(address(this), requests[requestHash].owner, requests[requestHash].tokenAmount), "invalid balance");
            }
        }
        requests[requestHash].status = status;
        emit StatusChanged(requestHash, status);
    }
}