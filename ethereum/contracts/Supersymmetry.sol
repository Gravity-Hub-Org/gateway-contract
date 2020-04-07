pragma solidity >=0.5.16 <=0.6.4;

import "./Token.sol";
import "./Models.sol";

contract Supersymmetry {
    event NewRequest(bytes32 requestHash, address owner, string target);
    event StatusChanged(bytes32 requestHash, Models.Status status);
    event Mint(bytes32 requestHash, address owner, uint amount);
    event Return(bytes32 requestHash, address owner, uint amount);

    uint timeout;
    address[5] public admins;
    uint8 public bftCoefficent;
    mapping(bytes32 => Models.Request) public requests;
    mapping(address => Models.Token) public tokens;

    constructor(address[5] memory newAdmins, uint8 newBftCoefficent, string memory ethAssetId, uint newTimeout) public {
        require(newAdmins[0] != address(0x00), "empty address");
        require(newAdmins[1] != address(0x00), "empty address");
        require(newAdmins[2] != address(0x00), "empty address");
        require(newAdmins[3] != address(0x00), "empty address");
        require(newAdmins[4] != address(0x00), "empty address");

        tokens[address(0x00)] = Models.Token(ethAssetId, Models.TokenType.NativeToken, Models.Status.Success, 18);
        timeout = newTimeout;
        bftCoefficent = newBftCoefficent;
        admins = newAdmins;
    }
    
    function registerNativeToken(address tokenAddress, string memory assetId) public {
        require(tokens[tokenAddress].status == Models.Status.None, "ivalid tokenAddress");
        tokens[tokenAddress] = Models.Token(assetId, Models.TokenType.NativeToken, Models.Status.New, Token(tokenAddress).decimals());
    }

    function registerInputToken(string memory assetId, string memory name, string memory symbol, uint8 decimals) public {
        address tokenAddress = address(new Token(name, symbol, decimals));
        tokens[tokenAddress] = Models.Token(assetId, Models.TokenType.InputToken, Models.Status.New, decimals);
    }

    function changeTokenStatus(address tokenAddress, uint8[5] memory v, bytes32[5] memory r, bytes32[5] memory s, Models.Status status) public {
        require(tokens[tokenAddress].status == Models.Status.New, "ivalid tokenAddress");
        require(status == Models.Status.Success, "ivalid status");
        require(status == Models.Status.Rejected, "ivalid status");
        
        int count = 0;
        for(uint i = 0; i < 5; i++) {
            count += ecrecover(keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n22", abi.encodePacked(tokenAddress, uint8(status)))),
                v[i], r[i], s[i]) == admins[i] ? 1 : 0;
        }
        require(count >= bftCoefficent, "admins vote count is less bftCoefficent");

        tokens[tokenAddress].status = status;
    }

    function withdrawPastDueReqest(bytes32 requestHash) public {
        require(requests[requestHash].status == Models.Status.New, "request not exist");
        require(block.number > requests[requestHash].height + timeout, "rq not expire");

        if (requests[requestHash].tokenAddress == address(0x00)) {
            requests[requestHash].owner.transfer(requests[requestHash].tokenAmount);
        } else {
            require(Token(requests[requestHash].tokenAddress).transferFrom(address(this), requests[requestHash].owner, requests[requestHash].tokenAmount), "invalid balance");
        }
        requests[requestHash].status = Models.Status.Returned;
    }

    function createEthRequest(string memory target, Models.RqType rqType, uint unlockAmount) public payable returns (bytes32) {
        require(rqType == Models.RqType.Lock || rqType == Models.RqType.Unlock, "rq type is not equals token type");

        uint amount = rqType == Models.RqType.Lock ? msg.value : unlockAmount;
        bytes32 requestHash = sha256(abi.encodePacked(this, msg.sender, target, amount, block.number));
        require(requests[requestHash].status == Models.Status.None, "request exist");

        requests[requestHash] = Models.Request(
            {status: Models.Status.New, rType: rqType, owner: msg.sender, target: target, height: block.number, tokenAmount: amount, tokenAddress: address(0x00) });
        emit NewRequest(requestHash, msg.sender, target);
        return requestHash;
    }

    function createTokenRequest(string memory target, Models.RqType rqType,
            uint256 amount, address tokenAddress) public returns (bytes32) {

        require(tokenAddress != address(0x00), "invalid tokenAddress");
        require(tokens[tokenAddress].status == Models.Status.Success, "ivalid tokenAddress");
        require(amount > 0, "value less or equals 0");

        require(tokens[tokenAddress].tokenType == Models.TokenType.NativeToken &&
            (rqType == Models.RqType.Lock || rqType == Models.RqType.Unlock), "rq type is not equals token type");
        require(tokens[tokenAddress].tokenType == Models.TokenType.InputToken &&
            (rqType == Models.RqType.Burn || rqType == Models.RqType.Mint), "rq type is not equals token type");

        require((rqType == Models.RqType.Burn || rqType == Models.RqType.Lock) &&
            Token(tokenAddress).transferFrom(msg.sender, address(this), amount), "invalid balance");


        bytes32 requestHash = sha256(abi.encodePacked(this, msg.sender, target, amount, block.number));
        require(requests[requestHash].status == Models.Status.None, "request exist");

        requests[requestHash] = Models.Request(
            {status: Models.Status.New, rType: rqType, owner: msg.sender, target: target, height: block.number, tokenAmount: amount, tokenAddress: tokenAddress });
        emit NewRequest(requestHash, msg.sender, target);
        return requestHash;
    }

    function changeStatus(bytes32 requestHash, uint8[5] memory v, bytes32[5] memory r, bytes32[5] memory s, Models.Status status) public {
        require(requests[requestHash].status == Models.Status.New, "status is now new");
        require(status == Models.Status.Success || status == Models.Status.Rejected, "invalid status");
        require(block.number <= requests[requestHash].height + timeout, "rq expire");
        
        address tokenAddress = requests[requestHash].tokenAddress;
        int count = 0;
        for(uint i = 0; i < 5; i++) {
            count += ecrecover(keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n33", abi.encodePacked(requestHash, uint8(status)))),
                v[i], r[i], s[i]) == admins[i] ? 1 : 0;
        }

        require(count >= bftCoefficent, "admins vote count is less bftCoefficent");

        if (status == Models.Status.Success) {
            if (requests[requestHash].rType == Models.RqType.Mint) {
                require(Token(tokenAddress).mint(requests[requestHash].owner, requests[requestHash].tokenAmount), "invalid balance");
            } else if (requests[requestHash].rType == Models.RqType.Unlock) {
                if (tokenAddress == address(0x00)) {
                    requests[requestHash].owner.transfer(requests[requestHash].tokenAmount);
                } else {
                    require(Token(tokenAddress).transferFrom(address(this), requests[requestHash].owner, requests[requestHash].tokenAmount), "invalid balance");
                }
            }
        } else if (status == Models.Status.Rejected) {
            if (requests[requestHash].rType == Models.RqType.Burn) {
                require(Token(tokenAddress).transferFrom(address(this), requests[requestHash].owner, requests[requestHash].tokenAmount), "invalid balance");
            } else if (requests[requestHash].rType == Models.RqType.Lock) {
                if (tokenAddress == address(0x00)) {
                    requests[requestHash].owner.transfer(requests[requestHash].tokenAmount);
                } else {
                   require(Token(tokenAddress).transferFrom(address(this), requests[requestHash].owner, requests[requestHash].tokenAmount), "invalid balance");
                }
            }
        }
        requests[requestHash].status = status;
        emit StatusChanged(requestHash, status);
    }
}